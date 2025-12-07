package main

import (
	"aoc-2025/common"
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func getOperations(data []byte) []byte {
	// 42 == *
	// 43 == +

	byteArrs := bytes.Split(data, []byte{10})
	operations := byteArrs[len(byteArrs)-1]

	var ops []byte
	for _, op := range operations {
		if op != 32 {
			ops = append(ops, op)
		}
	}

	return ops

}

func getValues(data []byte) ([][]int, error) {

	byteArrs := bytes.Split(data, []byte{10})
	var values [][]int

	for i := 0; i < len(byteArrs)-1; i++ {
		var lineValues []int
		var curr []byte
		j := 0
		byteArr := byteArrs[i]
		for j < len(byteArr) {
			if byteArr[j] == 32 {
				if len(curr) > 0 {
					val, err := strconv.Atoi(string(curr))
					if err != nil {
						return [][]int{}, err
					}
					lineValues = append(lineValues, val)
				}

				curr = []byte{}
			} else {
				curr = append(curr, byteArr[j])
			}
			j++
		}
		if len(curr) > 0 {
			val, err := strconv.Atoi(string(curr))
			if err != nil {
				return [][]int{}, err
			}
			lineValues = append(lineValues, val)
		}

		if len(lineValues) > 0 {
			values = append(values, lineValues)
		}

	}

	return values, nil

}

func calculateColumns(values [][]int, ops []byte) ([]int, error) {
	var result []int

	if len(values) == 0 {
		return result, errors.New("empty values")
	}

	if len(ops) == 0 {
		return result, errors.New("empty ops")
	}

	if len(values[0]) != len(ops) {
		return result, errors.New("mismatching values and ops length")
	}

	for col, op := range ops {

		var val int
		if op == 42 {
			val = 1
		}

		for row := range values {
			v := values[row][col]
			switch op {
			case 42:
				val *= v
			default:
				val += v
			}
		}

		result = append(result, val)

	}

	return result, nil

}

func sumIntSlice(slc []int) int {
	var res int
	for _, v := range slc {
		res += v
	}
	return res
}

func main() {
	// filePath := "./inputExample.txt"
	filePath := "./input.txt"

	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	ops := getOperations(data)

	values, err := getValues(data)
	if err != nil {
		log.Fatal(err)
	}

	colValues, err := calculateColumns(values, ops)
	if err != nil {
		log.Fatal(err)
	}

	res := sumIntSlice(colValues)
	fmt.Println(res)

}
