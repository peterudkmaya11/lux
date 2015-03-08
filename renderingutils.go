package noname

import (
	"io/ioutil"
)

type RenderProgram struct {
	Prog                            Program
	M, V, P, Diffuse, Light, N, Eye UniformLocation
}

//Load a vertex-fragment program and gathers:
//"M": model matrix uniform
//"V": view matrix uniform
//"P": projection matrix uniform
//"N": normal matrix uniform
//"diffuse":diffuse texture sampler2d
//"pointlight":array of vec3 for light position
func LoadProgram(vertexfile, fragfile string) (out RenderProgram, err error) {
	vssource, err := ioutil.ReadFile(vertexfile)
	if err != nil {
		return
	}
	fssource, err := ioutil.ReadFile(fragfile)
	if err != nil {
		return
	}
	vs, err := CompileShader(string(vssource)+"\x00", VertexShader)
	defer vs.Delete()
	if err != nil {
		return
	}
	fs, err := CompileShader(string(fssource)+"\x00", FragmentShader)
	defer fs.Delete()
	if err != nil {
		return
	}
	p, err := NewProgram(vs, fs)
	if err != nil {
		return
	}
	out.Prog = p
	out.M = p.GetUniformLocation("M")
	out.V = p.GetUniformLocation("V")
	out.P = p.GetUniformLocation("P")
	out.N = p.GetUniformLocation("N")
	out.Diffuse = p.GetUniformLocation("diffuse")
	out.Light = p.GetUniformLocation("pointlight")
	out.Eye = p.GetUniformLocation("eye")
	return
}

func (this *RenderProgram) Delete() {
	this.Prog.Delete()
}

/*
package gogine

import (
	"github.com/go-gl/examples/mathgl/opengl-tutorial/helper"
	"github.com/go-gl/gl"
	glm "github.com/go-gl/mathgl/mgl32"
)

type GLProgram struct {
	MVP, Diffuse            gl.UniformLocation
	Positions, Uvs, Normals gl.AttribLocation
	Prog                    gl.Program
}

func NewGLProgram(vertex, fragment string) GLProgram {
	out := GLProgram{}
	out.Prog = helper.MakeProgram(vertex, fragment)
	out.MVP = out.Prog.GetUniformLocation("MVP")
	out.Positions = out.Prog.GetAttribLocation("position")
	out.Uvs = out.Prog.GetAttribLocation("uv")
	out.Normals = out.Prog.GetAttribLocation("normal")
	out.Diffuse = out.Prog.GetUniformLocation("diffuse") //rename to diffuse
	return out
}

func (this *GLProgram) Render(mvp glm.Mat4, m Model, diffuse gl.Texture) {
	this.Prog.Use()
	defer gl.ProgramUnuse()

	gl.ActiveTexture(gl.TEXTURE0)
	diffuse.Bind(gl.TEXTURE_2D)
	defer diffuse.Unbind(gl.TEXTURE_2D)
	this.Diffuse.Uniform1i(0)

	this.Positions.EnableArray()
	defer this.Positions.DisableArray()
	m.Positions.Bind(gl.ARRAY_BUFFER)
	defer m.Positions.Unbind(gl.ARRAY_BUFFER)
	this.Positions.AttribPointer(3, gl.FLOAT, false, 0, nil)

	this.Uvs.EnableArray()
	defer this.Uvs.DisableArray()
	m.Uvs.Bind(gl.ARRAY_BUFFER)
	defer m.Uvs.Unbind(gl.ARRAY_BUFFER)
	this.Uvs.AttribPointer(2, gl.FLOAT, false, 0, nil)

	this.Normals.EnableArray()
	defer this.Normals.DisableArray()
	m.Normals.Bind(gl.ARRAY_BUFFER)
	defer m.Normals.Unbind(gl.ARRAY_BUFFER)
	this.Normals.AttribPointer(3, gl.FLOAT, false, 0, nil)

	m.Indices.Bind(gl.ELEMENT_ARRAY_BUFFER)
	defer m.Indices.Unbind(gl.ELEMENT_ARRAY_BUFFER)

	this.MVP.UniformMatrix4fv(false, mvp)
	gl.DrawElements(gl.TRIANGLES, m.Size, gl.UNSIGNED_SHORT, nil)
}
*/
