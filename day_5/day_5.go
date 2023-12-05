package day5

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Ranges struct {
	destination_range int
	source_range      int
	range_length      int
}

func parseRanges(line string) (Ranges, error) {
	re := regexp.MustCompile("[0-9]+")

	ranges_string := re.FindAllString(line, -1)

	if len(ranges_string) != 3 {
		return *new(Ranges), fmt.Errorf("expected 3 ranges, got %v. For line: %s", len(ranges_string), line)
	}

	destination_range, err := strconv.Atoi(ranges_string[0])
	if err != nil {
		return *new(Ranges), err
	}

	source_range, err := strconv.Atoi(ranges_string[1])
	if err != nil {
		return *new(Ranges), err
	}

	range_length, err := strconv.Atoi(ranges_string[2])
	if err != nil {
		return *new(Ranges), err
	}

	return Ranges{destination_range, source_range, range_length}, nil
}

func readSeedsFromInput(line string) []int {
	re := regexp.MustCompile("[0-9]+")

	seeds_string := re.FindAllString(line, -1)

	output := make([]int, len(seeds_string))

	for i, seed := range seeds_string {
		seed, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}

		output[i] = seed
	}

	return output
}

func mapRange(id int, ranges Ranges) int {
	is_in_range := ranges.source_range <= id && ranges.source_range+ranges.range_length >= id

	if is_in_range {
		return ranges.destination_range + (id - ranges.source_range)
	}

	return id
}

func PartOne(file_path string) int {
	lowest_location_number := math.MaxInt
	seeds, steps, err := parseInput(file_path)

	seed_steps := make([][]int, len(seeds))

	for i, seed := range seeds {
		seed_steps[i] = make([]int, len(steps)+1)
		seed_steps[i][0] = seed
	}

	if err != nil {
		panic(err)
	}

	for step_index, step := range steps {
	seed_label:
		for index, seed := range seeds {
			for _, ranges := range step {
				maped_seed := mapRange(seed, ranges)
				seeds[index] = maped_seed
				seed_steps[index][step_index+1] = maped_seed
				if seed != maped_seed {
					continue seed_label
				}
			}
		}
	}

	for _, seed := range seed_steps {
		last_step := seed[len(seed)-1]
		lowest_location_number = min(last_step, lowest_location_number)
	}

	return lowest_location_number
}

func PartTwo(file_path string) int {
	lowest_location_number := math.MaxInt
	seeds, steps, err := parseInput(file_path)

	seeds = convertSeedsToRanges(seeds)

	if err != nil {
		panic(err)
	}

	for _, step := range steps {
	seed_label:
		for index, seed := range seeds {
			for _, ranges := range step {
				maped_seed := mapRange(seed, ranges)
				seeds[index] = maped_seed
				if seed != maped_seed {
					continue seed_label
				}
			}
		}
	}

	for _, seed := range seeds {
		lowest_location_number = min(seed, lowest_location_number)
	}

	return lowest_location_number
}

func parseInput(file_path string) ([]int, [][]Ranges, error) {
	steps := make([][]Ranges, 7)
	current_step_index := -1

	fileBytes, err := os.ReadFile(file_path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	seeds := readSeedsFromInput(sliceData[0])
	sliceData = sliceData[2:]

	for _, line := range sliceData {
		if strings.Contains(line, "map") {
			current_step_index++
			continue
		}

		if line == "" {
			continue
		}

		ranges, err := parseRanges(line)
		if err != nil {
			return nil, nil, err
		}

		current_step := steps[current_step_index]

		if current_step == nil {
			current_step = make([]Ranges, 0)
		}

		steps[current_step_index] = append(current_step, ranges)
	}

	return seeds, steps, nil
}

func convertSeedsToRanges(seeds []int) []int {
	if len(seeds)%2 != 0 {
		panic("seeds must be even")
	}

	output := make([]int, 0)

	for i := 0; i < len(seeds); i = i + 2 {
		start := seeds[i]
		length := seeds[i+1]

		for j := 0; j < length; j++ {
			output = append(output, start+j)
		}
	}

	return output
}
