package day_17

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/rdleal/go-priorityq/kpq"
)

type Point struct {
	i int
	j int
}

type Node struct {
	point     Point
	direction Point
	step      int
	heat_loss int
}

func (point Point) IsPointInBoard(rows, cols int) bool {
	return point.i >= 0 && point.j >= 0 && point.i < rows && point.j < cols
}

func (p1 Point) IsPointAdjacent(p2 Point) bool {
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

func parseLines(lines []string) map[Point]int {
	rows := len(lines)
	cols := len(lines[0])
	output := map[Point]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			elm := string(lines[i][j])
			parsed_elm, err := strconv.Atoi(elm)
			if err != nil {
				panic(err)
			}
			output[Point{i, j}] = parsed_elm
		}
	}
	return output
}

func PartOne(input_file string) int {
	directions := []Point{
		{1, 0},  // DOWN
		{-1, 0}, // UP
		{0, 1},  // RIGHT
		{0, -1}, // LEFT
	}
	stoped := Point{0, 0}
	lines, err := getLines(input_file)
	if err != nil {
		panic(err)
	}
	cmp := func(a, b Node) bool {
		return a.heat_loss < b.heat_loss
	}

	rows := len(lines)
	cols := len(lines[0])
	grid := parseLines(lines)
	seen := map[string]bool{}

	pq := kpq.NewKeyedPriorityQueue[string](cmp)
	start := Node{
		point:     Point{0, 0},
		direction: Point{0, 0},
		heat_loss: 0,
		step:      0,
	}

	pq.Push(getKey(start), start)

	for pq.Len() > 0 {
		current_key, current_node, _ := pq.Pop()

		if current_node.point.i == rows-1 && current_node.point.j == cols-1 {
			return current_node.heat_loss
		}

		if !current_node.point.IsPointInBoard(rows, cols) {
			continue
		}

		if seen[current_key] {
			continue
		}

		seen[current_key] = true

		if current_node.step < 3 && current_node.direction != stoped {
			next_point := Point{
				i: current_node.point.i + current_node.direction.i,
				j: current_node.point.j + current_node.direction.j,
			}

			if next_point.IsPointInBoard(rows, cols) {
				found_heat_loss, present := grid[next_point]
				if !present {
					panic(found_heat_loss)
				}

				next_node := Node{
					point:     next_point,
					direction: current_node.direction,
					step:      current_node.step + 1,
					heat_loss: current_node.heat_loss + found_heat_loss,
				}
				pq.Push(getKey(next_node), next_node)
			}
		}

		for _, direction := range directions {
			if direction.i == current_node.direction.i && direction.j == current_node.direction.j {
				continue
			}

			if direction.i == -current_node.direction.i && direction.j == -current_node.direction.j {
				continue
			}

			next_point := Point{
				i: current_node.point.i + direction.i,
				j: current_node.point.j + direction.j,
			}

			if next_point.IsPointInBoard(rows, cols) {
				found_heat_loss, present := grid[next_point]
				if !present {
					panic(found_heat_loss)
				}

				next_node := Node{
					point:     next_point,
					direction: direction,
					step:      1,
					heat_loss: current_node.heat_loss + found_heat_loss,
				}
				pq.Push(getKey(next_node), next_node)
			}
		}
	}

	return 0
}

func getKey(start Node) string {
	key := fmt.Sprintf("i:%dj:%ddi:%ddj:%dh:%d", start.point.i, start.point.j, start.direction.i, start.direction.j, start.step)
	return key
}

func PartTwo(input_file string) int {
	directions := []Point{
		{1, 0},  // DOWN
		{-1, 0}, // UP
		{0, 1},  // RIGHT
		{0, -1}, // LEFT
	}
	stoped := Point{0, 0}
	lines, err := getLines(input_file)
	if err != nil {
		panic(err)
	}
	cmp := func(a, b Node) bool {
		return a.heat_loss < b.heat_loss
	}

	rows := len(lines)
	cols := len(lines[0])
	grid := parseLines(lines)
	seen := map[string]bool{}

	pq := kpq.NewKeyedPriorityQueue[string](cmp)
	start := Node{
		point:     Point{0, 0},
		direction: Point{0, 1},
		heat_loss: 0,
		step:      0,
	}

	pq.Push(getKey(start), start)

	for pq.Len() > 0 {
		current_key, current_node, _ := pq.Pop()

		if current_node.point.i == rows-1 && current_node.point.j == cols-1 {
			return current_node.heat_loss
		}

		if !current_node.point.IsPointInBoard(rows, cols) {
			continue
		}

		if seen[current_key] {
			continue
		}

		seen[current_key] = true

		if current_node.step < 10 && current_node.direction != stoped {
			next_point := Point{
				i: current_node.point.i + current_node.direction.i,
				j: current_node.point.j + current_node.direction.j,
			}

			if next_point.IsPointInBoard(rows, cols) {
				found_heat_loss, present := grid[next_point]
				if !present {
					panic(found_heat_loss)
				}

				next_node := Node{
					point:     next_point,
					direction: current_node.direction,
					step:      current_node.step + 1,
					heat_loss: current_node.heat_loss + found_heat_loss,
				}
				pq.Push(getKey(next_node), next_node)
			}
		}
		if current_node.step > 3 {
			for _, direction := range directions {
				if direction.i == current_node.direction.i && direction.j == current_node.direction.j {
					continue
				}

				if direction.i == -current_node.direction.i && direction.j == -current_node.direction.j {
					continue
				}

				next_point := Point{
					i: current_node.point.i + direction.i,
					j: current_node.point.j + direction.j,
				}

				if next_point.IsPointInBoard(rows, cols) {
					found_heat_loss, present := grid[next_point]
					if !present {
						panic(found_heat_loss)
					}

					next_node := Node{
						point:     next_point,
						direction: direction,
						step:      1,
						heat_loss: current_node.heat_loss + found_heat_loss,
					}
					pq.Push(getKey(next_node), next_node)
				}
			}
		}
	}

	return 0
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
