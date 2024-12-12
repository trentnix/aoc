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

	Garden     [][]GardenNode
	GardenNode struct {
		row, col              int
		left, up, right, down bool
		val                   rune
	}

	GardenSections       map[GardenNode][]GardenNode
	GardenSectionDetails struct {
		plant     rune
		id        GardenNode
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

	// part 1
	totalPrice = d.Part2(garden)
	w.Write([]byte(fmt.Sprintf("Day 12 - Part 2 - The total price of fencing all regions is %d.\n", totalPrice)))
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
			plant: garden[coordinate.row][coordinate.col].val,
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

// Part2 calculates the price of fence by calculating the area of a section
// and the number of straight fence runs in that section (which corresponds to
// the number of corners)
func (d *Day12) Part2(garden Garden) int {
	sections := d.extractSections(garden)
	var sectionDetails []GardenSectionDetails

	for coordinate, section := range sections {
		details := GardenSectionDetails{
			id:    coordinate,
			plant: garden[coordinate.row][coordinate.col].val,
			area:  len(section),
		}

		sectionWithDirectionalFlags := d.setDirectionalFlags(garden, section)
		details.perimeter = d.countCorners(sectionWithDirectionalFlags)

		sectionDetails = append(sectionDetails, details)
	}

	price := 0
	for _, details := range sectionDetails {
		price += details.area * details.perimeter
	}

	return price
}

func (d *Day12) parseInput(input []string) Garden {
	if len(input) == 0 {
		return nil
	}

	rows := len(input)
	cols := len(input[0])

	gardenNodes := make(Garden, rows)
	for i := 0; i < rows; i++ {
		gardenNodes[i] = make([]GardenNode, cols)
		for j, char := range input[i] {
			gardenNodes[i][j].val = char
			gardenNodes[i][j].row = i
			gardenNodes[i][j].col = j
		}
	}

	return gardenNodes
}

// Helper function to check if a coordinate is within bounds
func (d *Day12) isInBounds(garden Garden, row, col int) bool {
	return row >= 0 && row < len(garden) && col >= 0 && col < len(garden[0])
}

// findSection finds sections using flood-fill
func (d *Day12) findSection(garden Garden, visited map[GardenNode]bool, row, col int) []GardenNode {
	runeValue := garden[row][col].val
	section := []GardenNode{}
	stack := []GardenNode{{row: row, col: col}}

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
		directions := []GardenNode{
			{row: 0, col: 1},  // right
			{row: 1, col: 0},  // down
			{row: 0, col: -1}, // left
			{row: -1, col: 0}, // up
		}

		for _, dir := range directions {
			nrow, ncol := coord.row+dir.row, coord.col+dir.col
			if d.isInBounds(garden, nrow, ncol) && !visited[GardenNode{row: nrow, col: ncol}] && garden[nrow][ncol].val == runeValue {
				stack = append(stack, GardenNode{row: nrow, col: ncol})
			}
		}
	}

	return section
}

// extractSections creates an instance of GardenSections from the specified Garden
func (d *Day12) extractSections(garden Garden) GardenSections {
	sections := make(GardenSections)
	visited := make(map[GardenNode]bool)

	for row := 0; row < len(garden); row++ {
		for col := 0; col < len(garden[0]); col++ {
			coord := GardenNode{row: row, col: col}

			// If not visited, it's a new section
			if !visited[coord] {
				section := d.findSection(garden, visited, row, col)
				sections[coord] = section
			}
		}
	}

	return sections
}

// calculatePerimeter calculates the perimeter of a given Garden section using
func (d *Day12) calculatePerimeter(garden Garden, section []GardenNode) int {
	// set used to check if a coordinate belongs to the section
	sectionSet := make(map[GardenNode]bool)
	for _, coord := range section {
		sectionSet[coord] = true
	}

	perimeter := 0

	directions := []GardenNode{
		{row: 0, col: 1},  // right
		{row: 1, col: 0},  // down
		{row: 0, col: -1}, // left
		{row: -1, col: 0}, // up
	}

	for _, coord := range section {
		for _, dir := range directions {
			neighbor := GardenNode{row: coord.row + dir.row, col: coord.col + dir.col}
			if !d.isInBounds(garden, neighbor.row, neighbor.col) || !sectionSet[neighbor] {
				// neighbor is not out of bounds and is part of the section
				perimeter++
			}
		}
	}

	return perimeter
}

