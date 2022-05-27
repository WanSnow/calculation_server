package common

import (
	"fmt"
	"github.com/wansnow/calculation_server/model/game"
	"github.com/wansnow/calculation_server/model/mission"
	"github.com/wansnow/calculation_server/server/calculation_server/model/trigger_msg"
	"github.com/wansnow/calculation_server/server/calculation_server/service/logic_service"
	"log"
)

func Move(id string, param byte) {
	gu := game.NewUse()
	mu := mission.NewUse()
	gameId := GameMap[id]
	g, err := gu.GetGame(gameId)
	if err != nil {
		TriggerChanMap[id] <- &logic_service.End{}
		return
	}
	playerPoint := g.PlayerPoint
	missionId := g.MissionId
	condition := g.Condition
	switch g.PlayerDirection {
	case game.Direction_DIRECTION_UP:
		nextPlayerPoint(playerPoint, mu, condition, param, id, missionId, func() *mission.Point {
			return &mission.Point{
				Type: mission.PointType_POINT_TYPE_BLOCK,
				X:    playerPoint.X,
				Y:    playerPoint.Y + 1,
			}
		})
	case game.Direction_DIRECTION_RIGHT:
		nextPlayerPoint(playerPoint, mu, condition, param, id, missionId, func() *mission.Point {
			return &mission.Point{
				Type: mission.PointType_POINT_TYPE_BLOCK,
				X:    playerPoint.X + 1,
				Y:    playerPoint.Y,
			}
		})
	case game.Direction_DIRECTION_DOWN:
		nextPlayerPoint(playerPoint, mu, condition, param, id, missionId, func() *mission.Point {
			return &mission.Point{
				Type: mission.PointType_POINT_TYPE_BLOCK,
				X:    playerPoint.X,
				Y:    playerPoint.Y - 1,
			}
		})
	case game.Direction_DIRECTION_LEFT:
		nextPlayerPoint(playerPoint, mu, condition, param, id, missionId, func() *mission.Point {
			return &mission.Point{
				Type: mission.PointType_POINT_TYPE_BLOCK,
				X:    playerPoint.X - 1,
				Y:    playerPoint.Y,
			}
		})
	}
	g.PlayerPoint = playerPoint
	err = gu.SaveGame(gameId, g)
	if err != nil {
		log.Println(err)
		TriggerChanMap[id] <- &logic_service.End{}
		return
	}
}

func nextPlayerPoint(oldPoint *mission.Point, mu *mission.Use, condition mission.WinnerCondition, param byte, playerId, missionId string, nextPoint func() *mission.Point) bool {
	for i := byte(0); i < param; i++ {
		nextP := nextPoint()
		if mu.CanMove(missionId, nextP) {
			oldPoint.X = nextP.X
			oldPoint.Y = nextP.Y
		} else {
			TriggerChanMap[playerId] <- TriggerMap[fmt.Sprintf("%s_%d", playerId, trigger_msg.CRASH)]
			break
		}
		if condition == mission.WinnerCondition_WINNER_CONDITION_REACH_TARGET_POSITION {
			if mu.IsArriveTarget(missionId, &mission.Point{
				Type: mission.PointType_POINT_TYPE_TARGET,
				X:    oldPoint.X,
				Y:    oldPoint.Y,
			}) {
				TriggerChanMap[playerId] <- &logic_service.End{}
				return true
			}
		}
	}
	return false
}

func Shot(id string) {
	gameId := GameMap[id]
	gu := game.NewUse()
	g, err := gu.GetGame(gameId)
	if err != nil {
		TriggerChanMap[id] <- &logic_service.End{}
		return
	}

	bullet := &game.Bullet{
		Direction: g.WeaponDirection,
	}
	bullet.Point.Type = mission.PointType_POINT_TYPE_BULLET
	switch g.WeaponDirection {
	case game.Direction_DIRECTION_UP:
		bullet.Point.X = g.PlayerPoint.X
		bullet.Point.Y = g.PlayerPoint.Y + 1
	case game.Direction_DIRECTION_RIGHT:
		bullet.Point.X = g.PlayerPoint.X + 1
		bullet.Point.Y = g.PlayerPoint.Y
	case game.Direction_DIRECTION_DOWN:
		bullet.Point.X = g.PlayerPoint.X
		bullet.Point.Y = g.PlayerPoint.Y - 1
	case game.Direction_DIRECTION_LEFT:
		bullet.Point.X = g.PlayerPoint.X - 1
		bullet.Point.Y = g.PlayerPoint.Y
	default:
		return
	}
	g.Bullets = append(g.Bullets, bullet)
	err = gu.SaveGame(gameId, g)
	if err != nil {
		log.Println(err)
		TriggerChanMap[id] <- &logic_service.End{}
		return
	}
}

