package lux

import (
	"fmt" //for error
	"github.com/go-gl/gl/v3.3-core/gl"
	"log"
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

func MustNotGLError() {
	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Fatal(GLErrorToString(err))
	}
}

func GLErrorToString(err uint32) string {
	switch err {
	case gl.NO_ERROR:
		return "GL_NO_ERROR"
	case gl.INVALID_ENUM:
		return "GL_INVALID_ENUM"
	case gl.INVALID_VALUE:
		return "GL_INVALID_VALUE"
	case gl.INVALID_OPERATION:
		return "GL_INVALID_OPERATION"
	case gl.INVALID_FRAMEBUFFER_OPERATION:
		return "GL_INVALID_FRAMEBUFFER_OPERATION"
	case gl.OUT_OF_MEMORY:
		return "GL_OUT_OF_MEMORY"
	case gl.STACK_UNDERFLOW:
		return "GL_STACK_UNDERFLOW"
	case gl.STACK_OVERFLOW:
		return "GL_STACK_OVERFLOW"
	default:
		return "Unknown Error Code"
	}
}
