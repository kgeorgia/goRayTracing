package camera

import (
	"errors"
	"fmt"
	. "goRayTracing/types/vector"
	"math"
)

type Camera struct {
	Position		Vector
	Rotation		Vector
	Width, Height	uint
	Fov				uint8
}

func (c Camera) Print() {
	c.Position.Print()
	c.Rotation.Print()
	fmt.Println(c.Fov)
}

func (c Camera) CastRay(x, y uint) (pos, dir Vector, err error) {
	if x >= c.Width || y >= c.Height {
		return c.Position, c.Rotation, errors.New("seg fault")
	}

	viewportX := (float64(x) / float64(c.Width)) - 0.5
	viewportY := 0.5 - (float64(y) / float64(c.Height))
	aspect := float64(c.Width) / float64(c.Height)
	if aspect > 1 {
		viewportX *= aspect
	} else {
		viewportY /= aspect
	}

	viewportX *= math.Tan((float64(c.Fov) / 2) * math.Pi / 180)
	viewportY *= math.Tan((float64(c.Fov) / 2) * math.Pi / 180)

	result := Vector{X: viewportX, Y: viewportY, Z: 1}
	return c.Position, result, nil
}