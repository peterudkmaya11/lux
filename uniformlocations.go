package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//UniformLocation is the high level representation of openGL shader uniform location, generated via the program object.
type UniformLocation int32

//Uniform1f is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform1f(v0 float32) {
	gl.Uniform1f(int32(ul), v0)
}

//Uniform2f is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform2f(v0, v1 float32) {
	gl.Uniform2f(int32(ul), v0, v1)
}

//Uniform3f is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform3f(v0, v1, v2 float32) {
	gl.Uniform3f(int32(ul), v0, v1, v2)
}

//Uniform4f is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform4f(v0, v1, v2, v3 float32) {
	gl.Uniform4f(int32(ul), v0, v1, v2, v3)
}

//Uniform1i is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform1i(v0 int32) {
	gl.Uniform1i(int32(ul), v0)
}

//Uniform2i is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform2i(v0, v1 int32) {
	gl.Uniform2i(int32(ul), v0, v1)
}

//Uniform3i is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform3i(v0, v1, v2 int32) {
	gl.Uniform3i(int32(ul), v0, v1, v2)
}

//Uniform4i is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform4i(v0, v1, v2, v3 int32) {
	gl.Uniform4i(int32(ul), v0, v1, v2, v3)
}

//Uniform1ui is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform1ui(v0 uint32) {
	gl.Uniform1ui(int32(ul), v0)
}

//Uniform2ui is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform2ui(v0, v1 uint32) {
	gl.Uniform2ui(int32(ul), v0, v1)
}

//Uniform3ui is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform3ui(v0, v1, v2 uint32) {
	gl.Uniform3ui(int32(ul), v0, v1, v2)
}

//Uniform4ui is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform4ui(v0, v1, v2, v3 uint32) {
	gl.Uniform4ui(int32(ul), v0, v1, v2, v3)
}

//Uniform1fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform1fv(count int32, value *float32) {
	gl.Uniform1fv(int32(ul), count, value)
}

//Uniform2fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform2fv(count int32, value *float32) {
	gl.Uniform2fv(int32(ul), count, value)
}

//Uniform3fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform3fv(count int32, value *float32) {
	gl.Uniform3fv(int32(ul), count, value)
}

//Uniform4fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform4fv(count int32, value *float32) {
	gl.Uniform4fv(int32(ul), count, value)
}

//Uniform1iv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform1iv(count int32, value *int32) {
	gl.Uniform1iv(int32(ul), count, value)
}

//Uniform2iv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform2iv(count int32, value *int32) {
	gl.Uniform2iv(int32(ul), count, value)
}

//Uniform3iv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform3iv(count int32, value *int32) {
	gl.Uniform3iv(int32(ul), count, value)
}

//Uniform4iv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform4iv(count int32, value *int32) {
	gl.Uniform4iv(int32(ul), count, value)
}

//Uniform1uiv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform1uiv(count int32, value *uint32) {
	gl.Uniform1uiv(int32(ul), count, value)
}

//Uniform2uiv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform2uiv(count int32, value *uint32) {
	gl.Uniform2uiv(int32(ul), count, value)
}

//Uniform3uiv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform3uiv(count int32, value *uint32) {
	gl.Uniform3uiv(int32(ul), count, value)
}

//Uniform4uiv is an alias for the Opengl function of the same name.
func (ul UniformLocation) Uniform4uiv(count int32, value *uint32) {
	gl.Uniform4uiv(int32(ul), count, value)
}

//UniformMatrix2fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) UniformMatrix2fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2fv(int32(ul), count, transpose, value)
}

//UniformMatrix3fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) UniformMatrix3fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3fv(int32(ul), count, transpose, value)
}

//UniformMatrix4fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) UniformMatrix4fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4fv(int32(ul), count, transpose, value)
}

//UniformMatrix2x3fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) UniformMatrix2x3fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2x3fv(int32(ul), count, transpose, value)
}

//UniformMatrix3x2fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) UniformMatrix3x2fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3x2fv(int32(ul), count, transpose, value)
}

