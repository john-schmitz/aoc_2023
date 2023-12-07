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
			expected: Hand{bid: 765, cards: "23456", strength: HighCard},
		}, {
			line:     "32T3K 765",
			expected: Hand{bid: 765, cards: "32T3K", strength: OnePair},
		}, {
			line:     "23432 765",
			expected: Hand{bid: 765, cards: "23432", strength: TwoPair},
		}, {
			line:     "TTT98 765",
			expected: Hand{bid: 765, cards: "TTT98", strength: ThreeOfaKind},
		}, {
			line:     "23332 765",
			expected: Hand{bid: 765, cards: "23332", strength: FullHouse},
		}, {
			line:     "AAAA8 765",
			expected: Hand{bid: 765, cards: "AAAA8", strength: FourOfAKind},
		}, {
			line:     "AAAAA 765",
			expected: Hand{bid: 765, cards: "AAAAA", strength: FiveOfAKind},
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
	testCases := []struct {
		hand     Hand
		hand2    Hand
		expected bool
	}{
		{
			hand:     Hand{bid: 765, cards: "12345", strength: HighCard},
			hand2:    Hand{bid: 765, cards: "AAAA8", strength: FourOfAKind},
			expected: true,
		},
		{
			hand:     Hand{bid: 765, cards: "KKKKK", strength: FiveOfAKind},
			hand2:    Hand{bid: 765, cards: "AAAAA", strength: FiveOfAKind},
			expected: true,
		},
		{
			hand:     Hand{bid: 765, cards: "AAAAA", strength: FiveOfAKind},
			hand2:    Hand{bid: 765, cards: "KKKKK", strength: FiveOfAKind},
			expected: false,
		},
		{
			hand:     Hand{bid: 765, cards: "KTJJT", strength: FiveOfAKind},
			hand2:    Hand{bid: 765, cards: "KK677", strength: FiveOfAKind},
			expected: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.hand.cards, func(t *testing.T) {
			actual := compareHands(tC.hand, tC.hand2)
			if actual != tC.expected {
				t.Errorf("Expected compareHands(%v, %v) = %t. Got %t", tC.hand, tC.hand2, tC.expected, actual)
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
