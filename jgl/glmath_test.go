package jgl

import (
	// "fmt"
	"github.com/Stymphalian/go.math/lmath"
	"math"
	"testing"
)

func TestOrtho(t *testing.T) {
	cases := []struct {
		left, right, bottom, top, near, far float64
		want                                [16]float32
	}{
		{-1, 1, -1, 1, 1.0, 1500.0,
			[16]float32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, -0.0013342228, 0, -0, -0, -1.0013342, 1}},
		{-2, 2, -2, 2, 5.0, 25.0,
			[16]float32{0.5, 0, 0, 0, 0, 0.5, 0, 0, 0, 0, -0.1, 0, -0, -0, -1.5, 1}},
		{0, 4, 0, 4, 0.01, 1000.0,
			[16]float32{0.5, 0, 0, 0, 0, 0.5, 0, 0, 0, 0, -0.00200002, 0, -1, -1, -1.00002, 1}},
		{-4, 0, -4, 0, 1.0, 300.0,
			[16]float32{0.5, 0, 0, 0, 0, 0.5, 0, 0, 0, 0, -0.006688963, 0, 1, 1, -1.006689, 1}},
		{-4, 0, 0, 8, 1.0, 300.0,
			[16]float32{0.5, 0, 0, 0, 0, 0.25, 0, 0, 0, 0, -0.006688963, 0, 1, -1, -1.006689, 1}},
	}

	for testIndex, c := range cases {
		m := Ortho(c.left, c.right, c.bottom, c.top, c.near, c.far)
		get := m.DumpOpenGLf32()

		for i, _ := range get {
			if !closeEq(get[i], c.want[i], epsilon) {
				t.Errorf("TestOrthoMat4 %d %d", testIndex, i)
				break
			}
		}
	}
}

func TestFrustum(t *testing.T) {
	cases := []struct {
		left, right, bottom, top, near, far float64
		want                                [16]float32
	}{
		{-1, 1, -1, 1, 1.0, 1500.0,
			[16]float32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, -1.0013342, -1, 0, 0, -2.0013342, 0}},
		{-2, 2, -2, 2, 5.0, 25.0,
			[16]float32{2.5, 0, 0, 0, 0, 2.5, 0, 0, 0, 0, -1.5, -1, 0, 0, -12.5, 0}},
		{0, 4, 0, 4, 0.01, 1000.0,
			[16]float32{0.005, 0, 0, 0, 0, 0.005, 0, 0, 1, 1, -1.00002, -1, 0, 0, -0.0200002, 0}},
		{-4, 0, -4, 0, 1.0, 300.0,
			[16]float32{0.5, 0, 0, 0, 0, 0.5, 0, 0, -1, -1, -1.006689, -1, 0, 0, -2.006689, 0}},
		{-4, 0, 0, 8, 250.0, 300.0,
			[16]float32{125, 0, 0, 0, 0, 62.5, 0, 0, -1, 1, -11, -1, 0, 0, -3000, 0}},
		{-2, 10, -6, 40, 250.0, 300.0,
			[16]float32{41.666668, 0, 0, 0, 0, 10.869565, 0, 0, 0.6666667, 0.73913044, -11, -1, 0, 0, -3000, 0}},
	}

	for testIndex, c := range cases {
		m := Frustum(c.left, c.right, c.bottom, c.top, c.near, c.far)
		get := m.DumpOpenGLf32()

		for i, _ := range get {
			if !closeEq(get[i], c.want[i], epsilon) {
				t.Errorf("TestFrustum %d %d", testIndex, i)
				break
			}
		}
	}
}

