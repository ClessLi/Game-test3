package main

import (
	"github.com/ClessLi/Game-test3/game"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

const (
	width  = 800
	height = 600
)

var (
	windowName = "Test Game"
	section    = 0
	game2D     = game.NewC1(width, height)
	deltaTime  = 0.0
	lastFrame  = 0.0
)

func main() {
	runtime.LockOSThread()
	window := initGlfw()
	defer glfw.Terminate()
	initOpenGL()
	for _, m := range game2D.Maps {
		if m != nil {
			m.Create()
		}
	}
	//game2D.Init()
	for !window.ShouldClose() {
		secMap := game2D.Maps[section]
		currentFrame := glfw.GetTime()
		deltaTime = currentFrame - lastFrame
		lastFrame = currentFrame

		glfw.PollEvents()
		//secMap.ProcessInput(float32(deltaTime))
		secMap.Update(deltaTime)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		secMap.Draw()
		window.SwapBuffers()
	}
}

func initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.Viewport(0, 0, width, height)
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.BLEND)
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := glfw.CreateWindow(width, height, windowName, nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetKeyCallback(KeyCallback)

	window.MakeContextCurrent()
	return window
}

func KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch action {
	case glfw.Press:
		game2D.Maps[section].SetKeyDown(key)
	case glfw.Release:
		game2D.Maps[section].ReleaseKey(key)
	}
}
