package pastor

import (
	"github.com/gin-gonic/gin"

	"scheduler/module/pastor/controller"
)

func AvaiableRoutes(pastors *gin.RouterGroup) {
	pastors.GET("/", controller.Find)
	pastors.POST("/", controller.Create)
	pastors.PATCH("/", controller.Update)
	pastors.DELETE("/", controller.Delete)
}
