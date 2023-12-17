package day_16

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	testCases := []struct {
		input_file string
		expected   int
	}{
		{
			input_file: "sample.txt",
			expected:   46,
		},
		{
			input_file: "input.txt",
			expected:   7562,
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
			expected:   51,
		},
		{
			input_file: "input.txt",
			expected:   7793,
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
