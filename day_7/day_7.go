package day7

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Strength int64

const (
	HighCard Strength = iota
	OnePair
	TwoPair
	ThreeOfaKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	cards    string
	bid      int
	strength Strength
}

func countCharacters(str string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range str {
		frequency[char] = frequency[char] + 1
	}

	return frequency
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

func handStrength(cards string) Strength {
	map_of_occurences := countCharacters(cards)
	map_size := len(map_of_occurences)
	if map_size == 5 {
		return HighCard
	}

	if map_size == 4 {
		return OnePair
	}

	if map_size == 1 {
		return FiveOfAKind
	}

	if map_size == 2 {
		for _, v := range map_of_occurences {
			if v == 1 || v == 4 {
				return FourOfAKind
			}

			return FullHouse
		}
	}

	if map_size == 3 {
		pair_count := 0
		for _, v := range map_of_occurences {
			if v == 2 {
				pair_count++
			}
		}

		if pair_count == 2 {
			return TwoPair
		}
		return ThreeOfaKind
	}
	panic("Invalid strength")
}

func parseLine(line string) Hand {
	fields := strings.Fields(line)

	cards := fields[0]
	bid, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}

	return Hand{cards: cards, bid: bid, strength: handStrength(cards)}
}

func compareHands(hand1, hand2 Hand) bool {
	card_strengths := map[rune]int{
		rune('A'): 12,
		rune('K'): 11,
		rune('Q'): 10,
		rune('J'): 9,
		rune('T'): 8,
		rune('9'): 7,
		rune('8'): 6,
		rune('7'): 5,
		rune('6'): 4,
		rune('5'): 3,
		rune('4'): 2,
		rune('3'): 1,
		rune('2'): 0,
	}

	if hand1.strength == hand2.strength {
		for index := range hand1.cards {
			rune_1 := rune(hand1.cards[index])
			rune_2 := rune(hand2.cards[index])
			if card_strengths[rune_1] == card_strengths[rune_2] {
				continue
			}

			return card_strengths[rune_1] < card_strengths[rune_2]
		}
	}

	return hand1.strength < hand2.strength
}

func PartOne(input_path string) int {
	lines, err := getLines(input_path)
	if err != nil {
		panic(err)
	}
	hands := make([]Hand, 0)

	for _, line := range lines {
		hands = append(hands, parseLine(line))
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})

	total_bids := 0

	for hand_index, hand := range hands {
		total_bids = total_bids + (hand_index+1)*hand.bid
	}

	return total_bids
}
