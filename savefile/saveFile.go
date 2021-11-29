package savefile

import (
	"fmt"
	"os"
)

func SaveBMP(arrPixels [][]int) {
	header := make([]byte, 54)
	sizeImage := len(arrPixels) * len(arrPixels[0]) * 3
	widthImage := len(arrPixels)
	heightImage := len(arrPixels[0])
	image := make([]byte, sizeImage)

	file, err := os.Create("output.bmp")
	if err != nil {
		fmt.Println("Error: could not a create file")
		return
	}

	header[0] = 'B'
	header[1] = 'M'
	writeIntToByte(header, 2, sizeImage)
	writeIntToByte(header, 10, 54)
	writeIntToByte(header, 14, 40)
	writeIntToByte(header, 18, widthImage)
	writeIntToByte(header, 22, heightImage)
	writeIntToByte(header, 26, 1)
	writeIntToByte(header, 28, 24)
	writeIntToByte(header, 34, sizeImage)
	writeIntToByte(header, 38, 3780)
	writeIntToByte(header, 42, 3780)

	idx := 0
	for y := heightImage - 1; y >= 0; y-- {
		for x := 0; x < widthImage; x++ {
			writeColorToByte(image, idx, arrPixels[x][y])
			idx += 3
		}
	}

	output := append(header, image...)

	if _, err := file.Write(output); err != nil {
		fmt.Println("flag1")
	}
	if err := file.Close(); err != nil {
		fmt.Println("flag3")
	}
}

func writeColorToByte(arr []byte, idx, number int) {
	arr[idx] = byte(number)
	arr[idx + 1] = byte(number >> 8)
	arr[idx + 2] = byte(number >> 16)
}

func writeIntToByte(arr []byte, idx, number int ) {
	arr[idx] = byte(number)
	arr[idx + 1] = byte(number >> 8)
	arr[idx + 2] = byte(number >> 16)
	arr[idx + 3] = byte(number >> 24)
}