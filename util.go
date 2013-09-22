package glvec

import (
	"fmt"
	"math"
)

func M4Perspective(result *Mat4, fovy, aspect, near, far float32) {
	nmf := near - far
	f := float32( 1./math.Tan( float64(fovy)/2.0 ) )

	result[0] = float32(f / aspect)
	result[1] = 0
	result[2] = 0
	result[3] = 0
	result[4] = 0
	result[5] = float32(f)
	result[6] = 0
	result[7] = 0
	result[8] = 0
	result[9] = 0
	result[10] = float32((near + far) / nmf)
	result[11] = -1
	result[12] = 0
	result[13] = 0
	result[14] = float32((2. * far * near) / nmf)
	result[15] = 0
}

func M4LookAt(result *Mat4, eye, center, up Vec3) {
	var F Vec3
	V3Sub(&F, center, eye)
	V3Normalize(&F, F)

	var S, U Vec3
	V3Cross(&S, F, up)
	V3Normalize(&S, S)

	V3Cross(&U, S, F)

	M := Mat4 {
		S[0], S[1], S[2], 0,
		U[0], U[1], U[2], 0, 
		-F[0], -F[1], -F[2], 0, 
		0, 0, 0, 1,
	}
	M4Transpose(&M, &M)
	
	var MT Mat4
	V3Neg(&eye, eye)
	M4MakeTransform(&MT, eye)
	M4MulM4(result, &M, &MT)
}


func (m Mat4) Print() string {
	return fmt.Sprintf("[ %2.2f, %2.2f, %2.2f, %2.2f,\n  %2.2f, %2.2f, %2.2f, %2.2f,\n  %2.2f, %2.2f, %2.2f, %2.2f,\n  %2.2f, %2.2f, %2.2f, %2.2f ]",
		m[0], m[1], m[2], m[3], m[4], m[5], m[6], m[7], 
		m[8], m[9], m[10], m[11], m[12], m[13], m[14], m[15])
}
