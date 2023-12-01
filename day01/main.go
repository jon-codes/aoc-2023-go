package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func PartOne(data string) int {
	acc := 0

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		var first, last *int

		for _, r := range scanner.Text() {
			if isDigit(r) {
				digit := int(r - '0')
				if first == nil {
					first = &digit
				}
				last = &digit
			}
		}

		acc += (*first * 10) + *last
	}

	return acc
}

func PartTwo(data string) int {
	acc := 0

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()

		var first, last *int

		for start, r := range line {
			var digit int
			if isDigit(r) {
				digit = int(r - '0')
			} else {
				rest := line[start:]

				if strings.HasPrefix(rest, "one") {
					digit = 1
				} else if strings.HasPrefix(rest, "two") {
					digit = 2
				} else if strings.HasPrefix(rest, "three") {
					digit = 3
				} else if strings.HasPrefix(rest, "four") {
					digit = 4
				} else if strings.HasPrefix(rest, "five") {
					digit = 5
				} else if strings.HasPrefix(rest, "six") {
					digit = 6
				} else if strings.HasPrefix(rest, "seven") {
					digit = 7
				} else if strings.HasPrefix(rest, "eight") {
					digit = 8
				} else if strings.HasPrefix(rest, "nine") {
					digit = 9
				}
			}
			if digit != 0 {
				if first == nil {
					first = &digit
				}
				last = &digit
			}
		}

		acc += (*first * 10) + *last
	}

	return acc
}

func main() {
	fmt.Println("Part 1:", PartOne(input))
	fmt.Println("Part 2:", PartTwo(input))
}
