package main

import (
    "fmt"
    "goRayTracing/parser"
    "goRayTracing/savefile"
    "os"
    "strconv"
)

func main() {
    args := os.Args

    if len(args) != 3 {
        fmt.Println("Error: bad number of arguments")
        return
    }

    mainScene, err := parser.Parser(args[1])
    if err != nil {
        fmt.Println("Error: could not read file")
        return
    }
    countThreads, _ := strconv.Atoi(args[2])
    arrPixels := mainScene.Render(countThreads)
    savefile.SaveBMP(arrPixels)
}