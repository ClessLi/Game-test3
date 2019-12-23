package model

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
)

type Weapon interface {
	Attack(X, Y int32, vec2 mgl32.Vec2, isXReverse int32) Shape
	CoolDown(delta float64)
}

type LongRangeWeapon struct {
	BoltName string
	CD       float64
	CDDelta  float64
	Speed    float32
	BoltSize int32
}

func (lw *LongRangeWeapon) Attack(X, Y int32, vec2 mgl32.Vec2, isXReverse int32) Shape {
	if lw.CDDelta > 0 {
		return nil
	}
	lw.CDDelta = lw.CD

	SpdX, SpdY := lw.initSpd(vec2)

	bolt := NewCircle(X, Y, lw.BoltSize, 0, nil, []string{lw.BoltName})
	bolt.isXReverse = isXReverse
	bolt.SetSpd(SpdX, SpdY)
	bolt.AddTags("isMove")
	fmt.Println("shooting, x:", bolt.Shape.X, "y:", bolt.Shape.Y, "spdX:", bolt.SpeedX, "spdY:", bolt.SpeedY)
	return bolt
}

func (lw *LongRangeWeapon) CoolDown(delta float64) {
	lw.CDDelta -= delta
}

func (lw *LongRangeWeapon) initSpd(vec2 mgl32.Vec2) (float32, float32) {
	vecLen := vec2.Len()
	return vec2[0] * lw.Speed / vecLen, vec2[1] * lw.Speed / vecLen
}

func NewFireBolt() *LongRangeWeapon {
	weapon := &LongRangeWeapon{
		BoltName: "FireBolt",
		CD:       1.0,
		CDDelta:  0,
		Speed:    20,
		BoltSize: 10,
	}
	return weapon
}
