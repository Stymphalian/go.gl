package jgl

import (
	"fmt"
	// "math"

	"github.com/Stymphalian/go.math/lmath"
	"github.com/go-gl/gl/v3.3-compatibility/gl"
)

type Box struct {
	// LightFlag bool
	Mat   Material
	verts []lmath.Vec3
	tris  [][3]int
}

func NewBox(scale float64) *Box {
	out := &Box{}
	out.verts = make([]lmath.Vec3, 0, 8)
	out.tris = make([][3]int, 0, 12)

	out.verts = append(out.verts, lmath.Vec3{-scale, -scale, scale})
	out.verts = append(out.verts, lmath.Vec3{scale, -scale, scale})
	out.verts = append(out.verts, lmath.Vec3{scale, scale, scale})
	out.verts = append(out.verts, lmath.Vec3{-scale, scale, scale})
	out.verts = append(out.verts, lmath.Vec3{-scale, -scale, -scale})
	out.verts = append(out.verts, lmath.Vec3{scale, -scale, -scale})
	out.verts = append(out.verts, lmath.Vec3{scale, scale, -scale})
	out.verts = append(out.verts, lmath.Vec3{-scale, scale, -scale})

	out.tris = append(out.tris, [3]int{0, 1, 2})
	out.tris = append(out.tris, [3]int{0, 2, 3})
	out.tris = append(out.tris, [3]int{1, 5, 6})
	out.tris = append(out.tris, [3]int{1, 6, 2})
	out.tris = append(out.tris, [3]int{4, 6, 5})
	out.tris = append(out.tris, [3]int{4, 7, 6})
	out.tris = append(out.tris, [3]int{0, 4, 7})
	out.tris = append(out.tris, [3]int{0, 7, 3})
	out.tris = append(out.tris, [3]int{3, 2, 6})
	out.tris = append(out.tris, [3]int{3, 6, 7})
	out.tris = append(out.tris, [3]int{0, 1, 5})
	out.tris = append(out.tris, [3]int{0, 5, 4})

	fmt.Println("alskjlfsadf")
	for _, v := range out.verts {
		fmt.Println(v)
	}

	return out
}

func (this *Box) Draw(transform lmath.Mat4) {
	gl.Color3f(1.0, 1.0, 0)
	gl.Begin(gl.TRIANGLES)
	for i := 0; i < len(this.tris); i++ {
		v := transform.MultVec3(this.verts[this.tris[i][0]])
		gl.Vertex3f(float32(v.X), float32(v.Y), float32(v.Z))
		v = transform.MultVec3(this.verts[this.tris[i][1]])
		gl.Vertex3f(float32(v.X), float32(v.Y), float32(v.Z))
		v = transform.MultVec3(this.verts[this.tris[i][2]])
		gl.Vertex3f(float32(v.X), float32(v.Y), float32(v.Z))
	}
	gl.End()
}

func (this *Box) Intersects(ray Ray, hit HitRecord, transform lmath.Mat4) (h HitRecord) {
	h.Hit = false
	return h
}

func (this *Box) Normal(hitPoint lmath.Vec3, hit HitRecord) lmath.Vec3 {
	return lmath.Vec3{0, 0, 0}
}

func (this Box) Material() Material {
	return this.Mat
}
