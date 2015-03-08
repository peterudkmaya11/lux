package mask

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

var Depth DepthMask

type DepthMask struct{}

func (this DepthMask) Total() {
	gl.DepthMask(false)
}

func (this DepthMask) None() {
	gl.DepthMask(true)
}

func (this DepthMask) Mask(mask bool) {
	gl.DepthMask(mask)
}

var Color ColorMask

type ColorMask struct{}

func (this ColorMask) Total() {
	gl.ColorMask(false, false, false, false)
}

func (this ColorMask) None() {
	gl.ColorMask(true, true, true, true)
}

func (this ColorMask) Mask(r, g, b, a bool) {
	gl.ColorMask(r, g, b, a)
}
