package day11

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	i int
	j int
}

func PartOne(file_path string) int {
	lines, _ := getLines(file_path)
	lines = transposeBoard(lines)
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
		distances := traverse(galaxy, lines)
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

			distances_map[key] = distances[galaxy_.i][galaxy_.j]
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

func traverse(start Point, board []string) [][]int {
	unvisited := map[string]bool{}

	distances := make([][]int, len(board))
	for i := range distances {
		distances[i] = make([]int, len(board[0]))
		for j := range distances[i] {
			distances[i][j] = math.MaxInt
			key := getKey(i, j)
			unvisited[key] = false
		}
	}
	distances[start.i][start.j] = 0
	current_point := start
	directions := []Point{
		{1, 0},  // UP
		{0, 1},  // RIGHT
		{-1, 0}, // DOWN
		{0, -1}, // LEFT
	}

	for {
		for _, direction := range directions {
			point := Point{current_point.i + direction.i, current_point.j + direction.j}
			if !IsPointInBoard(point, board) {
				continue
			}

			_, present := unvisited[getKey(point.i, point.j)]
			if present {
				current_distance := distances[current_point.i][current_point.j]
				new_distance := current_distance + 1
				point_distance := distances[point.i][point.j]

				if new_distance < point_distance {
					distances[point.i][point.j] = new_distance
				}
			}
		}
		delete(unvisited, getKey(current_point.i, current_point.j))
		current_point = next_point_unvisited(unvisited, distances)
		if !IsPointInBoard(current_point, board) {
			break
		}
	}

	return distances
}

func next_point_unvisited(unvisited map[string]bool, distances [][]int) Point {
	min_distance := math.MaxInt
	min_i := -1
	min_j := -1

	for k := range unvisited {
		i, j := getPoints(k)
		distance := distances[i][j]
		if distance < min_distance {
			min_distance = distance
			min_i = i
			min_j = j
		}

	}

	return Point{min_i, min_j}
}

func getKey(i, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
}

func getPoints(key string) (i, j int) {
	result := strings.Split(key, ":")

	i, err := strconv.Atoi(result[0])
	if err != nil {
		panic(err)
	}

	j, err = strconv.Atoi(result[1])
	if err != nil {
		panic(err)
	}

	return i, j
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

	empty_row := strings.Repeat(".", len(lines[0]))
	rows_to_be_added := len(rows_with_no_galaxies)

	for len(rows_with_no_galaxies) > 0 {
		row_index := rows_with_no_galaxies[0] + rows_to_be_added - len(rows_with_no_galaxies)
		rows_with_no_galaxies = rows_with_no_galaxies[1:]
		lines = slices.Insert(lines, row_index, empty_row)
	}

	for index, row := range lines {
		new_row := row

		for i := 0; i < len(columns_with_no_galaxies); i++ {
			column_index := i + columns_with_no_galaxies[i]
			new_row = new_row[:column_index] + "." + new_row[column_index:]
		}

		lines[index] = new_row
	}

	return lines
}
