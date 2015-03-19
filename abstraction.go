package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Deletable interface {
	Delete()
}

//Representation of openGL VertexArray
type VertexArray uint32

//Generate 1 vertex array
func GenVertexArray() VertexArray {
	var va uint32
	gl.GenVertexArrays(1, &va)
	return VertexArray(va)
}

//Binds this vertex array, Alias to gl.BindVertexArray(uint32(va))
func (va VertexArray) Bind() {
	gl.BindVertexArray(uint32(va))
}

//Binds vertex array 0(zero). Alias to gl.BindVertexArray(0)
func (va VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

//Deletes this vertex array
func (va VertexArray) Delete() {
	gl.DeleteVertexArrays(1, (*uint32)(&va))
}

//Representation of openGL Buffer
//TODO: make subtype buffers with restrained functions
//TODO: make an enum for targets
type Buffer uint32

//Generates 1 Buffer
func GenBuffer() Buffer {
	var va uint32
	gl.GenBuffers(1, &va)
	return Buffer(va)
}

//Binds this buffer, requires target. Alias to gl.BindBuffer(target, uint32(b))
func (b Buffer) Bind(target uint32) {
	gl.BindBuffer(target, uint32(b))
}

//Binds buffer 0(zero). Alias to gl.BindBuffer(target, 0)
func (b Buffer) Unbind(target uint32) {
	gl.BindBuffer(target, 0)
}

//Deletes this buffer
func (b Buffer) Delete() {
	gl.DeleteBuffers(1, (*uint32)(&b))
}

//Representation of openGL shader uniform location, generated via the program object.
type UniformLocation int32

//upload 4x4 matrix
func (ul UniformLocation) UniformMatrix4fv(count int32, transpose bool, values *float32) {
	gl.UniformMatrix4fv(int32(ul), count, transpose, values)
}

//upload an integer
func (ul UniformLocation) Uniform1i(v0 int32) {
	gl.Uniform1i(int32(ul), v0)
}

//upload a vec3 of floats via vector
func (ul UniformLocation) Uniform3fv(count int32, values *float32) {
	gl.Uniform3fv(int32(ul), count, values)
}

//upload a vec2 of floats via values
func (ul UniformLocation) Uniform2f(v0, v1 float32) {
	gl.Uniform2f(int32(ul), v0, v1)
}

//upload a float
func (ul UniformLocation) Uniform1f(v0 float32) {
	gl.Uniform1f(int32(ul), v0)
}

//upload a vec2 of ints via values
func (ul UniformLocation) Uniform2i(v0, v1 int32) {
	gl.Uniform2i(int32(ul), v0, v1)
}

//upload a vec2 of floats via vector
func (ul UniformLocation) Uniform2fv(v *float32) {
	gl.Uniform2fv(int32(ul), 1, &v[0])
}

//Generate a single texture2D
func GetCurrentTexture2D() Texture2D {
	var i int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &i)
	return Texture2D(i)
}
