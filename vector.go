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

func V3Add(result *Vec3, left, right Vec3) {
	result[0] = left[0] + right[0]
	result[1] = left[1] + right[1]
	result[2] = left[2] + right[2]
}

func V3Sub(result *Vec3, left, right Vec3) {
	result[0] = left[0] - right[0]
	result[1] = left[1] - right[1]
	result[2] = left[2] - right[2]	
}

func V3ScalarMul(result *Vec3, vec Vec3, scalar float32) {
	result[0] = vec[0] * scalar
	result[1] = vec[1] * scalar
	result[2] = vec[2] * scalar
}

func V3ScalarDiv(result *Vec3, vec Vec3, scalar float32) {
	result[0] = vec[0] / scalar
	result[1] = vec[1] / scalar
	result[2] = vec[2] / scalar
}

func V3Neg(result *Vec3, vec Vec3) {
	result[0] = -vec[0]
	result[1] = -vec[1]
	result[2] = -vec[2]
}

func V3Dot(left, right Vec3) float32 {
	result := left[0] * right[0]
	result += left[1] * right[1]
	result += left[2] * right[2]
	return result
}

func V3LengthSquared(vec Vec3) float32 {
	result := vec[0] * vec[0]
	result += vec[1] * vec[1]
	result += vec[2] * vec[2]
	return result
}

func V3Length(vec Vec3) float32 {
	return float32(math.Sqrt(float64(V3LengthSquared(vec))))
}

func V3Normalize(result *Vec3, vec Vec3) {
	lenInv := 1.0 / V3Length(vec)
	result[0] = vec[0] * lenInv
	result[1] = vec[1] * lenInv
	result[2] = vec[2] * lenInv
}

func V3Cross(result *Vec3, vec0, vec1 Vec3) {
	result[0] = vec0[1] * vec1[2] - vec0[2] * vec1[1]
	result[1] = vec0[2] * vec1[0] - vec0[0] * vec1[2]
	result[2] = vec0[0] * vec1[1] - vec0[1] * vec1[0]
}

func (v *Vec3) String() string {
	return fmt.Sprintf("( %f %f %f )\n", v[0], v[1], v[2])
}


// -------------------------


type Vec4 [4]float32

func V4Copy(result *Vec4, v Vec4) {
	result[0] = v[0]
	result[1] = v[1]
	result[2] = v[2]
	result[3] = v[3]
}	

func V4CopyV3(result *Vec3, v Vec4) {
	result[0] = v[0]
	result[1] = v[1]
	result[2] = v[2]
}

func V4Add(result *Vec4, left, right Vec4) {
	result[0] = left[0] + right[0]
	result[1] = left[1] + right[1]
	result[2] = left[2] + right[2]
	result[3] = left[3] + right[3]
}

func V4Sub(result *Vec4, left, right Vec4) {
	result[0] = left[0] - right[0]
	result[1] = left[1] - right[1]
	result[2] = left[2] - right[2]	
	result[3] = left[3] - right[3]	
}

func V4ScalarDiv(result *Vec4, vec Vec4, scalar float32) {
	result[0] = vec[0] / scalar
	result[1] = vec[1] / scalar
	result[2] = vec[2] / scalar
	result[3] = vec[3] / scalar
}

func V4ScalarMul(result *Vec4, vec Vec4, scalar float32) {
	result[0] = vec[0] * scalar
	result[1] = vec[1] * scalar
	result[2] = vec[2] * scalar
	result[3] = vec[3] * scalar
}


func V4Dot(left, right Vec4) float32 {
	result := left[0] * right[0]
	result += left[1] * right[1]
	result += left[2] * right[2]
	result += left[3] * right[3]
	return result
}


func V4LengthSquared(vec Vec4) float32 {
	result := vec[0] * vec[0]
	result += vec[1] * vec[1]
	result += vec[2] * vec[2]
	result += vec[3] * vec[3]
	return result
}

func V4Length(vec Vec4) float32 {
	return float32(math.Sqrt(float64(V4LengthSquared(vec))))
}

func V4Normalize(result *Vec4, vec Vec4) {
	lenInv := 1.0 / V4Length(vec)
	result[0] = vec[0] * lenInv
	result[1] = vec[1] * lenInv
	result[2] = vec[2] * lenInv
	result[3] = vec[3] * lenInv
}
