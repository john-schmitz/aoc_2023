package day11

import (
	"slices"
	"testing"
)

func TestTransposeBoard(t *testing.T) {
	testCases := []struct {
		input_file string
		expected   []string
	}{
		{
			input_file: "sample.txt",
			expected: []string{
				"....#........",
				".........#...",
				"#............",
				".............",
				".............",
				"........#....",
				".#...........",
				"............#",
				".............",
				".............",
				".........#...",
				"#....#.......",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input_file, func(t *testing.T) {
			lines, _ := getLines(tC.input_file)
			actual := transposeBoard(lines)

			if len(actual) != len(tC.expected) {
				t.Errorf("len(actual) = %d. len(tC.expected) = %d", len(actual), len(tC.expected))
			}

			if len(actual[0]) != len(tC.expected[0]) {
				t.Errorf("len(actual[0]) = %d. len(tC.expected[0]) = %d", len(actual[0]), len(tC.expected[0]))
			}

			if !slices.Equal(actual, tC.expected) {
				t.Errorf("FAILED")
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	testCases := []struct {
		file_path string
		expected  int
	}{
		{
			file_path: "sample.txt",
			expected:  374,
		},
		// {
		// 	file_path: "input.txt",
		// 	expected:  374,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.file_path, func(t *testing.T) {
			actual := PartOne(tC.file_path)
			if actual != tC.expected {
				t.Errorf("Expected PartOne(%s) = %d. Got %d", tC.file_path, tC.expected, actual)
			}
		})
	}
}
