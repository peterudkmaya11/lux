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

func (va VertexArray) Bind() {
	gl.BindVertexArray(uint32(va))
}

func (va VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

func (va VertexArray) Delete() {
	gl.DeleteVertexArrays(1, (*uint32)(&va))
}

type Buffer uint32

func GenBuffer() Buffer {
	var va uint32
	gl.GenBuffers(1, &va)
	return Buffer(va)
}

func (b Buffer) Bind(target uint32) {
	gl.BindBuffer(target, uint32(b))
}

func (b Buffer) Unbind(target uint32) {
	gl.BindBuffer(target, 0)
}

func (b Buffer) Delete() {
	gl.DeleteBuffers(1, (*uint32)(&b))
}

//UNIFORMS

type UniformLocation int32

func (ul UniformLocation) UniformMatrix4fv(count int32, transpose bool, values *float32) {
	gl.UniformMatrix4fv(int32(ul), count, transpose, values)
}

func (ul UniformLocation) Uniform1i(v0 int32) {
	gl.Uniform1i(int32(ul), v0)
}

func (ul UniformLocation) Uniform3fv(count int32, values *float32) {
	gl.Uniform3fv(int32(ul), count, values)
}

func (ul UniformLocation) Uniform2f(v0, v1 float32) {
	gl.Uniform2f(int32(ul), v0, v1)
}

func (ul UniformLocation) Uniform1f(v0 float32) {
	gl.Uniform1f(int32(ul), v0)
}
func (ul UniformLocation) Uniform2i(v0, v1 int32) {
	gl.Uniform2i(int32(ul), v0, v1)
}

func (ul UniformLocation) Uniform2fv(v *[2]float32) {
	gl.Uniform2fv(int32(ul), 1, &v[0])
}

//Utility
func GetCurrentTexture2D() Texture2D {
	var i int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &i)
	return Texture2D(i)
}
