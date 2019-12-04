package main

import (
	"fmt"

	"github.com/kisom/aoc/advent2019/inst"
)

var minRange = 235741
var maxRange = 706948

func twoAdjacent(n int) bool {
	s := fmt.Sprintf("%0d", n)
	for i := 1; i < 6; i++ {
		if s[i-1] == s[i] {
			return true
		}
	}

	return false
}

func onlyTwoAdjacent(n int) bool {
	s := fmt.Sprintf("%0d", n)
	run := s[0]
	count := 1
	ok := false

	for i := 1; i < 6; i++ {
		if s[i] == run {
			count++
		} else {
			if count == 2 {
				ok = true
			}
			run = s[i]
			count = 1
		}
	}

	if count == 2 {
		ok = true
	}
	return ok
}

func isMonotic(n int) bool {
	s := fmt.Sprintf("%d", n)

	// relies on the fact that ASCII numerics are digit + 0x30, so
	// they can be compared directly without parsing.
	for i := 1; i < 6; i++ {
		if s[i] < s[i-1] {
			return false
		}
	}

	return true
}

func findCandidate1(min, max int) int {
	count := 0
	for i := min; i <= max; i++ {
		if twoAdjacent(i) && isMonotic(i) {
			count++
		}
	}

	return count
}

func findCandidate2(min, max int) int {
	count := 0
	for i := min; i <= max; i++ {
		if onlyTwoAdjacent(i) && isMonotic(i) {
			count++
		}
	}

	return count
}

func part1() {
	fmt.Println(findCandidate1(minRange, maxRange))
}

func part2() {
	fmt.Println(findCandidate2(minRange, maxRange))
}

func main() {
	inst.Run("day04p1", part1)
	inst.Run("day04p2", part2)
}
