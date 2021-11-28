package parser

import (
	"goRayTracing/types/camera"
	"goRayTracing/types/light"
	"goRayTracing/types/shapes"
	"strconv"
)

func parseCanvas(input []string) camera.Canvas {
	var result camera.Canvas

	if len(input) != 3 {
		return result
	}

	width, err := strconv.Atoi(input[1])
	if err == nil {
		result.Width = uint(width)
	}

	height, err := strconv.Atoi(input[2])
	if err == nil {
		result.Height = uint(height)
	}

	return result
}

func parseCamera(input []string) camera.Camera {
	var result camera.Camera

	if len(input) == 4 {
		result.Position.ParseVector(input[1])
		result.Rotation.ParseVector(input[2])
		fov, err := strconv.Atoi(input[3])
		if err == nil {
			result.Fov = uint8(fov)
		}
	}
	return result
}

func parsePointLight(input []string) light.PointLight {
	var result light.PointLight

	if len(input) == 3 {
		result.Position.ParseVector(input[1])
		result.Color.ParseColor(input[2])
	}

	return result
}

func parseSphere(input []string) shapes.Object {
	var result shapes.Sphere

	if len(input) == 4 {
		result.Position.ParseVector(input[1])
		result.Color.ParseColor(input[2])
		diameter, err := strconv.ParseFloat(input[3], 64)
		if err == nil {
			result.Diameter = diameter
		}
	}

	return result
}