func TestPerspective(t *testing.T) {
	cases := []struct {
		fov_y, aspect, near, far float64
		want                     [16]float32
	}{
		{math.Pi / 2, 1, 1.0, 1500.0, [16]float32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, -1.0013342, -1, 0, 0, -2.0013342, 0}},
		{lmath.Radians(45), 6, 1400.0, 1500.0, [16]float32{0.40236893, 0, 0, 0, 0, 2.4142137, 0, 0, 0, 0, -29, -1, 0, 0, -42000, 0}},
		{lmath.Radians(87), 0.5, 0.01, 40, [16]float32{2.1075602, 0, 0, 0, 0, 1.0537801, 0, 0, 0, 0, -1.0005001, -1, 0, 0, -0.020005, 0}},
		{lmath.Radians(165), 2, 1.0, 45, [16]float32{0.06582625, 0, 0, 0, 0, 0.1316525, 0, 0, 0, 0, -1.0454545, -1, 0, 0, -2.0454545, 0}},
	}

	for testIndex, c := range cases {
		m := Perspective(c.fov_y, c.aspect, c.near, c.far)
		get := m.DumpOpenGLf32()

		for i, _ := range get {
			if !closeEq(get[i], c.want[i], epsilon) {
				t.Errorf("TestPerspective %d %d", testIndex, i)
				break
			}
		}
	}
}

func TestFocalLength(t *testing.T) {
	cases := []struct {
		fov_y, height float64
		want          float32
	}{
		{math.Pi, 480, 0},
		{lmath.Radians(180), 6000, 0},
		{lmath.Radians(200), 55, -4.848991969482785},
		{lmath.Radians(50), 1, 1.0722534602547793},
		{lmath.Radians(90), 0.5, 0.25},
		{lmath.Radians(123), 60, 16.288670989153108},
	}

	for testIndex, c := range cases {
		get := FocalLength(c.fov_y, c.height)
		if !closeEq(float32(get), c.want, epsilon) {
			t.Errorf("TestFocalLength %d", testIndex)
		}
	}
}

