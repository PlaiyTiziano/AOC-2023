package chapters

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

	inputFilepath := fmt.Sprintf("%s/../sample_input_chapter_1.txt", wd)
	games := utils.ReadFileAndSplitLines(inputFilepath)

	value := Chapter1(games)
	expectedValue := 8

	if value != expectedValue {
		t.Fatalf("chapter1(games) = %d, expected %d", value, expectedValue)
	}
}

func TestChapter1AfterCompletingIt(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf(`Could not get current working directory`)
	}

	inputFilepath := fmt.Sprintf("%s/../input.txt", wd)
	games := utils.ReadFileAndSplitLines(inputFilepath)

	value := Chapter1(games)
	expectedValue := 2156

	if value != expectedValue {
		t.Fatalf("chapter1(games) = %d, expected %d", value, expectedValue)
	}
}
