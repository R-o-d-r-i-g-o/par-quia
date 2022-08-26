package event

import (
	"github.com/gin-gonic/gin"
)

func AvaiableRoutes(events *gin.RouterGroup) {
	events.GET("/", Find)
	events.POST("/", Create)
	events.PATCH("/", Update)
	events.DELETE("/", Delete)
}
