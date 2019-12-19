package model

import (
	"github.com/ClessLi/Game-test/resource"
	"github.com/ClessLi/Game-test/sprite"
	"github.com/ClessLi/resolvForGame/resolv"
	"github.com/go-gl/mathgl/mgl32"
)

type Rectangle struct {
	Shape *resolv.Rectangle
	*MoveObj
	friction float32
	maxSpd   float32
}

// NewRectangle returns a new Rectangle instance.
func NewRectangle(x, y, w, h int32, friction float32, moveTextures []*resource.Texture2D, standTextures []*resource.Texture2D) *Rectangle {
	rec := &Rectangle{
		Shape:    resolv.NewRectangle(x, y, w, h),
		friction: friction,
	}
	rec.MoveObj = NewMoveObj(*NewGameBasicObj(standTextures[0], &mgl32.Vec2{float32(w), float32(h)}, 0, &mgl32.Vec3{1, 1, 1}), moveTextures, standTextures)
	return rec
}

func (r *Rectangle) IsColliding(other Shape) bool {
	return r.Shape.IsColliding(other.GetShapeObj())
}

func (r *Rectangle) WouldBeColliding(other Shape, dx, dy int32) bool {
	return r.Shape.WouldBeColliding(other.GetShapeObj(), dx, dy)
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

func (r *Rectangle) GetXY2() (int32, int32) {
	x2 := r.Shape.X + r.Shape.W
	y2 := r.Shape.Y + r.Shape.H
	return x2, y2
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

func (r *Rectangle) GetShapeObj() resolv.Shape {
	return r.Shape
}

func (r *Rectangle) GetFriction() float32 {
	return r.friction
}

func (r *Rectangle) SetFriction(friction float32) {
	r.friction = friction
}

func (r *Rectangle) GetMaxSpd() float32 {
	return r.maxSpd
}

func (r *Rectangle) SetMaxSpd(spd float32) {
	r.maxSpd = spd
}
