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
	var starts []pair
	M := strings.Split(input, "\n")
	R := len(M)
	C := len(M[0])
	visited := make(map[string]bool)
	goals := make(map[string]map[string]bool)
	paths := make(map[string]int)
	directions := [][2]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	for r, line := range M {
		for c, val := range line {
			if int(val)-48 == 0 {
				starts = append(starts, pair{r, c})
			}
		}
	}
	totalPaths := 0
	totalGoals := 0
	for _, start := range starts {
		curR, curC, height := start.R, start.C, 0
		nextStep(visited, goals, paths, directions, curR, curC, M, height, R, C)
		key := coordsKey(curR, curC)
		totalGoals += len(goals[key])
		totalPaths += paths[key]
	}
	if part2 {
		return totalPaths
	}
	return totalGoals
}

func nextStep(visited map[string]bool, goals map[string]map[string]bool, paths map[string]int, directions [][2]int, curR int, curC int, M []string, height int, R int, C int) {
	key := coordsKey(curR, curC)
	if visited[key] {
		return
	}
	goals[key] = make(map[string]bool)
	if height == 9 {
		goals[key][key] = true
		paths[key] = 1
		visited[key] = true
		return
	}
	for _, direction := range directions {
		newR, newC := curR+direction[0], curC+direction[1]
		if isInMap(C, R, newC, newR) && int(M[newR][newC])-48 == height+1 {
			nextStep(visited, goals, paths, directions, newR, newC, M, height+1, R, C)
			nextKey := coordsKey(newR, newC)
			for k, _ := range goals[nextKey] {
				goals[key][k] = true
			}
			paths[key] += paths[nextKey]
		}
	}
	visited[key] = true
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
