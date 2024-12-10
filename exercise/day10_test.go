package exercise

import (
	"testing"
)

func TestDay10Part1Simple(t *testing.T) {
	input := []string{
		"9990999",
		"9991999",
		"9992999",
		"6543456",
		"7111117",
		"8111118",
		"9111119",
	}

	d10 := Day10{}

	topo := d10.parseInput(input)

	calculatedValue := d10.Part1(topo)
	expectedValue := 2

	if calculatedValue != expectedValue {
		t.Errorf("Day 10 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay10Part1Simple2(t *testing.T) {
	input := []string{
		"1190669",
		"7771498",
		"4312117",
		"6543456",
		"7651987",
		"8761111",
		"9871111",
	}

	d10 := Day10{}

	topo := d10.parseInput(input)

	calculatedValue := d10.Part1(topo)
	expectedValue := 4

	if calculatedValue != expectedValue {
		t.Errorf("Day 10 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay10Part1Simple3(t *testing.T) {
	input := []string{
		"1066966",
		"2882811",
		"3111711",
		"4567654",
		"1118113",
		"1119662",
		"7777701",
	}

	d10 := Day10{}

	topo := d10.parseInput(input)

	calculatedValue := d10.Part1(topo)
	expectedValue := 3

	if calculatedValue != expectedValue {
		t.Errorf("Day 10 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay10Part1(t *testing.T) {
	input := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}

	d10 := Day10{}

	topo := d10.parseInput(input)

	calculatedValue := d10.Part1(topo)
	expectedValue := 36

	if calculatedValue != expectedValue {
		t.Errorf("Day 10 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay10Part2Simple(t *testing.T) {
	input := []string{
		"8888808",
		"8843218",
		"1158828",
		"1165438",
		"1171148",
		"1187658",
		"1191111",
	}

	d10 := Day10{}

	topo := d10.parseInput(input)

	calculatedValue := d10.Part2(topo)
	expectedValue := 3

	if calculatedValue != expectedValue {
		t.Errorf("Day 10 - Part 2 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay10Part2Simple2(t *testing.T) {
	input := []string{
		"1190669",
		"6661698",
		"1112117",
		"6543456",
		"7651987",
		"8761111",
		"9871111",
	}

	d10 := Day10{}

	topo := d10.parseInput(input)

	calculatedValue := d10.Part2(topo)
	expectedValue := 13

	if calculatedValue != expectedValue {
		t.Errorf("Day 10 - Part 2 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay10Part2Simple3(t *testing.T) {
	input := []string{
		"012345",
		"123456",
		"234567",
		"345678",
		"416789",
		"567891",
	}

	d10 := Day10{}

	topo := d10.parseInput(input)

	calculatedValue := d10.Part2(topo)
	expectedValue := 227

	if calculatedValue != expectedValue {
		t.Errorf("Day 10 - Part 2 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay10Part2(t *testing.T) {
	input := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}

	d10 := Day10{}

	topo := d10.parseInput(input)

	calculatedValue := d10.Part2(topo)
	expectedValue := 81

	if calculatedValue != expectedValue {
		t.Errorf("Day 10 - Part 2 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}
