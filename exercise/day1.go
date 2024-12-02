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
	name string
	file string
}

// GetName returns the name of the Day 1 exercise
func (d *Day1) GetName() string {
	return d.name
}

// Run executes the solution for Day 1 by retrieving the default file contents and uses that data
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

// RunFromInput executs the Day 1 solution using the provided input data
func (d *Day1) RunFromInput(w io.Writer, input []string) {
	// parse the input into two arrays
	left, right, err := d.parseIntoLists(input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error trying to parse the input: %v.", err)))
		return
	}

	sumDifferences := d.Part1(left, right)

	// part 1
	w.Write([]byte(fmt.Sprintf("Day 1 - Part 1 - The sum of the distances between the left and right list is %d.\n", sumDifferences)))

	sumSimilarityScores := d.Part2(left, right)

	// part 2
	w.Write([]byte(fmt.Sprintf("Day 1 - Part 2 - The sum of the similarity scores between the left and right lists is %d.\n", sumSimilarityScores)))
}

// parseIntoLists parses the string input into two integer slices. An error is returned
// if there was an error during the parsing effort.
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

// Part1 calculates the differences between the smallest value in each list, then moves to the
// second smallest number in each and adds that differences to the first differences, and so on.
// The value returned is the sum of all of the differences.
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

// Part2 discerns a "similary score" by counting the instances of each left value in the right list
// and multiplying the number of instances it appears times the value itself. So if 3 is the value on
// the left side and it appears 4 times in the right list, 3 * 4 = 12. This method returns the sum of
// similarity scores.
func (d *Day1) Part2(l []int, r []int) int {
	repeatCount := make(map[int]int)
	for _, val := range r {
		repeatCount[val]++
	}

	sum := 0
	for _, val := range l {
		sum += repeatCount[val] * val
	}

	return sum
}
