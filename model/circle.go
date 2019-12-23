package model

import (
	"github.com/ClessLi/Game-test/sprite"
	"github.com/ClessLi/resolvForGame/resolv"
	"github.com/go-gl/mathgl/mgl32"
)

type Circle struct {
	Shape *resolv.Circle
	*MoveObj
	friction float32
	maxSpd   float32
}

func NewCircle(x, y, radius int32, friction float32, moveList []string, standList []string) *Circle {
	circle := &Circle{
		Shape:    resolv.NewCircle(x, y, radius),
		friction: friction,
	}
	var texture = ""
	if len(standList) >= 0 {
		texture = standList[0]
	}
	circle.MoveObj = NewMoveObj(*NewGameBasicObj(texture, &mgl32.Vec2{float32(2 * radius), float32(2 * radius)}, 0, &mgl32.Vec3{1, 1, 1}), moveList, standList)
	return circle
}

func (c *Circle) Clear() {
	*c = Circle{}
}

func (c *Circle) IsColliding(other Shape) bool {
	return c.Shape.IsColliding(other.GetShapeObj())
}

func (c *Circle) WouldBeColliding(other Shape, dx, dy int32) bool {
	return c.Shape.WouldBeColliding(other.GetShapeObj(), dx, dy)
}

func (c *Circle) GetBoundingRect() *Rectangle {
	return &Rectangle{
		Shape:    c.Shape.GetBoundingRect(),
		MoveObj:  c.MoveObj,
		friction: c.friction,
		maxSpd:   c.maxSpd,
	}
}

func (c *Circle) GetTags() []string {
	return c.Shape.GetTags()
}

func (c *Circle) ClearTags() {
	c.Shape.ClearTags()
}

func (c *Circle) AddTags(tags ...string) {
	c.Shape.AddTags(tags...)
}

func (c *Circle) RemoveTags(tags ...string) {
	c.Shape.RemoveTags(tags...)
}

func (c *Circle) HasTags(tags ...string) bool {
	return c.Shape.HasTags(tags...)
}

func (c *Circle) GetData() interface{} {
	return c.Shape.GetData()
}

func (c *Circle) SetData(data interface{}) {
	c.Shape.SetData(data)
}

func (c *Circle) GetXY() (int32, int32) {
	return c.Shape.GetXY()
}

func (c *Circle) GetXY2() (int32, int32) {
	x2 := c.Shape.X + c.Shape.Radius
	y2 := c.Shape.Y + c.Shape.Radius
	return x2, y2
}

func (c *Circle) SetXY(x, y int32) {
	c.Shape.SetXY(x, y)
}

func (c *Circle) Move(x, y int32) {
	c.Shape.Move(x, y)
}

func (c *Circle) Draw(renderer *sprite.SpriteRenderer) {
	renderer.DrawSprite(c.texture, &mgl32.Vec2{float32(c.Shape.X - c.Shape.Radius), float32(c.Shape.Y - c.Shape.Radius)}, c.size, c.rotate, c.color, c.isXReverse)
}

func (c *Circle) GetShapeObj() resolv.Shape {
	return c.Shape
}

func (c *Circle) GetFriction() float32 {
	return c.friction
}

func (c *Circle) SetFriction(friction float32) {
	c.friction = friction
}

func (c *Circle) GetMaxSpd() float32 {
	return c.maxSpd
}

func (c *Circle) SetMaxSpd(spd float32) {
	c.maxSpd = spd
}

func (c *Circle) GetSpd() (float32, float32) {
	return c.SpeedX, c.SpeedY
}

func (c *Circle) SetSpd(x, y float32) {
	c.SpeedX = x
	c.SpeedY = y
}
