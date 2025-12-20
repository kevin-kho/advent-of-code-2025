package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Coord struct {
	X int
	Y int
}

type Input struct {
	FilePath string
}

func createCoords(data []byte) ([]Coord, error) {
	var res []Coord

	for entry := range bytes.SplitSeq(data, []byte{10}) {
		c := bytes.Split(entry, []byte{44})

		x, err := strconv.Atoi(string(c[0]))
		y, err := strconv.Atoi(string(c[1]))

		if err != nil {
			return res, err
		}

		res = append(res, Coord{
			X: x,
			Y: y,
		})

	}

	return res, nil
}

func getMaxRectangle(coords []Coord) int {
	var res int
	for i, c1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			c2 := coords[j]

			xSide := common.IntAbs(c2.X-c1.X) + 1
			ySide := common.IntAbs(c2.Y-c1.Y) + 1

			res = max(res, xSide*ySide)

		}

	}

	return res

}
func solvePartOne(input Input) {

	data, err := common.ReadInput(input.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	coords, err := createCoords(data)
	if err != nil {
		log.Fatal(err)
	}
	maxRectangle := getMaxRectangle(coords)
	fmt.Println(maxRectangle)

}

func main() {

	exampleInput := Input{
		FilePath: "./inputExample.txt",
	}

	solvePartOne(exampleInput)

	input := Input{
		FilePath: "./input.txt",
	}

	solvePartOne(input)

}
