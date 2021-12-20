package common

import (
	"errors"
	"fmt"
	middleware_nsq "github.com/wansnow/calculation_server/middleware/nsq"
	"github.com/wansnow/calculation_server/server/calculation_server/model/func_msg"
	"github.com/wansnow/calculation_server/server/calculation_server/service/logic_service"
	"log"
	"time"
)

func StartGame(gameId, playerId string) {
	cmdChan := make(chan []byte, 10)
	triggerChan := make(chan logic_service.Trigger)
	TriggerChanMap[playerId] = triggerChan
	go middleware_nsq.StartNewProducer(fmt.Sprintf("topic_%s", gameId), cmdChan)
	go middleware_nsq.StartNewConsumer(fmt.Sprintf("topic_%s", gameId), "main", &CalculationMessageHandler{})

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
	close(triggerChan)
	delete(TriggerChanMap, playerId)
	close(cmdChan)
}
