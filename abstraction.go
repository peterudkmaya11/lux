package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Deletable interface {
	Delete()
}

type VertexArray uint32

func GenVertexArray() VertexArray {
	var va uint32
	gl.GenVertexArrays(1, &va)
	return VertexArray(va)
}

func (this VertexArray) Bind() {
	gl.BindVertexArray(uint32(this))
}

func (this VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

func (this VertexArray) Delete() {
	gl.DeleteVertexArrays(1, (*uint32)(&this))
}

type Buffer uint32

func GenBuffer() Buffer {
	var va uint32
	gl.GenBuffers(1, &va)
	return Buffer(va)
}

func (this Buffer) Bind(target uint32) {
	gl.BindBuffer(target, uint32(this))
}

func (this Buffer) Unbind(target uint32) {
	gl.BindBuffer(target, 0)
}

func (this Buffer) Delete() {
	gl.DeleteBuffers(1, (*uint32)(&this))
}

//UNIFORMS

type UniformLocation int32

func (this UniformLocation) UniformMatrix4fv(count int32, transpose bool, values *float32) {
	gl.UniformMatrix4fv(int32(this), count, transpose, values)
}

func (this UniformLocation) Uniform1i(v0 int32) {
	gl.Uniform1i(int32(this), v0)
}

func (this UniformLocation) Uniform3fv(count int32, values *float32) {
	gl.Uniform3fv(int32(this), count, values)
}

func (this UniformLocation) Uniform2f(v0, v1 float32) {
	gl.Uniform2f(int32(this), v0, v1)
}

func (this UniformLocation) Uniform1f(v0 float32) {
	gl.Uniform1f(int32(this), v0)
}
func (this UniformLocation) Uniform2i(v0, v1 int32) {
	gl.Uniform2i(int32(this), v0, v1)
}

func (this UniformLocation) Uniform2fv(v *[2]float32) {
	gl.Uniform2fv(int32(this), 1, &v[0])
}
