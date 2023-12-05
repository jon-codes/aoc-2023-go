package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func readLines(data string) []string {
	lines := []string{}
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func PartOne(data string) int {
	acc := 0
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		score := 0

		line := scanner.Text()
		line = strings.Split(line, ": ")[1]
		parts := strings.Split(line, " | ")
		winning := strings.Fields(parts[0])
		have := strings.Fields(parts[1])

		for _, item := range have {
			for _, cand := range winning {
				if item == cand {
					if score == 0 {
						score = 1
					} else {
						score = score * 2
					}
					break
				}
			}
		}

		acc += score
	}

	return acc
}

func PartTwo(data string) int {
	acc := 0
	counts := map[int]int{}

	lines := readLines(data)
	for row := range lines {
		counts[row] = 1
	}

	for row, line := range lines {
		copies := counts[row]
		wins := 0

		line := strings.Split(line, ": ")[1]
		parts := strings.Split(line, " | ")
		winning := strings.Fields(parts[0])
		have := strings.Fields(parts[1])

		for _, item := range have {
			for _, cand := range winning {
				if item == cand {
					wins++
				}
			}
		}

		for i := row + 1; i <= row+wins && i < len(lines); i++ {
			counts[i] = counts[i] + copies
		}
	}

	for _, val := range counts {
		acc += val
	}

	return acc
}

func main() {
	fmt.Println("Part 1:", PartOne(input))
	fmt.Println("Part 2:", PartTwo(input))
}
