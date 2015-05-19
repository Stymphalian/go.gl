package jgl

import (
	// "fmt"
	"math"

	"github.com/Stymphalian/go.math/lmath"
	"github.com/go-gl/gl/v3.3-compatibility/gl"
)

type Sphere struct {
	Pos    lmath.Vec3
	Radius float64

	LightFlag bool
	Mat       Material

	radius_vec lmath.Vec3
	vertices   []lmath.Vec3
}

func NewSphere(radius float64) *Sphere {
	out := &Sphere{}
	out.Pos = lmath.Vec3{0, 0, 0}
	out.Radius = radius

	out.radius_vec = lmath.Vec3{radius, 0, 0}

	num_lats := 20
	num_longs := 20

	var x, y, z, r float64
	long_divs := float64(math.Pi) / float64(num_longs)
	lat_divs := float64(2*math.Pi) / float64(num_lats)
	out.vertices = make([]lmath.Vec3, 0, num_longs*num_lats)

	for longs := 0; longs < num_longs; longs++ {
		y = math.Cos(long_divs*float64(longs)) * radius
		r = math.Sin(long_divs*float64(longs)) * radius

		for lat := 0; lat < num_lats; lat++ {
			x = math.Cos(lat_divs*float64(lat)) * r
			z = math.Sin(lat_divs*float64(lat)) * r

			v := lmath.Vec3{x, y, z}
			out.vertices = append(out.vertices, v)
		}
	}

	return out
}

func (this *Sphere) Draw(transform lmath.Mat4) {
	gl.Begin(gl.POINTS)
	gl.Color3f(float32(this.Mat.Color[0]),
		float32(this.Mat.Color[1]),
		float32(this.Mat.Color[2]))

	for i := 0; i < len(this.vertices); i++ {
		v := transform.MultVec3(this.vertices[i])
		// v := this.vertices[i].MultMat4(transform)
		gl.Vertex3f(float32(v.X), float32(v.Y), float32(v.Z))
	}
	gl.End()
}

func (this *Sphere) Intersects(ray Ray, hit HitRecord, transform lmath.Mat4) HitRecord {
	// origin - center
	trans_pos := transform.MultVec3(this.Pos)
	trans_rad_vec := transform.MultVec3(this.radius_vec)
	// trans_pos := this.Pos.MultMat4(transform)
	// trans_rad_vec := this.radius_vec.MultMat4(transform)

	omc := ray.Origin.Sub(trans_pos)
	radius := trans_rad_vec.Sub(trans_pos).Length()

	A := ray.Dir.Dot(ray.Dir)
	B := (ray.Dir.MultScalar(2)).Dot(omc)
	C := (omc.Dot(omc)) - radius*radius

	// calculate the descriminant and make sure that it has a solution
	descriminant := B*B - 4*A*C
	if descriminant < epsilon {
		hit.Hit = false
		return hit
	}

	// calculate the two possible solutions
	descriminant = math.Sqrt(descriminant)
	dist1 := (-B + descriminant) / (2 * A)
	dist2 := (-B - descriminant) / (2 * A)
	dist := 0.0
	if dist2 < dist1 {
		dist = dist2
	} else {
		dist = dist1
	}

	// make sure the dist is in range.
	if dist > hit.MinDist && dist < hit.MaxDist {
		hit.Hit = true
		hit.Dist = dist

		hit.MaxDist = hit.Dist
		return hit
	}

	hit.Hit = false
	return hit
}

func (this *Sphere) Direction(hitPoint lmath.Vec3, transform lmath.Mat4) lmath.Vec3 {
	trans_pos := transform.MultVec3(this.Pos)
	return trans_pos.Sub(hitPoint).Normalize()
	// return this.Pos.Sub(hitPoint).Normalize()
}

func (this *Sphere) Normal(hitPoint lmath.Vec3, hit HitRecord) lmath.Vec3 {
	trans_pos := hit.Transform.MultVec3(this.Pos)
	trans_rad_vec := hit.Transform.MultVec3(this.radius_vec)
	radius := trans_rad_vec.Sub(trans_pos).Length()
	return hitPoint.Sub(trans_pos).DivScalar(radius)
}

func (this *Sphere) IsLight() bool {
	return this.LightFlag
}

func (this Sphere) Material() Material {
	return this.Mat
}
