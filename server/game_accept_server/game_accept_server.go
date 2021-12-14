package game_accept_server

import (
	"context"
	"github.com/rs/xid"
	"github.com/wansnow/calculation_server/config"
	"github.com/wansnow/calculation_server/model/mission"
	"github.com/wansnow/calculation_server/model/player_info"
	"github.com/wansnow/calculation_server/server/calculation_server/service/common"
	"github.com/wansnow/calculation_server/service/accept_game"
)

type StartGameServer struct {
	accept_game.AcceptGameServer
}

func (s *StartGameServer) AcceptGame(_ context.Context, req *accept_game.GameMsg_Request) (*accept_game.GameMsg_Response, error) {
	gameId := xid.New().String()
	err := initNewGame(req, gameId)
	if err != nil {
		return nil, err
	}
	go common.StartGame(gameId, req.PlayerInfo.Id)
	return &accept_game.GameMsg_Response{
		Port:   config.ServerC.StartGameServerPort,
		GameId: gameId,
	}, nil
}

func initNewGame(req *accept_game.GameMsg_Request, gameId string) error {
	pu := player_info.NewUse()
	pu.SetPlayer(req.PlayerInfo)
	mu := mission.NewUse()
	_, err := mu.GetMission(req.MissionId)
	if err != nil {
		return err
	}

	common.GameMap[req.PlayerInfo.Id] = gameId
	return nil
}
