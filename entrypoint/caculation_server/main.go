package main

import (
	"fmt"
	"github.com/wansnow/calculation_server/config"
	"github.com/wansnow/calculation_server/middleware/redis"
	"github.com/wansnow/calculation_server/server/calculation_server"
	"os"
	"time"
)

func main() {
	projectPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = config.LoadConfig(projectPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	redis.InitRedisClient()

	cs := calculation_server.NewCalculationServer()
	go cs.Start()
	for {
		time.Sleep(1 * time.Second)
	}
}
