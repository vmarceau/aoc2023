package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	dst int
	src int
	len int
}

type Transform struct {
	ranges []Range
}

func (t Transform) Run(x int) int {
	for _, r := range t.ranges {
		if x >= r.src && x < r.src+r.len {
			return x - r.src + r.dst
		}
	}

	return x
}

func parse(scanner *bufio.Scanner) (Transform, error) {
	t := Transform{ranges: []Range{}}
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) == 0 {
			break
		}

		if len(fields) != 3 {
			return Transform{}, fmt.Errorf("invalid range: %v", line)
		}

		vals := []int{}
		for _, field := range fields {
			val, err := strconv.Atoi(field)
			if err != nil {
				return Transform{}, fmt.Errorf("invalid range: %v", line)
			}

			vals = append(vals, val)
		}

		r := Range{dst: vals[0], src: vals[1], len: vals[2]}
		t.ranges = append(t.ranges, r)

		// fmt.Println(line, r)
	}

	return t, nil
}

func main() {
	file, err := os.Open("data/05-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Seed numbers
	if !scanner.Scan() {
		panic("invalid input: no seeds")
	}

	seeds := []int{}
	seedline := scanner.Text()
	seedfields := strings.Fields(strings.Split(seedline, ":")[1])
	for _, field := range seedfields {
		seed, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}

		seeds = append(seeds, seed)
	}

	// fmt.Println(seedline, seeds)

	transforms := []Transform{}
	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, "map:") {
			continue
		}

		// fmt.Println(line)

		t, err := parse(scanner)
		if err != nil {
			panic(err)
		}

		transforms = append(transforms, t)
	}

	minloc := math.MaxInt
	locations := make([]int, 0, len(seeds))
	for _, seed := range seeds {
		loc := seed
		for _, t := range transforms {
			loc = t.Run(loc)
		}

		fmt.Println(seed, "->", loc)

		locations = append(locations, loc)

		if loc < minloc {
			minloc = loc
		}
	}

	fmt.Println("minloc:", minloc)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
