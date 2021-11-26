package light

import (
	. "goRayTracing/types/color"
	. "goRayTracing/types/shapes"
	. "goRayTracing/types/vector"
)

type PointLight struct {
	Position	Vector
	Color
}

func (p PointLight) AddLight(shapeColor Color, surfNorm float64) Color {
	if surfNorm < 0 {
		surfNorm = 0
	}

	intenseR := (float64(p.Color.R) / 255) * surfNorm
	intenseG := (float64(p.Color.G) / 255) * surfNorm
	intenseB := (float64(p.Color.B) / 255) * surfNorm

	retR := uint8(float64(shapeColor.R) * intenseR)
	retG := uint8(float64(shapeColor.G) * intenseG)
	retB := uint8(float64(shapeColor.B) * intenseB)
	return Color{R: retR, G: retG, B: retB }
}

func (p PointLight) IntersectLight(objects []Object, currShape Object, surfPoint Vector) (float64, bool) {
	lengthRay := surfPoint.Sub(p.Position).Length()
	lightRay := surfPoint.Sub(p.Position).Normalize()

	for _, shape := range objects {
		if &shape != &currShape {
			currLength := shape.Intersect(p.Position, lightRay)
			if currLength != -1 && currLength < lengthRay {
				return 0.0, false
			}
		}
	}

	shapeNormal := currShape.GetNormal(surfPoint)

	return shapeNormal.Dot(lightRay), true
}