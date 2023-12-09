package main

import (
	"fmt"

	"github.com/PlaiyTiziano/AOC-2023/day-7/chapters"
	"github.com/PlaiyTiziano/AOC-2023/utils"
)

func main() {
	fmt.Println("Day 7:")

	inputFilepath := utils.BuildInputFilepath(7)
	data := utils.ReadFile(inputFilepath)

	value := chapters.Chapter1(data)
	fmt.Println("- Chapter 1:", value)

	value = chapters.Chapter2(data)
	fmt.Println("- Chapter 2:", value)
}
