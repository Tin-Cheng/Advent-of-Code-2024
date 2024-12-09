package main

import (
	"sort"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type pair struct {
	I, V int
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	var s []*string
	id := 0
	empty := "."
	originalIndex := 0
	var spaces []pair
	var data []pair
	var ids []string
	for pos, char := range input {
		val := int(char) - 48
		if pos%2 == 0 {
			data = append(data, pair{originalIndex, val})
			ids = append(ids, strconv.Itoa(id))
			for i := 0; i < val; i++ {
				s = append(s, &ids[id])
			}
			id += 1
		} else {
			spaces = append(spaces, pair{originalIndex, val})
			for i := 0; i < val; i++ {
				s = append(s, &empty)
			}
		}
		originalIndex += val
	}
	if !part2 {
		l, r := 0, len(s)-1
		for l < r {
			for *s[l] != "." {
				l += 1
			}
			for *s[r] == "." {
				r -= 1
			}
			if l < r {
				s[l], s[r] = s[r], s[l]
			}
		}
	} else {
		for i := len(data) - 1; i >= 0; i-- {
			dataI, size := data[i].I, data[i].V
			if size > 0 {
				if checkSpace(spaces, size, s, dataI) {
					newSpace := pair{dataI, size}
					insertSpace(spaces, newSpace)
				}
			}
		}
	}

	var checkSum int64 = 0
	for i, v := range s {
		if *v == "." {
			continue
		}
		val, err := strconv.Atoi(*v)
		if err != nil {
			panic("!")
		}
		sum := i * val
		checkSum += int64(sum)
	}
	return checkSum
}

func checkSpace(spaces []pair, size int, s []*string, dataI int) bool {
	for i, pair := range spaces {
		index, space := pair.I, pair.V
		if index > dataI {
			return false
		}
		if space >= size {
			for x := 0; x < size; x++ {
				s[dataI+x], s[index+x] = s[index+x], s[dataI+x]
			}
			spaces[i].I += size
			spaces[i].V -= size
			return true
		}
	}
	return false
}

func insertSpace(spaces []pair, space pair) {
	i := sort.Search(len(spaces), func(i int) bool {
		return spaces[i].I >= space.I
	})
	if i > 0 && spaces[i-1].I+spaces[i-1].V == space.I {
		spaces[i-1].V += space.V
	} else {
		spaces = append(spaces[:i], append([]pair{space}, spaces[i:]...)...)
	}
	if i < len(spaces)-1 && spaces[i].I+spaces[i].V == spaces[i+1].I {
		spaces[i].V += spaces[i+1].V
		spaces = append(spaces[:i+1], spaces[i+2:]...)
	}
}
