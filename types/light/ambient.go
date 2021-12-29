package light

import . "goRayTracing/types/color"
import . "goRayTracing/types/vector"

type AmbientLight struct {
    Color   Color
}

func (a AmbientLight) AddLight(shapeColor Color, surfNorm, reflect float64) Color {
    surfNorm = 1
    reflect = 0

    intenseR := (float64(a.Color.R) / 255) * surfNorm
    intenseG := (float64(a.Color.G) / 255) * surfNorm
    intenseB := (float64(a.Color.B) / 255) * surfNorm

    retR := uint8(float64(shapeColor.R) * intenseR)
    retG := uint8(float64(shapeColor.G) * intenseG)
    retB := uint8(float64(shapeColor.B) * intenseB)
    return Color{R: retR, G: retG, B: retB }
}

func (a AmbientLight) CreateLightRay(surfPoint Vector) (Vector, float64) {
    return Vector{}, 0.0
}