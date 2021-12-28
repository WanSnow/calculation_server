package game_accept_server

import (
	"context"
	"fmt"
	"github.com/rs/xid"
	"github.com/wansnow/calculation_server/config"
	"github.com/wansnow/calculation_server/model/game"
	"github.com/wansnow/calculation_server/model/mission"
	"github.com/wansnow/calculation_server/model/player_info"
	"github.com/wansnow/calculation_server/server/calculation_server/service/common"
	"github.com/wansnow/calculation_server/service/accept_game"
	"google.golang.org/grpc"
	"log"
	"net"
)

type AcceptGameServer struct {
	accept_game.AcceptGameServer
}

func NewAcceptGameServer() *AcceptGameServer {
	return &AcceptGameServer{}
}

func StartAcceptGameServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.ServerC.AcceptServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accept_game.RegisterAcceptGameServer(grpcServer, NewAcceptGameServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func (s *AcceptGameServer) AcceptGame(_ context.Context, req *accept_game.GameMsg_Request) (*accept_game.GameMsg_Response, error) {
	gameId := xid.New().String()
	err := initNewGame(req, gameId)
	if err != nil {
		return nil, err
	}
	go common.StartGame(gameId, req.PlayerInfo.Id, req.MissionId)
	return &accept_game.GameMsg_Response{
		Port:   config.ServerC.StartGameServerPort,
		GameId: gameId,
	}, nil
}

func initNewGame(req *accept_game.GameMsg_Request, gameId string) error {
	pu := player_info.NewUse()
	pu.SetPlayer(req.PlayerInfo)
	mu := mission.NewUse()
	m, err := mu.GetMission(req.MissionId)
	if err != nil {
		return err
	}

	gu := game.NewUse()
	err = gu.InitGame(gameId, m)
	if err != nil {
		return err
	}

	common.GameMap[req.PlayerInfo.Id] = gameId
	return nil
}
