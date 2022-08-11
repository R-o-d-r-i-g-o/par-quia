package user

import (
	"github.com/gin-gonic/gin"

	"scheduler/module/user/controller"
)

func AvaiableRoutes(users *gin.RouterGroup) {
	users.GET("/", controller.Find)
	users.POST("/", controller.Create)
	users.PATCH("/", controller.Update)
	users.DELETE("/", controller.Delete)
}
