package main

import (
	"aoc-2025/common"
	"bytes"
	"fmt"
	"log"
	"strconv"
)

type Machine struct {
	DesiredState map[int]bool
	CurrentState map[int]bool
	Buttons      []Button
	Joltage      []int
}

func (m *Machine) Toggle(s Button) {
	for _, b := range s.Bulbs {
		m.CurrentState[b] = !m.CurrentState[b]
	}
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

func main() {
	filePath := "./inputExample.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	machines, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(machines)

}
