package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ASH rune = 46  // "."
var ROCK rune = 35 // "#"

var cache = map[string][]int{}

func flip(in string, i int) string {
	out := []rune(in)

	switch r := out[i]; r {
	case ASH:
		out[i] = ROCK
	case ROCK:
		out[i] = ASH
	}

	return string(out)
}

func check(line string) []int {
	if c, ok := cache[line]; ok {
		return c
	}

	chars := strings.Split(line, "")

	invalid := []int{}
	for i := 0; i < len(line)-1; i++ {
		l := ""
		r := ""
		for j := 0; j < len(line)/2+1; j++ {
			lidx := i - j
			if lidx < 0 {
				break
			}

			ridx := i + j + 1
			if ridx >= len(line) {
				break
			}

			l += chars[lidx]
			r += chars[ridx]
		}

		if l != r {
			invalid = append(invalid, i)
		}
	}

	cache[line] = invalid
	return invalid
}

func reflections(lines []string, skip int) int {
	valid := make([]bool, len(lines[0])-1)
	for i := 0; i < len(lines[0])-1; i++ {
		valid[i] = true
	}

	for _, line := range lines {
		invalid := check(line)
		for _, v := range invalid {
			valid[v] = false
		}
	}

	for i, v := range valid {
		if v && i != skip {
			return i
		}
	}

	return -1
}

func solve(lines, cols []string) int {
	l0 := reflections(lines, -1)
	c0 := reflections(cols, -1)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(cols); j++ {

			lines[i] = flip(lines[i], j)
			cols[j] = flip(cols[j], i)

			l := reflections(lines, l0)
			c := reflections(cols, c0)
			if l != -1 || c != -1 {
				return (l + 1) + 100*(c+1)
			}

			// undo
			lines[i] = flip(lines[i], j)
			cols[j] = flip(cols[j], i)
		}
	}

	panic("no solutions found")
}

func main() {
	file, err := os.Open("data/13-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	lines := []string{}
	cols := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			s := solve(lines, cols)
			fmt.Println(s)
			sum += s

			lines = []string{}
			cols = []string{}
			continue
		}

		lines = append(lines, line)
		if len(cols) == 0 {
			for _, c := range strings.Split(line, "") {
				cols = append(cols, c)
			}
		} else {
			for i, c := range strings.Split(line, "") {
				cols[i] += c
			}
		}

		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	s := solve(lines, cols)
	fmt.Println(s)
	sum += s

	fmt.Println("sum:", sum)
}
