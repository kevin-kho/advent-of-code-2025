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
	if len(slc) == 0 {
		return res
	}
	for _, v := range slc {
		res += v
	}
	return res
}

func productIntSlice(slc []int) int {
	if len(slc) == 0 {
		return 0
	}
	res := 1
	for _, v := range slc {
		res *= v
	}

	return res
}

func calculateReverse(data []byte) int {
	// 42 == *
	// 43 == +
	// whitespace == 32

	var longestRow int
	rows := bytes.Split(data, []byte{10})
	for _, row := range rows {
		longestRow = max(longestRow, len(row))
	}

	// Pad the rows to have the same length
	var rowsNew [][]byte
	for i := range rows {
		r := make([]byte, longestRow)
		copy(r, rows[i])
		for j := len(rows[i]); j < longestRow; j++ {
			r[j] = byte(32)
		}
		rowsNew = append(rowsNew, r)

	}

	var colNums []int
	var nums []int
	for x := longestRow - 1; x >= 0; x-- {
		var num int
		for y := range rowsNew {
			// Handle sum or product op
			if y == len(rowsNew)-1 && num != 0 {
				nums = append(nums, num)
				switch rowsNew[y][x] {
				case 42:
					colNums = append(colNums, productIntSlice(nums))
					num = 0
					nums = []int{}

				case 43:
					colNums = append(colNums, sumIntSlice(nums))
					num = 0
					nums = []int{}
				}

			} else {
				// Handle calculating number
				val := rowsNew[y][x]
				if val != 32 {
					num = num*10 + int(val-48)
				}
			}
		}
	}

	return sumIntSlice(colNums)

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

	res2 := calculateReverse(data)
	fmt.Println(res2)

}
