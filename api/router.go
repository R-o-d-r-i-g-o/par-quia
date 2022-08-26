package api

import (
	"scheduler/module/home/event"
	"scheduler/module/home/pastor"
	"scheduler/module/home/schedule"
	"scheduler/module/home/user"
	"scheduler/module/section"

	"github.com/gin-gonic/gin"
)

func Avaible(r *gin.Engine) {

	webSite := r.Group("/Scheduler.com")
	{
		// Landing Page
		//<<< space reserved to lading page routers >>>

		// Entry Page
		sections := webSite.Group("/Welcome")
		section.AvaiableRoutes(sections)

		// Home Page
		home := webSite.Group("/Home")
		homePageRouters(home)
	}
}

func homePageRouters(home *gin.RouterGroup) {

	users := home.Group("/Users")
	user.AvaiableRoutes(users)

	events := home.Group("/Events")
	event.AvaiableRoutes(events)

	pastors := home.Group("/Pastors")
	{
		pastor.AvaiableRoutes(pastors)
		pastorInternalRouters(pastors)
	}
}

func pastorInternalRouters(pastors *gin.RouterGroup) {

	schedules := pastors.Group("/Schedule")
	schedule.AvaiableRoutes(schedules)
}
