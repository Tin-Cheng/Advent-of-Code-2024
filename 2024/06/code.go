package main

import (
	"fmt"
	"maps"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func run(part2 bool, input string) any {
	// solve part 1 here
	visited := make(map[string]bool)
	part2Blocks := make(map[string]bool)
	M := strings.Split(input, "\n")
	width, height := len(M[0]), len(M)
	var startX, startY, direction int

	for i, line := range M {
		index := strings.Index(line, "^")
		if index >= 0 {
			startX, startY = i, index
			break
		}
	}
	simulate(part2, true, width, height, startX, startY, direction, visited, M, part2Blocks, -999, -999)

	if part2 {
		positions := len(part2Blocks)
		return positions
	}

	return len(visited)
}

func simulate(part2 bool, firstLevel bool, width int, height int, x int, y int, d int, visited map[string]bool, M []string, part2Blocks map[string]bool, extraX int, extraY int) {
	if part2Blocks[strconv.Itoa(extraX)+","+strconv.Itoa(extraY)] {
		return
	}
	localVisited := make(map[string]bool)
	path := make(map[string]bool)
	for {
		key := strconv.Itoa(x) + "," + strconv.Itoa(y)
		pKey := strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(d)
		if path[pKey] {
			part2Blocks[strconv.Itoa(extraX)+","+strconv.Itoa(extraY)] = true
			return
		}
		localVisited[key] = true
		path[pKey] = true
		nextX, nextY := x+directions[d][0], y+directions[d][1]

		if isInMap(width, height, nextX, nextY) {
			if part2 && firstLevel && M[nextX][nextY] != '#' && !localVisited[coordsKey(nextX, nextY)] {
				simulate(part2, false, width, height, x, y, d, visited, M, part2Blocks, nextX, nextY)
			}
			if M[nextX][nextY] == '#' || (nextX == extraX && nextY == extraY) {
				d = (d + 1) % 4
			} else {
				x, y = nextX, nextY
			}
		} else {
			if !part2 {
				maps.Copy(visited, localVisited)
			}
			return
		}
	}
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

func coordsKeyWithDirection(x, y, direction int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, direction)
}
