package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	duration_milliseconds int
	distance_millimeters  int
}

func isRaceWon(distance, record int) bool {
	return distance > record
}

func parseRaces(input []string) []Race {
	durations_f, _ := strings.CutPrefix(input[0], "Time:")
	distances_f, _ := strings.CutPrefix(input[1], "Distance:")

	durations := strings.Fields(durations_f)
	distances := strings.Fields(distances_f)

	if len(durations) != len(distances) {
		fmt.Println("Error: distances and durations are not the same length")
		os.Exit(1)
	}

	outputs := make([]Race, len(durations))

	for i := range durations {
		n, err := strconv.Atoi(distances[i])
		if err != nil {
			panic(err)
		}

		distance := n

		n, err = strconv.Atoi(durations[i])
		if err != nil {
			panic(err)
		}

		duration := n

		outputs[i] = Race{duration, distance}
	}

	return outputs
}

func PartOne(file_path string) int {
	acc := 1
	races := parseRaces(getLines(file_path))

	for _, race := range races {
		acc = acc * waysRaceCouldBeWon(race)
	}

	return acc
}

func getLines(file_path string) []string {
	fileBytes, err := os.ReadFile(file_path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	return sliceData
}

func waysRaceCouldBeWon(race Race) int {
	count := 0
	for i := 0; i < race.duration_milliseconds; i++ {
		duration := i
		remaning_time := race.duration_milliseconds - i
		distance := duration * remaning_time
		if isRaceWon(distance, race.distance_millimeters) {
			count++
		}
	}

	return count
}
