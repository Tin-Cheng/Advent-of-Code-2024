package main

import (
	"math"
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
	lines := strings.Split(input, "\n")
	var result int64
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		target, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic("Error parsing target!")
		}
		numbersStr := strings.Split(parts[1], " ")
		var numbers []int64
		for _, numberStr := range numbersStr {
			number, err1 := strconv.ParseInt(numberStr, 10, 64)
			if err1 != nil {
				panic("Error parsing number!")
			}
			numbers = append(numbers, number)
		}
		if solvable(part2, target, numbers, 1, numbers[0]) {
			result += target
		}
	}
	return result
}

func solvable(part2 bool, tar int64, nums []int64, index int, cur int64) bool {
	if index == len(nums) {
		return cur == tar
	}
	return solvable(part2, tar, nums, index+1, cur+nums[index]) ||
		solvable(part2, tar, nums, index+1, cur*nums[index]) ||
		(part2 && solvable(part2, tar, nums, index+1, combine(cur, nums[index])))
}

func combine(a int64, b int64) int64 {
	c := strconv.FormatInt(b, 10)
	d := a*int64(math.Pow(10, float64(len(c)))) + b
	return d
}
