package main

import (
	"aoc-2025/common"
	"bytes"
	"fmt"
	"log"
)

func buildGrid(data []byte) [][]byte {
	trimWhiteSpace, _ := bytes.CutSuffix(data, []byte{10})
	return bytes.Split(trimWhiteSpace, []byte{10})
}

func isValidRoll(x, y int, grid [][]byte) bool {

	X := len(grid[0])
	Y := len(grid)
	deltas := [][2]int{
		{-1, 1},
		{0, 1},
		{1, 1},
		{-1, 0},
		{1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	var surrounds int
	for _, delta := range deltas {
		newX := x + delta[0]
		newY := y + delta[1]

		// case: out of bounds
		if !(0 <= newX && newX < X) || !(0 <= newY && newY < Y) {
			continue
		}
		if grid[newY][newX] == 64 {
			surrounds++
		}

	}

	return surrounds < 4

}

func countValidRolls(grid [][]byte) int {
	// @ = 64
	var count int
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 64 && isValidRoll(x, y, grid) {
				count++
			}

		}
	}

	return count
}

func countValidRollsPartTwo(grid [][]byte) int {
	// @ = 64

	var count int

	var removed int = 1
	for removed > 0 {
		removed = 0
		for y, row := range grid {
			for x := range row {
				if grid[y][x] == 64 && isValidRoll(x, y, grid) {
					count++
					removed++
					grid[y][x] = 46
				}
			}
		}
	}

	return count

}

func main() {

	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	grid := buildGrid(data)

	ct := countValidRolls(grid)
	fmt.Println(ct)

	ct2 := countValidRollsPartTwo(grid)
	fmt.Println(ct2)

}
