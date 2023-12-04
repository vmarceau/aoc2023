package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	id   int
	want []int
	got  []int
}

func (c Card) matches() int {
	m := make(map[int]bool, len(c.want))
	for _, w := range c.want {
		m[w] = true
	}

	s := 0
	for _, g := range c.got {
		if !m[g] {
			continue
		}

		s += 1
	}

	return s
}

func parse(line string) (Card, error) {
	// Example:
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return Card{}, fmt.Errorf("Malformed card: %s", line)
	}

	id, err := strconv.Atoi(strings.TrimSpace(parts[0][4:]))
	if err != nil {
		return Card{}, fmt.Errorf("Malformed card: %s", line)
	}

	wantgot := strings.Split(parts[1], "|")
	if len(wantgot) != 2 {
		return Card{}, fmt.Errorf("Malformed card: %s", line)
	}

	want := strings.Fields(wantgot[0])
	got := strings.Fields(wantgot[1])
	card := Card{id: id, want: make([]int, 0, len(want)), got: make([]int, 0, len(got))}

	for _, w := range want {
		v, err := strconv.Atoi(w)
		if err != nil {
			return Card{}, fmt.Errorf("Malformed card: %s", line)
		}

		card.want = append(card.want, v)
	}

	for _, g := range got {
		v, err := strconv.Atoi(g)
		if err != nil {
			return Card{}, fmt.Errorf("Malformed card: %s", line)
		}

		card.got = append(card.got, v)
	}

	return card, nil
}

func main() {
	file, err := os.Open("data/04-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := make([]Card, 0)
	for scanner.Scan() {
		line := scanner.Text()

		card, err := parse(line)
		if err != nil {
			panic(err)
		}

		fmt.Println(line)
		cards = append(cards, card)
	}

	sum := 0
	replicas := make([]int, len(cards))
	for i, card := range cards {
		// Original
		replicas[i]++

		// Copies
		for j := i + 1; j < min(len(cards), i+1+card.matches()); j++ {
			replicas[j] += replicas[i]
		}

		sum += replicas[i]
	}

	fmt.Println("sum:", sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
