package routers

import (
	"github.com/gin-gonic/gin"

	"weaving_net/internal/handler"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		workexperiencesRouter(group, handler.NewWorkexperiencesHandler())
	})
}

func workexperiencesRouter(group *gin.RouterGroup, h handler.WorkexperiencesHandler) {
	//group.Use(middleware.Auth()) // all of the following routes use jwt authentication
	// or group.Use(middleware.Auth(middleware.WithVerify(verify))) // token authentication

	group.POST("/workexperiences", h.Create)
	group.DELETE("/workexperiences/:id", h.DeleteByID)
	group.POST("/workexperiences/delete/ids", h.DeleteByIDs)
	group.PUT("/workexperiences/:id", h.UpdateByID)
	group.GET("/workexperiences/:id", h.GetByID)
	group.POST("/workexperiences/condition", h.GetByCondition)
	group.POST("/workexperiences/list/ids", h.ListByIDs)
	group.GET("/workexperiences/list", h.ListByLastID)
	group.POST("/workexperiences/list", h.List)
}
