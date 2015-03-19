package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//Program is a high level representation of OpenGL program object.
type Program uint32

//Delete the program from OpenGL memory. Becomes unusable after this is invoked
func (p Program) Delete() {
	gl.DeleteProgram(uint32(p))
}

//Use is an alias for glUseProgram(p)
func (p Program) Use() {
	gl.UseProgram(uint32(p))
}

//Unuse is an alias for glUseProgram(0)
func (p Program) Unuse() {
	gl.UseProgram(0)
}

//GetUniformLocation is an alias for glGetUniformLocation(p, name)
func (p Program) GetUniformLocation(name string) UniformLocation {
	return UniformLocation(gl.GetUniformLocation(uint32(p), gl.Str(name+"\x00")))
}

//BindFragDataLocation is an alias for glBindFragDataLocation(p, color, name)
func (p Program) BindFragDataLocation(color uint32, name string) {
	gl.BindFragDataLocation(uint32(p), color, gl.Str(name+"\x00"))
}
