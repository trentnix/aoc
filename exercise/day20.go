// day20.go is the implementation for the twentieth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day20 represents the data necessary to process the Exercise
	Day20 struct {
		name string
		file string
	}
)

// GetName returns the name of the Day 20 exercise
func (d *Day20) GetName() string {
	return d.name
}

// Run executes the solution for Day 20 by retrieving the default file contents and uses that data
func (d *Day20) Run(w io.Writer) {
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

// RunFromInput executs the Day 20 solution using the provided input data
func (d *Day20) RunFromInput(w io.Writer, input []string) {
	raceTrack := d.parseInput(input)

	// part 1
	numberOfCheatsThatSave100 := d.Part1(raceTrack, 100)
	w.Write([]byte(fmt.Sprintf("Day 20 - Part 1 - The number of cheats that save 100 is %d.\n", numberOfCheatsThatSave100)))
}

type RaceTrackPosition struct {
	Y, X int
}

// Part1 determines the single path of the race track provided that works, then goes
// through the path and calculates the distance saved for every wall removed, providing
// a sum of the removed walls that provide benefits exceeding the specified threshold
func (d *Day20) Part1(raceTrack Maze, threshold int) int {
	start := raceTrack.findLocation('S')
	end := raceTrack.findLocation('E')

	path := d.GetMazePath(raceTrack, start, end)
	// dictionary of positions for quick lookups
	positionDict := make(map[RaceTrackPosition]int)

	originalPathLength := len(path)
	for i := 0; i < originalPathLength; i++ {
		position := RaceTrackPosition{Y: path[i].Y, X: path[i].X}
		positionDict[position] = path[i].pointCost
	}

	// cheats is a map of positions where the specified key is the distance saved
	cheats := make(map[int][]RaceTrackPosition)

	for i := 0; i < originalPathLength; i++ {
		for _, direction := range directionDeltas {
			// position where we should look for a wall
			wallPositionY := path[i].Y + direction.dy
			wallPositionX := path[i].X + direction.dx
			// position where we should look for a continuation of the path
			pathCheckY := path[i].Y + (direction.dy * 2)
			pathCheckX := path[i].X + (direction.dx * 2)

			if raceTrack[wallPositionY][wallPositionX].val == '#' {
				// there is an adjacent wall

				positionToCheck := RaceTrackPosition{Y: pathCheckY, X: pathCheckX}
				if positionDict[positionToCheck] > 0 {
					// new cheat found, calculate the distance
					distanceBetween := positionDict[positionToCheck] - path[i].pointCost
					// we have to travel to the cheat location and to the continuation point, so remove 2 from the
					// distance between the points to calculate the distance saved
					distanceSaved := distanceBetween - 2
					if distanceSaved > 0 {
						// add the new position to the cheats slice at the specified distance saved
						cheats[distanceSaved] = append(cheats[distanceSaved], positionToCheck)
					}
				}
			}
		}
	}

	sumCheatsMeetingThresholdForSavings := 0
	for saved, points := range cheats {
		if saved >= threshold {
			// the distance saved exceeds or equals the threshold, add the number of cheats to the running total
			sumCheatsMeetingThresholdForSavings += len(points)
		}
	}

	return sumCheatsMeetingThresholdForSavings
}

// GetMazePath calculates the path through the specified maze (as long as there is a single, valid path)
// and returns a slice of the points traveled (and their relative distance from the provided start point)
func (d *Day20) GetMazePath(raceTrack Maze, start MazePoint, end MazePoint) []MazePoint {
	currentLocation := start

	// use RaceTrackPosition since MazePoint includes a pointCost that might cause issues doing lookups
	startPosition := RaceTrackPosition{Y: start.Y, X: start.X}

	// keeps track of the positions visited
	visited := make(map[RaceTrackPosition]bool)
	visited[startPosition] = true
	// builds the path that works
	var path []MazePoint
	path = append(path, start)

	distance := 0
	for currentLocation.X != end.X || currentLocation.Y != end.Y {
		// the next position will be 1 position further from the start
		distance++

		for _, direction := range directionDeltas {
			newY, newX := currentLocation.Y+direction.dy, currentLocation.X+direction.dx
			newPosition := RaceTrackPosition{Y: newY, X: newX}

			// check whether we can move in each direction to an open, non-visited point on the track
			if raceTrack[newY][newX].val != '#' && !visited[newPosition] {
				visited[newPosition] = true

				newLocation := MazePoint{Y: newY, X: newX, pointCost: distance}
				path = append(path, newLocation)

				currentLocation = newLocation
				break
			}
		}
	}

	return path
}

// Part2
func (d *Day20) Part2() int {
	return 0
}

// parseInput
func (d *Day20) parseInput(input []string) Maze {
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
