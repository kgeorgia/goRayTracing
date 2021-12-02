package light

import (
	. "goRayTracing/types/color"
	. "goRayTracing/types/vector"
)

type Light interface {
	AddLight(shapeColor Color, surfNorm float64) Color
	CreateLightRay(surfPoint Vector) (Vector, float64)
}