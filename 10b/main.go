package main

import (
	"bufio"
	"fmt"
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

func solve(M [][]int, spos [2]int) [][2]int {
	for _, d := range [4]int{0b1000, 0b0100, 0b0010, 0b0001} {
		dir := d
		pos := [2]int{spos[0], spos[1]}
		contour := [][2]int{pos}

		for dir > 0 {
			npos, ndir := move(M, pos, dir)
			contour = append(contour, pos)

			if npos[0] == spos[0] && npos[1] == spos[1] {
				// Set value for starting point.
				M[spos[0]][spos[1]] = d | rev(dir)
				return contour
			}

			pos = npos
			dir = ndir
		}
	}

	return [][2]int{}
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

func inner(M [][]int, contour [][2]int) int {
	// Register contour.
	C := map[string]int{}
	for i := 0; i < len(contour); i++ {
		C[key(contour[i])] = M[contour[i][0]][contour[i][1]]
	}

	// Ray casting, west.
	area := 0
	for i := 0; i < len(M); i++ {
		edges := 0
		for j := 0; j < len(M[0]); j++ {
			// Point is on contour.
			if c, ok := C[key([2]int{i, j})]; ok {
				// Vertical edge crossed.
				if c&0b1000 > 0 {
					edges++
				}

				continue
			}

			if int(edges)%2 == 1 {
				area += 1
			}
		}
	}

	return area
}

func key(p [2]int) string {
	return fmt.Sprintf("%d:%d", p[0], p[1])
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

	c := solve(M, spos)
	a := inner(M, c)

	fmt.Println("inner:", a)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
