package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	got := PartOne(input)
	want := 4361

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	got := PartTwo(input)
	want := 467835

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
