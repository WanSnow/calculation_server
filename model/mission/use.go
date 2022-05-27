package mission

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/wansnow/calculation_server/middleware/redis_client"
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

func (u *Use) GetAllMissionId() ([]string, error) {
	var result []string
	ids, err := u.redisClient.Keys("mission_*").Result()
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		result = append(result, strings.Split(id, "_")[1])
	}
	return result, nil
}

func (u *Use) GetMission(missionId string) (*Mission, error) {
	mission := new(Mission)
	mission.Id = missionId
	missionMap, err := u.redisClient.HGetAll(fmt.Sprintf("mission_%s", missionId)).Result()
	if err != nil {
		return nil, err
	}
	for k := range missionMap {
		vals := strings.Split(k, "_")
		switch vals[0] {
		case "block":
			temp, err := strconv.ParseUint(vals[1], 0, 0)
			if err != nil {
				log.Println(fmt.Sprintf("Mission: %s has invalid field: %s", missionId, k))
				return nil, err
			}
			point := DecodeUint64ToPoint(temp)
			point.Type = PointType_POINT_TYPE_BLOCK
			mission.Points = append(mission.Points, point)
		case "position":
			temp, err := strconv.ParseUint(vals[1], 0, 0)
			if err != nil {
				log.Println(fmt.Sprintf("Mission: %s has invalid field: %s", missionId, k))
				return nil, err
			}
			point := DecodeUint64ToPoint(temp)
			point.Type = PointType_POINT_TYPE_POSITION
			mission.Points = append(mission.Points, point)
		case "target":
			temp, err := strconv.ParseUint(vals[1], 0, 0)
			if err != nil {
				log.Println(fmt.Sprintf("Mission: %s has invalid field: %s", missionId, k))
				return nil, err
			}
			point := DecodeUint64ToPoint(temp)
			point.Type = PointType_POINT_TYPE_TARGET
			mission.Points = append(mission.Points, point)
		case "enemy":
			temp, err := strconv.ParseUint(vals[1], 0, 0)
			if err != nil {
				log.Println(fmt.Sprintf("Mission: %s has invalid field: %s", missionId, k))
				return nil, err
			}
			point := DecodeUint64ToPoint(temp)
			point.Type = PointType_POINT_TYPE_ENEMY
			mission.Points = append(mission.Points, point)
		case "version":
			version, err := strconv.Atoi(missionMap[k])
			if err != nil {
				log.Println(fmt.Sprintf("Mission: %s has invalid field: %s", missionId, k))
				return nil, err
			}
			mission.Version = int64(version)
		case "condition":
			condition, err := strconv.Atoi(missionMap[k])
			if err != nil {
				log.Println(fmt.Sprintf("Mission: %s has invalid field: %s", missionId, k))
				return nil, err
			}
			mission.Condition = WinnerCondition(condition)
		default:
			log.Println(fmt.Sprintf("Mission: %s has invalid field: %s", missionId, k))
			return nil, errors.New("invalid mission field")
		}
	}

	return mission, nil
}

func (u *Use) SetMission(mission *Mission) error {
	_, err := u.redisClient.HSet(fmt.Sprintf("mission_%s", mission.Id), "version", mission.Version).Result()
	if err != nil {
		return err
	}
	for _, v := range mission.Points {
		prefix, err := GetPointPrefix(v)
		if err != nil {
			return err
		}
		_, err = u.redisClient.HSet(fmt.Sprintf("mission_%s", mission.Id), fmt.Sprintf("%s_%d", prefix, EncodePointToUint64(v)), int32(v.Type)).Result()
		if err != nil {
			return err
		}
	}
	_, err = u.redisClient.HSet(fmt.Sprintf("mission_%s", mission.Id), "condition", int32(mission.Condition)).Result()
	if err != nil {
		return err
	}
	return nil
}

func (u *Use) GetVersion(missionId string) (int64, error) {
	version, err := u.redisClient.HGet(fmt.Sprintf("mission_%s", missionId), "version").Result()
	if err != nil {
		return 0, err
	}
	intV, err := strconv.Atoi(version)
	if err != nil {
		return 0, err
	}
	return int64(intV), nil
}

func (u *Use) CanMove(missionId string, point *Point) bool {
	return !u.redisClient.HExists(fmt.Sprintf("mission_%s", missionId), fmt.Sprintf("block_%d", EncodePointToUint64(point))).Val()
}

func (u *Use) IsArriveTarget(missionId string, point *Point) bool {
	return u.redisClient.HExists(fmt.Sprintf("mission_%s", missionId), fmt.Sprintf("target_%d", EncodePointToUint64(point))).Val()
}

func (u *Use) IsEnemy(missionId string, point *Point) bool {
	return u.redisClient.HExists(fmt.Sprintf("mission_%s", missionId), fmt.Sprintf("target_%d", EncodePointToUint64(point))).Val()
}
