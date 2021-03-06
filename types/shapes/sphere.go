package shapes

import (
    . "goRayTracing/types/color"
    . "goRayTracing/types/vector"
    "math"
)

type Sphere struct {
    Position    Vector
    Color       Color
    Diameter    float64
}

func (sp Sphere) Intersect(origin, direction Vector) (float64, bool) {
    var a, b, c, x1, x2, disc float64
    sub := origin.Sub(sp.Position)

    a = direction.Dot(direction)
    b = 2 * direction.Dot(sub)
    c = sub.Dot(sub) - math.Pow(sp.Diameter / 2, 2)

    disc = math.Pow(b, 2) - (4 * a * c)
    if disc < 0 {
        return 0.0, false
    }

    disc = math.Sqrt(disc)
    a *= 2
    x1 = (-b - disc) / a
    x2 = (-b + disc) / a

    if x1 >= 0 && x1 < x2 {
        return x1, true
    } else if x2 >= 0 && x2 < x1 {
        return x2, true
    } else {
        return 0.0, false
    }
}

func (sp Sphere) GetNormal(surfPoint, _ Vector) Vector {
    sub := surfPoint.Sub(sp.Position)

    return sub.Normalize()
}

func (sp Sphere) GetColor() Color {
    return sp.Color
}