package noname

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
func (this RenderBuffer) Bind() {
	gl.BindRenderbuffer(gl.RENDERBUFFER, uint32(this))
}

//Alias for glBindRenderbuffer(gl.RENDERBUFFER, 0)
func (this RenderBuffer) Unbind() {
	gl.BindRenderbuffer(gl.RENDERBUFFER, 0)
}

//Alias for glDeleteRenderbuffers(1, (*uint32)(&this))
func (this RenderBuffer) Delete() {
	gl.DeleteRenderbuffers(1, (*uint32)(&this))
}

func (this RenderBuffer) Storage(internalformat uint32, width, height int32) {
	//RENDERBUFFER is the only possible value
	gl.RenderbufferStorage(gl.RENDERBUFFER, internalformat, width, height)
}
