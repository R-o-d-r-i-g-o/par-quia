package pastor

import (
	"github.com/gin-gonic/gin"
)

func AvaiableRoutes(pastors *gin.RouterGroup) {
	pastors.GET("/List", Find)
	pastors.POST("/Create", Create)
	pastors.PATCH("/{}", Update)
	pastors.DELETE("/{}", Delete)
}
