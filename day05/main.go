package main

import (
	"aoc-2025/common"
	"bytes"
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type IdRange struct {
	Start int
	End   int
}

func getIngredientsRange(data []byte) ([]IdRange, []int, error) {

	var idRanges []IdRange
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
			return idRanges, ingredients, err
		}

		idRanges = append(idRanges, IdRange{
			Start: start,
			End:   end,
		})

	}

	for id := range byteArrs {
		if len(id) == 0 {
			continue
		}
		idInt, err := strconv.Atoi(string(id))
		if err != nil {
			return idRanges, ingredients, err
		}
		ingredients = append(ingredients, idInt)
	}

	return idRanges, ingredients, nil

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
	var count int

	for _, ing := range ingredients {
		if freshIngredients[ing] {
			count++
		}
	}

	return count

}

func countFreshIngredientsRange(idRanges []IdRange, ingredients []int) int {
	var count int

	// Sort both slices to ensure in-order
	slices.SortFunc(idRanges, func(a, b IdRange) int {
		return cmp.Compare(a.Start, b.Start)
	})
	slices.Sort(ingredients)

	// two pointer
	r := 0
	i := 0

	for r < len(idRanges) && i < len(ingredients) {
		idRange := idRanges[r]

		if idRange.Start <= ingredients[i] && ingredients[i] <= idRange.End {
			count++
			i++
		} else if ingredients[i] < idRange.Start {
			i++
		} else {
			r++
		}

	}

	return count

}

func main() {
	// filePath := "inputExample.txt"
	filePath := "input.txt"

	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	idRanges, ingredients, err := getIngredientsRange(data)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(idRanges)
	// fmt.Println(ingredients)

	res := countFreshIngredientsRange(idRanges, ingredients)
	fmt.Println(res)

}
