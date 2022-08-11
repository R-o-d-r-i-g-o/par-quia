package main

import (
	"scheduler/configs"
	"scheduler/database"
	"scheduler/database/model"
	// "time"
	// "scheduler/configs"
	// "scheduler/router"
	// "github.com/gin-gonic/gin"
)

func init() {
	configs.Load()
}

func main() {

	configs.Load()
	database.StartDatabase()
	model.Handler(database.GetGormDB())

	// fmt.Println(configs.DBase)

	// currentTime := time.Now()
	// fmt.Println(currentTime.Format("2006-1-2"))

	// r := gin.Default()

	// router.Avaible(r)

	// r.Run(fmt.Sprintf("%s:%s", configs.Server.HOST, configs.Server.PORT))

}
