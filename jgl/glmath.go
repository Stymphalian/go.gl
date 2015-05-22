package jgl

import (
	"github.com/Stymphalian/go.math/lmath"
	"math"
	// "fmt"
)

// The below projection/perspective matrices use the derivations from here:
//http://www.songho.ca/opengl/gl_projectionmatrix.html

//Creates and orthographic projection matrix from the given parameters.
//The matrix is given in a right-handed system, but will transform
//a vector into a left-handed system (to match the NDC space of OpenGL).
//	precondition : far > near > 0.
func Ortho(left, right, bottom, top, near, far float64) (out lmath.Mat4) {
	out.Load([16]float64{
		2 / (right - left), 0, 0, -(right + left) / (right - left),
		0, 2 / (top - bottom), 0, -(top + bottom) / (top - bottom),
		//0,0,2/(n-f), -(n+f)/(n-f)
		0, 0, -2 / (far - near), -(far + near) / (far - near),
		0, 0, 0, 1})
	return
}

//Create a perspective frustum matrix from the provided parameters.
//The matrix is given in a right-handed system, but will transform
//a vector into a left-handed sytsem ( to match the NDC space of OpenGL).
//	precondition : far > near > 0
func Frustum(left, right, bottom, top, near, far float64) (out lmath.Mat4) {
	out.Load([16]float64{
		2 * near / (right - left), 0, (right + left) / (right - left), 0,
		0, 2 * near / (top - bottom), (top + bottom) / (top - bottom), 0,
		0, 0, -(far + near) / (far - near), -2 * far * near / (far - near),
		0, 0, -1, 0})
	return
}

//Creates a normalized viewing frustum using the given perspective parameters.
//	fov ( y-direction)  angle in radians
//	aspect - ratio between the width and the height (width/height)
//	precondition: far > near > 0
func Perspective(fov_y, aspect, near, far float64) (out lmath.Mat4) {
	top := math.Tan(fov_y/2) * near
	right := top * aspect
	return Frustum(-right, right, -top, top, near, far)
}

// Retrieve the Field of View from the Perspective Matrix
// http://paulbourke.net/miscellaneous/lens/
func FocalLength(fov_y, height float64) float64 {
	return (0.5 * height) / (math.Tan(fov_y / 2))
}

//Create a LookAt rotation matrix.
//	eye is the position of the camera.
//	at is the position in which the camera "looksAt".
//	up is the direction which is considered up. It is up to the user to ensure that the forward dir is not parallel to up.
func LookAt(eye, at, up lmath.Vec3) (out lmath.Mat4) {
	// up := &Vec3{0, 1, 0}
	// if at.Eq(&Vec3{0, 1, 0}) {
	//  up = &Vec3{0, 0, -1}
	// } else if at.Eq(&Vec3{0, -1, 0}) {
	//  up = &Vec3{0, 0, 1}
	// }

	forward := at.Sub(eye).Normalize().MultScalar(-1.0)
	right := up.Cross(forward).Normalize()
	up = forward.Cross(right).Normalize()

	translate := lmath.Mat4{}
	translate.ToTranslate(-eye.X, -eye.Y, -eye.Z)
	rot := lmath.Mat4{}
	rot.Load([16]float64{right.X, right.Y, right.Z, 0,
		up.X, up.Y, up.Z, 0,
		forward.X, forward.Y, forward.Z, 0,
		0, 0, 0, 1})
	return rot.Mult(translate)
}

// Return the forward, right and up basis formed by the specified vectors
//	eye is the position of the camera.
//	at is the position in which the camera "looks at".
//	up is the direction which is considered up. It is up to the user to ensure that the forward dir is not parallel to up.
func LookAtVec3(eye, at, up_in lmath.Vec3) (forward, right, up lmath.Vec3) {
	forward = at.Sub(eye).Normalize().MultScalar(-1.0)
	// forward = at.Sub(eye).Normalize()
	right = up_in.Cross(forward).Normalize()
	up = forward.Cross(right).Normalize()
	return
}

// Take the given screen coordinates on the image-plane and convert
// into world-space coordinates
func UnProject(screenPos lmath.Vec3,
	perspectiveMat lmath.Mat4,
	modelViewMat lmath.Mat4,
	viewport [4]int32,
) (worldPoint lmath.Vec3) {

	// calcualte the inverse matrix to undot he modelView and perspective matrices
	finalMat := modelViewMat.Mult(perspectiveMat).Inverse()

	x, y, z := screenPos.X, float64(viewport[3])-screenPos.Y, 0.0
	// translate to origin, then scale down into the range [0,1]
	x = (x - float64(viewport[0])) / float64(viewport[2])
	y = (y - float64(viewport[1])) / float64(viewport[3])
	// scale to width 2, then translate by -1, so that coordinates are in [-1,1]
	x = x*2 - 1
	y = y*2 - 1
	z = z*2 - 1
	worldPoint.X = x
	worldPoint.Y = y
	worldPoint.Z = z
	// Multiply the inverse matrix against the point
	worldPoint = worldPoint.MultMat4(finalMat)
	return
}

//Specify a matrix which transforms a vector from NDC space into screen space
// Idea is to translate from the range [-1,1] -> [0,2]
// Apply a scale from 2 -> width (or height)
// translate by the specified x and y
// What about z??
func ViewportMat(x, y, width, height int) (out lmath.Mat4) {
	// fmt.Println("Unimplemented!")
	// out.Load([16]float64{
	// 		w/2, 0,0,(w-1)/2,
	// 		0, h/2,0,(h-1)/2,
	// 		0,0,1,0,
	// 		0,0,0,1
	// })
	return
}
