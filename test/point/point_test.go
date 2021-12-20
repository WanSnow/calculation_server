package point

import (
	"fmt"
	"github.com/wansnow/calculation_server/model/mission"
	"testing"
)

func TestPoint(t *testing.T) {
	point := uint64(1152921516418007040)
	m := mission.DecodeUint64ToPoint(point)
	fmt.Println(m)
	point2 := &mission.Point{
		Type: 1,
		X:    1,
		Y:    1,
	}
	fmt.Println(mission.EncodePointToUint64(point2))
}
