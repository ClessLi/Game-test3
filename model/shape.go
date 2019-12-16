package model

import (
	"github.com/ClessLi/Game-test/sprite"
	"github.com/ClessLi/resolvForGame/resolv"
)

type Shape interface {
	IsColliding(Shape) bool
	WouldBeColliding(Shape, int32, int32) bool
	GetTags() []string
	ClearTags()
	AddTags(...string)
	RemoveTags(...string)
	HasTags(...string) bool
	GetData() interface{}
	SetData(interface{})
	GetXY() (int32, int32)
	SetXY(int32, int32)
	Move(int32, int32)
	Draw(*sprite.SpriteRenderer)
}

type Circle struct {
	Shape *resolv.Circle
	MoveObj
}