func TestLookAt(t *testing.T) {
	cases := []struct {
		eye, at, up lmath.Vec3
		want        [16]float32
	}{
		{lmath.Vec3{1, 0, 0}, lmath.Vec3{0, 0, 0}, lmath.Vec3{0, 1, 0},
			[16]float32{0, 0, 1, 0, 0, 1, 0, 0, -1, 0, 0, 0, 0, 0, -1, 1}},
		{lmath.Vec3{0, 0, 1}, lmath.Vec3{40, 0, 0}, lmath.Vec3{0, 1, 0},
			[16]float32{0.02499219, 0, -0.9996877, 0, 0, 1, 0, 0, 0.9996877, 0, 0.02499219, 0, -0.9996877, 0, -0.02499219, 1}},
		{lmath.Vec3{10, 0, 0}, lmath.Vec3{0, 0, 0.5}, lmath.Vec3{0, 1, 0},
			[16]float32{-0.049937617, 0, 0.99875236, 0, 0, 1, 0, 0, -0.99875236, 0, -0.049937617, 0, 0.49937618, 0, -9.987523, 1}},
		{lmath.Vec3{0, 0, 10}, lmath.Vec3{6, 0, 9}, lmath.Vec3{0, 1, 0},
			[16]float32{0.16439898, 0, -0.9863939, 0, 0, 1, 0, 0, 0.9863939, 0, 0.16439898, 0, -9.863939, 0, -1.6439899, 1}},
		{lmath.Vec3{0.5, 0, 0}, lmath.Vec3{0, 2, 0}, lmath.Vec3{0, 1, 0},
			[16]float32{0, 0.9701425, 0.24253562, 0, 0, 0.24253562, -0.9701425, 0, -1, 0, 0, 0, 0, -0.48507124, -0.12126781, 1}},
		{lmath.Vec3{0, 0, 0.5}, lmath.Vec3{0, 4, 5}, lmath.Vec3{0, 1, 0},
			[16]float32{-1, 0, 0, 0, 0, 0.74740934, -0.66436386, 0, 0, -0.66436386, -0.74740934, 0, 0, 0.33218193, 0.37370467, 1}},
		{lmath.Vec3{0.5, 5, 0}, lmath.Vec3{1, 1, 1}, lmath.Vec3{0, 1, 0},
			[16]float32{-0.8944272, 0.43070552, -0.120385855, 0, 0, 0.26919094, 0.96308684, 0, 0.4472136, 0.86141104, -0.24077171, 0, 0.4472136, -1.5613075, -4.7552414, 1}},
		{lmath.Vec3{0, 0.5, 5}, lmath.Vec3{0, 0, 2}, lmath.Vec3{0, 1, 0},
			[16]float32{1, 0, 0, 0, 0, 0.9863939, 0.16439898, 0, 0, -0.16439898, 0.9863939, 0, 0, 0.32879797, -5.014169, 1}},
		{lmath.Vec3{0, 1, 0.5}, lmath.Vec3{0, 50, 0}, lmath.Vec3{0, 1, 0},
			[16]float32{1, 0, 0, 0, 0, 0.010203551, -0.99994797, 0, 0, 0.99994797, 0.010203551, 0, 0, -0.5101775, 0.99484617, 1}},
		{lmath.Vec3{1, 0, 0}, lmath.Vec3{0, 0, 0}, lmath.Vec3{0, 1, 1},
			[16]float32{0, 0, 1, 0, 0.70710677, 0.70710677, 0, 0, -0.70710677, 0.70710677, 0, 0, 0, 0, -1, 1}},
		{lmath.Vec3{0, 1, 0}, lmath.Vec3{0, 0, 0}, lmath.Vec3{1, 0, 0},
			[16]float32{0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, -1, 1}},
		{lmath.Vec3{0, 0, 1}, lmath.Vec3{40, 0, 0}, lmath.Vec3{0, 1, 5},
			[16]float32{0.0049028443, 0.024506565, -0.9996877, 0, -0.9805689, 0.19617505, 0, 0, 0.19611378, 0.9802626, 0.02499219, 0, -0.19611378, -0.9802626, -0.02499219, 1}},
		{lmath.Vec3{10, 0, 0}, lmath.Vec3{0, 0, 0.5}, lmath.Vec3{4, 3, -4},
			[16]float32{-0.030967355, -0.03917638, 0.99875236, 0, -0.7845063, 0.6201208, 0, 0, -0.6193471, -0.78352755, -0.049937617, 0, 0.30967355, 0.39176378, -9.987523, 1}},
	}

	for testIndex, c := range cases {
		m := LookAt(c.eye, c.at, c.up)
		get := m.DumpOpenGLf32()

		for i, _ := range get {
			if !closeEq(get[i], c.want[i], epsilon) {
				t.Errorf("TestLookAt %d %d", testIndex, i)
				break
			}
		}
	}
}

