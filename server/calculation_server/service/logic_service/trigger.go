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
	AttackedFunc string
}

func (a *Attacked) Call() ([]func_msg.Msg, error) {
	return RunFunc(a.AttackedFunc, "")
}

type Crash struct {
	CrashFunc string
}

func (c *Crash) Call() ([]func_msg.Msg, error) {
	return RunFunc(c.CrashFunc, "")
}

type Discover struct {
	Sight        string
	DiscoverFunc string
}

func (d *Discover) Call() ([]func_msg.Msg, error) {
	return RunFunc(d.DiscoverFunc, d.Sight)
}

type End struct {
}

func (e *End) Call() ([]func_msg.Msg, error) {
	return nil, ErrEndGame
}
