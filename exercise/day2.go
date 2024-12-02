// day2.go is the implementation for the second day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day2 represents the data necessary to process the Exercise
	Day2 struct {
		name string
		file string
	}

	Level  int
	Report []Level
)

// GetName returns the name of the Day 2 exercise
func (d *Day2) GetName() string {
	return d.name
}

// Run executes the solution for Day 2 by retrieving the default file contents and uses that data
func (d *Day2) Run(w io.Writer) {
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
func (d *Day2) RunFromInput(w io.Writer, input []string) {
	reports, err := d.parseIntoReports(input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error trying to parse the input: %v.", err)))
		return
	}

	numSafeReports := d.Part1(reports)

	// part 1
	w.Write([]byte(fmt.Sprintf("Day 2 - Part 1 - The sum of safe reports is %d.\n", numSafeReports)))
}

// Part1 counts which Report entries are "safe"
func (d *Day2) Part1(reports []Report) int {
	numSafeReports := 0
	for _, report := range reports {
		if d.isReportSafe(report) {
			numSafeReports++
		}
	}

	return numSafeReports
}

// parseIntoReports takes the input string array and converts it into an array of Report
// structures (which is an array of Level structures, each of which is just an int)
func (d *Day2) parseIntoReports(input []string) ([]Report, error) {
	var result []Report

	for _, line := range input {
		// Split the line into fields by spaces
		parts := strings.Fields(line)
		var report Report

		for _, part := range parts {
			// Convert each part to an integer
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error parsing '%s': %w", part, err)
			}
			report = append(report, Level(num))
		}

		// Append the row of integers to the result
		result = append(result, report)
	}

	return result, nil
}

// isReportSafe determins whether a given Report (an array of integers, each
// called a 'Level') shows:
//
//	all levels are either all increasing or all decreasing
//	any two adjacent levels differ by at least one and at most three
func (d *Day2) isReportSafe(r Report) bool {
	if len(r) <= 1 {
		return false
	}

	var isIncreasing bool

	levelDifference := int(r[0]) - int(r[1])
	if levelDifference < 0 {
		isIncreasing = true
	}

	if !d.isLevelDifferenceSafe(levelDifference, 1, 3) {
		return false
	}

	for index := 1; index < len(r)-1; index++ {
		levelDifference = int(r[index]) - int(r[index+1])
		if isIncreasing {
			if levelDifference > 0 {
				return false
			}
		} else {
			if levelDifference < 0 {
				return false
			}
		}

		if !d.isLevelDifferenceSafe(levelDifference, 1, 3) {
			return false
		}
	}

	return true
}

// isLevelDifferenceSafe determines whether the absolute value of the difference value
// is between the min and max
func (d *Day2) isLevelDifferenceSafe(difference int, min int, max int) bool {
	if difference < 0 {
		difference *= -1
	}

	if difference > max || difference < min {
		return false
	}

	return true
}
