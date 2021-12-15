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
	GameId       string
	AttackedFunc string
}

func (a *Attacked) Call() ([]func_msg.Msg, error) {
	return RunFunc(a.AttackedFunc, a.GameId)
}

type Crash struct {
	GameId    string
	CrashFunc string
}

func (c *Crash) Call() ([]func_msg.Msg, error) {
	return RunFunc(c.CrashFunc, c.GameId)
}

type Discover struct {
	GameId       string
	DiscoverFunc string
}

func (d *Discover) Call() ([]func_msg.Msg, error) {
	return RunFunc(d.DiscoverFunc, d.GameId)
}

type End struct {
}

func (e *End) Call() ([]func_msg.Msg, error) {
	return nil, ErrEndGame
}
