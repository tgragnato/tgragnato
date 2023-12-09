package camelcards_test

import (
	"testing"

	"github.com/tgragnato/aoc23/day-07/camelcards"
)

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
