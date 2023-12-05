package day5

import (
	"testing"
)

const sample_path = "sample.txt"
const input_path = "input.txt"

func TestPartOne(t *testing.T) {
	test_cases := []struct {
		file_path string
		expected  int
	}{
		{sample_path, 35},
		{input_path, 35},
	}

	for _, test_case := range test_cases {
		actual := PartOne(test_case.file_path)
		if actual != test_case.expected {
			t.Errorf("Expected %v, got %v", test_case.expected, actual)
		}
	}
}

func TestParseRanges(t *testing.T) {
	line := "50 98 2"
	expected := Ranges{50, 98, 2}
	actual, _ := parseRanges(line)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestReadSeedsFromInput(t *testing.T) {
	line := "seeds: 79 14 55 13 2"
	expected := []int{79, 14, 55, 13, 2}
	actual := readSeedsFromInput(line)
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestMapRange(t *testing.T) {
	test_cases := []struct {
		id     int
		range_ Ranges
		output int
	}{
		{99, Ranges{50, 98, 2}, 51},
		{98, Ranges{50, 98, 2}, 50},
		{53, Ranges{52, 50, 48}, 55},
		{13, Ranges{52, 50, 48}, 13},
	}

	for _, test_case := range test_cases {
		actual := mapRange(test_case.id, test_case.range_)
		if actual != test_case.output {
			t.Errorf("Expected %v, got %v", test_case.output, actual)
		}
	}
}
