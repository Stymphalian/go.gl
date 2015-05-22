package jgl

import (
	"github.com/Stymphalian/go.math/lmath"
)

type Basis struct {
	forward, right, up lmath.Vec3
}

// Create a basis using the given eye, at and up vectors
func MakeBasis(eye, at, up lmath.Vec3) (out Basis) {
	out.forward = at.Sub(eye).Normalize()
	out.right = up.Cross(out.forward).Normalize()
	out.up = out.forward.Cross(out.right).Normalize()
	return
}

// Set the basis' forward, right and up vectors
// Make sure to set these properly
func (this *Basis) Set(forward, right, up lmath.Vec3) *Basis {
	this.forward = forward
	this.right = right
	this.up = up
	return this
}

// Accessor methods for the forward, right and up vectors of the basis
func (this Basis) Forward() lmath.Vec3 {
	return this.forward
}
func (this Basis) Right() lmath.Vec3 {
	return this.right
}
func (this Basis) Up() lmath.Vec3 {
	return this.up
}

//  Rotate the basis about the axis with the given angle
// This applies the rotation to each of the forward, right, and up vectors
func (this *Basis) Rotate(angle float64, x, y, z float64) *Basis {
	q := lmath.Quat{}
	q.FromAxisAngle(angle, x, y, z)
	this.forward = q.RotateVec3(this.forward).Normalize()
	this.right = q.RotateVec3(this.right).Normalize()
	this.up = q.RotateVec3(this.up).Normalize()
	return this
}
