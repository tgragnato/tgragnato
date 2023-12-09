package camelcards

import (
	"fmt"
	"sort"
)

func SameScore(char rune, joker bool) uint {
	switch char {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if joker {
			return 1
		}
		return 11
	case 'T':
		return 10
	case '9':
		return 9
	case '8':
		return 8
	case '7':
		return 7
	case '6':
		return 6
	case '5':
		return 5
	case '4':
		return 4
	case '3':
		return 3
	case '2':
		return 2
	default:
		return 0
	}
}

type Hand struct {
	Cards [5]rune
	Bid   uint
}

func (h *Hand) GetMatches(joker bool) uint {
	counts := make(map[rune]int)
	for _, r := range h.Cards {
		counts[r]++
	}

	if joker {
		for counts['J'] > 0 {
			vMaxNotJoker := -1
			var card rune
			for k, v := range counts {
				if k == 'J' {
					continue
				}
				if v > vMaxNotJoker {
					vMaxNotJoker = v
					card = k
				}
			}
			counts['J']--
			counts[card]++
		}
	}

	values := sort.IntSlice{}
	for _, v := range counts {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(values))

	switch values[0] {
	case 5:
		return 7
	case 4:
		return 6
	case 3:
		if values[1] == 2 {
			return 5
		}
		return 4
	case 2:
		if values[1] == 2 {
			return 3
		}
		return 2
	case 1:
		return 1
	default:
		return 0
	}
}

func (h *Hand) String() string {
	return fmt.Sprintf("Cards: [%s %s %s %s %s], Bid: %d",
		string(h.Cards[0]), string(h.Cards[1]),
		string(h.Cards[2]), string(h.Cards[3]),
		string(h.Cards[4]), h.Bid,
	)
}

type HandList struct {
	Hands []Hand
}

func (l *HandList) GetWinnings(joker bool) uint {
	sort.Slice(l.Hands, func(i, j int) bool {

		matches1 := l.Hands[i].GetMatches(joker)
		matches2 := l.Hands[j].GetMatches(joker)
		if matches1 < matches2 {
			return true
		}
		if matches1 > matches2 {
			return false
		}

		for c := 0; c < 5; c++ {
			card1 := SameScore(l.Hands[i].Cards[c], joker)
			card2 := SameScore(l.Hands[j].Cards[c], joker)
			if card1 < card2 {
				return true
			}
			if card1 > card2 {
				return false
			}
		}

		return true
	})

	winnings := uint(0)
	for i := 0; i < len(l.Hands); i++ {
		winnings += l.Hands[i].Bid * (uint(i) + 1)
	}
	return winnings
}

func (l *HandList) String() string {
	stringList := ""
	for pos, hand := range l.Hands {
		stringList += hand.String() + fmt.Sprintf(", Winnings: %d\n", hand.Bid*(uint(pos)+1))
	}
	return stringList
}
