// day3.go is the implementation for the third day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day3 represents the data necessary to process the Exercise
	Day3 struct {
		name string
		file string
	}

	instruction struct {
		operand string
		v1      int
		v2      int
	}
)

// GetName returns the name of the Day 3 exercise
func (d *Day3) GetName() string {
	return d.name
}

// Run executes the solution for Day 3 by retrieving the default file contents and uses that data
func (d *Day3) Run(w io.Writer) {
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

// RunFromInput executs the Day 2 solution using the provided input data
func (d *Day3) RunFromInput(w io.Writer, input []string) {
	instructions, err := d.parseInput(input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error trying to parse the input: %v.", err)))
		return
	}

	sumOfMultiplications := d.Part1(instructions)

	// part 1
	w.Write([]byte(fmt.Sprintf("Day 3 - Part 1 - The sum of mul() instructions is %d.\n", sumOfMultiplications)))
}

// Part1
func (d *Day3) Part1(instructions []instruction) int {
	sum := 0

	for _, i := range instructions {
		if i.operand == "mul" {
			sum += i.v1 * i.v2
		}
	}

	return sum
}

// Part2
func (d *Day3) Part2(instructions []instruction) int {
	return 0
}

// parseInput
func (d *Day3) parseInput(input []string) ([]instruction, error) {
	var instructions []instruction

	// make a single long string
	var data string
	for _, s := range input {
		data += s
	}

	// Regular expression to match valid mul(x, y) patterns
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(data, -1)

	// Iterate over matches and parse them into instructions
	for _, match := range matches {
		v1, err1 := strconv.Atoi(match[1])
		v2, err2 := strconv.Atoi(match[2])
		if err1 == nil && err2 == nil {
			instructions = append(instructions, instruction{
				operand: "mul",
				v1:      v1,
				v2:      v2,
			})
		} else {
			return instructions, fmt.Errorf("There was an error parsing the input")
		}
	}

	return instructions, nil
}
