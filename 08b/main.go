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

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	if len(values) == 1 {
		return values[0]
	}

	result := values[0] * values[1] / gcd(values[0], values[1])

	for i := 2; i < len(values); i++ {
		result = lcm(result, values[i])
	}

	return result
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

	keys := []string{}
	network := map[string][2]string{}
	for scanner.Scan() {
		line := scanner.Text()
		key, val := parseNode(line)

		network[key] = val

		if strings.HasSuffix(key, "A") {
			keys = append(keys, key)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(instructions)
	fmt.Println(network)
	fmt.Println(keys)

	N := len(instructions)
	steps := make([]int, 0, len(keys))
	for _, key := range keys {
		k := key
		step := 0

		for !strings.HasSuffix(k, "Z") {
			idx := instructions[step%N]

			val, ok := network[k]
			if !ok {
				panic("invalid key:" + k)
			}

			k = val[idx]
			step++
		}

		steps = append(steps, step)
	}

	fmt.Println("steps:", steps)
	fmt.Println("lcm:", lcm(steps...))
}
