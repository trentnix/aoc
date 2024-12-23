package exercise

import (
	"testing"
)

func TestDay23Part1(t *testing.T) {
	input := []string{
		"kh-tc",
		"qp-kh",
		"de-cg",
		"ka-co",
		"yn-aq",
		"qp-ub",
		"cg-tb",
		"vc-aq",
		"tb-ka",
		"wh-tc",
		"yn-cg",
		"kh-ub",
		"ta-co",
		"de-co",
		"tc-td",
		"tb-wq",
		"wh-td",
		"ta-ka",
		"td-qp",
		"aq-cg",
		"wq-ub",
		"ub-vc",
		"de-ta",
		"wq-aq",
		"wq-vc",
		"wh-yn",
		"ka-de",
		"kh-ta",
		"co-tc",
		"wh-qp",
		"tb-vc",
		"td-yn",
	}

	d23 := Day23{}

	numInterconnectedComputersWithT := d23.Part1(input)

	expectedNumInterconnectedComputersWithT := 7

	if numInterconnectedComputersWithT != expectedNumInterconnectedComputersWithT {
		t.Errorf("Day 23 - Part 1 (interconnected computers that start with t) Test:\nwant %v\ngot %v\n", expectedNumInterconnectedComputersWithT, numInterconnectedComputersWithT)
	}
}

func TestDay23Part2(t *testing.T) {
}
