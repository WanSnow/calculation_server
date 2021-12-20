package mission

import (
	"errors"
)

var (
	ErrPointType = errors.New("error point type")
)

func EncodePointToUint64(p *Point) uint64 {
	point := uint64(p.Type) << 60
	point += uint64(p.X) << 30
	point += uint64(p.Y)
	return point
}

func DecodeUint64ToPoint(point uint64) (p *Point) {
	p.Y = uint32(0xfffffff & point)
	p.X = uint32(0xfffffff & point >> 30)
	p.Type = PointType(point >> 60)
	return
}

func GetPointPrefix(point *Point) (string, error) {
	switch point.Type {
	case PointType_POINT_TYPE_BLOCK:
		return "block", nil
	case PointType_POINT_TYPE_POSITION:
		return "position", nil
	case PointType_POINT_TYPE_TARGET:
		return "target", nil
	case PointType_POINT_TYPE_ENEMY:
		return "enemy", nil
	default:
		return "", ErrPointType
	}
}
