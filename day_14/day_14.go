package day14

import (
	"bufio"
	"os"
)

func PartOne(input_file string) int {
	lines, err := getLines(input_file)
	if err != nil {
		panic(err)
	}

	grid := make([][]rune, 0)

	for _, row := range lines {
		grid = append(grid, []rune(row))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			elm := grid[i][j]
			if elm == rune('#') || elm == rune('.') {
				continue
			}

			if i-1 < 0 {
				continue
			}

			above_count := 0
			for z := 1; z <= i; z++ {
				x := grid[i-z][j]
				if x == rune('.') {
					above_count++
				} else {
					break
				}
			}

			if above_count > 0 {
				grid[i][j], grid[i-above_count][j] = grid[i-above_count][j], grid[i][j]
			}
		}
	}

	acc := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			elm := grid[i][j]
			if elm == rune('O') {
				score := len(grid) - i
				acc = acc + score
			}
		}
	}

	return acc
}

func PartTwo(input_file string) int {
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
