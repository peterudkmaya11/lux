package noname

import (
	glm "github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	View       glm.Mat4
	Projection glm.Mat4
	Pos        glm.Vec3
}

func (this *Camera) SetPerspective(angle, ratio, zNear, zFar float32) {
	this.Projection = glm.Perspective(angle, ratio, zNear, zFar)
}

//func (this Camera) SetOrtho(left, right, bottom, top, near, far float32){}

//func to project from 2d to 3d
//func to project from 3d to 2d

//by value
func (this *Camera) LookAtval(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float32) {
	this.View = glm.LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ)
	this.Pos[0], this.Pos[1], this.Pos[2] = eyeX, eyeY, eyeZ
}

//by vector
func (this *Camera) LookAtVec(eye, center, up *glm.Vec3) {
	this.View = glm.LookAt(eye[0], eye[1], eye[2], center[0], center[1], center[2], up[0], up[1], up[2])
	this.Pos = *eye
}
