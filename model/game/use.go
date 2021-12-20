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
			intC, err := strconv.Atoi(condition)
			if err != nil {
				return nil, err
			}
			game.Condition = mission.WinnerCondition(intC)
		case "mission":
			game.MissionId = gameMap[k]
		case "player_point":
			temp, err := strconv.ParseUint(gameMap[k], 0, 0)
			if err != nil {
				log.Println(fmt.Sprintf("Game: %s has invalid field: %s", gameId, k))
				return nil, err
			}
			game.PlayerPoint = mission.DecodeUint64ToPoint(temp)
		case "player_direction":
			direction := gameMap[k]
			intD, err := strconv.Atoi(direction)
			if err != nil {
				return nil, err
			}
			game.PlayerDirection = Direction(intD)
		case "weapon_direction":
			direction := gameMap[k]
			intD, err := strconv.Atoi(direction)
			if err != nil {
				return nil, err
			}
			game.WeaponDirection = Direction(intD)
		default:
			log.Println(fmt.Sprintf("Game: %s has invalid field: %s", gameId, k))
			return nil, errors.New("invalid Game field")
		}
	}
	return game, nil
}

func (u *Use) SaveGame(gameId string, game *Game) error {
	point := game.PlayerPoint
	_, err := u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "player_point", mission.EncodePointToUint64(point)).Result()
	if err != nil {
		return err
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "player_direction", Direction_DIRECTION_UP).Result()
	if err != nil {
		return err
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "weapon_direction", Direction_DIRECTION_UP).Result()
	if err != nil {
		return err
	}
	for _, v := range game.Points {
		_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), fmt.Sprintf("point_%d", mission.EncodePointToUint64(v)), int64(v.Type)).Result()
		if err != nil {
			return err
		}
	}
	return nil
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
	keys, _, err := u.redisClient.HScan(fmt.Sprintf("mission_%s", m.Id), 0, "position_*", 1).Result()
	if err != nil {
		return err
	}
	point := strings.Split(keys[0], "_")[1]
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "player_point", point).Result()
	if err != nil {
		return err
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "player_direction", Direction_DIRECTION_UP).Result()
	if err != nil {
		return err
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "weapon_direction", Direction_DIRECTION_UP).Result()
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
