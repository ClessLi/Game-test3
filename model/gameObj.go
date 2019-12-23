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

func NewGameBasicObj(texture string, size *mgl32.Vec2, rotate float32, color *mgl32.Vec3) *GameBasicObj {
	return &GameBasicObj{
		texture:    resource.GetTexture(texture),
		size:       size,
		rotate:     rotate,
		color:      color,
		isXReverse: 1,
	}
}

func GetTexturesByName(names ...string) []*resource.Texture2D {
	if names == nil {
		return nil
	}
	textures := make([]*resource.Texture2D, 0)
	for _, name := range names {
		textures = append(textures, resource.GetTexture(name))
	}
	return textures
}
