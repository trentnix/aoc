// priority_queue.go defines a PriorityQueue for use in various exercises
// solution
package exercise

// Priority Queue Item
type State struct {
	node      *MazeNode
	prev      *State
	direction int
	cost      int
	index     int // For heap management
}

// Priority Queue definition
type PriorityQueue []*State

// Len returns the length of the specified PriorityQueue
func (pq PriorityQueue) Len() int { return len(pq) }

// Less determines which of two State instances have a higher priority
// (based on having a lower cost)
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

// Swap swaps two State instances
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds an item to the specified PriorityQueue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes an item from the specified PriorityQueue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // For safety
	*pq = old[0 : n-1]
	return item
}

// IsTurn determines whether the particular node in the graph represents a change in direction
func (s *State) IsTurn() bool {
	if s.prev == nil {
		return false
	}

	if s.direction == s.prev.direction {
		return false
	}

	return true
}
