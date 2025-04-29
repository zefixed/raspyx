package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"raspyx/config"
	mw "raspyx/internal/delivery/http/middleware"
	v1 "raspyx/internal/delivery/http/v1"
	"raspyx/internal/domain/services"
	"raspyx/internal/repository/postgres"
	myredis "raspyx/internal/repository/redis"
	"raspyx/internal/usecase"
)

func NewRouter(r *gin.Engine, log *slog.Logger, conn *pgxpool.Pool, redisClient *redis.Client, cfg *config.Config) {
	apiV1GroupUser := r.Group("/raspyx/api/v1")
	apiV1GroupModerator := r.Group("/raspyx/api/v1")
	apiV1GroupModerator.Use(mw.AuthMiddleware(cfg.JWT), mw.AccessLevelMiddleware(50))
	apiV1GroupAdmin := r.Group("/raspyx/api/v1")
	apiV1GroupAdmin.Use(mw.AuthMiddleware(cfg.JWT), mw.AccessLevelMiddleware(99))

	groupUseCase := usecase.NewGroupUseCase(
		postgres.NewGroupRepository(conn),
		*services.NewGroupService(),
	)

	v1.NewGroupRouteCreate(apiV1GroupModerator, groupUseCase, log)
	v1.NewGroupRouteGet(apiV1GroupModerator, groupUseCase, log)
	v1.NewGroupRouteGetByUUID(apiV1GroupModerator, groupUseCase, log)
	v1.NewGroupRouteGetByNumber(apiV1GroupModerator, groupUseCase, log)
	v1.NewGroupRouteUpdate(apiV1GroupModerator, groupUseCase, log)
	v1.NewGroupRouteDelete(apiV1GroupModerator, groupUseCase, log)

	locationUseCase := usecase.NewLocationUseCase(
		postgres.NewLocationRepository(conn),
		*services.NewLocationService(),
	)

	v1.NewLocationRouteCreate(apiV1GroupModerator, locationUseCase, log)
	v1.NewLocationRouteGet(apiV1GroupModerator, locationUseCase, log)
	v1.NewLocationRouteGetByUUID(apiV1GroupModerator, locationUseCase, log)
	v1.NewLocationRouteGetByName(apiV1GroupModerator, locationUseCase, log)
	v1.NewLocationRouteUpdate(apiV1GroupModerator, locationUseCase, log)
	v1.NewLocationRouteDelete(apiV1GroupModerator, locationUseCase, log)

	roomUseCase := usecase.NewRoomUseCase(
		postgres.NewRoomRepository(conn),
		*services.NewRoomService(),
	)

	v1.NewRoomRouteCreate(apiV1GroupModerator, roomUseCase, log)
	v1.NewRoomRouteGet(apiV1GroupModerator, roomUseCase, log)
	v1.NewRoomRouteGetByUUID(apiV1GroupModerator, roomUseCase, log)
	v1.NewRoomRouteGetByNumber(apiV1GroupModerator, roomUseCase, log)
	v1.NewRoomRouteUpdate(apiV1GroupModerator, roomUseCase, log)
	v1.NewRoomRouteDelete(apiV1GroupModerator, roomUseCase, log)

	subjectUseCase := usecase.NewSubjectUseCase(
		postgres.NewSubjectRepository(conn),
		*services.NewSubjectService(),
	)

	v1.NewSubjectRouteCreate(apiV1GroupModerator, subjectUseCase, log)
	v1.NewSubjectRouteGet(apiV1GroupModerator, subjectUseCase, log)
	v1.NewSubjectRouteGetByUUID(apiV1GroupModerator, subjectUseCase, log)
	v1.NewSubjectRouteGetByName(apiV1GroupModerator, subjectUseCase, log)
	v1.NewSubjectRouteUpdate(apiV1GroupModerator, subjectUseCase, log)
	v1.NewSubjectRouteDelete(apiV1GroupModerator, subjectUseCase, log)

	subjectTypeUseCase := usecase.NewSubjectTypeUseCase(
		postgres.NewSubjectTypeRepository(conn),
		*services.NewSubjectTypeService(),
	)

	v1.NewSubjectTypeRouteCreate(apiV1GroupModerator, subjectTypeUseCase, log)
	v1.NewSubjectTypeRouteGet(apiV1GroupModerator, subjectTypeUseCase, log)
	v1.NewSubjectTypeRouteGetByUUID(apiV1GroupModerator, subjectTypeUseCase, log)
	v1.NewSubjectTypeRouteGetByType(apiV1GroupModerator, subjectTypeUseCase, log)
	v1.NewSubjectTypeRouteUpdate(apiV1GroupModerator, subjectTypeUseCase, log)
	v1.NewSubjectTypeRouteDelete(apiV1GroupModerator, subjectTypeUseCase, log)

	teacherUseCase := usecase.NewTeacherUseCase(
		postgres.NewTeacherRepository(conn),
		*services.NewTeacherService(),
	)

	v1.NewTeacherRouteCreate(apiV1GroupModerator, teacherUseCase, log)
	v1.NewTeacherRouteGet(apiV1GroupModerator, teacherUseCase, log)
	v1.NewTeacherRouteGetByUUID(apiV1GroupModerator, teacherUseCase, log)
	v1.NewTeacherRouteGetByFullName(apiV1GroupModerator, teacherUseCase, log)
	v1.NewTeacherRouteUpdate(apiV1GroupModerator, teacherUseCase, log)
	v1.NewTeacherRouteDelete(apiV1GroupModerator, teacherUseCase, log)

	scheduleUseCase := usecase.NewScheduleUseCase(
		postgres.NewScheduleRepository(conn),
		postgres.NewGroupRepository(conn),
		postgres.NewSubjectRepository(conn),
		postgres.NewSubjectTypeRepository(conn),
		postgres.NewLocationRepository(conn),
		postgres.NewTeacherRepository(conn),
		postgres.NewRoomRepository(conn),
		postgres.NewTeachersToScheduleRepository(conn),
		postgres.NewRoomsToScheduleRepository(conn),
		*services.NewScheduleService(),
		myredis.NewRedisCache(redisClient),
	)

	v1.NewScheduleRouteCreate(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGet(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGetByUUID(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGetByTeacher(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGetByTeacherUUID(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGetByGroup(apiV1GroupUser, scheduleUseCase, log)
	v1.NewScheduleRouteGetByGroupUUID(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGetByRoom(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGetByRoomUUID(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGetBySubject(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGetBySubjectUUID(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGetByLocation(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteGetByLocationUUID(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteUpdate(apiV1GroupModerator, scheduleUseCase, log)
	v1.NewScheduleRouteDelete(apiV1GroupModerator, scheduleUseCase, log)

	userUseCase := usecase.NewUserUseCase(
		postgres.NewUserRepository(conn),
		*services.NewUserService(),
	)

	v1.NewUserRouteRegister(apiV1GroupUser, userUseCase, log)
	v1.NewUserRouteLogin(apiV1GroupUser, userUseCase, log, cfg.JWT)
	v1.NewUserRouteGet(apiV1GroupModerator, userUseCase, log)
	v1.NewUserRouteGetByUUID(apiV1GroupModerator, userUseCase, log)
	v1.NewUserRouteGetByUsername(apiV1GroupModerator, userUseCase, log)
	v1.NewUserRouteGetByAccessLevel(apiV1GroupModerator, userUseCase, log)
	v1.NewUserRouteUpdate(apiV1GroupModerator, userUseCase, log)
	v1.NewUserRouteDelete(apiV1GroupModerator, userUseCase, log)
}
