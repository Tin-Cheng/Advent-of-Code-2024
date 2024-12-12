package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type pair struct {
	R, C int
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// solve part 1 here
	M := strings.Split(input, "\n")
	R := len(M)
	C := len(M[0])
	visited := make(map[string]bool)
	startMap := make(map[string]string)
	directions := [][2]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	result := 0
	areas := make(map[string]int)
	sides := make(map[string]int)
	for r, line := range M {
		for c, _ := range line {
			area, parameter := search(visited, startMap, directions, r, c, M, R, C, line[c], coordsKey(r, c))
			areas[coordsKey(r, c)] = area
			result += area * parameter
		}
	}
	for r := 0; r < R; r++ {
		checkAddSideByFirstLastCol(R, C, r, 0, M, sides, startMap)
		cur := M[r][0]
		for c := 1; c < C; c++ {
			if cur != M[r][c] {
				checkAddSideByCol(R, C, r, c-1, M, sides, startMap, true)
				checkAddSideByCol(R, C, r, c, M, sides, startMap, false)
				cur = M[r][c]
			}
		}
		checkAddSideByFirstLastCol(R, C, r, C-1, M, sides, startMap)
	}
	for c := 0; c < C; c++ {
		checkAddSideByFirstLastRow(R, C, 0, c, M, sides, startMap)
		cur := M[0][c]
		for r := 1; r < R; r++ {
			if cur != M[r][c] {
				checkAddSideByRow(R, C, r-1, c, M, sides, startMap, true)
				checkAddSideByRow(R, C, r, c, M, sides, startMap, false)
				cur = M[r][c]
			}
		}
		checkAddSideByFirstLastRow(R, C, R-1, c, M, sides, startMap)
	}
	if part2 {
		result = 0
		for k, v := range sides {
			result += areas[k] * v
		}
	}
	return result
}

func checkMatch(R int, C int, r1, c1 int, M []string, r2, c2 int) bool {
	return isInMap(R, C, r1, c1) && M[r1][c1] == M[r2][c2]
}

func checkAddSideByFirstLastCol(R int, C int, r int, c int, M []string, sides map[string]int, startMap map[string]string) {
	if !(checkMatch(R, C, r-1, c, M, r, c)) {
		sides[startMap[coordsKey(r, c)]] += 1
	}
}

func checkAddSideByCol(R int, C int, r int, c int, M []string, sides map[string]int, startMap map[string]string, from bool) {
	if !isInMap(R, C, r-1, c) || M[r-1][c] != M[r][c] {
		sides[startMap[coordsKey(r, c)]] += 1
	} else if checkMatch(R, C, r-1, c, M, r, c) &&
		((from && checkMatch(R, C, r-1, c+1, M, r, c)) ||
			(!from && checkMatch(R, C, r-1, c-1, M, r, c))) {
		sides[startMap[coordsKey(r, c)]] += 1
	}
}

func checkAddSideByFirstLastRow(R int, C int, r int, c int, M []string, sides map[string]int, startMap map[string]string) {
	if !(isInMap(R, C, r, c-1) && M[r][c-1] == M[r][c]) {
		sides[startMap[coordsKey(r, c)]] += 1
	}
}

func checkAddSideByRow(R int, C int, r int, c int, M []string, sides map[string]int, startMap map[string]string, from bool) {
	if !isInMap(R, C, r, c-1) || M[r][c-1] != M[r][c] {
		sides[startMap[coordsKey(r, c)]] += 1
	} else if checkMatch(R, C, r, c-1, M, r, c) &&
		(from && checkMatch(R, C, r+1, c-1, M, r, c) ||
			(!from && checkMatch(R, C, r-1, c-1, M, r, c))) {
		sides[startMap[coordsKey(r, c)]] += 1
	}
}
func search(visited map[string]bool, startMap map[string]string, directions [][2]int, r int, c int, M []string, R int, C int, v byte, start string) (int, int) {
	key := coordsKey(r, c)
	if visited[key] {
		return 0, 0
	}
	visited[key] = true
	startMap[key] = start
	size := 1
	parameter := 4
	for _, direction := range directions {
		newR, newC := r+direction[0], c+direction[1]
		if isInMap(C, R, newC, newR) && M[newR][newC] == v {
			if visited[coordsKey(newR, newC)] {
				parameter -= 1
			} else {
				nextSize, nextParameter := search(visited, startMap, directions, newR, newC, M, R, C, v, start)
				size += nextSize
				parameter += nextParameter - 1
			}
		}
	}
	return size, parameter
}
func isInMap(width int, height int, x int, y int) bool {
	if x < 0 || x >= width || y < 0 || y >= height {
		return false
	}
	return true
}

func coordsKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
