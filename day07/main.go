package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos [2]int

func getGrid(data []byte) [][]byte {
	return bytes.Split(data, []byte{10})
}

func getStartPos(grid [][]byte) (Pos, error) {
	var pos Pos
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 83 {
				return Pos{x, y}, nil
			}
		}
	}

	return pos, errors.New("no starting position found")

}

func countSplit(grid [][]byte, start Pos) int {
	var count int
	seen := make(map[Pos]bool)
	Y := len(grid)
	X := len(grid[0])

	var dfs func(x, y int)
	dfs = func(x, y int) {
		// case: out of bounds
		if !(0 <= x && x < X) || !(0 <= y && y < Y) {
			return
		}

		// case: already seen this pos
		if seen[Pos{x, y}] {
			return
		}

		// Visit the square and increment counter if it's a splitter
		// Traverse based on spliter or empty square
		seen[Pos{x, y}] = true
		if grid[y][x] == 94 {
			count++
			dfs(x-1, y+1)
			dfs(x+1, y+1)
		} else {
			dfs(x, y+1)
		}

	}
	dfs(start[0], start[1])

	return count
}

func countTimelines(grid [][]byte, start Pos) int {

	Y := len(grid)
	X := len(grid[0])

	cache := make(map[Pos]int)
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		// case: out of bounds
		if !(0 <= x && x < X) || !(0 <= y && y < Y) {
			return 0
		}

		// case: we already solved
		if val, ok := cache[Pos{x, y}]; ok {
			return val
		}

		// case: we reach the end
		if y == Y-1 {
			return 1
		}

		// Traverse based on splitter or empty square
		// Memoize the result
		if grid[y][x] == 94 {
			cache[Pos{x, y}] = dfs(x-1, y+1) + dfs(x+1, y+1)
			return cache[Pos{x, y}]
		}

		cache[Pos{x, y}] = dfs(x, y+1)
		return cache[Pos{x, y}]

	}

	return dfs(start[0], start[1])

}

func main() {
	// filePath := "./inputExample.txt"
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	grid := getGrid(data)
	startPos, err := getStartPos(grid)
	if err != nil {
		log.Fatal(err)
	}

	res := countSplit(grid, startPos)
	fmt.Println(res)

	res2 := countTimelines(grid, startPos)
	fmt.Println(res2)

}
