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

	// part 1
	maze = d.parseInput(input)
	visitedNodes := d.Part2(maze)
	w.Write([]byte(fmt.Sprintf("Day 16 - Part 2 - The number of nodes visited on the cheapest path(s) is %d.\n", visitedNodes)))
}

// Part1 traverses the maze and finds the shortest path cost from the start to
// the end positions
func (d *Day16) Part1(maze ReindeerMaze) int {
	start := maze.findLocation('S')
	end := maze.findLocation('E')
	startDirection := east

	reindeerMazeGraph := buildGraph(maze)

	// Find the cheapest path
	cheapestPath := findLowestCostPath(reindeerMazeGraph, start, end, startDirection)

	return cheapestPath
}

// Part2 finds the number of nodes visited along every path that happens to share
// the lowest cost
func (d *Day16) Part2(maze ReindeerMaze) int {
	start := maze.findLocation('S')
	end := maze.findLocation('E')
	startDirection := east

	reindeerMazeGraph := buildGraph(maze)

	_, allPaths := findAllMinimumPaths(reindeerMazeGraph, start, end, startDirection)
	// allPaths returns nodes, but not every position. We need to expand it to have every
	// position in the maze, not just the graph nodes
	expandedPaths := expandAllPaths(allPaths, reindeerMazeGraph)

	visitedPositions := make(map[ReindeerMazePoint]bool)
	for _, p := range expandedPaths {
		for _, pos := range p {
			// we visited this position on the specified path, set it to true
			visitedPositions[pos] = true
		}
	}

	// return the number of positions that were visited
	return len(visitedPositions)
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

// findLowestCostPath implement's Dijkstra's Algorithm to find the lowest cost path
// in the specified ReindeerMazeGraph. The parameters are:
// - graph is the ReindeerMazeGraph being traversed
// - start is the start node in the graph
// - end is the end node in the graph
// - startDirection determines which direction from the starting point the traversal will begin
//
// The return value is the cost of the path that was found.
func findLowestCostPath(graph ReindeerMazeGraph, start, end ReindeerMazePoint, startDirection int) int {
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

// findAllMinimumPaths implement's Dijkstra's Algorithm to find the lowest cost path and returns
// all of the paths that have the same minimum cost in the specified ReindeerMazeGraph. The parameters are:
// - graph is the ReindeerMazeGraph being traversed
// - start is the start node in the graph
// - end is the end node in the graph
// - startDirection determines which direction from the starting point the traversal will begin
//
// The return value is the cost of the path that was found and
func findAllMinimumPaths(graph ReindeerMazeGraph, start, end ReindeerMazePoint, startDirection int) (int, [][]ReindeerMazePoint) {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// visited[node][direction] = minimal cost to reach that node with that direction
	visited := make(map[ReindeerMazePoint]map[int]int)

	// parents[node][direction] = list of (node,direction) from which we arrived at this node/direction at minimal cost
	parents := make(map[ReindeerMazePoint]map[int][]struct {
		node      ReindeerMazePoint
		direction int
	})

	// Initialize with the start node and direction
	heap.Push(pq, &State{
		node:      graph[start],
		direction: startDirection,
		cost:      0,
	})
	// We know the cost to reach start with startDirection is 0
	if visited[start] == nil {
		visited[start] = make(map[int]int)
	}
	visited[start][startDirection] = 0

	// minimalEndCost: track the minimal cost to reach the end, once found.
	minimalEndCost := -1

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*State)

		// If we have found an end cost, and this state's cost is greater than that minimal cost, we can stop.
		// (Because the priority queue always gives us states in ascending order of cost, no cheaper or equal path
		// to the end can appear after we exceed the minimalEndCost.)
		if minimalEndCost != -1 && current.cost > minimalEndCost {
			break
		}

		// Check if we've reached the end node
		if current.node.point == end {
			// If this is the first time we reach the end, record this cost as minimal
			if minimalEndCost == -1 {
				minimalEndCost = current.cost
			}
			// Don't return immediately; continue to process other possible equal-cost paths
			continue
		}

		// If this state is not better than a previously found cost, skip
		if visited[current.node.point] != nil {
			if bestCost, ok := visited[current.node.point][current.direction]; ok && bestCost < current.cost {
				continue
			}
		}

		// Process neighbors
		for _, edge := range current.node.edges {
			turnCost := 0
			if edge.direction != current.direction {
				turnCost = 1000
			}
			newCost := current.cost + turnCost + edge.cost

			// If we already have an end cost and this new cost is worse, skip
			if minimalEndCost != -1 && newCost > minimalEndCost {
				continue
			}

			// Check if this new path to edge.to.node & edge.direction is better or equal to previously known paths
			if visited[edge.to.point] == nil {
				visited[edge.to.point] = make(map[int]int)
			}

			prevCost, found := visited[edge.to.point][edge.direction]
			if !found || newCost < prevCost {
				// Found a strictly better path
				visited[edge.to.point][edge.direction] = newCost

				if parents[edge.to.point] == nil {
					parents[edge.to.point] = make(map[int][]struct {
						node      ReindeerMazePoint
						direction int
					})
				}
				// Reset parents for this state because we found a strictly better path
				parents[edge.to.point][edge.direction] = []struct {
					node      ReindeerMazePoint
					direction int
				}{
					{current.node.point, current.direction},
				}

				heap.Push(pq, &State{
					node:      edge.to,
					direction: edge.direction,
					cost:      newCost,
				})
			} else if newCost == prevCost {
				// Found another minimal path of the same cost
				parents[edge.to.point][edge.direction] = append(
					parents[edge.to.point][edge.direction],
					struct {
						node      ReindeerMazePoint
						direction int
					}{current.node.point, current.direction},
				)
				// No need to push to pq because this cost is already known
			}
		}
	}

	// If we never found the end, return -1 and no paths
	if minimalEndCost == -1 {
		return -1, nil
	}

	// Reconstruct all minimal-cost paths
	allPaths := reconstructAllPaths(parents, visited, end, minimalEndCost)

	return minimalEndCost, allPaths
}

