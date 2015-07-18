package lux

import (
	glm "github.com/go-gl/mathgl/mgl32"
)

//Transform represent a single local-to-world transformation matrix. It might be upgraded to be used in a tree. (parent-child relation)
type Transform struct {
	LocalToWorld glm.Mat4
}

//NewTransform creates a new Transform
func NewTransform() *Transform {
	out := Transform{}
	out.Iden()
	return &out
}

//Translate add the translation (x,y,z) to the current transform.
func (t *Transform) Translate(x, y, z float32) {
	t.LocalToWorld = t.LocalToWorld.Mul4(glm.Translate3D(x, y, z))
}

//SetTranslate reset this transform to represent only the translation transform given by (x,y,z).
func (t *Transform) SetTranslate(x, y, z float32) {
	t.LocalToWorld = glm.Translate3D(x, y, z)
}

//QuatRotate add the rotation represented by this (angle,quat) to the current transform.
func (t *Transform) QuatRotate(angle float32, axis glm.Vec3) {
	t.LocalToWorld = t.LocalToWorld.Mul4(glm.HomogRotate3D(angle, axis))
}

//SetQuatRotate will reset this transform to represent the rotation represented by this (angle,quat).
func (t *Transform) SetQuatRotate(angle float32, axis glm.Vec3) {
	t.LocalToWorld = glm.HomogRotate3D(angle, axis)
}

//Scale add a scaling operation to the currently stored transform.
//I do not allow non-uniform scaling to prevent ending up with matrices without an inverse.
func (t *Transform) Scale(amount float32) {
	t.LocalToWorld = t.LocalToWorld.Mul4(glm.Scale3D(amount, amount, amount))
}

//SetScale reset this transform to represent only the scaling transform of `amount`
//I do not allow non-uniform scaling to prevent ending up with matrices without an inverse.
func (t *Transform) SetScale(amount float32) {
	t.LocalToWorld = glm.Scale3D(amount, amount, amount)
}

//Iden set this transform to the identity 4x4 matrix
func (t *Transform) Iden() {
	t.LocalToWorld = glm.Ident4()
}

//Mat4 returns the mathgl 4x4 matrix that represents this transform local-to-world transformation matrix.
func (t *Transform) Mat4() glm.Mat4 {
	return t.LocalToWorld
}

func (t *Transform) SetMatrix(m *[16]float32) {
	t.LocalToWorld = (glm.Mat4)(*m)
}
