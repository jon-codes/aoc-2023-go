package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func parseSets(game string) []string {
	start := strings.Index(game, ":") + 2
	return strings.Split(game[start:], "; ")
}

func parseCube(cube string) (int, string) {
	parts := strings.Split(cube, " ")
	val, _ := strconv.Atoi(parts[0])
	color := parts[1]
	return val, color
}

func PartOne(data string) int {
	acc := 0
	id := 0

	allowed := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		id++
		valid := true

		sets := parseSets(scanner.Text())

		for _, set := range sets {
			for _, cube := range strings.Split(set, ", ") {
				val, color := parseCube(cube)
				if allowed[color] < val {
					valid = false
				}
			}
		}
		if valid {
			acc += id
		}
	}

	return acc
}

func PartTwo(data string) int {
	acc := 0

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		mins := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		sets := parseSets(scanner.Text())

		for _, set := range sets {
			for _, cube := range strings.Split(set, ", ") {
				val, color := parseCube(cube)
				if mins[color] < val {
					mins[color] = val
				}
			}
		}

		p := mins["red"] * mins["green"] * mins["blue"]
		acc += p
	}

	return acc
}

func main() {
	fmt.Println("Part 1:", PartOne(input))
	fmt.Println("Part 2:", PartTwo(input))
}