//UniformMatrix2x4fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) UniformMatrix2x4fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2x4fv(int32(ul), count, transpose, value)
}

//UniformMatrix4x2fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) UniformMatrix4x2fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4x2fv(int32(ul), count, transpose, value)
}

//UniformMatrix3x4fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) UniformMatrix3x4fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3x4fv(int32(ul), count, transpose, value)
}

//UniformMatrix4x3fv is an alias for the Opengl function of the same name.
func (ul UniformLocation) UniformMatrix4x3fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4x3fv(int32(ul), count, transpose, value)
}

/*============================
=========Sub uniforms=========
============================*/

//UniformLocation1f is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation1f UniformLocation

//UniformLocation2f is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation2f UniformLocation

//UniformLocation3f is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation3f UniformLocation

//UniformLocation4f is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation4f UniformLocation

//UniformLocation1i is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation1i UniformLocation

//UniformLocation2i is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation2i UniformLocation

//UniformLocation3i is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation3i UniformLocation

//UniformLocation4i is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation4i UniformLocation

//UniformLocation1ui is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation1ui UniformLocation

//UniformLocation2ui is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation2ui UniformLocation

//UniformLocation3ui is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation3ui UniformLocation

//UniformLocation4ui is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocation4ui UniformLocation

//UniformLocationMatrix2 is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocationMatrix2 UniformLocation

//UniformLocationMatrix3 is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocationMatrix3 UniformLocation

//UniformLocationMatrix4 is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocationMatrix4 UniformLocation

//UniformLocationMatrix2x3 is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocationMatrix2x3 UniformLocation

//UniformLocationMatrix3x2 is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocationMatrix3x2 UniformLocation

//UniformLocationMatrix2x4 is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocationMatrix2x4 UniformLocation

//UniformLocationMatrix4x2 is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocationMatrix4x2 UniformLocation

//UniformLocationMatrix3x4 is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocationMatrix3x4 UniformLocation

//UniformLocationMatrix4x3 is a restrained UniformLocation to represent ONLY the function it should use.
type UniformLocationMatrix4x3 UniformLocation

//To1f transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To1f() UniformLocation1f {
	return UniformLocation1f(ul)
}

//To2f transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To2f() UniformLocation2f {
	return UniformLocation2f(ul)
}

//To3f transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To3f() UniformLocation3f {
	return UniformLocation3f(ul)
}

//To4f transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To4f() UniformLocation4f {
	return UniformLocation4f(ul)
}

//To1i transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To1i() UniformLocation1i {
	return UniformLocation1i(ul)
}

//To2i transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To2i() UniformLocation2i {
	return UniformLocation2i(ul)
}

//To3i transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To3i() UniformLocation3i {
	return UniformLocation3i(ul)
}

//To4i transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To4i() UniformLocation4i {
	return UniformLocation4i(ul)
}

//To1ui transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To1ui() UniformLocation1ui {
	return UniformLocation1ui(ul)
}

//To2ui transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To2ui() UniformLocation2ui {
	return UniformLocation2ui(ul)
}

//To3ui transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To3ui() UniformLocation3ui {
	return UniformLocation3ui(ul)
}

//To4ui transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) To4ui() UniformLocation4ui {
	return UniformLocation4ui(ul)
}

//ToMatrix2 transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) ToMatrix2() UniformLocationMatrix2 {
	return UniformLocationMatrix2(ul)
}

//ToMatrix3 transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) ToMatrix3() UniformLocationMatrix3 {
	return UniformLocationMatrix3(ul)
}

//ToMatrix4 transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) ToMatrix4() UniformLocationMatrix4 {
	return UniformLocationMatrix4(ul)
}

//ToMatrix2x3 transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) ToMatrix2x3() UniformLocationMatrix2x3 {
	return UniformLocationMatrix2x3(ul)
}

