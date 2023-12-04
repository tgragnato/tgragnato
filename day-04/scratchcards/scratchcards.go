package scratchcards

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func NewScratchcard(line string) *Scratchcard {
	line = strings.Split(line, ": ")[1]
	s := &Scratchcard{
		Winning: []uint{},
		Having:  []uint{},
	}

	piped := strings.Split(line, "| ")
	for index := 0; index < len(piped); index++ {
		for _, number := range strings.Split(piped[index], " ") {
			if intNumber, err := strconv.Atoi(number); err == nil {
				if index == 0 {
					s.Winning = append(s.Winning, uint(intNumber))
				} else if index == 1 {
					s.Having = append(s.Having, uint(intNumber))
				} else {
					log.Println("parsing error in NewScratchcard")
				}
			}
		}
	}

	return s
}

type Scratchcard struct {
	Winning []uint
	Having  []uint
}

func (s *Scratchcard) Have() uint {
	var sum uint = 0
	for _, have := range s.Having {
		for _, win := range s.Winning {
			if have == win {
				sum++
			}
		}
	}
	return sum
}

func (s *Scratchcard) Points() uint {
	sum := s.Have()
	if sum == 0 {
		return 0
	}
	return uint(math.Pow(2, float64(sum-1)))
}

type ScratchcardList struct {
	List []Scratchcard
}

func (l *ScratchcardList) Populate(lines []string) {
	for _, line := range lines {
		l.List = append(l.List, *NewScratchcard(line))
	}
}

func (l *ScratchcardList) Points() uint {
	var sum uint = 0
	for _, s := range l.List {
		sum += s.Points()
	}
	return sum
}

func (l *ScratchcardList) Matches() uint {
	partial := map[int]uint{}
	for index := 0; index < len(l.List); index++ {
		partial[index]++
		if have := l.List[index].Have(); have != 0 {
			for riddle := 1; riddle <= int(have); riddle++ {
				partial[index+riddle] += partial[index]
			}
		}
	}

	sum := uint(0)
	for index := 0; index < len(l.List); index++ {
		sum += partial[index]
	}
	return sum
}
