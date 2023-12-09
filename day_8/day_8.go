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
	for v := range history {
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

	new_slice := make([]int, len(history)-1)

	for i := 0; i < len(history)-1; i++ {
		current := history[i]
		next := history[i+1]

		new_slice[i] = next - current
	}

	

	return 0
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

	return len(lines)
}
