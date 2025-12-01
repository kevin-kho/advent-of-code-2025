package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Direction int // negative = L, positive = R
	Quantity  int
}

func (c Command) getValue() int {
	return c.Direction * c.Quantity
}

func readInput(filePath string) ([]string, error) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	stringSlice := strings.SplitSeq(string(data), "\n")

	var finalSlice []string

	// remove nil string entries
	for v := range stringSlice {
		if v != "" {
			finalSlice = append(finalSlice, v)
		}
	}

	return finalSlice, nil

}

func constructCommand(cmd string) (Command, error) {
	var c Command
	direction := string(cmd[0])
	quantity := string(cmd[1:])

	var d int
	if direction == "L" {
		d = -1
	} else {
		d = 1
	}

	q, err := strconv.Atoi(quantity)
	if err != nil {
		return c, err
	}

	c = Command{
		Direction: d,
		Quantity:  q,
	}

	return c, nil

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
		if ct < 0 {
			ct = -ct
		}
		count += ct

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

	stringCommands, err := readInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var commands []Command
	for _, sc := range stringCommands {
		cmd, err := constructCommand(sc)
		if err != nil {
			log.Fatal(err)
		}
		commands = append(commands, cmd)
	}
	resLanded := solveLanded(commands)

	resPasses := solvePasses(commands)

	fmt.Println(resLanded)
	fmt.Println(resPasses)

}
