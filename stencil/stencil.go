package noname

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

func Enable() {
	gl.Enable(gl.STENCIL_TEST)
}

func Disable() {
	gl.Disable(gl.STENCIL_TEST)
}

func Func(f StencilFunc, ref int32, mask uint32) {
	gl.StencilFunc(uint32(f), ref, mask)
}

func Op(sfail, zfail, zpass StencilOp) {
	gl.StencilOp(uint32(sfail), uint32(zfail), uint32(zpass))
}

func Mask(mask uint32) {
	gl.StencilMask(mask)
}

type StencilFunc uint32

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

type StencilOp uint32

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

/*
	gl.StencilOp(gl.KEEP, gl.KEEP, gl.REPLACE)
	gl.StencilMask(0xFF)

	gl.ColorMask(false, false, false, false)
	gl.DepthMask(false)

	gl.StencilFunc(gl.EQUAL, 0x1, 0xFF)
	gl.StencilOp(gl.KEEP, gl.KEEP, gl.KEEP)
	gl.ColorMask(true, true, true, true)
	gl.DepthMask(true)


*/
