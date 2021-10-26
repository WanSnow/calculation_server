package heart_beat_client

import (
	"context"
	"github.com/wansnow/calculation_server/service/heart_beat"
	"google.golang.org/grpc"
	"log"
	"time"
)

type HeartBeatClient struct {
	heart_beat.HeartBeatClient
}

func (c *HeartBeatClient) Start() {
	for {
		// Contact the server and print out its response.
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		r, err := c.HeartBeat(ctx, &heart_beat.HeartBeatMes_Req{
			Ip:   "127.0.0.1",
			Port: 6969,
		})
		if err != nil {
			log.Fatalf("could not send heart beat: %v", err)
		}
		log.Printf("status: %s", r.GetStatus())
		time.Sleep(3 * time.Second)
	}
}

func NewClient(conn grpc.ClientConnInterface) *HeartBeatClient {
	c := &HeartBeatClient{
		heart_beat.NewHeartBeatClient(conn),
	}
	return c
}
