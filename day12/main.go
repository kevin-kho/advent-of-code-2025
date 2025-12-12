package main

import (
	"aoc-2025/common"
	"fmt"
	"log"
)

func transpose(matrix [][]int) [3][3]int {
	// Transposes a 3x3 matrix 90 degrees clockwise
	var res [3][3]int

	x := 0
	for i := len(matrix) - 1; i >= 0; i-- {
		row := matrix[i]

		r := 0
		for y := range len(res[0]) {
			res[y][x] = row[r]
			r++
		}
		x++

	}

	return res

}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	res := transpose(matrix)

	fmt.Println(res)

	data, err := common.ReadInput("./inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

}
