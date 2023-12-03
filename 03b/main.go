package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var DOT rune = 46  // "."
var GEAR rune = 42 // "*"

type Gear [2]int // [line, col]

type Part struct {
	number string
	pos    int
	gear   Gear
}

func isGearSymbol(r rune) bool {
	return r == GEAR
}

func getAttachedGear(part Part, prev, curr, next []rune, idx int) (Gear, bool) {
	jl := part.pos - 1
	jr := part.pos + len(part.number)

	check := func(part Part, runes []rune, idx int) ([2]int, bool) {
		if runes == nil {
			return Gear{}, false
		}

		for j := max(0, jl); j <= min(len(runes)-1, jr); j++ {
			r := runes[j]
			if isGearSymbol(r) {
				return Gear{idx, j}, true
			}
		}

		return Gear{}, false
	}

	if gear, ok := check(part, prev, idx-1); ok {
		return gear, ok
	}

	if gear, ok := check(part, next, idx+1); ok {
		return gear, ok
	}

	if jl >= 0 && isGearSymbol(curr[jl]) {
		return Gear{idx, jl}, true
	}

	if jr < len(curr) && isGearSymbol(curr[jr]) {
		return Gear{idx, jr}, true
	}

	return Gear{}, false
}

func extractPartsOfGear(prev, curr, next []rune, idx int) []Part {
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
		gear, ok := getAttachedGear(candidate, prev, curr, next, idx)
		if !ok {
			continue
		}

		candidate.gear = gear
		parts = append(parts, candidate)
	}

	return parts
}

func updateGearRegistry(parts []Part, gears map[string][]int) {
	for _, part := range parts {
		pn, err := strconv.Atoi(part.number)
		if err != nil {
			panic(err)
		}

		key := fmt.Sprintf("%d,%d", part.gear[0], part.gear[1])
		if pns, ok := gears[key]; ok {
			gears[key] = append(pns, pn)
		} else {
			gears[key] = []int{pn}
		}
	}
}

func main() {
	file, err := os.Open("data/03-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan plan using a buffer of 3 lines.
	// Extract and validate gear parts for the center line.
	gears := map[string][]int{}
	var prev, curr, next []rune

	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		prev = curr
		curr = next
		next = runes

		if curr == nil {
			continue
		}

		parts := extractPartsOfGear(prev, curr, next, idx)
		updateGearRegistry(parts, gears)
		fmt.Println(string(curr), parts)

		idx++
	}

	// Process last line.
	parts := extractPartsOfGear(curr, next, nil, idx)
	updateGearRegistry(parts, gears)
	fmt.Println(string(next), parts)

	// Sum gear ratios.
	sum := 0
	for _, pns := range gears {
		if len(pns) != 2 {
			continue
		}

		sum += pns[0] * pns[1]
	}

	fmt.Println("sum:", sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
