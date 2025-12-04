package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type ParsedInput [][]int

func main() {
	data, err := os.ReadFile("day03/input.txt")
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
	ret := 0

	for _, battery := range input {
		num1 := 0
		num2 := 0
		for i, num := range battery {
			if num > num1 && i < len(battery)-1 {
				num1 = num
				num2 = 0
			} else if num > num2 {
				num2 = num
			}
		}
		ret += num1*10 + num2
	}

	return ret
}

func part2(input ParsedInput) int {
	ret := 0

	for _, battery := range input {
		joltage := 0
		nextDigitPos := 0
		for digit := 12 - 1; digit >= 0; digit-- {
			currentNum := 0
			for pos := nextDigitPos; pos < len(battery)-digit; pos++ {
				if battery[pos] > currentNum {
					currentNum = battery[pos]
					nextDigitPos = pos + 1
				}
			}
			joltage += currentNum * int(math.Pow10(digit))
		}
		ret += joltage
	}

	return ret
}

func parseInput(input string) ParsedInput {
	strBatteries := strings.Fields(input)

	batteries := make(ParsedInput, 0, len(strBatteries))
	for _, strBattery := range strBatteries {
		battery := make([]int, 0, len(strBattery))

		for _, charNum := range strBattery {
			battery = append(battery, int(charNum-'0'))
		}

		batteries = append(batteries, battery)
	}

	return batteries
}
