package mission

import "fmt"

const (
	Block = int64(1) << iota
	TargetPosition
	Enemy
	Bullet
	Player
)

type Point struct {
	Type int64
	X    uint32
	Y    uint32
}

func (p *Point) String() string {

	//prefix
	s := "Point_"

	//type
	switch p.Type {
	case Block:
	case TargetPosition:
		s += "TargetPosition"
	case Enemy:
		s += "Enemy"
	case Bullet:
		s += "Bullet"
	case Player:
		s += "Player"
	default:
		s += "UnKnow"
	}

	//point
	point := int64(p.X)<<32 + int64(p.Y)
	s = fmt.Sprintf("%s_%d", s, point)
	return s
}
