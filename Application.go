package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

//obvious ryze quote
func letsglowLETSGLOW() {
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}
}

func SetContext() {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    // Necessary for OS X
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // Necessary for OS X
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)
}

func glbs() {
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.3, 0.3, 0.3, 1.0)
	gl.Enable(gl.CULL_FACE)
}

//Will call glfw.Init and panic if it fails
func InitGLFW() {
	if err := glfw.Init(); err != nil {
		panic("failed to initialize glfw")
	}
}

//Alias to glfw.Terminate
func TerminateGLFW() {
	glfw.Terminate()
}

//Create a new glfw window. If fullscreen will place the screen on primary monitor.
func CreateWindow(width, height int, title string, fullscreen bool) (window *glfw.Window) {
	SetContext()
	var x *glfw.Monitor
	if fullscreen {
		x = glfw.GetPrimaryMonitor()
	}
	window, err := glfw.CreateWindow(width, height, title, x, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	letsglowLETSGLOW()
	glbs()
	QueryExtentions()
	return
}
