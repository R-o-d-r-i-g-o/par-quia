package section

import (
	"github.com/gin-gonic/gin"

	"scheduler/core/user"
)

func AvaiableRoutes(users *gin.RouterGroup) {
	users.GET("/Log-in", user.Find)
	users.POST("/Sign-up", user.Create)
}
