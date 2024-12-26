// day25.go is the implementation for the twenty-fifth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day25 represents the data necessary to process the Exercise
	Day25 struct {
		name string
		file string
	}

	Keys  []Schematic
	Locks []Schematic

	Schematic struct {
		val     [][]rune
		heights []int
	}
)

// GetName returns the name of the Day 25 exercise
func (d *Day25) GetName() string {
	return d.name
}

// Run executes the solution for Day 25 by retrieving the default file contents and uses that data
func (d *Day25) Run(w io.Writer) {
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

// RunFromInput executs the Day 25 solution using the provided input data
func (d *Day25) RunFromInput(w io.Writer, input []string) {
	locks, keys := d.parseInput(input)
	fits := d.Part1(locks, keys)
	w.Write([]byte(fmt.Sprintf("Day 25 - Part 1 - The number of fits between locks and keys is %d.\n", fits)))
}

// Part1 determines how many fits there are between the specified locks and keys. A fit is when
// a particular key height and particular lock height don't exceed the total number of pins in a
// column, as per the assignment
func (d *Day25) Part1(locks Locks, keys Keys) int {
	fits := 0

	heightPins := 5
	numColumns := 5
	for _, lock := range locks {
		for _, key := range keys {
			isFit := true
			for i := 0; i < numColumns; i++ {
				if lock.heights[i]+key.heights[i] > heightPins {
					isFit = false
				}
			}

			if isFit {
				fits++
			}
		}
	}

	return fits
}

// Part2
func (d *Day25) Part2() int {
	return 0
}

// parseInput parses the input into Locks and Keys
func (d *Day25) parseInput(input []string) (Locks, Keys) {
	var (
		locks            []Schematic
		keys             []Schematic
		currentSchematic Schematic
	)

	for i, line := range input {
		if line == "" {
			// process the completed schematic
			if len(currentSchematic.val) > 0 {
				currentSchematic.heights = calculateHeights(currentSchematic.val)
				if isLock(currentSchematic) {
					locks = append(locks, currentSchematic)
				} else if isKey(currentSchematic) {
					keys = append(keys, currentSchematic)
				}
			}

			// Reset for the next schematic
			currentSchematic = Schematic{}
			continue
		}

		// add the current line to the schematic
		currentSchematic.val = append(currentSchematic.val, []rune(line))

		// handle the last schematic
		if i == len(input)-1 && len(currentSchematic.val) > 0 {
			currentSchematic.heights = calculateHeights(currentSchematic.val)
			if isLock(currentSchematic) {
				locks = append(locks, currentSchematic)
			} else if isKey(currentSchematic) {
				keys = append(keys, currentSchematic)
			}
		}
	}

	return locks, keys
}

// Helper to check if a schematic is a lock
func isLock(s Schematic) bool {
	if len(s.val) == 0 {
		return false
	}

	for _, cell := range s.val[0] {
		if cell != '#' {
			return false
		}
	}

	return true
}

// isKey checks if a schematic is a key
func isKey(s Schematic) bool {
	if len(s.val) == 0 {
		return false
	}

	bottomRow := s.val[len(s.val)-1]
	for _, cell := range bottomRow {
		if cell != '#' {
			return false
		}
	}

	return true
}

// calculateHeights calculates the heights of each column in a schematic (number of '#' values
// in each column excluding first and last rows)
func calculateHeights(val [][]rune) []int {
	if len(val) <= 2 {
		return []int{}
	}

	heights := make([]int, len(val[0]))
	for i := 1; i < len(val)-1; i++ {
		for j, cell := range val[i] {
			if cell == '#' {
				heights[j]++
			}
		}
	}

	return heights
}
