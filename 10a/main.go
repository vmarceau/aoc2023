package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func parse(line string) ([]int, int, bool) {
	sok := false
	scol := 0
	vals := make([]int, 0, len(line))
	tiles := strings.Split(line, "")

	for i, tile := range tiles {
		var val int
		switch tile {
		case "|":
			val = 0b1100
		case "-":
			val = 0b0011
		case "L":
			val = 0b1010
		case "J":
			val = 0b1001
		case "7":
			val = 0b0101
		case "F":
			val = 0b0110
		case ".":
			val = 0b0000
		case "S":
			val = 0b1111
			sok = true
			scol = i
		default:
			panic("invalid char:" + tile)
		}

		vals = append(vals, val)
	}

	return vals, scol, sok
}

func solve(M [][]int, spos [2]int) int {
	for _, d := range [4]int{0b1000, 0b0100, 0b0010, 0b0001} {
		dir := d
		pos := [2]int{spos[0], spos[1]}
		steps := 0

		for dir > 0 {
			pos, dir = move(M, pos, dir)
			steps++

			if pos[0] == spos[0] && pos[1] == spos[1] {
				return steps
			}
		}
	}

	return 0
}

func move(M [][]int, posIn [2]int, dirIn int) ([2]int, int) {
	// Directions:
	// ^: 0b1000
	// v: 0b0100
	// >: 0b0010
	// <: 0b0001
	var posOut [2]int
	var dirOut int

	// Move north.
	if dirIn&0b1000 > 0 && posIn[0] > 0 {
		posOut[0] = posIn[0] - 1
		posOut[1] = posIn[1]
	}

	// Move south.
	if dirIn&0b0100 > 0 && posIn[0] < len(M)-1 {
		posOut[0] = posIn[0] + 1
		posOut[1] = posIn[1]
	}

	// Move east.
	if dirIn&0b0010 > 0 && posIn[1] < len(M[posIn[0]])-1 {
		posOut[0] = posIn[0]
		posOut[1] = posIn[1] + 1
	}

	// Move west.
	if dirIn&0b0001 > 0 && posIn[1] > 0 {
		posOut[0] = posIn[0]
		posOut[1] = posIn[1] - 1
	}

	// No move was made.
	if posOut[0] == posIn[0] && posOut[1] == posIn[1] {
		return posOut, dirOut
	}

	rdir := rev(dirIn)
	x := M[posOut[0]][posOut[1]]

	// Can't make the move because pipe doesn't enter tile.
	if x&rdir == 0 {
		return posOut, dirOut
	}

	dirOut = x &^ rdir

	return posOut, dirOut
}

func rev(dir int) int {
	switch dir {
	case 0b1000:
		return 0b0100
	case 0b0100:
		return 0b1000
	case 0b0010:
		return 0b0001
	case 0b0001:
		return 0b0010
	}

	return 0
}

func main() {
	file, err := os.Open("data/10-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	M := [][]int{}
	sline := 0
	spos := [2]int{}
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		vals, scol, sok := parse(line)

		if sok {
			spos[0] = sline
			spos[1] = scol
		}

		M = append(M, vals)

		sline++
	}

	fmt.Println("S:", spos)

	steps := solve(M, spos)
	fmt.Println("steps:", steps)
	fmt.Println("max:", math.Ceil(float64(steps)/2))

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
