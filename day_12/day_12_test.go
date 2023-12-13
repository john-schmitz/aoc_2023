package main

import "testing"

func TestDebug(t *testing.T) {
	input := "??..## 1,2"
	expected := 2
	actual := PossibleArrangements(input)

	if actual != expected {
		t.Errorf("PossibleArrangements(%s) = %d. Got %d", input, expected, actual)
	}
}

func TestPossibleArrangements(t *testing.T) {
	testCases := []struct {
		line     string
		expected int
	}{
		{
			line:     "?...## 1,2",
			expected: 1,
		},
		{
			line:     "??..## 1,2",
			expected: 2,
		},
		{
			line:     "???..## 1,2",
			expected: 3,
		},
		{
			line:     "???..## 1,2",
			expected: 3,
		},
		{
			line:     "????.## 1,2",
			expected: 4,
		},
		{
			line:     "????.#...#... 4,1,1",
			expected: 1,
		},
		{
			line:     "?.?.### 1,1,3",
			expected: 1,
		},
		{
			line:     "???.### 1,1,3",
			expected: 1,
		},
		{
			line:     "????.######..#####. 1,6,5",
			expected: 4,
		},
		{
			line:     ".??..??...?##. 1,1,3",
			expected: 4,
		},
		{
			line:     "?#?#?#?#?#?#?#? 1,3,1,6",
			expected: 1,
		},
		{
			line:     "?###???????? 3,2,1",
			expected: 10,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.line, func(t *testing.T) {
			actual := PossibleArrangements(tC.line)

			if actual != tC.expected {
				t.Errorf("PossibleArrangements(%s) = %d. Got %d", tC.line, tC.expected, actual)
			}
		})
	}
}
