package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

//texture interface maybe ?
/*
type Texture interface{
	Bind()
	Unbind()
	Delete()
	//think about parameters
	//Parameteri()
}
*/

type Texture uint32
type Texture2D Texture

func (this Texture2D) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, uint32(this))
}

func (this Texture2D) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (this Texture2D) Delete() {
	gl.DeleteTextures(1, (*uint32)(&this))
}

func (this Texture2D) TexImage2D(level, internalformat, width, height, border int32, format, xtype uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(gl.TEXTURE_2D, level, internalformat, width, height, border, format, xtype, pixels)
}

func (this Texture2D) TexParameteri(target, pname uint32, param int32) {
	gl.TexParameteri(target, pname, param)
}

func GenTexture2D() Texture2D {
	var tex uint32
	gl.GenTextures(1, &tex)
	return Texture2D(tex)
}

func GenRGBTexture2D(width, height int32) Texture2D {
	tex := GenTexture2D()
	tex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, width, height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	return tex
}
