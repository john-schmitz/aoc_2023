package day4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func getIntersection(a, b map[string]bool) map[string]bool {
	if len(a) > len(b) {
		a, b = b, a
	}
	s_intersection := map[string]bool{}
	for k := range a {
		if b[k] {
			s_intersection[k] = true
		}
	}
	return s_intersection

}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func parseGame(line string) int {
	a := strings.Split(line, ":")

	numbers := strings.Split(a[1], "|")

	winning_numbers := removeEmptyStrings(strings.Split(strings.TrimSpace(numbers[0]), " "))
	owned_numbers := removeEmptyStrings(strings.Split(strings.TrimSpace(numbers[1]), " "))

	winning_numbers_map := make(map[string]bool)
	owned_numbers_map := make(map[string]bool)

	// filter empty strings

	for _, v := range winning_numbers {
		winning_numbers_map[v] = true
	}

	for _, v := range owned_numbers {
		owned_numbers_map[v] = true
	}

	intersection := getIntersection(winning_numbers_map, owned_numbers_map)

	intersection_length := len(intersection)
	if intersection_length == 0 {
		return 0
	}

	if intersection_length == 1 {
		return 1
	}

	return int(math.Pow(float64(2), float64(intersection_length-1)))
}

func PartOne(file_path string) int {
	total_sum := 0
	readFile, err := os.Open(file_path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		total_sum += parseGame(line)
	}

	readFile.Close()

	return total_sum
}

func PartTwo(file_path string) int {
	total_sum := 0
	readFile, err := os.Open(file_path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		total_sum += parseGame(line)
	}

	readFile.Close()

	return total_sum
}