func Turn(id string, param byte) {
	gameId := GameMap[id]
	gu := game.NewUse()
	g, err := gu.GetGame(gameId)
	if err != nil {
		TriggerChanMap[id] <- &logic_service.End{}
		return
	}
	g.PlayerDirection = game.Direction(param)
	err = gu.SaveGame(gameId, g)
	if err != nil {
		TriggerChanMap[id] <- &logic_service.End{}
		return
	}
}

func TurnWeapon(id string, param byte) {
	gameId := GameMap[id]
	gu := game.NewUse()
	g, err := gu.GetGame(gameId)
	if err != nil {
		TriggerChanMap[id] <- &logic_service.End{}
		return
	}
	g.WeaponDirection = game.Direction(param)
	err = gu.SaveGame(gameId, g)
	if err != nil {
		TriggerChanMap[id] <- &logic_service.End{}
		return
	}
}

func MoveBullet(id string) {
	gameId := GameMap[id]
	gu := game.NewUse()
	mu := mission.NewUse()
	g, err := gu.GetGame(gameId)
	if err != nil {
		TriggerChanMap[id] <- &logic_service.End{}
		return
	}
	bullets := g.GetBullets()
	destoryBullets := make([]int, 0)
	for i, bullet := range bullets {
		switch bullet.GetDirection() {
		case game.Direction_DIRECTION_UP:
			if !nextBulletPoint(bullet, mu, g.Condition, id, g.MissionId, func() *mission.Point {
				return &mission.Point{
					Type: mission.PointType_POINT_TYPE_BLOCK,
					X:    bullet.Point.X,
					Y:    bullet.Point.Y + 1,
				}
			}) {
				destoryBullets = append(destoryBullets, i)
			}
		case game.Direction_DIRECTION_RIGHT:
			if !nextBulletPoint(bullet, mu, g.Condition, id, g.MissionId, func() *mission.Point {
				return &mission.Point{
					Type: mission.PointType_POINT_TYPE_BLOCK,
					X:    bullet.Point.X + 1,
					Y:    bullet.Point.Y,
				}
			}) {
				destoryBullets = append(destoryBullets, i)
			}
		case game.Direction_DIRECTION_DOWN:
			if !nextBulletPoint(bullet, mu, g.Condition, id, g.MissionId, func() *mission.Point {
				return &mission.Point{
					Type: mission.PointType_POINT_TYPE_BLOCK,
					X:    bullet.Point.X,
					Y:    bullet.Point.Y - 1,
				}
			}) {
				destoryBullets = append(destoryBullets, i)
			}
		case game.Direction_DIRECTION_LEFT:
			if !nextBulletPoint(bullet, mu, g.Condition, id, g.MissionId, func() *mission.Point {
				return &mission.Point{
					Type: mission.PointType_POINT_TYPE_BLOCK,
					X:    bullet.Point.X - 1,
					Y:    bullet.Point.Y,
				}
			}) {
				destoryBullets = append(destoryBullets, i)
			}
		default:
			return
		}
	}
	oldBullets := make([]*game.Bullet, len(g.Bullets))
	newBullets := make([]*game.Bullet, 0)

	copy(oldBullets, g.Bullets)
	if len(destoryBullets) > 0 {
		newBullets = append(newBullets, oldBullets[:destoryBullets[0]]...)
	}
	i := 1
	for ; i < len(destoryBullets); i++ {
		newBullets = append(newBullets, oldBullets[destoryBullets[0]+1:destoryBullets[i]]...)
	}
	g.Bullets = newBullets
	err = gu.SaveGame(gameId, g)
	if err != nil {
		log.Println(err)
		TriggerChanMap[id] <- &logic_service.End{}
		return
	}
}

func nextBulletPoint(oldBullet *game.Bullet, mu *mission.Use, condition mission.WinnerCondition, playerId, missionId string, nextPoint func() *mission.Point) bool {
	for i := byte(0); i < 10; i++ {
		nextP := nextPoint()
		if condition == mission.WinnerCondition_WINNER_CONDITION_DESTROY_ENEMY {
			if mu.IsEnemy(missionId, &mission.Point{
				Type: mission.PointType_POINT_TYPE_ENEMY,
				X:    nextP.X,
				Y:    nextP.Y,
			}) {
				TriggerChanMap[playerId] <- &logic_service.End{}
				return false
			}
		}
		if mu.CanMove(missionId, nextP) {
			oldBullet.Point.X = nextP.X
			oldBullet.Point.Y = nextP.Y
		} else {
			return false
		}
	}
	return true
}
