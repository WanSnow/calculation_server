package game

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/wansnow/calculation_server/middleware/redis_client"
	"github.com/wansnow/calculation_server/model/mission"
	"log"
	"strconv"
	"strings"
)

type Use struct {
	redisClient *redis.Client
}

func NewUse() *Use {
	return &Use{
		redisClient: redis_client.RedisClient,
	}
}

func (u *Use) GetGame(gameId string) (*Game, error) {
	game := new(Game)
	game.GameId = gameId
	gameMap, err := u.redisClient.HGetAll(fmt.Sprintf("game_%s", gameId)).Result()
	if err != nil {
		return nil, err
	}
	for k := range gameMap {
		vals := strings.Split(k, "_")
		switch vals[0] {
		case "point":
			temp, err := strconv.ParseUint(vals[1], 0, 0)
			if err != nil {
				log.Println(fmt.Sprintf("Game: %s has invalid field: %s", gameId, k))
				return nil, err
			}
			point := mission.DecodeUint64ToPoint(temp)
			game.Points = append(game.Points, point)
		case "condition":
			condition := gameMap[k]
			intC, err2 := strconv.Atoi(condition)
			if err2 != nil {
				return nil, err2
			}
			game.Condition = mission.WinnerCondition(intC)
		case "mission":
			game.MissionId = gameMap[k]
		default:
			log.Println(fmt.Sprintf("Game: %s has invalid field: %s", gameId, k))
			return nil, errors.New("invalid Game field")
		}
	}
	return game, nil
}

func (u *Use) InitGame(gameId string, m *mission.Mission) error {
	_, err := u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "condition", int64(m.Condition)).Result()
	if err != nil {
		return err
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "mission", m.Id).Result()
	if err != nil {
		return err
	}
	for _, v := range m.Points {
		_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), fmt.Sprintf("point_%d", mission.EncodePointToUint64(v)), int64(v.Type)).Result()
		if err != nil {
			return err
		}
	}
	return nil
}
