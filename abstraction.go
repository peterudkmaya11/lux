package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//Deletable is the interface to represent everything that can be unallocated.
type Deletable interface {
	Delete()
}

//VertexArray is the high level representation of openGL VertexArray.
type VertexArray uint32

//GenVertexArray generates 1 vertex array. Alias for glGenVertexArrays(1, &va).
func GenVertexArray() VertexArray {
	var va uint32
	gl.GenVertexArrays(1, &va)
	return VertexArray(va)
}

//Bind this vertex array, Alias to gl.BindVertexArray(uint32(va)).
func (va VertexArray) Bind() {
	gl.BindVertexArray(uint32(va))
}

//Unbind bind vertex array 0(zero). Alias to gl.BindVertexArray(0).
func (va VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

//Delete this vertex array.
func (va VertexArray) Delete() {
	gl.DeleteVertexArrays(1, (*uint32)(&va))
}

//Buffer is the high level representation of OpenGL Buffer.
//TODO: make subtype buffers with restrained functions
//TODO: make an enum for targets
type Buffer uint32

//GenBuffer generates 1 Buffer.
func GenBuffer() Buffer {
	var va uint32
	gl.GenBuffers(1, &va)
	return Buffer(va)
}

//Bind this buffer, requires target. Alias to gl.BindBuffer(target, uint32(b)).
func (b Buffer) Bind(target uint32) {
	gl.BindBuffer(target, uint32(b))
}

//Unbind binds buffer 0(zero). Alias to gl.BindBuffer(target, 0).
func (b Buffer) Unbind(target uint32) {
	gl.BindBuffer(target, 0)
}

//Delete this buffer
func (b Buffer) Delete() {
	gl.DeleteBuffers(1, (*uint32)(&b))
}

//GetCurrentTexture2D generate a single texture2D.
func GetCurrentTexture2D() Texture2D {
	var i int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &i)
	return Texture2D(i)
}
