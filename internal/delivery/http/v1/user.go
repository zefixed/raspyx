package v1

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"raspyx/config"
	"raspyx/internal/dto"
	"raspyx/internal/usecase"
)

type userRoutes struct {
	uc  *usecase.UserUseCase
	log *slog.Logger
}

// NewUserRouteRegister
// @Summary Creating a new user
// @Description Creates a new user in the database and returns its uuid
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.RegisterUserRequest true "User"
// @Success 200 {object} ResponseOK{response=dto.RegisterUserRequest}
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/users/register [post]
func NewUserRouteRegister(apiV1Group *gin.RouterGroup, uc *usecase.UserUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewUserRouteRegister"
	log = log.With(slog.String("op", op))

	r := &userRoutes{uc, log}

	userGroup := apiV1Group.Group("/users")

	userGroup.POST("/register", func(c *gin.Context) {
		var userDTO dto.RegisterUserRequest
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Create(c, &userDTO)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err: err,
				c:   c,
				log: log,
				// Potential leakage of credentials to logs
				logKey:   "user_dto",
				logValue: userDTO,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewUserRouteLogin
// @Summary User authentication
// @Description Authenticate user and return access token
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.LoginUserRequest true "User"
// @Success 200 {object} ResponseOK{response=dto.LoginUserRequest}
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/users/login [post]
func NewUserRouteLogin(apiV1Group *gin.RouterGroup, uc *usecase.UserUseCase, log *slog.Logger, jwt config.JWT) {
	const op = "delivery.http.v1.NewUserRouteLogin"
	log = log.With(slog.String("op", op))

	r := &userRoutes{uc, log}

	userGroup := apiV1Group.Group("/users")

	userGroup.POST("/login", func(c *gin.Context) {
		var userDTO dto.LoginUserRequest
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		resp, err := r.uc.Login(c, jwt, &userDTO)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "user_name",
				logValue: userDTO.Username,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})

}

// NewUserRouteGet
// @Summary Getting users
// @Description Get all users from database
// @Security ApiKeyAuth
// @Tags user
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseOK{response=dto.GetUsersResponse}
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/users/ [get]
func NewUserRouteGet(apiV1Group *gin.RouterGroup, uc *usecase.UserUseCase, log *slog.Logger) {
	const op = "delivery.http.v1.NewUserRouteGet"
	log = log.With(slog.String("op", op))

	r := &userRoutes{uc, log}

	userGroup := apiV1Group.Group("/users")

	userGroup.GET("/", func(c *gin.Context) {
		resp, err := r.uc.Get(c)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "error",
				logValue: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewUserRouteGetByUUID
// @Summary Getting user by uuid
// @Description Get user from database with given uuid
// @Security ApiKeyAuth
// @Tags user
// @Accept */*
// @Produce json
// @Param uuid path string true "User uuid"
// @Success 200 {object} ResponseOK{response=models.User}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/users/uuid/{uuid} [get]
func NewUserRouteGetByUUID(apiV1Group *gin.RouterGroup, uc *usecase.UserUseCase, log *slog.Logger) {
	r := &userRoutes{uc, log}

	userGroup := apiV1Group.Group("/users")

	userGroup.GET("/uuid/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		resp, err := r.uc.GetByUUID(c, reqUUID)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "user_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewUserRouteGetByUsername
// @Summary Getting user by username
// @Description Get user from database with given username
// @Security ApiKeyAuth
// @Tags user
// @Accept */*
// @Produce json
// @Param username path string true "User username"
// @Success 200 {object} ResponseOK{response=models.User}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/users/username/{username} [get]
func NewUserRouteGetByUsername(apiV1Group *gin.RouterGroup, uc *usecase.UserUseCase, log *slog.Logger) {
	r := &userRoutes{uc, log}

	userGroup := apiV1Group.Group("/users")

	userGroup.GET("/username/:username", func(c *gin.Context) {
		reqUsername := c.Param("username")
		resp, err := r.uc.GetByUsername(c, reqUsername)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "user_username",
				logValue: reqUsername,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewUserRouteGetByAccessLevel
// @Summary Getting user by access level
// @Description Get user from database with access level less than or equal to given AccessLevel
// @Security ApiKeyAuth
// @Tags user
// @Accept */*
// @Produce json
// @Param al path string true "Access level"
// @Success 200 {object} ResponseOK{response=[]models.User}
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/users/al/{al} [get]
func NewUserRouteGetByAccessLevel(apiV1Group *gin.RouterGroup, uc *usecase.UserUseCase, log *slog.Logger) {
	r := &userRoutes{uc, log}

	userGroup := apiV1Group.Group("/users")

	userGroup.GET("/al/:al", func(c *gin.Context) {
		reqAccessLevel := c.Param("al")
		resp, err := r.uc.GetByAccessLevel(c, reqAccessLevel)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "user_access_level",
				logValue: reqAccessLevel,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(resp))
	})
}

// NewUserRouteUpdate
// @Summary Updating user
// @Description Update user in database
// @Description 0 - user, 50 - moderator, 99 - admin
// @Security ApiKeyAuth
// @Tags user
// @Accept json
// @Produce json
// @Param uuid path string true "User uuid"
// @Param user body dto.UpdateUserRequest true "User"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/users/{uuid} [put]
func NewUserRouteUpdate(apiV1Group *gin.RouterGroup, uc *usecase.UserUseCase, log *slog.Logger) {
	r := &userRoutes{uc, log}

	userGroup := apiV1Group.Group("/users")

	userGroup.PUT("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")

		var userDTO dto.UpdateUserRequest
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			log.Warn(ErrWrongDataStructure, slog.String("error", err.Error()))
			c.JSON(http.StatusBadRequest, RespError(ErrWrongDataStructure))
			return
		}

		err := r.uc.Update(c, reqUUID, &userDTO)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "user",
				logValue: map[string]any{"uuid": reqUUID, "user_dto": userDTO},
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}

// NewUserRouteDelete
// @Summary Deleting existing user
// @Description Deleting existing user from the database
// @Security ApiKeyAuth
// @Tags user
// @Accept */*
// @Produce json
// @Param uuid path string true "User uuid"
// @Success 200 {object} ResponseOK
// @Failure 400 {object} ResponseError
// @Failure 401 {object} ResponseError
// @Failure 403 {object} ResponseError
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/users/{uuid} [delete]
func NewUserRouteDelete(apiV1Group *gin.RouterGroup, uc *usecase.UserUseCase, log *slog.Logger) {
	r := &userRoutes{uc, log}

	userGroup := apiV1Group.Group("/users")

	userGroup.DELETE("/:uuid", func(c *gin.Context) {
		reqUUID := c.Param("uuid")
		err := r.uc.Delete(c, reqUUID)
		if err != nil {
			makeErrResponse(c, &ErrResp{
				err:      err,
				c:        c,
				log:      log,
				logKey:   "user_uuid",
				logValue: reqUUID,
			})
			return
		}

		c.JSON(http.StatusOK, RespOK(nil))
	})
}
