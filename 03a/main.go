package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var DOT rune = 46 // "."

type Part struct {
	number string
	pos    int
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != DOT
}

func isPartValid(part Part, prev, curr, next []rune) bool {
	jl := part.pos - 1
	jr := part.pos + len(part.number)

	check := func(part Part, runes []rune) bool {
		if runes == nil {
			return false
		}

		for j := max(0, jl); j <= min(len(runes)-1, jr); j++ {
			r := runes[j]
			if isSymbol(r) {
				return true
			}
		}

		return false
	}

	if check(part, prev) {
		return true
	}

	if check(part, next) {
		return true
	}

	if jl >= 0 && isSymbol(curr[jl]) {
		return true
	}

	if jr < len(curr) && isSymbol(curr[jr]) {
		return true
	}

	return false
}

func extractParts(prev, curr, next []rune) []Part {
	found := false
	candidates := []Part{}
	candidate := Part{}

	// Extract candidate parts.
	for j, r := range curr {
		// This is a candidate.
		if unicode.IsDigit(r) {
			if !found {
				found = true
				candidate.pos = j
			}

			candidate.number += string(r)

			continue
		}

		// This is not a candidate, reset if needed.
		if found {
			candidates = append(candidates, candidate)
			candidate = Part{}
			found = false
		}
	}

	if found {
		candidates = append(candidates, candidate)
	}

	// Filter out invalid parts.
	parts := make([]Part, 0, len(candidates))
	for _, candidate := range candidates {
		if !isPartValid(candidate, prev, curr, next) {
			continue
		}

		parts = append(parts, candidate)
	}

	return parts
}

func sumPartNumbers(parts []Part) int {
	sum := 0
	for _, part := range parts {
		pn, err := strconv.Atoi(part.number)
		if err != nil {
			panic(err)
		}

		sum += pn
	}

	return sum
}

func main() {
	file, err := os.Open("data/03-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan plan using a buffer of 3 lines.
	// Extract and validate parts for the center line.
	sum := 0
	var prev, curr, next []rune

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		prev = curr
		curr = next
		next = runes

		if curr == nil {
			continue
		}

		parts := extractParts(prev, curr, next)
		sum += sumPartNumbers(parts)
		fmt.Println(string(curr), parts)
	}

	// Process last line.
	parts := extractParts(curr, next, nil)
	sum += sumPartNumbers(parts)
	fmt.Println(string(next), parts)

	fmt.Println("sum:", sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
