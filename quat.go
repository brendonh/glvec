package glvec

import "math"

type Quat struct {
	W float32
	V Vec3
}

func QIdent(result *Quat) {
	result.W = 1.0
	result.V[0] = 0.0
	result.V[1] = 0.0
	result.V[2] = 0.0
}

func QRotAng(result *Quat, angle float32, axis *Vec3) {
    sin_a := float32(math.Sin( float64(angle) / 2 ))
    cos_a := float32(math.Cos( float64(angle) / 2 ))

	V3ScalarMul(&result.V, axis, sin_a)
	result.W = cos_a

	QNorm(result, result)
}

func QConj(result, q *Quat) {
	V3Neg(&result.V, &q.V)
	result.W =  q.W
}

func QMag(q *Quat) float32 {
    return float32(math.Sqrt(float64(
		q.W * q.W +
		q.V[0] * q.V[0] + 
		q.V[1] * q.V[1] +
		q.V[2] * q.V[2])))
}

func QNorm(result, q *Quat) {
	magInv := 1.0 / QMag(q)
	V3ScalarMul(&result.V, &q.V, magInv)
	result.W = q.W * magInv
}


func QMul(result, left, right *Quat) {
	var temp Quat
	temp.W = V3Dot(&left.V, &right.V)

	var va, vb, vc Vec3
	
	V3Cross(&va, &left.V, &right.V)
	V3ScalarMul(&vb, &left.V, right.W)
	V3ScalarMul(&vc, &right.V, left.W)

	V3Add(&va, &va, &vb)
	V3Add(&temp.V, &va, &vc)

	QNorm(result, &temp)
}

func QMat3(result *Mat3, q *Quat) {
	xx := q.V[0] * q.V[0]
    xy := q.V[0] * q.V[1]
    xz := q.V[0] * q.V[2]
    xw := q.V[0] * q.W

    yy := q.V[1] * q.V[1]
    yz := q.V[1] * q.V[2]
    yw := q.V[1] * q.W

    zz := q.V[2] * q.V[2]
    zw := q.V[2] * q.W

    result[0] = 1 - 2 * ( yy + zz )
    result[1] =     2 * ( xy - zw )
    result[2] =     2 * ( xz + yw )

    result[3] =     2 * ( xy + zw )
    result[4] = 1 - 2 * ( xx + zz )
    result[5] =     2 * ( yz - xw )

    result[6] =     2 * ( xz - yw )
    result[7] =     2 * ( yz + xw )
    result[8] = 1 - 2 * ( xx + yy )
}

func QMat4(result *Mat4, q *Quat) {
	xx := q.V[0] * q.V[0]
    xy := q.V[0] * q.V[1]
    xz := q.V[0] * q.V[2]
    xw := q.V[0] * q.W

    yy := q.V[1] * q.V[1]
    yz := q.V[1] * q.V[2]
    yw := q.V[1] * q.W

    zz := q.V[2] * q.V[2]
    zw := q.V[2] * q.W

    result[0]  = 1 - 2 * ( yy + zz )
    result[1]  =     2 * ( xy - zw )
    result[2]  =     2 * ( xz + yw )

    result[4]  =     2 * ( xy + zw )
    result[5]  = 1 - 2 * ( xx + zz )
    result[6]  =     2 * ( yz - xw )

    result[8]  =     2 * ( xz - yw )
    result[9]  =     2 * ( yz + xw )
    result[10] = 1 - 2 * ( xx + yy )

    result[3] = 0
	result[7] = 0
	result[11] = 0
	result[12] = 0
	result[13] = 0
	result[14] = 0
    result[15] = 1
}
