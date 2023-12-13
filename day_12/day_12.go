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
	if expression == "" {
		if len(groups) > 0 {
			return 0
		}
		return 1
	}

	if len(groups) == 0 {
		if strings.ContainsRune(expression, '#') {
			return 0
		}
		return 1
	}

	result := 0
	current_group := groups[0]
	switch expression[0] {
	case '.':
		result += solve(expression[1:], groups)
	case '?':
		result += solve(expression[1:], groups) + solve(appendAtStart(expression, "#"), groups)
	case '#':
		if current_group > len(expression) {
			break
		}

		if strings.ContainsRune(expression[:current_group], '.') {
			break
		}

		if current_group == len(expression) || expression[current_group] != '#' {
			x := ""
			if len(expression) >= current_group+1 {
				x = expression[current_group+1:]
			}
			result += solve(x, groups[1:])
		}
	}

	return result
}

func PossibleArrangements(line string) int {
	slitted := strings.Split(line, " ")
	numbers := toIntSlice(slitted[1])

	return solve(slitted[0], numbers)
}
