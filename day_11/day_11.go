package day11

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	i int
	j int
}

func PartOne(file_path string) int {
	lines, _ := getLines(file_path)
	lines = transposeBoard(lines, 2)
	galaxies := make([]Point, 0)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			element := string(lines[i][j])
			if element == "#" {
				galaxies = append(galaxies, Point{i, j})
			}
		}
	}
	fmt.Println("Found galaxies: ", len(galaxies))

	distances_map := map[string]int{}
	expected_pairs := len(galaxies) * (len(galaxies) - 1) / 2

	for _, galaxy := range galaxies {
		for _, galaxy_ := range galaxies {
			if galaxy_ == galaxy {
				continue
			}
			key := fmt.Sprintf("%d:%d|%d:%d", galaxy.i, galaxy.j, galaxy_.i, galaxy_.j)

			_, present := distances_map[key]
			if present {
				continue
			}
			key = fmt.Sprintf("%d:%d|%d:%d", galaxy_.i, galaxy_.j, galaxy.i, galaxy.j)
			_, present = distances_map[key]
			if present {
				continue
			}

			manhattan_i := galaxy.i - galaxy_.i
			manhattan_j := galaxy.j - galaxy_.j

			if manhattan_i < 0 {
				manhattan_i = -manhattan_i
			}

			if manhattan_j < 0 {
				manhattan_j = -manhattan_j
			}

			distances_map[key] = manhattan_i + manhattan_j
		}
	}

	if len(distances_map) != expected_pairs {
		panic(fmt.Sprintf("Expected %d pairs, got %d", expected_pairs, len(distances_map)))
	}
	acc := 0
	for _, element := range distances_map {
		acc += element
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

func IsPointInBoard(p Point, board []string) bool {
	return p.i >= 0 && p.j >= 0 && p.i < len(board) && p.j < len(board[p.i])
}

func transposeBoard(lines []string, increment int) []string {
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

	empty_row := strings.Repeat(".", len(lines[0]))
	rows_to_be_added := len(rows_with_no_galaxies)

	for len(rows_with_no_galaxies) > 0 {
		row_index := rows_with_no_galaxies[0] + rows_to_be_added - len(rows_with_no_galaxies)
		rows_with_no_galaxies = rows_with_no_galaxies[1:]
		for i := 0; i < increment-1; i++ {
			lines = slices.Insert(lines, row_index, empty_row)
		}
	}

	for index, row := range lines {
		new_row := row

		for i := 0; i < len(columns_with_no_galaxies); i++ {
			column_index := i + columns_with_no_galaxies[i]
			new_value := strings.Repeat(".", increment-1)
			new_row = new_row[:column_index] + new_value + new_row[column_index:]
		}

		lines[index] = new_row
	}

	return lines
}
