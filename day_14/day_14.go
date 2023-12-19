package day14

import (
	"bufio"
	"os"
)

type Point struct {
	i int
	j int
}

var directions = []Point{
	{i: -1, j: 0}, // NORTH
	{i: 0, j: -1}, // WEST
	{i: 1, j: 0},  // SOUTH
	{i: 0, j: 1},  // EAST
}

func isPointInBoard(board [][]rune, point Point) bool {
	return point.i >= 0 && point.j >= 0 && point.i < len(board) && point.j < len(board[point.i])
}

func PartOne(input_file string) int {
	lines, err := getLines(input_file)
	if err != nil {
		panic(err)
	}

	direction := directions[0]
	grid := make([][]rune, 0)

	for _, row := range lines {
		grid = append(grid, []rune(row))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			elm := grid[i][j]
			if elm == rune('#') || elm == rune('.') {
				continue
			}

			next_point := Point{i: i + (1 * direction.i), j: j + (1 * direction.j)}

			if !isPointInBoard(grid, next_point) {
				continue
			}

			above_count := 0
			for z := 1; z <= i; z++ {
				point_x := Point{i: i + (z * direction.i), j: j + (z * direction.j)}
				x := grid[point_x.i][point_x.j]
				if x == rune('.') {
					above_count++
				} else {
					break
				}
			}

			if above_count > 0 {
				grid[i][j], grid[i-above_count][j] = grid[i-above_count][j], grid[i][j]
			}
		}
	}

	acc := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			elm := grid[i][j]
			if elm == rune('O') {
				score := len(grid) - i
				acc = acc + score
			}
		}
	}

	return acc
}

func PartTwo(input_file string) int {
	return 0
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
