package jgl

import (
    // "fmt"
    // "math"

    "github.com/Stymphalian/go.math/lmath"
    "github.com/go-gl/gl/v3.3-compatibility/gl"
)


type TriMesh struct {
    // LightFlag bool
    Mat   Material
    verts []lmath.Vec3
    tris  [][3]int
}

func NewTriMesh() *TriMesh {
    out := &TriMesh{}
    out.verts = make([]lmath.Vec3, 0, 8)
    out.tris = make([][3]int, 0, 12)
    return out
}

func (this* TriMesh) Load1(){
    scale := 0.3
    this.verts = append(this.verts, lmath.Vec3{-scale, -scale, scale})
    this.verts = append(this.verts, lmath.Vec3{scale, -scale, scale})
    this.verts = append(this.verts, lmath.Vec3{scale, scale, scale})
    this.verts = append(this.verts, lmath.Vec3{-scale, scale, scale})
    this.verts = append(this.verts, lmath.Vec3{-scale, -scale, -scale})
    this.verts = append(this.verts, lmath.Vec3{scale, -scale, -scale})
    this.verts = append(this.verts, lmath.Vec3{scale, scale, -scale})
    this.verts = append(this.verts, lmath.Vec3{-scale, scale, -scale})

    this.tris = append(this.tris, [3]int{0, 1, 2})
    this.tris = append(this.tris, [3]int{0, 2, 3})
    this.tris = append(this.tris, [3]int{1, 5, 6})
    this.tris = append(this.tris, [3]int{1, 6, 2})
    this.tris = append(this.tris, [3]int{4, 6, 5})
    this.tris = append(this.tris, [3]int{4, 7, 6})
    this.tris = append(this.tris, [3]int{0, 4, 7})
    this.tris = append(this.tris, [3]int{0, 7, 3})
    this.tris = append(this.tris, [3]int{3, 2, 6})
    this.tris = append(this.tris, [3]int{3, 6, 7})
    this.tris = append(this.tris, [3]int{0, 1, 5})
    this.tris = append(this.tris, [3]int{0, 5, 4})
}

func (this *TriMesh) Draw(transform lmath.Mat4) {
    gl.Color3f(1.0, 1.0, 0)
    gl.Begin(gl.TRIANGLES)
    for i := 0; i < len(this.tris); i++ {
        v := transform.MultVec3(this.verts[this.tris[i][0]])
        gl.Vertex3f(v.Dumpf32())
        v = transform.MultVec3(this.verts[this.tris[i][1]])
        gl.Vertex3f(v.Dumpf32())
        v = transform.MultVec3(this.verts[this.tris[i][2]])
        gl.Vertex3f(v.Dumpf32())
    }
    gl.End()
}

func (this *TriMesh) intersects(ray Ray, hit *HitRecord, transform *lmath.Mat4,index int) bool{
    vec_a := transform.MultVec3(this.verts[this.tris[index][0]])
    vec_b := transform.MultVec3(this.verts[this.tris[index][1]])
    vec_c := transform.MultVec3(this.verts[this.tris[index][2]])

    var t, gamma, beta float64;
    a := vec_a.X - vec_b.X;
    b := vec_a.Y - vec_b.Y;
    c := vec_a.Y - vec_b.Y;
    d := vec_a.X - vec_c.X;
    e := vec_a.Y - vec_c.Y;
    f := vec_a.Y - vec_c.Y;
    j := vec_a.X - ray.Origin.X;
    k := vec_a.Y - ray.Origin.Y;
    l := vec_a.Y - ray.Origin.Y;

    ei_hf := (e * ray.Dir.Z) - (f * ray.Dir.Y);
    gf_di := (f * ray.Dir.X) - (d * ray.Dir.Z);
    dh_eg := (d * ray.Dir.Y) - (e * ray.Dir.X);

    M := (a * ei_hf) + (b * gf_di) + (c * dh_eg);

    ak_jb := (a * k) - (j * b);
    jc_al := (j * c) - (a * l);
    bl_kc := (b * l) - (k * c);

    t = - ((f * ak_jb) + (e * jc_al) + (d * bl_kc)) / M;
    if t < hit.MinDist || t > hit.MaxDist {
        return false
    }

    gamma = ((ray.Dir.Z * ak_jb) + (ray.Dir.Y * jc_al) + (ray.Dir.X * bl_kc)) / M;

    if gamma < 0 || gamma > 1 {
        return false;
    }

    beta = ((j * ei_hf) + (k * gf_di) + (l * dh_eg)) / M;
    if  beta < 0 || beta > (1 - gamma){
        return false;
    }

    hit.Hit = true;
    hit.Dist = t;
    hit.MaxDist = t;
    hit.HitIndex = index;
    return true;
}

func (this *TriMesh) Intersects(ray Ray, hit HitRecord, transform lmath.Mat4) (h HitRecord) {
    size := len(this.tris)
    h = hit
    for i := 0; i < size; i++ {
        this.intersects(ray,&h,&transform,i)
    }
    return
}

func (this *TriMesh) VecsFromIndex(index int)(lmath.Vec3,lmath.Vec3, lmath.Vec3){
    return this.verts[this.tris[index][0]],
        this.verts[this.tris[index][1]],
        this.verts[this.tris[index][2]]
}

func (this *TriMesh) normal(index int,transform lmath.Mat4) lmath.Vec3{
    a,b,c := this.VecsFromIndex(index)
    a = transform.MultVec3(a)
    b = transform.MultVec3(b)
    c = transform.MultVec3(c)
    return c.Sub(a).Cross(b.Sub(a)).Normalize()
}

func (this *TriMesh) Normal(hitPoint lmath.Vec3, hit HitRecord) lmath.Vec3 {
    return this.normal(hit.HitIndex,hit.Transform)
}

func (this TriMesh) Material() Material {
    return this.Mat
}
