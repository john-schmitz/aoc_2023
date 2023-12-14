package day_13

import (
	"bufio"
	"os"
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

func PartOne(input_file string) int {
	lines, err := getLines(input_file)
	if err != nil {
		panic(err)
	}

	acc := 0
	current_slice := make([]string, 0)

	for index, line := range lines {
		if line != "" {
			current_slice = append(current_slice, line)
		}

		if line == "" || index == len(lines)-1 {
			acc = acc + FindReflection(current_slice)
			current_slice = make([]string, 0)
		}
	}

	return acc
}

func PartTwo(input_file string) int {
	return 0
}

func getColumn(grid []string, index int) string {
	column_size := len(grid[0])
	row_count := len(grid)

	if index > column_size-1 {
		panic("Index out of bounds")
	}

	column := ""

	for i := 0; i < row_count; i++ {
		column = column + string(grid[i][index])
	}

	return column
}

func FindReflection(grid []string) int {
	output := 0

	column_size := len(grid[0])
	row_count := len(grid)

vertical:
	for i := 0; i < column_size-1; i++ {
		currentRow := getColumn(grid, i)
		nextRow := getColumn(grid, i+1)

		if nextRow == currentRow {
			current_index := i
			next_index := i + 1
			for j := 1; j <= current_index; j++ {
				if (next_index+j) >= column_size || (current_index-j) < 0 {
					continue
				}

				if getColumn(grid, current_index-j) != getColumn(grid, next_index+j) {
					continue vertical
				}
			}
			output = i + 1
		}
	}

	if output > 0 {
		return output
	}

horizontal:
	for i := 0; i < row_count-1; i++ {
		currentRow := grid[i]
		nextRow := grid[i+1]

		if nextRow == currentRow {
			current_index := i
			next_index := i + 1
			for j := 1; j <= current_index; j++ {
				if (next_index+j) >= row_count || (current_index-j) < 0 {
					continue
				}

				if grid[current_index-j] != grid[next_index+j] {
					continue horizontal
				}
			}
			output = i + 1
		}
	}

	return output * 100
}
