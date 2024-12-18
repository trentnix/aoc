// day18.go is the implementation for the eighteenth day of the Advent of Code 2024
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
	// Day18 represents the data necessary to process the Exercise
	Day18 struct {
		name string
		file string
	}

	FallingBlocks []MazePoint
)

// GetName returns the name of the Day 18 exercise
func (d *Day18) GetName() string {
	return d.name
}

// Run executes the solution for Day 18 by retrieving the default file contents and uses that data
func (d *Day18) Run(w io.Writer) {
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

// RunFromInput executs the Day 18 solution using the provided input data
func (d *Day18) RunFromInput(w io.Writer, input []string) {
	startStep := 1024
	gridSize := 71
	fallingBlocks := d.parseInput(input)

	// part 1
	steps := d.Part1(fallingBlocks, gridSize, startStep)
	w.Write([]byte(fmt.Sprintf("Day 18 - Part 1 - The shortest path distance is: %d\n", steps)))

	y, x := d.Part2(fallingBlocks, gridSize, startStep)
	w.Write([]byte(fmt.Sprintf("Day 18 - Part 2 - The coordinates of the block that breaks the map is: %d,%d\n", y, x)))
}

// Part1
func (d *Day18) Part1(fallingBlocks FallingBlocks, gridSize int, startStep int) int {
	if startStep >= len(fallingBlocks) {
		log.Fatalf("invalid input: startStep is invalid")
	}

	// build the grid
	memoryMaze := Maze(make([][]MazeLocation, gridSize))
	for i := 0; i < gridSize; i++ {
		memoryMaze[i] = make([]MazeLocation, gridSize)
		for j := 0; j < gridSize; j++ {
			// set default value for each MazeLocation
			memoryMaze[i][j] = MazeLocation{val: '.'} // '.' represents an empty cell
		}
	}

	for i := 0; i < startStep; i++ {
		// add the falling blocks to the grid at each fallingBlocks location
		blockLocation := fallingBlocks[i]
		memoryMaze[blockLocation.Y][blockLocation.X].val = '#'
	}

	memoryMazeGraph := buildMazeGraph(memoryMaze)

	start := MazePoint{Y: 0, X: 0}
	end := MazePoint{Y: gridSize - 1, X: gridSize - 1}

	cheapestPath := findLowestCostMazePath(memoryMazeGraph, start, end, south, calculateMemoryMazeCost)

	return cheapestPath
}

// Part2
func (d *Day18) Part2(fallingBlocks FallingBlocks, gridSize int, startStep int) (y int, x int) {
	if startStep >= len(fallingBlocks) {
		log.Fatalf("invalid input: startStep is invalid")
	}

	// build the grid
	memoryMaze := Maze(make([][]MazeLocation, gridSize))
	for i := 0; i < gridSize; i++ {
		memoryMaze[i] = make([]MazeLocation, gridSize)
		for j := 0; j < gridSize; j++ {
			// set default value for each MazeLocation
			memoryMaze[i][j] = MazeLocation{val: '.'} // '.' represents an empty cell
		}
	}

	for i := 0; i < startStep; i++ {
		// add the falling blocks to the grid at each fallingBlocks location
		blockLocation := fallingBlocks[i]
		memoryMaze[blockLocation.Y][blockLocation.X].val = '#'
	}

	start := MazePoint{Y: 0, X: 0}
	end := MazePoint{Y: gridSize - 1, X: gridSize - 1}

	remainingBlocks := fallingBlocks[startStep:]

	for i := 0; i < len(remainingBlocks); i++ {
		// set the maze location
		block := remainingBlocks[i]
		memoryMaze[block.Y][block.X].val = '#'

		// rebuild the graph and find a path
		memoryMazeGraph := buildMazeGraph(memoryMaze)
		if -1 == findLowestCostMazePath(memoryMazeGraph, start, end, south, calculateMemoryMazeCost) {
			// not path found - we have the block
			y = block.Y
			x = block.X
			break
		}
	}

	return y, x
}

// parseInput takes the specified input and converts it into a FallingBlocks structure
func (d *Day18) parseInput(input []string) FallingBlocks {
	var fallingBlocks FallingBlocks

	for _, line := range input {
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			log.Fatalf("the input is invalid")
		}

		// convert the Y and X values
		y, err1 := strconv.Atoi(coords[0])
		x, err2 := strconv.Atoi(coords[1])

		if err1 != nil || err2 != nil {
			log.Fatalf("the input is invalid")
		}

		// add the MazePoint to the FallingBlocks slice
		fallingBlocks = append(fallingBlocks, MazePoint{Y: y, X: x})
	}

	return fallingBlocks
}

// calculateMemoryMazeCost determines the cost of a particular state in a Maze instance
// for the MemoryMaze assignment
func calculateMemoryMazeCost(s *State, e *MazeEdge) int {
	return s.cost + e.cost
}
