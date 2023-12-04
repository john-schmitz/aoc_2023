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

func parseGame(line string) (int, int) {
	a := strings.Split(line, ":")

	numbers := strings.Split(a[1], "|")

	winning_numbers := removeEmptyStrings(strings.Split(strings.TrimSpace(numbers[0]), " "))
	owned_numbers := removeEmptyStrings(strings.Split(strings.TrimSpace(numbers[1]), " "))

	winning_numbers_map := make(map[string]bool)
	owned_numbers_map := make(map[string]bool)

	for _, v := range winning_numbers {
		winning_numbers_map[v] = true
	}

	for _, v := range owned_numbers {
		owned_numbers_map[v] = true
	}

	intersection := getIntersection(winning_numbers_map, owned_numbers_map)

	numbers_won := len(intersection)

	if numbers_won < 1 {
		return numbers_won, numbers_won
	}

	return numbers_won, int(math.Pow(float64(2), float64(numbers_won-1)))
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
		_, points := parseGame(line)
		total_sum += points
	}

	readFile.Close()

	return total_sum
}

func ProcessCopies() {

}

type Card struct {
	id          int
	owned_count int
	won_count   int
}

func PartTwo(file_path string) int {
	total_sum := 0
	readFile, err := os.Open(file_path)
	line_count := 0

	cards := make([]Card, 0)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		won_count, _ := parseGame(line)
		cards = append(cards, Card{owned_count: 1, won_count: won_count, id: line_count + 1})
		line_count += 1

	}

	for index, card := range cards {
		cards = playCard(card, index, cards)
	}

	for _, card := range cards {
		total_sum += card.owned_count
	}

	readFile.Close()

	return total_sum
}

func playCard(card Card, index int, cards []Card) []Card {
	if card.won_count == 0 {
		return cards
	}

	for a := 0; a < card.owned_count; a++ {
		for i := 0; i < card.won_count; i++ {
			cards[index+i+1].owned_count += 1
		}
	}
	return cards
}