func TestLookAtVec3(t *testing.T) {
	cases := []struct {
		eye, at, up                       lmath.Vec3
		want_forward, want_right, want_up lmath.Vec3
	}{
		{lmath.Vec3{1, 0, 0}, lmath.Vec3{0, 0, 0}, lmath.Vec3{0, 1, 0},
			lmath.Vec3{1, -0, -0}, lmath.Vec3{0, 0, -1}, lmath.Vec3{0, 1, 0}},
		{lmath.Vec3{0, 0, 1}, lmath.Vec3{40, 0, 0}, lmath.Vec3{0, 1, 0},
			lmath.Vec3{-0.9996876464081228, -0, 0.02499219116020307}, lmath.Vec3{0.02499219116020307, -0, 0.9996876464081228}, lmath.Vec3{0, 1, 0}},
		{lmath.Vec3{10, 0, 0}, lmath.Vec3{0, 0, 0.5}, lmath.Vec3{0, 1, 0},
			lmath.Vec3{0.9987523388778445, -0, -0.04993761694389223}, lmath.Vec3{-0.049937616943892246, 0, -0.9987523388778448}, lmath.Vec3{0, 1, 0}},
		{lmath.Vec3{0, 0, 10}, lmath.Vec3{6, 0, 9}, lmath.Vec3{0, 1, 0},
			lmath.Vec3{-0.9863939238321437, -0, 0.1643989873053573}, lmath.Vec3{0.1643989873053573, -0, 0.9863939238321437}, lmath.Vec3{0, 1, 0}},
		{lmath.Vec3{0.5, 0, 0}, lmath.Vec3{0, 2, 0}, lmath.Vec3{0, 1, 0},
			lmath.Vec3{0.24253562503633297, -0.9701425001453319, -0}, lmath.Vec3{0, 0, -1}, lmath.Vec3{0.9701425001453319, 0.24253562503633297, 0}},
		{lmath.Vec3{0, 0, 0.5}, lmath.Vec3{0, 4, 5}, lmath.Vec3{0, 1, 0},
			lmath.Vec3{-0, -0.6643638388299198, -0.7474093186836597}, lmath.Vec3{-1, -0, 0}, lmath.Vec3{-0, 0.7474093186836597, -0.6643638388299198}},
		{lmath.Vec3{0.5, 5, 0}, lmath.Vec3{1, 1, 1}, lmath.Vec3{0, 1, 0},
			lmath.Vec3{-0.1203858530857692, 0.9630868246861536, -0.2407717061715384}, lmath.Vec3{-0.8944271909999159, -0, 0.4472135954999579}, lmath.Vec3{0.4307055216465324, 0.2691909510290828, 0.8614110432930648}},
		{lmath.Vec3{0, 0.5, 5}, lmath.Vec3{0, 0, 2}, lmath.Vec3{0, 1, 0},
			lmath.Vec3{-0, 0.1643989873053573, 0.9863939238321437}, lmath.Vec3{1, -0, 0}, lmath.Vec3{0, 0.9863939238321437, -0.1643989873053573}},
		{lmath.Vec3{0, 1, 0.5}, lmath.Vec3{0, 50, 0}, lmath.Vec3{0, 1, 0},
			lmath.Vec3{-0, -0.999947942424286, 0.010203550432900877}, lmath.Vec3{1, -0, 0}, lmath.Vec3{0, 0.010203550432900879, 0.9999479424242861}},
		{lmath.Vec3{1, 0, 0}, lmath.Vec3{0, 0, 0}, lmath.Vec3{0, 1, 1},
			lmath.Vec3{1, -0, -0}, lmath.Vec3{0, 0.7071067811865475, -0.7071067811865475}, lmath.Vec3{0, 0.7071067811865476, 0.7071067811865476}},
		{lmath.Vec3{0, 1, 0}, lmath.Vec3{0, 0, 0}, lmath.Vec3{1, 0, 0},
			lmath.Vec3{-0, 1, -0}, lmath.Vec3{-0, -0, 1}, lmath.Vec3{1, 0, 0}},
		{lmath.Vec3{0, 0, 1}, lmath.Vec3{40, 0, 0}, lmath.Vec3{0, 1, 5},
			lmath.Vec3{-0.9996876464081228, -0, 0.02499219116020307}, lmath.Vec3{0.004902844450389859, -0.9805688900779719, 0.19611377801559435}, lmath.Vec3{0.024506565146576818, 0.1961750539983474, 0.9802626058630728}},
		{lmath.Vec3{10, 0, 0}, lmath.Vec3{0, 0, 0.5}, lmath.Vec3{4, 3, -4},
			lmath.Vec3{0.9987523388778445, -0, -0.04993761694389223}, lmath.Vec3{-0.030967355248758714, -0.784506332968554, -0.6193471049751742}, lmath.Vec3{-0.03917637674584123, 0.6201208055953552, -0.7835275349168245}},
	}

	for testIndex, c := range cases {
		get_forward, get_right, get_up := LookAtVec3(c.eye, c.at, c.up)
		if !get_forward.Eq(c.want_forward) ||
			!get_right.Eq(c.want_right) ||
			!get_up.Eq(c.want_up) {
			t.Errorf("TestLookAtVec3 %d", testIndex)
		}

	}
}
