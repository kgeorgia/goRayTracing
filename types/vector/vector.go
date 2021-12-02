package vector

import (
	"fmt"
	"strconv"
	"strings"
)
import "math"

type Vector struct {
	X, Y, Z float64
}

type Ray struct {
	Origin, Direction	Vector
	Length				float64
}

func (a Vector) Print() {
	fmt.Println(a.X, a.Y, a.Z)
}

func (a Vector) Sum(b Vector) Vector {
	return Vector{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

func (a Vector) Multi(b float64) Vector {
	return Vector{X: a.X * b, Y: a.Y * b, Z: a.Z * b}
}

func (a Vector) Dot(b Vector) float64 {
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Vector) Normalize() Vector {
	lenVector := a.Length()

	return Vector{X: a.X / lenVector, Y: a.Y / lenVector, Z: a.Z / lenVector}
}

func (a Vector) Cross(b Vector) Vector {
	var x, y, z float64

	x = a.Y * b.Z - a.Z * b.Y
	y = a.Z * b.X - a.X * b.Z
	z = a.X * b.Y - a.Y * b.X
	return Vector{X: x, Y: y, Z: z}
}

func (a *Vector) ParseVector(input string) {
	values := strings.Split(input, ",")

	if len(values) == 3 {
		x, err := strconv.ParseFloat(values[0], 64)
		if err == nil {
			a.X = x
		}

		y, err := strconv.ParseFloat(values[1], 64)
		if err == nil {
			a.Y = y
		}

		z, err := strconv.ParseFloat(values[2], 64)
		if err == nil {
			a.Z = z
		}
	}
}