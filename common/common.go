package common

import (
	"math"
	"os"
)

func ReadInput(filePath string) ([]byte, error) {
	var data []byte

	data, err := os.ReadFile(filePath)
	if err != nil {
		return data, err
	}

	return data, nil

}

func IntPow(x int, pow int) int {

	return int(math.Pow(float64(x), float64(pow)))

}
