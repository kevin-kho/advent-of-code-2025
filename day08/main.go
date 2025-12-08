package main

import (
	"aoc-2025/common"
	"bytes"
	"cmp"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
)

type pos [3]int

func (p pos) X() int {
	return p[0]
}

func (p pos) Y() int {
	return p[1]
}

func (p pos) Z() int {
	return p[2]
}

type edge struct {
	src      pos
	dst      pos
	distance float64
}

func calculateDistance(a, b pos) float64 {

	xDistance := common.IntPow(a[0]-b[0], 2)
	yDistance := common.IntPow(a[1]-b[1], 2)
	zDistance := common.IntPow(a[2]-b[2], 2)

	return math.Sqrt(float64(xDistance) + float64(yDistance) + float64(zDistance))

}

func createPositions(data []byte) ([]pos, error) {
	var res []pos

	for ln := range bytes.SplitSeq(data, []byte{10}) {

		coord := bytes.Split(ln, []byte{44})

		x, err := strconv.Atoi(string(coord[0]))
		y, err := strconv.Atoi(string(coord[1]))
		z, err := strconv.Atoi(string(coord[2]))
		if err != nil {
			return res, err
		}
		res = append(res, [3]int{x, y, z})

	}

	return res, nil
}

func createSeenMap(positions []pos) map[pos]map[pos]bool {

	res := make(map[pos]map[pos]bool)
	for _, p := range positions {
		res[p] = make(map[pos]bool)
	}

	return res

}

func createEdges(positions []pos) []edge {

	var res []edge
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			src := positions[i]
			dst := positions[j]

			distance := calculateDistance(src, dst)

			res = append(res, edge{
				src:      src,
				dst:      dst,
				distance: distance,
			})

			res = append(res, edge{
				src:      dst,
				dst:      src,
				distance: distance,
			})

		}
	}

	return res

}

func connectBoxes(n int, edges []edge, seen map[pos]map[pos]bool) {

	var connections int
	var i int

	slices.SortFunc(edges, func(a, b edge) int {
		return cmp.Compare(a.distance, b.distance)
	})

	for connections < n {
		e := edges[i]
		if !seen[e.src][e.dst] {
			seen[e.src][e.dst] = true
			seen[e.dst][e.src] = true
			connections++
		}

		i++
	}

	// fmt.Println(connections)
	// fmt.Println(seen)

	for k, v := range seen {
		fmt.Println(k, len(v))
	}

}

func main() {

	filePath := "./inputExample.txt"
	data, err := common.ReadInput(filePath)

	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	positions, err := createPositions(data)
	if err != nil {
		log.Fatal(err)
	}

	seenMap := createSeenMap(positions)
	fmt.Println(seenMap)

	edges := createEdges(positions)

	connectBoxes(10, edges, seenMap)

}
