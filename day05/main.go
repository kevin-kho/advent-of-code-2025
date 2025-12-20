package main

import (
	"bytes"
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type IdInterval struct {
	Start int
	End   int
}

func getIdIntervals(data []byte) ([]IdInterval, []int) {

	var idRanges []IdInterval
	var ingredients []int
	byteArrs := bytes.SplitSeq(data, []byte{10})

	for byteArr := range byteArrs {
		if len(byteArr) == 0 {
			break
		}

		idRange := bytes.Split(byteArr, []byte{45})

		var start int
		var end int

		for _, digit := range idRange[0] {
			start = start*10 + int(digit-48)
		}

		for _, digit := range idRange[1] {
			end = end*10 + int(digit-48)
		}

		idRanges = append(idRanges, IdInterval{
			Start: start,
			End:   end,
		})

	}

	for id := range byteArrs {
		if len(id) == 0 {
			continue
		}
		var idInt int
		for _, digit := range id {
			idInt = idInt*10 + int(digit-48)
		}

		ingredients = append(ingredients, idInt)

	}

	return idRanges, ingredients

}

func getIngredients(data []byte) (map[int]bool, []int, error) {
	// Too slow, the ranges are huge

	freshIngredients := make(map[int]bool)
	var ingredients []int
	d := common.TrimNewLineSuffix(data)

	byteArrs := bytes.SplitSeq(d, []byte{10})

	for byteArr := range byteArrs {
		if len(byteArr) == 0 {
			break
		}

		idRange := strings.Split(string(byteArr), "-")
		start, err := strconv.Atoi(idRange[0])
		end, err := strconv.Atoi(idRange[1])

		if err != nil {
			return freshIngredients, ingredients, err
		}
		for i := start; i < end+1; i++ {
			freshIngredients[i] = true
		}
	}

	for id := range byteArrs {
		if len(id) == 0 {
			continue
		}
		idInt, err := strconv.Atoi(string(id))
		if err != nil {
			return freshIngredients, ingredients, err
		}
		ingredients = append(ingredients, idInt)
	}

	return freshIngredients, ingredients, nil

}

func countFreshIngredients(freshIngredients map[int]bool, ingredients []int) int {
	// Not used
	var count int

	for _, ing := range ingredients {
		if freshIngredients[ing] {
			count++
		}
	}

	return count

}

func countFreshIngredientsRange(idIntervals []IdInterval, ingredients []int) int {
	var count int

	// Sort both slices to ensure in-order
	slices.SortFunc(idIntervals, func(a, b IdInterval) int {
		return cmp.Compare(a.Start, b.Start)
	})
	slices.Sort(ingredients)

	// two pointer
	i := 0 // idInterval
	j := 0 // ingredients

	for i < len(idIntervals) && j < len(ingredients) {
		interval := idIntervals[i]

		if interval.Start <= ingredients[j] && ingredients[j] <= interval.End {
			count++
			j++
		} else if ingredients[j] < interval.Start {
			j++
		} else {
			i++
		}
	}

	return count

}

func calculatePossibleFreshIngredients(idIntervals []IdInterval) int {

	// Requires non-overlapping intervals
	var count int
	for _, interval := range idIntervals {
		count += (interval.End - interval.Start + 1)
	}

	return count

}

func mergeIntervals(idIntervals []IdInterval) []IdInterval {

	slices.SortFunc(idIntervals, func(a, b IdInterval) int {
		return cmp.Compare(a.Start, b.Start)
	})

	mergedIntervals := []IdInterval{idIntervals[0]}
	i := 1
	for i < len(idIntervals) {

		prev := &mergedIntervals[len(mergedIntervals)-1]
		curr := idIntervals[i]

		if prev.Start <= curr.Start && curr.Start <= prev.End {
			prev.End = max(prev.End, curr.End)
		} else {
			mergedIntervals = append(mergedIntervals, curr)
		}

		i++
	}

	return mergedIntervals

}

func main() {
	filePath := "input.txt"

	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	idRanges, ingredients := getIdIntervals(data)

	mergedIntervals := mergeIntervals(idRanges)
	res := countFreshIngredientsRange(mergedIntervals, ingredients)
	fmt.Println(res)

	res2 := calculatePossibleFreshIngredients(mergedIntervals)
	fmt.Println(res2)

}
