package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/wansnow/calculation_server/config"
	middleware_nsq "github.com/wansnow/calculation_server/middleware/nsq"
	"github.com/wansnow/calculation_server/server/calculation_server/model/func_msg"
	"github.com/wansnow/calculation_server/server/calculation_server/service/logic_service"
	"github.com/wansnow/calculation_server/service/game_result"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

func StartGame(gameId, playerId, missionId string) {
	cmdChan := make(chan []byte, 10)
	triggerChan := make(chan logic_service.Trigger)
	stopChan := make(chan int)
	waitGroup := &sync.WaitGroup{}

	TriggerChanMap[playerId] = triggerChan
	go middleware_nsq.StartNewProducer(fmt.Sprintf("topic_%s", gameId), cmdChan, waitGroup)
	go middleware_nsq.StartNewConsumer(fmt.Sprintf("topic_%s", gameId), "main", &CalculationMessageHandler{}, waitGroup, stopChan)

	waitGroup.Add(1)
	startTime := time.Now()
loop:
	for {
		var commands []func_msg.Msg
		select {
		case trigger := <-triggerChan:
			var err error
			commands, err = trigger.Call()
			if err != nil {
				if errors.Is(err, logic_service.ErrEndGame) {
					break loop
				}
				log.Println(err)
			}
		default:
			var err error
			commands, err = logic_service.ExecOnceMainLogic(MainLogicMap[playerId], "")
			if err != nil {
				log.Println(err)
			}
		}
		for _, command := range commands {
			if command.Id != playerId {
				continue
			}
			cmdChan <- func_msg.Encode(command)
		}
		time.Sleep(time.Duration(len(commands) * 1000 * int(time.Millisecond)))
	}
	stopChan <- 1
	cmdChan <- []byte("end")
	waitGroup.Done()
	waitGroup.Wait()
	endTime := time.Now()
	close(triggerChan)
	delete(TriggerChanMap, playerId)
	delete(GameMap, playerId)
	close(cmdChan)

	allTime := endTime.Sub(startTime)
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", config.ServerC.HeartBeatServerPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := game_result.NewGameResultClient(conn)
	_, err = client.GameResult(context.Background(), &game_result.ResultMsg_Request{
		GameId:    gameId,
		PlayerId:  playerId,
		MissionId: missionId,
		Time:      int64(allTime),
	})
	if err != nil {
		return
	}
}
