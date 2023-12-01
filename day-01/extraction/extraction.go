package extraction

import "reflect"

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
