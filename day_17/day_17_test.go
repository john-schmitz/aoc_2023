package day_17

import "testing"

func TestPartOne(t *testing.T) {
	testCases := []struct {
		input_file string
		expected   int
	}{
		{
			input_file: "sample.txt",
			expected:   102,
		},
		{
			input_file: "sample2.txt",
			expected:   9,
		},
		{
			input_file: "sample3.txt",
			expected:   68,
		},
		{
			input_file: "input.txt",
			expected:   928,
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
			expected:   94,
		},
		{
			input_file: "input.txt",
			expected:   1104,
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
