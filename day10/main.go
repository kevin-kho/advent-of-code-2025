package main

import (
	"aoc-2025/common"
	"bytes"
	"fmt"
	"log"
	"maps"
	"math"
	"strconv"
)

type Machine struct {
	DesiredState map[int]bool
	CurrentState map[int]bool
	Buttons      []Button
	Joltage      []int
}

func toggle(currentState map[int]bool, b Button) map[int]bool {

	m := maps.Clone(currentState)

	for _, s := range b.Bulbs {
		m[s] = !m[s]
	}

	return m

}

type Button struct {
	Bulbs []int // bulbs to toggle
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

func getJoltages(data []byte) ([]int, error) {

	var res []int

	// Trim braces
	d := bytes.TrimPrefix(data, []byte{123})
	d = bytes.TrimSuffix(d, []byte{125})

	for num := range bytes.SplitSeq(d, []byte{44}) {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			return res, err
		}
		res = append(res, n)
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
			Bulbs: bulbs,
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
	desiredState := getBulbDesiredState(items[0])
	currentState := map[int]bool{}
	for k := range desiredState {
		currentState[k] = false
	}

	joltages, err := getJoltages(items[len(items)-1])
	if err != nil {
		return nil, err
	}

	buttons, err := getButtons(items[1 : len(items)-1])
	if err != nil {
		return nil, err
	}

	m := &Machine{
		DesiredState: desiredState,
		CurrentState: currentState,
		Buttons:      buttons,
		Joltage:      joltages,
	}

	return m, nil

}

func solveMachine(m Machine) int {
	// TODO: make more efficient
	res := math.MaxInt
	var recurse func(currState map[int]bool, presses int, buttons []Button)
	recurse = func(currState map[int]bool, presses int, buttons []Button) {

		if maps.Equal(currState, m.DesiredState) {
			res = min(res, presses)
			return
		}

		if presses > len(m.DesiredState) {
			return
		}

		if len(buttons) == 0 {
			return
		}

		// press
		b := buttons[0]
		newButtons := append(buttons[1:], buttons[0])
		newState := toggle(currState, b)
		recurse(newState, presses+1, newButtons)

		// no press
		recurse(currState, presses, buttons[1:])

	}
	recurse(m.CurrentState, 0, m.Buttons)

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
		sum += solveMachine(m)
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

}
