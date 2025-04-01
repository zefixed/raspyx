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
	apiV1Group := r.Group("/api/v1")

	groupUseCase := usecase.NewGroupUseCase(
		postgres.NewGroupRepository(conn),
		*services.NewGroupService(),
	)

	v1.NewGroupRouteCreate(apiV1Group, groupUseCase, log)
	v1.NewGroupRouteGet(apiV1Group, groupUseCase, log)
	v1.NewGroupRouteGetByUUID(apiV1Group, groupUseCase, log)
	v1.NewGroupRouteGetByNumber(apiV1Group, groupUseCase, log)
	v1.NewGroupRouteUpdate(apiV1Group, groupUseCase, log)
	v1.NewGroupRouteDelete(apiV1Group, groupUseCase, log)

	locationUseCase := usecase.NewLocationUseCase(
		postgres.NewLocationRepository(conn),
		*services.NewLocationService(),
	)

	v1.NewLocationRouteCreate(apiV1Group, locationUseCase, log)
	v1.NewLocationRouteGet(apiV1Group, locationUseCase, log)
	v1.NewLocationRouteGetByUUID(apiV1Group, locationUseCase, log)
	v1.NewLocationRouteGetByName(apiV1Group, locationUseCase, log)
	v1.NewLocationRouteUpdate(apiV1Group, locationUseCase, log)
	v1.NewLocationRouteDelete(apiV1Group, locationUseCase, log)
}
