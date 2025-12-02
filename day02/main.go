package main

import (
	"aoc-2025/common"
	"fmt"
	"log"
)

func main() {

	filePath := "./input.txt"

	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

}
