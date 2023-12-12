package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const EXPANSION_FACTOR = 1e6

func dist(src, dst [2]int, ld, cd []int) int {
	d := 0

	lfrom, lto := min(src[0], dst[0]), max(src[0], dst[0])
	for l := lfrom + 1; l <= lto; l++ {
		d += ld[l]
	}

	cfrom, cto := min(src[1], dst[1]), max(src[1], dst[1])
	for c := cfrom + 1; c <= cto; c++ {
		d += cd[c]
	}

	return d
}

func main() {
	file, err := os.Open("data/11-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0
	g := [][2]int{}
	ld := []int{}
	cd := []int{}
	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, "")

		if i == 0 {
			n := len(tokens)
			ld = make([]int, n)
			cd = make([]int, n)
			for j := 0; j < n; j++ {
				ld[j] = EXPANSION_FACTOR
				cd[j] = EXPANSION_FACTOR
			}
		}

		for j, tok := range tokens {
			if tok == "#" {
				g = append(g, [2]int{i, j})
				ld[i] = 1
				cd[j] = 1
			}
		}

		i++
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sum := 0
	for i := 0; i < len(g); i++ {
		for j := i + 1; j < len(g); j++ {
			sum += dist(g[i], g[j], ld, cd)
		}
	}

	fmt.Println("sum:", sum)
}
