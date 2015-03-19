package stencil

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//Alias to gl.Enable(gl.STENCIL_TEST)
func Enable() {
	gl.Enable(gl.STENCIL_TEST)
}

//Alias to gl.Disable(gl.STENCIL_TEST)
func Disable() {
	gl.Disable(gl.STENCIL_TEST)
}

//Alias to gl.StencilFunc(f, ref, mask)
func Func(f StencilFunc, ref int32, mask uint32) {
	gl.StencilFunc(uint32(f), ref, mask)
}

//Alias to gl.StencilOp(sfail, zfail, zpass)
func Op(sfail, zfail, zpass StencilOp) {
	gl.StencilOp(uint32(sfail), uint32(zfail), uint32(zpass))
}

//Alias to gl.StencilMask(mask)
func Mask(mask uint32) {
	gl.StencilMask(mask)
}

//Enum to represent all possible stencil func, prevents bad arguments
type StencilFunc uint32

//Stencil func possible values
const (
	Never    StencilFunc = gl.NEVER
	Less                 = gl.LESS
	Lequal               = gl.LEQUAL
	Greater              = gl.GREATER
	Gequal               = gl.GEQUAL
	Equal                = gl.EQUAL
	NotEqual             = gl.NOTEQUAL
	Always               = gl.ALWAYS
)

//Enum to represent all possible stencil op, prevent bad arguments
type StencilOp uint32

//Stencil op possible values
const (
	Keep     StencilOp = gl.KEEP
	Zero               = gl.ZERO
	Replace            = gl.REPLACE
	Incr               = gl.INCR
	IncrWrap           = gl.INCR_WRAP
	Decr               = gl.DECR
	DecrWrap           = gl.DECR_WRAP
	Invert             = gl.INVERT
)
