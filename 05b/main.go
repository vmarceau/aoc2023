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

func (t Transform) tr(x int, r Range) int {
	return x - r.src + r.dst
}

// Run transforms an input range into a list of output ranges.
func (t Transform) Run(x [2]int) [][2]int {
	curr := [2]int{}
	out := [][2]int{}
	rem := [][2]int{x}

	for len(rem) > 0 {
		curr, rem = rem[0], rem[1:]

		found := false
		for _, r := range t.ranges {
			// No overlap.
			if curr[0]+curr[1] <= r.src || curr[0] >= r.src+r.len {
				continue
			}

			// Full overlap.
			if curr[0] >= r.src && curr[0]+curr[1] <= r.src+r.len {
				out = append(out, [2]int{t.tr(curr[0], r), curr[1]})
				found = true
				break
			}

			// Partial left overlap.
			if curr[0] >= r.src {
				overlap := r.src + r.len - curr[0]
				out = append(out, [2]int{t.tr(curr[0], r), overlap})
				rem = append(rem, [2]int{curr[0] + overlap, curr[1] - overlap})
				found = true
				break
			}

			// Partial right overlap.
			if curr[0]+curr[1] <= r.src+r.len {
				overlap := curr[0] + curr[1] - r.src
				out = append(out, [2]int{t.tr(r.src, r), overlap})
				rem = append(rem, [2]int{curr[0], curr[1] - overlap})
				found = true
				break
			}

			// Center overlap
			out = append(out, [2]int{t.tr(r.src, r), r.len})
			rem = append(rem, [2]int{curr[0], r.src - curr[0]})
			rem = append(rem, [2]int{r.src + r.len, curr[0] + curr[1] - r.src - r.len})
			found = true
			break
		}

		if !found {
			out = append(out, curr)
		}
	}

	return out
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

func minimize(seed [2]int, transforms []Transform) int {
	in := [][2]int{seed}
	out := [][2]int{}
	for _, t := range transforms {
		for _, r := range in {
			out = append(out, t.Run(r)...)
		}

		in = out
		out = [][2]int{}
	}

	minloc := math.MaxInt
	for _, r := range in {
		if r[0] < minloc {
			minloc = r[0]
		}
	}

	fmt.Printf("minloc [%v]: %d\n", seed, minloc)

	return minloc
}

func main() {
	file, err := os.Open("data/05-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Seed ranges
	if !scanner.Scan() {
		panic("invalid input: no seed ranges")
	}

	seeds := [][2]int{}
	seedline := scanner.Text()
	seedfields := strings.Fields(strings.Split(seedline, ":")[1])
	if len(seedfields)%2 != 0 {
		panic("invalid input")
	}

	curr := [2]int{}
	for i, field := range seedfields {
		val, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			curr[0] = val
			continue
		}

		curr[1] = val
		seeds = append(seeds, curr)
	}

	fmt.Println(seedline, seeds)

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
	for _, seed := range seeds {
		loc := minimize(seed, transforms)
		if loc < minloc {
			minloc = loc
		}
	}

	fmt.Println("minloc:", minloc)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
