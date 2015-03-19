package mask

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

//Global variable to encapsulate depth mask related functions.
var Depth DepthMask

//type to encapsulate depth mask related functions.
type DepthMask struct{}

//Alias to gl.DepthMask(false)
func (this DepthMask) Total() {
	gl.DepthMask(false)
}

//Alias to gl.DepthMask(false)
func (this DepthMask) None() {
	gl.DepthMask(true)
}

//Alias to gl.DepthMask(mask)
func (this DepthMask) Mask(mask bool) {
	gl.DepthMask(mask)
}

//Global varialbe to encapsulate color-mask related functions.
var Color ColorMask

//Type to encapsulate color mask related functions.
type ColorMask struct{}

//Alias to gl.ColorMask(false, false, false, false)
func (this ColorMask) Total() {
	gl.ColorMask(false, false, false, false)
}

//Alias to gl.ColorMask(true, true, true, true)
func (this ColorMask) None() {
	gl.ColorMask(true, true, true, true)
}

//Alias to gl.ColorMask(r, g, b, a)
func (this ColorMask) Mask(r, g, b, a bool) {
	gl.ColorMask(r, g, b, a)
}
