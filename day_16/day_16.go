package day_16

import (
	"bufio"
	"math"
	"os"
)

type Point struct {
	i int
	j int
}

func PartOne(input_file string) int {
	lines, err := getLines(input_file)
	if err != nil {
		panic(err)
	}

	return energized(lines)
}

func energized(lines []string) int {
	grid := map[Point]rune{}
	unvisited := map[Point]bool{}
	distances := map[Point]int{}

	rows := len(lines)
	cols := len(lines[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			point := Point{i: i, j: j}
			unvisited[point] = false
			distances[point] = math.MaxInt
			grid[point] = rune(lines[i][j])
		}
	}

	start := Point{i: 0, j: 0}
	distances[start] = 0

	current_point := start
	for {
		point := Point{i: current_point.i, j: current_point.j + 1}
		if !IsPointInBoard(point, rows, cols) {
			break
		}
		_, present := unvisited[point]

		if present {
			current_distance := distances[current_point]
			new_distance := current_distance + 1
			point_distance := distances[point]

			if new_distance < point_distance {
				distances[point] = new_distance
			}
		}

		delete(unvisited, current_point)
		current_point = next_point_unvisited(unvisited, distances)
		if !IsPointInBoard(current_point, rows, cols) {
			break
		}
	}

	return rows*cols - len(unvisited) + 1
}

func PartTwo(input_file string) int {
	return 0
}

func IsPointInBoard(p Point, rows int, cols int) bool {
	return p.i >= 0 && p.j >= 0 && p.i < rows && p.j < cols
}

func next_point_unvisited(unvisited map[Point]bool, distances map[Point]int) Point {
	min_distance := math.MaxInt
	min_i := -1
	min_j := -1

	for point := range unvisited {
		distance := distances[point]
		if distance < min_distance {
			min_distance = distance
			min_i = point.i
			min_j = point.j
		}
	}

	return Point{min_i, min_j}
}

func getLines(input_file string) ([]string, error) {
	file, err := os.Open(input_file)
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
