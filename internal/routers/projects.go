package routers

import (
	"github.com/gin-gonic/gin"

	"weaving_net/internal/handler"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		projectsRouter(group, handler.NewProjectsHandler())
	})
}

func projectsRouter(group *gin.RouterGroup, h handler.ProjectsHandler) {
	//group.Use(middleware.Auth()) // all of the following routes use jwt authentication
	// or group.Use(middleware.Auth(middleware.WithVerify(verify))) // token authentication

	group.POST("/projects", h.Create)
	group.DELETE("/projects/:id", h.DeleteByID)
	group.POST("/projects/delete/ids", h.DeleteByIDs)
	group.PUT("/projects/:id", h.UpdateByID)
	group.GET("/projects/:id", h.GetByID)
	group.POST("/projects/condition", h.GetByCondition)
	group.POST("/projects/list/ids", h.ListByIDs)
	group.GET("/projects/list", h.ListByLastID)
	group.POST("/projects/list", h.List)
}
