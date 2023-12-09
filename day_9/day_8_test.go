package day8

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLastValueHistory(t *testing.T) {
	testCases := []struct {
		history  []int
		expected int
	}{
		{
			history: []int{
				3, 3, 5, 9, 15,
			},
			expected: 5,
		},
		{
			history: []int{
				2, 2, 2,
			},
			expected: 2,
		},
		{
			history: []int{
				0, 2, 4, 6,
			},
			expected: -2,
		},
		{
			history: []int{
				10, 13, 16, 21, 30, 45,
			},
			expected: 5,
		},
		{
			history: []int{
				0, 3, 6, 9, 12, 15,
			},
			expected: -3,
		},
		{
			history: []int{
				1, 3, 6, 10, 15, 21,
			},
			expected: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprint(tC.history), func(t *testing.T) {
			actual := lastValueInHistory(tC.history)
			if actual != tC.expected {
				t.Errorf("got %d, want %d", actual, tC.expected)
			}
		})
	}
}

func TestNextValueHistory(t *testing.T) {
	testCases := []struct {
		history  []int
		expected int
	}{
		{
			history: []int{
				0, 3, 6, 9, 12, 15,
			},
			expected: 18,
		},
		{
			history: []int{
				1, 3, 6, 10, 15, 21,
			},
			expected: 28,
		},
		{
			history: []int{
				10, 13, 16, 21, 30, 45,
			},
			expected: 68,
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprint(tC.history), func(t *testing.T) {
			actual := nextValueInHistory(tC.history)
			if actual != tC.expected {
				t.Errorf("got %d, want %d", actual, tC.expected)
			}
		})
	}
}
func TestParseLine(t *testing.T) {
	testCases := []struct {
		line     string
		expected []int
	}{
		{
			line: "0 3 6 9 12 15",
			expected: []int{
				0, 3, 6, 9, 12, 15,
			},
		},
		{
			line: "1 3 6 10 15 21",
			expected: []int{
				1, 3, 6, 10, 15, 21,
			},
		},
		{
			line: "10 13 16 21 30 45",
			expected: []int{
				10, 13, 16, 21, 30, 45,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.line, func(t *testing.T) {
			actual := parseLine(tC.line)
			if !reflect.DeepEqual(actual, tC.expected) {
				t.Errorf("got %d, want %d", actual, tC.expected)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	testCases := []struct {
		input_path string
		expected   int
	}{
		{
			input_path: "sample.txt",
			expected:   114,
		},
		{
			input_path: "input.txt",
			expected:   1842168671,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input_path, func(t *testing.T) {
			actual := PartOne(tC.input_path)
			if actual != tC.expected {
				t.Errorf("got %d, want %d", actual, tC.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		input_path string
		expected   int
	}{
		{
			input_path: "sample.txt",
			expected:   2,
		},
		{
			input_path: "input.txt",
			expected:   12,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input_path, func(t *testing.T) {
			actual := PartTwo(tC.input_path)
			if actual != tC.expected {
				t.Errorf("got %d, want %d", actual, tC.expected)
			}
		})
	}
}
