package lux

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"strings"
)

type Shader struct {
	Stype ShaderType
	Loc   uint32
}

func (s Shader) Delete() {
	gl.DeleteShader(s.Loc)
}

func CompileShader(source string, shaderType ShaderType) (Shader, error) {
	shader := gl.CreateShader(uint32(shaderType))

	csource := gl.Str(source)
	gl.ShaderSource(shader, 1, &csource, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return Shader{}, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return Shader{shaderType, shader}, nil
}
