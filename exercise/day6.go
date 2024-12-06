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
		position [][]rune
	}

	Coordinate struct {
		x int
		y int
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

	// part 1
	numberPositionsVisitedByGuard := d.Part1(g.Copy())
	w.Write([]byte(fmt.Sprintf("Day 6 - Part 1 - The number of positions visited by the guard is %d.\n", numberPositionsVisitedByGuard)))

	// part 2
	numLoops := d.Part2(g.Copy())
	w.Write([]byte(fmt.Sprintf("Day 6 - Part 2 - The number of new blocks that result in a loop is %d.\n", numLoops)))
}

// Part1 moves the guard through the map (grid) and counts how many positions
// the guard covers
func (d *Day6) Part1(g *Grid) int {
	guardPositionX, guardPositionY, direction := d.findGuardPositionAndDirection(g)
	_, err := d.traverseGridLoop(g, guardPositionX, guardPositionY, direction)
	if err != nil {
		fmt.Printf("there was an error traversing the grid")
		return -1
	}

	return d.countVisited(g)
}

// countVisited counts the number of positions whose value is 'X', indicating they
// were visited by the guard
func (d *Day6) countVisited(g *Grid) int {
	numVisited := 0
	for _, row := range g.position {
		for _, c := range row {
			if c == 'X' {
				numVisited++
			}
		}
	}

	return numVisited
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
	for y, row := range g.position {
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

// Part2 adds an obstruction to each point in the grid and looks for scenarios where
// the guard loops due to the added obstruction
func (d *Day6) Part2(g *Grid) int {
	numObstructionsThatCauseLoops := 0
	guardPositionX, guardPositionY, direction := d.findGuardPositionAndDirection(g)

	for currentY, row := range g.position {
		for currentX, c := range row {
			if c != 'X' && c != '#' && !(currentX == guardPositionX && currentY == guardPositionY) {
				newGrid := g.Copy()

				// add a block to the current position
				newGrid.position[currentY][currentX] = 'O'

				isLooped, err := d.traverseGridLoop(newGrid, guardPositionX, guardPositionY, direction)
				if err != nil {
					fmt.Printf("there was an error traversing the grid")
					return -1
				}

				if isLooped {
					numObstructionsThatCauseLoops++
				}
			}
		}
	}

	return numObstructionsThatCauseLoops
}

// traverseGridLoop takes a guard, given the start position, and moves them through the grid
// according to the rules outlined in the assignment. Namely:
//
// If there is something directly in front of you, turn right 90 degrees.
// Otherwise, take a step forward.
//
// It returns 'true' when a loop is detected.
func (d *Day6) traverseGridLoop(g *Grid, startX int, startY int, direction string) (bool, error) {
	sizeX := len(g.position[0])
	sizeY := len(g.position)

	// set the initial position as visited
	g.position[startY][startX] = 'X'

	currentX, currentY := startX, startY
	var nextX, nextY int

	// create a slice to store coordinates of the grid to help detect for loops
	var loopCoordinates []Coordinate

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
			return false, fmt.Errorf("invalid direction value: %s", direction)
		}

		if d.isInvalidPosition(nextX, nextY, sizeX, sizeY) {
			g.position[currentY][currentX] = 'X'
			break
		}

		// figure out if blocked
		if g.position[nextY][nextX] == '#' || g.position[nextY][nextX] == 'O' {
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
				return false, fmt.Errorf("invalid direction value: %s", direction)
			}

			if g.position[nextY][nextX] == 'O' {
				if len(loopCoordinates)%4 == 0 && len(loopCoordinates) > 0 {
					if loopCoordinates[0].x == currentX && loopCoordinates[0].y == currentY {
						// we've made a loop back to the original element - report a looping event
						return true, nil
					}
				}

				loopCoordinates = append(loopCoordinates, Coordinate{x: currentX, y: currentY})
			} else {
				// reset the loopBlockCounter since we encountered a new block
				loopCoordinates = nil
				g.position[nextY][nextX] = 'O'
			}
		} else {
			// set the current position to visited and move the current position
			if g.position[currentY][currentX] == '.' {
				g.position[currentY][currentX] = 'X'
			}
			currentX, currentY = nextX, nextY
		}
	}

	return false, nil
}

// parseInput takes the string array input and converts it into a Grid
func (d *Day6) parseInput(input []string) *Grid {
	var grid Grid

	// Convert each string in the input to a slice of runes
	for _, line := range input {
		row := []rune(line) // Convert string to rune slice
		grid.position = append(grid.position, row)
	}

	return &grid
}

// Copy() returns a deep copy of the source Grid
func (g *Grid) Copy() *Grid {
	copyGrid := &Grid{
		position: make([][]rune, len(g.position)),
	}

	// Deep copy each row of the grid
	for i, row := range g.position {
		// Create a new slice for each row and copy the contents
		copyGrid.position[i] = make([]rune, len(row))
		copy(copyGrid.position[i], row)
	}

	return copyGrid
}

// printGrid provides a pretty-print of the grid to stdout
func (g *Grid) Print() {
	fmt.Println("Grid:")
	for _, row := range g.position {
		for _, cell := range row {
			fmt.Printf("%c ", cell) // Print each rune with a space
		}
		fmt.Println() // Newline after each row
	}

	fmt.Println()
}
