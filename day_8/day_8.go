package day8

import (
	"bufio"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
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

func parseLine(line string) (string, Node) {
	parts := strings.Split(line, " = ")

	input := parts[1]
	input = strings.Trim(input, "()")
	parts_2 := strings.Split(input, ", ")
	node := Node{parts_2[0], parts_2[1]}

	return parts[0], node
}

func PartOne(input_path string) int {
	lines, err := getLines(input_path)
	if err != nil {
		panic(err)
	}

	map_nodes := make(map[string]Node)

	for _, v := range lines[2:] {
		key, node := parseLine(v)
		map_nodes[key] = node
	}

	i := 0
	steps := 0
	instructions := lines[0]
	current_node := map_nodes["AAA"]
	current_key := "AAA"
	for {

		v := rune(instructions[i])
		if v == rune('L') {
			current_key = current_node.left
		} else if v == rune('R') {
			current_key = current_node.right
		} else {
			panic("Invalid instruction")
		}

		if i < len(instructions)-1 {
			i++
		} else {
			i = 0
		}

		steps++
		current_node = map_nodes[current_key]
		if current_key == "ZZZ" {
			break
		}
	}

	return steps
}

func PartTwo(input_path string) int {
	lines, err := getLines(input_path)
	if err != nil {
		panic(err)
	}

	map_nodes := make(map[string]Node)

	for _, v := range lines[2:] {
		key, node := parseLine(v)
		map_nodes[key] = node
	}

	i := 0
	steps := 0
	instructions := lines[0]
	current_node := map_nodes["AAA"]
	current_key := "AAA"
	for {

		v := rune(instructions[i])
		if v == rune('L') {
			current_key = current_node.left
		} else if v == rune('R') {
			current_key = current_node.right
		} else {
			panic("Invalid instruction")
		}

		if i < len(instructions)-1 {
			i++
		} else {
			i = 0
		}

		steps++
		current_node = map_nodes[current_key]
		if current_key == "ZZZ" {
			break
		}
	}

	return steps
}
