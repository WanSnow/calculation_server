package calculation_server

import (
	"github.com/wansnow/calculation_server/client/heart_beat_client"
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
	conn, err := grpc.Dial("127.0.0.1:9696", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	cs := &CalculationServer{
		HBClient: heart_beat_client.NewClient(conn),
	}
	return cs
}
