// Day11.go is the implementation for the eleventh day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day11 represents the data necessary to process the Exercise
	Day11 struct {
		name string
		file string
	}
)

const (
	// this is the value used to multiply a stone engraving
	defaultMultiplier = uint64(2024)
)

// GetName returns the name of the day 11 exercise
func (d *Day11) GetName() string {
	return d.name
}

// Run executes the solution for day 11 by retrieving the default file contents and uses that data
func (d *Day11) Run(w io.Writer) {
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

// RunFromInput executs the day 11 solution using the provided input data
func (d *Day11) RunFromInput(w io.Writer, input []string) {
	if len(input) != 1 {
		log.Fatalf("the input is invalid")
	}

	stones := d.parseInput(input[0])

	// part 1
	numBlinks := 25
	numStones := d.Part1(stones, numBlinks)
	w.Write([]byte(fmt.Sprintf("Day 11 - Part 1 - The number of stones after %d blinks is %d.\n", numBlinks, numStones)))
}

// Part1 applies the rules of a 'blink' (per the day's assignment) and determines the number of
// stones after 25 blinks
func (d *Day11) Part1(stones []string, blinks int) int {
	for i := 0; i < blinks; i++ {
		stones = d.blink(stones)
		if stones == nil {
			log.Fatalf("blink failed")
		}
	}

	return len(stones)
}

// Part2
func (d *Day11) Part2() int {
	return 0
}

// blink applies the rules of a 'blink' to the set of stones specified
//
//   - a stone with 0 engraved turns into a stone with 1 engraved
//   - a stone with an even number of digits engraved turns into two stones engraved with
//     the left half of the digits and the right half of the digits
//   - a stone that doesn't fall into one of those categories has the value engraved
//     multiplied by 2024
func (d *Day11) blink(stones []string) []string {
	var stonesAfterBlink []string
	for _, stone := range stones {
		switch {
		case stone == "0":
			stonesAfterBlink = append(stonesAfterBlink, "1")
		case len(stone)%2 == 0:
			middle := len(stone) / 2

			stonesAfterBlink = append(stonesAfterBlink, d.removeLeadingZeroes(stone[:middle]))
			stonesAfterBlink = append(stonesAfterBlink, d.removeLeadingZeroes(stone[middle:]))

		default:
			value, err := strconv.ParseUint(stone, 10, 64)
			if err != nil {
				fmt.Println("Error converting string to uint64:", err)
				return nil
			}

			value *= defaultMultiplier
			newStone := strconv.FormatUint(value, 10)
			stonesAfterBlink = append(stonesAfterBlink, newStone)
		}
	}

	return stonesAfterBlink
}

// removeLeadingZeroes takes a string of digits and removes any leading 0s
func (d *Day11) removeLeadingZeroes(stone string) string {
	for i := 0; i < len(stone); i++ {
		if stone[i] != '0' {
			return stone[i:]
		}
	}

	return "0"
}

// parseInput parses the specified input into a slice of string values by separating
// the input by a space
func (d *Day11) parseInput(input string) []string {
	return strings.Split(input, " ")
}
