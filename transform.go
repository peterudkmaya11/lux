package noname

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

func (this *Transform) Translate(x, y, z float32) {
	this.LocalToWorld = this.LocalToWorld.Mul4(glm.Translate3D(x, y, z))
}

func (this *Transform) SetTranslate(x, y, z float32) {
	this.LocalToWorld = glm.Translate3D(x, y, z)
}

func (this *Transform) QuatRotate(angle float32, axis glm.Vec3) {
	this.LocalToWorld = this.LocalToWorld.Mul4(glm.HomogRotate3D(angle, axis))
}

func (this *Transform) SetQuatRotate(angle float32, axis glm.Vec3) {
	this.LocalToWorld = glm.HomogRotate3D(angle, axis)
}

//only uniform scaling to prevent 'not having an inverse' and eventually accellerate matrix inversion
func (this *Transform) Scale(amount float32) {
	this.LocalToWorld = this.LocalToWorld.Mul4(glm.Scale3D(amount, amount, amount))
}

func (this *Transform) SetScale(amount float32) {
	this.LocalToWorld = glm.Scale3D(amount, amount, amount)
}

func (this *Transform) Iden() {
	this.LocalToWorld = glm.Ident4()
}

func (this *Transform) Mat4() glm.Mat4 {
	return this.LocalToWorld
}
