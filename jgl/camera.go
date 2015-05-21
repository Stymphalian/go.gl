package jgl

import (
	// "fmt"
	"github.com/Stymphalian/go.math/lmath"
)

type Camera struct {
	Eye lmath.Vec3
	At  lmath.Vec3
	Up  lmath.Vec3

	//Forward,Right,Up lmath.Vec3
	// projectionMat lmath.Mat4
	// modelViewMat lmath.Mat4
	// IsOrtho
}

// move(eye)
// look(at)
// projection
// modelview(eye, at)

// GetModelView
// GetProjection

func (this *Camera) RotateEye(rot lmath.Quat) {
	// axis := rot.Axis()
	// theta := rot.Angle()

	this.Eye = rot.RotateVec3(this.Eye)
	this.Up = rot.RotateVec3(this.Up)
}

func (this *Camera) MoveEye(eye lmath.Vec3) {
	old_v := this.Eye.Sub(this.At).Normalize()
	new_v := eye.Sub(this.At).Normalize()

	axis := old_v.Cross(new_v)
	theta := old_v.Dot(new_v)
	rot := lmath.Quat{}
	rot.FromAxisAngle(theta, axis.X, axis.Y, axis.Z)

	this.Up = rot.RotateVec3(this.Up)
	this.Eye = eye
}

func (this *Camera) MoveAt(at lmath.Vec3) {
	old_v := this.At.Sub(this.Eye).Normalize()
	new_v := at.Sub(this.Eye).Normalize()

	axis := new_v.Cross(old_v)
	theta := old_v.Dot(new_v)
	rot := lmath.Quat{}
	rot.FromAxisAngle(theta, axis.X, axis.Y, axis.Z)

	this.Up = rot.RotateVec3(this.Up)
	this.At = at
}

func (this Camera) LookAt() lmath.Mat4 {
	return LookAt(this.Eye, this.At, this.Up)
}
