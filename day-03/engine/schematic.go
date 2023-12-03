package engine

import (
	"strconv"
)

type PartNumber struct {
	Start int
	End   int
	Row   int
}

func (p *PartNumber) isAdjacent(row int, column int) bool {
	if p.Row < row-1 || p.Row > row+1 {
		return false
	}

	for i := p.Start; i <= p.End; i++ {
		if i >= column-1 && i <= column+1 {
			return true
		}
	}

	return false
}

type Schematic struct {
	Schema      [][]rune
	partNumbers []PartNumber
}

func (s *Schematic) IsAdjacent(pn PartNumber) bool {
	for row := pn.Row - 1; row <= pn.Row+1; row++ {
		if row < 0 || row >= len(s.Schema) {
			continue
		}

		for column := pn.Start - 1; column <= pn.End+1; column++ {
			if column < 0 || column >= len(s.Schema[row]) {
				continue
			}

			switch s.Schema[row][column] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
				continue
			default:
				return true
			}
		}
	}

	return false
}

func (s *Schematic) ExtractPartNumber(pn PartNumber) uint {
	extracted := []rune{}
	for column := pn.Start; column <= pn.End; column++ {
		extracted = append(extracted, s.Schema[pn.Row][column])
	}
	number, err := strconv.Atoi(string(extracted))
	if err != nil {
		return 0
	}
	return uint(number)
}

func (s *Schematic) SumPartNumber() uint {
	var (
		sum   uint = 0
		start int  = 0
		end   int  = 0
		found bool = false
	)

	for row := 0; row < len(s.Schema); row++ {
		for column := 0; column < len(s.Schema[row]); column++ {
			if num := s.Schema[row][column] - '0'; num >= 0 && num <= 9 {
				if !found {
					start = column
					found = true
				}
				end = column
			} else if found {
				partNumber := PartNumber{
					Start: start,
					End:   end,
					Row:   row,
				}
				if s.IsAdjacent(partNumber) {
					sum += s.ExtractPartNumber(partNumber)
				}
				found = false
			}
		}

		start = 0
		end = 0
		found = false
	}

	return sum
}

func (s *Schematic) ExtractGearRatio(gearRow int, gearColumn int) uint {
	gears := []PartNumber{}

	for _, pn := range s.partNumbers {
		if pn.isAdjacent(gearRow, gearColumn) {
			gears = append(gears, pn)
		}
	}

	if len(gears) == 2 {
		return s.ExtractPartNumber(gears[0]) * s.ExtractPartNumber(gears[1])
	}

	return 0
}

func (s *Schematic) SumGearRatio() uint {
	var (
		sum   uint = 0
		start int  = 0
		end   int  = 0
		found bool = false
	)

	for row := 0; row < len(s.Schema); row++ {
		for column := 0; column < len(s.Schema[row]); column++ {
			if num := s.Schema[row][column] - '0'; num >= 0 && num <= 9 {
				if !found {
					start = column
					found = true
				}
				end = column
			} else if found {
				partNumber := PartNumber{
					Start: start,
					End:   end,
					Row:   row,
				}
				s.partNumbers = append(s.partNumbers, partNumber)
				found = false
			}
		}

		start = 0
		end = 0
		found = false
	}

	for row := 0; row < len(s.Schema); row++ {
		for column := 0; column < len(s.Schema[row]); column++ {
			if s.Schema[row][column] == '*' {
				sum += s.ExtractGearRatio(row, column)
			}
		}
	}

	return sum
}
