package jgl

import (
	"github.com/Stymphalian/go.math/lmath"
)

type Primitive interface {
	Draw(transform lmath.Mat4)
	Intersects(r Ray, hit HitRecord, transform lmath.Mat4) HitRecord
	Normal(hitPoint lmath.Vec3, hit HitRecord) lmath.Vec3
	Material() Material
}
