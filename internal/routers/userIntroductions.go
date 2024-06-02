package routers

import (
	"github.com/gin-gonic/gin"

	"weaving_net/internal/handler"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		userIntroductionsRouter(group, handler.NewUserIntroductionsHandler())
	})
}

func userIntroductionsRouter(group *gin.RouterGroup, h handler.UserIntroductionsHandler) {
	//group.Use(middleware.Auth()) // all of the following routes use jwt authentication
	// or group.Use(middleware.Auth(middleware.WithVerify(verify))) // token authentication

	group.POST("/userIntroductions", h.Create)
	group.DELETE("/userIntroductions/:id", h.DeleteByID)
	group.POST("/userIntroductions/delete/ids", h.DeleteByIDs)
	group.PUT("/userIntroductions/:id", h.UpdateByID)
	group.GET("/userIntroductions/:id", h.GetByID)
	group.POST("/userIntroductions/condition", h.GetByCondition)
	group.POST("/userIntroductions/list/ids", h.ListByIDs)
	group.GET("/userIntroductions/list", h.ListByLastID)
	group.POST("/userIntroductions/list", h.List)
}
