package jgl

import (
    "github.com/Stymphalian/go.math/lmath"
)

type Mesh struct{
    Verts []lmath.Vec4
    Normals []lmath.Vec4
    Texs []lmath.Vec4

    TriVerts [][3]int
    TriNormals [][3]int
    TriTexs [][3]int
}

func MakeMesh() (m Mesh){
    m.Verts = make([]lmath.Vec4,0,32)
    m.Normals = make([]lmath.Vec4,0,32)
    m.Texs = make([]lmath.Vec4,0,32)

    m.TriVerts = make([][3]int,0,32)
    m.TriNormals = make([][3]int,0,32)
    m.TriTexs= make([][3]int,0,32)
    return
}

func ( this* Mesh) AddVert(vert lmath.Vec4) int {
    this.Verts = append(this.Verts,vert)
    return len(this.Verts) -1
}
func ( this* Mesh) AddNormal(norm lmath.Vec4) int {
    this.Normals = append(this.Normals,norm)
    return len(this.Normals) -1
}
func ( this* Mesh) AddTex(tex lmath.Vec4) int {
    this.Texs = append(this.Texs,tex)
    return len(this.Texs) -1
}

func ( this* Mesh) AddTriVert(a,b,c int) {
    this.TriVerts = append(this.TriVerts,[3]int{a,b,c})
}
func ( this* Mesh) AddTriNormal(a,b,c int) {
    this.TriNormals = append(this.TriNormals,[3]int{a,b,c})
}
func ( this* Mesh) AddTriTex(a,b,c int) {
    this.TriTexs = append(this.TriTexs,[3]int{a,b,c})
}

func (this *Mesh) TriVert4(index int)(lmath.Vec4,lmath.Vec4,lmath.Vec4){
    return this.Verts[this.TriVerts[index][0]],
        this.Verts[this.TriVerts[index][1]],
        this.Verts[this.TriVerts[index][2]]
}

func (this *Mesh) TriNormal4(index int)(lmath.Vec4,lmath.Vec4,lmath.Vec4){
    return this.Normals[this.TriNormals[index][0]],
        this.Normals[this.TriNormals[index][1]],
        this.Normals[this.TriNormals[index][2]]
}
func (this *Mesh) TriTex4(index int)(lmath.Vec4,lmath.Vec4,lmath.Vec4){
    return this.Texs[this.TriTexs[index][0]],
        this.Texs[this.TriTexs[index][1]],
        this.Texs[this.TriTexs[index][2]]
}

func (this *Mesh) TriVert3(index int)(lmath.Vec3,lmath.Vec3,lmath.Vec3){
    return this.Verts[this.TriVerts[index][0]].Vec3(),
        this.Verts[this.TriVerts[index][1]].Vec3(),
        this.Verts[this.TriVerts[index][2]].Vec3()

}
func (this *Mesh) TriNormal3(index int)(lmath.Vec3,lmath.Vec3,lmath.Vec3){
    return this.Normals[this.TriNormals[index][0]].Vec3(),
        this.Normals[this.TriNormals[index][1]].Vec3(),
        this.Normals[this.TriNormals[index][2]].Vec3()
}
func (this *Mesh) TriTex3(index int)(lmath.Vec3,lmath.Vec3,lmath.Vec3){
    return this.Texs[this.TriTexs[index][0]].Vec3(),
        this.Texs[this.TriTexs[index][1]].Vec3(),
        this.Texs[this.TriTexs[index][2]].Vec3()
}