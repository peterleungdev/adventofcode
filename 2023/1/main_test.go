package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	sum := PartOne("input_example_part_1")
	if sum != 142 {
		t.Fatal("Part One example sum should be 142")
	}
}

func TestPartTwo(t *testing.T) {
	sum := PartTwo("input_example_part_2")
	if sum != 281 {
		t.Fatal("Part Two example sum should be 281")
	}
}
