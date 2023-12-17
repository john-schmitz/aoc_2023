package day_16

import (
	"fmt"
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

func TestEnergized(t *testing.T) {
	testCases := []struct {
		lines    []string
		expected int
	}{
		{
			expected: 15,
			lines: []string{
				"....|.....",
				"....-.....",
			},
		},
		{
			expected: 10,
			lines: []string{
				"..........",
			},
		},
		{
			expected: 11,
			lines: []string{
				".........|",
				"..........",
			},
		},
		{
			expected: 20,
			lines: []string{
				".........|",
				".........-",
			},
		},
		{
			expected: 20,
			lines: []string{
				".........|",
				".........-",
			},
		},
		{
			expected: 11,
			lines: []string{
				".|.........",
				".-....|....",
				"...........",
			},
		},
		{
			expected: 11,
			lines: []string{
				"\\.........",
				"-.........",
			},
		},
	}

	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%s", tC.lines), func(t *testing.T) {
			actual := energized(tC.lines)
			if actual != tC.expected {
				t.Errorf("Expected energized(%s) = %d. Got %d", tC.lines, tC.expected, actual)
			}
		})
	}
}
