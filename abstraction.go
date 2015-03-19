package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//Deletable is the interface to represent everything that can be unallocated.
type Deletable interface {
	Delete()
}

//VertexArray is the high level representation of openGL VertexArray.
type VertexArray uint32

//GenVertexArray generates 1 vertex array. Alias for glGenVertexArrays(1, &va).
func GenVertexArray() VertexArray {
	var va uint32
	gl.GenVertexArrays(1, &va)
	return VertexArray(va)
}

//Bind this vertex array, Alias to gl.BindVertexArray(uint32(va)).
func (va VertexArray) Bind() {
	gl.BindVertexArray(uint32(va))
}

//Unbind bind vertex array 0(zero). Alias to gl.BindVertexArray(0).
func (va VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

//Delete this vertex array.
func (va VertexArray) Delete() {
	gl.DeleteVertexArrays(1, (*uint32)(&va))
}

//Buffer is the high level representation of OpenGL Buffer.
//TODO: make subtype buffers with restrained functions
//TODO: make an enum for targets
type Buffer uint32

//GenBuffer generates 1 Buffer.
func GenBuffer() Buffer {
	var va uint32
	gl.GenBuffers(1, &va)
	return Buffer(va)
}

//Bind this buffer, requires target. Alias to gl.BindBuffer(target, uint32(b)).
func (b Buffer) Bind(target uint32) {
	gl.BindBuffer(target, uint32(b))
}

//Unbind binds buffer 0(zero). Alias to gl.BindBuffer(target, 0).
func (b Buffer) Unbind(target uint32) {
	gl.BindBuffer(target, 0)
}

//Delete this buffer
func (b Buffer) Delete() {
	gl.DeleteBuffers(1, (*uint32)(&b))
}

//UniformLocation is the high level representation of openGL shader uniform location, generated via the program object.
type UniformLocation int32

//UniformMatrix4fv upload a 4x4 matrix.
func (ul UniformLocation) UniformMatrix4fv(count int32, transpose bool, values *float32) {
	gl.UniformMatrix4fv(int32(ul), count, transpose, values)
}

//Uniform1i upload an int32.
func (ul UniformLocation) Uniform1i(v0 int32) {
	gl.Uniform1i(int32(ul), v0)
}

//Uniform3fv upload a vec3 of floats via vector.
func (ul UniformLocation) Uniform3fv(count int32, values *float32) {
	gl.Uniform3fv(int32(ul), count, values)
}

//Uniform2f upload a vec2 of floats via values.
func (ul UniformLocation) Uniform2f(v0, v1 float32) {
	gl.Uniform2f(int32(ul), v0, v1)
}

//Uniform1f upload a float.
func (ul UniformLocation) Uniform1f(v0 float32) {
	gl.Uniform1f(int32(ul), v0)
}

//Uniform2i upload a vec2 of ints via values.
func (ul UniformLocation) Uniform2i(v0, v1 int32) {
	gl.Uniform2i(int32(ul), v0, v1)
}

//Uniform2fv upload a vec2 of floats via vector.
func (ul UniformLocation) Uniform2fv(v *[2]float32) {
	gl.Uniform2fv(int32(ul), 1, &v[0])
}

//GetCurrentTexture2D generate a single texture2D.
func GetCurrentTexture2D() Texture2D {
	var i int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &i)
	return Texture2D(i)
}
