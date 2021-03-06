package light

import (
    . "goRayTracing/types/color"
    . "goRayTracing/types/vector"
)

type Light interface {
    AddLight(shapeColor Color, surfNorm, reflect float64) Color
    CreateLightRay(surfPoint Vector) (Vector, float64)
}