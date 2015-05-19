package jgl

import (
	"github.com/Stymphalian/go.math/lmath"
)

// Vectors are specified in World coordinates
type Ray struct {
	Origin, Dir lmath.Vec3

	u, v, w lmath.Vec3
}

func (this *Ray) CalcRay(
	renderWidth, renderHeight float64,
	col, row float64,
	origin lmath.Vec3,
	focalLength float64) {

	this.Origin = origin
	r := renderWidth / 2
	l := -r
	t := renderHeight / 2
	b := -t
	subu := l + (r-l)*(col+0.5)/renderWidth
	subv := b + (t-b)*(row+0.5)/renderHeight

	wcomp := this.w.MultScalar(-focalLength)
	ucomp := this.u.MultScalar(subu)
	vcomp := this.v.MultScalar(subv)
	this.Dir = wcomp.Add(ucomp).Add(vcomp)
}

func (this *Ray) CalcUVW(gaze, up lmath.Vec3) {
	this.w = gaze.MultScalar(-1).Normalize()
	this.u = up.Cross(this.w).Normalize()
	this.v = this.w.Cross(this.u).Normalize()
}
