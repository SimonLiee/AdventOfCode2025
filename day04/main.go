package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type ParsedInput []string

func main() {
	data, err := os.ReadFile("day04/input.txt")
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
	sum := 0
	for x := range input {
		for y := range len(input[x]) {
			if input[x][y] == '@' && checkAccess(input, x, y) < 4 {
				sum += 1
			}
		}
	}
	return sum
}

func part2(input ParsedInput) int {
	sum := 0
	accesible := [][2]int{}
	for {
		for x := range input {
			for y := range len(input[x]) {
				if input[x][y] == '@' && checkAccess(input, x, y) < 4 {
					sum += 1
					accesible = append(accesible, [2]int{x, y})
				}
			}
		}
		if len(accesible) == 0 {
			break
		}
		for _, roll := range accesible {
			setChar(input, roll[0], roll[1], '.')
		}
		accesible = [][2]int{}
	}

	return sum
}

func setChar(input []string, x, y int, c byte) {
	row := []byte(input[x])
	row[y] = c
	input[x] = string(row)
}

func checkAccess(input ParsedInput, x, y int) int {
	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	maxX := len(input)
	if maxX == 0 {
		return 0
	}
	maxY := len(input[0])

	for _, d := range dirs {
		nx := x + d[0]
		ny := y + d[1]

		if nx < 0 || ny < 0 || nx >= maxX || ny >= maxY {
			continue
		}

		if input[nx][ny] == '@' {
			count++
		}
	}

	return count
}

func parseInput(input string) ParsedInput {
	return strings.Fields(input)
}
