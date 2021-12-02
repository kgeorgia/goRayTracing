package shapes

import (
	. "goRayTracing/types/color"
	. "goRayTracing/types/vector"
)

type Object interface {
	Intersect(origin, direction Vector) (float64, bool)
	GetNormal(surfPoint, origin Vector) Vector
	GetColor() Color
}