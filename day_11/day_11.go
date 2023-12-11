package day11

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func transposeBoard(lines []string) []string {

	rows_with_no_galaxies := make([]int, 0)
	columns_with_no_galaxies := make([]int, 0)
	for row_index, row := range lines {
		if !strings.Contains(row, "#") {
			rows_with_no_galaxies = append(rows_with_no_galaxies, row_index)

		}
	}

	for j := 0; j < len(lines[0]); j++ {
		cur_column := ""
		for i := 0; i < len(lines); i++ {
			cur_column += string(lines[i][j])
		}
		if !strings.Contains(cur_column, "#") {
			columns_with_no_galaxies = append(columns_with_no_galaxies, j)
		}
	}

	fmt.Println("Rows with no galaxies: ", rows_with_no_galaxies)
	fmt.Println("Columns with no galaxies: ", columns_with_no_galaxies)

	empty_row := strings.Repeat(".", len(lines[0]))
	rows_to_be_added := len(rows_with_no_galaxies)

	for len(rows_with_no_galaxies) > 0 {
		row_index := rows_with_no_galaxies[0] + rows_to_be_added - len(rows_with_no_galaxies)
		rows_with_no_galaxies = rows_with_no_galaxies[1:]
		fmt.Println(row_index, rows_with_no_galaxies)
		lines = slices.Insert(lines, row_index, empty_row)
	}

	return lines
}
