package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type CoOrds struct {
	X, Y int
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	antenna := make(map[string]bool)
	antinodes := make(map[string]bool)
	antennaGroups := make(map[rune][]CoOrds)
	M := strings.Split(input, "\n")
	width := len(M[0])
	height := len(M)
	for x, line := range M {
		for y, val := range line {
			if val != '.' {
				antenna[coordsKey(x, y)] = true
				antennaGroups[val] = append(antennaGroups[val], CoOrds{x, y})
				if part2 {
					antinodes[coordsKey(x, y)] = true
				}
			}
		}
	}
	for _, v := range antennaGroups {
		for l := 0; l < len(v)-1; l++ {
			for r := l + 1; r < len(v); r++ {
				xDiff, yDiff := v[l].X-v[r].X, v[l].Y-v[r].Y
				createAntinodes(v, l, xDiff, yDiff, width, height, antinodes, part2)
				createAntinodes(v, r, -xDiff, -yDiff, width, height, antinodes, part2)
			}
		}
	}

	return len(antinodes)
}

func createAntinodes(v []CoOrds, l int, xDiff int, yDiff int, width int, height int, antinodes map[string]bool, part2 bool) {
	x, y := v[l].X, v[l].Y
	for {
		x, y = x+xDiff, y+yDiff
		if isInMap(width, height, x, y) {
			antinodes[coordsKey(x, y)] = true
			if !part2 {
				break
			}
		} else {
			break
		}
	}
}

func coordsKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func isInMap(width int, height int, x int, y int) bool {
	if x < 0 || x >= width || y < 0 || y >= height {
		return false
	}
	return true
}