// reconstructAllPaths reconstructs all of the minimal paths from the parents map
func reconstructAllPaths(
	parents map[ReindeerMazePoint]map[int][]struct {
		node      ReindeerMazePoint
		direction int
	},
	visited map[ReindeerMazePoint]map[int]int,
	end ReindeerMazePoint,
	minCost int,
) [][]ReindeerMazePoint {
	allPaths := [][]ReindeerMazePoint{}
	// The end node could be reached with multiple directions
	for dir, c := range visited[end] {
		if c == minCost {
			// Reconstruct all paths that end at (end, dir)
			paths := reconstructPaths(parents, end, dir)
			allPaths = append(allPaths, paths...)
		}
	}
	return allPaths
}

// reconstructPaths finds the paths recursively from the parents map
func reconstructPaths(
	parents map[ReindeerMazePoint]map[int][]struct {
		node      ReindeerMazePoint
		direction int
	},
	current ReindeerMazePoint,
	currentDirection int,
) [][]ReindeerMazePoint {
	// If no parents, this might be the start node
	if parents[current] == nil || len(parents[current][currentDirection]) == 0 {
		// We assume the current node is the start node
		// Adjust this if you explicitly know the start
		return [][]ReindeerMazePoint{{current}}
	}

	var allPaths [][]ReindeerMazePoint
	for _, p := range parents[current][currentDirection] {
		subPaths := reconstructPaths(parents, p.node, p.direction)
		for _, sp := range subPaths {
			allPaths = append(allPaths, append(sp, current))
		}
	}
	return allPaths
}

// expandAllPaths takes the provided graph nodes and finds every maze point that
// the path touches
func expandAllPaths(allPaths [][]ReindeerMazePoint, graph ReindeerMazeGraph) [][]ReindeerMazePoint {
	expandedPaths := make([][]ReindeerMazePoint, 0, len(allPaths))

	for _, path := range allPaths {
		// Each path is a list of graph nodes (decision points)
		if len(path) == 0 {
			continue
		}
		expandedPath := []ReindeerMazePoint{path[0]} // start from the first node

		for i := 0; i < len(path)-1; i++ {
			startNode := path[i]
			endNode := path[i+1]

			// Find the edge connecting startNode -> endNode
			edge := findEdge(graph, startNode, endNode)
			if edge == nil {
				// This should not happen if the graph is consistent
				continue
			}

			// We have startNode and know edge.direction and edge.cost
			// Let's expand intermediate cells
			curX, curY := startNode.X, startNode.Y
			dx := directionDeltas[edge.direction].dx
			dy := directionDeltas[edge.direction].dy

			for step := 0; step < edge.cost; step++ {
				curX += dx
				curY += dy
				expandedPath = append(expandedPath, ReindeerMazePoint{X: curX, Y: curY})
			}
		}

		expandedPaths = append(expandedPaths, expandedPath)
	}

	return expandedPaths
}

// findEdge returns the edge that was traveled
func findEdge(graph ReindeerMazeGraph, from, to ReindeerMazePoint) *ReindeerMazeEdge {
	node := graph[from]
	if node == nil {
		return nil
	}
	for _, e := range node.edges {
		if e.to.point == to {
			return &e
		}
	}
	return nil
}
