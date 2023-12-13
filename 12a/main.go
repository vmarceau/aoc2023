package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(line string) ([]int, []int, int, int) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		panic("invalid input:" + line)
	}

	dam := 0
	unk := 0
	vals := []int{}
	for _, tok := range strings.Split(parts[0], "") {
		val := 0
		switch tok {
		case "#":
			val = 1
			dam++
		case "?":
			val = 2
			unk++
		}

		vals = append(vals, val)
	}

	groups := []int{}
	for _, tok := range strings.Split(parts[1], ",") {
		group, err := strconv.Atoi(tok)
		if err != nil {
			panic("invalid token:" + tok)
		}

		groups = append(groups, group)
	}

	return vals, groups, dam, unk
}

func solve(vals []int, groups []int, dam, unk int) int {
	damcnt := 0
	for _, g := range groups {
		damcnt += g
	}

	solution := make([]int, unk)
	for i := 0; i < unk; i++ {
		if i < unk-(damcnt-dam) {
			solution[i] = 0
			continue
		}

		solution[i] = 1
	}

	ways := 0
	solutions := permutations(unk, damcnt-dam)
	for _, solution := range solutions {
		valid := evaluate(vals, groups, solution)
		if valid {
			ways++
		}
	}

	return ways
}

func permutations(unk, dam int) [][]int {
	if unk == 0 {
		return [][]int{}
	}

	if dam == 0 {
		solution := make([]int, unk)
		for i := 0; i < unk; i++ {
			solution[i] = 0
		}

		return [][]int{solution}
	}

	if dam == unk {
		solution := make([]int, unk)
		for i := 0; i < unk; i++ {
			solution[i] = 1
		}

		return [][]int{solution}
	}

	solutions := [][]int{}

	for _, sol := range permutations(unk-1, dam) {
		sol = append(sol, 0)
		solutions = append(solutions, sol)
	}

	for _, sol := range permutations(unk-1, dam-1) {
		sol = append(sol, 1)
		solutions = append(solutions, sol)
	}

	return solutions
}

func evaluate(vals []int, groups []int, solution []int) bool {
	gotvals := make([]int, len(vals))
	copy(gotvals, vals)

	k := 0
	count := 0
	gotgroups := []int{}
	for i := 0; i < len(gotvals); i++ {
		if gotvals[i] == 2 {
			gotvals[i] = solution[k]
			k++
		}

		if gotvals[i] == 1 {
			count++
			continue
		}

		if count > 0 {
			gotgroups = append(gotgroups, count)
			count = 0
		}
	}

	if count > 0 {
		gotgroups = append(gotgroups, count)
	}

	if len(groups) != len(gotgroups) {
		return false
	}

	for i := 0; i < len(groups); i++ {
		if groups[i] != gotgroups[i] {
			return false
		}
	}

	return true
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

		vals, groups, dam, unk := parse(line)

		ways := solve(vals, groups, dam, unk)
		sum += ways

		fmt.Println(line, ":", ways)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("sum:", sum)
}
