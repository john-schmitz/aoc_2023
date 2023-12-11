package day11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	i int
	j int
}

func solve(file_path string, increments int) int {
	lines, _ := getLines(file_path)
	galaxies := transposeBoard(lines, increments)

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

func PartOne(file_path string) int {
	return solve(file_path, 2)
}

func PartTwo(file_path string) int {
	return solve(file_path, 1000000)
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

func transposeBoard(lines []string, increment int) []Point {
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

	galaxies := make([]Point, 0)
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			element := string(lines[i][j])
			if element == "#" {
				new_i := i
				new_j := j
				for _, v := range rows_with_no_galaxies {
					if v < i {
						new_i = new_i + increment - 1
					}
				}

				for _, v := range columns_with_no_galaxies {
					if v < j {
						new_j = new_j + increment - 1
					}
				}
				galaxies = append(galaxies, Point{new_i, new_j})
			}
		}
	}
	fmt.Println("Found galaxies: ", len(galaxies))

	return galaxies
}
