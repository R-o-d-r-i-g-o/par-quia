package main

import (
	"fmt"
	"scheduler/configs"
	"time"
	// "time"
	// "scheduler/configs"
	// "scheduler/router"
	// "github.com/gin-gonic/gin"
)

func init() {
	configs.Load()
}

func main() {
	// database.StartDatabase()
	// model.Handler(database.GetGormDB())

	fmt.Println(configs.Time)

	currentTime := time.Now()

	format1 := "2006-01-02"
	format2 := "2006-jan-02"

	timeTest1, _ := time.Parse(format1, currentTime.Format("2006-01-02"))
	timeTest2, _ := time.Parse(format2, currentTime.Format("2006-jan-02"))

	fmt.Println(timeTest1)
	fmt.Println(timeTest2)
	fmt.Println(time.Hour.String())
	fmt.Println(currentTime)

	// r := gin.Default()

	// router.Avaible(r)

	// r.Run(fmt.Sprintf("%s:%s", configs.Server.HOST, configs.Server.PORT))

}
