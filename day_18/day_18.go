package day_18

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

var DOWN Point = Point{i: 1, j: 0}
var UP Point = Point{i: -1, j: 0}
var LEFT Point = Point{i: 0, j: -1}
var RIGHT Point = Point{i: 0, j: 1}

func PartOne(input_path string) int {
	lines, err := getLines(input_path)
	if err != nil {
		panic(err)
	}

	points := []Point{
		{
			i: 0, j: 0,
		},
	}

	current_point := Point{
		i: 0, j: 0,
	}

	for _, line := range lines {
		splited := strings.Split(line, " ")
		direction := GetDirection(splited[0])
		length, err := strconv.Atoi(splited[1])
		if err != nil {
			panic(err)
		}

		new_point := Point{
			i: current_point.i + (length * direction.i),
			j: current_point.j + (length * direction.j),
		}

		points = append(points, new_point)
		current_point = new_point
	}

	slices.Reverse(points)
	return areaPoints(points) + 1 + (perimeter(points))/2
}

func PartTwo(input_path string) int {
	lines, err := getLines(input_path)
	if err != nil {
		panic(err)
	}

	points := []Point{
		{
			i: 0, j: 0,
		},
	}

	current_point := Point{
		i: 0, j: 0,
	}

	for _, line := range lines {
		splited := strings.Split(line, " ")
		direction_s, length := convertHexToInstruction(splited[2][1:][:7])
		direction := GetDirection(direction_s)
		new_point := Point{
			i: current_point.i + (length * direction.i),
			j: current_point.j + (length * direction.j),
		}

		points = append(points, new_point)
		current_point = new_point
	}

	slices.Reverse(points)
	return areaPoints(points) + 1 + (perimeter(points))/2
}

func convertHexToInstruction(hex string) (direction string, length int) {

	hexadecimal_num := hex[1:6]
	decimal_num, err := strconv.ParseInt(hexadecimal_num, 16, 64)

	if err != nil {
		panic(err)
	}

	var direction_result string

	switch hex[len(hex)-1:] {
	case "0":
		direction_result = "R"
	case "1":
		direction_result = "D"
	case "2":
		direction_result = "L"
	case "3":
		direction_result = "U"
	}

	return direction_result, int(decimal_num)
}
func GetDirection(s string) Point {
	if s == "R" {
		return RIGHT
	}

	if s == "L" {
		return LEFT
	}

	if s == "U" {
		return UP
	}

	if s == "D" {
		return DOWN
	}

	panic(fmt.Sprintf("Invalid direction %s", s))
}

func perimeter(points []Point) int {
	var distance float64
	for i := 0; i < len(points); i++ {
		current := points[i]
		next := points[(i+1)%len(points)]
		distance += math.Sqrt(math.Pow(float64(current.i-next.i), 2) + math.Pow(float64(current.j-next.j), 2))
	}
	return int(distance)
}

func areaPoints(points []Point) int {
	var area int
	for i := 0; i < len(points); i++ {
		current := points[i]
		next := points[(i+1)%len(points)]
		area += current.i*next.j - next.i*current.j
	}
	return area / 2
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
