package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//Framebuffer is a high level representation of OpenGL framebuffer object.
type Framebuffer uint32

//GenFramebuffer is an alias to glGenFramebuffers(1, &fb)
func GenFramebuffer() Framebuffer {
	var fb uint32
	gl.GenFramebuffers(1, &fb)
	return Framebuffer(fb)
}

//Bind binds this framebuffer. Alias to glBindFramebuffer(gl.FRAMEBUFFER, fb)
func (fb Framebuffer) Bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, uint32(fb))
}

//Unbind binds the 0 framebuffer. Alias to glBindFramebuffer(gl.FRAMEBUFFER, 0)
func (fb Framebuffer) Unbind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

//RenderBuffer attaches a Render to this Framebuffer object. Alias for glFramebufferRenderbuffer(gl.FRAMEBUFFER, attachement, gl.RENDERBUFFER, renderbuffer)
func (fb Framebuffer) RenderBuffer(attachement FramebufferAttachement, renderbuffer RenderBuffer) {
	//FRAMEBUFFER and RENDERBUFFER are the only possible values
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, uint32(attachement), gl.RENDERBUFFER, uint32(renderbuffer))
}

//Texture attach a Texture2D. Alias to glFramebufferTexture(target, attachement, texture, level)
func (fb Framebuffer) Texture(target FramebufferTarget, attachement FramebufferAttachement, texture Texture2D, level int32) {
	gl.FramebufferTexture(uint32(target), uint32(attachement), uint32(texture), level)
}

//DrawBuffers attachs all the FramebufferAttachement to this FBO. Alias for glDrawBuffers(len(attachements), &attachements[0])
func (fb Framebuffer) DrawBuffers(attachements ...FramebufferAttachement) {
	gl.DrawBuffers(int32(len(attachements)), (*uint32)(&attachements[0]))
}

//DrawBuffer attach a single FramebufferAttachement. Alias for glDrawBuffer(attachement)
func (fb Framebuffer) DrawBuffer(attachement FramebufferAttachement) {
	gl.DrawBuffer(uint32(attachement))
}

//ReadBuffer is as alias to glReadBuffer(attachement)
func (fb Framebuffer) ReadBuffer(attachement FramebufferAttachement) {
	gl.ReadBuffer(uint32(attachement))
}

//Delete is an alias to glDeleteFramebuffers(1, &fb)
func (fb Framebuffer) Delete() {
	gl.DeleteFramebuffers(1, (*uint32)(&fb))
}
