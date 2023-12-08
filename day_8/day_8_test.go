package day8

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	testCases := []struct {
		file_path string
		expected  int
	}{
		{
			file_path: "sample.txt",
			expected:  6,
		},
		{
			file_path: "input.txt",
			expected:  2,
		},
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

func TestParseLine(t *testing.T) {
	testCases := []struct {
		input         string
		expected_key  string
		expected_node Node
	}{
		{
			input:         "AAA = (BBB, CCC)",
			expected_key:  "AAA",
			expected_node: Node{left: "BBB", right: "CCC"},
		},
		{
			input:         "XXX = (BBB, DDD)",
			expected_key:  "XXX",
			expected_node: Node{left: "BBB", right: "DDD"},
		},
		{
			input:         "ZZZ = (BBB, CCC)",
			expected_key:  "ZZZ",
			expected_node: Node{left: "BBB", right: "CCC"},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			key, actual := parseLine(tC.input)
			if actual != tC.expected_node || key != tC.expected_key {
				t.Errorf("Expected parseLine(%s) = %v. Got %v", tC.input, tC.expected_node, actual)
			}
		})
	}
}
