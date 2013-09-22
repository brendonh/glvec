package glvec

import "math"

type Mat4 [16]float32

func M4Ident(result *Mat4) {
	result[0] = 1.0
	result[1] = 0.0
	result[2] = 0.0
	result[3] = 0.0

	result[4] = 0.0
	result[5] = 1.0
	result[6] = 0.0
	result[7] = 0.0

	result[8] = 0.0
	result[9] = 0.0
	result[10] = 1.0
	result[11] = 0.0

	result[12] = 0.0
	result[13] = 0.0
	result[14] = 0.0
	result[15] = 1.0
}

func M4Copy(result, m *Mat4) {
	for i := 0; i < 16; i++ {
		result[i] = m[i]
	}
}

func M4MakeTransform(result *Mat4, trans Vec3) {
	M4Ident(result)
	M4SetTransform(result, trans)
}

func M4SetTransform(result *Mat4, trans Vec3) {
	result[12] = trans[0]
	result[13] = trans[1]
	result[14] = trans[2]
}


func M4MakeScale(result *Mat4, scale float32) {
	M4Ident(result)
	result[0] = scale
	result[5] = scale
	result[10] = scale
}


func M4MakeRotation(result *Mat4, radians float32, unitVec Vec3) {
	s := float32(math.Sin(float64(radians)))
	c := float32(math.Cos(float64(radians)))
	x := unitVec[0]
	y := unitVec[1]
	z := unitVec[2]
	xy := x * y
	yz := y * z
	zx := z * x
	oneMinusC := 1.0 - c

	result[0] = ((x * x) * oneMinusC) + c
	result[1] = (xy * oneMinusC) + (z * s)
	result[2] = (zx * oneMinusC) - (y * s)
	result[3] = 0.0

    result[4] = (xy * oneMinusC) - (z * s)
	result[5] = ((y * y) * oneMinusC) + c
	result[6] = (yz * oneMinusC) + (x * s)
	result[7] = 0.0

	result[8] = (zx * oneMinusC) + (y * s)
	result[9] = (yz * oneMinusC) - (x * s)
	result[10] = ((z * z) * oneMinusC) + c
	result[11] = 0.0

	result[12] = 0.0
	result[13] = 0.0
	result[14] = 0.0
	result[15] = 1.0
}


func M4Transpose(result, m *Mat4) {
	var temp = Mat4{
		m[0],  m[4],  m[8], m[12],
		m[1],  m[5],  m[9], m[13],
		m[2],  m[6],  m[10], m[14],
		m[3],  m[7],  m[11], m[15],
	}
	M4Copy(result, &temp)
}


func M4Inverse(result, m *Mat4) {
	var inv Mat4

    inv[0] = m[5]  * m[10] * m[15] -
             m[5]  * m[11] * m[14] -
             m[9]  * m[6]  * m[15] +
             m[9]  * m[7]  * m[14] +
             m[13] * m[6]  * m[11] -
             m[13] * m[7]  * m[10]

    inv[4] = -m[4]  * m[10] * m[15] +
              m[4]  * m[11] * m[14] +
              m[8]  * m[6]  * m[15] -
              m[8]  * m[7]  * m[14] -
              m[12] * m[6]  * m[11] +
              m[12] * m[7]  * m[10]

    inv[8] = m[4]  * m[9] * m[15] -
             m[4]  * m[11] * m[13] -
             m[8]  * m[5] * m[15] +
             m[8]  * m[7] * m[13] +
             m[12] * m[5] * m[11] -
             m[12] * m[7] * m[9]

    inv[12] = -m[4]  * m[9] * m[14] +
               m[4]  * m[10] * m[13] +
               m[8]  * m[5] * m[14] -
               m[8]  * m[6] * m[13] -
               m[12] * m[5] * m[10] +
               m[12] * m[6] * m[9]

    inv[1] = -m[1]  * m[10] * m[15] +
              m[1]  * m[11] * m[14] +
              m[9]  * m[2] * m[15] -
              m[9]  * m[3] * m[14] -
              m[13] * m[2] * m[11] +
              m[13] * m[3] * m[10]

    inv[5] = m[0]  * m[10] * m[15] -
             m[0]  * m[11] * m[14] -
             m[8]  * m[2] * m[15] +
             m[8]  * m[3] * m[14] +
             m[12] * m[2] * m[11] -
             m[12] * m[3] * m[10]

    inv[9] = -m[0]  * m[9] * m[15] +
              m[0]  * m[11] * m[13] +
              m[8]  * m[1] * m[15] -
              m[8]  * m[3] * m[13] -
              m[12] * m[1] * m[11] +
              m[12] * m[3] * m[9]

    inv[13] = m[0]  * m[9] * m[14] -
              m[0]  * m[10] * m[13] -
              m[8]  * m[1] * m[14] +
              m[8]  * m[2] * m[13] +
              m[12] * m[1] * m[10] -
              m[12] * m[2] * m[9]

    inv[2] = m[1]  * m[6] * m[15] -
             m[1]  * m[7] * m[14] -
             m[5]  * m[2] * m[15] +
             m[5]  * m[3] * m[14] +
             m[13] * m[2] * m[7] -
             m[13] * m[3] * m[6]

    inv[6] = -m[0]  * m[6] * m[15] +
              m[0]  * m[7] * m[14] +
              m[4]  * m[2] * m[15] -
              m[4]  * m[3] * m[14] -
              m[12] * m[2] * m[7] +
              m[12] * m[3] * m[6]

    inv[10] = m[0]  * m[5] * m[15] -
              m[0]  * m[7] * m[13] -
              m[4]  * m[1] * m[15] +
              m[4]  * m[3] * m[13] +
              m[12] * m[1] * m[7] -
              m[12] * m[3] * m[5]

    inv[14] = -m[0]  * m[5] * m[14] +
               m[0]  * m[6] * m[13] +
               m[4]  * m[1] * m[14] -
               m[4]  * m[2] * m[13] -
               m[12] * m[1] * m[6] +
               m[12] * m[2] * m[5]

    inv[3] = -m[1] * m[6] * m[11] +
              m[1] * m[7] * m[10] +
              m[5] * m[2] * m[11] -
              m[5] * m[3] * m[10] -
              m[9] * m[2] * m[7] +
              m[9] * m[3] * m[6]

    inv[7] = m[0] * m[6] * m[11] -
             m[0] * m[7] * m[10] -
             m[4] * m[2] * m[11] +
             m[4] * m[3] * m[10] +
             m[8] * m[2] * m[7] -
             m[8] * m[3] * m[6]

    inv[11] = -m[0] * m[5] * m[11] +
               m[0] * m[7] * m[9] +
               m[4] * m[1] * m[11] -
               m[4] * m[3] * m[9] -
               m[8] * m[1] * m[7] +
               m[8] * m[3] * m[5]

    inv[15] = m[0] * m[5] * m[10] -
              m[0] * m[6] * m[9] -
              m[4] * m[1] * m[10] +
              m[4] * m[2] * m[9] +
              m[8] * m[1] * m[6] -
              m[8] * m[2] * m[5]

    det := m[0] * inv[0] + m[1] * inv[4] + m[2] * inv[8] + m[3] * inv[12]

	if (math.Abs(float64(det)) < 0.00005) {
		M4Ident(result)
		return
	}

    detInv := 1.0 / det

	for i := 0; i < 16; i++ {
        result[i] = inv[i] * detInv
	}
}


