package render

import (
	. "github.com/kgeorgia/goRayTracing/types/camera"
	. "github.com/kgeorgia/goRayTracing/types/color"
	. "github.com/kgeorgia/goRayTracing/types/light"
	. "github.com/kgeorgia/goRayTracing/types/shapes"
	. "github.com/kgeorgia/goRayTracing/types/vector"
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
	Lights		[]Light
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
		minDist := math.MaxFloat64

		viewportCord := scene.Viewport.TranslateCordSys(value.X, value.Y)
		position, direction := scene.Cameras[0].CastRay(viewportCord)
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
			value.Value = scene.ApplyLight(closestShape, surfPoint, position)
		}
		color <- value
	}
}

func (scene Scene) ApplyLight(idxShape int, surfPoint, origin Vector) Color {
	var resultColor Color
	currShape := scene.Objects[idxShape]

	for idx, light := range scene.Lights {

		if idx == 1 {

		}

		var isIntersect bool
		direction, length := light.CreateLightRay(surfPoint)

		for idx, shape := range scene.Objects {
			if idx != idxShape {
				currDist, ok := shape.Intersect(surfPoint, direction)
				if ok && currDist < length {
					isIntersect = true
					break
				}
			}
		}

		if isIntersect == false {
			shapeNormal := currShape.GetNormal(surfPoint, origin)
			reflectRay := shapeNormal.Multi(shapeNormal.Dot(direction))
			reflectRay = (reflectRay.Multi(2)).Sub(direction)

			reflect := reflectRay.Dot(origin.Sub(surfPoint).Normalize())

			surfNormal := (currShape.GetNormal(surfPoint, origin)).Dot(direction)
			tmpIntense := light.AddLight(currShape.GetColor(), surfNormal, reflect)
			resultColor = resultColor.Sum(tmpIntense)
		}
	}

	return resultColor
}
