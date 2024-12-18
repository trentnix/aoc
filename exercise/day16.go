// day16.go is the implementation for the sixteenth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day16 represents the data necessary to process the Exercise
	Day16 struct {
		name string
		file string
	}
)

// GetName returns the name of the Day 16 exercise
func (d *Day16) GetName() string {
	return d.name
}

// Run executes the solution for Day 16 by retrieving the default file contents and uses that data
func (d *Day16) Run(w io.Writer) {
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

// RunFromInput executs the Day 16 solution using the provided input data
func (d *Day16) RunFromInput(w io.Writer, input []string) {
	maze := d.parseInput(input)

	// part 1
	cheapestPathCost := d.Part1(maze)
	w.Write([]byte(fmt.Sprintf("Day 16 - Part 1 - The cheapest path through the maze costs %d.\n", cheapestPathCost)))

	// part 1
	maze = d.parseInput(input)
	visitedNodes := d.Part2(maze)
	w.Write([]byte(fmt.Sprintf("Day 16 - Part 2 - The number of nodes visited on the cheapest path(s) is %d.\n", visitedNodes)))
}

// Part1 traverses the maze and finds the shortest path cost from the start to
// the end positions
func (d *Day16) Part1(maze Maze) int {
	start := maze.findLocation('S')
	end := maze.findLocation('E')
	startDirection := east

	reindeerMazeGraph := buildMazeGraph(maze)

	// Find the cheapest path
	cheapestPath := findLowestCostMazePath(reindeerMazeGraph, start, end, startDirection, calculateReindeerMazeCost)

	return cheapestPath
}

// Part2 finds the number of nodes visited along every path that happens to share
// the lowest cost
func (d *Day16) Part2(maze Maze) int {
	start := maze.findLocation('S')
	end := maze.findLocation('E')
	startDirection := east

	reindeerMazeGraph := buildMazeGraph(maze)

	_, allPaths := findAllMinimumMazePaths(reindeerMazeGraph, start, end, startDirection, calculateReindeerMazeCost)
	// allPaths returns nodes, but not every position. We need to expand it to have every
	// position in the maze, not just the graph nodes
	expandedPaths := expandAllMazePaths(allPaths, reindeerMazeGraph)

	visitedPositions := make(map[MazePoint]bool)
	for _, p := range expandedPaths {
		for _, pos := range p {
			// we visited this position on the specified path, set it to true
			visitedPositions[pos] = true
		}
	}

	// return the number of positions that were visited
	return len(visitedPositions)
}

// parseInput converts the input into a Maze
func (d *Day16) parseInput(input []string) Maze {
	maze := make([][]MazeLocation, len(input))

	for i, line := range input {
		row := make([]MazeLocation, len(line))

		for j, char := range line {
			row[j] = MazeLocation{
				val: char,
			}
		}

		maze[i] = row
	}

	return maze
}

// findLocation will find the y,x location of the specified val in the specified
// Maze
func (reindeerMaze *Maze) findLocation(val rune) MazePoint {
	r := *reindeerMaze
	for y := 0; y < len(r); y++ {
		for x := 0; x < len(r[0]); x++ {
			if r[y][x].val == val {
				return MazePoint{Y: y, X: x}
			}
		}
	}

	return MazePoint{Y: -1, X: -1}
}

func calculateReindeerMazeCost(s *State, e *MazeEdge) int {
	turnCost := 0
	if e.direction != s.direction {
		// this is a turn
		turnCost = 1000
	}

	return s.cost + turnCost + e.cost
}
