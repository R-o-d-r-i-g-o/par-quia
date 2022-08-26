package user

import (
	"github.com/gin-gonic/gin"
)

func AvaiableRoutes(users *gin.RouterGroup) {
	users.GET("/Log-in", Find)
	users.POST("/Sign-up", Create)
	users.PATCH("/Profile", Update)
	users.DELETE("/Profile", Delete)
}
