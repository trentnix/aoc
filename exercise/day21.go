// day21.go is the implementation for the twenty-first day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day21 represents the data necessary to process the Exercise
	Day21 struct {
		name string
		file string
	}

	KeyLocation struct {
		x, y int
	}

	Keypad map[string]KeyLocation
)

// GetName returns the name of the Day 21 exercise
func (d *Day21) GetName() string {
	return d.name
}

// Run executes the solution for Day 21 by retrieving the default file contents and uses that data
func (d *Day21) Run(w io.Writer) {
	if d.file == "" {
		w.Write([]byte(fmt.Sprintf("A default input file is not specified.")))
		return
	}

	input, err := fileprocessing.ReadFile(d.file)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error trying to read the input file %s: %v.", d.file, err)))
		return
	}

	d.RunFromInput(w, input)
}

// RunFromInput executs the Day 21 solution using the provided input data
func (d *Day21) RunFromInput(w io.Writer, input []string) {
	// part 1
	sumCodeComplexity := d.CalculateComplexity(input, 2)
	w.Write([]byte(fmt.Sprintf("Day 21 - Part 1 - The sum of the complexity of the provided codes (2 robots) is %d.\n", sumCodeComplexity)))

	// part 2
	sumCodeComplexity = d.CalculateComplexity(input, 25)
	w.Write([]byte(fmt.Sprintf("Day 21 - Part 2 - The sum of the complexity of the provided codes (25 robots) is %d.\n", sumCodeComplexity)))
}

// CalculateComplexity determines the complexity calculation by finding the shortest
// functioning input for each keypadPresses value and multiplying it by the
// keypadPresses value (the numeric part). So if keypadPresses[x] is "029A", the
// complexity is:
//
// 29 * len(input on first directional pad)
//
// The value returned is the sum of all the individual complexity values. Every sequence ends in
// 'A', so the approach is to keep track of the current level and solve for every section that ends
// with 'A' all the way to the bottom (last robot), then adding the result length all the way back
// up to the top.
func (d *Day21) CalculateComplexity(keypadCodes []string, numRobots int) int {
	// the numeric keypad
	numMap := Keypad{
		"0": {1, 0},
		"A": {2, 0},
		"1": {0, 1},
		"2": {1, 1},
		"3": {2, 1},
		"4": {0, 2},
		"5": {1, 2},
		"6": {2, 2},
		"7": {0, 3},
		"8": {1, 3},
		"9": {2, 3},
	}

	// the directional keypad
	dirMap := Keypad{
		"^": {1, 1},
		"A": {2, 1},
		"<": {0, 0},
		"v": {1, 0},
		">": {2, 0},
	}

	sumCodeComplexity := 0
	cache := make(map[string][]int)

	for _, code := range keypadCodes {
		moves := doorSequence(code, "A", numMap)
		length := countSequences(moves, numRobots, 1, cache, dirMap)

		numericCode := strings.TrimLeft(code[:3], "0")
		multiplier, _ := strconv.Atoi(numericCode)

		sumCodeComplexity += multiplier * length
	}

	return sumCodeComplexity
}

// doorSequence takes the specified input that needs to be typed on the door and returns
// the robot keypad sequence that will build it
func doorSequence(input string, start string, numMap Keypad) string {
	chars := strings.Split(input, "")

	current := numMap[start]
	sequence := ""

	for _, char := range chars {
		destination := numMap[char]
		dx, dy := destination.x-current.x, destination.y-current.y

		horizontal, vertical := "", ""

		// build horizontal moves
		for i := 0; i < absInt(dx); i++ {
			if dx >= 0 {
				horizontal += string('>')
			} else {
				horizontal += string('<')
			}
		}

		// build vertical moves
		for i := 0; i < absInt(dy); i++ {
			if dy >= 0 {
				vertical += string('^')
			} else {
				vertical += string('v')
			}
		}

		// avoid the empty spot
		if current.y == 0 && destination.x == 0 {
			sequence += vertical
			sequence += horizontal
		} else if current.x == 0 && destination.y == 0 {
			sequence += horizontal
			sequence += vertical
		} else if dx < 0 {
			sequence += horizontal
			sequence += vertical
		} else {
			sequence += vertical
			sequence += horizontal
		}

		current = destination
		sequence += string('A')
	}

	return sequence
}

// robotSequence takes the specified input that needs to be typed on the robot keypad and returns
// the robot keypad sequence that will build it
func robotSequence(input string, start string, dirMap Keypad) string {
	current := dirMap[start]
	sequence := ""

	chars := strings.Split(input, "")

	for _, char := range chars {
		destination := dirMap[char]
		dx := destination.x - current.x
		dy := destination.y - current.y

		horizontal, vertical := "", ""

		// build horizontal moves
		for i := 0; i < absInt(dx); i++ {
			if dx >= 0 {
				horizontal += string('>')
			} else {
				horizontal += string('<')
			}
		}

		// build vertical moves
		for i := 0; i < absInt(dy); i++ {
			if dy >= 0 {
				vertical += string('^')
			} else {
				vertical += string('v')
			}
		}

		// avoid the empty spot
		if current.x == 0 && destination.y == 1 {
			sequence += horizontal
			sequence += vertical
		} else if current.y == 1 && destination.x == 0 {
			sequence += vertical
			sequence += horizontal
		} else if dx < 0 {
			sequence += horizontal
			sequence += vertical
		} else {
			sequence += vertical
			sequence += horizontal
		}

		current = destination
		sequence += string('A')
	}

	return sequence
}

// countSequences caches a specific sequence and adds its length to a running total for a given input
func countSequences(input string, maxRobots, robotLevel int, cache map[string][]int, dirMap Keypad) int {
	if val, ok := cache[input]; ok && robotLevel <= len(val) && val[robotLevel-1] != 0 {
		// the input and level have already been cached - use the cached value
		return val[robotLevel-1]
	}

	// initialize the cache
	if _, ok := cache[input]; !ok {
		cache[input] = make([]int, maxRobots)
	}

	// get the next sequence of movements
	sequence := robotSequence(input, "A", dirMap)
	if robotLevel == maxRobots {
		// this is the final keypad - return the length of the sequence
		return len(sequence)
	}

	// split the sequence into a full parent keypad sequence (each step ends with an 'A')
	steps := splitSequence(sequence)
	count := 0
	for _, step := range steps {
		// for each split sequence, run countSequences for the next level
		sequenceCount := countSequences(step, maxRobots, robotLevel+1, cache, dirMap)
		// sum the counts returned
		count += sequenceCount
	}

	// cache and return the total count
	cache[input][robotLevel-1] = count
	return count
}

// splitSequence splits the specified input into various string chunks where
// the last character is an 'A', indicating a complete move sequence
func splitSequence(input string) []string {
	var result []string
	var current string

	for _, char := range input {
		current += string(char)
		if char == 'A' {
			result = append(result, current)
			current = ""
		}
	}
	return result
}
