package lux

import (
	glm "github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	LocalToWorld glm.Mat4
}

func NewTransform() *Transform {
	out := Transform{}
	out.Iden()
	return &out
}

func (t *Transform) Translate(x, y, z float32) {
	t.LocalToWorld = t.LocalToWorld.Mul4(glm.Translate3D(x, y, z))
}

func (t *Transform) SetTranslate(x, y, z float32) {
	t.LocalToWorld = glm.Translate3D(x, y, z)
}

func (t *Transform) QuatRotate(angle float32, axis glm.Vec3) {
	t.LocalToWorld = t.LocalToWorld.Mul4(glm.HomogRotate3D(angle, axis))
}

func (t *Transform) SetQuatRotate(angle float32, axis glm.Vec3) {
	t.LocalToWorld = glm.HomogRotate3D(angle, axis)
}

//only uniform scaling to prevent 'not having an inverse' and eventually accellerate matrix inversion
func (t *Transform) Scale(amount float32) {
	t.LocalToWorld = t.LocalToWorld.Mul4(glm.Scale3D(amount, amount, amount))
}

func (t *Transform) SetScale(amount float32) {
	t.LocalToWorld = glm.Scale3D(amount, amount, amount)
}

func (t *Transform) Iden() {
	t.LocalToWorld = glm.Ident4()
}

func (t *Transform) Mat4() glm.Mat4 {
	return t.LocalToWorld
}
