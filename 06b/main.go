package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parse(line string) int {
	raw := strings.Split(line, ":")[1]
	raw = strings.ReplaceAll(raw, " ", "")

	val, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}

	return val
}

func solve(t, d int) int {
	// d = h * (t - h) where h = hold time
	// h**2 - t*h  + d = 0
	tf := float64(t)
	df := float64(d)
	h1 := 0.5 * (tf + math.Sqrt(math.Pow(tf, 2)-4*df))
	h2 := 0.5 * (tf - math.Sqrt(math.Pow(tf, 2)-4*df))

	hmax := int(math.Ceil(h1)) - 1
	hmin := int(math.Floor(h2)) + 1
	ways := max(hmax-hmin+1, 0)

	fmt.Printf("(t:%d, d:%d) => %d ways\n", t, d, ways)

	return ways
}

func main() {
	file, err := os.Open("data/06-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		panic("invalid input")
	}
	t := parse(scanner.Text())
	fmt.Println("time:", t)

	if !scanner.Scan() {
		panic("invalid input")
	}
	d := parse(scanner.Text())
	fmt.Println("distance:", d)

	ways := solve(t, d)

	fmt.Println("total ways:", ways)
}
