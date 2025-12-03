package main

import (
	"fmt"
	"slices"
	"testing"
)

func createIntSlice(input int) []int {

	var res []int
	for input != 0 {
		digit := input % 10
		input = input / 10
		res = append(res, digit)
	}

	slices.Reverse(res)

	return res

}

func TestFindMaxPower(t *testing.T) {

	input0 := createIntSlice(987654321111111)
	input1 := createIntSlice(811111111111119)
	input2 := createIntSlice(234234234234278)
	input3 := createIntSlice(818181911112111)

	res0 := findMaxPower(input0)
	res1 := findMaxPower(input1)
	res2 := findMaxPower(input2)
	res3 := findMaxPower(input3)

	fmt.Println(res0)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

}

func TestFindMaxPowPartTwoRecursive(t *testing.T) {
	input0 := createIntSlice(987654321111111)
	input1 := createIntSlice(811111111111119)
	input2 := createIntSlice(234234234234278)
	input3 := createIntSlice(818181911112111)

	res0 := findMaxPowerPartTwo(input0)
	res1 := findMaxPowerPartTwo(input1)
	res2 := findMaxPowerPartTwo(input2)
	res3 := findMaxPowerPartTwo(input3)

	fmt.Println(res0)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

}
