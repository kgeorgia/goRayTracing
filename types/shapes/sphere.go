package shapes

import (
	"math"
	color "goRayTracing/types/color"
	vector "goRayTracing/types/vector"
)

type Sphere struct {
	Position vector.Vector
	Color color.Color
	Diameter float64
}

func (sp Sphere) Intersect(ray, pos vector.Vector) (float64) {
	var a, b, c, x1, x2, disc float64
	sub := pos.Sub(sp.Position)

	a = ray.Dot(ray)
	b = 2 * ray.Dot(sub)
	c = sub.Dot(sub) - math.Pow(sp.Diameter / 2, 2)

	disc = math.Pow(b, 2) - (4 * a * c)
	if (disc < 0) {
		return -1
	}

	disc = math.Sqrt(disc)
	a *= 2
	x1 = (-b - disc) / a
	x2 = (-b + disc) / a

	if x1 >= 0 && x1 < x2 {
		return x1
	} else if x2 >= 0 && x2 < x1 {
		return x2
	} else {
		return -1
	}
}

func (sp Sphere) GetNormal(surf_point vector.Vector)(vector.Vector) {
	sub := surf_point.Sub(sp.Position)

	return sub.Normalize()
}