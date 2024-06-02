package routers

import (
	"github.com/gin-gonic/gin"

	"weaving_net/internal/handler"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		educationsRouter(group, handler.NewEducationsHandler())
	})
}

func educationsRouter(group *gin.RouterGroup, h handler.EducationsHandler) {
	//group.Use(middleware.Auth()) // all of the following routes use jwt authentication
	// or group.Use(middleware.Auth(middleware.WithVerify(verify))) // token authentication

	group.POST("/educations", h.Create)
	group.DELETE("/educations/:id", h.DeleteByID)
	group.POST("/educations/delete/ids", h.DeleteByIDs)
	group.PUT("/educations/:id", h.UpdateByID)
	group.GET("/educations/:id", h.GetByID)
	group.POST("/educations/condition", h.GetByCondition)
	group.POST("/educations/list/ids", h.ListByIDs)
	group.GET("/educations/list", h.ListByLastID)
	group.POST("/educations/list", h.List)
}
