package logic_service

import (
	"fmt"
	"github.com/wansnow/calculation_server/server/calculation_server/model/func_msg"
	"os/exec"
	"strconv"
	"strings"
)

const PyCmd = `
def move(id, param):
	print("logic", 1, param, id)
def shot(id):
	print("logic", 2, 0, id)
def turn(id, param):
	print("logic", 4, param, id)
def turn_weapon(id, param):
	print("logic", 8, param, id)

sight={%s}

%s`

func RunFunc(playerFunc, sight string) ([]func_msg.Msg, error) {
	pyLogic := exec.Command("python3", "-c", fmt.Sprintf(PyCmd, sight, playerFunc))
	output, err := pyLogic.Output()
	if err != nil {
		return nil, err
	}

	cmds := strings.Split(strings.TrimSpace(string(output)), "\n")
	var commands []func_msg.Msg
	for _, cmd := range cmds {
		msgs := strings.Split(strings.TrimSpace(cmd), " ")
		if msgs[0] != "logic" {
			continue
		}
		var funcType, param int
		funcType, err = strconv.Atoi(msgs[1])
		if err != nil {
			continue
		}
		param, err = strconv.Atoi(msgs[2])
		if err != nil {
			continue
		}
		commands = append(commands, func_msg.Msg{
			FuncType: byte(funcType),
			Param:    byte(param),
			Id:       msgs[3],
		})
	}
	return commands, nil
}
