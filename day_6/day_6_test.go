package day6

import (
	"testing"
)

const input = "input.txt"
const sample = "sample.txt"

func TestPartOne(t *testing.T) {
	test_cases := []struct {
		input    string
		expected int
	}{
		{sample, 288},
		{input, 0},
	}

	for _, test_case := range test_cases {
		actual := PartOne(test_case.input)

		if actual != test_case.expected {
			t.Errorf("Test failed, expected: '%d', got:  '%d'", test_case.expected, actual)
		}
	}
}

func TestParseRaces(t *testing.T) {
	expected := []Race{
		{duration_milliseconds: 7, distance_millimeters: 9},
		{duration_milliseconds: 15, distance_millimeters: 40},
		{duration_milliseconds: 30, distance_millimeters: 200},
		{duration_milliseconds: 71530, distance_millimeters: 940200},
	}

	actual := parseRaces(getLines(sample))

	if len(actual) != len(expected) {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", len(expected), len(actual))
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Test failed, expected: '%v', got:  '%v'", expected[i], actual[i])
		}
	}

}

func TestRace(t *testing.T) {
	test_cases := []struct {
		input    Race
		expected int
	}{
		{Race{duration_milliseconds: 7, distance_millimeters: 9}, 4},
		{Race{duration_milliseconds: 15, distance_millimeters: 40}, 8},
		{Race{duration_milliseconds: 30, distance_millimeters: 200}, 9},
		{Race{duration_milliseconds: 30, distance_millimeters: 200}, 9},
		{Race{duration_milliseconds: 71530, distance_millimeters: 940200}, 71503},
		{Race{duration_milliseconds: 44899691, distance_millimeters: 277113618901768}, 71503},
	}

	for _, test_case := range test_cases {
		actual := waysRaceCouldBeWon(test_case.input)
		if actual != test_case.expected {
			t.Errorf("Test failed %v, expected: '%d', got:  '%d'", test_case.input, test_case.expected, actual)
		}
	}
}
