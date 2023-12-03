package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	i int
	j int
}

type Range struct {
	start Point
	end   Point
}

func isPointInBoard(board [][]string, point Point) bool {
	return point.i >= 0 && point.j >= 0 && point.i < len(board) && point.j < len(board[point.i])
}

func isPointSymbol(board [][]string, point Point) bool {
	if !isPointInBoard(board, point) {
		return false
	}

	x := board[point.i][point.j]
	if _, err := strconv.Atoi(x); err == nil || x == "." {
		return false
	}

	return true
}

func IsPointAdjacentToSymbol(board [][]string, point Point) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			neighbor := Point{point.i + i, point.j + j}
			if isPointSymbol(board, neighbor) {
				return true
			}
		}
	}
	return false
}

func IsNumberAdjacentToSymbol(board [][]string, input Range) bool {
	if input.start.i != input.end.i {
		panic("Range must be horizontal.")
	}

	for i := input.start.j; i <= input.end.j; i++ {
		if IsPointAdjacentToSymbol(board, Point{input.start.i, i}) {
			return true
		}
	}

	return false
}

func parseInput(input_path string) [][]string {
	f, err := os.Open(input_path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	board := make([][]string, 0)
	line_count := 0

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, make([]string, 0))
		for _, char := range line {
			board[line_count] = append(board[line_count], string(char))
		}
		line_count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return board
}

func SumPartNumbers(input_path string) int {

	board := parseInput(input_path)
	total := 0

	for index, line := range board {
		currently_parsing_number := false
		current_number := ""

		last_index := len(line) - 1
		for jindex, char := range line {
			if _, err := strconv.Atoi(char); err == nil {
				currently_parsing_number = true
				current_number += char

				if last_index != jindex {
					continue
				}
			}

			if currently_parsing_number {
				start_point := Point{index, jindex - len(current_number)}
				end_point := Point{index, jindex - 1}

				if IsNumberAdjacentToSymbol(board, Range{start_point, end_point}) {
					parsed_number, _ := strconv.Atoi(current_number)
					total += parsed_number
				}

				currently_parsing_number = false
				current_number = ""
			}
		}
	}

	return total
}

func SumGearRatios(input_path string) int {
	board := parseInput(input_path)
	total := 0
	numbers := getNumbers(board)

	for index, line := range board {
		for jindex, char := range line {
			if char != "*" {
				continue
			}

			current_point := Point{index, jindex}
			found := make([]Range, 0)
			for _, number := range numbers {
				for jindex := number.start.j; jindex <= number.end.j; jindex++ {
					if IsPointAdjacent(current_point, Point{number.start.i, jindex}) {
						found = append(found, number)
						break
					}
				}
			}

			if len(found) == 2 {
				first_number := found[0]
				second_number := found[1]

				first_number_value, err := strconv.Atoi(strings.Join(board[first_number.start.i][first_number.start.j:first_number.end.j+1], ""))
				if err != nil {
					panic(err)
				}

				second_number_value, err := strconv.Atoi(strings.Join(board[second_number.start.i][second_number.start.j:second_number.end.j+1], ""))
				if err != nil {
					panic(err)
				}
				total += first_number_value * second_number_value
			}
		}
	}

	return total
}

func isPointDigit(board [][]string, point Point) bool {
	if !isPointInBoard(board, point) {
		return false
	}

	x := board[point.i][point.j]
	if _, err := strconv.Atoi(x); err == nil {
		return true
	}

	return false
}

func getNumbers(board [][]string) []Range {
	numbers := make([]Range, 0)

	for index, line := range board {
		current_number := ""
		for jindex, char := range line {
			if isPointDigit(board, Point{index, jindex}) {
				current_number += char
				is_next_point_digit := isPointDigit(board, Point{index, jindex + 1})

				if !is_next_point_digit {
					start_point := Point{index, jindex - len(current_number) + 1}
					end_point := Point{index, jindex}

					numbers = append(numbers, Range{start_point, end_point})

					current_number = ""
				}
			}
		}
	}
	return numbers
}

func IsPointAdjacent(p1 Point, p2 Point) bool {
	dx := absDiffInt(p1.i - p2.i)
	dy := absDiffInt(p1.j - p2.j)

	return dx+dy == 1 || (dx == 1 && dy == 1)
}

func absDiffInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
