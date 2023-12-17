package chapters

import (
	"slices"
)

const (
	xSize = 200
	ySize = 200
)

type QueueItem struct {
	x        int
	y        int
	distance int
}

// Also loop over all chars to map the galaxies
// Store the location of the galaxies in a 2D int array [[x,y],[x,y]]
// Only move NESW, can move through other galaxies
// Store a multiplier in a map for each x and y index, so every y gets a multiplier and every x as well.

func findDistanceOfshortestPath(a, b []int) int {
	var visited [ySize][xSize]bool
	source := QueueItem{x: a[0], y: a[1]}
	queue := []QueueItem{source}

	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]

		if item.y == b[1] && item.x == b[0] {
			return item.distance
		}

		// Up
		if item.y > 0 && !visited[item.y-1][item.x] {
			queue = append(queue, QueueItem{
				y:        item.y - 1,
				x:        item.x,
				distance: item.distance + 1,
			})
			visited[item.y-1][item.x] = true
		}

		// Right
		if item.x < xSize-1 && !visited[item.y][item.x+1] {
			queue = append(queue, QueueItem{
				y:        item.y,
				x:        item.x + 1,
				distance: item.distance + 1,
			})
			visited[item.y][item.x+1] = true
		}

		// Down
		if item.y < ySize-1 && !visited[item.y+1][item.x] {
			queue = append(queue, QueueItem{
				y:        item.y + 1,
				x:        item.x,
				distance: item.distance + 1,
			})
			visited[item.y+1][item.x] = true
		}

		// Left
		if item.x > 0 && !visited[item.y][item.x-1] {
			queue = append(queue, QueueItem{
				y:        item.y,
				x:        item.x - 1,
				distance: item.distance + 1,
			})
			visited[item.y][item.x-1] = true
		}
	}

	return -1
}

func expandGrid(grid [][]rune, galaxies [][]int) [][]rune {
	var allY []int
	var allX []int
	for _, galaxy := range galaxies {
		allX = append(allX, galaxy[0])
		allY = append(allY, galaxy[1])
	}

	for y := 0; y < len(grid); y++ {
		if slices.Contains(allY, y) {
			continue
		}

		grid = append(grid[:y+1], grid[y:]...)
		grid[y] = []rune{}
		for range grid[0] {
			grid[y] = append(grid[y], '.')
		}
		y++
		for idx := range allY {
			allY[idx]++
		}
	}

	for x := 0; x < len(grid[0]); x++ {
		if slices.Contains(allX, x) {
			continue
		}

		for y := range grid {
			grid[y] = append(grid[y][:x+1], grid[y][x:]...)
			grid[y][x] = '.'
		}
		x++
		for idx := range allX {
			allX[idx]++
		}
	}

	return grid
}

func locateGalaxies(grid [][]rune) [][]int {
	var galaxies [][]int

	for y := range grid {
		for x, char := range grid[y] {
			if char == '#' {
				galaxies = append(galaxies, []int{x, y})
			}
		}
	}

	return galaxies
}

func Chapter1(grid [][]rune) int {
	grid = grid[:len(grid)-1]

	galaxies := locateGalaxies(grid)

	grid = expandGrid(grid, galaxies)

	galaxies = locateGalaxies(grid)

	var value int
	for len(galaxies) != 0 {
		galaxyA := galaxies[0]
		galaxies = galaxies[1:]

		for _, galaxyB := range galaxies {
			value += findDistanceOfshortestPath(galaxyA, galaxyB)
		}
	}

	return value
}
