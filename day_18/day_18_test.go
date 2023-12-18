package day_18

import (
	"fmt"
	"testing"
)

func TestAreaPoints(t *testing.T) {
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

func Test(t *testing.T) {
	testCases := []struct {
		desc      string
		length    int
		direction string
	}{
		{
			desc:      "#70c710",
			length:    461937,
			direction: "R",
		},
		{
			desc:      "#0dc571",
			length:    56407,
			direction: "D",
		},
		{
			desc:      "#caa173",
			length:    829975,
			direction: "U",
		},
		{
			desc:      "#8ceee2",
			length:    577262,
			direction: "L",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			direction, length := convertHexToInstruction(tC.desc)
			if length != tC.length {
				t.Errorf("Expected convertHexToInstruction(%s) to have length %d. Got %d", tC.desc, tC.length, length)
			}

			if direction != tC.direction {
				t.Errorf("Expected convertHexToInstruction(%s) to have direction %s. Got %s", tC.desc, tC.direction, direction)
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
		{
			input_file: "input.txt",
			expected:   47675,
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
			expected:   952408144115,
		},
		{
			input_file: "input.txt",
			expected:   122103860427465,
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
