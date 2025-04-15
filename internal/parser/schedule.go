package parser

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"io"
	"log/slog"
	"net/http"
	"raspyx/config"
	"raspyx/internal/domain/services"
	"raspyx/internal/dto"
	"raspyx/internal/repository"
	"raspyx/internal/repository/postgres"
	"raspyx/internal/usecase"
	"regexp"
	"strings"
	"time"
)

type ScheduleParser struct {
	client    *http.Client
	conn      *pgx.Conn
	log       *slog.Logger
	cfg       config.Parser
	groupRepo *postgres.GroupRepository
	groupSVC  *services.GroupService
}

func NewScheduleParser(timeout time.Duration, conn *pgx.Conn, log *slog.Logger, cfg config.Parser) *ScheduleParser {
	return &ScheduleParser{
		client: &http.Client{Timeout: timeout},
		conn:   conn,
		log:    log,
		cfg:    cfg,
	}
}

func (p *ScheduleParser) New(ctx context.Context) {
	p.log = p.log.With(slog.String("module", "ScheduleParser"))

	p.groupRepo = postgres.NewGroupRepository(p.conn)
	p.groupSVC = services.NewGroupService()

	// Init parsing schedule
	err := p.parse(ctx, p.conn)
	if err != nil {
		p.log.Error(fmt.Sprintf("error parsing schedule: %v", err))
	}

	// Set timeout to 1 minute if it too small
	if p.cfg.Timeout < 1 {
		p.cfg.Timeout = 1
	}

	// Ticker for parsing schedule
	ticker := time.NewTicker(time.Duration(p.cfg.Timeout) * time.Minute)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				p.log.Error(fmt.Sprintf("cancel schedule parser"))
				return
			case <-ticker.C:
				err := p.parse(ctx, p.conn)
				if err != nil {
					p.log.Error(fmt.Sprintf("error parsing schedule: %v", err))
				}
			}
		}
	}()

	<-ctx.Done()
}

func (p *ScheduleParser) parse(ctx context.Context, conn *pgx.Conn) error {
	// Parsing groups
	groups, err := p.parseGroups(ctx)
	if err != nil {
		return err
	}

	// Adding groups to db
	p.addGroupsToDB(ctx, groups)

	return nil
}

func (p *ScheduleParser) parseGroups(ctx context.Context) ([]string, error) {
	// New request to rasp.dmami.ru
	req, err := http.NewRequestWithContext(ctx, "GET", "https://rasp.dmami.ru/", nil)
	if err != nil {
		return nil, err
	}

	// Set referer
	req.Header.Set("Referer", "https://rasp.dmami.ru/")

	// Sending request
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Reading response
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Collect groups from response
	re := regexp.MustCompile(`\d{2}[0-9a-zA-Zа-яА-Я]-\d{3}(\s{1}[a-zA-Zа-яА-я]{3})?`)
	matches := re.FindAll(raw, -1)

	// Deleting repeats from groups
	gm := make(map[string]int)
	for _, m := range matches {
		if _, ok := gm[string(m)]; !ok {
			gm[string(m)] = 1
		}
	}

	// From map of groups to []string
	groups := make([]string, 0, len(gm))
	for group, _ := range gm {
		groups = append(groups, group)
	}

	return groups, nil
}

func (p *ScheduleParser) addGroupsToDB(ctx context.Context, groups []string) {
	groupUC := usecase.NewGroupUseCase(p.groupRepo, *p.groupSVC)
	addedGroups := 0
	for _, group := range groups {
		// Adding group to db
		_, err := groupUC.Create(ctx, &dto.CreateGroupRequest{Group: group})

		// If error != group exist
		if err != nil {
			if !strings.Contains(err.Error(), repository.ErrExist.Error()) {
				p.log.Error(fmt.Sprintf("error adding group %v to db: %v", group, err))
			}
		} else {
			addedGroups++
		}
	}

	p.log.Info(fmt.Sprintf("added %v new groups to db", addedGroups))
}
