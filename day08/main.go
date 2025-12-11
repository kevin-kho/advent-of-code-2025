package main

import (
	"aoc-2025/common"
	"bytes"
	"cmp"
	"errors"
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

}

func getCircuitSizes(seen map[pos]map[pos]bool) []int {
	visited := make(map[pos]bool)
	var res []int

	var dfs func(p pos) int
	dfs = func(p pos) int {
		if visited[p] {
			return 0
		}

		visited[p] = true
		var sum int
		for child := range seen[p] {
			sum += dfs(child)
		}

		return 1 + sum

	}

	for node := range seen {
		size := dfs(node)
		res = append(res, size)
	}

	return res

}

func productTopN(sizes []int, n int) (int, error) {
	var product int
	if n < 0 {
		return product, errors.New("n must be positive")
	}

	if n > len(sizes) {
		return product, errors.New("n is greater than length of sizes")
	}

	sizesClone := slices.Clone(sizes)
	slices.Sort(sizesClone)

	product = 1
	for i := len(sizesClone) - 1; i >= len(sizesClone)-n; i-- {
		product *= sizesClone[i]
	}

	return product, nil
}

func main() {

	// filePath := "./inputExample.txt"
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)

	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	// Part 1
	positions, err := createPositions(data)
	if err != nil {
		log.Fatal(err)
	}
	seenMap := createSeenMap(positions)
	edges := createEdges(positions)
	connectBoxes(1000, edges, seenMap)
	sizes := getCircuitSizes(seenMap)
	prod, err := productTopN(sizes, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prod)

}
