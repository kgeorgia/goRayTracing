package shapes

import . "goRayTracing/types/vector"
import . "goRayTracing/types/color"

type Plane struct {
	Position, Rotation	Vector
	Color				Color
}

func (p Plane) Intersect(origin, direction Vector) (float64, bool) {
	var a, b, ret float64

	sub := origin.Sub(p.Position)
	a = sub.Dot(p.Rotation.Normalize())
	b = direction.Dot(p.Rotation.Normalize())
	if b == 0 || (a < 0 && b < 0) || (a > 0 && b > 0) {
		return 0.0, false
	}
	ret = -a / b
	if ret < 0 {
		return 0.0, false
	}
	return ret, true
}

func (p Plane) GetNormal(surfPoint, origin Vector) Vector {
	normRot := p.Rotation.Normalize()
	sub := surfPoint.Sub(origin)
	if normRot.Dot(sub) < 0 {
		return normRot
	}
	return normRot.Multi(-1)
}

func (p Plane) GetColor() Color {
	return p.Color
}