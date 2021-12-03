package shapes

import (
	. "goRayTracing/types/color"
	. "goRayTracing/types/vector"
	"math"
)

type Cylinder struct {
	Position, Rotation	Vector
	Diameter, Height	float64
	Color				Color
}

func (cy Cylinder) Intersect(origin, direction Vector) (float64, bool) {
	var a, b, c, disc, x1, x2, m1, m2 float64

	a = direction.Dot(direction) - math.Pow(direction.Dot(cy.Rotation.Normalize()), 2)
	sub := origin.Sub(cy.Position)
	b = direction.Dot(cy.Rotation.Normalize()) * sub.Dot(cy.Rotation.Normalize())
	b = 2 * (direction.Dot(sub) - b)
	c = sub.Dot(sub) - math.Pow(sub.Dot(cy.Rotation.Normalize()), 2) - math.Pow(cy.Diameter / 2, 2)
	disc = math.Pow(b, 2) - (4 * a * c)
	if disc < 0 {
		return 0.0, false
	}
	disc = math.Sqrt(disc)
	x1 = (-b - disc) / (2 * a)
	x2 = (-b + disc) / (2 * a)

	m1 = (direction.Dot(cy.Rotation.Normalize()) * x1) + sub.Dot(cy.Rotation.Normalize())
	m2 = (direction.Dot(cy.Rotation.Normalize()) * x2) + sub.Dot(cy.Rotation.Normalize())

	if x1 < x2 {
		if m1 >= 0 && m1 <= cy.Height && x1 >= 0 {
			return x1, true
		} else if m2 >= 0 && m2 <= cy.Height && x2 >= 0 {
			return x2, true
		}
	} else if x2 < x1 {
		if m2 >= 0 && m2 <= cy.Height && x2 >= 0 {
			return x2, true
		} else if m1 >= 0 && m1 <= cy.Height && x1 >= 0 {
			return x1, true
		}
	}

	return 0.0, false
}

func (cy Cylinder) GetNormal(surfPoint, origin Vector) Vector {
	lenRay := origin.Sub(surfPoint).Length()
	ray := surfPoint.Sub(origin).Normalize()
	m := (ray.Dot(cy.Rotation.Normalize()) * lenRay) + (origin.Sub(cy.Position)).Dot(cy.Rotation.Normalize())
	norm := surfPoint.Sub(cy.Position)
	norm = norm.Sub((cy.Rotation.Normalize()).Multi(m))
	norm = norm.Normalize()
	if ray.Dot(norm) < 0 {
		return norm
	}
	return norm.Multi(-1)
}

func (cy Cylinder) GetColor() Color {
	return cy.Color
}