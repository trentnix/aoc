// day4.go is the implementation for the fourth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"log"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day4 represents the data necessary to process the Exercise
	Day4 struct {
		name string
		file string
	}

	coordinate struct {
		x int
		y int
	}
)

// GetName returns the name of the Day 4 exercise
func (d *Day4) GetName() string {
	return d.name
}

// Run executes the solution for Day 4 by retrieving the default file contents and uses that data
func (d *Day4) Run(w io.Writer) {
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

// RunFromInput executs the Day 4 solution using the provided input data
func (d *Day4) RunFromInput(w io.Writer, input []string) {
	grid := make([][]rune, len(input))

	for i, str := range input {
		// Convert each string into a slice of runes
		grid[i] = []rune(str)
	}

	numberOfXmasInstances := d.Part1(grid)

	// part 1
	w.Write([]byte(fmt.Sprintf("Day 4 - Part 1 - The number of 'XMAS' instances is %d.\n", numberOfXmasInstances)))

	numberOfMasXInstances := d.Part2(grid)

	// part 1
	w.Write([]byte(fmt.Sprintf("Day 4 - Part 2 - The number of 'MAS' in an X instances is %d.\n", numberOfMasXInstances)))
}

// Part1 counts the number of instances of the word 'XMAS' in the input character (rune) grid
func (d *Day4) Part1(input [][]rune) int {
	numColumns := len(input)
	if numColumns == 0 {
		log.Fatalf("invalid input")
	}

	numRows := len(input[0])
	if numRows < 4 {
		log.Fatalf("invalid input")
	}

	lengthTerm := len("XMAS")

	countXMAS := 0

	for y := 0; y < numColumns; y++ {
		for x := 0; x < numRows; x++ {
			// build coordinates for right
			if x <= numRows-lengthTerm {
				if d.matchXmas(input, coordinate{x, y}, coordinate{x + 1, y}, coordinate{x + 2, y}, coordinate{x + 3, y}) {
					countXMAS++
				}
			}

			// build coordinates for left
			if x >= lengthTerm-1 {
				if d.matchXmas(input, coordinate{x, y}, coordinate{x - 1, y}, coordinate{x - 2, y}, coordinate{x - 3, y}) {
					countXMAS++
				}
			}

			// build coordinates for up
			if y >= lengthTerm-1 {
				if d.matchXmas(input, coordinate{x, y}, coordinate{x, y - 1}, coordinate{x, y - 2}, coordinate{x, y - 3}) {
					countXMAS++
				}
			}

			// build coordinates for down
			if y <= numColumns-lengthTerm {
				if d.matchXmas(input, coordinate{x, y}, coordinate{x, y + 1}, coordinate{x, y + 2}, coordinate{x, y + 3}) {
					countXMAS++
				}
			}

			// build coordinates for up-left
			if y >= lengthTerm-1 && x >= lengthTerm-1 {
				if d.matchXmas(input, coordinate{x, y}, coordinate{x - 1, y - 1}, coordinate{x - 2, y - 2}, coordinate{x - 3, y - 3}) {
					countXMAS++
				}
			}

			// build coordinates for down-left
			if y <= numColumns-lengthTerm && x >= lengthTerm-1 {
				if d.matchXmas(input, coordinate{x, y}, coordinate{x - 1, y + 1}, coordinate{x - 2, y + 2}, coordinate{x - 3, y + 3}) {
					countXMAS++
				}
			}

			// build coordinates for down-right
			if y <= numColumns-lengthTerm && x <= numRows-lengthTerm {
				if d.matchXmas(input, coordinate{x, y}, coordinate{x + 1, y + 1}, coordinate{x + 2, y + 2}, coordinate{x + 3, y + 3}) {
					countXMAS++
				}
			}

			// build coordinates for up-right
			if y >= lengthTerm-1 && x <= numRows-lengthTerm {
				if d.matchXmas(input, coordinate{x, y}, coordinate{x + 1, y - 1}, coordinate{x + 2, y - 2}, coordinate{x + 3, y - 3}) {
					countXMAS++
				}
			}
		}
	}

	return countXMAS
}

// matchXmas determines whether the word 'XMAS' is found at the four specified coordinates of
// the input grid
func (d *Day4) matchXmas(input [][]rune, x, m, a, s coordinate) bool {
	if input[x.y][x.x] != 'X' {
		return false
	}
	if input[m.y][m.x] != 'M' {
		return false
	}
	if input[a.y][a.x] != 'A' {
		return false
	}
	if input[s.y][s.x] != 'S' {
		return false
	}
	return true
}

// Part2 counts the number of instances of the word 'MAS' that make an X
func (d *Day4) Part2(input [][]rune) int {
	numColumns := len(input)
	if numColumns == 0 {
		log.Fatalf("invalid input")
	}

	numRows := len(input[0])
	if numRows < 4 {
		log.Fatalf("invalid input")
	}

	countMAS := 0

	for y := 1; y < numColumns-1; y++ {
		for x := 1; x < numRows-1; x++ {
			if d.matchMasX(input, coordinate{x, y}) {
				countMAS++
			}
		}
	}

	return countMAS
}

// matchMasX returns true if, at the given coordinate, the word 'MAS' makes an overlapping
// X with itself ('MAS' and 'MAS', 'MAS' and 'SAM', 'SAM' and 'SAM', or 'SAM' and 'MAS')
func (d *Day4) matchMasX(input [][]rune, a coordinate) bool {
	if input[a.y][a.x] != 'A' {
		return false
	}

	if ((input[a.y-1][a.x-1] == 'M' && input[a.y+1][a.x+1] == 'S') || (input[a.y-1][a.x-1] == 'S' && input[a.y+1][a.x+1] == 'M')) &&
		((input[a.y-1][a.x+1] == 'M' && input[a.y+1][a.x-1] == 'S') || (input[a.y-1][a.x+1] == 'S' && input[a.y+1][a.x-1] == 'M')) {
		return true
	}

	return false
}
