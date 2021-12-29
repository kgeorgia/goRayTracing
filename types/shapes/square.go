package shapes

import (
    . "goRayTracing/types/vector"
    "math"
)
import . "goRayTracing/types/color"

type Square struct {
    Position, Rotation  Vector
    Size                float64
    Color               Color
}

func (s Square) Intersect(origin, direction Vector) (float64, bool) {
    var a, b, c, disc float64

    a = (origin.Sub(s.Position)).Dot(s.Rotation.Normalize())
    b = direction.Dot(s.Rotation.Normalize())
    if a == 0 || (a < 0 && b < 0) || (a > 0 && b > 0) {
        return 0.0, false
    }
    c = -a / b
    mult := direction.Multi(c)
    sub := (mult.Sum(origin)).Sub(s.Position)
    disc = s.Size / 2
    if math.Abs(sub.X) > disc || math.Abs(sub.Y) > disc || math.Abs(sub.Z) > disc {
        return 0.0, false
    }
    if c > 0 {
        return c, true
    }
    return 0.0, false
}

func (s Square) GetNormal(surfPoint, origin Vector) Vector {
    normRot := s.Rotation.Normalize()
    sub := surfPoint.Sub(origin)
    if normRot.Dot(sub) < 0 {
        return normRot
    }
    return normRot.Multi(-1)
}

func (s Square) GetColor() Color {
    return s.Color
}