package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program uint32

//Deletes the program from OpenGL memory. Becomes unusable after this is invoked
func (this Program) Delete() {
	gl.DeleteProgram(uint32(this))
}

//Use this program
func (this Program) Use() {
	gl.UseProgram(uint32(this))
}

//alias for UseProgram(0)
func (this Program) Unuse() {
	gl.UseProgram(0)
}

//return the uniform location of 'name'
func (this Program) GetUniformLocation(name string) UniformLocation {
	return UniformLocation(gl.GetUniformLocation(uint32(this), gl.Str(name+"\x00")))
}

func (this Program) BindFragDataLocation(color uint32, name string) {
	gl.BindFragDataLocation(uint32(this), color, gl.Str(name+"\x00"))
}
