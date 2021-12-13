package common

import "github.com/wansnow/calculation_server/server/calculation_server/service/logic_service"

var (
	// TriggerChanMap k: [gameId], v: trigger chan
	TriggerChanMap = make(map[string]chan logic_service.Trigger)
	// TriggerMap k: [playerId]_[triggerType] v: Trigger
	TriggerMap = make(map[string]*logic_service.Trigger)
	// MainLogicMap k: [playerId], v: [MainLogicFunc]
	MainLogicMap = make(map[string]string)
)
