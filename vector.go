package glvec

import (
	"fmt"
	"math"
)

type Vec3 [3]float32

func V3Copy(result, v *Vec3) {
	result[0] = v[0]
	result[1] = v[1]
	result[2] = v[2]
}	

func V3Add(result, left, right *Vec3) {
	result[0] = left[0] + right[0]
	result[1] = left[1] + right[1]
	result[2] = left[2] + right[2]
}

func V3Sub(result, left, right *Vec3) {
	result[0] = left[0] - right[0]
	result[1] = left[1] - right[1]
	result[2] = left[2] - right[2]	
}

func V3ScalarMul(result, vec *Vec3, scalar float32) {
	result[0] = vec[0] * scalar
	result[1] = vec[1] * scalar
	result[2] = vec[2] * scalar
}

func V3ScalarDiv(result, vec *Vec3, scalar float32) {
	result[0] = vec[0] / scalar
	result[1] = vec[1] / scalar
	result[2] = vec[2] / scalar
}

func V3Neg(result, vec *Vec3) {
	result[0] = -vec[0]
	result[1] = -vec[1]
	result[2] = -vec[2]
}

func V3Dot(left, right *Vec3) float32 {
	result := left[0] * right[0]
	result += left[1] * right[1]
	result += left[2] * right[2]
	return result
}

func V3LengthSquared(vec *Vec3) float32 {
	result := vec[0] * vec[0]
	result += vec[1] * vec[1]
	result += vec[2] * vec[2]
	return result
}

func V3Length(vec *Vec3) float32 {
	return float32(math.Sqrt(float64(V3LengthSquared(vec))))
}

func V3Normalize(result, vec *Vec3) {
	lenInv := 1.0 / V3Length(vec)
	result[0] = vec[0] * lenInv
	result[1] = vec[1] * lenInv
	result[2] = vec[2] * lenInv
}

func V3Cross(result, vec0, vec1 *Vec3) {
	tmpX := vec0[1] * vec1[2] - vec0[2] * vec1[1]
	tmpY := vec0[2] * vec1[0] - vec0[0] * vec1[2]
	tmpZ := vec0[0] * vec1[1] - vec0[1] * vec1[0]
	result[0] = tmpX
	result[1] = tmpY
	result[2] = tmpZ
}

func (v *Vec3) String() string {
	return fmt.Sprintf("( %f %f %f )\n", v[0], v[1], v[2])
}
