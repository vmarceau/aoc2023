package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	HighCard int = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKing
)

var index = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

type Hand struct {
	cards []string
	score int
	bet   int
}

func parse(line string) Hand {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		panic("invalid input: " + line)
	}

	cards := strings.Split(fields[0], "")
	if len(cards) != 5 {
		panic("invalid input: " + line)
	}

	bet, _ := strconv.Atoi(fields[1])

	jokers := 0
	counts := make([]int, 13)
	for _, c := range cards {
		if c == "J" {
			jokers++
			continue
		}

		idx := index[c]
		counts[idx]++
	}
	slices.Sort(counts)
	slices.Reverse(counts)
	counts[0] += jokers

	score := HighCard
	if counts[0] == 5 {
		score = FiveOfAKing
	} else if counts[0] == 4 {
		score = FourOfAKind
	} else if counts[0] == 3 && counts[1] == 2 {
		score = FullHouse
	} else if counts[0] == 3 {
		score = ThreeOfAKind
	} else if counts[0] == 2 && counts[1] == 2 {
		score = TwoPairs
	} else if counts[0] == 2 {
		score = OnePair
	}

	return Hand{cards: cards, score: score, bet: bet}
}

func cmp(a, b Hand) int {
	if a.score < b.score {
		return -1
	}

	if a.score > b.score {
		return 1
	}

	for i := 0; i < 5; i++ {
		ca := index[a.cards[i]]
		cb := index[b.cards[i]]
		if ca < cb {
			return -1
		}

		if ca > cb {
			return 1
		}
	}

	return 0
}

func main() {
	file, err := os.Open("data/07-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		hand := parse(line)

		fmt.Println(line, hand)

		hands = append(hands, hand)
	}

	slices.SortFunc(hands, cmp)

	pts := 0
	for i, hand := range hands {
		pts += (i + 1) * hand.bet
	}

	fmt.Println("total:", pts)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
