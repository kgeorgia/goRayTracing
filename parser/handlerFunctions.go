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

func parsePointLight(input []string) light.Light {
	var result light.PointLight

	if len(input) == 3 {
		result.Position.ParseVector(input[1])
		result.Color.ParseColor(input[2])
	}

	return result
}

func parseAmbientLight(input []string) light.Light {
	var result light.AmbientLight

	if len(input) == 2 {
		result.Color.ParseColor(input[1])
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

func parsePlane(input []string) shapes.Object {
	var result shapes.Plane

	if len(input) == 4 {
		result.Position.ParseVector(input[1])
		result.Rotation.ParseVector(input[2])
		result.Color.ParseColor(input[3])
	}

	return result
}

func parseSquare(input []string) shapes.Object {
	var result shapes.Square

	if len(input) == 5 {
		result.Position.ParseVector(input[1])
		result.Rotation.ParseVector(input[2])
		size, err := strconv.ParseFloat(input[3], 64)
		if err == nil {
			result.Size = size
		}
		result.Color.ParseColor(input[4])
	}

	return result
}

func parseTriangle(input []string) shapes.Object {
	var result shapes.Triangle

	if len(input) == 5 {
		result.A.ParseVector(input[1])
		result.B.ParseVector(input[2])
		result.C.ParseVector(input[3])
		result.Color.ParseColor(input[4])
	}

	return result
}

func parseCylinder(input []string) shapes.Object {
	var result shapes.Cylinder

	if len(input) == 6 {
		result.Position.ParseVector(input[1])
		result.Rotation.ParseVector(input[2])
		diameter, err := strconv.ParseFloat(input[3], 64)
		if err == nil {
			result.Diameter = diameter
		}
		height, err := strconv.ParseFloat(input[4], 64)
		if err == nil {
			result.Height = height
		}
		result.Color.ParseColor(input[5])
	}

	return result
}