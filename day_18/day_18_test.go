package day_18

import (
	"fmt"
	"testing"
)



func Test(t *testing.T) {
	testCases := []struct {
		input    []Point
		expected int
	}{
		{
			expected: 16,
			input: []Point{
				{0, 0},
				{4, 0},
				{4, 4},
				{0, 4},
			},
		},
		{
			expected: 8,
			input: []Point{
				{0, 0},
				{4, 0},
				{2, 4},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("areaPoints(%v)", tC.input), func(t *testing.T) {
			actual := areaPoints(tC.input)
			if actual != tC.expected {
				t.Errorf("Expected areaPoints(%v) = %d. Got %d", tC.input, tC.expected, actual)
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
			expected:   62,
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
			input_file: "sample.txt",
			expected:   62,
		},
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
