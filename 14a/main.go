package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var NONE rune = 46   // "."
var SQUARE rune = 35 // "#"
var ROUND rune = 79  // "O"

func rev(a, b rune) int {
	return int(b - a)
}

func roll(line string) string {
	parts := strings.Split(line, "#")
	for i, part := range parts {
		if len(part) == 0 {
			continue
		}

		runes := []rune(part)
		slices.SortFunc(runes, rev)
		parts[i] = string(runes)
	}

	return strings.Join(parts, "#")
}

func weight(line string) int {
	weight := 0
	for i, r := range line {
		if r == ROUND {
			weight += len(line) - i
		}
	}

	return weight
}

func solve(lines []string) int {
	w := 0
	for _, col := range lines {
		rolled := roll(col)
		w += weight(rolled)
	}

	return w
}

func main() {
	file, err := os.Open("data/14-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cols := []string{}
	for scanner.Scan() {
		line := scanner.Text()

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

	sum := solve(cols)
	fmt.Println("sum:", sum)
}
