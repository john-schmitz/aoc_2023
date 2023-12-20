package day_19

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	x, m, a, s int
}

func (part Part) Sum() int {
	return part.x + part.a + part.m + part.s
}

func PartOne(input_file string) int {
	lines, err := getLines(input_file)
	if err != nil {
		panic(err)
	}

	parts := []Part{}
	expressions := map[string]string{}
	parsing_expressions := true

	for _, line := range lines {
		if line == "" {
			parsing_expressions = false
			continue
		}
		if parsing_expressions {
			index := strings.IndexAny(line, "{")
			expression := line[index:]
			expression = expression[1:][:len(expression)-2]
			token := line[:index]
			expressions[token] = expression
		} else {
			part, err := parsePart(line)
			if err != nil {
				panic(err)
			}

			parts = append(parts, part)
		}
	}
	acc := 0
	for i := 0; i < len(parts); i++ {
		start := "in"
		part := parts[i]
		result := evaluateWorkflow(part, expressions[start])

		for result != "A" && result != "R" {
			result = evaluateWorkflow(part, expressions[result])
		}

		if result == "A" {
			acc = acc + part.Sum()
		}
	}

	return acc
}

func getLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parsePart(line string) (Part, error) {
	without := strings.Split(line[1:][:len(line)-2], ",")
	x, err := strconv.Atoi(strings.Split(without[0], "=")[1])
	if err != nil {
		return Part{}, err
	}

	m, err := strconv.Atoi(strings.Split(without[1], "=")[1])
	if err != nil {
		return Part{}, err
	}

	a, err := strconv.Atoi(strings.Split(without[2], "=")[1])
	if err != nil {
		return Part{}, err
	}

	s, err := strconv.Atoi(strings.Split(without[3], "=")[1])
	if err != nil {
		return Part{}, err
	}

	return Part{x, m, a, s}, nil
}

func getOperator(expression string) string {
	if strings.Contains(expression, "<") {
		return "<"
	}

	return ">"
}

func getComparator(part Part, expression string) int {
	if strings.Contains(expression, "a") {
		return part.a
	}

	if strings.Contains(expression, "s") {
		return part.s
	}

	if strings.Contains(expression, "m") {
		return part.m
	}

	return part.x
}

func evaluateExpression(part Part, expression string) bool {
	operator := getOperator(expression)

	result := strings.Split(expression, operator)
	parsedInt, err := strconv.Atoi(result[1])
	if err != nil {
		panic(err)
	}

	comparator := getComparator(part, expression)

	if operator == ">" {
		return comparator > parsedInt
	}

	return comparator < parsedInt
}

func evaluateWorkflow(part Part, line string) string {
	expressions := strings.Split(line, ",")
	for _, expression := range expressions {
		if !strings.Contains(expression, ":") {
			return expression
		}

		splited := strings.Split(expression, ":")

		next_token := splited[1]

		if evaluateExpression(part, splited[0]) {
			return next_token
		}
	}

	panic("end of workflow")
}
