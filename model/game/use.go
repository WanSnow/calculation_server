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
		case "player-point":
			temp, err := strconv.ParseUint(gameMap[k], 0, 0)
			if err != nil {
				log.Println(fmt.Sprintf("Game: %s has invalid field: %s", gameId, k))
				return nil, err
			}
			game.PlayerPoint = mission.DecodeUint64ToPoint(temp)
		case "player-direction":
			direction := gameMap[k]
			intD, err := strconv.Atoi(direction)
			if err != nil {
				return nil, err
			}
			game.PlayerDirection = Direction(intD)
		case "weapon-direction":
			direction := gameMap[k]
			intD, err := strconv.Atoi(direction)
			if err != nil {
				return nil, err
			}
			game.WeaponDirection = Direction(intD)
		case "bullet":
			temp, err := strconv.ParseUint(vals[1], 0, 0)
			if err != nil {
				log.Println(fmt.Sprintf("Game: %s has invalid field: %s", gameId, k))
				return nil, err
			}
			point := mission.DecodeUint64ToPoint(temp)
			direction := gameMap[k]
			intD, err := strconv.Atoi(direction)
			if err != nil {
				return nil, err
			}
			game.Bullets = append(game.Bullets, &Bullet{
				Point:     point,
				Direction: Direction(intD),
			})
		default:
			log.Println(fmt.Sprintf("Game: %s has invalid field: %s", gameId, k))
			return nil, errors.New("invalid Game field")
		}
	}
	return game, nil
}

func (u *Use) SaveGame(gameId string, game *Game) error {
	point := game.PlayerPoint
	_, err := u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "player-point", mission.EncodePointToUint64(point)).Result()
	if err != nil {
		return err
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "player-direction", fmt.Sprintf("%d", game.PlayerDirection)).Result()
	if err != nil {
		return err
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "weapon-direction", fmt.Sprintf("%d", game.WeaponDirection)).Result()
	if err != nil {
		return err
	}
	for _, v := range game.Points {
		_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), fmt.Sprintf("point_%d", mission.EncodePointToUint64(v)), int64(v.Type)).Result()
		if err != nil {
			return err
		}
	}
	for _, v := range game.Bullets {
		_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), fmt.Sprintf("bullet_%d", mission.EncodePointToUint64(v.Point)), int64(v.Direction)).Result()
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
	point := strconv.FormatUint(mission.EncodePointToUint64(&mission.Point{
		Type: mission.PointType_POINT_TYPE_POSITION,
		X:    1,
		Y:    1,
	}), 10)
	if len(keys) > 0 {
		point = strings.Split(keys[0], "_")[1]
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "player-point", point).Result()
	if err != nil {
		return err
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "player-direction", fmt.Sprintf("%d", Direction_DIRECTION_UP)).Result()
	if err != nil {
		return err
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("game_%s", gameId), "weapon-direction", fmt.Sprintf("%d", Direction_DIRECTION_UP)).Result()
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

func (u *Use) GetMissionId(gameId string) (string, error) {
	return u.redisClient.HGet(fmt.Sprintf("game_%s", gameId), "mission").Result()
}
