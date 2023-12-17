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

	start := Point{i: 0, j: -1}
	delete(unvisited, start)

	vectors := []struct {
		direction     Point
		current_point Point
	}{
		{
			direction:     Point{i: 0, j: 1},
			current_point: start,
		},
	}

	for len(vectors) > 0 {
		for index := 0; index < len(vectors); index++ {
			vector := vectors[index]
			current_point := Point{i: vector.current_point.i + vector.direction.i, j: vector.current_point.j + vector.direction.j}
			if !IsPointInBoard(current_point, rows, cols) {
				vectors = append(vectors[:index], vectors[index+1:]...)
				index--
				continue
			}

			element := grid[current_point]
			previous_point := Point{i: current_point.i + (-1 * vector.direction.i), j: current_point.j + (-1 * vector.direction.j)}

			if element == '.' {
				delete(unvisited, current_point)
				vectors[index].current_point = current_point
				continue
			}

			if element == '|' {
				if previous_point.j != current_point.j {
					vectors[index].direction = Point{1, 0}
					vectors = append(vectors, struct {
						direction     Point
						current_point Point
					}{
						direction:     Point{-1, 0},
						current_point: current_point,
					})
				}

				delete(unvisited, current_point)
				vectors[index].current_point = current_point
				continue
			}

			if element == '-' {
				if previous_point.i != current_point.i {
					vectors[index].direction = Point{0, -1}
					vectors = append(vectors, struct {
						direction     Point
						current_point Point
					}{
						direction:     Point{0, 1},
						current_point: current_point,
					})
				}

				delete(unvisited, current_point)
				vectors[index].current_point = current_point
				continue
			}
		}
	}

	return rows*cols - len(unvisited)
}

func PartTwo(input_file string) int {
	return 0
}

func IsPointInBoard(p Point, rows int, cols int) bool {
	return p.i >= 0 && p.j >= 0 && p.i < rows && p.j < cols
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
