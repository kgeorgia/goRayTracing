package shapes

import (
	. "goRayTracing/types/vector"
	. "goRayTracing/types/color"
)

type Object interface {
	Intersect(origin, direction Vector) float64
	GetNormal(surfPoint Vector) Vector
	GetColor() Color
}