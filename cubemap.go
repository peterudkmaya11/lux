package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type CubeMap uint32

func GenCubeMap() CubeMap {
	var tex uint32
	gl.GenTextures(1, &tex)
	return CubeMap(tex)
}
