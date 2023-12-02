package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_game_id(line string) (int, error) {
	game_string := strings.Split(line, ":")
	result := strings.Split(game_string[0], " ")

	game_id, err := strconv.Atoi(result[1])
	if err != nil {
		return 0, err
	}

	return game_id, nil
}

func part_one(file_path string) (int, error) {
	total_sum := 0
	readFile, err := os.Open(file_path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		possible, err := is_possible_game(line)
		if err != nil {
			return 0, err
		}

		if possible {
			game_id, err := get_game_id(line)
			if err != nil {
				return 0, err
			}
			total_sum += game_id
		}
	}

	readFile.Close()

	return total_sum, nil
}

func is_possible_game(line string) (bool, error) {
	cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	result := strings.Split(line, ":")
	for _, turn := range result[1:] {
		for _, play := range strings.Split(turn, ";") {
			for _, step := range strings.Split(play, ",") {
				result := strings.Split(strings.TrimSpace(step), " ")
				count, err := strconv.Atoi(result[0])
				if err != nil {
					return false, err
				}

				color := result[1]

				if count > cubes[color] {
					return false, nil
				}
			}
		}
	}

	return true, nil
}
