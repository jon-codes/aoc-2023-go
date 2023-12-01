package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	got := PartOne(input)
	want := 142

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	got := PartTwo(input)
	want := 281

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
