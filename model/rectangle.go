package model

import (
	"github.com/ClessLi/Game-test/sprite"
	"github.com/ClessLi/resolvForGame/resolv"
	"github.com/go-gl/mathgl/mgl32"
)

type Rectangle struct {
	Shape *resolv.Rectangle
	MoveObj
}

func (r *Rectangle) IsColliding(other Shape) bool {
	return r.Shape.IsColliding(other.(resolv.Shape))
}

func (r *Rectangle) WouldBeColliding(other Shape, dx, dy int32) bool {
	return r.Shape.WouldBeColliding(other.(resolv.Shape), dx, dy)
}

func (r *Rectangle) GetTags() []string {
	return r.Shape.GetTags()
}

func (r *Rectangle) ClearTags() {
	r.Shape.ClearTags()
}

func (r *Rectangle) AddTags(tags ...string) {
	r.Shape.AddTags(tags...)
}

func (r *Rectangle) RemoveTags(tags ...string) {
	r.Shape.RemoveTags(tags...)
}

func (r *Rectangle) HasTags(tags ...string) bool {
	return r.Shape.HasTags(tags...)
}

func (r *Rectangle) GetData() interface{} {
	return r.Shape.GetData()
}

func (r *Rectangle) SetData(data interface{}) {
	r.Shape.SetData(data)
}

func (r *Rectangle) GetXY() (int32, int32) {
	return r.Shape.GetXY()
}

func (r *Rectangle) SetXY(x, y int32) {
	r.Shape.SetXY(x, y)
}

func (r *Rectangle) Move(x, y int32) {
	r.Shape.Move(x, y)
}

func (r *Rectangle) Draw(renderer *sprite.SpriteRenderer) {
	renderer.DrawSprite(r.texture, &mgl32.Vec2{float32(r.Shape.X), float32(r.Shape.Y)}, r.size, r.rotate, r.color, r.isXReverse)
}
