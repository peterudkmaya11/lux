package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//RenderBuffer is the high level representation of OpenGL RenderBuffer
type RenderBuffer uint32

//GenRenderBuffer generates one renderbuffer. Alias to glGenRenderbuffers(1, &buf)
func GenRenderBuffer() RenderBuffer {
	var buf uint32
	gl.GenRenderbuffers(1, &buf)
	return RenderBuffer(buf)
}

//Bind is an alias for glBindRenderbuffer(gl.RENDERBUFFER, uint32(rb))
func (rb RenderBuffer) Bind() {
	gl.BindRenderbuffer(gl.RENDERBUFFER, uint32(rb))
}

//Unbind is an alias for glBindRenderbuffer(gl.RENDERBUFFER, 0)
func (rb RenderBuffer) Unbind() {
	gl.BindRenderbuffer(gl.RENDERBUFFER, 0)
}

//Delete is an alias for glDeleteRenderbuffers(1, (*uint32)(&rb))
func (rb RenderBuffer) Delete() {
	gl.DeleteRenderbuffers(1, (*uint32)(&rb))
}

//Storage set the width,height and internal format of this RenderBuffer
func (rb RenderBuffer) Storage(internalformat uint32, width, height int32) {
	//RENDERBUFFER is the only possible value
	gl.RenderbufferStorage(gl.RENDERBUFFER, internalformat, width, height)
}
