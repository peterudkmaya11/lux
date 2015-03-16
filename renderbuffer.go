package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type RenderBuffer uint32

//Generates one renderbuffer. You cannot use a renderbuffer as texture in shader, use texture for that :P
func GenRenderBuffer() RenderBuffer {
	var buf uint32
	gl.GenRenderbuffers(1, &buf)
	return RenderBuffer(buf)
}

//Alias for glBindRenderbuffer(gl.RENDERBUFFER, uint32(this))
func (rb RenderBuffer) Bind() {
	gl.BindRenderbuffer(gl.RENDERBUFFER, uint32(rb))
}

//Alias for glBindRenderbuffer(gl.RENDERBUFFER, 0)
func (rb RenderBuffer) Unbind() {
	gl.BindRenderbuffer(gl.RENDERBUFFER, 0)
}

//Alias for glDeleteRenderbuffers(1, (*uint32)(&rb))
func (rb RenderBuffer) Delete() {
	gl.DeleteRenderbuffers(1, (*uint32)(&rb))
}

func (rb RenderBuffer) Storage(internalformat uint32, width, height int32) {
	//RENDERBUFFER is the only possible value
	gl.RenderbufferStorage(gl.RENDERBUFFER, internalformat, width, height)
}
