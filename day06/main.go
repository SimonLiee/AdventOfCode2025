package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	nums []int
	op   string
}

type ParsedInput []Row

func main() {
	data, err := os.ReadFile("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var input = parseInput(string(data))
	var input2 = parseInput2(string(data))

	var p1 = part1(input)
	var p2 = part2(input2)

	fmt.Println("Part1: ", p1)
	fmt.Println("Part2: ", p2)
}

func part1(input ParsedInput) int {
	ret := 0
	for _, problem := range input {
		sum := 0
		for _, num := range problem.nums {
			switch problem.op {
			case "*":
				if sum > 0 {
					sum *= num
				} else {
					sum = num
				}
			case "+":
				sum += num
			default:
				log.Fatal("Unexpected operator: ", problem.op)
			}
		}

		ret += sum
	}
	return ret
}

func part2(input ParsedInput) int {
	ret := part1(input)

	return ret
}

func parseInput2(input string) ParsedInput {
	lines := strings.FieldsFunc(input, func(r rune) bool {
		return r == '\n'
	})

	ops := strings.Fields(lines[len(lines)-1])
	ret := make(ParsedInput, len(ops))

	for i, op := range ops {
		ret[i].op = op
	}

	numbers := make([]string, len(lines[0]))
	for _, line := range lines {
		for j, c := range line {
			if c >= '0' && c <= '9' {
				numbers[j] = numbers[j] + string(c)
			}
		}
	}

	operation := 0
	for _, numStr := range numbers {
		if strings.TrimSpace(numStr) == "" {
			operation++
			continue
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}

		ret[operation].nums = append(ret[operation].nums, num)
	}
	return ret
}

func parseInput(input string) ParsedInput {
	tmp := strings.FieldsFunc(input, func(r rune) bool {
		return r == '\n'
	})

	var tmp2 [][]string
	for _, i := range tmp {
		tmp2 = append(tmp2, strings.Fields(i))
	}

	ret := make(ParsedInput, len(tmp2[len(tmp2)-1]))
	for i := range tmp2[:len(tmp2)-1] {
		for j := range tmp2[i] {
			num, err := strconv.Atoi(tmp2[i][j])
			if err != nil {
				log.Fatal(err)
				break
			}

			ret[j].nums = append(ret[j].nums, num)
		}
	}

	for i, op := range tmp2[len(tmp2)-1] {
		ret[i].op = op
	}

	return ret
}
