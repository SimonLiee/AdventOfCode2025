package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ParsedInput [][2]int

func main() {
	data, err := os.ReadFile("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var input = parseInput(string(data))

	var p1 = part1(input)
	var p2 = part2(input)

	fmt.Println("part1: ", p1)
	fmt.Println("part2: ", p2)
}

func part1(input ParsedInput) int {
	ret := 0
	for _, r := range input {
		for i := r[0]; i <= r[1]; i++ {
			numStr := strconv.Itoa(i)
			if numStr[0:len(numStr)/2] == numStr[len(numStr)/2:] {
				ret += i
			}
		}
	}

	return ret
}

func part2(input ParsedInput) int {
	ret := 0
	for _, numRange := range input {
		for num := numRange[0]; num <= numRange[1]; num++ {
			numStr := strconv.Itoa(num)
			for j := 1; j <= (len(numStr))/2; j++ {
				if len(numStr)%j == 0 {
					success := true
					startNum := numStr[0:j]
					for strIdx := 0; strIdx < len(numStr); strIdx += j {
						if startNum != numStr[strIdx:strIdx+j] {
							success = false
							break
						}
					}
					if success {
						ret += num
						break
					}
				}
			}
		}
	}

	return ret
}

func parseInput(input string) ParsedInput {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	ret := make(ParsedInput, 0, len(ranges))

	for _, r := range ranges {
		rsplit := strings.Split(r, "-")
		r1, err := strconv.Atoi(rsplit[0])
		if err != nil {
			log.Fatal(err)
		}
		r2, err := strconv.Atoi(rsplit[1])
		if err != nil {
			log.Fatal(err)
		}

		ret = append(ret, [2]int{r1, r2})
	}

	return ret
}
