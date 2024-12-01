// day1.go is the implementation for the first day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

// Day1 represents the data necessary to process an Exercise for the first day
type Day1 struct {
	name  string
	file  string
	order int
}

// GetName returns the name of the Day1 exercise
func (d *Day1) GetName() string {
	return d.name
}

// Run executes the solution for Day1 by retrieving the default file contents and uses that data
func (d *Day1) Run(w io.Writer) {
	if d.file != "" {
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

// RunFromInput executs the Day1 solution using the provided input data
func (d *Day1) RunFromInput(w io.Writer, input []string) {
}
