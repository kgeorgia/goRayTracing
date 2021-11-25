package main

import (
	"fmt"
	color "goRayTracing/types/color"
	vector "goRayTracing/types/vector"
)

func main() {
	col := color.Color{100, 255, 100}
	col2 := color.Color{255, 200, 180}
	col3 := col.Sum(col2)
	vector := vector.Vector{1, 2, 3}

	vector.Print()
	col.Print()
	col3.Print()
	fmt.Println(col.ToInt())
	fmt.Println("ght")
}