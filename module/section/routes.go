package section

import (
	"github.com/gin-gonic/gin"

	"scheduler/module/user/controller"
)

func AvaiableRoutes(users *gin.RouterGroup) {
	users.GET("/Log-in", controller.Find)
	users.POST("/Sign-up", controller.Create)
}
