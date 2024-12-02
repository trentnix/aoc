// day1.go is the implementation for the first day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

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

// RunFromInput executs the Day1 solution using the provided input data
func (d *Day1) RunFromInput(w io.Writer, input []string) {
	// parse the input into two arrays
	left, right, err := d.parseIntoLists(input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error trying to parse the input: %v.", err)))
		return
	}

	sumDifferences := d.Part1(left, right)

	// part 1
	w.Write([]byte(fmt.Sprintf("Day 1 - Part 1 - The sum of the distances between the left and right list is %d.", sumDifferences)))
}

func (d *Day1) parseIntoLists(input []string) ([]int, []int, error) {
	var left, right []int

	// Parse the integers from the strings and populate left and right slices
	for rowNumber, line := range input {
		// Split the line into parts
		parts := strings.Fields(line)
		if len(parts) == 2 {
			// Convert strings to integers
			l, errL := strconv.Atoi(parts[0])
			r, errR := strconv.Atoi(parts[1])
			if errL == nil && errR == nil {
				left = append(left, l)
				right = append(right, r)
			} else {
				return nil, nil, fmt.Errorf("There was an error parsing the input on row %d: l(%v), r(%v)", rowNumber, errL, errR)
			}
		} else {
			return nil, nil, fmt.Errorf("There was an error parsing the input on row %d: invalid format", rowNumber)
		}
	}

	return left, right, nil
}

func (d *Day1) Part1(l []int, r []int) int {
	// Sort the left and right slices
	sort.Ints(l)
	sort.Ints(r)

	if len(l) != len(r) {
		return -1
	}

	var sum int

	for i := 0; i < len(l); i++ {
		if l[i] > r[i] {
			sum += l[i] - r[i]
		} else {
			sum += r[i] - l[i]
		}
	}

	return sum
}
