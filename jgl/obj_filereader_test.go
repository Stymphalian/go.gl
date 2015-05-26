package jgl

import(
    "fmt"
    "testing"
)

func TestReadOBJFileReader(t *testing.T){
    out,_ := ReadOBJFile("testResources/test.obj");

    fmt.Println("Verts")
    for _,v := range(out.Verts){
        fmt.Println(v)
    }
    fmt.Println("Normal")
    for _,v := range(out.Normals){
        fmt.Println(v)
    }
    fmt.Println("Texs")
    for _,v := range(out.Texs){
        fmt.Println(v)
    }

    fmt.Println("triVerts")
    for _,v := range(out.TriVerts){
        fmt.Println(v)
    }
    fmt.Println("triNormals")
    for _,v := range(out.TriNormals){
        fmt.Println(v)
    }
    fmt.Println("triTexs")
    for _,v := range(out.TriTexs){
        fmt.Println(v)
    }
}