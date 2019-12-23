package model

import "github.com/go-gl/mathgl/mgl32"

type Liver interface {
	Attack() Shape
}

type Player struct {
	*Rectangle
	weapon Weapon
	AtkVec mgl32.Vec2
}

func NewRecPlayer(x, y, w, h int32, friction, drawMulti float32, moveList []string, standList []string) *Player {
	rec := NewRectangle(x, y, w, h, friction, drawMulti, moveList, standList)
	player := &Player{
		Rectangle: rec,
		weapon:    nil,
	}
	return player
}

func (p *Player) Attack() Shape {
	x, y := p.GetXY()
	if p.isXReverse > 0 {
		x += p.Shape.W * 2 / 3
	} else {
		x += p.Shape.W / 3
	}
	y += p.Shape.H / 4

	switch p.weapon.(type) {
	case *LongRangeWeapon:
		return p.weapon.Attack(x, y, p.AtkVec, p.isXReverse)
	}
	return nil
}

//func (p *Player) Draw(renderer *sprite.SpriteRenderer) {
//	p.Rectangle.Draw(renderer)
//}
