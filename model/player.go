package model

import (
	"github.com/ClessLi/Game-test/resource"
	"github.com/ClessLi/Game-test/sprite"
	"github.com/ClessLi/resolvForGame/resolv"
	"github.com/go-gl/mathgl/mgl32"
)

type Player struct {
	*Rectangle
}

func NewRecPlayer(x, y, w, h int32, mTextures []*resource.Texture2D, sTextures []*resource.Texture2D, rotate float32, color *mgl32.Vec3) *Player {
	size := &mgl32.Vec2{float32(w), float32(h)}
	rec := resolv.NewRectangle(x, y, w, h)
	moveObj := NewMoveObj(*NewGameBasicObj(sTextures[0], size, rotate, color), mTextures, sTextures)
	player := &Player{&Rectangle{
		Shape:   rec,
		MoveObj: *moveObj,
	}}
	return player
}

func (p *Player) Draw(renderer *sprite.SpriteRenderer) {
	x, y := p.Shape.GetXY()
	renderer.DrawSprite(p.texture, &mgl32.Vec2{float32(x), float32(y)}, p.size, p.rotate, p.color, p.isXReverse)
}
