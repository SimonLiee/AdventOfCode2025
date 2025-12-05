package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ParsedInput struct {
	ranges [][2]int
	ids    []int
}

func main() {
	data, err := os.ReadFile("day05/input.txt")
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
	for _, id := range input.ids {
		if checkFresh(id, input.ranges) {
			ret++
		}
	}
	return ret
}

func checkFresh(id int, ranges [][2]int) bool {
	for _, rng := range ranges {
		if id >= rng[0] && id <= rng[1] {
			return true
		}
	}
	return false
}

func part2(input ParsedInput) int {
	noOverlap := removeOverlap(input.ranges)
	validIds := 0
	for _, rng := range noOverlap {
		validIds += rng[1] - rng[0] + 1
	}
	return validIds
}

func removeOverlap(ranges [][2]int) [][2]int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := [][2]int{ranges[0]}

	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]

		if r[0] <= last[1] {
			if r[1] > last[1] {
				last[1] = r[1]
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}

func parseInput(input string) ParsedInput {
	var ret ParsedInput

	tmp := strings.Split(input, "\n\n")
	strRanges := strings.Fields(tmp[0])
	for _, strRange := range strRanges {
		splitRange := strings.Split(strRange, "-")
		start, err := strconv.Atoi(splitRange[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(splitRange[1])
		if err != nil {
			log.Fatal(err)
		}
		ret.ranges = append(ret.ranges, [2]int{start, end})
	}

	for _, strId := range strings.Fields(tmp[1]) {
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Fatal(err)
		}
		ret.ids = append(ret.ids, id)
	}

	return ret
}
