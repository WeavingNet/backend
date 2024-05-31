package routers

import (
	"github.com/gin-gonic/gin"

	"weaving_net/internal/handler"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		usersRouter(group, handler.NewUsersHandler())
	})
}

func usersRouter(group *gin.RouterGroup, h handler.UsersHandler) {
	//group.Use(middleware.Auth()) // all of the following routes use jwt authentication
	// or group.Use(middleware.Auth(middleware.WithVerify(verify))) // token authentication

	group.POST("/users", h.Create)
	group.DELETE("/users/:id", h.DeleteByID)
	group.POST("/users/delete/ids", h.DeleteByIDs)
	group.PUT("/users/:id", h.UpdateByID)
	group.GET("/users/:id", h.GetByID)
	group.POST("/users/condition", h.GetByCondition)
	group.POST("/users/list/ids", h.ListByIDs)
	group.GET("/users/list", h.ListByLastID)
	group.POST("/users/list", h.List)
}
