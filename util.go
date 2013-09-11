package glvec

import "math"

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
