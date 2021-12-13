package mission

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"strings"
)

type Use struct {
	redisClient *redis.Client
}

func (u *Use) GetMission(missionId string) (*Mission, error) {
	mission := new(Mission)
	mission.Id = missionId
	missionMap := u.redisClient.HGetAll(missionId).Val()
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
			version, err := strconv.Atoi(vals[0])
			if err != nil {
				log.Println(fmt.Sprintf("Mission: %s has invalid field: %s", missionId, k))
				return nil, err
			}
			mission.Version = int64(version)
		default:
			log.Println(fmt.Sprintf("Mission: %s has invalid field: %s", missionId, k))
			return nil, errors.New("invalid mission field")
		}
	}

	return mission, nil
}

func (u *Use) SetMission(mission *Mission) error {
	u.redisClient.HSet(mission.Id, "version", mission.Version)
	for _, v := range mission.Points {
		prefix, err := GetPointPrefix(v)
		if err != nil {
			return err
		}
		u.redisClient.HSet(mission.Id, fmt.Sprintf("%s_%d", prefix, EncodePointToUint64(v)), struct{}{})
	}
	return nil
}
