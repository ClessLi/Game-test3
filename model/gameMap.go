package model

import (
	"github.com/ClessLi/Game-test/camera"
	"github.com/ClessLi/Game-test/resource"
	"github.com/ClessLi/Game-test/sprite"
	"github.com/ClessLi/resolvForGame/resolv"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type WorldInterface interface {
	Create()
	Update(float64)
	Draw()
	Destroy()
	// 调整WorldInterface的SetKey(glfw.Key, bool)方法为SetKeyDown(glfw.Key)
	SetKeyDown(glfw.Key)
	// 新增IsPressed、PressedKey方法，用于判断按键是否已按下和标记按键已按下
	IsPressed(...glfw.Key) bool
	PressedKey(glfw.Key)
	ReleaseKey(glfw.Key)
	// 调整WorldInterface的GetKey(glfw.Key) bool方法为HasOneKeyDown(...glfw.Key) bool，
	// 用于判断查询多个或单个按键中是否存在已按下的
	HasOneKeyDown(...glfw.Key) bool
}

type GameMap struct {
	Player *Player
	Map    *Space
	//精灵渲染器
	renderer *sprite.SpriteRenderer
	//摄像头
	camera     *camera.Camera2D
	Keys       [1024]bool
	LockedKeys [1024]bool
	Init       func()
	WorldSize
}

func (gm *GameMap) Create() {
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

	// 初始化地图
	gm.Init()
	//创建摄像头,将摄像头同步到玩家位置
	gm.camera = camera.NewDefaultCamera(float32(gm.WorldHeight),
		float32(gm.WorldWidth),
		float32(gm.ScreenWidth),
		float32(gm.ScreenHeight),
		mgl32.Vec2{float32(gm.WorldWidth/2 - gm.ScreenWidth/2), float32(gm.WorldHeight/2 - gm.ScreenHeight/2)})

	//创建测试游戏人物
	gm.Player = NewRecPlayer(gm.WorldWidth/2, gm.WorldHeight/2, 100, 100, 0.5, 0.5, []string{"0", "1", "2", "3", "4", "5", "6", "7"}, []string{"0", "1", "2", "3", "4", "5", "6", "7"})
	gm.Player.SetMaxSpd(5)

	gm.Player.weapon = NewFireBolt()

	gm.Map.Add(gm.Player)
}

func (gm *GameMap) Update(delta float64) {
	gm.Player.SpeedY += 0.5

	// 更新移动物体
	gm.updateMove()

	// Check for a collision downwards by just attempting a resolution downwards and seeing if it collides with something.
	down := gm.Map.Filter(func(shape Shape) bool {
		if (shape.HasTags("solid") || shape.HasTags("ramp")) && !shape.HasTags("destroyed") {
			return true
		}
		return false
	}).Resolve(gm.Player, 0, 4)
	onGround := down.Colliding()
	gm.Player.isMove = false

	// 角色左右移动
	gm.playerMove(down)

	// JUMP

	gm.playerJump(onGround)

	// Attack
	gm.playerAttack(delta)

	if !gm.Player.isMove {
		gm.Player.MoveObj.Stand(float32(delta))
	} else {
		gm.Player.MoveObj.Move(float32(delta))
	}

	x := int32(gm.Player.SpeedX)
	y := int32(gm.Player.SpeedY)

	solids := gm.Map.FilterByTags("solid")
	ramps := gm.Map.FilterByTags("ramp")
	dangers := gm.Map.FilterByTags("dangerous")

	//fmt.Println("check player is dead or not.")
	// 判断用户是否已死亡
	if res := dangers.Resolve(gm.Player, x, y); res.Colliding() {
		//fmt.Println("player is dead.")
		gm.Player.AddTags("isDead")
	}

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

	//if gm.Player.HasTags("isDead") {
	//	gm.Player.SpeedX = 0
	//}

}

func (gm *GameMap) Draw() {

	resource.GetShader("sprite").SetMatrix4fv("view", gm.camera.GetViewMatrix())
	//game.player.MoveBy(float32(delta))
	// 判断角色是否死亡
	if gm.Player.HasTags("isDead") {
		gm.Player.texture = resource.GetTexture("x")
	}
	gm.Player.Draw(gm.renderer)
	//摄像头跟随
	px, py := gm.Player.GetXY()
	size := gm.Player.GetSize()
	gm.camera.InPosition(float32(px-gm.ScreenWidth/2)+size[0], float32(py-gm.ScreenHeight/2)+size[1])

	// TO-DO: 由于渲染依赖camera，暂时将space内各个对象渲染放在这个位置
	for _, shape := range *gm.Map {
		switch shape.(type) {
		case *Rectangle:
			if shape != gm.Player.Rectangle && gm.isInCamera(shape) && !shape.HasTags("hide") {

				shape.Draw(gm.renderer)

			}
		case *Line:
			if gm.isInCamera(shape) && !shape.HasTags("hide") {

				shape.Draw(gm.renderer)

			}
		case *Circle:
			if gm.isInCamera(shape) && !shape.HasTags("hide") && !shape.HasTags("destroyed") && !shape.HasTags("init") {

				shape.Draw(gm.renderer)

			}
		}

		if shape.HasTags("destroy") {
			shape.RemoveTags("destroy")
			shape.AddTags("destroyed")
		}

	}

	//if gm.DrawHelpText {
	//    DrawText(32, 16,
	//        "-Platformer test-",
	//        "You are the green square.",
	//        "Use the arrow keys to move.",
	//        "Press X to playerJump.",
	//        "You can playerJump through blue ramps / platforms.")
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

func (gm *GameMap) SetKeyDown(key glfw.Key) {
	gm.Keys[key] = true
}

func (gm *GameMap) IsPressed(keys ...glfw.Key) bool {
	for _, key := range keys {
		if gm.LockedKeys[key] {
			return true
		}
	}
	return false
}

func (gm *GameMap) PressedKey(key glfw.Key) {
	gm.LockedKeys[key] = true
}

func (gm *GameMap) ReleaseKey(key glfw.Key) {
	gm.Keys[key] = false
	gm.LockedKeys[key] = false
}

func (gm *GameMap) HasOneKeyDown(keys ...glfw.Key) bool {
	for _, key := range keys {
		if gm.Keys[key] {
			return true
		}
	}
	return false
}

func (gm *GameMap) playerJump(onGround bool) {
	if gm.HasOneKeyDown(glfw.KeyUp, glfw.KeyW) && !gm.IsPressed(glfw.KeyUp, glfw.KeyW) && onGround && !gm.Player.HasTags("isDead") {
		gm.Player.isMove = true
		// 现在跳跃按键按下后重复跳跃
		if gm.HasOneKeyDown(glfw.KeyUp) {
			gm.PressedKey(glfw.KeyUp)
		}
		if gm.HasOneKeyDown(glfw.KeyW) {
			gm.PressedKey(glfw.KeyW)
		}
		gm.Player.SpeedY = -16
	}
}

func (gm *GameMap) playerMove(down Collision) {
	onGround := down.Colliding()
	friction := float32(0.01)
	if onGround {
		ground := down.ShapeB
		if ground.GetFriction() <= gm.Player.GetFriction() {
			friction = ground.GetFriction()
		} else {
			friction = gm.Player.GetFriction()
		}
	}
	accel := gm.Player.GetFriction() + friction

	if gm.Player.SpeedX > friction {
		gm.Player.SpeedX -= friction
	} else if gm.Player.SpeedX < -friction {
		gm.Player.SpeedX += friction
	} else {
		gm.Player.SpeedX = 0
	}

	if gm.HasOneKeyDown(glfw.KeyLeft, glfw.KeyRight, glfw.KeyA, glfw.KeyD) {
		gm.Player.isMove = true
	}

	if gm.HasOneKeyDown(glfw.KeyRight, glfw.KeyD) && onGround && !gm.Player.HasTags("isDead") {
		gm.Player.isXReverse = 1
		gm.Player.SpeedX += accel
	}

	if gm.HasOneKeyDown(glfw.KeyLeft, glfw.KeyA) && onGround && !gm.Player.HasTags("isDead") {
		gm.Player.isXReverse = -1
		gm.Player.SpeedX -= accel
	}

	//fmt.Println(gm.Player.SpeedX)
	if gm.Player.SpeedX > gm.Player.GetMaxSpd() {
		gm.Player.SpeedX = gm.Player.GetMaxSpd()
	}

	if gm.Player.SpeedX < -gm.Player.GetMaxSpd() {
		gm.Player.SpeedX = -gm.Player.GetMaxSpd()
	}
	//fmt.Println(gm.Player.SpeedX)
}

func (gm *GameMap) playerAttack(delta float64) {
	gm.Player.weapon.CoolDown(delta)

	// 调整角色攻击矢量
	if gm.Player.isXReverse < 0 {
		gm.Player.AtkVec[0] = -1
	} else {
		gm.Player.AtkVec[0] = 1
	}

	if gm.HasOneKeyDown(glfw.KeyUp, glfw.KeyW) {
		gm.Player.AtkVec[1] = -1
	} else if gm.HasOneKeyDown(glfw.KeyDown, glfw.KeyS) {
		gm.Player.AtkVec[1] = 1
	} else {
		gm.Player.AtkVec[1] = 0
	}

	if gm.HasOneKeyDown(glfw.KeyJ, glfw.KeySpace) && !gm.Player.HasTags("isDead") {
		bolt := gm.Player.Attack()
		if bolt != nil {
			gm.Map.Add(bolt)
		}
	}
}

func (gm *GameMap) updateMove() {
	move := gm.Map.FilterByTags("isMove")
	for i := 0; i < move.Length(); i++ {
		shape := move.Get(i)
		X, Y := shape.GetXY()
		x, y := shape.GetSpd()
		if res := gm.Map.FilterByTags("solid").Resolve(shape, int32(x), int32(y)); res.Colliding() {
			x = float32(res.ResolveX)
			y = float32(res.ResolveY)
			shape.SetSpd(x, y)
			shape.AddTags("destroy")
		}
		shape.SetXY(X+int32(x), Y+int32(y))
	}
}
