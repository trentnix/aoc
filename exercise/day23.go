// day23.go is the implementation for the twenty-third day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"log"
	"sort"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day23 represents the data necessary to process the Exercise
	Day23 struct {
		name string
		file string
	}

	// Graph represents an undirected graph using an adjacency list
	ComputerGraph struct {
		adjacency map[string]map[string]bool
	}
)

// GetName returns the name of the Day 23 exercise
func (d *Day23) GetName() string {
	return d.name
}

// Run executes the solution for Day 23 by retrieving the default file contents and uses that data
func (d *Day23) Run(w io.Writer) {
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

// RunFromInput executs the Day 23 solution using the provided input data
func (d *Day23) RunFromInput(w io.Writer, input []string) {
	// part 1
	numSets := d.Part1(input)
	w.Write([]byte(fmt.Sprintf("Day 22 - Part 1 - The number of three inter-connected computers that start with 't' is %d.\n", numSets)))
}

// Part1 computers the number of interconnected computers where at least one computer starts with 't'
func (d *Day23) Part1(input []string) int {
	computerGraph := NewComputerGraph(input)
	setsOfThree := computerGraph.FindTrianglesThatStartWith('t')

	return len(setsOfThree)
}

// Part2
func (d *Day23) Part2() int {
	return 0
}

// NewComputerGraph takes the input and builds a graph of the computers and their
// connections
func NewComputerGraph(input []string) *ComputerGraph {
	computerGraph := ComputerGraph{
		adjacency: make(map[string]map[string]bool),
	}

	for _, pair := range input {
		nodes := strings.Split(pair, "-")
		if len(nodes) != 2 {
			log.Fatalf("Invalid pair: %s", pair)
		}

		node1 := strings.TrimSpace(nodes[0])
		node2 := strings.TrimSpace(nodes[1])

		computerGraph.AddEdge(node1, node2)
	}

	return &computerGraph
}

// AddEdge adds an undirected edge between node1 and node2
func (g *ComputerGraph) AddEdge(node1, node2 string) {
	if g.adjacency[node1] == nil {
		g.adjacency[node1] = make(map[string]bool)
	}

	if g.adjacency[node2] == nil {
		g.adjacency[node2] = make(map[string]bool)
	}

	g.adjacency[node1][node2] = true
	g.adjacency[node2][node1] = true
}

// FindTrianglesThatStartWith finds all of the sets of three connected computers on the
// specified graph where at least one of the computers starts with the value 't'
func (g *ComputerGraph) FindTrianglesThatStartWith(startsWith rune) [][]string {
	var triangles [][]string
	seen := make(map[[3]string]bool) // To ensure uniqueness

	for node, neighbors := range g.adjacency {
		neighborList := make([]string, 0, len(neighbors))
		for neighbor := range neighbors {
			neighborList = append(neighborList, neighbor)
		}

		// find mutual connections
		for i := 0; i < len(neighborList); i++ {
			for j := i + 1; j < len(neighborList); j++ {
				neighbor1 := neighborList[i]
				neighbor2 := neighborList[j]

				if g.adjacency[neighbor1][neighbor2] {
					// found a triangle, sort it for uniqueness
					triangle := [3]string{node, neighbor1, neighbor2}
					sort.Strings(triangle[:])

					if !seen[triangle] && HasEntryThatStartsWith(startsWith, triangle[:]) {
						// this is a new set and it has an entry that starts with the specified value
						seen[triangle] = true
						triangles = append(triangles, triangle[:])
					}
				}
			}
		}
	}

	return triangles
}

// HasEntryThatStartsWith cycles through the nodeSet and looks for an entry whose first
// character starts with the specified startsWith value
func HasEntryThatStartsWith(startsWith rune, nodeSet []string) bool {
	for _, entry := range nodeSet {
		if len(entry) > 0 && entry[0] == byte(startsWith) {
			return true
		}
	}

	return false
}

// PrintGraph prints the adjacency list of the graph
func (g *ComputerGraph) PrintGraph() {
	for node, neighbors := range g.adjacency {
		fmt.Printf("%s: ", node)
		var neighborList []string
		for neighbor := range neighbors {
			neighborList = append(neighborList, neighbor)
		}
		fmt.Println(strings.Join(neighborList, ", "))
	}
}
