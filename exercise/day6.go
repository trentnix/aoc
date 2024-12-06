// day6.go is the implementation for the TBD day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day6 represents the data necessary to process the Exercise
	Day6 struct {
		name string
		file string
	}

	Grid struct {
		coordinate [][]rune
	}
)

// GetName returns the name of the Day 6 exercise
func (d *Day6) GetName() string {
	return d.name
}

// Run executes the solution for Day 6 by retrieving the default file contents and uses that data
func (d *Day6) Run(w io.Writer) {
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

// RunFromInput executes the Day 6 solution using the provided input data
func (d *Day6) RunFromInput(w io.Writer, input []string) {
	g := d.parseInput(input)

	numberPositionsVisitedByGuard := d.Part1(g)

	// part 1
	w.Write([]byte(fmt.Sprintf("Day 6 - Part 1 - The number of coordinates visited by the guard is %d.\n", numberPositionsVisitedByGuard)))
}

// Part1 moves the guard through the map (grid) and counts how many coordinates
// the guard covers
func (d *Day6) Part1(g *Grid) int {
	guardPositionX, guardPositionY, direction := d.findGuardPositionAndDirection(g)
	err := d.traverseGrid(g, guardPositionX, guardPositionY, direction)
	if err != nil {
		fmt.Printf("there was an error traversing the grid")
		return -1
	}

	return d.countVisited(g)
}

// countVisited counts the number of coordinates whose value is 'X', indicating they
// were visited by the guard
func (d *Day6) countVisited(g *Grid) int {
	numVisited := 0
	for _, row := range g.coordinate {
		for _, c := range row {
			if c == 'X' {
				numVisited++
			}
		}
	}

	return numVisited
}

// traverseGrid takes a guard, given the start position, and moves them through the grid
// according to the rules outlined in the assignment. Namely:
//
// If there is something directly in front of you, turn right 90 degrees.
// Otherwise, take a step forward.
func (d *Day6) traverseGrid(g *Grid, startX int, startY int, direction string) error {
	sizeX := len(g.coordinate[0])
	sizeY := len(g.coordinate)

	if sizeX <= 0 || sizeY <= 0 {
		return fmt.Errorf("the specified grid is invalid")
	}

	if d.isInvalidPosition(startX, startY, sizeX, sizeY) {
		return fmt.Errorf("the specified initial coordinate position is invalid")
	}

	// set the initial position as visited
	g.coordinate[startY][startX] = 'X'

	currentX, currentY := startX, startY
	var nextX, nextY int

	for {
		nextX, nextY = currentX, currentY

		// update the next position
		switch direction {
		case "north":
			nextY--
		case "east":
			nextX++
		case "south":
			nextY++
		case "west":
			nextX--
		default:
			return fmt.Errorf("invalid direction value: %s", direction)
		}

		if d.isInvalidPosition(nextX, nextY, sizeX, sizeY) {
			g.coordinate[currentY][currentX] = 'X'
			break
		}

		// figure out if blocked
		if g.coordinate[nextY][nextX] == '#' {
			// need to turn
			switch direction {
			case "north":
				direction = "east"
			case "east":
				direction = "south"
			case "south":
				direction = "west"
			case "west":
				direction = "north"
			default:
				return fmt.Errorf("invalid direction value: %s", direction)
			}
		} else {
			// set the current position to visited and move the current position
			g.coordinate[currentY][currentX] = 'X'
			currentX, currentY = nextX, nextY
		}
	}

	return nil
}

// isInvalidPosition determines whether the x,y is within the grid whose size is
// specified by sizeX, sizeY
func (d *Day6) isInvalidPosition(x, y, sizeX, sizeY int) bool {
	if x >= sizeX || x < 0 {
		return true
	}

	if y >= sizeY || y < 0 {
		return true
	}

	return false
}

// findGuardPosition returns the position of the guard assuming the upper-leftmost position
// on the grid is 0, 0 and both x and y increase as you move down and right on the grid
func (d *Day6) findGuardPositionAndDirection(g *Grid) (int, int, string) {
	for y, row := range g.coordinate {
		for x, c := range row {
			if c == '^' {
				return x, y, "north"
			}

			if c == '>' {
				return x, y, "east"
			}

			if c == '<' {
				return x, y, "west"
			}

			if c == 'v' {
				return x, y, "south"
			}
		}
	}

	return -1, -1, ""
}

// Part2
func (d *Day6) Part2() int {
	return 0
}

// parseInput takes the string array input and converts it into a Grid
func (d *Day6) parseInput(input []string) *Grid {
	var grid Grid

	// Convert each string in the input to a slice of runes
	for _, line := range input {
		row := []rune(line) // Convert string to rune slice
		grid.coordinate = append(grid.coordinate, row)
	}

	return &grid
}

// printGrid provides a pretty-print of the grid to stdout
func (d *Day6) printGrid(g *Grid) {
	fmt.Println("Grid:")
	for _, row := range g.coordinate {
		for _, cell := range row {
			fmt.Printf("%c ", cell) // Print each rune with a space
		}
		fmt.Println() // Newline after each row
	}

	fmt.Println()
}
