package parser

import (
	"bufio"
	"fmt"
	"goRayTracing/render"
	"goRayTracing/types/camera"
	"goRayTracing/types/color"
	"goRayTracing/types/light"
	"goRayTracing/types/shapes"
	"os"
	"strings"
)

var (
	parsingViewport = map[string]func([]string) camera.Canvas {
		"Resolution": parseCanvas,
	}

	parsingCameras = map[string]func([]string) camera.Camera {
		"camera": parseCamera,
	}

	parsingObjects = map[string]func([]string) shapes.Object {
		"sphere": parseSphere,
		"plane": parsePlane,
		"square": parseSquare,
		"triangle": parseTriangle,
		"cylinder": parseCylinder,
	}
	parsingLights = map[string]func([]string) light.Light {
		"pointLight": parsePointLight,
		"Ambient": parseAmbientLight,
	}
)

func Parser(filename string)(render.Scene, error) {
	var mainScene render.Scene

	lines, err := ReadFile(filename)
	if err != nil {
		fmt.Println("Error: read file!")
		return mainScene, err
	}

	for _, str := range lines {
		arrSubStr := strings.Split(str, " ")

		if tmpFunc, ok := parsingViewport[arrSubStr[0]]; ok {
			mainScene.Viewport = tmpFunc(arrSubStr)
			continue
		}

		if tmpFunc, ok := parsingCameras[arrSubStr[0]]; ok {
			mainScene.Cameras = append(mainScene.Cameras, tmpFunc(arrSubStr))
			continue
		}

		if tmpFunc, ok := parsingObjects[arrSubStr[0]]; ok {
			mainScene.Objects = append(mainScene.Objects, tmpFunc(arrSubStr))
			continue
		}

		if tmpFunc, ok := parsingLights[arrSubStr[0]]; ok {
			mainScene.Lights = append(mainScene.Lights, tmpFunc(arrSubStr))
			continue
		}
	}
	mainScene.Background = color.Color{R: 20, G: 20, B: 50}

	return mainScene, err
}

func ReadFile(filename string)([]string, error) {
	var result []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}