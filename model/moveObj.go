package model

import (
	"github.com/ClessLi/Game-test/resource"
)

// 可移动的游戏对象
type MoveObj struct {
	GameBasicObj
	SpeedX      float32
	SpeedY      float32
	BounceFrame float32
	// 移动时的动画纹理
	moveTextures []*resource.Texture2D
	// 静止时的动画纹理
	standTextures []*resource.Texture2D
	//当前静止帧
	standIndex int
	//静止帧之间的切换阈值
	standDelta float32
	//当前运动帧
	moveIndex int
	//运动帧之间的切换阈值
	moveDelta float32
}

func NewMoveObj(obj GameBasicObj, moveTextures []*resource.Texture2D, standTextures []*resource.Texture2D) *MoveObj {
	return &MoveObj{
		GameBasicObj:  obj,
		moveTextures:  moveTextures,
		standTextures: standTextures,
		standIndex:    0,
		standDelta:    0,
		moveIndex:     0,
		moveDelta:     0,
	}
}

//恢复静止
func (m *MoveObj) Stand(delta float32) {
	if m.standIndex >= len(m.standTextures) {
		m.standIndex = 0
	}
	m.standDelta += delta
	if m.standDelta > 0.1 {
		m.standDelta = 0
		m.texture = m.standTextures[m.standIndex]
		m.standIndex += 1
	}
}

//由用户主动发起的运动
func (m *MoveObj) Move(delta float32) {
	if m.moveIndex >= len(m.moveTextures) {
		m.moveIndex = 0
	}
	m.moveDelta += delta
	if m.moveDelta > 0.05 {
		m.moveDelta = 0
		m.texture = m.moveTextures[m.moveIndex]
		m.moveIndex += 1
	}
}
