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
	// format2 := "0000-00-00"

	timeTest1, _ := time.Parse(format1, currentTime.Format("2006-01-02"))
	// timeTest2, _ := time.Parse(format2, currentTime.Format("2006-jan-02"))

	fmt.Println(timeTest1)
	// fmt.Println(timeTest2)
	// fmt.Println(currentTime.Format("2006-jan-02"))
	// fmt.Println(currentTime)

	// r := gin.Default()

	// router.Avaible(r)

	// r.Run(fmt.Sprintf("configs.Server.HOST:configs.Server.PORT"))

}

// 01   -> month with zero prefix
// 02   -> day with zero prefix
// 06   -> year (last two digits)
// 15   -> hour (24h based)
// 04   -> minutes with zero prefix
// 05   -> seconds with zero prefix
// 2006 -> long year
