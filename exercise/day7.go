// day7.go is the implementation for the seventh day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day7 represents the data necessary to process the Exercise
	Day7 struct {
		name string
		file string
	}

	Equation struct {
		Value  int64
		Inputs []int64
	}
)

// GetName returns the name of the Day 7 exercise
func (d *Day7) GetName() string {
	return d.name
}

// Run executes the solution for Day 7 by retrieving the default file contents and uses that data
func (d *Day7) Run(w io.Writer) {
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

// RunFromInput executs the Day 7 solution using the provided input data
func (d *Day7) RunFromInput(w io.Writer, input []string) {
	equations, err := d.parseInput(input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error trying to parse the input: %v.", err)))
		return
	}

	// part 1
	numTrueEquations := d.Part1(equations)
	w.Write([]byte(fmt.Sprintf("Day 7 - Part 1 - The number of true equations is %d.\n", numTrueEquations)))

	// part 2
	// numLoops := d.Part2(g.Copy())
	// w.Write([]byte(fmt.Sprintf("Day 7 - Part 2 - The number of new blocks that result in a loop is %d.\n", numLoops)))
}

// Part1 calculates the sum of solvable equations using an operator set of "+" and "*", if
// the equation is evaluated from left-to-right (ignoring operator precedence)
func (d *Day7) Part1(equations []Equation) uint64 {
	sumEquationValuesMadeTrue := uint64(0)
	operators := []string{"+", "*"}

	for _, e := range equations {
		operatorCombinations := d.generateCombinations(operators, len(e.Inputs)-1)
		for _, operatorCombination := range operatorCombinations {
			if d.evaluateEquationValues(e, operatorCombination) {
				sumEquationValuesMadeTrue += uint64(e.Value)
				break
			}
		}
	}

	return sumEquationValuesMadeTrue
}

// evaluateEquationValues determines whether the specified operators list will
// solve the specified equation (evaluating from left-to-right with no precedence
// rules)
func (d *Day7) evaluateEquationValues(e Equation, operators []string) bool {
	val := e.Inputs[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+":
			val += e.Inputs[i+1]
		case "*":
			val *= e.Inputs[i+1]
		}
	}

	return val == e.Value
}

// GenerateCombinations generates all possible combinations of the given operators
// for the specified length.
func (d *Day7) generateCombinations(operators []string, length int) [][]string {
	var results [][]string

	// Helper function for recursion
	var helper func(current []string)
	helper = func(current []string) {
		if len(current) == length {
			// If the current combination has the required length, add it to results
			results = append(results, append([]string(nil), current...))
			return
		}

		// Add each operator to the current combination and recurse
		for _, op := range operators {
			helper(append(current, op))
		}
	}

	// Start the recursion with an empty combination
	helper([]string{})
	return results
}

// Part2
func (d *Day7) Part2() int {
	return 0
}

// parseInput parses the input into a slice of Equation values
// e.g.:
//
//		3267: 81 40 27
//	 would result in Equation{Value: 3267, Inputs: {81, 40, 27}}
func (d *Day7) parseInput(input []string) ([]Equation, error) {
	var equations []Equation
	for _, line := range input {
		// Split the line into the value part and inputs part
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		// Parse the value (before the colon)
		value, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value format: %s", parts[0])
		}

		// Parse the inputs (after the colon)
		inputStrings := strings.Fields(parts[1])
		var inputs []int64
		for _, inputStr := range inputStrings {
			inputValue, err := strconv.ParseInt(inputStr, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid input format: %s", inputStr)
			}
			inputs = append(inputs, inputValue)
		}

		// Add the parsed Equation to the result
		equations = append(equations, Equation{Value: value, Inputs: inputs})
	}

	return equations, nil
}
