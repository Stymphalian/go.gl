package jgl

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
    "github.com/Stymphalian/go.math/lmath"
)

// Verts []lmath.Vec3
// Normals []lmath.Vec3
// Texs []lmath.Vec3
// Tris [][3]int

func ReadOBJFile(filename string) (out Mesh, ok bool){
    fin, err := os.Open(filename)
    if err != nil {
        fmt.Println("Unable to open file : ", filename)
        ok= false
        return
    }
    defer fin.Close()

    scanner := bufio.NewScanner(fin)
    lineNum := 0
    for scanner.Scan(){
        line := scanner.Text();
        line = strings.Trim(line," \t\n")
        lineNum += 1

        // skip empty lines, skip comments,
        if len(line) == 0 ||
            strings.HasPrefix(line,"#") {
            continue
        }

        tokens := strings.Split(line," ")
        header,rest := tokens[0],tokens[1:]
        var v lmath.Vec4
        if header == "v"  || header == "vt" || header == "vn" || header == "vp" {
            var x,y,z,w float64
            if len(rest) == 1 {
                x,_ = strconv.ParseFloat(rest[0],64)
            }else if len(rest) == 2 {
                x,_ = strconv.ParseFloat(rest[0],64)
                y,_ = strconv.ParseFloat(rest[1],64)
            }else if len(rest) == 3 {
                x,_ = strconv.ParseFloat(rest[0],64)
                y,_ = strconv.ParseFloat(rest[1],64)
                z,_ = strconv.ParseFloat(rest[2],64)
            }else if len(rest) == 4 {
                x,_ = strconv.ParseFloat(rest[0],64)
                y,_ = strconv.ParseFloat(rest[1],64)
                z,_ = strconv.ParseFloat(rest[2],64)
                w,_ = strconv.ParseFloat(rest[3],64)
            }
            v.Set(x,y,z,w)

            switch(header){
                case "v" :{
                    out.AddVert(v)
                }
                case "vt":{
                    out.AddTex(v)
                }
                case "vn":{
                    out.AddNormal(v)
                }
                case "vp":{
                    fmt.Printf("%d: We dont recognize vp\n",lineNum)
                }
            }

        }else if header == "f" {
            if( len(rest) != 3){
                fmt.Printf("%d: Too many values specified for f\n",lineNum)
                continue
            }

            var values [3][3]struct{
                v int
                ok bool
            }

            for i := 0; i < 3; i++ {
                var a,b,c int64
                var ok1,ok2,ok3 error
                r := strings.Split(rest[i],"/")
                fmt.Println(r)
                if len(r) == 1 {
                    a,ok1 = strconv.ParseInt(r[0],10,0)
                }else if len(r) == 2 {
                    a,ok1 = strconv.ParseInt(r[0],10,0)
                    b,ok2 = strconv.ParseInt(r[1],10,0)
                }else if len(r) == 3 {
                    a,ok1 = strconv.ParseInt(r[0],10,0)
                    b,ok2 = strconv.ParseInt(r[1],10,0)
                    c,ok3 = strconv.ParseInt(r[2],10,0)
                }

                values[i][0].v = int(a)
                values[i][0].ok = (ok1 == nil)
                values[i][1].v = int(b)
                values[i][1].ok = (ok2 == nil)
                values[i][2].v = int(c)
                values[i][2].ok = (ok3 == nil)
            }
            fmt.Println(values)
            // fill vertex tris
            if values[0][0].ok && values[1][0].ok && values[2][0].ok {
                out.AddTriVert(values[0][0].v,
                    values[1][0].v,
                    values[2][0].v)
            }
            // fill texture tris
            if values[0][1].ok && values[1][1].ok && values[2][1].ok {
                out.AddTriTex(values[0][1].v,
                    values[1][1].v,
                    values[2][1].v)
            }
            // fill normal tris
            if values[0][2].ok && values[1][2].ok && values[2][2].ok {
                out.AddTriNormal(values[0][2].v,
                    values[1][2].v,
                    values[2][2].v)
            }
        }else{
            fmt.Printf("Unrecognized tokens \"%s\":%d\n",header,lineNum)
        }
    }

    ok = true
    return
}
