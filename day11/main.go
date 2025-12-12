package main

import (
	"aoc-2025/common"
	"bytes"
	"fmt"
	"log"
	"strings"
)

type Input struct {
	FilePath string
}

func createAdjMatrix(data []byte) map[string][]string {

	adj := make(map[string][]string)
	for entry := range bytes.SplitSeq(data, []byte{10}) {
		line := strings.Split(string(entry), " ")

		node, _ := strings.CutSuffix(line[0], ":")
		neighbors := line[1:]

		adj[node] = neighbors

	}

	return adj

}

func countPaths(adj map[string][]string, startNode string, endNode string) int {

	var traverse func(node string) int
	solved := map[string]int{}
	traverse = func(node string) int {

		if node == endNode {
			return 1
		}

		if val, ok := solved[node]; ok {
			return val
		}

		var paths int
		for _, nxt := range adj[node] {
			paths += traverse(nxt)
		}

		solved[node] = paths

		return solved[node]

	}

	return traverse(startNode)

}

func countPathsWithFftDac(adj map[string][]string, startNode string) int {

	solved := map[string]int{}
	var traverse func(node string, visited map[string]bool) int
	traverse = func(node string, visited map[string]bool) int {
		if node == "out" {
			if visited["fft"] && visited["dac"] {
				return 1
			}
			return 0
		}

		// if val, ok := solved[node]; ok {
		// 	return val
		// }

		visited[node] = true

		var paths int
		for _, nxt := range adj[node] {
			paths += traverse(nxt, visited)
		}

		solved[node] = paths

		visited[node] = false

		return solved[node]

	}

	return traverse(startNode, map[string]bool{})

}

func solvePartOne(in Input, startNode string, endNode string) (int, error) {

	data, err := common.ReadInput(in.FilePath)
	if err != nil {
		return 0, err
	}

	data = common.TrimNewLineSuffix(data)

	adj := createAdjMatrix(data)
	res := countPaths(adj, startNode, endNode)

	return res, nil

}

func solvePartTwo(in Input, startNode string) (int, error) {
	data, err := common.ReadInput(in.FilePath)
	if err != nil {
		return 0, err
	}

	data = common.TrimNewLineSuffix(data)

	adj := createAdjMatrix(data)
	res := countPathsWithFftDac(adj, startNode)

	return res, nil
}

func main() {
	exampleInputPartOne := Input{FilePath: "./inputExamplePartOne.txt"}
	resExamplePartOne, err := solvePartOne(exampleInputPartOne, "you", "out")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resExamplePartOne)

	input := Input{FilePath: "./input.txt"}
	res, err := solvePartOne(input, "you", "out")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	exampleInputPartTwo := Input{FilePath: "./InputExamplePartTwo.txt"}
	resExamplePartTwo, err := solvePartTwo(exampleInputPartTwo, "svr")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resExamplePartTwo)

	// resPartTwo, err := solvePartTwo(input)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resPartTwo)

}
