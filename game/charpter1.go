package game

import (
	"github.com/ClessLi/Game-test/resource"
	"github.com/ClessLi/Game-test3/model"
)

type C1 game

func NewC1(width, height int32) *C1 {
	var (
		ww int32 = 1600
		wh int32 = 400
		sw int32 = 800
		sh int32 = 600
		cw int32 = 16
		ch int32 = 16
	)
	xF := float32(width) / float32(sw)
	yF := float32(height) / float32(sh)
	ww = int32(xF * float32(ww))
	wh = int32(yF * float32(wh))
	sw = int32(xF * float32(sw))
	sh = int32(yF * float32(sh))
	cw = int32(xF * float32(cw))
	ch = int32(yF * float32(ch))
	c1 := &C1{}
	s1 := model.GameMap{
		WorldSize: *model.NewWorldSize(ww, wh, sw, sh, false),
	}
	s1.Init = func() {
		s1.Map = model.NewSpace()
		s1.Map.Clear()

		// A ramp
		line := model.NewLine(ww/4+cw, wh-ch*4, ww/4+cw*11, wh-ch*10, 0.5, nil, []*resource.Texture2D{resource.GetTexture("line")})
		line.AddTags("ramp")
		s1.Map.Add(line)

		line = model.NewLine(ww/4+cw*11, wh-ch*10, ww/4+cw*40, wh-ch*10, 0.1, nil, []*resource.Texture2D{resource.GetTexture("line")})
		line.AddTags("ramp")
		s1.Map.Add(line)

		line = model.NewLine(ww/4+cw*40, wh-ch*10, ww/4+cw*50, wh-ch*4, 0.5, nil, []*resource.Texture2D{resource.GetTexture("line")})
		line.AddTags("ramp")
		s1.Map.Add(line)

		for y := int32(0); y < wh; y += ch {

			for x := int32(0); x < ww; x += cw {

				// 构建四周的墙
				if y <= ch*4 || y >= wh-ch*4 || x <= cw*4 || x >= ww-cw*4 {
					wall := model.NewRectangle(x, y, cw, ch, 0.5, nil, []*resource.Texture2D{
						resource.GetTexture("wall"),
					})
					wall.AddTags("isWall", "solid", "ramp")
					s1.Map.Add(wall)

				}

				// 构建顶部尖刺
				if y == ch*5 && x > cw*4 && x < ww-cw*4 {
					spike := model.NewRectangle(x, y, cw, ch, 0.01, nil, []*resource.Texture2D{
						resource.GetTexture("spike"),
					})
					spike.AddTags("isSpike", "injurious")
					s1.Map.Add(spike)
				}

			}

		}
	}

	c1.Maps = append(c1.Maps, &s1)
	return c1
}
