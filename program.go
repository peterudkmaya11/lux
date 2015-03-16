package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program uint32

//Deletes the program from OpenGL memory. Becomes unusable after this is invoked
func (p Program) Delete() {
	gl.DeleteProgram(uint32(p))
}

//Use p program
func (p Program) Use() {
	gl.UseProgram(uint32(p))
}

//alias for UseProgram(0)
func (p Program) Unuse() {
	gl.UseProgram(0)
}

//return the uniform location of 'name'
func (p Program) GetUniformLocation(name string) UniformLocation {
	return UniformLocation(gl.GetUniformLocation(uint32(p), gl.Str(name+"\x00")))
}

func (p Program) BindFragDataLocation(color uint32, name string) {
	gl.BindFragDataLocation(uint32(p), color, gl.Str(name+"\x00"))
}
