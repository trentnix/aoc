// day17.go is the implementation for the seventeenth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day17 represents the data necessary to process the Exercise
	Day17 struct {
		name string
		file string
	}

	DeviceProgram struct {
		A       int
		B       int
		C       int
		program []int
		output  string
	}
)

// GetName returns the name of the Day 17 exercise
func (d *Day17) GetName() string {
	return d.name
}

// Run executes the solution for Day 17 by retrieving the default file contents and uses that data
func (d *Day17) Run(w io.Writer) {
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

// RunFromInput executs the Day 17 solution using the provided input data
func (d *Day17) RunFromInput(w io.Writer, input []string) {
	// part1
	program := d.parseInput(input)
	programOutput := d.Part1(program)
	w.Write([]byte(fmt.Sprintf("Day 17 - Part 1 - The output of the program is: \n%s.\n", programOutput)))
}

// Part1 runs the program specified by the input and returns the output
func (d *Day17) Part1(program *DeviceProgram) string {
	return program.Run()
}

// Run navigates through the instruction set of the specified DeviceProgram and returns the resulting output
func (p *DeviceProgram) Run() string {
	nextOperation := 0
	for {
		nextOperation = p.DoInstruction(nextOperation)
		if nextOperation >= len(p.program)-1 {
			break
		}
	}

	return p.output
}

// DoInstruction identifies the opCode at the specified index and its subsequent operand,
// performs the operation, and returns the index of the next instruction and a bool whether
// the program should end
func (p *DeviceProgram) DoInstruction(index int) int {
	newIndex := index

	opCode := p.program[index]
	operand := p.program[index+1]

	jumpValue, doJump := p.RunOpCode(opCode, operand)
	if doJump {
		newIndex = jumpValue
	} else {
		newIndex += 2
	}

	return newIndex
}

// RunOpCode runs the specified operation using the specified operand
func (p *DeviceProgram) RunOpCode(opCode int, operand int) (int, bool) {
	comboOperand := p.getComboOperandValue(operand)
	switch opCode {
	case 0:
		p.A = p.dvOp(p.A, comboOperand)
	case 1:
		p.B = p.B ^ operand
	case 2:
		p.B = comboOperand % 8
	case 3:
		if p.A != 0 {
			return operand, true
		}
	case 4:
		p.B = p.B ^ p.C
	case 5:
		outResult := strconv.Itoa(comboOperand % 8)
		if len(p.output) == 0 {
			p.output += outResult
		} else {
			p.output += "," + outResult
		}
	case 6:
		p.B = p.dvOp(p.A, comboOperand)
	case 7:
		p.C = p.dvOp(p.A, comboOperand)
	}

	return 0, false
}

// the *dv instruction returns the divident / 2*operand (the truncated, not rounded value)
func (p *DeviceProgram) dvOp(dividend int, operand int) int {
	divisor := 1 << operand // calculate 2^operand

	// integer division in Go truncates automatically
	return dividend / divisor
}

// getComboOperandValue calculates the combo operand given the specified operand. Some
// instructions use the combo operand
func (p *DeviceProgram) getComboOperandValue(operand int) int {
	comboOperand := operand
	switch operand {
	case 4:
		comboOperand = p.A
	case 5:
		comboOperand = p.B
	case 6:
		comboOperand = p.C
	}

	return comboOperand
}

// Part2
func (d *Day17) Part2() int {
	return 0
}

// parseInput parses the specified input data and returns a corresponding DeviceProgram instance
func (d *Day17) parseInput(input []string) *DeviceProgram {
	var dp DeviceProgram

	// Parse register lines
	for _, line := range input {
		if strings.HasPrefix(line, "Register A:") {
			dp.A = parseRegisterValue(line)
		} else if strings.HasPrefix(line, "Register B:") {
			dp.B = parseRegisterValue(line)
		} else if strings.HasPrefix(line, "Register C:") {
			dp.C = parseRegisterValue(line)
		} else if strings.HasPrefix(line, "Program:") {
			parts := strings.Split(line, ":")
			valuesStr := strings.Split(strings.TrimSpace(parts[1]), ",")
			for _, v := range valuesStr {
				num, _ := strconv.Atoi(strings.TrimSpace(v))
				dp.program = append(dp.program, num)
			}
		}
	}

	return &dp
}

// parseRegisterValue extracts an integer value from a register line
func parseRegisterValue(line string) int {
	parts := strings.Split(line, ":")
	value, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	return value
}
