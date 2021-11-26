package camera

import (
	"fmt"
	. "goRayTracing/types/vector"
	"math"
)

type Canvas struct {
	Width, Height uint
}

type Camera struct {
	Position		Vector
	Rotation		Vector
	Fov				uint8
}

func (c Camera) Print() {
	c.Position.Print()
	c.Rotation.Print()
	fmt.Println(c.Fov)
}

func (c Camera) CastRay(x, y uint, canvas Canvas) (origin, direction Vector) {
	viewportX := (float64(x) / float64(canvas.Width)) - 0.5
	viewportY := 0.5 - (float64(y) / float64(canvas.Height))
	aspect := float64(canvas.Width) / float64(canvas.Height)
	if aspect > 1 {
		viewportX *= aspect
	} else {
		viewportY /= aspect
	}

	viewportX *= math.Tan((float64(c.Fov) / 2) * math.Pi / 180)
	viewportY *= math.Tan((float64(c.Fov) / 2) * math.Pi / 180)

	result := Vector{X: viewportX, Y: viewportY, Z: 1}
	return c.Position, result.Normalize()
}