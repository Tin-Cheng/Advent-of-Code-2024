package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"fmt"
	"strings"
	"strconv"
	"math"
	"slices"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	result := 0.0

	var s1 []float64
	var s2 []float64
	for _, line := range strings.Split(strings.TrimSpace(input), "\n"){
		pair := strings.SplitN(line, " ", 2)
		s1 = append(s1, toInt(pair[0]))
		s2 = append(s2, toInt(pair[1]))
	}
	slices.Sort(s1)
	slices.Sort(s2)
	if part2 {
		m1 := make(map[float64]int)
		for i := range s1 {
			m1[s1[i]] += 1
		}
		m2 := make(map[float64]int)
		for i := range s2 {
			m2[s2[i]] += 1
		}
		for k, v := range m1{
			result += k * float64(m2[k] * v)
		}		
		fmt.Printf("%f\n", result)
		return result
	}
	if len(s1) == len(s2) {
        for i := range s1 {
            result += math.Abs(s1[i] - s2[i])
        }
    }
	// solve part 1 here
	
	fmt.Printf("%f\n", result)
	return result
}

func toInt(s string) float64 {
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return float64(i)
}