// countCorners will count the number of corners in a given section. This effectively
// deterimines the number of *sides* a section has, which is necessary for calculating
// the cost in Part 2 of the assignment
func (d *Day12) countCorners(section []GardenNode) int {
	up := [][]GardenNode{}
	down := [][]GardenNode{}
	left := [][]GardenNode{}
	right := [][]GardenNode{}

	// Create a lookup map for quick node access by (row,col)
	nodeMap := make(map[[2]int]*GardenNode)
	for i := range section {
		n := &section[i]
		nodeMap[[2]int{n.row, n.col}] = n
	}

	visitedUp := make(map[*GardenNode]bool)
	visitedDown := make(map[*GardenNode]bool)
	visitedLeft := make(map[*GardenNode]bool)
	visitedRight := make(map[*GardenNode]bool)

	// For left/right vertices, we move vertically (up/down)
	// For up/down vertices, we move horizontally (left/right)

	// vertical vertices (for left/right)
	dfsVertical := func(start *GardenNode, checkLeft bool) []GardenNode {
		// checkLeft = true means we're forming a vertex of nodes with left=true
		// otherwise, right=true
		visited := visitedLeft
		dirFlag := func(n *GardenNode) bool { return n.left }
		if !checkLeft {
			visited = visitedRight
			dirFlag = func(n *GardenNode) bool { return n.right }
		}

		stack := []*GardenNode{start}
		vertexNodes := []GardenNode{}
		for len(stack) > 0 {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if visited[n] {
				continue
			}
			visited[n] = true
			vertexNodes = append(vertexNodes, *n)

			// vertical connections: move up and down
			// Up neighbor
			upNeighborPos := [2]int{n.row - 1, n.col}
			if nn, ok := nodeMap[upNeighborPos]; ok && !visited[nn] && dirFlag(nn) {
				stack = append(stack, nn)
			}

			// Down neighbor
			downNeighborPos := [2]int{n.row + 1, n.col}
			if nn, ok := nodeMap[downNeighborPos]; ok && !visited[nn] && dirFlag(nn) {
				stack = append(stack, nn)
			}
		}

		return vertexNodes
	}

	// horizontal vertices (for up/down)
	dfsHorizontal := func(start *GardenNode, checkUp bool) []GardenNode {
		// checkUp = true means we're forming a vertex of nodes with up=true
		// otherwise, down=true
		visited := visitedUp
		dirFlag := func(n *GardenNode) bool { return n.up }
		if !checkUp {
			visited = visitedDown
			dirFlag = func(n *GardenNode) bool { return n.down }
		}

		stack := []*GardenNode{start}
		vertexNodes := []GardenNode{}
		for len(stack) > 0 {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if visited[n] {
				continue
			}
			visited[n] = true
			vertexNodes = append(vertexNodes, *n)

			// horizontal connections: move left and right
			// Left neighbor
			leftNeighborPos := [2]int{n.row, n.col - 1}
			if nn, ok := nodeMap[leftNeighborPos]; ok && !visited[nn] && dirFlag(nn) {
				stack = append(stack, nn)
			}

			// Right neighbor
			rightNeighborPos := [2]int{n.row, n.col + 1}
			if nn, ok := nodeMap[rightNeighborPos]; ok && !visited[nn] && dirFlag(nn) {
				stack = append(stack, nn)
			}
		}

		return vertexNodes
	}

	for i := range section {
		n := &section[i]
		if n.left && !visitedLeft[n] {
			v := dfsVertical(n, true) // build vertical vertex for left
			if len(v) > 0 {
				left = append(left, v)
			}
		}

		if n.right && !visitedRight[n] {
			v := dfsVertical(n, false) // build vertical vertex for right
			if len(v) > 0 {
				right = append(right, v)
			}
		}

		if n.up && !visitedUp[n] {
			v := dfsHorizontal(n, true) // build horizontal vertex for up
			if len(v) > 0 {
				up = append(up, v)
			}
		}

		if n.down && !visitedDown[n] {
			v := dfsHorizontal(n, false) // build horizontal vertex for down
			if len(v) > 0 {
				down = append(down, v)
			}
		}
	}

	// return the sum of the vertices
	return len(up) + len(down) + len(left) + len(right)
}

// setDirectionalFlags sets the left, up, right, and down flags for each GardenCoordinate
func (d *Day12) setDirectionalFlags(garden Garden, coordinates []GardenNode) []GardenNode {
	// Create a map for quick lookup of coordinates in the given slice
	coordinateSet := make(map[GardenNode]bool)
	for _, coord := range coordinates {
		coordinateSet[coord] = true
	}

	// Define directions and their corresponding flag names
	directions := []struct {
		drow, dcol int
		flagName   string
		setFlag    func(*GardenNode, bool)
	}{
		{-1, 0, "up", func(c *GardenNode, val bool) { c.up = val }},
		{1, 0, "down", func(c *GardenNode, val bool) { c.down = val }},
		{0, -1, "left", func(c *GardenNode, val bool) { c.left = val }},
		{0, 1, "right", func(c *GardenNode, val bool) { c.right = val }},
	}

	// Iterate through each coordinate and set the directional flags
	for i := range coordinates {
		current := &coordinates[i]
		for _, dir := range directions {
			neighbor := GardenNode{row: current.row + dir.drow, col: current.col + dir.dcol}
			if !d.isInBounds(garden, neighbor.row, neighbor.col) || !coordinateSet[neighbor] {
				// Neighbor is out of bounds or not in the slice
				dir.setFlag(current, true)
			}
		}
	}

	return coordinates
}
