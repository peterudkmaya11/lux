package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	gl2 "github.com/luxengine/gl"
)

//GenDepthTexture is a utility function to generate a depth Texture2D
func GenDepthTexture(width, height int32) gl2.Texture2D {
	tex := gl2.GenTexture2D()
	tex.Bind()
	defer tex.Unbind()
	tex.MinFilter(gl2.NEAREST)
	tex.MagFilter(gl2.NEAREST)
	tex.WrapS(gl2.CLAMP_TO_EDGE)
	tex.WrapT(gl2.CLAMP_TO_EDGE)
	tex.TexImage2D(0, gl.DEPTH_COMPONENT, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	return tex
}
