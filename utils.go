package lux

import (
	"fmt" //for error
	"github.com/go-gl/gl/v3.3-core/gl"
	"log"
	gl2 "luxengine.net/gl"
	"strings" //for string manip
)

//NewProgram will create an OpenGL program from the given shaders, any combinations can be used
func NewProgram(shaders ...Shader) (gl2.Program, error) {
	program := gl2.CreateProgram()

	for _, shader := range shaders {
		program.AttachShader(shader.Loc)
		//gl.AttachShader(program, shader.Loc)
	}
	program.Link()

	if !program.GetLinkStatus() {
		var logLength int32
		gl.GetProgramiv(uint32(program), gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(uint32(program), logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	return gl2.Program(program), nil
}

//MustNotGLError will check opengl for error and panic if one was generated
func MustNotGLError() {
	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Panic(GLErrorToString(err))
	}
}

//GLErrorToString takes the OpenGL error code and return thr appropriate string describing this error
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

//GenRGBTexture2D is a utility function to generate an empty 2D textures of size (width,height), internal format RGB adn data type UNSIGNED_BYTE
func GenRGBTexture2D(width, height int32) gl2.Texture2D {
	tex := gl2.GenTexture2D()
	tex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, width, height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	return tex
}
