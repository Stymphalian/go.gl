package jgl

import (
	"math"
)

type FrustumDim struct {
	Left, Right, Bottom, Top, Near, Far float64
}

func (this *FrustumDim) Set(Left, Right, Bottom, Top, Near, Far float64) {
	this.Left = Left
	this.Right = Right
	this.Bottom = Bottom
	this.Top = Top
	this.Near = Near
	this.Far = Far
}

func (this *FrustumDim) FromPerspective(fov_y, aspect, near, far float64) {
	top := math.Tan(fov_y/2) * near
	right := top * aspect
	this.Set(-right, right, -top, top, near, far)
}

func (this FrustumDim) FOV() float64 {
	return 2 * math.Atan(this.Top/this.Near)
}
func (this FrustumDim) Aspect() float64 {
	return this.Right / this.Top
}
