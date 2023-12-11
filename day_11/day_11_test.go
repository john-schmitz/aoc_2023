package day11

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
			expected:  374,
		},
		{
			file_path: "input.txt",
			expected:  9418609,
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

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		file_path string
		expected  int
	}{
		{
			file_path: "sample.txt",
			expected:  82000210,
		},
		{
			file_path: "input.txt",
			expected:  593821230983,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.file_path, func(t *testing.T) {
			actual := PartTwo(tC.file_path)
			if actual != tC.expected {
				t.Errorf("Expected PartTwo(%s) = %d. Got %d", tC.file_path, tC.expected, actual)
			}
		})
	}
}
