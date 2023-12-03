package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/PlaiyTiziano/AOC-2023/utils"
)

func TestChapter1WithSampleInput(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf(`Could not get current working directory`)
	}

	inputFilepath := fmt.Sprintf("%s/sample_input_chapter_1.txt", wd)
	lines := utils.ReadFileAndSplitLines(inputFilepath)

	value := chapter1(lines)
	expectedValue := 142

	if value != expectedValue {
		t.Fatalf(`chapter1(lines) = %d, expected %d`, value, expectedValue)
	}
}

func TestChapter1AfterCompletingIt(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf(`Could not get current working directory`)
	}

	inputFilepath := fmt.Sprintf("%s/input.txt", wd)
	lines := utils.ReadFileAndSplitLines(inputFilepath)

	value := chapter1(lines)
	expectedValue := 54630

	if value != expectedValue {
		t.Fatalf(`chapter1(lines) = %d, expected %d`, value, expectedValue)
	}
}

func TestChapter2WithSampleInput(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf(`Could not get current working directory`)
	}

	inputFilepath := fmt.Sprintf("%s/sample_input_chapter_2.txt", wd)
	lines := utils.ReadFileAndSplitLines(inputFilepath)

	value := chapter2(lines)
	expectedValue := 281

	if value != expectedValue {
		t.Fatalf(`chapter1(lines) = %d, expected %d`, value, expectedValue)
	}
}

func TestChapter2AfterCompletingIt(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf(`Could not get current working directory`)
	}

	inputFilepath := fmt.Sprintf("%s/input.txt", wd)
	lines := utils.ReadFileAndSplitLines(inputFilepath)

	value := chapter2(lines)
	expectedValue := 54770

	if value != expectedValue {
		t.Fatalf(`chapter1(lines) = %d, expected %d`, value, expectedValue)
	}
}
