package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

var digits = [][]rune{
	[]rune("one"),
	[]rune("two"),
	[]rune("three"),
	[]rune("four"),
	[]rune("five"),
	[]rune("six"),
	[]rune("seven"),
	[]rune("eight"),
	[]rune("nine"),
}

var stigid = [][]rune{
	[]rune("eno"),
	[]rune("owt"),
	[]rune("eerht"),
	[]rune("ruof"),
	[]rune("evif"),
	[]rune("xis"),
	[]rune("neves"),
	[]rune("thgie"),
	[]rune("enin"),
}

func isLiteralDigit(word []rune, digits [][]rune) (int, bool) {
	candidates := []int{0, 1, 2, 3, 4, 5, 6, 7, 8} // correspond to literal digit - 1
	matches := make([]int, 0, len(candidates))

	for i := 0; i < len(word); i++ {
		if len(candidates) == 0 {
			break
		}

		for _, candidate := range candidates {
			// i-th letter doesn't match
			if word[i] != digits[candidate][i] {
				continue
			}

			// all letters match
			if i == len(digits[candidate])-1 {
				return candidate + 1, true
			}

			matches = append(matches, candidate)
		}

		candidates = matches
		matches = make([]int, 0, len(candidates))
	}

	return 0, false
}

func main() {
	file, err := os.Open("data/01-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	var first, last string

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				first = string(runes[i])
				break
			}

			word := runes[i:min(i+5, len(runes)-1)]
			if d, ok := isLiteralDigit(word, digits); ok {
				first = fmt.Sprintf("%d", d)
				break
			}
		}

		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				last = string(runes[i])
				break
			}

			word := make([]rune, min(i+1, 5))
			copy(word, runes[max(i-4, 0):i+1])
			slices.Reverse(word)
			if d, ok := isLiteralDigit(word, stigid); ok {
				last = fmt.Sprintf("%d", d)
				break
			}
		}

		code, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}

		sum += code

		fmt.Printf("%v: %v%v\n", line, first, last)
	}

	fmt.Println("sum:", sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
