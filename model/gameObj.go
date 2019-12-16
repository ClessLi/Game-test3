package model

import (
	"github.com/ClessLi/Game-test/resource"
	"github.com/go-gl/mathgl/mgl32"
)

type GameBasicObj struct {
	texture    *resource.Texture2D
	size       *mgl32.Vec2
	rotate     float32
	color      *mgl32.Vec3
	isXReverse int32
}

func (g *GameBasicObj) GetSize() mgl32.Vec2 {
	return mgl32.Vec2{g.size[0], g.size[1]}
}

func (g *GameBasicObj) ReverseX() {
	g.isXReverse = -1
}

func (g *GameBasicObj) ForWardX() {
	g.isXReverse = 1
}

func NewGameBasicObj(texture *resource.Texture2D, size *mgl32.Vec2, rotate float32, color *mgl32.Vec3) *GameBasicObj {
	return &GameBasicObj{
		texture:    texture,
		size:       size,
		rotate:     rotate,
		color:      color,
		isXReverse: 1,
	}
}
