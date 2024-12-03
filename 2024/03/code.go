package main

import (
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
	do := true
	result := 0
	i := 0
	for i < len(input)-7 {
		if part2 {
			if input[i:i+4] == "do()" {
				do = true
				i += 4
				continue
			}
			if input[i:i+7] == "don't()" {
				do = false
				i += 7
				continue
			}
			if !do {
				i += 1
				continue
			}
		}
		if input[i:i+4] != "mul(" {
			i += 1
			continue
		}
		j := strings.Index(input[i+4:], ",")
		if j < 0 {
			i += 4
			continue
		}
		x, err1 := strconv.Atoi(input[i+4 : i+4+j])
		if err1 != nil {
			i += 4
			continue
		}
		k := strings.Index(input[i+4+j+1:], ")")
		if k < 0 {
			i += 4
			continue
		}
		y, err2 := strconv.Atoi(input[i+4+j+1 : i+4+j+1+k])
		if err2 != nil {
			i += 4
			continue
		}
		result += x * y
		i = i + 4 + j + 1 + k + 1
	}
	return result
}
