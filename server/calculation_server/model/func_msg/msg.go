package func_msg

const (
	MOVE = byte(1) << iota
	SHOT
	TURN
	TURN_WEAPON
)

type Msg struct {
	FuncType byte
	Param    byte
	Id       string
}

func Encode(msg Msg) (mb []byte) {
	mb = append(mb, msg.FuncType, msg.Param)
	mb = append(mb, []byte(msg.Id)...)
	return
}

func Decode(mb []byte) (msg Msg) {
	msg.FuncType = mb[0]
	msg.Param = mb[1]
	msg.Id = string(mb[2:])
	return
}
