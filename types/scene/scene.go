package scene

import (
	. "goRayTracing/types/camera"
	. "goRayTracing/types/color"
	. "goRayTracing/types/light"
	. "goRayTracing/types/shapes"
	"math"
)

type Pixel struct {
	X, Y	uint
	Value	Color
}

type Scene struct {
	Viewport	Canvas
	Cameras		[]Camera
	Objects		[]Object
	Lights		[]PointLight
	Background	Color
}

func (scene Scene) Render(countThreads int) [][]int {

	pixelArray := make([][]int, scene.Viewport.Height)
	for i := range pixelArray {
		pixelArray[i] = make([]int, scene.Viewport.Width)
	}

	cordPixel := make(chan Pixel, scene.Viewport.Width * scene.Viewport.Height)
	colorPixel := make(chan Pixel)

	for countGo := 0; countGo < countThreads; countGo++ {
		go scene.Trace(cordPixel, colorPixel)
	}

	for y := uint(0); y < scene.Viewport.Height; y++ {
		for x := uint(0); x < scene.Viewport.Width; x++ {
			cordPixel <- Pixel{X: x, Y: y, Value: scene.Background}
		}
	}
	close(cordPixel)

	for value := range colorPixel {
			pixelArray[value.X][value.Y] = value.Value.ToInt()
	}

	return pixelArray
}

func (scene Scene) Trace(cord, color chan Pixel) {
	var closerShape Object
	dist := math.MaxFloat64

	for value := range cord {
		position, direrction := scene.Cameras[0].CastRay(value.X, value.Y, scene.Viewport)

		for _, shape := range scene.Objects {
			currDist := shape.Intersect(dir, pos)

			if currDist != -1 && currDist < dist {
				dist = currDist
				closerShape = shape
			}
		}


	}
}










