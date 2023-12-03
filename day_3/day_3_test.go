package day3

import (
	"fmt"
	"testing"
)

const SAMPLE_PATH = "sample.txt"
const INPUT_PATH = "input.txt"

func TestPartOneSample(t *testing.T) {
	expected := 4361
	actual := SumPartNumbers(SAMPLE_PATH)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPartOne(t *testing.T) {
	expected := 556367
	actual := SumPartNumbers(INPUT_PATH)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPartTwoSample(t *testing.T) {
	expected := 467835
	actual := SumGearRatios(SAMPLE_PATH)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 467835
	actual := SumGearRatios(INPUT_PATH)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
func TestPointAdjacent(t *testing.T) {
	test_cases := []struct {
		point1   Point
		point2   Point
		expected bool
	}{
		{Point{0, 0}, Point{0, 0}, false},
		{Point{0, 1}, Point{0, 0}, true},
		{Point{1, 1}, Point{0, 0}, true},
		{Point{2, 2}, Point{0, 0}, false},
	}

	for _, test_case := range test_cases {
		actual := IsPointAdjacent(test_case.point1, test_case.point2)
		if actual != test_case.expected {
			fmt.Print(test_case)
			t.Errorf("Expected %t, got %t", test_case.expected, actual)
		}
	}
}

func TestPartTwoSample2(t *testing.T) {
	expected := 276
	actual := SumGearRatios("sample2.txt")
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetNumbers(t *testing.T) {
	board := parseInput("sample2.txt")
	expected := []Range{
		Range{Point{0, 0}, Point{0, 1}},
		Range{Point{0, 3}, Point{0, 4}},
	}

	actual := getNumbers(board)

	if len(actual) != len(expected) {
		t.Errorf("Expected %d, got %d", len(expected), len(actual))
	}

	for index, _ := range expected {
		if actual[index] != expected[index] {
			t.Errorf("Expected %v, got %v", expected[index], actual[index])
		}
	}

}
