---
title: Trebuchet?!
description: Advent of Code 2023 [Day 1]
layout: default
lang: en
---

Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

```
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
```

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

```go
func ExtractValues1(line string) uint {
	var (
		first  uint = 0
		last   uint = 0
		ffound bool = false
	)
	for _, char := range line {
		value := int(char - '0')
		if value < 0 || value >= 10 {
			continue
		}
		if !ffound {
			first = uint(value)
			ffound = true
		}
		last = uint(value)
	}

	return first*10 + last
}

func TestExtractValues1(t *testing.T) {
	tests := []struct {
		name string
		line string
		want uint
	}{
		{
			name: "Line 1",
			line: "1abc2",
			want: 12,
		},
		{
			name: "Line 2",
			line: "pqr3stu8vwx",
			want: 38,
		},
		{
			name: "Line 3",
			line: "a1b2c3d4e5f",
			want: 15,
		},
		{
			name: "Line 4",
			line: "treb7uchet",
			want: 77,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extraction.ExtractValues1(tt.line); got != tt.want {
				t.Errorf("ExtractValues1() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

```
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
```

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?

```go
func ExtractValues2(line string) uint {
	var (
		first  uint            = 0
		last   uint            = 0
		ffound bool            = false
		digits map[string]uint = map[string]uint{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}
		substring []rune = []rune{}
	)

	for _, char := range line {
		if value := int(char - '0'); value >= 0 && value < 10 {
			if !ffound {
				first = uint(value)
				ffound = true
			}
			last = uint(value)
		} else {
			substring = append(substring, char)
			for key, value := range digits {
				x := []rune(key)
				if len(substring) < len(x) {
					continue
				}

				y := substring[len(substring)-len(x):]
				if !reflect.DeepEqual(x, y) {
					continue
				}

				if !ffound {
					first = uint(value)
					ffound = true
				}
				last = uint(value)
			}
		}
	}

	return first*10 + last
}

func TestExtractValues2(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		line string
		want uint
	}{
		{
			name: "Line 1",
			line: "two1nine",
			want: 29,
		},
		{
			name: "Line 2",
			line: "eightwothree",
			want: 83,
		},
		{
			name: "Line 3",
			line: "abcone2threexyz",
			want: 13,
		},
		{
			name: "Line 4",
			line: "xtwone3four",
			want: 24,
		},
		{
			name: "Line 5",
			line: "4nineeightseven2",
			want: 42,
		},
		{
			name: "Line 6",
			line: "zoneight234",
			want: 14,
		},
		{
			name: "Line 7",
			line: "7pqrstsixteen",
			want: 76,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extraction.ExtractValues2(tt.line); got != tt.want {
				t.Errorf("ExtractValues2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum1 := uint(0)
	sum2 := uint(0)
	for scanner.Scan() {
		line := scanner.Text()
		sum1 += extraction.ExtractValues1(line)
		sum2 += extraction.ExtractValues2(line)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}
	log.Printf("sum: %d, %d\n", sum1, sum2)
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/)

- [input.txt](/documents/2023-12-01-input.txt)
- [Challenge](https://adventofcode.com/2023/day/1)
