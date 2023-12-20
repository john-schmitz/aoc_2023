package day_19

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		expected int
	}{
		{
			desc:     "sample.txt",
			expected: 19114,
		},
		{
			desc:     "input.txt",
			expected: 432788,
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

func TestPartSum(t *testing.T) {
	testCases := []struct {
		input    Part
		expected int
	}{
		{
			input:    Part{787, 2655, 1222, 2876},
			expected: 7540,
		},
		{
			input:    Part{2036, 264, 79, 2244},
			expected: 4623,
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v", tC.input), func(t *testing.T) {

			if tC.input.Sum() != tC.expected {
				t.Errorf("Expected Sum of %v = %d Got %d", tC.input, tC.expected, tC.input.Sum())
			}
		})
	}
}

func TestEvaluateExpression(t *testing.T) {
	testCases := []struct {
		expression string
		part       Part
		expected   bool
	}{
		{
			expression: "a<2006",
			part:       Part{x: 0, m: 0, a: 100, s: 0},
			expected:   true,
		},
		{
			expression: "a>2006",
			part:       Part{x: 0, m: 0, a: 100, s: 0},
			expected:   false,
		},
		{
			expression: "m<20",
			part:       Part{x: 0, m: 200, a: 0, s: 0},
			expected:   false,
		},
		{
			expression: "m>10",
			part:       Part{x: 0, m: 11, a: 0, s: 0},
			expected:   true,
		},
		{
			expression: "s<20",
			part:       Part{s: 200},
			expected:   false,
		},
		{
			expression: "s>10",
			part:       Part{s: 11},
			expected:   true,
		},
		{
			expression: "x<20",
			part:       Part{x: 200},
			expected:   false,
		},
		{
			expression: "x>10",
			part:       Part{x: 11},
			expected:   true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.expression, func(t *testing.T) {
			actual := evaluateExpression(tC.part, tC.expression)

			if actual != tC.expected {
				t.Errorf("Expected evaluateExpression(%v, %s) = %t. Got %t", tC.part, tC.expression, tC.expected, actual)
			}
		})
	}
}

func TestWorkflow(t *testing.T) {
	testCases := []struct {
		expression string
		part       Part
		expected   string
	}{
		{
			expression: "a<2006:qkq,m>2090:A,rfg",
			part:       Part{x: 1679, m: 44, a: 2067, s: 496},
			expected:   "rfg",
		},
		{
			expression: "s<1351:px,qqz",
			part:       Part{x: 787, m: 2655, a: 1222, s: 2876},
			expected:   "qqz",
		},
		{
			expression: "s<1351:px,qqz",
			part:       Part{x: 2461, m: 1339, a: 466, s: 291},
			expected:   "px",
		},
		{
			expression: "m>1548:A,A",
			part:       Part{x: 2461, m: 1339, a: 466, s: 291},
			expected:   "A",
		},
		{
			expression: "m>1548:R,R",
			part:       Part{x: 2461, m: 1339, a: 466, s: 291},
			expected:   "R",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.expression, func(t *testing.T) {
			actual := evaluateWorkflow(tC.part, tC.expression)

			if actual != tC.expected {
				t.Errorf("Expected evaluateWorkflow(%v, %s) = %s. Got %s", tC.part, tC.expression, tC.expected, actual)
			}
		})
	}
}
