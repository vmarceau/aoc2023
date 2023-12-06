package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func parse(line string) []int {
	re := regexp.MustCompile(`\d+`)
	vals := []int{}

	for _, m := range re.FindAllString(line, -1) {
		val, _ := strconv.Atoi(m)
		vals = append(vals, val)
	}

	return vals
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
	times := parse(scanner.Text())
	fmt.Println("times:", times)

	if !scanner.Scan() {
		panic("invalid input")
	}
	distances := parse(scanner.Text())
	fmt.Println("distances:", distances)

	if len(times) != len(distances) {
		panic("invalid input")
	}

	ways := 1
	for i := 0; i < len(times); i++ {
		ways *= solve(times[i], distances[i])
	}

	fmt.Println("total ways:", ways)
}
