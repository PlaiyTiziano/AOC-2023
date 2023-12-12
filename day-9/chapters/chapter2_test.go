package chapters

import (
	"fmt"
	"os"
	"testing"

	"github.com/PlaiyTiziano/AOC-2023/utils"
)

func TestChapter2WithSampleInput(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get current working directory")
	}

	inputFilepath := fmt.Sprintf("%s/../sample_input_1.txt", wd)
	data := utils.ReadFile(inputFilepath)

	value := Chapter2(data)
	expectedValue := 2

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
	data := utils.ReadFile(inputFilepath)

	value := Chapter2(data)
	expectedValue := 1100

	if value != expectedValue {
		t.Fatalf("chapter2(games) = %d, expected %d", value, expectedValue)
	}
}
