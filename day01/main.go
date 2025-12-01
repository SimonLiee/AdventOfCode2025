package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type ParsedInput []int

func main() {
	data, err := os.ReadFile("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := parseInput(string(data))

	fmt.Println("Part1:", part1(input))
	fmt.Println("Part2:", part2(input))
}

func part1(input ParsedInput) int {
	dial := 50
	zeroCount := 0
	for _, step := range input {
		dial += step

		dial = ((dial % 100) + 100) % 100

		if dial == 0 {
			zeroCount++
		}
	}
	return zeroCount
}

func part2(input ParsedInput) int {
	dial := 50
	zeroCount := 0
	for _, step := range input {
		for range int(math.Abs(float64(step))) {
			if step > 0 {
				dial = (dial + 1) % 100
			} else {
				dial = (dial - 1 + 100) % 100
			}
			if dial == 0 {
				zeroCount++
			}
		}
	}
	return zeroCount
}

func parseInput(input string) ParsedInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ret := make(ParsedInput, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		step, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		if line[0] == 'L' {
			step = -step
		}

		ret = append(ret, step)
	}

	return ret
}
