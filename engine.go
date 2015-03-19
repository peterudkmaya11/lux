package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//ViewportChange is an alias to glViewport(0, 0, width, height)
func ViewportChange(width, height int32) {
	gl.Viewport(0, 0, width, height)
}
