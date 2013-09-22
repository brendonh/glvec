package glvec

type Plane struct {
	Point Vec3
	Normal Vec3
}

type Ray struct {
	Origin Vec3
	Dir Vec3
}

func (r *Ray) PlaneIntersect(p Plane) (pos Vec3, ok bool) {
	// var planePoint = Vec3{ 0.0, 0.0, 1.0 }
	// var normal = Vec3 { 0.0, 0.0, 1.0 }

	var origToPlane Vec3
	V3Sub(&origToPlane, r.Origin, p.Point)

	var dist = -(V3Dot(origToPlane, p.Normal)) / V3Dot(r.Dir, p.Normal)

	if dist <= 0 {
		return Vec3{ 0, 0, 0 }, false
	}

	V3ScalarMul(&pos, r.Dir, dist)
	V3Add(&pos, pos, r.Origin)

	return pos, true
}
