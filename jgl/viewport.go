package jgl

import (
	"github.com/go-gl/gl/v3.3-compatibility/gl"
)

type Viewport [4]int32

func (this *Viewport) LoadFromGL() {
	gl.GetIntegerv(gl.VIEWPORT, &this[0])
}
func (this Viewport) X() float64 {
	return float64(this[0])
}
func (this Viewport) Y() float64 {
	return float64(this[1])
}
func (this Viewport) W() float64 {
	return float64(this[2])
}
func (this Viewport) H() float64 {
	return float64(this[3])
}
