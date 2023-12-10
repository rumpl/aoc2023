package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var s = map[byte]int{
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func countCards(cards string) map[string]int {
	m := make(map[string]int)
	for _, card := range cards {
		m[string(card)]++
	}
	if _, ok := m["J"]; !ok {
		return m
	}
	if m["J"] == 5 {
		return map[string]int{"AAAAA": 5}
	}

	max := 0
	kmax := " "
	for k, v := range m {
		if k == "J" {
			continue
		}
		if v == max && s[k[0]] > s[kmax[0]] {
			max = v
			kmax = k
			continue
		}
		if v > max {
			max = v
			kmax = k
		}
	}
	m[kmax] += m["J"]
	delete(m, "J")
	return m
}

func rank(cards string) int {
	c := countCards(cards)

	// five of a kind
	if len(c) == 1 {
		return 7
	}
	if len(c) == 2 {
		for _, v := range c {
			// 4 of a kind
			if v == 4 {
				return 6
			}
			if v == 3 {
				for _, v := range c {
					// full house
					if v == 2 {
						return 5
					}
				}
			}
		}
	}

	// two pair
	if len(c) == 3 {
		for _, v := range c {
			// three of a kind
			if v == 3 {
				return 4
			}
		}
		// two pair
		return 3
	}

	// one pair
	if len(c) == 4 {
		return 2
	}
	// high card
	return 1
}

type hand struct {
	Cards string
	Bid   int
}

func main() {
	b, _ := os.ReadFile("input.txt")
	data := string(b)
	lines := strings.Split(data, "\n")
	hands := []hand{}
	for _, line := range lines {
		cards := strings.Split(line, " ")
		n, _ := strconv.Atoi(cards[1])
		hands = append(hands, hand{Cards: cards[0], Bid: n})
	}

	slices.SortFunc[[]hand](hands, func(i, j hand) int {
		irank := rank(i.Cards)
		jrank := rank(j.Cards)
		if irank > jrank {
			return 1
		}
		if irank < jrank {
			return -1
		}

		for k := 0; k < 5; k++ {
			if s[i.Cards[k]] > s[j.Cards[k]] {
				return 1
			}
			if s[i.Cards[k]] < s[j.Cards[k]] {
				return -1
			}
		}
		return 0
	})

	total := 0
	for i := 0; i < len(hands); i++ {
		total += hands[i].Bid * (i + 1)
	}
	fmt.Println(total)
}
