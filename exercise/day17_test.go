package exercise

import (
	"testing"
)

func TestDay17Part1Sample1(t *testing.T) {
	input := []string{
		"Register A: 0",
		"Register B: 0",
		"Register C: 9",
		"",
		"Program: 2,6",
	}

	d17 := Day17{}

	program := d17.parseInput(input)

	_ = d17.Part1(program)
	expectedOutput := uint64(1)

	if program.B != expectedOutput {
		t.Errorf("Day 17 - Part 1 (C=9) Test:\nwant %v\ngot %v\n", expectedOutput, program.B)
	}
}

func TestDay17Part1Sample2(t *testing.T) {
	input := []string{
		"Register A: 10",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 5,0,5,1,5,4",
	}

	d17 := Day17{}

	program := d17.parseInput(input)

	output := d17.Part1(program)
	expectedOutput := "0,1,2"

	if output != expectedOutput {
		t.Errorf("Day 17 - Part 1 (A=10) Test:\nwant %v\ngot %v\n", expectedOutput, output)
	}
}

func TestDay17Part1Sample3(t *testing.T) {
	input := []string{
		"Register A: 2024",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,1,5,4,3,0",
	}

	d17 := Day17{}

	program := d17.parseInput(input)

	output := d17.Part1(program)
	expectedOutput := "4,2,5,6,7,7,7,7,3,1,0"

	if output != expectedOutput {
		t.Errorf("Day 17 - Part 1 (A=2024) Test:\nwant %v\ngot %v\n", expectedOutput, output)
	}

	if program.A != 0 {
		t.Errorf("Day 17 - Part 1 (A=2024) Test:\nwant %v\ngot %v\n", 0, program.A)
	}
}

func TestDay17Part1Sample4(t *testing.T) {
	input := []string{
		"Register A: 0",
		"Register B: 29",
		"Register C: 0",
		"",
		"Program: 1,7",
	}

	d17 := Day17{}

	program := d17.parseInput(input)

	_ = d17.Part1(program)
	expectedOutput := uint64(26)

	if program.B != expectedOutput {
		t.Errorf("Day 17 - Part 1 (B=29) Test:\nwant %v\ngot %v\n", expectedOutput, program.B)
	}
}

func TestDay17Part1Sample5(t *testing.T) {
	input := []string{
		"Register A: 0",
		"Register B: 2024",
		"Register C: 43690",
		"",
		"Program: 4,0",
	}

	d17 := Day17{}

	program := d17.parseInput(input)

	_ = d17.Part1(program)
	expectedOutput := uint64(44354)

	if program.B != expectedOutput {
		t.Errorf("Day 17 - Part 1 (B=2024 and C=43690) Test:\nwant %v\ngot %v\n", expectedOutput, program.B)
	}
}

func TestDay17Part1(t *testing.T) {
	input := []string{
		"Register A: 729",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,1,5,4,3,0",
	}

	d17 := Day17{}

	program := d17.parseInput(input)

	output := d17.Part1(program)
	expectedOutput := "4,6,3,5,6,3,5,2,1,0"

	if output != expectedOutput {
		t.Errorf("Day 17 - Part 1 (last example) Test:\nwant %v\ngot %v\n", expectedOutput, output)
	}
}

func TestDay17Part1Input(t *testing.T) {
	input := []string{
		"Register A: 47006051",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 2,4,1,3,7,5,1,5,0,3,4,3,5,5,3,0",
	}

	d17 := Day17{}

	program := d17.parseInput(input)

	output := d17.Part1(program)
	expectedOutput := "6,2,7,2,3,1,6,0,5"

	if output != expectedOutput {
		t.Errorf("Day 17 - Part 1 (last example) Test:\nwant %v\ngot %v\n", expectedOutput, output)
	}
}

func TestDay17Part2TestConfirm(t *testing.T) {
	input := []string{
		"Register A: 117440",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,3,5,4,3,0",
	}

	d17 := Day17{}

	program := d17.parseInput(input)

	output := d17.Part1(program)
	expectedOutput := "0,3,5,4,3,0"

	if output != expectedOutput {
		t.Errorf("Day 17 - Part 2 (find A where the output matches) Test:\nwant %v\ngot %v\n", expectedOutput, output)
	}
}

func TestDay17Part2(t *testing.T) {
	input := []string{
		"Register A: 2024",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,3,5,4,3,0",
	}

	d17 := Day17{}

	program := d17.parseInput(input)

	output := d17.Part2(program)
	expectedOutput := uint64(117440)

	if output != expectedOutput {
		t.Errorf("Day 17 - Part 2 (find A where the output matches) Test:\nwant %v\ngot %v\n", expectedOutput, output)
	}
}

func TestDay17Part2Input(t *testing.T) {
	input := []string{
		"Register A: 236548287712877",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 2,4,1,3,7,5,1,5,0,3,4,3,5,5,3,0",
	}

	d17 := Day17{}

	program := d17.parseInput(input)

	output := d17.Part1(program)
	expectedOutput := "2,4,1,3,7,5,1,5,0,3,4,3,5,5,3,0"

	if output != expectedOutput {
		t.Errorf("Day 17 - Part 2 (find A where the output matches -input) Test:\nwant %v\ngot %v\n", expectedOutput, output)
	}
}
