package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseInstructions(line string) []int {
	chars := strings.Split(line, "")

	instructions := make([]int, 0, len(chars))
	for _, c := range chars {
		switch c {
		case "L":
			instructions = append(instructions, 0)
		case "R":
			instructions = append(instructions, 1)
		}
	}

	return instructions
}

func parseNode(line string) (string, [2]string) {
	// eg. AAA = (BBB, CCC)
	re := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)

	matches := re.FindStringSubmatch(line)
	if len(matches) != 4 {
		panic("invalid line:" + line)
	}

	return matches[1], [2]string{matches[2], matches[3]}
}

func main() {
	file, err := os.Open("data/08-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	header := scanner.Text()
	instructions := parseInstructions(header)

	scanner.Scan()

	network := map[string][2]string{}
	for scanner.Scan() {
		line := scanner.Text()
		key, val := parseNode(line)

		network[key] = val
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(instructions)
	fmt.Println(network)

	key := "AAA"
	steps := 0
	N := len(instructions)
	for key != "ZZZ" {
		idx := instructions[steps%N]

		val, ok := network[key]
		if !ok {
			panic("invalid key:" + key)
		}

		key = val[idx]
		steps++
		fmt.Println(key, steps)
	}

	fmt.Println("steps:", steps)
}
