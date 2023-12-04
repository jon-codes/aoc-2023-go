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

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isBlank(r rune) bool {
	return r == '.'
}

type Position struct {
	row int
	col int
}

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
	parts := map[Position]int{}

	lines := readLines(data)

	for row, line := range lines {
		for col, char := range line {
			if isDigit(char) || isBlank(char) {
				continue
			}

			// check the surrounding 3-by grid for numbers
			for i := max(row-1, 0); i <= min(row+1, len(lines)-1); i++ {
				for j := max(col-1, 0); j <= min(col+1, len(line)-1); j++ {
					if isDigit(rune(lines[i][j])) {
						// look for the start of the number
						start := j
						for k := start; k >= 0; k-- {
							if isDigit(rune(lines[i][k])) {
								start = k
								continue
							}
							break
						}
						// seek to the end of the number
						str := ""
						for k := start; k < len(line); k++ {
							if isDigit(rune(lines[i][k])) {
								str = str + string(lines[i][k])
								continue
							}
							break
						}
						// we have the number
						val, _ := strconv.Atoi(str)
						pos := Position{row: i, col: start}
						parts[pos] = val
					}
				}
			}
		}
	}

	for _, val := range parts {
		acc += val
	}

	return acc
}

func PartTwo(data string) int {
	acc := 0

	lines := readLines(data)

	for row, line := range lines {
		for col, char := range line {
			if isDigit(char) || isBlank(char) {
				continue
			}

			parts := map[int]bool{}

			// check the surrounding 3-by grid for numbers
			for i := max(row-1, 0); i <= min(row+1, len(lines)-1); i++ {
				for j := max(col-1, 0); j <= min(col+1, len(line)-1); j++ {
					if isDigit(rune(lines[i][j])) {
						// look for the start of the number
						start := j
						for k := start; k >= 0; k-- {
							if isDigit(rune(lines[i][k])) {
								start = k
								continue
							}
							break
						}
						// seek to the end of the number
						str := ""
						for k := start; k < len(line); k++ {
							if isDigit(rune(lines[i][k])) {
								str = str + string(lines[i][k])
								continue
							}
							break
						}
						// we have the number
						val, _ := strconv.Atoi(str)
						parts[val] = true
					}
				}
			}

			if len(parts) == 2 {
				product := 1
				for val := range parts {
					product = product * val
				}
				acc += product
			}
		}
	}

	return acc
}

func main() {
	fmt.Println("Part 1:", PartOne(input))
	fmt.Println("Part 2:", PartTwo(input))
}
