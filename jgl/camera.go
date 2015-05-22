package jgl

import (
	// "fmt"
	"github.com/Stymphalian/go.math/lmath"
)

type Camera struct {
	eye lmath.Vec3
	at  lmath.Vec3
	up  lmath.Vec3
}

func MakeCamera(eye, at, up lmath.Vec3) (c Camera) {
	c.Set(eye, at, up)
	return
}

func (this *Camera) Set(eye, at, up lmath.Vec3) *Camera {
	this.eye = eye
	this.at = at
	this.up = up
	return this
}

// rotate the eye about the given quaternion
func (this *Camera) RotateEye(rot lmath.Quat) {
	this.eye = rot.RotateVec3(this.eye)
	this.up = rot.RotateVec3(this.up)
}

func (this *Camera) MoveEye(eye lmath.Vec3) {
	old_v := this.eye.Sub(this.at).Normalize()
	new_v := eye.Sub(this.at).Normalize()

	axis := old_v.Cross(new_v)
	theta := old_v.Dot(new_v)
	rot := lmath.Quat{}
	rot.FromAxisAngle(theta, axis.X, axis.Y, axis.Z)

	this.up = rot.RotateVec3(this.up)
	this.eye = eye
}

func (this *Camera) MoveAt(at lmath.Vec3) {
	old_v := this.at.Sub(this.eye).Normalize()
	new_v := at.Sub(this.eye).Normalize()

	axis := new_v.Cross(old_v)
	theta := old_v.Dot(new_v)
	rot := lmath.Quat{}
	rot.FromAxisAngle(theta, axis.X, axis.Y, axis.Z)

	this.up = rot.RotateVec3(this.up)
	this.at = at
}

func (this *Camera) Zoom(factor float64) {
	this.eye.AddIn(this.Forward().MultScalar(factor))
}

func (this *Camera) Forward() lmath.Vec3 {
	return this.at.Sub(this.eye).Normalize()
}

// TODO: Make sure this returns the correct vec
func (this *Camera) Right() lmath.Vec3 {
	return this.up.Cross(this.Forward()).Normalize()
}

func (this Camera) Eye() lmath.Vec3 {
	return this.eye
}
func (this Camera) At() lmath.Vec3 {
	return this.at
}
func (this *Camera) Up() lmath.Vec3 {
	return this.up
}

func (this Camera) LookAt() lmath.Mat4 {
	return LookAt(this.eye, this.at, this.up)
}
