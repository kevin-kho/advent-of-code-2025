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

	res0 := pickTwo(input0)
	res1 := pickTwo(input1)
	res2 := pickTwo(input2)
	res3 := pickTwo(input3)

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

	res0 := pickTwelve(input0)
	res1 := pickTwelve(input1)
	res2 := pickTwelve(input2)
	res3 := pickTwelve(input3)

	fmt.Println(res0)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

}
