package day8

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func allZeros(history []int) bool {
	for _, v := range history {
		if v != 0 {
			return false
		}
	}

	return true
}

func nextValueInHistory(history []int) int {
	if allZeros(history) {
		return 0
	}

	last_value := history[len(history)-1]
	new_history := make([]int, len(history)-1)

	for i := 0; i < len(history)-1; i++ {
		current := history[i]
		next := history[i+1]

		new_history[i] = next - current
	}

	return nextValueInHistory(new_history) + last_value
}

func parseLine(line string) []int {

	fields := strings.Fields(line)
	output := make([]int, len(fields))

	for i := range fields {
		number, err := strconv.Atoi(fields[i])
		if err != nil {
			panic(err)
		}
		output[i] = number
	}

	return output
}

func PartOne(input_path string) int {
	lines, _ := getLines(input_path)

	acc := 0

	for _, line := range lines {
		acc = acc + nextValueInHistory(parseLine(line))
	}

	return acc
}