//ToMatrix3x2 transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) ToMatrix3x2() UniformLocationMatrix3x2 {
	return UniformLocationMatrix3x2(ul)
}

//ToMatrix2x4 transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) ToMatrix2x4() UniformLocationMatrix2x4 {
	return UniformLocationMatrix2x4(ul)
}

//ToMatrix4x2 transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) ToMatrix4x2() UniformLocationMatrix4x2 {
	return UniformLocationMatrix4x2(ul)
}

//ToMatrix3x4 transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) ToMatrix3x4() UniformLocationMatrix3x4 {
	return UniformLocationMatrix3x4(ul)
}

//ToMatrix4x3 transforms this uniform into a restrained version that only has the functions it should use.
func (ul UniformLocation) ToMatrix4x3() UniformLocationMatrix4x3 {
	return UniformLocationMatrix4x3(ul)
}

//Uniform1f is an alias for the Opengl function of the same name.
func (ul UniformLocation1f) Uniform1f(v0 float32) {
	gl.Uniform1f(int32(ul), v0)
}

//Uniform2f is an alias for the Opengl function of the same name.
func (ul UniformLocation2f) Uniform2f(v0, v1 float32) {
	gl.Uniform2f(int32(ul), v0, v1)
}

//Uniform3f is an alias for the Opengl function of the same name.
func (ul UniformLocation3f) Uniform3f(v0, v1, v2 float32) {
	gl.Uniform3f(int32(ul), v0, v1, v2)
}

//Uniform4f is an alias for the Opengl function of the same name.
func (ul UniformLocation4f) Uniform4f(v0, v1, v2, v3 float32) {
	gl.Uniform4f(int32(ul), v0, v1, v2, v3)
}

//Uniform1i is an alias for the Opengl function of the same name.
func (ul UniformLocation1i) Uniform1i(v0 int32) {
	gl.Uniform1i(int32(ul), v0)
}

//Uniform2i is an alias for the Opengl function of the same name.
func (ul UniformLocation2i) Uniform2i(v0, v1 int32) {
	gl.Uniform2i(int32(ul), v0, v1)
}

//Uniform3i is an alias for the Opengl function of the same name.
func (ul UniformLocation3i) Uniform3i(v0, v1, v2 int32) {
	gl.Uniform3i(int32(ul), v0, v1, v2)
}

//Uniform4i is an alias for the Opengl function of the same name.
func (ul UniformLocation4i) Uniform4i(v0, v1, v2, v3 int32) {
	gl.Uniform4i(int32(ul), v0, v1, v2, v3)
}

//Uniform1ui is an alias for the Opengl function of the same name.
func (ul UniformLocation1ui) Uniform1ui(v0 uint32) {
	gl.Uniform1ui(int32(ul), v0)
}

//Uniform2ui is an alias for the Opengl function of the same name.
func (ul UniformLocation2ui) Uniform2ui(v0, v1 uint32) {
	gl.Uniform2ui(int32(ul), v0, v1)
}

//Uniform3ui is an alias for the Opengl function of the same name.
func (ul UniformLocation3ui) Uniform3ui(v0, v1, v2 uint32) {
	gl.Uniform3ui(int32(ul), v0, v1, v2)
}

//Uniform4ui is an alias for the Opengl function of the same name.
func (ul UniformLocation4ui) Uniform4ui(v0, v1, v2, v3 uint32) {
	gl.Uniform4ui(int32(ul), v0, v1, v2, v3)
}

//Uniform1fv is an alias for the Opengl function of the same name.
func (ul UniformLocation1f) Uniform1fv(count int32, value *float32) {
	gl.Uniform1fv(int32(ul), count, value)
}

//Uniform2fv is an alias for the Opengl function of the same name.
func (ul UniformLocation2f) Uniform2fv(count int32, value *float32) {
	gl.Uniform2fv(int32(ul), count, value)
}

