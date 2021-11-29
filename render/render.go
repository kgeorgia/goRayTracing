package render

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

	pixelArray := make([][]int, scene.Viewport.Width)
	for i := range pixelArray {
		pixelArray[i] = make([]int, scene.Viewport.Height)
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

	countCycle := int(scene.Viewport.Width * scene.Viewport.Height)
	for i := 0; i < countCycle; i++ {
		value, ok := <- colorPixel
		if !ok {
			return pixelArray
		}
		pixelArray[value.X][value.Y] = value.Value.ToInt()
	}

	return pixelArray
}

func (scene Scene) Trace(cord, color chan Pixel) {
	for value := range cord {
		var closestShape int
		var isIntersect bool
		var resultIntense Intense
		minDist := math.MaxFloat64

		position, direction := scene.Cameras[0].CastRay(value.X, value.Y, scene.Viewport)
		for idx, shape := range scene.Objects {
			currDist, ok := shape.Intersect(position, direction)
			if ok && currDist < minDist {
				isIntersect = true
				minDist = currDist
				closestShape = idx
			}
		}

		if isIntersect {
			surfPoint := direction.Multi(minDist).Sum(position)
			for _, light := range scene.Lights {
				shapeNormal, ok := light.IntersectLight(scene.Objects, closestShape, surfPoint)
				if ok {
					tmpIntense := light.AddLight(scene.Objects[closestShape].GetColor(), shapeNormal)
					resultIntense = resultIntense.Sum(tmpIntense)
				}
			}
			value.Value = resultIntense.ResultColor(scene.Objects[closestShape].GetColor())
		}
		color <- value
	}
}










