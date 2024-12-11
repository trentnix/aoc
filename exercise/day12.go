// day12.go is the implementation for the twelth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day12 represents the data necessary to process the Exercise
	Day12 struct {
		name string
		file string
	}
)

// GetName returns the name of the Day 12 exercise
func (d *Day12) GetName() string {
	return d.name
}

// Run executes the solution for Day 12 by retrieving the default file contents and uses that data
func (d *Day12) Run(w io.Writer) {
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

// RunFromInput executs the Day 12 solution using the provided input data
func (d *Day12) RunFromInput(w io.Writer, input []string) {
	// data, err := // parse the data
	// if err != nil {
	// 	w.Write([]byte(fmt.Sprintf("There was an error trying to parse the input: %v.", err)))
	// 	return
	// }
}

// Part1
func (d *Day12) Part1() int {
	return 0
}

// Part2
func (d *Day12) Part2() int {
	return 0
}

// parseInput
func (d *Day12) parseInput(input []string) /* return types */ {
}
