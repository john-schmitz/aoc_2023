package day10

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	i int
	j int
}

func PartOne(file_path string) int {
	lines, err := getLines(file_path)
	if err != nil {
		panic(err)
	}

	start := *new(Point)

	var pipes_map [][]rune
	unvisited_set := map[string]bool{}
	distances := make([][]int, len(lines))

	for i := range lines {
		for j := range lines[0] {
			key := getKey(i, j)
			unvisited_set[key] = false
		}
	}

	for i := range distances {
		distances[i] = make([]int, len(lines[0]))
		for j := range distances[i] {
			distances[i][j] = math.MaxInt
		}
	}

	for line_index, line := range lines {
		var char_coll []rune
		for char_index, c := range line {
			if c == rune('S') { // S
				start = Point{line_index, char_index}
			}

			char_coll = append(char_coll, c)
		}
		pipes_map = append(pipes_map, char_coll)
	}

	distances[start.i][start.j] = 0

	traverse(start, pipes_map, unvisited_set, distances)

	max_distance := math.MinInt

	for i := 0; i < len(distances); i++ {
		for j := 0; j < len(distances[0]); j++ {
			element := distances[i][j]
			if element != math.MaxInt {
				max_distance = max(element, max_distance)
			}
		}
	}

	return max_distance
}

func PartTwo(file_path string) int {
	return 0
}

func traverse(start Point, board [][]rune, unvisited map[string]bool, distances [][]int) {
	current_point := start
	for {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				point := Point{i: i + current_point.i, j: current_point.j + j}

				if !IsPointInBoard(point, board) {
					continue
				}

				_, present := unvisited[getKey(point.i, point.j)]

				can_access := canAccess(board[current_point.i][current_point.j], board[point.i][point.j], current_point, point)
				if IsPointAdjacent(current_point, point) && present && can_access {
					current_distance := distances[current_point.i][current_point.j]
					new_distance := current_distance + 1
					point_distance := distances[point.i][point.j]

					if new_distance < point_distance {
						distances[point.i][point.j] = new_distance
					}
				}
			}
		}

		delete(unvisited, getKey(current_point.i, current_point.j))
		current_point = next_point_unvisited(unvisited, distances)
		if !IsPointInBoard(current_point, board) {
			break
		}
	}
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

func IsPointInBoard(p Point, board [][]rune) bool {
	return p.i >= 0 && p.j >= 0 && p.i < len(board) && p.j < len(board[p.i])
}

func IsPointAdjacent(p1 Point, p2 Point) bool {
	dx := absDiffInt(p1.i - p2.i)
	dy := absDiffInt(p1.j - p2.j)

	return dx+dy == 1
}

func absDiffInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func canAccess(from rune, to rune, from_point Point, to_point Point) bool {
	if !IsPointAdjacent(from_point, to_point) {
		return false
	}
	if from == rune('.') || to == rune('.') {
		return false
	}

	map_pipe_directions := map[rune][]struct {
		i int
		j int
	}{
		rune('|'): {
			{i: -1, j: 0},
			{i: +1, j: 0},
		},
		rune('L'): {
			{i: -1, j: 0},
			{i: 0, j: 1},
		},
		rune('-'): {
			{i: 0, j: -1},
			{i: 0, j: 1},
		},
		rune('J'): {
			{i: -1, j: 0},
			{i: 0, j: -1},
		},
		rune('F'): {
			{i: 1, j: 0},
			{i: 0, j: 1},
		},
		rune('7'): {
			{i: 1, j: 0},
			{i: 0, j: -1},
		},
		rune('S'): {
			{i: 0, j: 1},
			{i: 0, j: -1},
			{i: -1, j: 0},
			{i: +1, j: 0},
		},
	}

	directions_from, present := map_pipe_directions[from]

	if !present {
		panic(fmt.Sprintf("Did not find directions for %c", from))
	}

	directions_to, present := map_pipe_directions[to]

	if !present {
		panic(fmt.Sprintf("Did not find directions for %c", to))
	}

	can_from := calculateDirections(directions_from, to_point, from_point)
	can_to := calculateDirections(directions_to, from_point, to_point)

	return can_from && can_to
}

func calculateDirections(directions_to []struct {
	i int
	j int
}, from_point Point, to_point Point) bool {
	for _, direction := range directions_to {
		result_i := from_point.i - to_point.i
		result_j := from_point.j - to_point.j

		if result_i == direction.i && result_j == direction.j {
			return true
		}
	}
	return false
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
