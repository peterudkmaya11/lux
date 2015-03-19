package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//CubeMap is a high level representation of OpenGL cubemaps. Not really implemented right now.
type CubeMap uint32

//GenCubeMap is an alias to glGenTextures(1, &tex) then cast to
func GenCubeMap() CubeMap {
	var tex uint32
	gl.GenTextures(1, &tex)
	return CubeMap(tex)
}
