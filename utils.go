package lux

import (
	"fmt" //for error
	"github.com/go-gl/gl/v3.3-core/gl"
	"strings" //for string manip
)

//NewProgram will create an OpenGL program from the given shaders, any combinations can be used
func NewProgram(shaders ...Shader) (Program, error) {
	program := gl.CreateProgram()

	for _, shader := range shaders {
		gl.AttachShader(program, shader.Loc)
	}
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	return Program(program), nil
}
