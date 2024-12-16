// day16.go is the implementation for the sixteenth day of the Advent of Code 2024
package exercise

import (
	"container/heap"
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

	ReindeerMazeNode struct {
		point ReindeerMazePoint
		edges []ReindeerMazeEdge
	}

	ReindeerMazeEdge struct {
		to        *ReindeerMazeNode // Destination node
		cost      int               // Distance to the destination
		direction int               // Direction of the edge (north, east, south, west)
	}

	ReindeerMazeGraph map[ReindeerMazePoint]*ReindeerMazeNode
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

// Part1 traverses the maze and finds the shortest path cost from the start to
// the end positions
func (d *Day16) Part1(maze ReindeerMaze) int {
	start := maze.findLocation('S')
	end := maze.findLocation('E')
	startDirection := east

	reindeerMazeGraph := buildGraph(maze)

	// Find the cheapest path
	cheapestPath := dijkstra(reindeerMazeGraph, start, end, startDirection)

	return cheapestPath
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

// buildGraph takes the specified ReindeerMaze and builds a graph structure out
// of the maze
func buildGraph(maze ReindeerMaze) ReindeerMazeGraph {
	graph := make(ReindeerMazeGraph)

	for y := range maze {
		for x := range maze[y] {
			current := ReindeerMazePoint{X: x, Y: y}

			// skip walls
			if maze[y][x].val == '#' {
				continue
			}

			// identify nodes: intersections, corners, endpoints, or dead ends
			if isNode(maze, current) {
				if graph[current] == nil {
					graph[current] = &ReindeerMazeNode{
						point: current,
						edges: []ReindeerMazeEdge{},
					}
				}

				// Explore paths from this node
				for dir := 0; dir < 4; dir++ { // Iterate over all 4 directions
					if neighbor, cost := findNextNode(maze, current, dir); neighbor != nil {
						if graph[*neighbor] == nil {
							graph[*neighbor] = &ReindeerMazeNode{
								point: *neighbor,
								edges: []ReindeerMazeEdge{},
							}
						}

						// Add an edge between the current node and the found neighbor
						graph[current].edges = append(graph[current].edges, ReindeerMazeEdge{
							to:        graph[*neighbor],
							cost:      cost,
							direction: dir,
						})
					}
				}
			}
		}
	}

	return graph
}

// findNextNode traverses the graph according to the specified direction, returning
// the next point to traverse and the next direction faced
func findNextNode(maze ReindeerMaze, start ReindeerMazePoint, direction int) (*ReindeerMazePoint, int) {
	x, y := start.X, start.Y
	dy, dx := directionDeltas[direction].dy, directionDeltas[direction].dx
	distance := 0

	// traverse the graph in the provided direction until a Node is reached
	for {
		x, y = x+dx, y+dy // dx changes columns, dy changes rows
		distance++

		// out of bounds
		if y < 0 || y >= len(maze) || x < 0 || x >= len(maze[y]) {
			return nil, 0
		}

		// hit a wall - stop
		if maze[y][x].val == '#' {
			return nil, 0
		}

		current := ReindeerMazePoint{X: x, Y: y}

		// Stop if reaching a node
		if isNode(maze, current) {
			return &current, distance
		}
	}
}

// IsNode checks if the current point in the graph is a node (intersection, corner,
// endpoint, or dead end)
func isNode(maze ReindeerMaze, point ReindeerMazePoint) bool {
	x, y := point.X, point.Y
	cell := maze[y][x].val

	// treat start or end as nodes
	if cell == 'S' || cell == 'E' {
		return true
	}

	var openPositions []ReindeerMazePoint
	for _, delta := range directionDeltas {
		ny, nx := y+delta.dy, x+delta.dx
		if ny >= 0 && ny < len(maze) && nx >= 0 && nx < len(maze[ny]) {
			if maze[ny][nx].val != '#' {
				openPositions = append(openPositions, ReindeerMazePoint{X: nx, Y: ny})
			}
		}
	}

	openCount := len(openPositions)
	if openCount != 2 {
		return true
	}

	// Exactly two neighbors
	// Check if they form a line or a corner
	p1, p2 := openPositions[0], openPositions[1]

	// If both neighbors share the same row (y) or the same column (x), it's a straight line
	if p1.Y == p2.Y || p1.X == p2.X {
		return false
	}

	// it's a corner: it's a node
	return true
}

// djikstra implement's Dijkstra's Algorithm to find the lowest cost path
// in the specified ReindeerMazeGraph. The parameters are:
// - graph is the ReindeerMazeGraph being traversed
// - start is the start node in the graph
// - end is the end node in the graph
// - startDirection determines which direction from the starting point the traversal will begin
//
// The return value is the cost of the path that was found.
func dijkstra(graph ReindeerMazeGraph, start, end ReindeerMazePoint, startDirection int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// store minimum costs to each node from each direction
	visited := make(map[ReindeerMazePoint]map[int]int)

	// initialize the priority queue with the start node and direction
	heap.Push(pq, &State{
		node:      graph[start],
		direction: startDirection,
		cost:      0,
	})

	for pq.Len() > 0 {
		// get the node with the smallest cost
		current := heap.Pop(pq).(*State)

		// we reached the end, return the cost
		if current.node.point == end {
			return current.cost
		}

		// Check if we've seen a better cost for this node and direction
		if visited[current.node.point] == nil {
			visited[current.node.point] = make(map[int]int)
		}
		if costSoFar, ok := visited[current.node.point][current.direction]; ok && costSoFar <= current.cost {
			// we found a cheaper cost before, skip this one
			continue
		}

		visited[current.node.point][current.direction] = current.cost

		for _, edge := range current.node.edges {
			// calculate the cost to move to the neighbor
			turnCost := 0
			if edge.direction != current.direction {
				// this is a turn
				turnCost = 1000
			}
			newCost := current.cost + turnCost + edge.cost

			// if we've visited the neighbor at this direction cheaper, skip
			if visited[edge.to.point] != nil {
				if prevCost, ok := visited[edge.to.point][edge.direction]; ok && prevCost <= newCost {
					continue
				}
			}

			// add the neighbor to the priority queue
			heap.Push(pq, &State{
				node:      edge.to,
				direction: edge.direction,
				cost:      newCost,
			})
		}
	}

	// end not found
	return -1
}
