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
		v1 int
		v2 int
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

// RunFromInput executs the Day 3 solution using the provided input data
func (d *Day3) RunFromInput(w io.Writer, input []string) {
	instructions, err := d.parseInputRaw(input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error trying to parse the input: %v.", err)))
		return
	}

	sumOfMultiplications := d.Part1(instructions)

	// part 1
	w.Write([]byte(fmt.Sprintf("Day 3 - Part 1 - The sum of mul() instructions is %d.\n", sumOfMultiplications)))

	instructions, err = d.parseInputApplyConditionals(input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error trying to parse the input: %v.", err)))
		return
	}

	sumOfMultiplications = d.Part2(instructions)

	// part 2
	w.Write([]byte(fmt.Sprintf("Day 3 - Part 2 - The sum of mul() instructions is %d.\n", sumOfMultiplications)))
}

// Part1 iterates over the instructions and returns the sum of the instruction multiples
func (d *Day3) Part1(instructions []instruction) int {
	sum := 0

	for _, i := range instructions {
		sum += i.v1 * i.v2
	}

	return sum
}

// Part2 performs the same functionality as part 1, but has a different input. As a result,
// we can just call the Part1 method.
func (d *Day3) Part2(instructions []instruction) int {
	return d.Part1(instructions)
}

// parseInputRaw parses the input by extracting the mul(x,y) instructions into
// an []instruction type
func (d *Day3) parseInputRaw(input []string) ([]instruction, error) {
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
				v1: v1,
				v2: v2,
			})
		} else {
			return instructions, fmt.Errorf("There was an error parsing the input")
		}
	}

	return instructions, nil
}

// parseInputApplyConditionals parses the input by extracting the mul(x,y) instructions
// while applying conditional logic don't() and do() which enables or disables the
// instructions. It is assumed that instructions are enabled initially.
func (d *Day3) parseInputApplyConditionals(input []string) ([]instruction, error) {
	var instructions []instruction

	// make a single long string
	var data string
	for _, s := range input {
		data += s
	}

	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)
	doRe := regexp.MustCompile(`do\(\)`)

	// Initialize a flag to track whether to ignore mul() instructions
	ignore := false

	// Iterate over the input using regex matches in sequence
	re := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(data, -1)

	for _, match := range matches {
		if dontRe.MatchString(match[0]) {
			// If "don't()" is encountered, set the ignore flag to true
			ignore = true
		} else if doRe.MatchString(match[0]) {
			// If "do()" is encountered, reset the ignore flag
			ignore = false
		} else if mulRe.MatchString(match[0]) && !ignore {
			// If a valid mul(x, y) is encountered and not ignoring
			v1, err1 := strconv.Atoi(match[1])
			v2, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				instructions = append(instructions, instruction{
					v1: v1,
					v2: v2,
				})
			}
		}
	}

	return instructions, nil
}
