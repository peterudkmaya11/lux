package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

type Texture uint32

func (this Texture) Bind(target uint32) {
	gl.BindTexture(target, uint32(this))
}

func (this Texture) Unbind(target uint32) {
	gl.BindTexture(target, 0)
}

func (this Texture) Delete() {
	gl.DeleteTextures(1, (*uint32)(&this))
}

func (this Texture) TexImage2D(level, internalformat, width, height, border int32, format, xtype uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(gl.TEXTURE_2D, level, internalformat, width, height, border, format, xtype, pixels)
}

func (this Texture) TexParameteri(target, pname uint32, param int32) {
	gl.TexParameteri(target, pname, param)
}

func GenTexture() Texture {
	var tex uint32
	gl.GenTextures(1, &tex)
	return Texture(tex)
}

func GenRGBTexture(width, height int32) Texture {
	tex := GenTexture()
	tex.Bind(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, width, height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	return tex
}

//utility

func HeightTexture(width, height int32, data []float32) Texture {
	tex := GenTexture()
	tex.Bind(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH_COMPONENT32F, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, gl.Ptr(data))
	return tex
}
