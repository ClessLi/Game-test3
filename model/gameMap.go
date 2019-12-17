package model

import (
	"github.com/ClessLi/Game-test/camera"
	"github.com/ClessLi/Game-test/resource"
	"github.com/ClessLi/Game-test/sprite"
	"github.com/ClessLi/resolvForGame/resolv"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type WorldInterface interface {
	Create()
	Update(float64)
	Draw()
	Destroy()
	SetKey(glfw.Key, bool)
	GetKey(glfw.Key) bool
}

type GameMap struct {
	Player                *Player
	Map                   *Space
	FloatingPlatform      *Line
	FloatingPlatformY     float64
	FloatingPlatformDelta float64
	//精灵渲染器
	renderer *sprite.SpriteRenderer
	//摄像头
	camera *camera.Camera2D
	Keys   [1024]bool
	Init   func()
	WorldSize
}

func (gm *GameMap) Create() {
	c := int32(16)

	//初始化着色器
	resource.LoadShader("./glsl/shader.vs", "./glsl/shader.fs", "sprite")
	shader := resource.GetShader("sprite")
	shader.Use()
	shader.SetInt("image", 0)
	//初始化精灵渲染器
	gm.renderer = sprite.NewSpriteRenderer(shader)
	//设置投影
	projection := mgl32.Ortho(0, float32(gm.ScreenWidth), float32(gm.ScreenHeight), 0, -1, 1)
	shader.SetMatrix4fv("projection", &projection[0])

	//加载资源
	resource.LoadTexture(gl.TEXTURE0, "./image/platformLine.png", "platformLine")
	resource.LoadTexture(gl.TEXTURE0, "./image/line.png", "line")
	resource.LoadTexture(gl.TEXTURE0, "./image/spike.png", "spike")
	resource.LoadTexture(gl.TEXTURE0, "./image/wall.png", "wall")
	resource.LoadTexture(gl.TEXTURE0, "./image/bat/x.png", "x")
	resource.LoadTexture(gl.TEXTURE0, "./image/bat/0.png", "0")
	resource.LoadTexture(gl.TEXTURE0, "./image/bat/1.png", "1")
	resource.LoadTexture(gl.TEXTURE0, "./image/bat/2.png", "2")
	resource.LoadTexture(gl.TEXTURE0, "./image/bat/3.png", "3")
	resource.LoadTexture(gl.TEXTURE0, "./image/bat/4.png", "4")
	resource.LoadTexture(gl.TEXTURE0, "./image/bat/5.png", "5")
	resource.LoadTexture(gl.TEXTURE0, "./image/bat/6.png", "6")
	resource.LoadTexture(gl.TEXTURE0, "./image/bat/7.png", "7")

	// 初始化地图
	gm.Init()
	//创建摄像头,将摄像头同步到玩家位置
	gm.camera = camera.NewDefaultCamera(float32(gm.WorldHeight),
		float32(gm.WorldWidth),
		float32(gm.ScreenWidth),
		float32(gm.ScreenHeight),
		mgl32.Vec2{float32(gm.WorldWidth/2 - gm.ScreenWidth/2), float32(gm.WorldHeight/2 - gm.ScreenHeight/2)})

	//创建测试游戏人物
	gm.Player = NewRecPlayer(gm.WorldWidth/2, gm.WorldHeight/2, 100, 100, []*resource.Texture2D{
		resource.GetTexture("0"),
		resource.GetTexture("1"),
		resource.GetTexture("2"),
		resource.GetTexture("3"),
		resource.GetTexture("4"),
		resource.GetTexture("5"),
		resource.GetTexture("6"),
		resource.GetTexture("7"),
	}, []*resource.Texture2D{
		resource.GetTexture("0"),
		resource.GetTexture("1"),
		resource.GetTexture("2"),
		resource.GetTexture("3"),
		resource.GetTexture("4"),
		resource.GetTexture("5"),
		resource.GetTexture("6"),
		resource.GetTexture("7"),
	}, 0, &mgl32.Vec3{1, 1, 1})

	gm.Map.Add(gm.Player)
	gm.FloatingPlatform = NewLine(c*10, gm.WorldHeight-c*7, c*30, gm.WorldHeight-c*8, nil, []*resource.Texture2D{
		resource.GetTexture("platformLine"),
	})
	gm.FloatingPlatform.AddTags("ramp")
	gm.Map.Add(gm.FloatingPlatform)
	gm.FloatingPlatformY = float64(gm.FloatingPlatform.Shape.Y)
	gm.FloatingPlatformDelta = 0.0
}

func (gm *GameMap) Update(delta float64) {
	gm.Player.SpeedY += 0.5

	friction := float32(0.5)
	accel := 0.5 + friction

	maxSpd := float32(3)
	gm.FloatingPlatformDelta += delta
	if gm.FloatingPlatformDelta > 60 {
		gm.FloatingPlatformDelta -= 60
	}
	gm.FloatingPlatformY += math.Sin(gm.FloatingPlatformDelta/1000) * .5

	gm.FloatingPlatform.Shape.Y = int32(gm.FloatingPlatformY)
	gm.FloatingPlatform.Shape.Y2 = int32(gm.FloatingPlatformY) - 16

	if gm.Player.SpeedX > friction {
		gm.Player.SpeedX -= friction
	} else if gm.Player.SpeedX < -friction {
		gm.Player.SpeedX += friction
	} else {
		gm.Player.SpeedX = 0
	}

	playerMove := false

	if gm.Keys[glfw.KeyRight] || gm.Keys[glfw.KeyD] {
		playerMove = true
		gm.Player.isXReverse = -1
		gm.Player.SpeedX += accel
	}

	if gm.Keys[glfw.KeyLeft] || gm.Keys[glfw.KeyA] {
		playerMove = true
		gm.Player.isXReverse = 1
		gm.Player.SpeedX -= accel
	}

	if gm.Player.SpeedX > maxSpd {
		gm.Player.SpeedX = maxSpd
	}

	if gm.Player.SpeedX < -maxSpd {
		gm.Player.SpeedX = -maxSpd
	}

	// JUMP

	// Check for a collision downwards by just attempting a resolution downwards and seeing if it collides with something.
	down := gm.Map.Resolve(gm.Player, 0, 4)
	onGround := down.Colliding()

	if (gm.Keys[glfw.KeyUp] || gm.Keys[glfw.KeyW]) && onGround {
		playerMove = true
		gm.Player.SpeedY = -8
	}

	if !playerMove {
		gm.Player.MoveObj.Stand(float32(delta))
	} else {
		gm.Player.MoveObj.Move(float32(delta))
	}

	x := int32(gm.Player.SpeedX)
	y := int32(gm.Player.SpeedY)

	solids := gm.Map.FilterByTags("solid")
	ramps := gm.Map.FilterByTags("ramp")

	// X-movement. We only want to collide with solid objects (not ramps) because we want to be able to move up them
	// and don't need to be inhibited on the x-axis when doing so.

	if res := solids.Resolve(gm.Player, x, 0); res.Colliding() {
		x = res.ResolveX
		gm.Player.SpeedX = 0
	}

	gm.Player.Shape.X += x

	// Y movement. We check for ramp collision first; if we find it, then we just automatically will
	// slide up the ramp because the player is moving into it.

	// We look for ramps a little aggressively downwards because when walking down them, we want to stick to them.
	// If we didn't do this, then you would "bob" when walking down the ramp as the Player moves too quickly out into
	// space for gravity to push back down onto the ramp.
	res := ramps.Resolve(gm.Player, 0, y+4)

	if y < 0 || (res.Teleporting && res.ResolveY < -gm.Player.Shape.H/2) {
		res = Collision{}
	}

	if !res.Colliding() {
		res = solids.Resolve(gm.Player, 0, y)
	}

	if res.Colliding() {
		y = res.ResolveY
		gm.Player.SpeedY = 0
	}

	gm.Player.Shape.Y += y

}

func (gm *GameMap) Draw() {

	resource.GetShader("sprite").SetMatrix4fv("view", gm.camera.GetViewMatrix())
	//game.player.MoveBy(float32(delta))
	gm.Player.Draw(gm.renderer)
	//摄像头跟随
	px, py := gm.Player.GetXY()
	size := gm.Player.GetSize()
	gm.camera.InPosition(float32(px-gm.ScreenWidth/2)+size[0], float32(py-gm.ScreenHeight/2)+size[1])

	// TO-DO: 由于渲染依赖camera，暂时将space内各个对象渲染放在这个位置
	for _, shape := range *gm.Map {
		switch shape.(type) {
		case *Rectangle:
			if shape != gm.Player.Rectangle && gm.isInCamera(shape) && (shape.HasTags("isWall") || shape.HasTags("isSpike")) {

				shape.Draw(gm.renderer)

			}
		case *Line:
			if gm.isInCamera(shape) {

				shape.Draw(gm.renderer)

			}
		}

	}

	//if gm.DrawHelpText {
	//    DrawText(32, 16,
	//        "-Platformer test-",
	//        "You are the green square.",
	//        "Use the arrow keys to move.",
	//        "Press X to jump.",
	//        "You can jump through blue ramps / platforms.")
	//}

}

func (gm *GameMap) Destroy() {
	gm.Map.Clear()
}

func (gm *GameMap) isInCamera(shape Shape) bool {
	cp := gm.camera.GetPosition()
	cx := int32(cp.X())
	cy := int32(cp.Y())
	cameraRec := resolv.NewRectangle(cx, cy, gm.ScreenWidth, gm.ScreenHeight)
	return cameraRec.IsColliding(shape.GetShapeObj())
}

func (gm *GameMap) SetKey(key glfw.Key, press bool) {
	gm.Keys[key] = press
}

func (gm *GameMap) GetKey(key glfw.Key) bool {
	return gm.Keys[key]
}
