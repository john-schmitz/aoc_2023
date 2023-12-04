package day4

import "testing"

func TestParseGame(t *testing.T) {
	test_cases := []struct {
		input string
		want  int
	}{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
		{"Card 177: 78 14 51 53 26 44 36  2 33 23 | 53 26 48 24  5 88 33 90 36 22 29 89 51 45 70 44 23  2 73 91 57 92 78 47 14", 512},
	}

	for _, tc := range test_cases {
		got := parseGame(tc.input)
		if got != tc.want {
			t.Errorf("parseGame(%q) = %d; want %d", tc.input, got, tc.want)
		}
	}
}

func TestPartOneSample(t *testing.T) {
	got := PartOne("sample.txt")
	want := 13
	if got != want {
		t.Errorf("PartOne(sample.txt) = %d; want %d", got, want)
	}
}

func TestPartOneInput(t *testing.T) {
	got := PartOne("input.txt")
	want := 19135
	if got != want {
		t.Errorf("PartOne(input.txt) = %d; want %d", got, want)
	}
}

func TestPartTwoSample(t *testing.T) {
	got := PartTwo("sample.txt")
	want := 30
	if got != want {
		t.Errorf("PartTwo(sample.txt) = %d; want %d", got, want)
	}
}

func TestPartTwoInput(t *testing.T) {
	got := PartTwo("input.txt")
	want := 0
	if got != want {
		t.Errorf("PartTwo(input.txt) = %d; want %d", got, want)
	}
}
