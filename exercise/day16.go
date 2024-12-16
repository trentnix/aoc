// day16.go is the implementation for the sixteenth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"math"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day16 represents the data necessary to process the Exercise
	Day16 struct {
		name string
		file string
	}

	ReindeerMaze [][]ReindeerMazeLocation

	ReindeerMazeLocation struct {
		val rune
	}

	ReindeerMazePoint struct {
		Y, X int
	}

	ReindeerMazePath struct {
		positions []ReindeerMazePoint
		cost      int
	}
)

const (
	north = iota
	east
	south
	west
)

var directionDeltas = []struct {
	dy, dx int
}{
	{-1, 0}, // north
	{0, 1},  // east
	{1, 0},  // south
	{0, -1}, // west
}

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
}

// Part1
func (d *Day16) Part1(maze ReindeerMaze) int {
	start := maze.findLocation('S')
	end := maze.findLocation('E')
	startDirection := east

	// Find the cheapest path
	cheapestPath := findCheapestPath(maze, start, end, startDirection)

	return cheapestPath.cost
}

// Part2
func (d *Day16) Part2() int {
	return 0
}

// parseInput converts the input into a ReindeerMaze
func (d *Day16) parseInput(input []string) ReindeerMaze {
	maze := make([][]ReindeerMazeLocation, len(input))

	for i, line := range input {
		row := make([]ReindeerMazeLocation, len(line))

		for j, char := range line {
			row[j] = ReindeerMazeLocation{
				val: char,
			}
		}

		maze[i] = row
	}

	return maze
}

// findLocation will find the y,x location of the specified val in the specified
// ReindeerMaze
func (reindeerMaze *ReindeerMaze) findLocation(val rune) ReindeerMazePoint {
	r := *reindeerMaze
	for y := 0; y < len(r); y++ {
		for x := 0; x < len(r[0]); x++ {
			if r[y][x].val == val {
				return ReindeerMazePoint{Y: y, X: x}
			}
		}
	}

	return ReindeerMazePoint{Y: -1, X: -1}
}

// Move calculates the next position and direction based on the move type
func Move(current ReindeerMazePoint, direction int, moveType string, maze ReindeerMaze) (ReindeerMazePoint, int, bool) {
	var newDirection int
	switch moveType {
	case "forward":
		newDirection = direction
	case "left":
		newDirection = (direction + 3) % 4
	case "right":
		newDirection = (direction + 1) % 4
	default:
		return current, direction, false // Invalid move
	}

	next := ReindeerMazePoint{
		X: current.X + directionDeltas[newDirection].dx,
		Y: current.Y + directionDeltas[newDirection].dy,
	}

	// Check bounds and walls
	if next.X < 0 || next.X >= len(maze) || next.Y < 0 || next.Y >= len(maze[0]) || maze[next.Y][next.X].val == '#' {
		return current, direction, false // Invalid move
	}

	return next, newDirection, true // Valid move
}

// Recursive function to explore all possible paths.
func findPaths(
	maze ReindeerMaze,
	start ReindeerMazePoint,
	end ReindeerMazePoint,
	direction int,
	visited map[ReindeerMazePoint]map[int]bool,
	path []ReindeerMazePoint,
	turns int,
	forwardMoves int,
	paths *[]ReindeerMazePath,
	bestCost *int,
) {
	// Calculate the current cost of the path
	currentCost := (turns * 1000) + forwardMoves

	// Prune paths with cost exceeding the best cost
	if currentCost >= *bestCost {
		return
	}

	// Check if we reached the end
	if start == end {
		// Update bestCost if this path is better
		if currentCost < *bestCost {
			*bestCost = currentCost
		}

		// Add the current path to the list of valid paths
		*paths = append(*paths, ReindeerMazePath{
			positions: append([]ReindeerMazePoint(nil), path...),
			cost:      currentCost,
		})
		return
	}

	// Mark this position and direction as visited
	if visited[start] == nil {
		visited[start] = make(map[int]bool)
	}
	visited[start][direction] = true

	// Try to move forward
	if next, nextDir, valid := Move(start, direction, "forward", maze); valid {
		if !visited[next][nextDir] {
			findPaths(maze, next, end, nextDir, visited, append(path, next), turns, forwardMoves+1, paths, bestCost)
		}
	}

	// Try to turn left and move forward
	if next, nextDir, valid := Move(start, direction, "left", maze); valid {
		if !visited[next][nextDir] {
			findPaths(maze, next, end, nextDir, visited, append(path, next), turns+1, forwardMoves+1, paths, bestCost)
		}
	}

	// Try to turn right and move forward
	if next, nextDir, valid := Move(start, direction, "right", maze); valid {
		if !visited[next][nextDir] {
			findPaths(maze, next, end, nextDir, visited, append(path, next), turns+1, forwardMoves+1, paths, bestCost)
		}
	}

	// Unmark this position and direction as visited (backtracking)
	delete(visited[start], direction)
	if len(visited[start]) == 0 {
		delete(visited, start)
	}
}

// Finds the cheapest path among all possible paths.
func findCheapestPath(maze ReindeerMaze, start ReindeerMazePoint, end ReindeerMazePoint, startDirection int) ReindeerMazePath {
	paths := []ReindeerMazePath{}
	visited := make(map[ReindeerMazePoint]map[int]bool)
	bestCost := math.MaxInt

	// Add the starting point to the path
	initialPath := []ReindeerMazePoint{start}

	// Explore all paths
	findPaths(maze, start, end, startDirection, visited, initialPath, 0, 0, &paths, &bestCost)

	// Find the path with the minimum cost
	var cheapestPath ReindeerMazePath
	for _, p := range paths {
		if p.cost == bestCost {
			cheapestPath = p
			break
		}
	}

	return cheapestPath
}
