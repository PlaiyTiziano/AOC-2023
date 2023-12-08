package main

import (
	"fmt"

	"github.com/PlaiyTiziano/AOC-2023/day-5/chapters"
	"github.com/PlaiyTiziano/AOC-2023/utils"
)

func main() {
	fmt.Println("Day 5:")

	inputFilepath := utils.BuildInputFilepath(5)
	content := utils.ReadFile(inputFilepath)

	value := chapters.Chapter1(content)
	fmt.Printf("\t- Chapter 1: %d\n", value)

	value = chapters.Chapter2(content)
	fmt.Printf("\t- Chapter 2: %d\n", value)
}
