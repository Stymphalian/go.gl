package jgl

import (
	// "fmt"
	// "math"
	"github.com/Stymphalian/go.math/lmath"
	// "github.com/go-gl/gl/v3.3-compatibility/gl"
)

type Box struct {
	TriMesh
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

	return out
}
