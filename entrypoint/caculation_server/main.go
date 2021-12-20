package main

import (
	"fmt"
	"github.com/wansnow/calculation_server/config"
	"github.com/wansnow/calculation_server/middleware/redis_client"
	"github.com/wansnow/calculation_server/server/calculation_server"
	"github.com/wansnow/calculation_server/server/game_accept_server"
	"github.com/wansnow/calculation_server/server/game_pub_server"
	"os"
	"os/signal"
	"syscall"
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
	redis_client.InitRedisClient()

	cs := calculation_server.NewCalculationServer()
	go cs.Start()
	go game_accept_server.StartAcceptGameServer()
	go game_pub_server.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
