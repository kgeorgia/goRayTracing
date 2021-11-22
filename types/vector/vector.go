package vector

import "fmt"
import "math"

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Print() {
	fmt.Println(v.X, v.Y, v.Z)
}

func (a Vector) Sum(b Vector)(Vector) {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vector) Sub(b Vector)(Vector) {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vector) Multi(b float64)(Vector) {
	return Vector{a.X * b, a.Y * b, a.Z * b}
}

func (a Vector) Dot(b Vector)(float64) {
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}

func (a Vector) Length()(float64) {
	return math.Sqrt(a.Dot(a))
}

func (a Vector) Normalize()(Vector) {
	len_vect := a.Length()

	return Vector{a.X / len_vect, a.Y / len_vect, a.Z / len_vect}
}

func (a Vector) Cross(b Vector)(Vector) {
	var x, y, z float64

	x = a.Y * b.Z - a.Z * b.Y
	y = a.Z * b.X - a.X * b.Z
	z = a.X * b.Y - a.Y * b.X
	return Vector{x, y, z}
}