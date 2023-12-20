package day_19

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	x, m, a, s int
}

func PartOne(input_file string) int {
	lines, err := getLines(input_file)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return 0
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
