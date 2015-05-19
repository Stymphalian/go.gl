package jgl

import (
	"github.com/Stymphalian/go.math/lmath"
)

type HitRecord struct {
	Hit     bool
	Dist    float64
	MinDist float64
	MaxDist float64

	// Set only if Hit is true
	HitObject Primitive
	Transform lmath.Mat4
}
