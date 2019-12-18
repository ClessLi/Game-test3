package model

import (
	"github.com/ClessLi/Game-test/resource"
	"github.com/ClessLi/Game-test/sprite"
	"github.com/ClessLi/resolvForGame/resolv"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type Line struct {
	Shape *resolv.Line
	MoveObj
}

// NewLine returns a new Line instance.
func NewLine(x, y, x2, y2 int32, moveTextures []*resource.Texture2D, standTextures []*resource.Texture2D) *Line {
	l := &Line{Shape: resolv.NewLine(x, y, x2, y2)}
	rotate := float32(math.Atan2(float64(l.Shape.Y-l.Shape.Y2), float64(l.Shape.X-l.Shape.X2)))
	l.MoveObj = *NewMoveObj(*NewGameBasicObj(standTextures[0], l.GetSize(), rotate, &mgl32.Vec3{1, 1, 1}), moveTextures, standTextures)
	return l
}

func (l *Line) GetSize() *mgl32.Vec2 {
	//return mgl32.Vec2{float32(math.Abs(float64(l.Shape.X - l.Shape.X2))), float32(math.Abs(float64(l.Shape.Y - l.Shape.Y2)))}
	return &mgl32.Vec2{float32(l.Shape.GetLength()), 2}
}

func (l *Line) IsColliding(other Shape) bool {
	return l.Shape.IsColliding(other.GetShapeObj())
}

func (l *Line) WouldBeColliding(other Shape, dx, dy int32) bool {
	return l.Shape.WouldBeColliding(other.GetShapeObj(), dx, dy)
}

func (l *Line) GetTags() []string {
	return l.Shape.GetTags()
}

func (l *Line) ClearTags() {
	l.Shape.ClearTags()
}

func (l *Line) AddTags(tags ...string) {
	l.Shape.AddTags(tags...)
}

func (l *Line) RemoveTags(tags ...string) {
	l.Shape.RemoveTags(tags...)
}

func (l *Line) HasTags(tags ...string) bool {
	return l.Shape.HasTags(tags...)
}

func (l *Line) GetData() interface{} {
	return l.Shape.GetData()
}

func (l *Line) SetData(data interface{}) {
	l.Shape.SetData(data)
}

func (l *Line) GetXY() (int32, int32) {
	return l.Shape.GetXY()
}

func (l *Line) GetXY2() (int32, int32) {
	x2 := l.Shape.X2
	y2 := l.Shape.Y2
	return x2, y2
}

func (l *Line) SetXY(x, y int32) {
	l.Shape.SetXY(x, y)
}

func (l *Line) Move(x, y int32) {
	l.Shape.Move(x, y)
}

func (l *Line) Draw(renderer *sprite.SpriteRenderer) {
	renderer.DrawSprite(l.texture, l.getDrawXY(), l.size, l.rotate, l.color, l.isXReverse)
}

func (l *Line) GetShapeObj() resolv.Shape {
	return l.Shape
}

func (l *Line) getDrawXY() *mgl32.Vec2 {
	centerX, centerY := l.Shape.Center()
	drawX := float32(centerX) - float32(l.Shape.GetLength())/2
	drawY := float32(centerY)
	return &mgl32.Vec2{drawX, drawY}
}
