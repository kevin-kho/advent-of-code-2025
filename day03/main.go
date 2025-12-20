package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func createBanks(data []byte) [][]int {

	var banks [][]int
	byteSeq := bytes.SplitSeq(data, []byte{10})
	for bs := range byteSeq {
		var bank []int
		for _, b := range bs {
			cell := int(b - 48) // Given that every value is the unicode repr of 1-9
			bank = append(bank, cell)
		}

		if len(bank) > 0 {
			banks = append(banks, bank)
		}

	}

	return banks

}

func pickTwo(bank []int) int {
	var maxPower int
	var tensPlace int // greedily keep the biggest battery as you go along

	for _, battery := range bank {
		// Check power generated with tensPlace and current battery
		power := (tensPlace*10 + battery)
		maxPower = max(maxPower, power)

		// Check if battery can be new tensPlace
		tensPlace = max(tensPlace, battery)

	}

	return maxPower
}

func pickTwelve(bank []int) int {

	// maximum index you're able to search in order to get a 12 digit number
	maxDigitIdx := make(map[int]int)
	for i := range 12 {
		maxDigitIdx[i] = len(bank) - 12 + i
	}

	var startIdx int
	digits := make([]int, 12)
	for di, _ := range digits {
		selectedDigit := -1
		selectedDigitIdx := -1

		for i := startIdx; i < maxDigitIdx[di]+1; i++ {
			if bank[i] > selectedDigit {
				selectedDigit = bank[i]
				selectedDigitIdx = i
			}
		}
		digits[di] = selectedDigit
		startIdx = selectedDigitIdx + 1

	}

	var maxPower int
	for i, digit := range digits {
		maxPower += digit * common.IntPow(10, len(digits)-i-1)
	}
	return maxPower

}

func pickTwelveRecursive(bank []int) int {
	// Too slow
	var maxPower int

	var recurse func(i int, curr int, digits int)
	recurse = func(i, curr, digits int) {
		// exit condtion: reach end of digits
		if !(i < len(bank)) {
			if digits == 0 {
				maxPower = max(maxPower, curr)
			}
			return
		}

		// exit condition: digits == 12
		if digits == 0 {
			maxPower = max(maxPower, curr)
			return
		}

		// no take
		recurse(i+1, curr, digits)

		// take
		curr += (bank[i] * common.IntPow(10, digits-1))
		recurse(i+1, curr, digits-1)

	}
	recurse(0, 0, 12)

	return maxPower

}

func findTotalPower(banks [][]int, findBankPowerStrategy func(bank []int) int) int {
	var totalPower int
	for _, bank := range banks {
		totalPower += findBankPowerStrategy(bank)
	}
	return totalPower

}

func main() {

	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	banks := createBanks(data)
	totalPower := findTotalPower(banks, pickTwo)
	fmt.Println(totalPower)

	totalPowerPartTwo := findTotalPower(banks, pickTwelve)
	fmt.Println(totalPowerPartTwo)

}
