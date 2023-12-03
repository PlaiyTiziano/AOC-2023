package main

import (
	"fmt"

	"github.com/PlaiyTiziano/AOC-2023/day-2/chapters"
	"github.com/PlaiyTiziano/AOC-2023/utils"
)

func main() {
	fmt.Println("Day 2:")

	inputFilepath := utils.BuildInputFilepath(2)
	games := utils.ReadFileAndSplitLines(inputFilepath)

	value := chapters.Chapter1(games)
	fmt.Printf("\t-Chapter 1: %d\n", value)

	value = chapters.Chapter2(games)
	fmt.Printf("\t-Chapter 2: %d\n", value)
}
