package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const R = 12
const G = 13
const B = 14

type Game struct {
	id int
	r  int
	g  int
	b  int
}

func parseGame(line string) (Game, error) {
	// Example:
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return Game{}, fmt.Errorf("Invalid game: %s", line)
	}

	if len(parts[0]) < 6 {
		return Game{}, fmt.Errorf("Invalid prefix: %s", parts[0])
	}

	id, err := strconv.Atoi(parts[0][5:])
	if err != nil {
		return Game{}, err
	}

	r, g, b := 0, 0, 0
	rounds := strings.Split(parts[1], ";")
	for _, round := range rounds {
		draws := strings.Split(round, ",")

		for _, draw := range draws {
			draw = strings.TrimSpace(draw)

			result := strings.Split(draw, " ")
			if len(result) != 2 {
				return Game{}, fmt.Errorf("Invalid draw: %s", draw)
			}

			n, err := strconv.Atoi(result[0])
			if err != nil {
				return Game{}, fmt.Errorf("Invalid draw: %s", draw)
			}

			color := result[1]
			switch color {
			case "red":
				if n > r {
					r = n
				}
			case "green":
				if n > g {
					g = n
				}
			case "blue":
				if n > b {
					b = n
				}
			}
		}
	}

	return Game{id: id, r: r, g: g, b: b}, nil
}

func main() {
	file, err := os.Open("data/02-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		game, err := parseGame(line)
		if err != nil {
			panic(err)
		}

		if game.r <= R && game.g <= G && game.b <= B {
			sum += game.id
		}

		fmt.Printf("%s: %+v\n", line, game)
	}

	fmt.Println("sum:", sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
