package day_19

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		expected int
	}{
		{
			desc:     "sample.txt",
			expected: 200,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := PartOne(tC.desc)
			if actual != tC.expected {
				t.Errorf("Expected PartOne(%s) = %d. Got %d", tC.desc, tC.expected, actual)
			}
		})
	}
}

func TestParsePart(t *testing.T) {
	testCases := []struct {
		input    string
		expected Part
	}{
		{
			input:    "{x=787,m=2655,a=1222,s=2876}",
			expected: Part{787, 2655, 1222, 2876},
		},
		{
			input:    "{x=1679,m=44,a=2067,s=496}",
			expected: Part{1679, 44, 2067, 496},
		},
		{
			input:    "{x=2036,m=264,a=79,s=2244}",
			expected: Part{2036, 264, 79, 2244},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			actual, _ := parsePart(tC.input)
			if actual != tC.expected {
				t.Errorf("Expected ParsePart(%s) = %v. Got %v", tC.input, tC.expected, actual)
			}
		})
	}
}
