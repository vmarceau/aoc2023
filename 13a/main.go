package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reflections(lines []string) int {
	valid := make([]bool, len(lines[0])-1)
	for i := 0; i < len(lines[0])-1; i++ {
		valid[i] = true
	}

	for _, line := range lines {
		chars := strings.Split(line, "")

		for i := range valid {
			if !valid[i] {
				continue
			}

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

			valid[i] = l == r
		}
	}

	sum := 0
	for i, v := range valid {
		if v {
			sum += i + 1
		}
	}

	return sum
}

func solve(lines, cols []string) int {
	return reflections(lines) + 100*reflections(cols)
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
