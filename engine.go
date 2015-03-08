package noname

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

func RenderThing(cam *Camera, mesh Mesh, prog RenderProgram, t *Transform, diffuse Texture) {
	prog.Prog.Use()

	model := t.Mat4()
	prog.M.UniformMatrix4fv(1, false, &model[0])

	normal := cam.View.Mul4(model).Inv()
	prog.N.UniformMatrix4fv(1, true, &normal[0])

	prog.P.UniformMatrix4fv(1, false, &cam.Projection[0])
	prog.V.UniformMatrix4fv(1, false, &cam.View[0])

	prog.Eye.Uniform3fv(1, &cam.Pos[0])

	gl.ActiveTexture(TextureUnitDiffuse)
	diffuse.Bind(gl.TEXTURE_2D)
	prog.Diffuse.Uniform1i(TextureUniformDiffuse)

	mesh.Bind()
	mesh.DrawCall()
}

func ViewPortChange(width, height int32) {
	gl.Viewport(0, 0, width, height)
}
