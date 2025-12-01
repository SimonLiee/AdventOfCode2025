package main

import (
	"fmt"
	"log"
	"os"
)

type ParsedInput []int

func main() {
	data, err := os.ReadFile("day00/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var input = parseInput(string(data))

	var p1 = part1(input)
	var p2 = part2(input)

	fmt.Println("Part1: ", p1)
	fmt.Println("Part2: ", p2)
}

func part1(input ParsedInput) int {
	return 0
}

func part2(input ParsedInput) int {
	return 0
}

func parseInput(input string) ParsedInput {
	return []int{0}
}
