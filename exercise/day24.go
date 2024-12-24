// day24.go is the implementation for the twenty-fourth day of the Advent of Code 2024
// day2 borrows heavily from https://github.com/dickeyy
package exercise

import (
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day24 represents the data necessary to process the Exercise
	Day24 struct {
		name string
		file string
	}

	Bits        map[string]bool
	Instruction struct {
		Source      [2]string
		Destination string
		Operation   int
	}

	WireGraph struct {
		Nodes map[string]*Instruction   // Map of wire -> Instruction
		Edges map[string][]*Instruction // Map of wire -> List of dependent Instructions
		Order []*Instruction            // Execution order (set after sorting)
	}
)

const (
	AND = iota
	OR
	XOR
)

// GetName returns the name of the Day 24 exercise
func (d *Day24) GetName() string {
	return d.name
}

// Run executes the solution for Day 24 by retrieving the default file contents and uses that data
func (d *Day24) Run(w io.Writer) {
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

// RunFromInput executs the Day 24 solution using the provided input data
func (d *Day24) RunFromInput(w io.Writer, input []string) {
	bits, instructions := d.parseInput(input)
	zVal := d.Part1(bits, instructions)
	w.Write([]byte(fmt.Sprintf("Day 24 - Part 1 - The value of the wires that start with 'z' is %d.\n", zVal)))

	swappedRegisters := d.Part2(bits, instructions)
	w.Write([]byte(fmt.Sprintf("Day 24 - Part 2 - The swapped registers are %s.\n", swappedRegisters)))
}

// Part1 runs the instructions (after their constituent wires have been loaded with values)
// and calculates the decimal value of the wires that start with 'z', as per the instructions
func (d *Day24) Part1(bits Bits, instructions []Instruction) int {
	graph := NewWireGraph()
	for _, instr := range instructions {
		graph.AddInstruction(&instr)
	}

	graph.OrderSort(bits)

	// Execute instructions
	finalBits := executeInstructions(graph, bits)
	zVal := getZVal(finalBits)
	return zVal
}

// Part2 finds all of the registers that are swapped and returns them in alphabetical order
func (d *Day24) Part2(bits Bits, instructions []Instruction) string {
	graph := NewWireGraph()
	for _, instr := range instructions {
		graph.AddInstruction(&instr)
	}

	graph.OrderSort(bits)

	swapRegisters := findSwapRegisters(graph)
	sort.Strings(swapRegisters)
	return strings.Join(swapRegisters, ",")
}

// parseInput takes the specified input and produces a Bits map and a slice of
// Instructions to run
func (d *Day24) parseInput(input []string) (Bits, []Instruction) {
	bits := make(Bits)
	var instructions []Instruction

	parsingBits := true

	for _, line := range input {
		line = strings.TrimSpace(line)

		if line == "" {
			parsingBits = false
			continue
		}

		if parsingBits {
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				log.Fatalf("invalid input")
			}
			key := strings.TrimSpace(parts[0])
			valueStr := strings.TrimSpace(parts[1])
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				log.Fatalf("invalid input")
			}
			bits[key] = (value != 0)
		} else {
			// x00 AND y00 -> z00

			parts := strings.Split(line, "->")
			if len(parts) != 2 {
				log.Fatalf("invalid input")
			}
			operationPart := strings.TrimSpace(parts[0])
			destination := strings.TrimSpace(parts[1])

			tokens := strings.Fields(operationPart)
			if len(tokens) != 3 {
				log.Fatalf("invalid input")
			}

			source0 := tokens[0]
			opStr := strings.ToUpper(tokens[1])
			source1 := tokens[2]

			// map the operation string to the corresponding operator
			var op int
			switch opStr {
			case "AND":
				op = AND
			case "OR":
				op = OR
			case "XOR":
				op = XOR
			default:
				log.Fatalf("invalid input")
			}

			instr := Instruction{
				Source:      [2]string{source0, source1},
				Destination: destination,
				Operation:   op,
			}
			instructions = append(instructions, instr)
		}
	}

	return bits, instructions
}

