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

	roomUseCase := usecase.NewRoomUseCase(
		postgres.NewRoomRepository(conn),
		*services.NewRoomService(),
	)

	v1.NewRoomRouteCreate(apiV1Group, roomUseCase, log)
	v1.NewRoomRouteGet(apiV1Group, roomUseCase, log)
	v1.NewRoomRouteGetByUUID(apiV1Group, roomUseCase, log)
	v1.NewRoomRouteGetByNumber(apiV1Group, roomUseCase, log)
	v1.NewRoomRouteUpdate(apiV1Group, roomUseCase, log)
	v1.NewRoomRouteDelete(apiV1Group, roomUseCase, log)

	subjectUseCase := usecase.NewSubjectUseCase(
		postgres.NewSubjectRepository(conn),
		*services.NewSubjectService(),
	)

	v1.NewSubjectRouteCreate(apiV1Group, subjectUseCase, log)
	v1.NewSubjectRouteGet(apiV1Group, subjectUseCase, log)
	v1.NewSubjectRouteGetByUUID(apiV1Group, subjectUseCase, log)
	v1.NewSubjectRouteGetByName(apiV1Group, subjectUseCase, log)
	v1.NewSubjectRouteUpdate(apiV1Group, subjectUseCase, log)
	v1.NewSubjectRouteDelete(apiV1Group, subjectUseCase, log)

	subjectTypeUseCase := usecase.NewSubjectTypeUseCase(
		postgres.NewSubjectTypeRepository(conn),
		*services.NewSubjectTypeService(),
	)

	v1.NewSubjectTypeRouteCreate(apiV1Group, subjectTypeUseCase, log)
	v1.NewSubjectTypeRouteGet(apiV1Group, subjectTypeUseCase, log)
	v1.NewSubjectTypeRouteGetByUUID(apiV1Group, subjectTypeUseCase, log)
	v1.NewSubjectTypeRouteGetByType(apiV1Group, subjectTypeUseCase, log)
	v1.NewSubjectTypeRouteUpdate(apiV1Group, subjectTypeUseCase, log)
	v1.NewSubjectTypeRouteDelete(apiV1Group, subjectTypeUseCase, log)
}
