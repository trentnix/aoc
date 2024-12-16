// day15.go is the implementation for the fifteenth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day15 represents the data necessary to process the Exercise
	Day15 struct {
		name string
		file string
	}

	BoxMap [][]rune

	Instructions string
)

// GetName returns the name of the Day 15 exercise
func (d *Day15) GetName() string {
	return d.name
}

// Run executes the solution for Day 15 by retrieving the default file contents and uses that data
func (d *Day15) Run(w io.Writer) {
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

// RunFromInput executs the Day 15 solution using the provided input data
func (d *Day15) RunFromInput(w io.Writer, input []string) {
	boxMap, instructions := d.parseInput(input)
	sumCoordinateValues := d.Part1(boxMap, instructions)
	w.Write([]byte(fmt.Sprintf("Day 15 - Part 1 - The sum of the box coordinate values is %d.", sumCoordinateValues)))
}

// Part1 calculates the sum of the coordinate values per the instructions by
// moving a robot (specified at the @ element) according to the instructions list
//
// - a coordinate value is y * 100 + x
// - boxes, specified by O, can be pushed into an open space
// - instructions are < (left), ^ (up), > (right), and v (down)
func (d *Day15) Part1(boxMap BoxMap, instructions Instructions) int {
	posY, posX := boxMap.Find('@')

	for _, instruction := range instructions {
		posY, posX = boxMap.Move(instruction, posY, posX)
	}

	return boxMap.calculateSumCoordinateValues()
}

// Move takes the value at the specified position and moves it (if possible) according
// to the specified instruction. Move returns the new y,x position after the move
// occurs
func (boxMap *BoxMap) Move(instruction rune, positionY, positionX int) (int, int) {
	b := *boxMap

	sizeY := len(b)
	sizeX := len(b[0])

	y, x := positionY, positionX
	newPosY, newPosX := positionY, positionX

	var move bool

	switch instruction {
	case '^':
		for y = positionY - 1; y > 0; y-- {
			if b[y][x] == '.' {
				move = true
				break
			}

			if b[y][x] == '#' {
				break
			}
		}

		if move {
			// there's an open spot
			if positionY-y > 1 {
				// we hit boxes, push the box into the open spot
				b[y][x] = b[positionY-1][positionX]
			}

			b[positionY-1][positionX] = b[positionY][positionX]
			b[positionY][positionX] = '.'

			newPosY = positionY - 1
			newPosX = positionX
		}
	case '>':
		for x = positionX + 1; x < sizeX-1; x++ {
			if b[y][x] == '.' {
				move = true
				break
			}

			if b[y][x] == '#' {
				break
			}
		}

		if move {
			// there's an open spot
			if x-positionX > 1 {
				// we hit boxes, push the box into the open spot
				b[y][x] = b[positionY][positionX+1]
			}

			b[positionY][positionX+1] = b[positionY][positionX]
			b[positionY][positionX] = '.'

			newPosY = positionY
			newPosX = positionX + 1
		}
	case 'v':
		for y = positionY + 1; y < sizeY-1; y++ {
			if b[y][x] == '.' {
				move = true
				break
			}

			if b[y][x] == '#' {
				break
			}
		}

		if move {
			// there's an open spot
			if y-positionY > 1 {
				// we hit boxes, push the box into the open spot
				b[y][x] = b[positionY+1][positionX]
			}

			b[positionY+1][positionX] = b[positionY][positionX]
			b[positionY][positionX] = '.'

			newPosY = positionY + 1
			newPosX = positionX
		}
	case '<':
		for x = positionX - 1; x > 0; x-- {
			if b[y][x] == '.' {
				move = true
				break
			}

			if b[y][x] == '#' {
				break
			}
		}

		if move {
			// there's an open spot
			if positionX-x > 1 {
				// we hit boxes, push the box into the open spot
				b[y][x] = b[positionY][positionX-1]
			}

			b[positionY][positionX-1] = b[positionY][positionX]
			b[positionY][positionX] = '.'

			newPosY = positionY
			newPosX = positionX - 1
		}
	}

	return newPosY, newPosX
}

// Find returns the position of the specified value (the first instance found)
func (boxMap *BoxMap) Find(val rune) (int, int) {
	b := *boxMap

	sizeY := len(b)
	if sizeY <= 0 {
		return -1, -1
	}

	sizeX := len(b[0])
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if b[y][x] == val {
				return y, x
			}
		}
	}

	return -1, -1
}

// calculateSumCoordinateValues finds all instances of a O value and, using
// the formula provided in the assignment, calculates the "coordinate value" of each.
// A coordinate value is y * 100 + x (with y,x being the grid location -y row, x column)
func (boxMap *BoxMap) calculateSumCoordinateValues() int {
	b := *boxMap
	sizeY := len(b)
	if sizeY <= 0 {
		return 0
	}

	sizeX := len(b[0])

	sumCoordinateValues := 0

	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if b[y][x] == 'O' {
				sumCoordinateValues += 100*y + x
			}
		}
	}

	return sumCoordinateValues
}

func (boxMap *BoxMap) Print() {
	b := *boxMap
	for y := 0; y < len(b); y++ {
		fmt.Printf("%s\n", string(b[y]))
	}
}

// Part2
func (d *Day15) Part2() int {
	return 0
}

// parseInput converts the input into a BoxMap and set of Instructions
func (d *Day15) parseInput(input []string) (BoxMap, Instructions) {
	var boxMap BoxMap
	var instructions string

	// Flag to identify whether we are processing the instructions part
	isInstructions := false

	for _, line := range input {
		// If an empty line is encountered, switch to instructions parsing
		if line == "" {
			isInstructions = true
			continue
		}

		if isInstructions {
			// Append instructions, ignoring newlines
			instructions += line
		} else {
			// Convert the line to a slice of runes and append to the BoxMap
			boxMap = append(boxMap, []rune(line))
		}
	}

	return boxMap, Instructions(instructions)
}
