package jgl

import (
	"github.com/Stymphalian/go.math/lmath"
)

type Camera struct {
	Eye lmath.Vec3
	At  lmath.Vec3

	FocalLength float64
	//Forward,Right,Up lmath.Vec3
}

func (this Camera) LookAt() lmath.Mat4 {
	up := lmath.Vec3{0, 1, 0}
	if this.At.Eq(lmath.Vec3{0, 1, 0}) {
		up = lmath.Vec3{0, 0, -1}
	} else if this.At.Eq(lmath.Vec3{0, -1, 0}) {
		up = lmath.Vec3{0, 0, 1}
	}
	return LookAt(this.Eye, this.At, up)
}
