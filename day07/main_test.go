package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	data, err := os.ReadFile("test1.txt")
	if err != nil {
		t.Error(err)
	}

	input := parseInput(string(data))

	result := part1(input)
	expected := 21

	if result != expected {
		t.Errorf("part1 = %d; want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	data, err := os.ReadFile("test1.txt")
	if err != nil {
		t.Error(err)
	}

	input := parseInput(string(data))

	result := part2(input)
	expected := 40

	if result != expected {
		t.Errorf("part2 = %d; want %d", result, expected)
	}
}
