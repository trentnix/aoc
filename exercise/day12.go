// day12.go is the implementation for the twelth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day12 represents the data necessary to process the Exercise
	Day12 struct {
		name string
		file string
	}

	Garden           [][]rune
	GardenCoordinate struct {
		x, y int
	}
	GardenSections       map[GardenCoordinate][]GardenCoordinate
	GardenSectionDetails struct {
		plant     rune
		id        GardenCoordinate
		area      int
		perimeter int
	}
)

// GetName returns the name of the Day 12 exercise
func (d *Day12) GetName() string {
	return d.name
}

// Run executes the solution for Day 12 by retrieving the default file contents and uses that data
func (d *Day12) Run(w io.Writer) {
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

// RunFromInput executs the Day 12 solution using the provided input data
func (d *Day12) RunFromInput(w io.Writer, input []string) {
	garden := d.parseInput(input)

	// part 1
	totalPrice := d.Part1(garden)
	w.Write([]byte(fmt.Sprintf("Day 12 - Part 1 - The total price of fencing all regions is %d.\n", totalPrice)))
}

// Part1 calculates the price of the fence by finding each Garden section, calculating
// the area, calculating the perimeter, multiplying the area * perimeter, and adding the
// product. The sum of the products for each area is the total price of the fencing,
// per the assignment.
func (d *Day12) Part1(garden Garden) int {
	sections := d.extractSections(garden)
	var sectionDetails []GardenSectionDetails

	for coordinate, section := range sections {
		details := GardenSectionDetails{
			id:    coordinate,
			plant: garden[coordinate.y][coordinate.x],
			area:  len(section),
		}

		details.perimeter = d.calculatePerimeter(garden, section)

		sectionDetails = append(sectionDetails, details)
	}

	price := 0
	for _, details := range sectionDetails {
		price += details.area * details.perimeter
	}

	return price
}

// Part2
func (d *Day12) Part2() int {
	return 0
}

// parseInput
func (d *Day12) parseInput(input []string) Garden {
	var garden Garden
	for _, row := range input {
		garden = append(garden, []rune(row))
	}
	return garden
}

// Helper function to check if a coordinate is within bounds
func (d *Day12) isInBounds(garden Garden, x, y int) bool {
	return x >= 0 && x < len(garden) && y >= 0 && y < len(garden[0])
}

// findSection finds sections using flood-fill
func (d *Day12) findSection(garden Garden, visited map[GardenCoordinate]bool, x, y int) []GardenCoordinate {
	runeValue := garden[x][y]
	section := []GardenCoordinate{}
	stack := []GardenCoordinate{{x, y}}

	// Perform DFS to collect all connected coordinates
	for len(stack) > 0 {
		coord := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// skip already visited coordinates
		if visited[coord] {
			continue
		}

		visited[coord] = true
		section = append(section, coord)

		// check neighbors
		directions := []GardenCoordinate{
			{0, 1},  // right
			{1, 0},  // down
			{0, -1}, // left
			{-1, 0}, // up
		}

		for _, dir := range directions {
			nx, ny := coord.x+dir.x, coord.y+dir.y
			if d.isInBounds(garden, nx, ny) && !visited[GardenCoordinate{nx, ny}] && garden[nx][ny] == runeValue {
				stack = append(stack, GardenCoordinate{nx, ny})
			}
		}
	}

	return section
}

// extractSections creates an instance of GardenSections from the specified Garden
func (d *Day12) extractSections(garden Garden) GardenSections {
	sections := make(GardenSections)
	visited := make(map[GardenCoordinate]bool)

	for x := 0; x < len(garden); x++ {
		for y := 0; y < len(garden[0]); y++ {
			coord := GardenCoordinate{x, y}

			// If not visited, it's a new section
			if !visited[coord] {
				section := d.findSection(garden, visited, x, y)
				sections[coord] = section
			}
		}
	}

	return sections
}

// calculatePerimeter calculates the perimeter of a given Garden section using
func (d *Day12) calculatePerimeter(garden Garden, section []GardenCoordinate) int {
	// set used to check if a coordinate belongs to the section
	sectionSet := make(map[GardenCoordinate]bool)
	for _, coord := range section {
		sectionSet[coord] = true
	}

	perimeter := 0

	directions := []GardenCoordinate{
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
		{-1, 0}, // up
	}

	for _, coord := range section {
		for _, dir := range directions {
			neighbor := GardenCoordinate{coord.x + dir.x, coord.y + dir.y}
			if !d.isInBounds(garden, neighbor.x, neighbor.y) || !sectionSet[neighbor] {
				// neighbor is not out of bounds and is part of the section
				perimeter++
			}
		}
	}

	return perimeter
}
