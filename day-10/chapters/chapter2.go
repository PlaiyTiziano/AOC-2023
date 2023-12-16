package chapters

import (
	"slices"
)

type location struct {
	partOfLoop bool
	value      rune
}

func markCycle(grid [][]location, x, y int, direction string) ([][]location, [][]int) {
	var newDirection string
	var polygon [][]int

	grid[y][x].partOfLoop = true
	polygon = append(polygon, []int{x, y})

	if direction == up {
		upperVal := grid[y-1][x].value

		if !slices.Contains([]rune{'|', '7', 'F'}, upperVal) {
			return grid, polygon
		}

		if upperVal == '|' {
			newDirection = up
		} else if upperVal == '7' {
			newDirection = left
		} else {
			newDirection = right
		}

		return markCycle(grid, x, y-1, newDirection)
	}

	if direction == right {
		rightVal := grid[y][x+1].value

		if !slices.Contains([]rune{'-', '7', 'J'}, rightVal) {
			return grid, polygon
		}

		if rightVal == '-' {
			newDirection = right
		} else if rightVal == '7' {
			newDirection = down
		} else {
			newDirection = up
		}

		return markCycle(grid, x+1, y, newDirection)
	}

	if direction == down {
		lowerVal := grid[y+1][x].value

		if !slices.Contains([]rune{'|', 'J', 'L'}, lowerVal) {
			return grid, polygon
		}

		if lowerVal == '|' {
			newDirection = down
		} else if lowerVal == 'J' {
			newDirection = left
		} else {
			newDirection = right
		}

		return markCycle(grid, x, y+1, newDirection)
	}

	if direction == left {
		leftVal := grid[y][x-1].value

		if !slices.Contains([]rune{'-', 'F', 'L'}, leftVal) {
			return grid, polygon
		}

		if leftVal == '-' {
			newDirection = left
		} else if leftVal == 'F' {
			newDirection = down
		} else {
			newDirection = up
		}

		return markCycle(grid, x-1, y, newDirection)
	}

	return grid, polygon
}

func floodFill(grid [][]location, x, y int, newV rune) {
	loc := grid[y][x]

	if loc.partOfLoop || loc.value == '0' {
		return
	}

	grid[y][x].value = newV

	if y != 0 {
		floodFill(grid, x, y-1, newV)
	}

	if x != len(grid[0])-1 {
		floodFill(grid, x+1, y, newV)
	}

	if x != 0 {
		floodFill(grid, x-1, y, newV)
	}

	if y != len(grid)-1 {
		floodFill(grid, x, y+1, newV)
	}
}

func floodFillEdges(grid [][]location) {
	for y := range grid {
		floodFill(grid, 0, y, '0')
		floodFill(grid, len(grid[0])-1, y, '0')
	}

	for x := range grid[0] {
		floodFill(grid, x, 0, '0')
		floodFill(grid, x, len(grid)-1, '0')
	}
}

func isPositionEnclosed(grid [][]location, x, y int) bool {
	if grid[y][x].value == '0' || grid[y][x].partOfLoop {
		return false
	}

	var left int
	var oneSide bool
	var sideOpener rune

	for idx := x + 1; idx < len(grid[y]); idx++ {
		loc := grid[y][idx]

		if !loc.partOfLoop || loc.value == '-' {
			continue
		}

		if loc.value == 'F' || loc.value == 'L' {
			oneSide = true
			sideOpener = loc.value
			left++
		}

		if !oneSide && loc.value == '|' {
			left++
		}

		if oneSide && (loc.value == 'J' || loc.value == '7') {
			if (sideOpener == 'F' && loc.value == '7') || (sideOpener == 'L' && loc.value == 'J') {
				left++
			}
			oneSide = false
		}
	}

	return left%2 != 0
}

func Chapter2(data [][]rune) int {
	data = data[:len(data)-1]

	var grid [][]location
	for y := range data {
		for x := range data[y] {
			loc := location{value: data[y][x]}

			if x == 0 {
				grid = append(grid, []location{loc})
			} else {
				grid[y] = append(grid[y], loc)
			}
		}
	}

	y, x := findStartingPosition(data)

	grid, _ = markCycle(grid, x, y, right)

	floodFillEdges(grid)

	var tiles int
	for y := range grid {
		for x := range grid[y] {
			if isPositionEnclosed(grid, x, y) {
				tiles++
			}
		}
	}

	return tiles
}
