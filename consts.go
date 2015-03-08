package noname

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type ShaderType uint32

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

const (
	TextureUnitDiffuse      = gl.TEXTURE0
	TextureUnitNormalMap    = gl.TEXTURE1
	TextureUnitDisplacement = gl.TEXTURE2
)

const (
	TextureUniformDiffuse      = 0
	TextureUniformNormalMap    = 1
	TextureUniformDisplacement = 2
)

type FramebufferTarget uint32

const (
	DrawFramebuffer     FramebufferTarget = gl.DRAW_FRAMEBUFFER
	ReadFramebuffer                       = gl.READ_FRAMEBUFFER
	ReadDrawFramebuffer                   = gl.FRAMEBUFFER
)

//
type FramebufferAttachement uint32

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
