package main

import (
	"fmt"
	"time"

	"scheduler/configs"
	// "time"
	// "scheduler/configs"
	// "scheduler/router"
	// "github.com/gin-gonic/gin"
)

func init() {
	configs.Load()
}

func main() {

	// getEnv := func(key string) {
	//     val, ok := os.LookupEnv(key)
	//     if !ok {
	//         fmt.Printf("%s not set\n", key)
	//     } else {
	//         fmt.Printf("%s=%s\n", key, val)
	//     }
	// }

	// getEnv("EDITOR")
	// getEnv("SHELL")

	fmt.Println(configs.DBase)

	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-1-2"))

	// r := gin.Default()

	// router.Avaible(r)

	// r.Run(fmt.Sprintf("%s:%s", configs.Server.HOST, configs.Server.PORT))

}
