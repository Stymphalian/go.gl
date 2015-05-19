package jgl

import (
	"github.com/Stymphalian/go.math/lmath"
)

type Mat4Builder struct {
	stack []lmath.Mat4
}

func NewMat4Builder() *Mat4Builder {
	this := &Mat4Builder{}
	this.stack = make([]lmath.Mat4, 1, 32)
	return this
}

func (this *Mat4Builder) Empty() bool {
	return (len(this.stack) == 0)
}
func (this *Mat4Builder) Push() {
	this.stack = append(this.stack, this.stack[len(this.stack)-1])
}
func (this *Mat4Builder) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}
func (this *Mat4Builder) Peek() lmath.Mat4 {
	return this.stack[len(this.stack)-1]
}

func (this *Mat4Builder) LoadIdentity() {
	this.stack[len(this.stack)-1].ToIdentity()
}

// x,y,z Radians
func (this *Mat4Builder) Rotate(x, y, z float64) {
	mat := lmath.Mat4{}
	mat.FromEuler(x, y, z)
	this.stack[len(this.stack)-1].MultIn(mat)
}
func (this *Mat4Builder) Translate(x, y, z float64) {
	mat := lmath.Mat4{}
	mat.ToTranslate(x, y, z)
	this.stack[len(this.stack)-1].MultIn(mat)
}
func (this *Mat4Builder) Scale(x, y, z float64) {
	mat := lmath.Mat4{}
	mat.ToScale(x, y, z)
	this.stack[len(this.stack)-1].MultIn(mat)
}

func (this *Mat4Builder) Load(mat lmath.Mat4) {
	this.stack[len(this.stack)-1] = mat
}
