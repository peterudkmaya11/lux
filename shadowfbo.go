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
	this := ShadowFBO{}
	fbo := GenFramebuffer()
	this.width, this.height = width, height
	this.framebuffer = fbo
	fbo.Bind()
	defer fbo.Unbind()

	shadowtex := GenTexture2D()
	this.texture = shadowtex
	shadowtex.Bind(gl.TEXTURE_2D)
	defer shadowtex.Unbind(gl.TEXTURE_2D)
	shadowtex.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	shadowtex.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	shadowtex.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	shadowtex.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	shadowtex.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_MODE, gl.COMPARE_REF_TO_TEXTURE)
	shadowtex.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_COMPARE_FUNC, gl.LEQUAL)
	shadowtex.TexImage2D(0, gl.DEPTH_COMPONENT16, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)

	fbo.Texture(ReadDrawFramebuffer, DepthAttachement, shadowtex, 0 /*level*/)

	fbo.DrawBuffer(None)
	fbo.ReadBuffer(None)

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		return &this, errors.New("framebuffer incomplete")
	}

	vs, err := CompileShader(_shadow_fbo_vertex_shader, VertexShader)
	if err != nil {
		return &this, err
	}
	defer vs.Delete()
	fs, err := CompileShader(_shadow_fbo_fragment_shader, FragmentShader)
	if err != nil {
		return &this, err
	}
	defer fs.Delete()
	prog, err := NewProgram(vs, fs)
	if err != nil {
		return &this, err
	}
	this.program = prog
	prog.Use()
	this.mvpUni = prog.GetUniformLocation("mvp")

	return &this, nil
}

func (this *ShadowFBO) SetOrtho(left, right, bottom, top, near, far float32) {
	this.projection = glm.Ortho(left, right, bottom, top, near, far)
}

func (this *ShadowFBO) LookAt(ex, ey, ez, tx, ty, tz float32) {
	this.view = glm.LookAt(ex, ey, ez, tx, ty, tz, 0, 1, 0)
}

func (this *ShadowFBO) BindForDrawing() {
	this.framebuffer.Bind()
	this.program.Use()
	this.vp = this.projection.Mul4(this.view)
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	ViewPortChange(this.width, this.height)
	gl.CullFace(gl.FRONT)
}

func (this *ShadowFBO) Unbind() {
	gl.CullFace(gl.BACK)
}

func (this *ShadowFBO) Render(mesh Mesh, transform *Transform) {
	mvpmat := this.vp.Mul4(transform.Mat4())
	this.mvpUni.UniformMatrix4fv(1, false, &mvpmat[0])
	mesh.Bind()
	mesh.DrawCall()
}

func (this *ShadowFBO) ShadowMap() Texture2D {
	return this.texture
}

func (this *ShadowFBO) ShadowMat() glm.Mat4 {
	return depthscaling.Mul4(this.vp)
}

func (this *ShadowFBO) Delete() {
	this.texture.Delete()
	this.framebuffer.Delete()
	this.program.Delete()
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
