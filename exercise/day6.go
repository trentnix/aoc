// day6.go is the implementation for the TBD day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day6 represents the data necessary to process the Exercise
	Day6 struct {
		name string
		file string
	}
)

// GetName returns the name of the Day 6 exercise
func (d *Day6) GetName() string {
	return d.name
}

// Run executes the solution for Day 6 by retrieving the default file contents and uses that data
func (d *Day6) Run(w io.Writer) {
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

// RunFromInput executs the Day 6 solution using the provided input data
func (d *Day6) RunFromInput(w io.Writer, input []string) {
	// data, err := // parse the data
	// if err != nil {
	// 	w.Write([]byte(fmt.Sprintf("There was an error trying to parse the input: %v.", err)))
	// 	return
	// }
}

// Part1
func (d *Day6) Part1() int {
	return 0
}

// Part2
func (d *Day6) Part2() int {
	return 0
}

// parseData
// func (d *Day6) parseData(input []string) /* return types */ {
// }
