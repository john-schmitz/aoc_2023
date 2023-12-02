package day2

import (
	"testing"
)

const SAMPLE_PATH = "sample.txt"
const INPUT_PATH = "input.txt"

func TestPossibleGames(t *testing.T) {
	test_cases := []struct {
		input  string
		output bool
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}

	for _, pair := range test_cases {
		expected := pair.output
		actual, err := is_possible_game(pair.input)

		if err != nil {
			t.Error(
				"For", pair.input,
				"expected", expected,
				"got", err,
			)
		}

		if actual != expected {
			t.Error(
				"For", pair.input,
				"expected", expected,
				"got", actual,
			)
		}
	}

}

func TestGameId(t *testing.T) {
	test_cases := []struct {
		input  string
		output int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 2},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 3},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 4},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 5},
	}

	for _, pair := range test_cases {
		expected := pair.output
		actual, err := get_game_id(pair.input)

		if err != nil {
			t.Error(
				"For", pair.input,
				"expected", expected,
				"got", err,
			)
		}

		if actual != expected {
			t.Error(
				"For", pair.input,
				"expected", expected,
				"got", actual,
			)
		}
	}
}

func TestSamplePartOne(t *testing.T) {
	expected := 8
	actual, err := part_one(SAMPLE_PATH)

	if err != nil {
		t.Error(
			"For", SAMPLE_PATH,
			"expected", expected,
			"got", err,
		)
	}

	if actual != expected {
		t.Error(
			"For", SAMPLE_PATH,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestPartOne(t *testing.T) {
	expected := 2176
	actual, err := part_one(INPUT_PATH)

	if err != nil {
		t.Error(
			"For", INPUT_PATH,
			"expected", expected,
			"got", err,
		)
	}

	if actual != expected {
		t.Error(
			"For", INPUT_PATH,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestFewestCubes(t *testing.T) {
	test_cases := []struct {
		input  string
		output int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
	}

	for _, pair := range test_cases {
		expected := pair.output
		actual, err := fewests_cubes_needed(pair.input)

		if err != nil {
			t.Error(
				"For", pair.input,
				"expected", expected,
				"got", err,
			)
		}

		if actual != expected {
			t.Error(
				"For", pair.input,
				"expected", expected,
				"got", actual,
			)
		}
	}
}

func TestPartTwoSample(t *testing.T) {
	expected := 2286
	actual, err := part_two(SAMPLE_PATH)

	if err != nil {
		t.Error(
			"For", SAMPLE_PATH,
			"expected", expected,
			"got", err,
		)
	}

	if actual != expected {
		t.Error(
			"For", SAMPLE_PATH,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 63700
	actual, err := part_two(INPUT_PATH)

	if err != nil {
		t.Error(
			"For", INPUT_PATH,
			"expected", expected,
			"got", err,
		)
	}

	if actual != expected {
		t.Error(
			"For", INPUT_PATH,
			"expected", expected,
			"got", actual,
		)
	}
}
