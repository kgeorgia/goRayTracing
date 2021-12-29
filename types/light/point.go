package light

import (
    . "goRayTracing/types/color"
    . "goRayTracing/types/vector"
    "math"
)

type PointLight struct {
    Position    Vector
    Color       Color
}

func (p PointLight) AddLight(shapeColor Color, surfNorm, reflect float64) Color {
    if surfNorm < 0 {
        surfNorm = 0
    }

    intenseR := (float64(p.Color.R) / 255) * surfNorm
    intenseG := (float64(p.Color.G) / 255) * surfNorm
    intenseB := (float64(p.Color.B) / 255) * surfNorm

    if reflect > 0 {
        intenseR += (float64(p.Color.R) / 255) * math.Pow(reflect, 50)
        intenseG += (float64(p.Color.G) / 255) * math.Pow(reflect, 50)
        intenseB += (float64(p.Color.B) / 255) * math.Pow(reflect, 50)
    }

    retR := uint8(float64(shapeColor.R) * intenseR)
    retG := uint8(float64(shapeColor.G) * intenseG)
    retB := uint8(float64(shapeColor.B) * intenseB)
    return Color{R: retR, G: retG, B: retB }
}

func (p PointLight) CreateLightRay(surfPoint Vector) (Vector, float64) {
    lightRay := p.Position.Sub(surfPoint)
    lenRay := lightRay.Length()

    return lightRay.Normalize(), lenRay
}
