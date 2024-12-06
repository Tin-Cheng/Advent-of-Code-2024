package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
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
	// solve part 1 here
	result := 0
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	rules := strings.Split(parts[0], "\n")
	for _, line := range strings.Split(parts[1], "\n") {
		arr := strings.Split(line, ",")
		if checkAndSwap(rules, arr) {
			if !part2 {
				result += getMidVal(arr)
			}
		} else {
			if !part2 {
				continue
			}
			for !checkAndSwap(rules, arr) {

			}
			result += getMidVal(arr)
		}
	}
	return result
}

func getMidVal(arr []string) int {
	midIndex := (len(arr)+1)/2 - 1
	val, err := strconv.Atoi(arr[midIndex])
	if err != nil {
		panic(err)
	}
	return val
}

func checkAndSwap(rules []string, s []string) bool {
	len := len(s)
	for l := 0; l < len-1; l++ {
		for r := l + 1; r < len; r++ {
			if slices.Contains(rules, s[r]+"|"+s[l]) {
				t := s[l]
				s[l] = s[r]
				s[r] = t
				return false
			}
		}
	}
	return true
}
