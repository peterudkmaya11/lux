package lux

import (
	"errors"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

//the post process fullscreen quad mesh

var fstri struct {
	vao     VertexArray
	pos, uv Buffer
}

func InitPostProcessSystem() {
	//init model
	quadvao := GenVertexArray()
	quadvao.Bind()
	defer quadvao.Unbind()

	vertpos := GenBuffer()
	vertpos.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(_quadvertpos)*2*4, ptr(_quadvertpos), gl.STATIC_DRAW)

	vertuv := GenBuffer()
	vertuv.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(_quadvertuv)*2*4, ptr(_quadvertuv), gl.STATIC_DRAW)

	fstri.vao = quadvao
	fstri.pos = vertpos
	fstri.uv = vertuv
}

type PostProcessFramebufferer interface {
	PreRender()
	Render(Texture2D)
	SetNext(PostProcessFramebufferer)
}

type PostProcessFramebuffer struct {
	Fb           Framebuffer
	Tex          Texture2D
	Prog         Program
	next         PostProcessFramebufferer
	time, source UniformLocation
}

//post process framebuffers are a tool to queue fullscreen fragment shader passes
func NewPostProcessFramebuffer(width, height int32, fragmentSource string) (*PostProcessFramebuffer, error) {
	ppf := PostProcessFramebuffer{}
	vs, err := CompileShader(_fullscreen_vertex_shader, VertexShader)
	if err != nil {
		return &ppf, err
	}
	defer vs.Delete()
	fs, err := CompileShader(fragmentSource, FragmentShader)
	if err != nil {
		return &ppf, err
	}
	defer fs.Delete()
	prog, err := NewProgram(vs, fs)
	if err != nil {
		return &ppf, err
	}
	ppf.Prog = prog
	prog.Use()
	res := prog.GetUniformLocation("resolution")
	ppf.source = prog.GetUniformLocation("tex")

	res.Uniform2f(float32(width), float32(height))

	ppf.time = prog.GetUniformLocation("time")

	ppf.Fb = GenFramebuffer()
	ppf.Tex = GenRGBTexture2D(width, height)

	ppf.Fb.Bind()
	defer ppf.Fb.Unbind()

	ppf.Fb.Texture(ReadDrawFramebuffer, ColorAttachement0, ppf.Tex, 0)
	ppf.Fb.DrawBuffers(ColorAttachement0)

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		return &ppf, errors.New("framebuffer incomplete")
	}
	return &ppf, nil
}

func (ppfb *PostProcessFramebuffer) PreRender() {
	if ppfb.next != nil {
		ppfb.Fb.Bind()
	} else {
		gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	}
	gl.Disable(gl.DEPTH_TEST)
}

func (ppfb *PostProcessFramebuffer) PostRender() {
	gl.Enable(gl.DEPTH_TEST)
}

//should be called a 'pass'?
func (ppfb *PostProcessFramebuffer) Render(t Texture2D) {
	ppfb.Prog.Use()
	ppfb.time.Uniform1f(float32(glfw.GetTime()))

	gl.ActiveTexture(TextureUnitDiffuse)
	t.Bind()
	ppfb.source.Uniform1i(TextureUniformDiffuse)
	Fstri()
	if ppfb.next != nil {
		ppfb.next.PreRender()
		ppfb.next.Render(ppfb.Tex)
	}
}

func (ppfb *PostProcessFramebuffer) SetNext(n PostProcessFramebufferer) {
	ppfb.next = n
}

func (ppfb *PostProcessFramebuffer) Delete() {
	ppfb.Prog.Delete()
	ppfb.Tex.Delete()
	ppfb.Fb.Delete()
}

func Fstri() {
	fstri.vao.Bind()
	defer fstri.vao.Unbind()
	gl.EnableVertexAttribArray(0)
	fstri.pos.Bind(gl.ARRAY_BUFFER)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	fstri.uv.Bind(gl.ARRAY_BUFFER)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

var _fullscreen_vertex_shader = `
#version 330
layout (location=0) in vec2 vert;
layout (location=1) in vec2 vertTexCoord;
out vec2 uv;
void main() {
	uv = vertTexCoord;
	gl_Position = vec4(vert,0,1);
}
` + "\x00"

//technicly a fullscreen triangle now
var _quadvertpos = []float32{3, 1, -1, 1, -1, -3}
var _quadvertuv = []float32{2, 1, 0, 1, 0, -1}

type UIFramebuffer struct {
	PostProcessFramebuffer
}
