package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log/slog"
	v1 "raspyx/internal/delivery/http/v1"
	"raspyx/internal/domain/services"
	"raspyx/internal/repository/postgres"
	"raspyx/internal/usecase"
)

func NewRouter(r *gin.Engine, log *slog.Logger, conn *pgx.Conn) {
	// Routers
	apiV1Group := r.Group("/api/v1")
	{
		v1.NewGroupRouteCreate(
			apiV1Group,
			usecase.NewGroupUseCase(
				postgres.NewGroupRepository(conn),
				*services.NewGroupService(),
			),
			log,
		)

		v1.NewGroupRouteDelete(
			apiV1Group,
			usecase.NewGroupUseCase(
				postgres.NewGroupRepository(conn),
				*services.NewGroupService(),
			),
			log,
		)
	}
}
