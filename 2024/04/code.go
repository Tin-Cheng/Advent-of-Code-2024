package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var directions = [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func main() {
	aoc.Harness(run)
}

// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	result := 0
	lines := strings.Split(input, "\n")
	X := len(lines[0])
	Y := len(lines)
	for x := range X {
		for y := range Y {
			if part2 {
				if checkXmasString(lines, X, Y, x, y) {
					result++
				}
			} else {
				result += check(lines, X, Y, x, y)
			}
		}
	}

	return result
}
func checkXmasString(arr []string, X int, Y int, x int, y int) bool {
	if x < 1 || y < 1 || x >= X-1 || y >= Y-1 || arr[x][y] != 'A' {
		return false
	}

	checkPattern := func(a, b, c, d byte) bool {
		return (arr[x-1][y-1] == a && arr[x-1][y+1] == b &&
			arr[x+1][y-1] == c && arr[x+1][y+1] == d) ||
			(arr[x-1][y-1] == a && arr[x+1][y-1] == b &&
				arr[x-1][y+1] == c && arr[x+1][y+1] == d) ||
			(arr[x-1][y-1] == c && arr[x-1][y+1] == d &&
				arr[x+1][y-1] == a && arr[x+1][y+1] == b) ||
			(arr[x-1][y-1] == c && arr[x+1][y-1] == d &&
				arr[x-1][y+1] == a && arr[x+1][y+1] == b)
	}

	return checkPattern('M', 'M', 'S', 'S') || checkPattern('S', 'S', 'M', 'M')
}

func check(arr []string, X int, Y int, x int, y int) int {
	result := 0
	for _, direction := range directions {
		if checkString(arr, X, Y, x, y, direction[0], direction[1], 0, "XMAS") {
			result++
		}
	}
	return result
}
func checkString(arr []string, X, Y, x, y, dx, dy, i int, target string) bool {
	if i == len(target) {
		return true
	}
	if x < 0 || y < 0 || x >= X || y >= Y {
		return false
	}
	if arr[x][y] == target[i] {
		return checkString(arr, X, Y, x+dx, y+dy, dx, dy, i+1, target)
	}
	return false
}
