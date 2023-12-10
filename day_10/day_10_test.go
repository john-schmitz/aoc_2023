package day10

import (
	"fmt"
	"testing"
)

func TestCanAccess(t *testing.T) {
	testCases := []struct {
		from       rune
		to         rune
		from_point Point
		to_point   Point
		expected   bool
	}{
		{
			from:       rune('S'),
			to:         rune('-'),
			from_point: Point{i: 0, j: 0},
			to_point:   Point{i: 0, j: 1},
			expected:   true,
		},
		{
			from:       rune('S'),
			to:         rune('-'),
			from_point: Point{i: 0, j: 0},
			to_point:   Point{i: 0, j: 2},
			expected:   false,
		},
		{
			from:       rune('S'),
			to:         rune('-'),
			from_point: Point{i: 0, j: 0},
			to_point:   Point{i: 1, j: 1},
			expected:   false,
		},
		{
			from:       rune('|'),
			to:         rune('-'),
			from_point: Point{i: 0, j: 0},
			to_point:   Point{i: 1, j: 0},
			expected:   false,
		},
		{
			from:       rune('|'),
			to:         rune('-'),
			from_point: Point{i: 0, j: 0},
			to_point:   Point{i: 1, j: 0},
			expected:   false,
		},
		{
			from:       rune('|'),
			to:         rune('|'),
			from_point: Point{i: 0, j: 0},
			to_point:   Point{i: 1, j: 0},
			expected:   true,
		},
		{
			from:       rune('|'),
			to:         rune('L'),
			from_point: Point{i: 0, j: 0},
			to_point:   Point{i: 1, j: 0},
			expected:   true,
		},
		{
			from:       rune('|'),
			to:         rune('L'),
			from_point: Point{i: 0, j: 0},
			to_point:   Point{i: 0, j: 1},
			expected:   false,
		},
		{
			from:       rune('L'),
			to:         rune('-'),
			from_point: Point{i: 0, j: 0},
			to_point:   Point{i: 0, j: 1},
			expected:   true,
		},
		{
			from:       rune('L'),
			to:         rune('-'),
			from_point: Point{i: 1, j: 1},
			to_point:   Point{i: 1, j: 0},
			expected:   false,
		},
	}
	for _, tC := range testCases {
		title := fmt.Sprintf("canAccess(%c,%c,%v,%v)", tC.from, tC.to, tC.from_point, tC.to_point)
		t.Run(title, func(t *testing.T) {
			actual := canAccess(tC.from, tC.to, tC.from_point, tC.to_point)

			if actual != tC.expected {
				t.Errorf("Expected canAccess(%c,%c,%v,%v) = %t. Got %t", tC.from, tC.to, tC.from_point, tC.to_point, tC.expected, actual)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	testCases := []struct {
		input_file string
		expected   int
	}{
		{
			input_file: "sample.txt",
			expected:   8,
		},
		{
			input_file: "input.txt",
			expected:   6640,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input_file, func(t *testing.T) {
			actual := PartOne(tC.input_file)

			if actual != tC.expected {
				t.Errorf("Expected PartOne(%s) = %d. Got %d", tC.input_file, tC.expected, actual)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		input_file string
		expected   int
	}{
		{
			input_file: "sample2.txt",
			expected:   10,
		},
		// {
		// 	input_file: "input.txt",
		// 	expected:   0,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.input_file, func(t *testing.T) {
			actual := PartTwo(tC.input_file)

			if actual != tC.expected {
				t.Errorf("Expected PartTwo(%s) = %d. Got %d", tC.input_file, tC.expected, actual)
			}
		})
	}
}
