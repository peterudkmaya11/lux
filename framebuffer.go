package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Framebuffer uint32

// The framebuffer, which regroups 0, 1, or more textures, and 0 or 1 depth buffer.
func GenFramebuffer() Framebuffer {
	var fb uint32
	gl.GenFramebuffers(1, &fb)
	return Framebuffer(fb)
}

//the target may only be GL_FRAMEBUFFER
func (this Framebuffer) Bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, uint32(this))
}

func (this Framebuffer) Unbind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (this Framebuffer) RenderBuffer(attachement FramebufferAttachement, renderbuffer RenderBuffer) {
	//FRAMEBUFFER and RENDERBUFFER are the only possible values
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, uint32(attachement), gl.RENDERBUFFER, uint32(renderbuffer))
}

//Must be called with a bound Framebuffer
func (this Framebuffer) Texture(target FramebufferTarget, attachement FramebufferAttachement, texture Texture, level int32) {
	gl.FramebufferTexture(uint32(target), uint32(attachement), uint32(texture), level)
}

//Must be called with a bound Framebuffer
func (this Framebuffer) DrawBuffers(attachements ...FramebufferAttachement) {
	gl.DrawBuffers(int32(len(attachements)), (*uint32)(&attachements[0]))
}

//Must be called with a bound Framebuffer
func (this Framebuffer) DrawBuffer(attachements FramebufferAttachement) {
	gl.DrawBuffer(uint32(attachements))
}

//Must be called with a bound Framebuffer
func (this Framebuffer) ReadBuffer(attachements FramebufferAttachement) {
	gl.ReadBuffer(uint32(attachements))
}

func (this Framebuffer) Delete() {
	gl.DeleteFramebuffers(1, (*uint32)(&this))
}
