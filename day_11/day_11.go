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

func (point Point) ManhattanDistance(p2 Point) int {
	manhattan_i := p2.i - point.i
	manhattan_j := p2.j - point.j

	if manhattan_i < 0 {
		manhattan_i = -manhattan_i
	}

	if manhattan_j < 0 {
		manhattan_j = -manhattan_j
	}

	return manhattan_i + manhattan_j
}

func solve(file_path string, increments int) int {
	totalDistance := 0
	lines, _ := getLines(file_path)
	galaxies := transposeBoard(lines, increments)

	distances_map := map[string]int{}

	for _, from := range galaxies {
		for _, to := range galaxies {
			if to == from {
				continue
			}
			key := fmt.Sprintf("%d:%d|%d:%d", from.i, from.j, to.i, to.j)

			_, present := distances_map[key]
			if present {
				continue
			}
			key = fmt.Sprintf("%d:%d|%d:%d", to.i, to.j, from.i, from.j)
			_, present = distances_map[key]
			if present {
				continue
			}
			totalDistance += from.ManhattanDistance(to)
		}
	}

	return totalDistance / 2
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

	return galaxies
}
