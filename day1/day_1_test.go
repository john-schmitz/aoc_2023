package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

const INPUT_FILE = "input.txt"

type test_case struct {
	input  string
	output int
}

func TestLineParse(t *testing.T) {
	test_cases := []test_case{
		{"one", 11},
		{"two", 22},
		{"three", 33},
		{"four", 44},
		{"five", 55},
		{"six", 66},
		{"seven", 77},
		{"eight", 88},
		{"nine", 99},
		{"two1nine", 29},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"sixgf2sssrzvqlsm9one26twonedzq", 61},
	}

	for _, test_case := range test_cases {

		parsed, err := lineParse(test_case.input)
		if err != nil {
			t.Errorf("Error parsing line: %s", err)
		}

		if parsed != test_case.output {
			t.Errorf("Expected %d, got %d", test_case.output, parsed)
		}
	}
}

func TestPartTwo(t *testing.T) {
	total_sum := 0
	readFile, err := os.Open(INPUT_FILE)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		result_number, err := lineParse(line)
		if err != nil {
			panic(err)
		}

		total_sum += result_number
	}

	fmt.Println("Part two solution: ", total_sum)
	readFile.Close()
}

func TestPartOne(t *testing.T) {
	total_sum := 0
	readFile, err := os.Open(INPUT_FILE)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		first_num := -1
		second_num := -1

		for _, char := range line {

			i, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			if first_num == -1 {
				first_num = i
			}

			second_num = i
		}

		result_number, err := strconv.Atoi(strings.Join([]string{fmt.Sprint(first_num), fmt.Sprint(second_num)}, ""))
		if err != nil {
			panic(err)
		}

		total_sum += result_number
	}

	fmt.Println("Part one solution: ", total_sum)
	readFile.Close()
}
