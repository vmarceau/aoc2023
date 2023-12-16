package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var M map[string]int = map[string]int{}

func key(v []string, g []int, f bool) string {
	return fmt.Sprintf("%v:%v:%v", v, g, f)
}

func parse(line string) ([]string, []int) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		panic("invalid input:" + line)
	}

	// Unfold
	parts[0] = strings.Join([]string{parts[0], parts[0], parts[0], parts[0], parts[0]}, "?")
	parts[1] = strings.Join([]string{parts[1], parts[1], parts[1], parts[1], parts[1]}, ",")

	vals := strings.Split(parts[0], "")

	groups := []int{}
	for _, tok := range strings.Split(parts[1], ",") {
		group, err := strconv.Atoi(tok)
		if err != nil {
			panic("invalid token:" + tok)
		}

		groups = append(groups, group)
	}

	return vals, groups
}

func solve(v []string, g []int, f bool) int {
	k := key(v, g, f)
	if s, ok := M[k]; ok {
		return s
	}

	s := 0
	v0, v := v[0], v[1:]

	gc := make([]int, len(g))
	copy(gc, g)

	switch v0 {
	case ".":
		if len(v) == 0 {
			if len(gc) == 1 && gc[0] == 0 {
				s = 1
			}

			break
		}

		if f && g[0] > 0 {
			break
		}

		if gc[0] == 0 && len(gc) > 1 {
			gc = gc[1:]
		}

		s = solve(v, gc, false)

	case "#":
		if gc[0] <= 0 {
			break
		}

		gc[0]--

		if len(v) == 0 {
			if len(gc) == 1 && gc[0] == 0 {
				s = 1
			}

			break
		}

		s = solve(v, gc, true)

	case "?":
		if len(v) == 0 {
			if len(gc) == 1 && (gc[0] == 0 || gc[0] == 1) {
				s = 1
			}

			break
		}

		sd := 0
		if gc[0] > 0 {
			gd := make([]int, len(gc))
			copy(gd, gc)

			gd[0]--
			sd = solve(v, gd, true)
		}

		snd := 0
		if !f || f && gc[0] == 0 {
			if gc[0] == 0 && len(gc) > 1 {
				gc = gc[1:]
			}

			snd = solve(v, gc, false)
		}

		s = sd + snd
	}

	M[k] = s
	return s
}

func main() {
	file, err := os.Open("data/12-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		vals, groups := parse(line)

		ways := solve(vals, groups, false)
		sum += ways

		fmt.Println(line, ":", ways)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("sum:", sum)
}
