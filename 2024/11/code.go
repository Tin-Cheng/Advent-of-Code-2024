package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//aoc.Harness(run)
	exampleByte, err1 := os.ReadFile("input-example.txt")
	if err1 != nil {
		panic("example err")
	}
	example := string(exampleByte)
	userInputByte, err2 := os.ReadFile("input-user.txt")
	if err2 != nil {
		panic("user input err")
	}
	userInput := string(userInputByte)
	run(false, example)
	run(false, userInput)
	run(true, example)
	run(true, userInput)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	result := 0
	note := make(map[string]int)
	blinks := 25
	if part2 {
		blinks = 75
	}
	for _, v := range strings.Split(input, " ") {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic("v err")
		}
		result += letsBlink(note, i, blinks)
	}
	println(part2, result)
	return "end"
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func letsBlink(note map[string]int, v, blinks int) int {
	if blinks == 0 {
		return 1
	}
	if note[key(v, blinks)] > 0 {
		return note[key(v, blinks)]
	}
	result := 0
	if v == 0 {
		result = letsBlink(note, 1, blinks-1)
	} else {
		s := strconv.Itoa(v)
		if len(s)%2 == 0 {
			l := len(s) / 2
			s1, s2 := s[:l], s[l:]
			v1, err1 := strconv.Atoi(s1)
			if err1 != nil {
				panic("pt1 err")
			}
			v2, err2 := strconv.Atoi(s2)
			if err2 != nil {
				panic("pt2 err")
			}
			result += letsBlink(note, v1, blinks-1) + letsBlink(note, v2, blinks-1)
		} else {
			result += letsBlink(note, v*2024, blinks-1)
		}
	}
	note[key(v, blinks)] = result
	return result
}

/*

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func run(part2 bool, input string, blinks int) any {
	list := LinkedList{}
	head := Node{0, nil}
	list.head = &head
	generateList(input, list.head)
	for i := 0; i < blinks; i++ {
		pt := list.head.next
		for pt != nil {
			if pt.data == 0 {
				pt.data = 1
			} else if len(strconv.Itoa(pt.data))%2 == 0 {
				s := strconv.Itoa(pt.data)
				l := len(s) / 2
				s1, s2 := s[:l], s[l:]
				v1, err1 := strconv.Atoi(s1)
				if err1 != nil {
					panic("pt1 err")
				}
				pt.data = v1
				v2, err2 := strconv.Atoi(s2)
				if err2 != nil {
					panic("pt1 err")
				}
				newNode := Node{v2, pt.next}
				pt.next = &newNode
				pt = pt.next
			} else {
				pt.data = pt.data * 2024
			}
			pt = pt.next
		}
	}
	println(countNodes(&list))
	return "end"
}

func generateList(input string, pt *Node) {
	for _, v := range strings.Split(input, " ") {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic("!")
		}
		n := Node{i, nil}
		pt.next = &n
		pt = pt.next
	}
}

func countNodes(list *LinkedList) (count int) {
	current := list.head
	for current != nil {
		current = current.next
		count++
	}
	count--
	return count
}
*/
