package scratchcards_test

import (
	"testing"

	"github.com/tgragnato/aoc23/day-04/scratchcards"
)

func TestScratchcard_Points(t *testing.T) {
	tests := []struct {
		name    string
		Winning []uint
		Having  []uint
		want    uint
	}{
		{
			name:    "Card 1",
			Winning: []uint{41, 48, 83, 86, 17},
			Having:  []uint{83, 86, 6, 31, 17, 9, 48, 53},
			want:    8,
		},
		{
			name:    "Card 2",
			Winning: []uint{13, 32, 20, 16, 61},
			Having:  []uint{61, 30, 68, 82, 17, 32, 24, 19},
			want:    2,
		},
		{
			name:    "Card 3",
			Winning: []uint{1, 21, 53, 59, 44},
			Having:  []uint{69, 82, 63, 72, 16, 21, 14, 1},
			want:    2,
		},
		{
			name:    "Card 4",
			Winning: []uint{41, 92, 73, 84, 69},
			Having:  []uint{59, 84, 76, 51, 58, 5, 54, 83},
			want:    1,
		},
		{
			name:    "Card 5",
			Winning: []uint{87, 83, 26, 28, 32},
			Having:  []uint{88, 30, 70, 12, 93, 22, 82, 36},
			want:    0,
		},
		{
			name:    "Card 6",
			Winning: []uint{31, 18, 13, 56, 72},
			Having:  []uint{74, 77, 10, 23, 35, 67, 36, 11},
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scratchcards.Scratchcard{
				Winning: tt.Winning,
				Having:  tt.Having,
			}
			if got := s.Points(); got != tt.want {
				t.Errorf("Scratchcard.Points() = %v, want %v", got, tt.want)
			}
		})
	}
}

var test = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

func TestScratchcardList_Points(t *testing.T) {
	t.Run("Example Test", func(t *testing.T) {
		l := &scratchcards.ScratchcardList{
			List: []scratchcards.Scratchcard{},
		}
		l.Populate(test)
		if got := l.Points(); got != 13 {
			t.Errorf("ScratchcardList.Points() = %v, want 13", got)
		}
	})
}

func TestScratchcardList_Matches(t *testing.T) {
	t.Run("Example Test", func(t *testing.T) {
		l := &scratchcards.ScratchcardList{
			List: []scratchcards.Scratchcard{},
		}
		l.Populate(test)
		if got := l.Matches(); got != 30 {
			t.Errorf("ScratchcardList.Matches() = %v, want 30", got)
		}
	})
}
