package main

import (
	"github.com/wansnow/calculation_server/server/calculation_server"
	"time"
)

func main() {
	cs := calculation_server.NewCalculationServer()
	go cs.Start()
	for {
		time.Sleep(1 * time.Second)
	}
}
