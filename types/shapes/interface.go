package shapes

import (
	. "goRayTracing/types/vector"
)

type Object interface {
	Intersect(origin, direction Vector) float64
	GetNormal(surfPoint Vector) Vector
}