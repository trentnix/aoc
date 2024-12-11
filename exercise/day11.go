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
	stonesP2 := d.parseInput(input[0])

	// part 1
	numBlinks := 25
	numStones := d.ProcessStones(stones, numBlinks)
	w.Write([]byte(fmt.Sprintf("Day 11 - Part 1 - The number of stones after %d blinks is %d.\n", numBlinks, numStones)))

	// part 1
	numBlinks = 75
	numStones = d.ProcessStones(stonesP2, numBlinks)
	w.Write([]byte(fmt.Sprintf("Day 11 - Part 2 - The number of stones after %d blinks is %d.\n", numBlinks, numStones)))
}

// ProcessStones applies the rules of a 'blink' (per the day's assignment) and determines the number of
// stones after 25 blinks
//
// Order doesn't matter, so we keep track of the number of instances of a stone with a specific inscription
// and apply the rules accordingly
func (d *Day11) ProcessStones(input []uint64, blinks int) uint64 {
	stones := make(map[uint64]uint64)
	for _, stone := range input {
		stones[stone]++
	}

	for i := 0; i < blinks; i++ {
		newStones := make(map[uint64]uint64)
		for key, numElements := range stones {
			if key == 0 {
				newStones[1] += numElements
				continue
			}

			keyString := strconv.FormatUint(key, 10)
			if len(keyString)%2 == 0 {
				// even number of digits
				middleIndex := len(keyString) / 2
				leftString := keyString[:middleIndex]
				rightString := keyString[middleIndex:]

				left, _ := strconv.ParseUint(leftString, 10, 64)
				right, _ := strconv.ParseUint(rightString, 10, 64)

				newStones[left] += numElements
				newStones[right] += numElements
			} else {
				// odd number of digits
				newStones[key*defaultMultiplier] += numElements
			}
		}

		stones = newStones
	}

	numStones := uint64(0)
	for key := range stones {
		numStones += stones[key]
	}

	return numStones
}

// parseInput parses the specified input into a slice of string values by separating
// the input by a space
func (d *Day11) parseInput(input string) []uint64 {
	var iStones []uint64

	stones := strings.Split(input, " ")
	for _, stone := range stones {
		iStone, err := strconv.ParseUint(stone, 10, 64)
		if err != nil {
			fmt.Printf("could not convert %s", stone)
		}
		iStones = append(iStones, iStone)
	}

	return iStones
}
