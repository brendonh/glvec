package glvec

import "math"

type Mat3 [9]float32

func M3Ident(result *Mat3) {
	result[0] = 1.0
	result[1] = 0.0
	result[2] = 0.0

	result[3] = 0.0
	result[4] = 1.0
	result[5] = 0.0

	result[6] = 0.0
	result[7] = 0.0
	result[8] = 1.0
}

func M3Copy(result, m *Mat3) {
	for i := 0; i < 9; i++ {
		result[i] = m[i]
	}
}

func M3Transpose(result, m *Mat3) {
	result[0] = m[0]
	result[1] = m[3]
	result[2] = m[6]

	result[3] = m[1]
	result[4] = m[4]
	result[5] = m[7]

	result[6] = m[2]
	result[7] = m[5]
	result[8] = m[8]
}


func M3Det(m *Mat3) float32 {
    return m[0] * ( m[4]*m[8] - m[7]*m[5] ) -
		m[1] * ( m[3]*m[8] - m[6]*m[5] ) +
		m[2] * ( m[3]*m[7] - m[6]*m[4] )
}


func M3Inv(result, m *Mat3) {
	det := M3Det(m)
	if (math.Abs(float64(det)) < 0.0005) {
		M3Ident(result)
		return
	}

	detInv := 1 / det

	result[0] =    m[4]*m[8] - m[5]*m[7]   * detInv
	result[1] = -( m[1]*m[8] - m[7]*m[2] ) * detInv
	result[2] =    m[1]*m[5] - m[4]*m[2]   * detInv

	result[3] = -( m[3]*m[8] - m[5]*m[6] ) * detInv
	result[4] =    m[0]*m[8] - m[6]*m[2]   * detInv
	result[5] = -( m[0]*m[5] - m[3]*m[2] ) * detInv

	result[6] =    m[3]*m[7] - m[6]*m[4]   * detInv
	result[7] = -( m[0]*m[7] - m[6]*m[1] ) * detInv
	result[8] =    m[0]*m[4] - m[1]*m[3]   * detInv
}


func M3MulM3(result, left, right *Mat3) {
	var temp = Mat3 {
		left[0]*right[0] + left[3]*right[1] + left[6]*right[2],
		left[1]*right[0] + left[4]*right[1] + left[7]*right[2],
		left[2]*right[0] + left[5]*right[1] + left[8]*right[2],
		left[0]*right[3] + left[3]*right[4] + left[6]*right[5],
		left[1]*right[3] + left[4]*right[4] + left[7]*right[5],
		left[2]*right[3] + left[5]*right[4] + left[8]*right[5],
		left[0]*right[6] + left[3]*right[7] + left[6]*right[8],
		left[1]*right[6] + left[4]*right[7] + left[7]*right[8],
		left[2]*right[6] + left[5]*right[7] + left[8]*right[8],
	}
	M3Copy(result, &temp)
}

func M3MulV3(result *Vec3, m *Mat3, v *Vec3) {
	var temp Vec3
	temp[0] = m[0]*v[0] + m[3]*v[1] + m[6]*v[2]
	temp[1] = m[1]*v[0] + m[4]*v[1] + m[7]*v[2]
	temp[2] = m[2]*v[0] + m[5]*v[1] + m[8]*v[2]
	V3Copy(result, &temp)
}
