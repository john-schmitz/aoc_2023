package main

import "testing"

func TestPossibleArrangements(t *testing.T) {
	testCases := []struct {
		line     string
		expected int
	}{
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
