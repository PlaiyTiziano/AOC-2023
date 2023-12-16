package chapters

import (
	"fmt"
	"os"
	"testing"

	"github.com/PlaiyTiziano/AOC-2023/utils"
)

func TestChapter2WithSampleInput3(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get current working directory")
	}

	inputFilepath := fmt.Sprintf("%s/../sample_input_3.txt", wd)
	data := utils.ReadFileSplitLinesAndChars(inputFilepath)

	value := Chapter2(data)
	expectedValue := 8

	if value != expectedValue {
		t.Fatalf("chapter2(games) = %d, expected %d", value, expectedValue)
	}
}

func TestChapter2WithSampleInput4(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get current working directory")
	}

	inputFilepath := fmt.Sprintf("%s/../sample_input_4.txt", wd)
	data := utils.ReadFileSplitLinesAndChars(inputFilepath)

	value := Chapter2(data)
	expectedValue := 10

	if value != expectedValue {
		t.Fatalf("chapter2(games) = %d, expected %d", value, expectedValue)
	}
}

func TestChapter2AfterCompletingIt(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get current working directory")
	}

	inputFilepath := fmt.Sprintf("%s/../input.txt", wd)
	data := utils.ReadFileSplitLinesAndChars(inputFilepath)

	value := Chapter2(data)
	expectedValue := 413

	if value != expectedValue {
		t.Fatalf("chapter2(games) = %d, expected %d", value, expectedValue)
	}
}
