package calculation_server

import (
	"fmt"
	"github.com/wansnow/calculation_server/client/heart_beat_client"
	"github.com/wansnow/calculation_server/config"
	"google.golang.org/grpc"
	"log"
)

type CalculationServer struct {
	HBClient *heart_beat_client.HeartBeatClient
}

func (cs *CalculationServer) Start() {
	go cs.HBClient.Start()
}

func NewCalculationServer() *CalculationServer {
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", config.ServerC.HeartBeatServerPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	cs := &CalculationServer{
		HBClient: heart_beat_client.NewClient(conn),
	}
	return cs
}
