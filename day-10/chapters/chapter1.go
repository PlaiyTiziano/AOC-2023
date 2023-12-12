package chapters

import "slices"

const (
	startingPositionChar = "S"
	up = "UP"
	right = "RIGHT"
	down = "DOWN"
	left = "LEFT"
)

// Looks for the starting positions and returns Y, X
func findStartingPosition(grid [][]rune) (int, int) {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'S' {
				return y, x
			}
		}
	}

	return -1, -1
}

func traverseGrid(grid [][]rune, x, y, steps int, direction string) int {
	var newDirection string

	steps++

	if direction == up {
		upperVal := grid[y-1][x]

		if !slices.Contains([]rune{'|', '7', 'F'}, upperVal) {
			return steps
		}

		if upperVal == '|' {
			newDirection = up
		} else if upperVal == '7' {
			newDirection = left
		} else {
			newDirection = right
		}

		return traverseGrid(grid, x, y-1, steps, newDirection)
	}

	if direction == right {
		rightVal := grid[y][x+1]

		if !slices.Contains([]rune{'-', '7', 'J'}, rightVal) {
			return steps
		}

		if rightVal == '-' {
			newDirection = right
		} else if rightVal == '7' {
			newDirection = down
		} else {
			newDirection = up
		}

		return traverseGrid(grid, x+1, y, steps, newDirection)
	}

	if direction == down {
		lowerVal := grid[y+1][x]

		if !slices.Contains([]rune{'|', 'J', 'L'}, lowerVal)  {
			return steps
		}

		if lowerVal == '|' {
			newDirection = down
		} else if lowerVal == 'J' {
			newDirection = left
		} else {
			newDirection = right
		}

		return traverseGrid(grid, x, y+1, steps, newDirection)
	}

	if direction == left {
		leftVal := grid[y][x-1]

		if !slices.Contains([]rune{'-', 'F', 'L'}, leftVal) {
			return steps
		}

		if leftVal == '-' {
			newDirection = left
		} else if leftVal == 'F' {
			newDirection = down
		} else {
			newDirection = up
		}

		return traverseGrid(grid, x-1, y, steps, newDirection)
	}

	return steps
}

func Chapter1(grid [][]rune) int {
	y, x := findStartingPosition(grid)

	steps := traverseGrid(grid, x, y, 0, right)

	return steps/2
}
