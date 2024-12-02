package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	//"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	result := 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		split := strings.Split(line, " ")
		if part2 {
			for i, _ := range split {
				oneLess := slices.Concat(split[:i], split[i+1:])
				if isSafe(oneLess) {
					result += 1
					break
				}
			}
		} else {
			if isSafe(split) {
				result += 1
			}
		}
	}
	return result
}

func isSafe(split []string) bool {
	isIncreasing := true
	last := 0
	for i, v := range split {
		cur := toInt(v)
		if i == 0 {
			last = cur
			continue
		}
		if i == 1 && cur < last {
			isIncreasing = false
		}
		diff := math.Abs(float64(cur - last))
		if (cur == last) ||
			(isIncreasing && cur < last) ||
			(!isIncreasing && cur > last) ||
			(diff > 3) {
			return false
		}
		last = cur
	}
	return true
}

func toInt(s string) int {
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
