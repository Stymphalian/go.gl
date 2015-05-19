package jgl

import (
	// "fmt"
	"github.com/Stymphalian/go.math/lmath"
)

type SceneNode struct {
	SceneObject    Primitive
	Children       []*SceneNode
	Parent         *SceneNode
	localTransform lmath.Mat4
	realTransform  lmath.Mat4
}

func NewSceneNode() (out *SceneNode) {
	out = &SceneNode{}
	out.Children = make([]*SceneNode, 0, 100)
	out.Parent = nil
	out.localTransform.ToIdentity()
	out.realTransform.ToIdentity()
	return
}

func (this *SceneNode) Draw() {
	if this.SceneObject != nil {
		this.SceneObject.Draw(this.realTransform)
	}

	for _, v := range this.Children {
		v.Draw()
	}
}

func (this *SceneNode) Query(r Ray, minDist, maxDist float64) HitRecord {
	hit := HitRecord{}
	hit.MinDist = minDist
	hit.MaxDist = maxDist
	hit.Hit = false
	hit.Transform.ToIdentity()

	return this.query(r, hit)
}

func (this *SceneNode) query(ray Ray, hit HitRecord) HitRecord {
	candHitRecord := hit
	candHitRecord.Hit = false

	if this.SceneObject != nil {
		candHitRecord = this.SceneObject.Intersects(ray, candHitRecord, this.realTransform)
		if candHitRecord.Hit {
			candHitRecord.HitObject = this.SceneObject
			candHitRecord.Transform = this.realTransform

			// record the results out in the hit record
			hit = candHitRecord
		}
	}

	// recursively query all the children
	for _, v := range this.Children {
		hit = v.query(ray, hit)
	}

	return hit
}

func (this *SceneNode) Add(s *SceneNode) {
	s.Parent = this
	s.realTransform = s.localTransform.Mult(this.realTransform)
	this.Children = append(this.Children, s)
}
func (this *SceneNode) Clear() {
	this.Children = make([]*SceneNode, 0, 100)
}

func (this *SceneNode) LocalTransform() lmath.Mat4 {
	return this.localTransform
}
func (this *SceneNode) SetLocalTransform(transform lmath.Mat4) {
	this.localTransform = transform
	if this.Parent != nil {
		this.realTransform = this.localTransform.Mult(this.Parent.realTransform)
	}
}
func (this *SceneNode) RealTransform() lmath.Mat4 {
	return this.realTransform
}
