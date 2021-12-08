package logic_service

import (
	"errors"
	"github.com/wansnow/calculation_server/server/calculation_server/model/func_msg"
)

var (
	ErrEndGame = errors.New("game ending")
)

type Trigger interface {
	Call() ([]func_msg.Msg, error)
}

type Attacked struct {
	attackedFunc string
}

func (a *Attacked) Call() ([]func_msg.Msg, error) {
	return RunFunc(a.attackedFunc)
}

type Crash struct {
	crashFunc string
}

func (c *Crash) Call() ([]func_msg.Msg, error) {
	return RunFunc(c.crashFunc)
}

type Discover struct {
	discoverFunc string
}

func (d *Discover) Call() ([]func_msg.Msg, error) {
	return RunFunc(d.discoverFunc)
}

type End struct {
}

func (e *End) Call() ([]func_msg.Msg, error) {
	return nil, ErrEndGame
}
