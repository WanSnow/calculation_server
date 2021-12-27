package heart_beat_client

import (
	"context"
	"github.com/wansnow/calculation_server/config"
	"github.com/wansnow/calculation_server/model/mission"
	"github.com/wansnow/calculation_server/service/heart_beat"
	"google.golang.org/grpc"
	"log"
	"time"
)

type HeartBeatClient struct {
	heart_beat.HeartBeatClient
}

func (c *HeartBeatClient) Start() {
	use := mission.NewUse()
	for {
		ids, err := use.GetAllMissionId()
		if err != nil {
			return
		}
		versionMap := make(map[string]int64)
		for _, id := range ids {
			v, err := use.GetVersion(id)
			if err != nil {
				log.Fatalln(err)
			}
			versionMap[id] = v
		}

		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		r, err := c.HeartBeat(ctx, &heart_beat.HeartBeatMes_Req{
			Ip:              "127.0.0.1",
			Port:            config.ServerC.AcceptServerPort,
			CheckPort:       config.ServerC.StartGameServerPort,
			MissionVersions: versionMap,
		})
		if err != nil {
			log.Fatalf("could not send heart beat: %v", err)
		}

		//log.Printf("res: %v", r.Missions)

		for _, m := range r.Missions {
			err := use.SetMission(m)
			if err != nil {
				log.Fatalln(err)
			}
		}
		//log.Printf("status: %s", r.Status)
		time.Sleep(3 * time.Second)
	}
}

func NewClient(conn grpc.ClientConnInterface) *HeartBeatClient {
	c := &HeartBeatClient{
		heart_beat.NewHeartBeatClient(conn),
	}
	return c
}
