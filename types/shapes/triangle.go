package shapes

import (
    . "goRayTracing/types/vector"
    "math"
)
import . "goRayTracing/types/color"

type Triangle struct {
    A, B, C	Vector
    Color	Color
}

func (t Triangle) Intersect(origin, direction Vector) (float64, bool) {
    vectorA := t.B.Sub(t.A)
    vectorB := t.C.Sub(t.A)
    vectorC := direction.Cross(vectorB)
    disc := vectorA.Dot(vectorC)
    if math.Abs(disc) < 0.000000001 {
        return 0.0, false
    }
    c := 1 / disc
    vectorD := origin.Sub(t.A)
    a := vectorD.Dot(vectorC) * c
    if a < 0 || a > 1 {
        return 0.0, false
    }
    vectorD = vectorD.Cross(vectorA)
    b := direction.Dot(vectorD) * c
    if b < 0 || a + b > 1 {
        return 0.0, false
    }
    if ret := vectorB.Dot(vectorD) * c; ret > 0 {
        return ret, true
    }
    return 0.0, false
}

func (t Triangle) GetNormal(surfPoint, origin Vector) Vector {
    vectorAB := t.A.Sub(t.B)
    vectorBC := t.B.Sub(t.C)
    norm := (vectorAB.Cross(vectorBC)).Normalize()
    sub := (surfPoint.Sub(origin)).Normalize()
    if norm.Dot(sub) < 0 {
        return norm
    }
    return norm.Multi(-1)
}

func (t Triangle) GetColor() Color {
    return t.Color
}