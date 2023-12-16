package main

import (
	"fmt"

	"github.com/PlaiyTiziano/AOC-2023/day-10/chapters"
	"github.com/PlaiyTiziano/AOC-2023/utils"
)

func main() {
	fmt.Println("Day 10:")

	inputFilepath := utils.BuildInputFilepath(10)
	data := utils.ReadFileSplitLinesAndChars(inputFilepath)

	value := chapters.Chapter1(data)
	fmt.Println("- Chapter 1:", value)

	value = chapters.Chapter2(data)
	fmt.Println("- Chapter 2:", value)
}
