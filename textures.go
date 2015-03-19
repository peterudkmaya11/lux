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

//Texture is a high level representation of the OpenGL texture object, can be any type (TEXTURE_2D, TEXTURE_1D, etc)
type Texture uint32

//Texture2D is the high level representation of OpenGL TEXTURE_2D object. it restrict availlable functions and automatically fills the GL_TEXTURE_2D target.
type Texture2D Texture

//Bind is an alias to gl.BindTexture(gl.TEXTURE_2D, uint32(t))
func (t Texture2D) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, uint32(t))
}

//Unbind is an alias to gl.BindTexture(gl.TEXTURE_2D, 0)
func (t Texture2D) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

//Delete is an alias to gl.DeleteTextures(1, (*uint32)(&t))
func (t Texture2D) Delete() {
	gl.DeleteTextures(1, (*uint32)(&t))
}

//TexImage2D is an alias to gl.TexImage2D(gl.TEXTURE_2D, level, internalformat, width, height, border, format, xtype, pixels)
func (t Texture2D) TexImage2D(level, internalformat, width, height, border int32, format, xtype uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(gl.TEXTURE_2D, level, internalformat, width, height, border, format, xtype, pixels)
}

//TexParameteriv is an alias to gl.TexParameteri(target, pname, param)
func (t Texture2D) TexParameteriv(target, pname uint32, param int32) {
	gl.TexParameteri(target, pname, param)
}

//GetTexParameteriv is an alias to gl.GetTexParameteriv(target, pname, params)
func (t Texture2D) GetTexParameteriv(target, pname uint32, params *int32) {
	gl.GetTexParameteriv(target, pname, params)
}

//GetTexLevelParameteriv is an alias to gl.GetTexLevelParameteriv(target, level, pname, params)
func (t Texture2D) GetTexLevelParameteriv(target uint32, level int32, pname uint32, params *int32) {
	gl.GetTexLevelParameteriv(target, level, pname, params)
}

//GetTexImage is an alias to gl.GetTexImage(gl.TEXTURE_2D, level, format, xtype, pixels)
func (t Texture2D) GetTexImage(level int32, format, xtype uint32, pixels unsafe.Pointer) {
	gl.GetTexImage(gl.TEXTURE_2D, level, format, xtype, pixels)
}

//ReadPixels is an alias to gl.ReadPixels(x, y, width, height, format, xtype, pixels)
func (t Texture2D) ReadPixels(x, y, width, height int32, format, xtype uint32, pixels unsafe.Pointer) {
	gl.ReadPixels(x, y, width, height, format, xtype, pixels)
}

//Width is an alias to gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_WIDTH, &w)
func (t Texture2D) Width(miplevel int32) int32 {
	var w int32
	gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_WIDTH, &w)
	return w
}

//Height is an alias to gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_HEIGHT, &h)
func (t Texture2D) Height(miplevel int32) int32 {
	var h int32
	gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_HEIGHT, &h)
	return h
}

//InternalFormat is an alias to gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_INTERNAL_FORMAT, &x)
func (t Texture2D) InternalFormat(miplevel int32) uint32 {
	var x int32
	gl.GetTexLevelParameteriv(gl.TEXTURE_2D, miplevel, gl.TEXTURE_INTERNAL_FORMAT, &x)
	return uint32(x)
}

//GenTexture2D generate a single Texture2D via gl.GenTextures(1,...).
func GenTexture2D() Texture2D {
	var tex uint32
	gl.GenTextures(1, &tex)
	return Texture2D(tex)
}

//GenRGBTexture2D is a utility function to generate an empty 2D textures of size (width,height), internal format RGB adn data type UNSIGNED_BYTE
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
