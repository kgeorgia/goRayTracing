package light

import (
	. "goRayTracing/types/color"
	. "goRayTracing/types/shapes"
	. "goRayTracing/types/vector"
)

type Intense struct {
	R, G, B float64
}

type PointLight struct {
	Position	Vector
	Color
}

//func (p PointLight) AddLight(shapeColor Color, surfNorm float64) Color {
//	if surfNorm < 0 {
//		surfNorm = 0
//	}
//
//	intenseR := (float64(p.Color.R) / 255) * surfNorm
//	intenseG := (float64(p.Color.G) / 255) * surfNorm
//	intenseB := (float64(p.Color.B) / 255) * surfNorm
//
//	retR := uint8(float64(shapeColor.R) * intenseR)
//	retG := uint8(float64(shapeColor.G) * intenseG)
//	retB := uint8(float64(shapeColor.B) * intenseB)
//	return Color{R: retR, G: retG, B: retB }
//}

func (a Intense) Sum(b Intense) Intense {
	retR := a.R + b.R
	retG := a.G + b.G
	retB := a.B + b.B

	if retR > 1.0 {
		retR = 1
	}
	if retG > 1.0 {
		retG = 1
	}
	if retB > 1.0 {
		retB = 1
	}

	return Intense{R: retR, G: retG, B: retB}
}

func (a Intense) ResultColor(b Color) Color {
	retR := uint8(float64(b.R) * a.R)
	retG := uint8(float64(b.G) * a.G)
	retB := uint8(float64(b.B) * a.B)

	return Color{R: retR, G: retG, B: retB}
}

func (p PointLight) AddLight(shapeColor Color, surfNorm float64) Intense {
	if surfNorm < 0 {
		surfNorm = 0
	}

	intenseR := (float64(p.Color.R) / 255) * surfNorm
	intenseG := (float64(p.Color.G) / 255) * surfNorm
	intenseB := (float64(p.Color.B) / 255) * surfNorm

	retR := (float64(shapeColor.R) / 255) * intenseR
	retG := (float64(shapeColor.G) / 255) * intenseG
	retB := (float64(shapeColor.B) / 255) * intenseB
	return Intense{R: retR, G: retG, B: retB }
}

func (p PointLight) IntersectLight(objects []Object, currShape int, surfPoint Vector) (float64, bool) {
	lengthRay := surfPoint.Sub(p.Position).Length()
	lightRay := surfPoint.Sub(p.Position).Normalize()

	for idx, shape := range objects {
		if idx != currShape {
			currLength, ok := shape.Intersect(p.Position, lightRay)
			if ok && currLength < lengthRay {
				return 0.0, false
			}
		}
	}
	lightRay = lightRay.Multi(-1)
	shapeNormal := objects[currShape].GetNormal(surfPoint)

	return shapeNormal.Dot(lightRay), true
}