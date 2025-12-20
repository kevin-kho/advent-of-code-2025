package main

import (
	"bytes"
	"fmt"
	"log"
	"maps"
	"math"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Machine struct {
	DesiredBulbState    map[int]bool
	CurrentBulbState    map[int]bool
	Buttons             []Button
	DesiredJoltageState map[int]int
	CurrentJoltageState map[int]int
}

func toggleBulbs(currentState map[int]bool, b Button) map[int]bool {

	m := maps.Clone(currentState)

	for _, s := range b.Indexes {
		m[s] = !m[s]
	}

	return m

}

func incrementJoltage(currentState map[int]int, b Button) map[int]int {
	m := maps.Clone(currentState)
	for _, s := range b.Indexes {
		m[s]++
	}

	return m
}

type Button struct {
	Indexes []int // bulbs/levers to toggle
}

func getBulbDesiredState(data []byte) map[int]bool {
	desiredState := map[int]bool{}
	var i int
	for _, b := range data {
		switch b {
		case 46:
			desiredState[i] = false
			i++

		case 35:
			desiredState[i] = true
			i++
		}
	}
	return desiredState

}

func getDesiredJoltages(data []byte) (map[int]int, error) {

	res := map[int]int{}

	// Trim braces
	d := bytes.TrimPrefix(data, []byte{123})
	d = bytes.TrimSuffix(d, []byte{125})

	for i, num := range bytes.Split(d, []byte{44}) {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			return res, err
		}
		res[i] = n
	}

	return res, nil

}

func getButtons(data [][]byte) ([]Button, error) {
	var buttons []Button

	for _, s := range data {
		// Remove curly braces
		s = bytes.TrimPrefix(s, []byte{40})
		s = bytes.TrimSuffix(s, []byte{41})

		var bulbs []int
		for b := range bytes.SplitSeq(s, []byte{44}) {
			i, err := strconv.Atoi(string(b))
			if err != nil {
				return buttons, err
			}
			bulbs = append(bulbs, i)

		}

		buttons = append(buttons, Button{
			Indexes: bulbs,
		})

	}

	return buttons, nil
}

func parseData(data []byte) ([]Machine, error) {

	var machines []Machine
	byteArrs := bytes.SplitSeq(data, []byte{10})
	for b := range byteArrs {
		m, err := createMachine(b)
		if err != nil {
			return machines, err
		}
		machines = append(machines, *m)

	}

	return machines, nil

}

func createMachine(data []byte) (*Machine, error) {
	items := bytes.Split(data, []byte{32})
	desiredBulbState := getBulbDesiredState(items[0])
	currentBulbState := map[int]bool{}
	for k := range desiredBulbState {
		currentBulbState[k] = false
	}

	desiredJoltageState, err := getDesiredJoltages(items[len(items)-1])
	if err != nil {
		return nil, err
	}
	currentJoltageState := map[int]int{}
	for k := range desiredJoltageState {
		currentJoltageState[k] = 0
	}

	buttons, err := getButtons(items[1 : len(items)-1])
	if err != nil {
		return nil, err
	}

	m := &Machine{
		DesiredBulbState:    desiredBulbState,
		CurrentBulbState:    currentBulbState,
		Buttons:             buttons,
		DesiredJoltageState: desiredJoltageState,
		CurrentJoltageState: currentJoltageState,
	}

	return m, nil

}

func solveMachineBulb(m Machine) int {
	// TODO: make more efficient
	res := math.MaxInt
	var recurse func(currState map[int]bool, presses int, buttons []Button)
	recurse = func(currState map[int]bool, presses int, buttons []Button) {

		if maps.Equal(currState, m.DesiredBulbState) {
			res = min(res, presses)
			return
		}

		if presses > len(m.DesiredBulbState) {
			return
		}

		if len(buttons) == 0 {
			return
		}

		// press
		b := buttons[0]
		newButtons := append(buttons[1:], buttons[0])
		newState := toggleBulbs(currState, b)
		recurse(newState, presses+1, newButtons)

		// no press
		recurse(currState, presses, buttons[1:])

	}
	recurse(m.CurrentBulbState, 0, m.Buttons)

	return res

}

func currStateExceedsDesired(currState, desiredState map[int]int) bool {
	for k := range currState {
		if currState[k] > desiredState[k] {
			return true
		}
	}
	return false
}

func solveMachineJoltage(m Machine) int {
	res := math.MaxInt

	var recurse func(currState map[int]int, presses int, buttons []Button)
	recurse = func(currState map[int]int, presses int, buttons []Button) {
		if maps.Equal(currState, m.DesiredJoltageState) {
			res = min(res, presses)
			return
		}

		if currStateExceedsDesired(currState, m.DesiredJoltageState) {
			return
		}

		if len(buttons) == 0 {
			return
		}

		// press
		recurse(incrementJoltage(currState, buttons[0]), presses+1, buttons)

		// don't press and go to next button
		recurse(currState, presses, buttons[1:])

	}
	recurse(m.CurrentJoltageState, 0, m.Buttons)

	return res
}

func solvePartOne(filePath string) (int, error) {
	var sum int

	data, err := common.ReadInput(filePath)
	if err != nil {
		return sum, err
	}
	data = common.TrimNewLineSuffix(data)
	machines, err := parseData(data)
	if err != nil {
		return sum, err
	}
	for _, m := range machines {
		sum += solveMachineBulb(m)
	}

	return sum, nil

}

func solvePartTwo(filePath string) (int, error) {
	var sum int

	data, err := common.ReadInput(filePath)
	if err != nil {
		return sum, err
	}
	data = common.TrimNewLineSuffix(data)
	machines, err := parseData(data)
	if err != nil {
		return sum, err
	}

	for _, m := range machines {
		c := solveMachineJoltage(m)
		fmt.Println(c)
		// sum += solveMachineJoltage(m)
		sum += c
	}

	return sum, nil
}

func main() {

	resExample, err := solvePartOne("./inputExample.txt")
	if err != nil {
		log.Printf("example error: %v\n", err)
	} else {
		fmt.Println(resExample)
	}

	res, err := solvePartOne("./input.txt")
	if err != nil {
		log.Printf("solvePartOne error: %v\n", err)
	} else {
		fmt.Println(res)
	}

	resExampleTwo, err := solvePartTwo("./inputExample.txt")
	if err != nil {
		log.Printf("example error part two: %v\n", err)
	} else {
		fmt.Println(resExampleTwo)
	}

}