//Uniform3fv is an alias for the Opengl function of the same name.
func (ul UniformLocation3f) Uniform3fv(count int32, value *float32) {
	gl.Uniform3fv(int32(ul), count, value)
}

//Uniform4fv is an alias for the Opengl function of the same name.
func (ul UniformLocation4f) Uniform4fv(count int32, value *float32) {
	gl.Uniform4fv(int32(ul), count, value)
}

//Uniform1iv is an alias for the Opengl function of the same name.
func (ul UniformLocation1i) Uniform1iv(count int32, value *int32) {
	gl.Uniform1iv(int32(ul), count, value)
}

//Uniform2iv is an alias for the Opengl function of the same name.
func (ul UniformLocation2i) Uniform2iv(count int32, value *int32) {
	gl.Uniform2iv(int32(ul), count, value)
}

//Uniform3iv is an alias for the Opengl function of the same name.
func (ul UniformLocation3i) Uniform3iv(count int32, value *int32) {
	gl.Uniform3iv(int32(ul), count, value)
}

//Uniform4iv is an alias for the Opengl function of the same name.
func (ul UniformLocation4i) Uniform4iv(count int32, value *int32) {
	gl.Uniform4iv(int32(ul), count, value)
}

//Uniform1uiv is an alias for the Opengl function of the same name.
func (ul UniformLocation1ui) Uniform1uiv(count int32, value *uint32) {
	gl.Uniform1uiv(int32(ul), count, value)
}

//Uniform2uiv is an alias for the Opengl function of the same name.
func (ul UniformLocation2ui) Uniform2uiv(count int32, value *uint32) {
	gl.Uniform2uiv(int32(ul), count, value)
}

//Uniform3uiv is an alias for the Opengl function of the same name.
func (ul UniformLocation3ui) Uniform3uiv(count int32, value *uint32) {
	gl.Uniform3uiv(int32(ul), count, value)
}

//Uniform4uiv is an alias for the Opengl function of the same name.
func (ul UniformLocation4ui) Uniform4uiv(count int32, value *uint32) {
	gl.Uniform4uiv(int32(ul), count, value)
}

//UniformMatrix2fv is an alias for the Opengl function of the same name.
func (ul UniformLocationMatrix2) UniformMatrix2fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2fv(int32(ul), count, transpose, value)
}

//UniformMatrix3fv is an alias for the Opengl function of the same name.
func (ul UniformLocationMatrix3) UniformMatrix3fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3fv(int32(ul), count, transpose, value)
}

//UniformMatrix4fv is an alias for the Opengl function of the same name.
func (ul UniformLocationMatrix4) UniformMatrix4fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4fv(int32(ul), count, transpose, value)
}

//UniformMatrix2x3fv is an alias for the Opengl function of the same name.
func (ul UniformLocationMatrix2x3) UniformMatrix2x3fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2x3fv(int32(ul), count, transpose, value)
}

//UniformMatrix3x2fv is an alias for the Opengl function of the same name.
func (ul UniformLocationMatrix3x2) UniformMatrix3x2fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3x2fv(int32(ul), count, transpose, value)
}

//UniformMatrix2x4fv is an alias for the Opengl function of the same name.
func (ul UniformLocationMatrix2x4) UniformMatrix2x4fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix2x4fv(int32(ul), count, transpose, value)
}

//UniformMatrix4x2fv is an alias for the Opengl function of the same name.
func (ul UniformLocationMatrix4x2) UniformMatrix4x2fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4x2fv(int32(ul), count, transpose, value)
}

//UniformMatrix3x4fv is an alias for the Opengl function of the same name.
func (ul UniformLocationMatrix3x4) UniformMatrix3x4fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix3x4fv(int32(ul), count, transpose, value)
}

//UniformMatrix4x3fv is an alias for the Opengl function of the same name.
func (ul UniformLocationMatrix4x3) UniformMatrix4x3fv(count int32, transpose bool, value *float32) {
	gl.UniformMatrix4x3fv(int32(ul), count, transpose, value)
}
