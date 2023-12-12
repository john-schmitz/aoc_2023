package main

import (
	"strconv"
	"strings"
)

func toIntSlice(input string) []int {
	output := make([]int, 0)

	for _, v := range strings.Split(input, ",") {
		number, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		output = append(output, number)
	}

	return output
}

func appendAtStart(s string, value string) string {
	if len(s) < 1 {
		panic("String has no length")
	}

	if len(value) != 1 {
		panic("Invalid length")
	}

	return s[:0] + value + s[0+1:]
}

func solve(expression string, groups []int) int {
	if expression == "" && len(groups) > 0 {
		return 0
	}

	if expression == "" && len(groups) == 0 {
		return 1
	}

	if len(groups) == 0 {
		return 1
	}

	first_char := string(expression[0])
	current_group := groups[0]

	if first_char == "?" {
		return solve(appendAtStart(expression, "#"), groups) + solve(appendAtStart(expression, "."), groups)
	}

	if first_char == "." {
		return solve(expression[1:], groups)
	}

	if first_char == "#" {
		if current_group > len(expression) {
			return 0
		}
		for i := 0; i < current_group; i++ {
			if string(expression[i]) == "." {
				return solve(expression[i:], groups)
			}
		}

		x := expression[current_group:]

		if x == "" {
			return solve(x, groups[1:])
		}

		return solve(x[1:], groups[1:])
	}

	return 0
}

func PossibleArrangements(line string) int {
	slitted := strings.Split(line, " ")
	numbers := toIntSlice(slitted[1])

	return solve(slitted[0], numbers)
}
