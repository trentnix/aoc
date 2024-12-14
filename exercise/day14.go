// day14.go is the implementation for the fourteenth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day14 represents the data necessary to process the Exercise
	Day14 struct {
		name string
		file string
	}

	Robot struct {
		x, y                 int
		velocityX, velocityY int
	}
)

// GetName returns the name of the Day 14 exercise
func (d *Day14) GetName() string {
	return d.name
}

// Run executes the solution for Day 14 by retrieving the default file contents and uses that data
func (d *Day14) Run(w io.Writer) {
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

// RunFromInput executs the Day 14 solution using the provided input data
func (d *Day14) RunFromInput(w io.Writer, input []string) {
	robots := d.parseInput(input)

	var seconds, gridX, gridY int

	// part 1
	seconds = 100
	gridX = 101
	gridY = 103
	safetyFactor := d.Part1(robots, seconds, gridX, gridY)
	w.Write([]byte(fmt.Sprintf("Day 14 - Part 1 - The safety factor after %d seconds for a %d by %d grid is %d.\n", seconds, gridX, gridY, safetyFactor)))

	// Part2
	robots = d.parseInput(input)
	secondsToTree := d.Part2(robots, gridX, gridY)
	w.Write([]byte(fmt.Sprintf("Day 14 - Part 2 - The tree is visible after %d seconds.\n", secondsToTree)))
}

// Part1 determines where the robots will be after the specified number of seconds and calculates
// their safety factor
func (d *Day14) Part1(robots []Robot, seconds, gridX, gridY int) int {
	for i := 0; i < len(robots); i++ {
		moveX := (robots[i].velocityX * seconds) % gridX
		moveY := (robots[i].velocityY * seconds) % gridY

		newX := robots[i].x + moveX
		newY := robots[i].y + moveY

		if newX < 0 {
			newX = newX + gridX
		} else if newX >= gridX {
			newX = newX % gridX
		}

		if newY < 0 {
			newY = newY + gridY
		} else if newY >= gridY {
			newY = newY % gridY
		}

		robots[i].x = newX
		robots[i].y = newY
	}

	robotMap := d.robotsToGrid(robots, gridX, gridY)

	middleY := gridY / 2
	middleX := gridX / 2

	quadrants := [4][4]int{
		// quadrant 1 (Top-left)
		{0, middleX - 1, 0, middleY - 1},
		// quadrant 2 (Top-right)
		{middleX + 1, gridX - 1, 0, middleY - 1},
		// quadrant 3 (Bottom-left)
		{0, middleX - 1, middleY + 1, gridY - 1},
		// quadrant 4 (Bottom-right)
		{middleX + 1, gridX - 1, middleY + 1, gridY - 1},
	}

	safetyFactor := 1
	for _, q := range quadrants {
		safetyFactor *= d.countRobots(robotMap, q[0], q[1], q[2], q[3])
	}

	return safetyFactor
}

// moveRobots moves the robots to a new position after 1 second
func (d *Day14) moveRobots(robots []Robot, sizeX, sizeY int) {
	for i := 0; i < len(robots); i++ {
		newX := robots[i].x + robots[i].velocityX
		newY := robots[i].y + robots[i].velocityY

		if newX < 0 {
			newX = newX + sizeX
		} else if newX >= sizeX {
			newX = newX % sizeX
		}

		if newY < 0 {
			newY = newY + sizeY
		} else if newY >= sizeY {
			newY = newY % sizeY
		}

		robots[i].x = newX
		robots[i].y = newY
	}
}

// countRobots counts the robots in the specified robots grid using the specified
// boundaries
func (d *Day14) countRobots(robots [][]int, startX, endX int, startY, endY int) int {
	if endY < startY || endX < startX {
		return -1
	}

	robotsInSpecifiedGrid := 0
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			robotsInSpecifiedGrid += robots[y][x]
		}
	}

	return robotsInSpecifiedGrid
}

// robotsToGrid converts a slice of Robots into a [][]int grid that shows the count
// of the robots at each grid location
func (d *Day14) robotsToGrid(robots []Robot, sizeX, sizeY int) [][]int {
	grid := make([][]int, sizeY) // Create a slice for the rows
	for i := range grid {
		grid[i] = make([]int, sizeX) // Allocate each row with sizeX columns
	}

	for _, robot := range robots {
		x := robot.x
		y := robot.y

		grid[y][x]++
	}

	return grid
}

// Part2 tries to find when the robots are arranged into a Christmas tree,
// which should correlate to when all robots are on a distinct location
func (d *Day14) Part2(robots []Robot, gridX, gridY int) int {
	seconds := 0
	for {
		seconds++
		d.moveRobots(robots, gridX, gridY)
		robotMap := d.robotsToGrid(robots, gridX, gridY)

		overlap := false
		for _, row := range robotMap {
			for _, element := range row {
				if element > 1 {
					overlap = true
				}
			}
		}

		if !overlap {
			// fmt.Printf("seconds: %d\n", seconds)
			// d.printGrid(robotMap)
			// fmt.Printf("\n\n\n")
			break
		}
	}

	return seconds
}

// parseInput takes the specified input and returns a slice of Robot instances
func (d *Day14) parseInput(input []string) []Robot {
	var robots []Robot

	for _, line := range input {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue // Skip if the format is invalid
		}

		// position (p=...)
		posPart := strings.TrimPrefix(parts[0], "p=")
		posCoords := strings.Split(posPart, ",")
		if len(posCoords) != 2 {
			continue
		}
		x, _ := strconv.Atoi(posCoords[0])
		y, _ := strconv.Atoi(posCoords[1])

		// velocity (v=...)
		velPart := strings.TrimPrefix(parts[1], "v=")
		velCoords := strings.Split(velPart, ",")
		if len(velCoords) != 2 {
			continue
		}
		velocityX, _ := strconv.Atoi(velCoords[0])
		velocityY, _ := strconv.Atoi(velCoords[1])

		robot := Robot{x: x, y: y, velocityX: velocityX, velocityY: velocityY}
		robots = append(robots, robot)
	}

	return robots
}

func (d *Day14) printGrid(robotMap [][]int) {
	for _, row := range robotMap {
		for _, element := range row {
			if element == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", element)
			}
		}
		fmt.Printf("\n")
	}
}
