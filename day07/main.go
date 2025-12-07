package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ParsedInput [][]rune

func main() {
	data, err := os.ReadFile("day07/input.txt")
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
	splits := 0
	for y, line := range input[1:] {
		y = y + 1
		for x := range line {
			if input[y-1][x] == 'S' || input[y-1][x] == '|' {
				if input[y][x] == '^' {
					splits += 1
					if x-1 >= 0 {
						input[y][x-1] = '|'
					}
					if x+1 < len(line) {
						input[y][x+1] = '|'
					}
				} else {
					input[y][x] = '|'
				}
			}
		}
	}

	return splits
}

func part2(input ParsedInput) int {
	newInput := make([][]string, len(input))
	for i, line := range input {
		newInput[i] = make([]string, len(line))
		for j, char := range line {
			if char == 'S' {
				newInput[i][j] = "1"
			} else {
				newInput[i][j] = string(char)
			}
		}
	}

	for y, line := range newInput[1:] {
		y = y + 1
		for x := range line {
			num, err := strconv.Atoi(newInput[y-1][x])
			if err != nil {
				continue
			}

			if newInput[y][x] == "^" {
				if x-1 >= 0 {
					sideNum, err := strconv.Atoi(newInput[y][x-1])
					if err != nil {
						newInput[y][x-1] = strconv.Itoa(num)
					} else {
						newInput[y][x-1] = strconv.Itoa(num + sideNum)
					}
				}
				if x+1 < len(line) {
					sideNum, err := strconv.Atoi(newInput[y][x+1])
					if err != nil {
						newInput[y][x+1] = strconv.Itoa(num)
					} else {
						newInput[y][x+1] = strconv.Itoa(num + sideNum)
					}
				}
			} else {
				num2, err := strconv.Atoi(newInput[y][x])
				if err != nil {
					newInput[y][x] = strconv.Itoa(num)
				} else {
					newInput[y][x] = strconv.Itoa(num + num2)
				}
			}
		}
	}

	splits := 0
	for _, strNum := range newInput[len(newInput)-1] {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			continue
		}
		splits += num
	}

	return splits
}

func parseInput(input string) ParsedInput {
	lines := strings.FieldsFunc(
		input,
		func(c rune) bool { return c == '\n' },
	)

	ret := make([][]rune, len(lines))
	for i, line := range lines {
		ret[i] = []rune(line)
	}

	return ret
}