// getZVal constructs the decimal value of all bits whose wire name starts with 'z' at
// the index provided in the name - e.g. "z03" = true results in 8 (2^3), "z05" = true
// results in 32 (2^5)
func getZVal(bits Bits) int {
	result := 0
	for wire, bitValue := range bits {
		if wire[0] == 'z' {
			zIndex, _ := strconv.Atoi(wire[1:])
			if bitValue {
				result += 1 << zIndex
			}
		}
	}

	return result
}

// NewGraph initializes a new graph.
func NewWireGraph() *WireGraph {
	return &WireGraph{
		Nodes: make(map[string]*Instruction),
		Edges: make(map[string][]*Instruction),
		Order: []*Instruction{},
	}
}

// AddInstruction adds an instruction to the graph.
func (g *WireGraph) AddInstruction(instr *Instruction) {
	g.Nodes[instr.Destination] = instr
	for _, source := range instr.Source {
		g.Edges[source] = append(g.Edges[source], instr)
	}
}

// OrderSort resolves the order of instructions based on dependencies.
func (g *WireGraph) OrderSort(initialWires map[string]bool) {
	visited := make(map[string]bool)
	stack := []*Instruction{}

	// DFS
	var visit func(wire string)
	visit = func(wire string) {
		if visited[wire] {
			return
		}
		visited[wire] = true

		// visit dependent instructions
		for _, instr := range g.Edges[wire] {
			visit(instr.Destination)
		}

		// add the producing instruction (if it exists)
		if instr, exists := g.Nodes[wire]; exists {
			stack = append([]*Instruction{instr}, stack...) // Prepend to stack
		}
	}

	for wire := range initialWires {
		visit(wire)
	}

	g.Order = stack
}

// Execute the instructions in the resolved order.
func executeInstructions(graph *WireGraph, initialBits map[string]bool) map[string]bool {
	bits := make(map[string]bool)

	// set the initial bits
	for k, v := range initialBits {
		bits[k] = v
	}

	// execute instructions in order
	for _, instr := range graph.Order {
		src0, src1 := instr.Source[0], instr.Source[1]
		val0, val1 := bits[src0], bits[src1]

		var result bool
		switch instr.Operation {
		case AND:
			result = val0 && val1
		case OR:
			result = val0 || val1
		case XOR:
			result = val0 != val1
		}

		bits[instr.Destination] = result
	}

	return bits
}

// findSwapRegisters navigates through every output to determine which of their parent registers are swapped
func findSwapRegisters(graph *WireGraph) []string {
	var swapped []string
	var carry string

	// check each bit position of the output - z45 is the largest of the input so we iterate
	// up to 45
	for i := 0; i < 45; i++ {
		xVal := fmt.Sprintf("x%02d", i)
		yVal := fmt.Sprintf("y%02d", i)

		var m1, n1, r1, z1, c1 string

		// Find half adder gates
		m1 = find(graph, xVal, yVal, XOR)
		n1 = find(graph, xVal, yVal, AND)

		if carry != "" {
			// a carry value exists
			// try the full adder
			r1 = find(graph, carry, m1, AND)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = find(graph, carry, m1, AND)
			}

			z1 = find(graph, carry, m1, XOR)

			// check for misplaced z wires
			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}
			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}
			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}

			c1 = find(graph, r1, n1, OR)
		}

		if strings.HasPrefix(c1, "z") && c1 != "z45" {
			// the last output - not subject to a carry
			c1, z1 = z1, c1
			swapped = append(swapped, c1, z1)
		}

		if carry == "" {
			carry = n1
		} else {
			carry = c1
		}
	}

	return swapped
}

// find returns the destination of the specified source registers that use the specified operator
func find(graph *WireGraph, reg1 string, reg2 string, operator int) string {
	for _, node := range graph.Nodes {
		registers := []string{node.Source[0], node.Source[1]}
		if valueExistsInStrings(reg1, registers) && valueExistsInStrings(reg2, registers) && node.Operation == operator {
			return node.Destination
		}
	}

	return ""
}

// valueExistsInStrings returns true if the value specified is in the specified set of strings
func valueExistsInStrings(value string, set []string) bool {
	for _, setVal := range set {
		if value == setVal {
			return true
		}
	}

	return false
}
