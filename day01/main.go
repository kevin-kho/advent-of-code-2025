package main

import (
	"aoc-2025/common"
	"bytes"
	"fmt"
	"log"
)

type Command struct {
	Direction int // negative = L, positive = R
	Quantity  int
}

func (c Command) getValue() int {
	return c.Direction * c.Quantity
}

func constructCommands(data []byte) []Command {
	// L = 76
	// R = 82
	var cmds []Command
	for b := range bytes.SplitSeq(data, []byte{10}) {
		var d int
		var num int
		direction := b[0]
		quantity := b[1:]

		if direction == 76 {
			d = -1
		} else {
			d = 1
		}

		for _, v := range quantity {
			num = num*10 + int(v-48)
		}

		cmd := Command{
			Direction: d,
			Quantity:  num,
		}
		cmds = append(cmds, cmd)

	}

	return cmds

}

func solveLanded(cmds []Command) int {

	// dial goes from 0-99 with wraparound
	// I think it has to be mod 100
	// count how many times the dial lands at 0
	var count int
	pos := 50
	for _, cmd := range cmds {
		v := cmd.getValue()
		pos += v
		pos %= 100

		if pos == 0 {
			count++
		}
	}

	return count

}

func solvePasses(cmds []Command) int {
	// count how many times the dial lands OR passes 0

	var count int
	pos := 50
	for _, cmd := range cmds {
		v := cmd.getValue()

		// case: turning left and it goes "negative
		if pos != 0 && pos+v < 0 {
			count++
		}

		// integer division to determine number of times it passes 0
		pos += v
		ct := pos / 100
		count += common.IntAbs(ct)

		// case: land exactly on zero
		if pos == 0 {
			count++
		}

		// Set position on the dial
		pos %= 100

		// modding a negative number by 100 doesn't change sign
		// Add 100 to make it positve
		if pos < 0 {
			pos += 100
		}

	}

	return count
}

func main() {

	data, err := common.ReadInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	commands := constructCommands(data)
	resLanded := solveLanded(commands)

	resPasses := solvePasses(commands)

	fmt.Println(resLanded)
	fmt.Println(resPasses)

}
