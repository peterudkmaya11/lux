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

type DrawMode uint32

const (
	POINTS                   DrawMode = gl.POINTS
	LINE_STRIP                        = gl.LINE_STRIP
	LINE_LOOP                         = gl.LINE_LOOP
	LINES                             = gl.LINES
	LINE_STRIP_ADJACENCY              = gl.LINE_STRIP_ADJACENCY
	LINES_ADJACENCY                   = gl.LINES_ADJACENCY
	TRIANGLE_STRIP                    = gl.TRIANGLE_STRIP
	TRIANGLE_FAN                      = gl.TRIANGLE_FAN
	TRIANGLES                         = gl.TRIANGLES
	TRIANGLE_STRIP_ADJACENCY          = gl.TRIANGLE_STRIP_ADJACENCY
	TRIANGLES_ADJACENCY               = gl.TRIANGLES_ADJACENCY
	PATCHES                           = gl.PATCHES
)

type BufferTarget uint32

const (
	ARRAY_BUFFER              BufferTarget = gl.ARRAY_BUFFER              //OpenGL 2+
	ELEMENT_ARRAY_BUFFER                   = gl.ELEMENT_ARRAY_BUFFER      //OpenGL 2+
	PIXEL_PACK_BUFFER                      = gl.PIXEL_PACK_BUFFER         //OpenGL 2+
	PIXEL_UNPACK_BUFFER                    = gl.PIXEL_UNPACK_BUFFER       //OpenGL 2+
	COPY_READ_BUFFER                       = gl.COPY_READ_BUFFER          //OpenGL 3+
	COPY_WRITE_BUFFER                      = gl.COPY_WRITE_BUFFER         //OpenGL 3+
	TEXTURE_BUFFER                         = gl.TEXTURE_BUFFER            //OpenGL 3+
	TRANSFORM_FEEDBACK_BUFFER              = gl.TRANSFORM_FEEDBACK_BUFFER //OpenGL 3+
	UNIFORM_BUFFER                         = gl.UNIFORM_BUFFER            //OpenGL 3+
	ATOMIC_COUNTER_BUFFER                  = gl.ATOMIC_COUNTER_BUFFER     //OpenGL 4+
	DRAW_INDIRECT_BUFFER                   = gl.DRAW_INDIRECT_BUFFER      //OpenGL 4+
	DISPATCH_INDIRECT_BUFFER               = gl.DISPATCH_INDIRECT_BUFFER  //OpenGL 4+
	QUERY_BUFFER                           = gl.QUERY_BUFFER              //OpenGL 4+
	SHADER_STORAGE_BUFFER                  = gl.SHADER_STORAGE_BUFFER     //OpenGL 4+
)

/* = gl.'standard' supported texture:
diffuse
normal
displacement
shadow

acknowledged but not supported:
light

anything else is special and will have to be manually coded
*/
