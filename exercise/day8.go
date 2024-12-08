// day8.go is the implementation for the eighth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day8 represents the data necessary to process the Exercise
	Day8 struct {
		name string
		file string
	}

	AntennaMap struct {
		frequency [][]rune
	}

	AntennaMapCoordinate struct {
		x int
		y int
	}
)

// GetName returns the name of the Day 8 exercise
func (d *Day8) GetName() string {
	return d.name
}

// Run executes the solution for Day 8 by retrieving the default file contents and uses that data
func (d *Day8) Run(w io.Writer) {
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

// RunFromInput executs the Day 8 solution using the provided input data
func (d *Day8) RunFromInput(w io.Writer, input []string) {
	antennaMap := d.parseInput(input)

	// part 1
	numberAntinodes := d.Part1(antennaMap)
	w.Write([]byte(fmt.Sprintf("Day 8 - Part 1 - The number of antinodes in the map is %d.\n", numberAntinodes)))

	// // part 2
	// numLoops := d.Part2(antennaMap)
	// w.Write([]byte(fmt.Sprintf("Day 6 - Part 2 - The number of new blocks that result in a loop is %d.\n", numLoops)))
}

// Part1 calculates antinode locations and counts the number of antinodes (the rules are specified in
// the readme for the day)
func (d *Day8) Part1(antennaMap *AntennaMap) int {
	antennaFrequencies := d.getUniqueFrequencies(antennaMap)

	antinodeMap := antennaMap.Copy()
	for _, antennas := range antennaFrequencies {
		for _, antenna := range antennas {
			d.SetAntinodes(antinodeMap, antenna, antennas, '#')
		}
	}

	antinodeCount := antinodeMap.countOccurences('#')

	return antinodeCount
}

// Part2
func (d *Day8) Part2() int {
	return 0
}

// parseData
func (d *Day8) parseInput(input []string) *AntennaMap {
	var antennaMap AntennaMap

	// Convert each string in the input to a slice of runes
	for _, line := range input {
		row := []rune(line) // Convert string to rune slice
		antennaMap.frequency = append(antennaMap.frequency, row)
	}

	return &antennaMap
}

// getUniqueFrequencies takes an AntennaMap and returns a map object with a list of coordinates
// where a specific frequency can be found
func (d *Day8) getUniqueFrequencies(a *AntennaMap) map[rune][]AntennaMapCoordinate {
	antennaFrequencies := make(map[rune][]AntennaMapCoordinate)

	for y, row := range a.frequency {
		for x, frequency := range row {
			if frequency != '.' {
				antennaFrequencies[frequency] = append(antennaFrequencies[frequency], AntennaMapCoordinate{x: x, y: y})
			}
		}
	}

	return antennaFrequencies
}

func (d *Day8) SetAntinodes(a *AntennaMap, sourceAntenna AntennaMapCoordinate, antennaPositions []AntennaMapCoordinate, marker rune) {
	if a == nil || len(a.frequency) == 0 || len(a.frequency[0]) == 0 {
		return
	}

	colSize := len(a.frequency)
	rowSize := len(a.frequency[0])

	for _, position := range antennaPositions {
		if position.x == sourceAntenna.x && position.y == sourceAntenna.y {
			// the source is the same as the destination
			continue
		}

		deltaX := position.x - sourceAntenna.x
		deltaY := position.y - sourceAntenna.y

		// set the position in line
		antinodePositionX := deltaX + position.x
		antinodePositionY := deltaY + position.y

		if antinodePositionX < 0 || antinodePositionY < 0 || antinodePositionX >= rowSize || antinodePositionY >= colSize {
			continue
		}

		a.frequency[antinodePositionY][antinodePositionX] = '#'
	}
}

// getAbsDifference returns the absolute value of the difference between p1 and p2
func getAbsDifference(p1, p2 int) int {
	difference := p1 - p2
	if difference < 0 {
		difference *= -1
	}

	return difference
}

func (a *AntennaMap) countOccurences(r rune) int {
	numOccurences := 0
	for _, row := range a.frequency {
		for _, frequency := range row {
			if r == frequency {
				numOccurences++
			}
		}
	}

	return numOccurences
}

// Copy() returns a deep copy of the source AntennaMap
func (a *AntennaMap) Copy() *AntennaMap {
	copyAntennaMap := &AntennaMap{
		frequency: make([][]rune, len(a.frequency)),
	}

	// Deep copy each row of the grid
	for i, row := range a.frequency {
		// Create a new slice for each row and copy the contents
		copyAntennaMap.frequency[i] = make([]rune, len(row))
		copy(copyAntennaMap.frequency[i], row)
	}

	return copyAntennaMap
}

// printGrid provides a pretty-print of the AntennaMap to stdout
func (a *AntennaMap) Print() {
	fmt.Println("Grid:")
	for _, row := range a.frequency {
		for _, cell := range row {
			fmt.Printf("%c ", cell) // Print each rune with a space
		}
		fmt.Println() // Newline after each row
	}

	fmt.Println()
}
