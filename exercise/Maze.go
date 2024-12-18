// Maze.go defines a generic maze that is used by (at least) Day16 and Day18
package exercise

import (
	"container/heap"
	"fmt"
)

type (
	MazePoint struct {
		Y, X      int
		pointCost int
	}

	MazeNode struct {
		point MazePoint
		edges []MazeEdge
	}

	MazeEdge struct {
		to        *MazeNode // Destination node
		cost      int       // Distance to the destination
		direction int       // Direction of the edge (north, east, south, west)
	}

	Maze [][]MazeLocation

	MazeLocation struct {
		val rune
	}

	MazePath struct {
		positions []MazePoint
		cost      int
	}

	MazeGraph map[MazePoint]*MazeNode
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

// buildGraph takes the specified MemoryMaze and builds a graph structure out
// of the maze
func buildMazeGraph(maze Maze) MazeGraph {
	graph := make(MazeGraph)

	for y := range maze {
		for x := range maze[y] {
			current := MazePoint{X: x, Y: y}

			// skip walls
			if maze[y][x].val == '#' {
				continue
			}

			// identify nodes: intersections, corners, endpoints, or dead ends
			if isMazeNode(maze, current) {
				if graph[current] == nil {
					graph[current] = &MazeNode{
						point: current,
						edges: []MazeEdge{},
					}
				}

				// Explore paths from this node
				for dir := 0; dir < 4; dir++ { // Iterate over all 4 directions
					if neighbor, cost := findNextMazeNode(maze, current, dir); neighbor != nil {
						if graph[*neighbor] == nil {
							graph[*neighbor] = &MazeNode{
								point: *neighbor,
								edges: []MazeEdge{},
							}
						}

						// Add an edge between the current node and the found neighbor
						graph[current].edges = append(graph[current].edges, MazeEdge{
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
func findNextMazeNode(maze Maze, start MazePoint, direction int) (*MazePoint, int) {
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

		current := MazePoint{X: x, Y: y}

		// Stop if reaching a node
		if isMazeNode(maze, current) {
			return &current, distance
		}
	}
}

// IsNode checks if the current point in the graph is a node (intersection, corner,
// endpoint, or dead end)
func isMazeNode(maze Maze, point MazePoint) bool {
	x, y := point.X, point.Y
	cell := maze[y][x].val

	// treat start or end as nodes
	if cell == 'S' || cell == 'E' {
		return true
	}

	var openPositions []MazePoint
	for _, delta := range directionDeltas {
		ny, nx := y+delta.dy, x+delta.dx
		if ny >= 0 && ny < len(maze) && nx >= 0 && nx < len(maze[ny]) {
			if maze[ny][nx].val != '#' {
				openPositions = append(openPositions, MazePoint{X: nx, Y: ny})
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

// findLowestCostMazePath implement's Dijkstra's Algorithm to find the lowest cost path
// in the specified MazeGraph. The parameters are:
// - graph is the MazeGraph being traversed
// - start is the start node in the graph
// - end is the end node in the graph
// - startDirection determines which direction from the starting point the traversal will begin
//
// The return value is the cost of the path that was found.
func findLowestCostMazePath(graph MazeGraph, start, end MazePoint, startDirection int, calculateCost func(s *State, e *MazeEdge) int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// store minimum costs to each node from each direction
	visited := make(map[MazePoint]map[int]int)

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
			newCost := calculateCost(current, &edge)

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

// findAllMinimumMazePaths implement's Dijkstra's Algorithm to find the lowest cost path and returns
// all of the paths that have the same minimum cost in the specified MazeGraph. The parameters are:
// - graph is the MazeGraph being traversed
// - start is the start node in the graph
// - end is the end node in the graph
// - startDirection determines which direction from the starting point the traversal will begin
//
// The return value is the cost of the path that was found and
func findAllMinimumMazePaths(graph MazeGraph, start, end MazePoint, startDirection int, calculateCost func(s *State, e *MazeEdge) int) (int, [][]MazePoint) {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// visited[node][direction] = minimal cost to reach that node with that direction
	visited := make(map[MazePoint]map[int]int)

	// parents[node][direction] = list of (node,direction) from which we arrived at this node/direction at minimal cost
	parents := make(map[MazePoint]map[int][]struct {
		node      MazePoint
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
			newCost := calculateCost(current, &edge)

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
						node      MazePoint
						direction int
					})
				}
				// Reset parents for this state because we found a strictly better path
				parents[edge.to.point][edge.direction] = []struct {
					node      MazePoint
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
						node      MazePoint
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
	allPaths := reconstructAllMazePaths(parents, visited, end, minimalEndCost)

	return minimalEndCost, allPaths
}

// reconstructAllMazePaths reconstructs all of the minimal paths from the parents map
func reconstructAllMazePaths(
	parents map[MazePoint]map[int][]struct {
		node      MazePoint
		direction int
	},
	visited map[MazePoint]map[int]int,
	end MazePoint,
	minCost int,
) [][]MazePoint {
	allPaths := [][]MazePoint{}
	// The end node could be reached with multiple directions
	for dir, c := range visited[end] {
		if c == minCost {
			// Reconstruct all paths that end at (end, dir)
			paths := reconstructMemoryMazePaths(parents, end, dir)
			allPaths = append(allPaths, paths...)
		}
	}
	return allPaths
}

// reconstructMemoryMazePaths finds the paths recursively from the parents map
func reconstructMemoryMazePaths(
	parents map[MazePoint]map[int][]struct {
		node      MazePoint
		direction int
	},
	current MazePoint,
	currentDirection int,
) [][]MazePoint {
	// If no parents, this might be the start node
	if parents[current] == nil || len(parents[current][currentDirection]) == 0 {
		// We assume the current node is the start node
		// Adjust this if you explicitly know the start
		return [][]MazePoint{{current}}
	}

	var allPaths [][]MazePoint
	for _, p := range parents[current][currentDirection] {
		subPaths := reconstructMemoryMazePaths(parents, p.node, p.direction)
		for _, sp := range subPaths {
			allPaths = append(allPaths, append(sp, current))
		}
	}
	return allPaths
}

// expandAllMazePaths takes the provided graph nodes and finds every maze point that
// the path touches
func expandAllMazePaths(allPaths [][]MazePoint, graph MazeGraph) [][]MazePoint {
	expandedPaths := make([][]MazePoint, 0, len(allPaths))

	for _, path := range allPaths {
		// Each path is a list of graph nodes (decision points)
		if len(path) == 0 {
			continue
		}
		expandedPath := []MazePoint{path[0]} // start from the first node

		for i := 0; i < len(path)-1; i++ {
			startNode := path[i]
			endNode := path[i+1]

			// Find the edge connecting startNode -> endNode
			edge := findMazeEdge(graph, startNode, endNode)
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
				pointCost := edge.cost + step + 1
				expandedPath = append(expandedPath, MazePoint{X: curX, Y: curY, pointCost: pointCost})
			}
		}

		expandedPaths = append(expandedPaths, expandedPath)
	}

	return expandedPaths
}

// findMazeEdge returns the edge that was traveled
func findMazeEdge(graph MazeGraph, from, to MazePoint) *MazeEdge {
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

// Print prints the current maze to stdout
func (maze *Maze) Print() {
	m := *maze
	if len(m) == 0 || len((m)[0]) == 0 {
		return
	}

	height := len(m)
	width := len((m)[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%c", m[y][x].val)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}
