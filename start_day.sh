#! /bin/sh
ROOT="$HOME/personal/AOC-2023/day-$1"

mkdir -p "$ROOT/chapters"

touch "$ROOT/main.go"

echo "package main

import (
	\"fmt\"

	\"github.com/PlaiyTiziano/AOC-2023/day-$1/chapters\"
	\"github.com/PlaiyTiziano/AOC-2023/utils\"
)

func main() {
	fmt.Println(\"Day $1:\")

	// inputFilepath := utils.BuildInputFilepath($1)
	// content := utils.ReadFile(inputFilepath)

	value := chapters.Chapter1()
	fmt.Println(\"- Chapter 1:\", value)

	value = chapters.Chapter2()
	fmt.Println(\"- Chapter 2:\", value)
}" >> "$ROOT/main.go"

touch "$ROOT/chapters/chapter1.go"

echo "package chapters

func Chapter1() int {
	return -1
}" >> "$ROOT/chapters/chapter1.go"

touch "$ROOT/chapters/chapter1_test.go"

echo "package chapters

import (
	\"fmt\"
	\"os\"
	\"testing\"

	\"github.com/PlaiyTiziano/AOC-2023/utils\"
)

func TestChapter1WithSampleInput(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf(\"Could not get current working directory\")
	}

	// inputFilepath := fmt.Sprintf(\"%s/../sample_input_chapter_1.txt\", wd)
	// data := utils.ReadFile(inputFilepath)

	value := Chapter1()
	expectedValue := 35

	if value != expectedValue {
		t.Fatalf(\"chapter1(games) = %d, expected %d\", value, expectedValue)
	}
}

// func TestChapter1AfterCompletingIt(t *testing.T) {
	// wd, err := os.Getwd()
	// if err != nil {
		// t.Fatalf(\"Could not get current working directory\")
	// }
	// 
	// inputFilepath := fmt.Sprintf(\"%s/../input.txt\", wd)
	// data := utils.ReadFile(inputFilepath)
	//
	// value := Chapter1(data)
	// expectedValue := 322500873
	//
	// if value != expectedValue {
		// t.Fatalf(\"chapter1(games) = %d, expected %d\", value, expectedValue)
	// }
// }" >> "$ROOT/chapters/chapter1_test.go"

touch "$ROOT/chapters/chapter2.go"

echo "package chapters

func Chapter2() int {
	return -1
}" >> "$ROOT/chapters/chapter2.go"

touch "$ROOT/chapters/chapter2_test.go"

echo "package chapters

import (
	\"fmt\"
	\"os\"
	\"testing\"

	\"github.com/PlaiyTiziano/AOC-2023/utils\"
)

func TestChapter2WithSampleInput(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf(\"Could not get current working directory\")
	}

	// inputFilepath := fmt.Sprintf(\"%s/../sample_input_chapter_1.txt\", wd)
	// data := utils.ReadFile(inputFilepath)

	value := Chapter2()
	expectedValue := 35

	if value != expectedValue {
		t.Fatalf(\"chapter2(games) = %d, expected %d\", value, expectedValue)
	}
}

// func TestChapter2AfterCompletingIt(t *testing.T) {
	// wd, err := os.Getwd()
	// if err != nil {
		// t.Fatalf(\"Could not get current working directory\")
	// }
	// }
	// 
	// inputFilepath := fmt.Sprintf(\"%s/../input.txt\", wd)
	// data := utils.ReadFile(inputFilepath)
	//
	// value := Chapter2(data)
	// expectedValue := 322500873
	//
	// if value != expectedValue {
		// t.Fatalf(\"chapter2(games) = %d, expected %d\", value, expectedValue)
	// }
// }" >> "$ROOT/chapters/chapter2_test.go"

# Fetch the input
curl -H "Cookie: session=`cat $HOME/personal/AOC-2023/cookie.txt`" \
	"https://adventofcode.com/2023/day/$1/input" \
	-o "$ROOT/input.txt"
