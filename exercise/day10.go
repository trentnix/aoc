// day10.go is the implementation for the tenth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"strconv"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day10 represents the data necessary to process the Exercise
	Day10 struct {
		name string
		file string
	}

	TopographicMap [][]int

	MapPosition struct {
		x, y, value int
	}
)

// GetName returns the name of the Day 10 exercise
func (d *Day10) GetName() string {
	return d.name
}

// Run executes the solution for Day 10 by retrieving the default file contents and uses that data
func (d *Day10) Run(w io.Writer) {
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

// RunFromInput executs the Day 10 solution using the provided input data
//
// note - part1 and part2 could be combined to only iterate over the trails a single time and then
// count the number of paths and unique paths to arrive at each solution. However, in the interest of separating
// the solutions more deliberately, we navigate the trails for each part
func (d *Day10) RunFromInput(w io.Writer, input []string) {
	topo := d.parseInput(input)

	// part 1
	sumTrailheadScores := d.Part1(topo)
	w.Write([]byte(fmt.Sprintf("Day 10 - Part 1 - The sum of the trailhead scores for the provided map is %d.\n", sumTrailheadScores)))

	// part 2
	sumTrailheadRatings := d.Part2(topo)
	w.Write([]byte(fmt.Sprintf("Day 10 - Part 2 - The sum of the trailhead ratings for the provided map is %d.\n", sumTrailheadRatings)))
}

// Part1 counts all of the unique trails that have the same origin and endpoint
func (d *Day10) Part1(topo TopographicMap) int {
	if topo == nil || len(topo) == 0 {
		fmt.Printf("The TopographicMap specified is invalid.")
		return -1
	}

	trailExit := 9
	sumTrails := 0

	trailheads := d.getTopographicMapPositions(topo, 0)
	for _, trailhead := range trailheads {
		trails := d.findNextPosition(topo, trailhead, trailExit)
		uniqueTrails := d.countUniqueTrails(trails)
		sumTrails += uniqueTrails
	}

	return sumTrails
}

// Part2 counts the sum of the unique trails that have their own unique path
func (d *Day10) Part2(topo TopographicMap) int {
	if topo == nil || len(topo) == 0 {
		fmt.Printf("The TopographicMap specified is invalid.")
		return -1
	}

	trailExit := 9
	sumTrails := 0

	trailheads := d.getTopographicMapPositions(topo, 0)
	for _, trailhead := range trailheads {
		trails := d.findNextPosition(topo, trailhead, trailExit)
		sumTrails += len(trails)
	}

	return sumTrails
}

// findNextPosition navigates the map (going up, down, left, or right) from the
// specified position until navigation is blocked (the next position is more than 1
// topographical value away) or the exitVal is reached
func (d *Day10) findNextPosition(topo TopographicMap, position MapPosition, exitVal int) []MapPosition {
	maxY := len(topo)
	maxX := len(topo[0])

	var lastPositions []MapPosition
	position.value = topo[position.y][position.x]

	adjacentPositions := d.findAdjacentPositions(position, maxY, maxX)
	for _, adjacentPosition := range adjacentPositions {
		currentValue := topo[adjacentPosition.y][adjacentPosition.x]
		adjacentPosition.value = currentValue

		// Check if this adjacent position advances by exactly 1
		if currentValue == position.value+1 {
			if currentValue >= exitVal {
				// append to results and continue
				lastPositions = append(lastPositions, adjacentPosition)
			} else {
				// Recurse to find all valid paths going forward
				lastPositions = append(lastPositions, d.findNextPosition(topo, adjacentPosition, exitVal)...)
			}
		}
	}

	return lastPositions
}

// countUniqueTrails counts the number of MapPosition in the specified array that have unique locations
func (d *Day10) countUniqueTrails(trails []MapPosition) int {
	uniquePositions := make(map[string]struct{})

	for _, trail := range trails {
		key := fmt.Sprintf("%d,%d", trail.x, trail.y)
		uniquePositions[key] = struct{}{}
	}

	return len(uniquePositions)
}

// findAdjacentPositions finds all of the possible positions that can be moved to from
// the specified position. maxY and maxX define the boundaries of the map. Movement is
// restricted to up, down, left, and right
func (d *Day10) findAdjacentPositions(position MapPosition, maxY int, maxX int) []MapPosition {
	// Define the possible directions for adjacency (including diagonals)
	directions := []struct{ dx, dy int }{
		{0, -1}, // Up
		{-1, 0}, // Left
		{1, 0},  // Right
		{0, 1},  // Down
	}

	var adjacentPositions []MapPosition

	for _, dir := range directions {
		// Compute the new position
		newX := position.x + dir.dx
		newY := position.y + dir.dy

		// Check if the new position is within bounds
		if newX >= 0 && newX < maxX && newY >= 0 && newY < maxY {
			// Add the new position with value -1
			adjacentPositions = append(adjacentPositions, MapPosition{
				x:     newX,
				y:     newY,
				value: -1, // Set the value to -1 as required
			})
		}
	}

	return adjacentPositions
}

// getTopographicMapPositions returns an array of MapPosition objects outlining all of the
// TopographicalMap coordinates where the specified value exists
func (d *Day10) getTopographicMapPositions(topo TopographicMap, value int) []MapPosition {
	numColumns := len(topo)
	if numColumns == 0 {
		return nil
	}

	numRows := len(topo[0])

	var positions []MapPosition
	for y := 0; y < numRows; y++ {
		for x := 0; x < numColumns; x++ {
			if topo[y][x] == value {
				positions = append(positions, MapPosition{x: x, y: y, value: value})
			}
		}
	}

	return positions
}

// parseInput parses the input array of strings into a TopoGraphicMap
func (d *Day10) parseInput(input []string) TopographicMap {
	numColumns := len(input)
	if numColumns == 0 {
		return nil
	}

	numRows := len(input[0])

	var err error

	var topo TopographicMap = make(TopographicMap, numRows)
	for y := 0; y < numRows; y++ {
		topo[y] = make([]int, numColumns)
		for x := 0; x < numColumns; x++ {
			topo[y][x], err = strconv.Atoi(string(input[y][x]))
			if err != nil {
				fmt.Printf("There was an error parsing converting the input character at [y:%d][x:%d] to its topographic value: %s: %v", y, x, string(input[y][x]), err)
				return nil
			}
		}
	}

	return topo
}
