package day_16

import (
	"bufio"
	"fmt"
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
	start := Point{i: 0, j: -1}
	direction := Point{i: 0, j: 1}
	return energized(lines, start, direction)
}

var DOWN Point = Point{i: 1, j: 0}
var UP Point = Point{i: -1, j: 0}
var LEFT Point = Point{i: 0, j: -1}
var RIGHT Point = Point{i: 0, j: 1}

func energized(lines []string, start Point, direction Point) int {
	grid := map[Point]rune{}
	unvisited := map[Point]bool{}
	distances := map[Point]int{}

	rows := len(lines)
	cols := len(lines[0])
	visited_with_diretion := map[Point]map[Point]bool{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			point := Point{i: i, j: j}
			visited_with_diretion[point] = map[Point]bool{}
			visited_with_diretion[point][UP] = false
			visited_with_diretion[point][DOWN] = false
			visited_with_diretion[point][RIGHT] = false
			visited_with_diretion[point][LEFT] = false
			unvisited[point] = false
			distances[point] = math.MaxInt
			grid[point] = rune(lines[i][j])
		}
	}

	vectors := []struct {
		direction     Point
		current_point Point
	}{
		{
			direction:     direction,
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
				not_visited := !visited_with_diretion[current_point][vector.direction]
				if !not_visited {
					vectors = append(vectors[:index], vectors[index+1:]...)
					index--
					continue
				}

				if previous_point.j != current_point.j && not_visited {
					vectors[index].direction = UP
					vectors = append(vectors, struct {
						direction     Point
						current_point Point
					}{
						direction:     DOWN,
						current_point: current_point,
					})
				}

				visited_with_diretion[current_point][vector.direction] = true
				delete(unvisited, current_point)
				vectors[index].current_point = current_point
				continue
			}

			if element == '-' {
				not_visited := !visited_with_diretion[current_point][vector.direction]
				if !not_visited {
					vectors = append(vectors[:index], vectors[index+1:]...)
					index--
					continue
				}

				if previous_point.i != current_point.i && not_visited {
					vectors[index].direction = LEFT
					vectors = append(vectors, struct {
						direction     Point
						current_point Point
					}{
						direction:     RIGHT,
						current_point: current_point,
					})
				}

				visited_with_diretion[current_point][vector.direction] = true
				delete(unvisited, current_point)
				vectors[index].current_point = current_point
				continue
			}

			if element == '\\' {
				if vector.direction == RIGHT {
					vectors[index].direction = DOWN
				} else if vector.direction == UP {
					vectors[index].direction = LEFT
				} else if vector.direction == DOWN {
					vectors[index].direction = RIGHT
				} else if vector.direction == LEFT {
					vectors[index].direction = UP
				} else {
					panic("INVALID DIRECTION")
				}

				delete(unvisited, current_point)
				vectors[index].current_point = current_point
				continue
			}

			if element == '/' {
				if vector.direction == RIGHT {
					vectors[index].direction = UP
				} else if vector.direction == UP {
					vectors[index].direction = RIGHT
				} else if vector.direction == DOWN {
					vectors[index].direction = LEFT
				} else if vector.direction == LEFT {
					vectors[index].direction = DOWN
				} else {
					panic("INVALID DIRECTION")
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
	lines, err := getLines(input_file)
	if err != nil {
		panic(err)
	}

	result := math.MinInt

	rows := len(lines)
	cols := len(lines[0])

	for i := 0; i < rows; i++ {
		start := Point{i: -1, j: i}
		if IsPointInBoard(start, rows, cols) {
			panic(fmt.Sprintf("Starting position cannot be in grid. %v for DOWN", start))
		}
		result = max(result, energized(lines, start, DOWN))

		start = Point{i: rows, j: i}
		if IsPointInBoard(start, rows, cols) {
			panic(fmt.Sprintf("Starting position cannot be in grid. %v for UP", start))
		}
		result = max(result, energized(lines, start, UP))
	}

	for i := 0; i < cols; i++ {
		start := Point{i: i, j: cols}
		if IsPointInBoard(start, rows, cols) {
			panic(fmt.Sprintf("Starting position cannot be in grid. %v for LEFT", start))
		}
		result = max(result, energized(lines, start, LEFT))

		start = Point{i: i, j: -1}
		if IsPointInBoard(start, rows, cols) {
			panic(fmt.Sprintf("Starting position cannot be in grid. %v for RIGHT", start))
		}
		result = max(result, energized(lines, start, RIGHT))
	}

	return result
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
