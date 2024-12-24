// day24.go is the implementation for the twenty-fourth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day24 represents the data necessary to process the Exercise
	Day24 struct {
		name string
		file string
	}

	Bits        map[string]bool
	Instruction struct {
		Source      [2]string
		Destination string
		Operation   int
	}
)

const (
	AND = iota
	OR
	XOR
)

// GetName returns the name of the Day 24 exercise
func (d *Day24) GetName() string {
	return d.name
}

// Run executes the solution for Day 24 by retrieving the default file contents and uses that data
func (d *Day24) Run(w io.Writer) {
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

// RunFromInput executs the Day 24 solution using the provided input data
func (d *Day24) RunFromInput(w io.Writer, input []string) {
	bits, instructions := d.parseInput(input)
	zVal := d.Part1(bits, instructions)
	w.Write([]byte(fmt.Sprintf("Day 24 - Part 1 - The value of the wires that start with 'z' is %d.\n", zVal)))
}

// Part1 runs the instructions (after their constituent wires have been loaded with values)
// and calculates the decimal value of the wires that start with 'z', as per the instructions
func (d *Day24) Part1(bits Bits, instructions []Instruction) int {
	hasRun := make(map[Instruction]bool)
	numInstructions := len(instructions)

	for {
		processedInstructions := false
		for i := 0; i < numInstructions; i++ {
			if hasRun[instructions[i]] {
				continue
			}

			wire0, ok0 := bits[instructions[i].Source[0]]
			wire1, ok1 := bits[instructions[i].Source[1]]

			if ok0 && ok1 {
				// the values have been loaded
				result, wire := instructions[i].Run(wire0, wire1)
				bits[wire] = result
				hasRun[instructions[i]] = true
				processedInstructions = true
			}
		}

		if !processedInstructions {
			break
		}
	}

	zVal := getZVal(bits)

	return zVal
}

// Part2
func (d *Day24) Part2() int {
	return 0
}

// parseInput takes the specified input and produces a Bits map and a slice of
// Instructions to run
func (d *Day24) parseInput(input []string) (Bits, []Instruction) {
	bits := make(Bits)
	var instructions []Instruction

	// Flag to indicate whether we're parsing bits or instructions
	parsingBits := true

	for _, line := range input {
		line = strings.TrimSpace(line)

		// Skip empty lines
		if line == "" {
			parsingBits = false
			continue
		}

		if parsingBits {
			// Parse Bits lines: key:value
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				// Handle invalid format if necessary
				continue
			}
			key := strings.TrimSpace(parts[0])
			valueStr := strings.TrimSpace(parts[1])
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				// Handle invalid value if necessary
				continue
			}
			bits[key] = (value != 0)
		} else {
			// Parse Instruction lines: source0 OP source1 -> destination
			// Example: x00 AND y00 -> z00

			// Split the line by "->" to separate the operation from the destination
			parts := strings.Split(line, "->")
			if len(parts) != 2 {
				// Handle invalid format if necessary
				continue
			}
			operationPart := strings.TrimSpace(parts[0])
			destination := strings.TrimSpace(parts[1])

			// Split the operation part into tokens
			tokens := strings.Fields(operationPart)
			if len(tokens) != 3 {
				// Handle invalid format if necessary
				continue
			}

			source0 := tokens[0]
			opStr := strings.ToUpper(tokens[1])
			source1 := tokens[2]

			// Map the operation string to the corresponding constant
			var op int
			switch opStr {
			case "AND":
				op = AND
			case "OR":
				op = OR
			case "XOR":
				op = XOR
			default:
				// Handle unknown operation if necessary
				continue
			}

			// Create the Instruction and append to the slice
			instr := Instruction{
				Source:      [2]string{source0, source1},
				Destination: destination,
				Operation:   op,
			}
			instructions = append(instructions, instr)
		}
	}

	return bits, instructions
}

// Run performs the Instruction's specified operation on the source values specified
// in the input values and returns the result and the destination wire name
func (i *Instruction) Run(wire0, wire1 bool) (result bool, outputWire string) {
	iResult, i0, i1 := 0, 0, 0
	if wire0 {
		i0 = 1
	}

	if wire1 {
		i1 = 1
	}

	switch i.Operation {
	case AND:
		iResult = i0 & i1
	case OR:
		iResult = i0 | i1
	case XOR:
		iResult = i0 ^ i1
	}

	if iResult != 0 {
		result = true
	}

	return result, i.Destination
}

// getZVal constructs the decimal value of all bits whose wire name starts with 'z' at
// the index provided in the name - e.g. "z03" = true results in 8 (2^3), "z05" = true
// results in 32 (2^5)
func getZVal(bits Bits) int {
	result := 0
	for wire, bitValue := range bits {
		if wire[0] == 'z' {
			zIndex, _ := strconv.Atoi(wire[1:])
			if bitValue {
				result += 1 << zIndex
			}
		}
	}

	return result
}
