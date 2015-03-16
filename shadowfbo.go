package lux

import (
	"errors"
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
)

type ShadowFBO struct {
	framebuffer          Framebuffer
	texture              Texture2D
	projection, view, vp glm.Mat4
	program              Program
	mvpUni               UniformLocation
	width, height        int32
}

func NewShadowFBO(width, height int32) (*ShadowFBO, error) {
	sfbo := ShadowFBO{}
	fbo := GenFramebuffer()
	sfbo.width, sfbo.height = width, height
	sfbo.framebuffer = fbo
	fbo.Bind()
	defer fbo.Unbind()

	shadowtex := GenTexture2D()
	sfbo.texture = shadowtex
	shadowtex.Bind()
	defer shadowtex.Unbind()
	shadowtex.TexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	shadowtex.TexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	shadowtex.TexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	shadowtex.TexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	shadowtex.TexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_MODE, gl.COMPARE_REF_TO_TEXTURE)
	shadowtex.TexParameteriv(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_FUNC, gl.LEQUAL)
	shadowtex.TexImage2D(0, gl.DEPTH_COMPONENT16, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)

	fbo.Texture(ReadDrawFramebuffer, DepthAttachement, shadowtex, 0 /*level*/)

	fbo.DrawBuffer(None)
	fbo.ReadBuffer(None)

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		return &sfbo, errors.New("framebuffer incomplete")
	}

	vs, err := CompileShader(_shadow_fbo_vertex_shader, VertexShader)
	if err != nil {
		return &sfbo, err
	}
	defer vs.Delete()
	fs, err := CompileShader(_shadow_fbo_fragment_shader, FragmentShader)
	if err != nil {
		return &sfbo, err
	}
	defer fs.Delete()
	prog, err := NewProgram(vs, fs)
	if err != nil {
		return &sfbo, err
	}
	sfbo.program = prog
	prog.Use()
	sfbo.mvpUni = prog.GetUniformLocation("mvp")

	return &sfbo, nil
}

func (sfbo *ShadowFBO) SetOrtho(left, right, bottom, top, near, far float32) {
	sfbo.projection = glm.Ortho(left, right, bottom, top, near, far)
}

func (sfbo *ShadowFBO) LookAt(ex, ey, ez, tx, ty, tz float32) {
	sfbo.view = glm.LookAt(ex, ey, ez, tx, ty, tz, 0, 1, 0)
}

func (sfbo *ShadowFBO) BindForDrawing() {
	sfbo.framebuffer.Bind()
	sfbo.program.Use()
	sfbo.vp = sfbo.projection.Mul4(sfbo.view)
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	ViewPortChange(sfbo.width, sfbo.height)
	gl.CullFace(gl.FRONT)
}

func (sfbo *ShadowFBO) Unbind() {
	gl.CullFace(gl.BACK)
}

func (sfbo *ShadowFBO) Render(mesh Mesh, transform *Transform) {
	mvpmat := sfbo.vp.Mul4(transform.Mat4())
	sfbo.mvpUni.UniformMatrix4fv(1, false, &mvpmat[0])
	mesh.Bind()
	mesh.DrawCall()
}

func (sfbo *ShadowFBO) ShadowMap() Texture2D {
	return sfbo.texture
}

func (sfbo *ShadowFBO) ShadowMat() glm.Mat4 {
	return depthscaling.Mul4(sfbo.vp)
}

func (sfbo *ShadowFBO) Delete() {
	sfbo.texture.Delete()
	sfbo.framebuffer.Delete()
	sfbo.program.Delete()
}

var _shadow_fbo_vertex_shader = `
#version 330 core
uniform mat4 mvp;

layout(location = 0) in vec3 position;
 
void main(){
	gl_Position =  mvp * vec4(position,1);
}
` + "\x00"

var _shadow_fbo_fragment_shader = `
#version 330 core

layout(location=0) out float x;

void main(){
	x=1;
}
` + "\x00"

var depthscaling = glm.Mat4{0.5, 0, 0, 0,
	0, 0.5, 0, 0,
	0, 0, 0.5, 0,
	0.5, 0.5, 0.5, 1}
