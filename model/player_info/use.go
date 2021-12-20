package player_info

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/wansnow/calculation_server/middleware/redis_client"
	"github.com/wansnow/calculation_server/server/calculation_server/model/trigger_msg"
	"github.com/wansnow/calculation_server/server/calculation_server/service/common"
	"github.com/wansnow/calculation_server/server/calculation_server/service/logic_service"
)

type Use struct {
	redisClient *redis.Client
}

func NewUse() *Use {
	return &Use{
		redisClient: redis_client.RedisClient,
	}
}

func (u *Use) SetPlayer(playerInfo *PlayerInfo) {
	common.MainLogicMap[playerInfo.Id] = playerInfo.MainLogicFunc
	common.TriggerMap[fmt.Sprintf("%s_%d", playerInfo.Id, trigger_msg.ATTACKED)] = &logic_service.Attacked{
		AttackedFunc: playerInfo.AttackedFunc,
	}
	common.TriggerMap[fmt.Sprintf("%s_%d", playerInfo.Id, trigger_msg.CRASH)] = &logic_service.Crash{
		CrashFunc: playerInfo.CrashFunc,
	}
	common.TriggerMap[fmt.Sprintf("%s_%d", playerInfo.Id, trigger_msg.DISCOVER)] = &logic_service.Discover{
		DiscoverFunc: playerInfo.DiscoverFunc,
	}
}
