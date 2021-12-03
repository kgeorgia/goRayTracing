package camera

import . "goRayTracing/types/vector"

func (c Canvas) TranslateCordSys(x, y uint) Vector {
	viewportX := (float64(x) / float64(c.Width)) - 0.5
	viewportY := 0.5 - (float64(y) / float64(c.Height))
	aspect := float64(c.Width) / float64(c.Height)
	if aspect > 1 {
		viewportX *= aspect
	} else {
		viewportY /= aspect
	}

	return Vector{X: viewportX, Y: viewportY}
}
