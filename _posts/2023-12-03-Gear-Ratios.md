---
title: Gear Ratios
description: Advent of Code 2023 [Day 3]
layout: default
lang: en
prefetch:
  - adventofcode.com
---

You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

```
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
```

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?

```go
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

var testSchema = [][]rune{
	{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
	{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
	{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
	{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
	{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
	{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
}

func TestSchematic_IsAdjacent(t *testing.T) {
	tests := []struct {
		name   string
		Schema [][]rune
		start  int
		end    int
		line   int
		want   bool
	}{
		{
			name:   "467",
			Schema: testSchema,
			start:  0,
			end:    2,
			line:   0,
			want:   true,
		},
		{
			name:   "114",
			Schema: testSchema,
			start:  5,
			end:    7,
			line:   0,
			want:   false,
		},
		{
			name:   "35",
			Schema: testSchema,
			start:  2,
			end:    3,
			line:   2,
			want:   true,
		},
		{
			name:   "633",
			Schema: testSchema,
			start:  6,
			end:    8,
			line:   2,
			want:   true,
		},
		{
			name:   "617",
			Schema: testSchema,
			start:  0,
			end:    2,
			line:   4,
			want:   true,
		},
		{
			name:   "58",
			Schema: testSchema,
			start:  7,
			end:    8,
			line:   5,
			want:   false,
		},
		{
			name:   "592",
			Schema: testSchema,
			start:  2,
			end:    4,
			line:   6,
			want:   true,
		},
		{
			name:   "755",
			Schema: testSchema,
			start:  6,
			end:    8,
			line:   7,
			want:   true,
		},
		{
			name:   "664",
			Schema: testSchema,
			start:  1,
			end:    3,
			line:   9,
			want:   true,
		},
		{
			name:   "598",
			Schema: testSchema,
			start:  5,
			end:    7,
			line:   9,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &engine.Schematic{
				Schema: tt.Schema,
			}
			if got := s.IsAdjacent(engine.PartNumber{
				Start: tt.start,
				End:   tt.end,
				Row:   tt.line,
			}); got != tt.want {
				t.Errorf("Schematic.isAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchematic_ExtractPartNumber(t *testing.T) {
	tests := []struct {
		name   string
		Schema [][]rune
		start  int
		end    int
		row    int
		want   uint
	}{
		{
			name:   "467",
			Schema: testSchema,
			start:  0,
			end:    2,
			row:    0,
			want:   467,
		},
		{
			name:   "114",
			Schema: testSchema,
			start:  5,
			end:    7,
			row:    0,
			want:   114,
		},
		{
			name:   "35",
			Schema: testSchema,
			start:  2,
			end:    3,
			row:    2,
			want:   35,
		},
		{
			name:   "633",
			Schema: testSchema,
			start:  6,
			end:    8,
			row:    2,
			want:   633,
		},
		{
			name:   "617",
			Schema: testSchema,
			start:  0,
			end:    2,
			row:    4,
			want:   617,
		},
		{
			name:   "58",
			Schema: testSchema,
			start:  7,
			end:    8,
			row:    5,
			want:   58,
		},
		{
			name:   "592",
			Schema: testSchema,
			start:  2,
			end:    4,
			row:    6,
			want:   592,
		},
		{
			name:   "755",
			Schema: testSchema,
			start:  6,
			end:    8,
			row:    7,
			want:   755,
		},
		{
			name:   "664",
			Schema: testSchema,
			start:  1,
			end:    3,
			row:    9,
			want:   664,
		},
		{
			name:   "598",
			Schema: testSchema,
			start:  5,
			end:    7,
			row:    9,
			want:   598,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &engine.Schematic{
				Schema: tt.Schema,
			}
			if got := s.ExtractPartNumber(engine.PartNumber{
				Start: tt.start,
				End:   tt.end,
				Row:   tt.row,
			}); got != tt.want {
				t.Errorf("Schematic.ExtractPartNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchematic_SumPartNumber(t *testing.T) {
	t.Run("part1-test", func(t *testing.T) {
		s := &engine.Schematic{
			Schema: testSchema,
		}
		if gotSum := s.SumPartNumber(); gotSum != 4361 {
			t.Errorf("Schematic.SumPartNumber() = %v, want 4361", gotSum)
		}
	})

}
```

The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

```
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
```

In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?

```go
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

func TestSchematic_SumGearRatio(t *testing.T) {
	t.Run("part2-test", func(t *testing.T) {
		s := &engine.Schematic{
			Schema: testSchema,
		}
		if got := s.SumGearRatio(); got != 467835 {
			t.Errorf("Schematic.SumGearRatio() = %v, want 467835", got)
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

	var (
		engineSchematic *(engine.Schematic) = &engine.Schematic{
			Schema: [][]rune{},
		}
	)

	for scanner.Scan() {
		runes := []rune(scanner.Text())
		runes = append(runes, '.')
		engineSchematic.Schema = append(engineSchematic.Schema, runes)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}
	log.Printf(
		"sum: %d, %d\n",
		engineSchematic.SumPartNumber(),
		engineSchematic.SumGearRatio(),
	)
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/)

- [input.txt](/documents/2023-12-03-input.txt)
- [Challenge](https://adventofcode.com/2023/day/3)
