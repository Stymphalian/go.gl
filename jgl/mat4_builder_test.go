package jgl

import (
	"testing"
	// "fmt"
	"github.com/Stymphalian/go.math/lmath"
)

func TestMakeMat4Builder(t *testing.T) {
	s := NewMat4Builder()
	if s.Empty() == true {
		t.Errorf("Empty stack")
	}
}

func TestMat4Builder(t *testing.T) {
	s := NewMat4Builder()
	s.LoadIdentity()

	m := lmath.Mat4{}
	m.Load([16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	s.Load(m)

	s.Push()
	s.LoadIdentity()

	s.Push()
	s.Translate(1, 2, 3)

	s.Push()
	s.LoadIdentity()
	s.Rotate(lmath.Radians(30), lmath.Radians(68), lmath.Radians(4))
	s.Scale(1, 1, 2)
	get := s.Peek()

	m1 := lmath.Mat4{}
	m1.FromEuler(lmath.Radians(30), lmath.Radians(68), lmath.Radians(4))
	m2 := lmath.Mat4{}
	m2.ToScale(1, 1, 2)
	want := m1.Mult(m2)

	if !get.Eq(want) {
		t.Errorf("Fail 1")
	}
	s.Pop()

	m1.ToTranslate(1, 2, 3)
	get = s.Peek()
	if !get.Eq(m1) {
		t.Errorf("Fail 2")
	}
	s.Pop()

	m1.ToIdentity()
	get = s.Peek()
	if !get.Eq(m1) {
		t.Errorf("Fail 3")
	}
	s.Pop()

	get = s.Peek()
	if !get.Eq(m) {
		t.Errorf("Fail 4")
	}

	s.Pop()
	if s.Empty() != true {
		t.Errorf("Fail 5")
	}
}
