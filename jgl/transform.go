package jgl

import (
	"github.com/Stymphalian/go.math/lmath"
)

type Transform struct {
	Trans_x, Trans_y, Trans_z    float64
	Rotate_x, Rotate_y, Rotate_z float64
	Scale_x, Scale_y, Scale_z    float64

	// cached values of the transform parameters
	trans_x, trans_y, trans_z    float64
	rotate_x, rotate_y, rotate_z float64
	scale_x, scale_y, scale_z    float64
	// cached-copy of the tranform matrix
	cachedMat lmath.Mat4
}

func (this *Transform) Translate(x, y, z float64) *Transform {
	this.Trans_x, this.Trans_y, this.Trans_z = x, y, z
	return this
}
func (this *Transform) Scale(x, y, z float64) *Transform {
	this.Scale_x, this.Scale_y, this.Scale_z = x, y, z
	return this
}
func (this *Transform) Rotate(x, y, z float64) *Transform {
	this.Rotate_x, this.Rotate_y, this.Rotate_z = x, y, z
	return this
}

// Dump the transform. Applies the transformations in the order
// scale => rotate => translate
func (this Transform) Dump() lmath.Mat4 {
	// Check to see if we can use the cached version of the matrix.
	if !closeEqf64(this.Trans_x, this.trans_x, epsilon) ||
		!closeEqf64(this.Trans_y, this.trans_y, epsilon) ||
		!closeEqf64(this.Trans_z, this.trans_z, epsilon) ||
		!closeEqf64(this.Rotate_x, this.rotate_x, epsilon) ||
		!closeEqf64(this.Rotate_y, this.rotate_y, epsilon) ||
		!closeEqf64(this.Rotate_z, this.rotate_z, epsilon) ||
		!closeEqf64(this.Scale_x, this.scale_x, epsilon) ||
		!closeEqf64(this.Scale_y, this.scale_y, epsilon) ||
		!closeEqf64(this.Scale_z, this.scale_z, epsilon) {
		this.trans_x, this.trans_y, this.trans_z = this.Trans_x, this.Trans_y, this.Trans_z
		this.rotate_x, this.rotate_y, this.rotate_z = this.Rotate_x, this.Rotate_y, this.Rotate_z
		this.scale_x, this.scale_y, this.scale_z = this.Scale_x, this.Scale_y, this.Scale_z

		translate := lmath.Mat4{}
		rotate := lmath.Mat4{}
		scale := lmath.Mat4{}

		translate.ToTranslate(this.trans_x, this.trans_y, this.trans_z)
		rotate.FromEuler(this.rotate_x, this.rotate_y, this.rotate_z)
		scale.ToScale(this.scale_x, this.scale_y, this.scale_z)
		this.cachedMat = translate.Mult(rotate).Mult(scale)
	}

	return this.cachedMat
}
