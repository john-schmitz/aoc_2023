package day7

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		line     string
		expected Hand
	}{
		{
			line:     "23456 765",
			expected: Hand{bid: 765, cards: "23456"},
		}, {
			line:     "32T3K 765",
			expected: Hand{bid: 765, cards: "32T3K"},
		}, {
			line:     "23432 765",
			expected: Hand{bid: 765, cards: "23432"},
		}, {
			line:     "TTT98 765",
			expected: Hand{bid: 765, cards: "TTT98"},
		}, {
			line:     "23332 765",
			expected: Hand{bid: 765, cards: "23332"},
		}, {
			line:     "AAAA8 765",
			expected: Hand{bid: 765, cards: "AAAA8"},
		}, {
			line:     "AAAAA 765",
			expected: Hand{bid: 765, cards: "AAAAA"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.line, func(t *testing.T) {
			actual := parseLine(tC.line)
			if actual != tC.expected {
				t.Errorf("Expected parseLine(%s) = %v. Got %v", tC.line, tC.expected, actual)
			}

		})
	}
}

func TestComparison(t *testing.T) {
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
	testCases := []struct {
		hand     Hand
		hand2    Hand
		expected bool
	}{
		{
			hand:     Hand{bid: 765, cards: "12345"},
			hand2:    Hand{bid: 765, cards: "AAAA8"},
			expected: true,
		},
		{
			hand:     Hand{bid: 765, cards: "KKKKK"},
			hand2:    Hand{bid: 765, cards: "AAAAA"},
			expected: true,
		},
		{
			hand:     Hand{bid: 765, cards: "AAAAA"},
			hand2:    Hand{bid: 765, cards: "KKKKK"},
			expected: false,
		},
		{
			hand:     Hand{bid: 765, cards: "KTJJT"},
			hand2:    Hand{bid: 765, cards: "KK677"},
			expected: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.hand.cards, func(t *testing.T) {
			actual := compareHands(tC.hand, tC.hand2, card_strengths)
			if actual != tC.expected {
				t.Errorf("Expected compareHands(%v, %v) = %t. Got %t", tC.hand, tC.hand2, tC.expected, actual)
			}
		})
	}
}

func TestJokerHandStrength(t *testing.T) {
	testCases := []struct {
		cards    string
		expected Strength
	}{
		{
			cards:    "KTJJT",
			expected: FourOfAKind,
		}, {
			cards:    "JJJJA",
			expected: FiveOfAKind,
		},
		{
			cards:    "JJJJJ",
			expected: FiveOfAKind,
		},
		{
			cards:    "T55J5",
			expected: FourOfAKind,
		},
		{
			cards:    "QQQJA",
			expected: FourOfAKind,
		},
		{
			cards:    "2233J",
			expected: FullHouse,
		},
		{
			cards:    "1233J",
			expected: ThreeOfaKind,
		},
		{
			cards:    "1234J",
			expected: OnePair,
		},
		{
			cards:    "1244J",
			expected: ThreeOfaKind,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.cards, func(t *testing.T) {
			actual := jokerHandStrength(tC.cards)
			if actual != tC.expected {
				t.Errorf("Expected jokerHandStrength(%s) = %d. Got %d", tC.cards, tC.expected, actual)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	testCases := []struct {
		file_path string
		expected  int
	}{
		{
			file_path: "sample.txt",
			expected:  6440,
		}, {
			file_path: "input.txt",
			expected:  251545216,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.file_path, func(t *testing.T) {
			actual := PartOne(tC.file_path)
			if actual != tC.expected {
				t.Errorf("Expected PartOne(%s) = %d. Got %d", tC.file_path, tC.expected, actual)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		file_path string
		expected  int
	}{
		{
			file_path: "sample.txt",
			expected:  5905,
		},
		{
			file_path: "input.txt",
			expected:  250384185,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.file_path, func(t *testing.T) {
			actual := PartTwo(tC.file_path)
			if actual != tC.expected {
				t.Errorf("Expected PartOne(%s) = %d. Got %d", tC.file_path, tC.expected, actual)
			}
		})
	}
}

func TestJokerComparison(t *testing.T) {
	card_strengths := map[rune]int{
		rune('A'): 12,
		rune('K'): 11,
		rune('Q'): 10,
		rune('T'): 8,
		rune('9'): 7,
		rune('8'): 6,
		rune('7'): 5,
		rune('6'): 4,
		rune('5'): 3,
		rune('4'): 2,
		rune('3'): 1,
		rune('2'): 0,
		rune('J'): -1,
	}
	testCases := []struct {
		hand     Hand
		hand2    Hand
		expected bool
	}{
		{
			hand:     Hand{bid: 765, cards: "T55J5"},
			hand2:    Hand{bid: 765, cards: "KTJJT"},
			expected: true,
		},
		{
			hand:     Hand{bid: 765, cards: "T55J5"},
			hand2:    Hand{bid: 765, cards: "QQQJA"},
			expected: true,
		},
		{
			hand:     Hand{bid: 765, cards: "QQQJA"},
			hand2:    Hand{bid: 765, cards: "KTJJT"},
			expected: true,
		},
		{
			hand:     Hand{bid: 765, cards: "KTJJT"},
			hand2:    Hand{bid: 765, cards: "QQQJA"},
			expected: false,
		},
		{
			hand:     Hand{bid: 765, cards: "32T3K"},
			hand2:    Hand{bid: 765, cards: "KTJJT"},
			expected: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.hand.cards, func(t *testing.T) {
			actual := compareHandsJoker(tC.hand, tC.hand2, card_strengths)
			if actual != tC.expected {
				t.Errorf("Expected compareHands(%v, %v) = %t. Got %t", tC.hand, tC.hand2, tC.expected, actual)
			}
		})
	}
}
