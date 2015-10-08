package lux

/*
import (
	"github.com/go-gl/gl/v3.3-core/gl"
	gl2 "github.com/luxengine/gl"
	"math/rand"
	"time"
	"unsafe"
)

//MakeRandomTexture will create a 1D texture of size "size" with all random only in the red channel, the values are 16f between [0.0,1.0)
func MakeRandomTexture(size int) gl2.Texture {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]float32, size)
	for x := 0; x < len(data); x++ {
		data[x] = r.Float32()
	}

	tex := gl2.GenTexture()
	tex.Bind(gl.TEXTURE_1D)
	tex.TexImage1D(gl.TEXTURE_1D, 0, gl.R16F, int32(size), 0, gl.RED, gl.FLOAT, unsafe.Pointer(&data[0]))
	gl.TexParameterf(gl.TEXTURE_1D, gl.TEXTURE_MIN_FILTER, gl2.LINEAR)
	gl.TexParameterf(gl.TEXTURE_1D, gl.TEXTURE_MAG_FILTER, gl2.LINEAR)
	gl.TexParameterf(gl.TEXTURE_1D, gl.TEXTURE_WRAP_S, gl2.REPEAT)

	return tex
}*/
