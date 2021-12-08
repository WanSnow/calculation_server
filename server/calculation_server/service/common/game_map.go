package common

import "github.com/wansnow/calculation_server/server/calculation_server/service/logic_service"

var (
	TriggerMap   = make(map[string]chan logic_service.Trigger)
	MainLogicMap = make(map[string]string)
)
