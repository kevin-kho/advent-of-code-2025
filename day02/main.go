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

func buildIntoRanges(byteArr []byte) ([]Range, error) {
	// 44 = ,
	// 45 = -

	var ranges []Range
	for b := range bytes.SplitSeq(byteArr, []byte{44}) {
		res := bytes.Split(b, []byte{45})

		s, _ := bytes.CutSuffix(res[0], []byte{10})
		e, _ := bytes.CutSuffix(res[1], []byte{10})

		start, err := strconv.Atoi(string(s))
		end, err := strconv.Atoi(string(e))
		if err != nil {
			return ranges, err
		}

		r := Range{
			Start: start,
			End:   end,
		}
		ranges = append(ranges, r)

	}

	return ranges, nil

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

	ranges, err := buildIntoRanges(data)
	if err != nil {
		log.Fatal(err)
	}

	res := totalizeInvalidIds(ranges)
	fmt.Println(res)

}
