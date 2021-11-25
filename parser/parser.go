package parser

import (
	"bufio"
	"fmt"
	"os"

)

func Parser(filename string)([]string, error) {
	lines, err := ReadFile(filename)
	if err != nil {
		fmt.Println("Error: read file!")
		
	}
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