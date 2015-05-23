package jgl

import (
	"encoding/json"
	"github.com/Stymphalian/go.math/lmath"
	"image/color"
	"math"
)

type color4 struct {
	lmath.Vec4 // unnamed field implies an "is-a" relationship
}

func (this color4) ToRGBA() color.RGBA {
	return color.RGBA{
		uint8(this.X * 255),
		uint8(this.Y * 255),
		uint8(this.Z * 255),
		uint8(this.W * 255),
	}
}

func (this color4) Map(f func(x float64) float64) (out color4) {
	out.X = f(this.X)
	out.Y = f(this.Y)
	out.Z = f(this.Z)
	out.W = f(this.W)
	return
}

func (this color4) Clamp(min, max float64) color4 {
	this.X = lmath.Clamp(this.X, min, max)
	this.Y = lmath.Clamp(this.Y, min, max)
	this.Z = lmath.Clamp(this.Z, min, max)
	this.W = lmath.Clamp(this.W, min, max)
	return this
}

func (this color4) Min(min float64) color4 {
	this.X = math.Min(this.X, min)
	this.Y = math.Min(this.Y, min)
	this.Z = math.Min(this.Z, min)
	this.W = math.Min(this.W, min)
	return this
}

func (this color4) Max(max float64) color4 {
	this.X = math.Max(this.X, max)
	this.Y = math.Max(this.Y, max)
	this.Z = math.Max(this.Z, max)
	this.W = math.Max(this.W, max)
	return this
}

func (this color4) MarshalJSON() ([]byte, error) {
	data := [4]float64{this.X, this.Y, this.Z, this.W}
	return json.Marshal(data)
}

func (this *color4) UnmarshalJSON(data []byte) (err error) {
	var d [4]float64
	err = json.Unmarshal(data, &d)
	if err != nil {
		return
	}

	this.X, this.Y, this.Z, this.W = d[0], d[1], d[2], d[3]
	return nil
}
