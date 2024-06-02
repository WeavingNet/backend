package routers

import (
	"github.com/gin-gonic/gin"

	"weaving_net/internal/handler"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		skillsRouter(group, handler.NewSkillsHandler())
	})
}

func skillsRouter(group *gin.RouterGroup, h handler.SkillsHandler) {
	//group.Use(middleware.Auth()) // all of the following routes use jwt authentication
	// or group.Use(middleware.Auth(middleware.WithVerify(verify))) // token authentication

	group.POST("/skills", h.Create)
	group.DELETE("/skills/:id", h.DeleteByID)
	group.POST("/skills/delete/ids", h.DeleteByIDs)
	group.PUT("/skills/:id", h.UpdateByID)
	group.GET("/skills/:id", h.GetByID)
	group.POST("/skills/condition", h.GetByCondition)
	group.POST("/skills/list/ids", h.ListByIDs)
	group.GET("/skills/list", h.ListByLastID)
	group.POST("/skills/list", h.List)
}
