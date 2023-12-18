package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"slices"
	"strings"
)

var NONE rune = 46   // "."
var SQUARE rune = 35 // "#"
var ROUND rune = 79  // "O"

type State struct {
	h uint64
	w int
}

var lineCache = map[string]string{}
var stateCache = map[uint64]State{}

func rev(a, b rune) int {
	return int(b - a)
}

func lineKey(line string, flip bool) string {
	return fmt.Sprintf("%s:%v", line, flip)
}

func hashState(lines []string) uint64 {
	h := fnv.New64()
	for _, line := range lines {
		h.Write([]byte(line))
	}

	return h.Sum64()
}

func roll(lines []string, flip bool) {
	for i, line := range lines {
		if c, ok := lineCache[lineKey(line, flip)]; ok {
			lines[i] = c
			continue
		}

		parts := strings.Split(line, "#")
		for i, part := range parts {
			if len(part) == 0 {
				continue
			}

			runes := []rune(part)

			if flip {
				slices.Sort(runes)
			} else {
				slices.SortFunc(runes, rev)
			}

			parts[i] = string(runes)
		}

		rolled := strings.Join(parts, "#")
		lineCache[lineKey(line, flip)] = rolled

		lines[i] = rolled
	}
}

func weight(lines []string) int {
	weight := 0
	for _, line := range lines {
		for i, r := range line {
			if r == ROUND {
				weight += len(line) - i
			}
		}
	}

	return weight
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

	h := hashState(cols)
	w := 0
	for cycle := 0; cycle < 1000000000-1; cycle++ {
		if cycle%10000000 == 0 {
			fmt.Println("cycle:", cycle, "h:", h)
		}

		if c, ok := stateCache[h]; ok {
			h = c.h
			continue
		}

		half := func(ini []string, flip bool) []string {
			// North (flip=false), South (flip=true)
			roll(cols, flip)

			lines := []string{}
			for _, col := range cols {
				if len(lines) == 0 {
					for _, c := range strings.Split(col, "") {
						lines = append(lines, c)
					}
				} else {
					for i, c := range strings.Split(col, "") {
						lines[i] += c
					}
				}
			}

			// West (flip=false), East (flip=true)
			roll(lines, flip)

			cols = []string{}
			for _, line := range lines {
				if len(cols) == 0 {
					for _, c := range strings.Split(line, "") {
						cols = append(cols, c)
					}
				} else {
					for i, c := range strings.Split(line, "") {
						cols[i] += c
					}
				}
			}

			return cols
		}

		cols = half(cols, false)
		cols = half(cols, true)

		nh := hashState(cols)
		w = weight(cols)

		stateCache[h] = State{h: nh, w: w}
		h = nh
	}

	final := stateCache[h]
	fmt.Println("sum:", final.w)
}
