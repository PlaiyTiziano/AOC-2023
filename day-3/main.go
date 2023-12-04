package main

import (
	"fmt"

	"github.com/PlaiyTiziano/AOC-2023/day-3/chapters"
	"github.com/PlaiyTiziano/AOC-2023/utils"
)

func main() {
	fmt.Println("Day 3:")

	inputFilepath := utils.BuildInputFilepath(3)
	data := utils.ReadFileSplitLinesAndChars(inputFilepath)

	value := chapters.Chapter1(data)
	fmt.Printf("\t- Chapter 1: %d\n", value)

	value = chapters.Chapter2(data)
	fmt.Printf("\t- Chapter 2: %d\n", value)
}
