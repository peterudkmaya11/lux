package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//ShaderType is a high level representation of OpenGL shader types.
type ShaderType uint32

//All Shader types
const (
	VertexShader   ShaderType = gl.VERTEX_SHADER
	FragmentShader            = gl.FRAGMENT_SHADER
	GeometryShader            = gl.GEOMETRY_SHADER
)

const (
	layoutVertexPosition = 0
	layoutVertexUv       = 1
	layoutVertexNormal   = 2
)

//lux standard Active texture layout.
const (
	TextureUnitDiffuse      = gl.TEXTURE0
	TextureUnitNormalMap    = gl.TEXTURE1
	TextureUnitDisplacement = gl.TEXTURE2
)

//lux standard active texture layout.
const (
	TextureUniformDiffuse      = 0
	TextureUniformNormalMap    = 1
	TextureUniformDisplacement = 2
)

//FramebufferTarget is a high level representation of OpenGL framebuffer targets.
type FramebufferTarget uint32

//All possible framebuffer target.
const (
	DrawFramebuffer     FramebufferTarget = gl.DRAW_FRAMEBUFFER
	ReadFramebuffer                       = gl.READ_FRAMEBUFFER
	ReadDrawFramebuffer                   = gl.FRAMEBUFFER
)

//FramebufferAttachement is a high level representation fo OpenGL framebuffer attachements
type FramebufferAttachement uint32

//all possible framebuffer attachement.
const (
	ColorAttachement0       FramebufferAttachement = gl.COLOR_ATTACHMENT0
	ColorAttachement1                              = gl.COLOR_ATTACHMENT1
	ColorAttachement2                              = gl.COLOR_ATTACHMENT2
	ColorAttachement3                              = gl.COLOR_ATTACHMENT3
	ColorAttachement4                              = gl.COLOR_ATTACHMENT4
	ColorAttachement5                              = gl.COLOR_ATTACHMENT5
	ColorAttachement6                              = gl.COLOR_ATTACHMENT6
	ColorAttachement7                              = gl.COLOR_ATTACHMENT7
	ColorAttachement8                              = gl.COLOR_ATTACHMENT8
	ColorAttachement9                              = gl.COLOR_ATTACHMENT9
	ColorAttachement10                             = gl.COLOR_ATTACHMENT10
	ColorAttachement11                             = gl.COLOR_ATTACHMENT11
	ColorAttachement12                             = gl.COLOR_ATTACHMENT12
	ColorAttachement13                             = gl.COLOR_ATTACHMENT13
	ColorAttachement14                             = gl.COLOR_ATTACHMENT14
	ColorAttachement15                             = gl.COLOR_ATTACHMENT15
	DepthAttachement                               = gl.DEPTH_ATTACHMENT
	StencilAttachement                             = gl.STENCIL_ATTACHMENT
	DepthStencilAttachement                        = gl.DEPTH_STENCIL_ATTACHMENT
	None                                           = gl.NONE
)

/*
'standard' supported texture:
diffuse
normal
displacement
shadow

acknowledged but not supported:
light

anything else is special and will have to be manually coded
*/
