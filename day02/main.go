package main

import (
	"aoc-2025/common"
	"bytes"
	"fmt"
	"log"
	"strconv"
)

type Range struct {
	Start int
	End   int
}

func (r *Range) findInvalidIds() []int {
	var invalidIds []int
	for i := r.Start; i < r.End+1; i++ {
		intStr := strconv.Itoa(i)

		left := intStr[:len(intStr)/2]
		right := intStr[len(intStr)/2:]
		if left == right {
			invalidIds = append(invalidIds, i)
		}

	}

	return invalidIds

}

func (r *Range) findInvalidIdsPart2() []int {
	// TODO: Finish eventually
	// An invalid Id is now
	// if it's of some sequence of digits repeated at least twice
	// Meaning the pattern can be of size 1 to len(n)/2 where n is the id
	var invalidIds []int
	for i := r.Start; i < r.End+1; i++ {
		intStr := strconv.Itoa(i)
		for j := 0; j < len(intStr)/2+1; j++ {

		}

	}

	return invalidIds

}

func buildIntoRanges(byteArr []byte) []Range {
	// 44 = ,
	// 45 = -

	var ranges []Range
	for b := range bytes.SplitSeq(byteArr, []byte{44}) {
		res := bytes.Split(b, []byte{45})

		var start int
		var end int
		for _, digit := range res[0] {
			start = start*10 + int(digit-48)
		}

		for _, digit := range res[1] {
			end = end*10 + int(digit-48)
		}

		r := Range{
			Start: start,
			End:   end,
		}
		ranges = append(ranges, r)

	}

	return ranges

}

func totalizeInvalidIds(ranges []Range) int {
	var total int
	for _, r := range ranges {
		for _, id := range r.findInvalidIds() {
			total += id
		}

	}

	return total

}

func main() {

	filePath := "./input.txt"

	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	ranges := buildIntoRanges(data)

	res := totalizeInvalidIds(ranges)
	fmt.Println(res)

}
