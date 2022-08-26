package schedule

import (
	"github.com/gin-gonic/gin"
)

func AvaiableRoutes(users *gin.RouterGroup) {
	users.GET("/", Find)
	users.POST("/", Create)
	users.PATCH("/", Update)
	users.DELETE("/", Delete)
}
