package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parse(line string) []int {
	re := regexp.MustCompile(`-?\d+`)
	vals := []int{}

	for _, m := range re.FindAllString(line, -1) {
		val, _ := strconv.Atoi(m)
		vals = append(vals, val)
	}

	return vals
}

func solve(vals []int) int {
	order := 0
	pred := vals[len(vals)-1]
	diffs := [][]int{vals}

	for !done(diffs[order]) {
		n := len(diffs[order]) - 1

		diff := make([]int, 0, n)
		for i := 0; i < n; i++ {
			diff = append(diff, diffs[order][i+1]-diffs[order][i])
		}

		pred += diff[len(diff)-1]

		diffs = append(diffs, diff)
		order += 1
	}

	return pred
}

func done(xs []int) bool {
	for _, x := range xs {
		if x != 0 {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("data/09-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		vals := parse(line)

		pred := solve(vals)
		sum += pred

		fmt.Println(line, ":", pred)
	}

	fmt.Println("sum:", sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
