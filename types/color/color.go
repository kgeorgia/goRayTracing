package color

import "fmt"

type Color struct {
	R, G, B uint8
}

func (x Color) Print() {
	fmt.Println(x.R, x.G, x.B)
}

func (x Color) Sum(y Color) Color {
	var r, g, b uint8

	if int(x.R) + int(y.R) > 255 {
		r = 255
	} else {
		r = x.R + y.R
	}

	if int(x.G) + int(y.G) > 255 {
		g = 255
	} else {
		g = x.G + y.G
	}

	if int(x.B) + int(y.B) > 255 {
		b = 255
	} else {
		b = x.B + y.B
	}
	return Color{r, g, b}
}

func (x Color) ToInt() int {
	return int(x.R) << 16 | int(x.G) << 8 | int(x.B)
}