package logic_service

import (
	"github.com/wansnow/calculation_server/server/calculation_server/model/func_msg"
)

func ExecOnceMainLogic(logicFunc, sight string) ([]func_msg.Msg, error) {
	return RunFunc(logicFunc, sight)
}
