package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

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
		}

		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				last = string(runes[i])
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
