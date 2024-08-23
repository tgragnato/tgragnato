---
title: Camel Cards
description: Advent of Code 2023 [Day 7]
layout: default
lang: en
---

Your all-expenses-paid trip turns out to be a one-way, five-minute ride in an airship. (At least it's a cool airship!) It drops you off at the edge of a vast desert and descends back to Island Island.

"Did you bring the parts?"

You turn around to see an Elf completely covered in white clothing, wearing goggles, and riding a large camel.

"Did you bring the parts?" she asks again, louder this time. You aren't sure what parts she's looking for; you're here to figure out why the sand stopped.

"The parts! For the sand, yes! Come with me; I will show you." She beckons you onto the camel.

After riding a bit across the sands of Desert Island, you can see what look like very large rocks covering half of the horizon. The Elf explains that the rocks are all along the part of Desert Island that is directly above Island Island, making it hard to even get there. Normally, they use big machines to move the rocks and filter the sand, but the machines have broken down because Desert Island recently stopped receiving the parts they need to fix the machines.

You've already assumed it'll be your job to figure out why the parts stopped when she asks if you can help. You agree automatically.

Because the journey will take a few days, she offers to teach you the game of Camel Cards. Camel Cards is sort of similar to poker except it's designed to be easier to play while riding a camel.

In Camel Cards, you get a list of hands, and your goal is to order them based on the strength of each hand. A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2. The relative strength of each card follows this order, where A is the highest and 2 is the lowest.

Every hand is exactly one type. From strongest to weakest, they are:

- Five of a kind, where all five cards have the same label: AAAAA
- Four of a kind, where four cards have the same label and one card has a different label: AA8AA
- Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
- Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
- Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
- One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
- High card, where all cards' labels are distinct: 23456

Hands are primarily ordered based on type; for example, every full house is stronger than any three of a kind.

If two hands have the same type, a second ordering rule takes effect. Start by comparing the first card in each hand. If these cards are different, the hand with the stronger first card is considered stronger. If the first card in each hand have the same label, however, then move on to considering the second card in each hand. If they differ, the hand with the higher second card wins; otherwise, continue with the third card in each hand, then the fourth, then the fifth.

So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because its first card is stronger. Similarly, 77888 and 77788 are both a full house, but 77888 is stronger because its third card is stronger (and both hands have the same first and second card).

To play Camel Cards, you are given a list of hands and their corresponding bid (your puzzle input). For example:

```
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
```

This example shows five hands; each hand is followed by its bid amount. Each hand wins an amount equal to its bid multiplied by its rank, where the weakest hand gets rank 1, the second-weakest hand gets rank 2, and so on up to the strongest hand. Because there are five hands in this example, the strongest hand will have rank 5 and its bid will be multiplied by 5.

So, the first step is to put the hands in order of strength:

- 32T3K is the only one pair and the other hands are all a stronger type, so it gets rank 1.
- KK677 and KTJJT are both two pair. Their first cards both have the same label, but the second card of KK677 is stronger (K vs T), so KTJJT gets rank 2 and KK677 gets rank 3.
- T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first card, so it gets rank 5 and T55J5 gets rank 4.

Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this example are 6440.

Find the rank of every hand in your set. What are the total winnings?

To make things a little more interesting, the Elf introduces one additional rule. Now, J cards are jokers - wildcards that can act like whatever card would make the hand the strongest type possible.

To balance this, J cards are now the weakest individual cards, weaker even than 2. The other cards stay in the same order: A, K, Q, T, 9, 8, 7, 6, 5, 4, 3, 2, J.

J cards can pretend to be whatever card is best for the purpose of determining hand type; for example, QJJQ2 is now considered four of a kind. However, for the purpose of breaking ties between two hands of the same type, J is always treated as J, not the card it's pretending to be: JKKK2 is weaker than QQQQ2 because J is weaker than Q.

Now, the above example goes very differently:

```
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
```

- 32T3K is still the only one pair; it doesn't contain any jokers, so its strength doesn't increase.
- KK677 is now the only two pair, making it the second-weakest hand.
- T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.

With the new joker rule, the total winnings in this example are 5905.

Using the new joker rule, find the rank of every hand in your set. What are the new total winnings?

```go
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

func TestHandList_GetWinnings(t *testing.T) {
	hl := &camelcards.HandList{
		Hands: []camelcards.Hand{
			{
				Cards: [5]rune{'3', '2', 'T', '3', 'K'},
				Bid:   765,
			},
			{
				Cards: [5]rune{'T', '5', '5', 'J', '5'},
				Bid:   684,
			},
			{
				Cards: [5]rune{'K', 'K', '6', '7', '7'},
				Bid:   28,
			},
			{
				Cards: [5]rune{'K', 'T', 'J', 'J', 'T'},
				Bid:   220,
			},
			{
				Cards: [5]rune{'Q', 'Q', 'Q', 'J', 'A'},
				Bid:   483,
			},
		},
	}

	t.Run("Example Test 2", func(t *testing.T) {
		l := hl
		if got := l.GetWinnings(true); got != 5905 {
			t.Errorf("HandList.GetWinnings() = %v, want 5905", got)
		}
	})

	t.Run("Example Test 1", func(t *testing.T) {
		l := hl
		if got := l.GetWinnings(false); got != 6440 {
			t.Errorf("HandList.GetWinnings() = %v, want 6440", got)
		}
	})
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	handList := camelcards.HandList{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		cards := []rune(line[0])

		bid, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalln(err.Error())
		}

		handList.Hands = append(handList.Hands, camelcards.Hand{
			Cards: [5]rune(cards),
			Bid:   uint(bid),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	points1 := handList.GetWinnings(false)
	dump1 := handList.String()
	points2 := handList.GetWinnings(true)
	dump2 := handList.String()
	log.Printf("\n---\n%s---\n%s---\nsum: %d, %d\n", dump1, dump2, points1, points2)
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/)

- [input.txt](/documents/2023-12-07-input.txt)
- [Challenge](https://adventofcode.com/2023/day/7)