func M4RotationMatrix(result *Mat3, m *Mat4) {
	result[0] = m[0]
	result[1] = m[1]
	result[2] = m[2]

	result[3] = m[4]
	result[4] = m[5]
	result[5] = m[6]

	result[6] = m[8]
	result[7] = m[9]
	result[8] = m[10]
}


func M4MulM4(result, left, right *Mat4) {
	var temp = Mat4{
		left[0]*right[0] + left[4]*right[1] + left[8]*right[2] + left[12]*right[3],
		left[1]*right[0] + left[5]*right[1] + left[9]*right[2] + left[13]*right[3],
		left[2]*right[0] + left[6]*right[1] + left[10]*right[2] + left[14]*right[3],
		left[3]*right[0] + left[7]*right[1] + left[11]*right[2] + left[15]*right[3],
		left[0]*right[4] + left[4]*right[5] + left[8]*right[6] + left[12]*right[7],
		left[1]*right[4] + left[5]*right[5] + left[9]*right[6] + left[13]*right[7],
		left[2]*right[4] + left[6]*right[5] + left[10]*right[6] + left[14]*right[7],
		left[3]*right[4] + left[7]*right[5] + left[11]*right[6] + left[15]*right[7],
		left[0]*right[8] + left[4]*right[9] + left[8]*right[10] + left[12]*right[11],
		left[1]*right[8] + left[5]*right[9] + left[9]*right[10] + left[13]*right[11],
		left[2]*right[8] + left[6]*right[9] + left[10]*right[10] + left[14]*right[11],
		left[3]*right[8] + left[7]*right[9] + left[11]*right[10] + left[15]*right[11],
		left[0]*right[12] + left[4]*right[13] + left[8]*right[14] + left[12]*right[15],
		left[1]*right[12] + left[5]*right[13] + left[9]*right[14] + left[13]*right[15],
		left[2]*right[12] + left[6]*right[13] + left[10]*right[14] + left[14]*right[15],
		left[3]*right[12] + left[7]*right[13] + left[11]*right[14] + left[15]*right[15],
	}
	M4Copy(result, &temp)
}

func M4MulV3(result *Vec3, m *Mat4, v Vec3) {
	result[0] = m[0]*v[0] + m[4]*v[1] + m[8] * v[2] + m[12]
	result[1] = m[1]*v[0] + m[5]*v[1] + m[9] * v[2] + m[13]
	result[2] = m[2]*v[0] + m[6]*v[1] + m[10] * v[2] + m[14]
}


func M4MulV4(result *Vec4, m *Mat4, v Vec4) {
	result[0] = m[0]*v[0] + m[4]*v[1] + m[8] * v[2] + m[12]
	result[1] = m[1]*v[0] + m[5]*v[1] + m[9] * v[2] + m[13]
	result[2] = m[2]*v[0] + m[6]*v[1] + m[10] * v[2] + m[14]
	result[3] = m[3]*v[0] + m[7]*v[1] + m[11] * v[2] + m[15]
}
