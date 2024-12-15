---
title: Cube Conundrum
description: Advent of Code 2023 [Day 2]
layout: default
lang: en
tag: aoc23
prefetch:
  - adventofcode.com
---

You're launched high into the atmosphere! The apex of your trajectory just barely reaches the surface of a large island floating in the sky. You gently land in a fluffy pile of leaves. It's quite cold, but you don't see much snow. An Elf runs over to greet you.

The Elf explains that you've arrived at Snow Island and apologizes for the lack of snow. He'll be happy to explain the situation, but it's a bit of a walk, so you have some time. They don't get many visitors up here; would you like to play a game in the meantime?

As you walk, the Elf shows you a small bag and some cubes which are either red, green, or blue. Each time you play this game, he will hide a secret number of cubes of each color in the bag, and your goal is to figure out information about the number of cubes.

To get information, once a bag has been loaded with cubes, the Elf will reach into the bag, grab a handful of random cubes, show them to you, and then put them back in the bag. He'll do this a few times per game.

You play several games and record the information from each game (your puzzle input). Each game is listed with its ID number (like the 11 in Game 11: ...) followed by a semicolon-separated list of subsets of cubes that were revealed from the bag (like 3 red, 5 green, 4 blue).

For example, the record of a few games might look like this:

```
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
```

In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.

The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration. However, game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once; similarly, game 4 would also have been impossible because the Elf showed you 15 blue cubes at once. If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?

```go
type Configuration struct {
	Red   uint
	Blue  uint
	Green uint
}

func (c *Configuration) Evaluate(line string) bool {
	line = strings.Split(line, ": ")[1]
	lastValue := uint(0)

	for _, set := range strings.Split(line, "; ") {
		for _, tuple := range strings.Split(set, ", ") {
			for _, word := range strings.Split(tuple, " ") {
				switch word {
				case "red":
					if lastValue > c.Red {
						return false
					}
				case "blue":
					if lastValue > c.Blue {
						return false
					}
				case "green":
					if lastValue > c.Green {
						return false
					}
				default:
					if value, err := strconv.Atoi(word); err == nil {
						lastValue = uint(value)
					} else {
						log.Println(err.Error())
					}
				}
			}
		}
	}

	return true
}

func TestConfiguration_Evaluate(t *testing.T) {
	type fields struct {
		Red   uint
		Blue  uint
		Green uint
	}
	tests := []struct {
		name   string
		fields fields
		line   string
		want   bool
	}{
		{
			name:   "Game 1 (low)",
			fields: fields{Red: 12, Blue: 13, Green: 14},
			line:   "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:   true,
		},
		{
			name:   "Game 2 (low)",
			fields: fields{Red: 12, Blue: 13, Green: 14},
			line:   "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:   true,
		},
		{
			name:   "Game 3 (low)",
			fields: fields{Red: 12, Blue: 13, Green: 14},
			line:   "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:   false,
		},
		{
			name:   "Game 4 (low)",
			fields: fields{Red: 12, Blue: 13, Green: 14},
			line:   "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want:   false,
		},
		{
			name:   "Game 5 (low)",
			fields: fields{Red: 12, Blue: 13, Green: 14},
			line:   "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want:   true,
		},
		{
			name:   "Game 1 (high)",
			fields: fields{Red: 100, Blue: 100, Green: 100},
			line:   "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:   true,
		},
		{
			name:   "Game 2 (high)",
			fields: fields{Red: 100, Blue: 100, Green: 100},
			line:   "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:   true,
		},
		{
			name:   "Game 3 (high)",
			fields: fields{Red: 100, Blue: 100, Green: 100},
			line:   "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:   true,
		},
		{
			name:   "Game 4 (high)",
			fields: fields{Red: 100, Blue: 100, Green: 100},
			line:   "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want:   true,
		},
		{
			name:   "Game 5 (high)",
			fields: fields{Red: 100, Blue: 100, Green: 100},
			line:   "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration.Configuration{
				Red:   tt.fields.Red,
				Blue:  tt.fields.Blue,
				Green: tt.fields.Green,
			}
			if got := c.Evaluate(tt.line); got != tt.want {
				t.Errorf("Configuration.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

The Elf says they've stopped producing snow because they aren't getting any water! He isn't sure why the water stopped; however, he can show you how to get to the water source to check it out for yourself. It's just up ahead!

As you continue your walk, the Elf poses a second question: in each game you played, what is the fewest number of cubes of each color that could have been in the bag to make the game possible?

Again consider the example games from earlier:

```
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
```

- In game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes. If any color had even one fewer cube, the game would have been impossible.
- Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
- Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
- Game 4 required at least 14 red, 3 green, and 15 blue cubes.
- Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.

The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together. The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five powers produces the sum 2286.

For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?

```go
func (c *Configuration) Power(line string) uint {
	c.Red = 0
	c.Blue = 0
	c.Green = 0

	line = strings.Split(line, ": ")[1]
	lastValue := uint(0)

	for _, set := range strings.Split(line, "; ") {
		for _, tuple := range strings.Split(set, ", ") {
			for _, word := range strings.Split(tuple, " ") {
				switch word {
				case "red":
					if lastValue > c.Red {
						c.Red = lastValue
					}
				case "blue":
					if lastValue > c.Blue {
						c.Blue = lastValue
					}
				case "green":
					if lastValue > c.Green {
						c.Green = lastValue
					}
				default:
					if value, err := strconv.Atoi(word); err == nil {
						lastValue = uint(value)
					} else {
						log.Println(err.Error())
					}
				}
			}
		}

	}

	return c.Red * c.Blue * c.Green
}

func TestConfiguration_Power(t *testing.T) {
	tests := []struct {
		name string
		line string
		want uint
	}{
		{
			name: "Game 1",
			line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: 48,
		},
		{
			name: "Game 2",
			line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: 12,
		},
		{
			name: "Game 3",
			line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want: 1560,
		},
		{
			name: "Game 4",
			line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: 630,
		},
		{
			name: "Game 5",
			line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: 36,
		},
	}
	for _, tt := range tests {
		c := &configuration.Configuration{}
		t.Run(tt.name, func(t *testing.T) {
			if got := c.Power(tt.line); got != tt.want {
				t.Errorf("Configuration.Power() = %v, want %v", got, tt.want)
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

	var (
		sum1   uint
		sum2   uint
		riddle uint
		c1     = &configuration.Configuration{
			Red:   12,
			Green: 13,
			Blue:  14,
		}
		c2 = &configuration.Configuration{}
	)

	for scanner.Scan() {
		line := scanner.Text()
		riddle++
		if c1.Evaluate(line) {
			sum1 += riddle
		}
		sum2 += c2.Power(line)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}
	log.Printf("sum: %d, %d\n", sum1, sum2)
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2023/day/2)

<details>
	<summary>Click to show the input</summary>
	<pre>
Game 1: 7 green, 14 red, 5 blue; 8 red, 4 green; 6 green, 18 red, 9 blue
Game 2: 3 blue, 15 red, 5 green; 1 blue, 14 red, 5 green; 11 red; 4 green, 1 blue, 3 red; 4 green, 1 blue; 10 red, 1 green
Game 3: 11 green, 3 red; 4 green, 15 blue; 14 blue, 2 red, 10 green; 1 red, 3 green, 10 blue
Game 4: 1 green, 6 red, 11 blue; 3 blue, 12 red; 1 green, 14 red, 8 blue; 3 blue, 7 red; 8 blue, 5 red; 7 red, 1 green
Game 5: 14 green, 3 red, 3 blue; 2 red, 1 green, 1 blue; 8 green, 3 blue, 1 red; 15 green, 8 blue, 1 red
Game 6: 4 blue, 8 green, 5 red; 9 green, 10 blue, 7 red; 11 blue, 10 red, 7 green; 8 red, 6 blue, 9 green
Game 7: 5 green, 11 blue, 9 red; 2 green, 6 red, 12 blue; 8 red, 4 blue, 3 green; 7 green, 8 red, 9 blue; 8 green, 5 red
Game 8: 7 red, 12 green; 9 blue, 15 red, 8 green; 3 blue, 11 green, 6 red; 8 blue, 12 red, 5 green
Game 9: 8 blue, 6 red, 7 green; 2 blue, 3 red, 10 green; 10 blue, 6 red, 7 green; 11 red, 7 blue, 5 green; 10 red, 11 green
Game 10: 5 red, 14 green; 2 red, 6 blue, 15 green; 3 red, 4 blue, 7 green; 6 red, 1 green, 4 blue
Game 11: 4 blue, 11 green, 6 red; 12 red, 1 blue, 5 green; 7 red, 1 blue; 11 red, 2 green, 3 blue; 2 blue, 6 red, 7 green
Game 12: 1 green, 8 red, 3 blue; 3 green, 2 red; 2 blue, 5 red, 1 green
Game 13: 2 green; 8 green, 1 blue, 12 red; 1 blue, 14 green, 2 red; 1 blue, 6 red, 6 green; 7 green, 10 red
Game 14: 9 green, 4 red, 1 blue; 5 red, 2 green; 17 green, 1 red; 6 red, 10 green; 4 green, 3 red, 1 blue
Game 15: 7 green, 13 blue, 4 red; 1 blue, 7 green, 9 red; 13 blue, 13 red, 7 green; 8 red, 9 blue; 9 red, 14 blue; 2 green, 7 red, 9 blue
Game 16: 6 green, 18 blue, 6 red; 5 green, 2 blue, 2 red; 6 green, 2 red, 17 blue; 2 red, 2 green, 8 blue; 2 red, 10 blue
Game 17: 17 red, 8 green; 4 blue, 10 green, 3 red; 8 red, 5 green, 3 blue; 12 green, 3 red
Game 18: 6 red, 1 green, 14 blue; 1 red, 10 blue, 1 green; 1 red, 17 blue, 1 green; 5 red, 1 blue; 5 red, 18 blue; 2 red, 3 blue
Game 19: 5 blue, 12 red; 6 blue, 3 red, 6 green; 8 blue, 6 red, 6 green; 8 green, 8 blue, 2 red; 4 green, 6 red, 6 blue; 1 green, 3 red, 13 blue
Game 20: 7 green, 2 blue; 4 blue, 12 red, 2 green; 7 red, 2 green, 6 blue
Game 21: 8 green, 1 red; 1 red, 9 green; 1 red, 6 green, 4 blue; 1 red, 3 green, 5 blue; 2 red, 6 green
Game 22: 11 green, 12 red, 5 blue; 5 blue, 9 red, 11 green; 8 green, 4 red, 5 blue; 7 green, 1 blue, 1 red
Game 23: 11 blue, 9 red, 5 green; 3 green, 3 blue; 11 blue, 9 red, 1 green; 2 red, 7 green; 4 green, 3 blue, 1 red; 5 green, 4 blue
Game 24: 1 green, 4 blue, 9 red; 1 green, 2 blue, 11 red; 1 green, 13 red; 1 green, 2 blue, 3 red
Game 25: 1 red, 7 green, 4 blue; 2 red, 1 green, 3 blue; 10 blue, 1 red; 7 blue, 2 red, 6 green; 7 green, 15 blue, 2 red; 14 green, 13 blue
Game 26: 5 red, 2 blue; 9 red, 2 green, 12 blue; 15 red, 1 green, 5 blue; 1 green, 16 blue, 17 red
Game 27: 2 green, 4 red; 4 red, 1 green; 1 blue, 3 red; 2 red
Game 28: 3 green; 8 green, 9 red; 9 red, 3 blue, 10 green; 16 green, 4 blue, 4 red
Game 29: 3 green, 1 red, 7 blue; 5 blue, 5 green, 2 red; 5 blue, 6 green, 2 red; 2 green, 2 red, 4 blue; 1 green, 3 red, 8 blue
Game 30: 7 red, 3 green, 7 blue; 3 green, 10 red; 5 red, 5 blue, 1 green; 9 blue, 2 green, 7 red; 1 red, 10 blue; 10 blue, 2 red, 4 green
Game 31: 9 green, 1 red; 9 blue, 6 red, 9 green; 17 blue, 4 green, 10 red; 19 blue, 11 green
Game 32: 1 red, 1 blue, 6 green; 10 blue, 4 green; 1 red, 5 blue; 9 blue, 3 green
Game 33: 4 red; 3 red; 2 red, 1 green, 1 blue; 1 green; 1 blue, 1 red
Game 34: 2 green, 9 blue, 1 red; 5 blue, 7 green, 1 red; 2 green, 1 red, 16 blue; 1 blue, 5 green, 6 red
Game 35: 11 red, 10 blue; 2 blue, 12 green, 12 red; 3 green, 6 red, 6 blue; 14 blue, 10 green, 1 red
Game 36: 2 blue, 3 red, 15 green; 2 blue, 6 green, 2 red; 14 blue, 4 red, 7 green; 13 blue, 12 green, 2 red
Game 37: 6 green, 14 blue, 7 red; 7 blue, 2 red, 6 green; 1 blue, 2 green, 6 red
Game 38: 2 green, 15 red, 2 blue; 14 red, 1 blue; 14 red, 2 green, 12 blue
Game 39: 5 green, 1 blue, 10 red; 4 red, 3 blue, 7 green; 2 red, 2 green, 4 blue; 10 blue, 5 green
Game 40: 7 red, 10 green, 2 blue; 7 green, 3 red, 2 blue; 10 red, 9 blue, 7 green; 3 green, 5 blue, 10 red
Game 41: 5 blue, 2 green, 11 red; 2 green, 18 red, 3 blue; 8 green, 10 red, 1 blue; 16 red, 13 green; 17 green, 2 blue, 17 red; 1 green, 1 blue, 9 red
Game 42: 5 red, 2 green, 1 blue; 6 red, 2 blue; 3 red, 1 blue; 9 red, 5 blue; 1 green, 8 red, 1 blue
Game 43: 1 red, 2 green; 12 red, 4 green, 5 blue; 4 blue, 9 red; 4 green, 10 red, 2 blue
Game 44: 2 blue, 9 green, 3 red; 6 red, 4 blue, 4 green; 3 red, 4 blue; 5 red, 2 green, 1 blue; 4 blue, 1 green; 8 green, 1 red, 4 blue
Game 45: 7 blue, 1 red; 2 red, 4 green, 9 blue; 3 red, 15 blue; 4 red, 4 green, 12 blue; 1 red, 18 blue
Game 46: 4 red, 14 blue, 11 green; 5 blue, 6 red, 17 green; 10 red, 8 green, 17 blue; 7 red, 10 blue, 19 green
Game 47: 7 blue, 3 red; 7 blue, 1 green, 2 red; 2 red, 6 blue; 1 green, 9 blue, 2 red; 3 red; 2 green, 1 blue
Game 48: 12 red, 6 blue, 6 green; 9 green, 19 red, 1 blue; 2 blue, 12 green, 8 red
Game 49: 1 green, 11 red, 11 blue; 10 red, 10 blue, 11 green; 4 red, 19 green, 6 blue; 11 blue, 19 green, 13 red; 9 green, 9 blue
Game 50: 1 blue, 12 green, 4 red; 1 blue, 18 green, 1 red; 1 blue, 12 green, 3 red; 1 blue, 4 green
Game 51: 10 red, 5 blue, 1 green; 10 red, 4 blue; 6 red, 8 blue
Game 52: 1 blue; 5 green, 9 red; 2 blue, 1 green, 11 red; 2 blue, 13 red, 5 green; 6 green, 1 blue, 9 red
Game 53: 8 blue, 15 red; 2 green, 4 red, 12 blue; 6 blue, 1 green, 15 red; 20 red, 12 blue; 6 red, 1 green, 2 blue
Game 54: 5 red, 16 blue; 5 green, 3 red, 17 blue; 5 red, 3 blue, 5 green; 4 green, 6 blue, 9 red; 2 blue, 6 green, 2 red
Game 55: 1 blue, 1 red; 1 green, 1 red, 3 blue; 4 blue, 1 green, 1 red; 5 blue; 2 blue
Game 56: 4 red, 4 blue; 7 blue, 11 red; 1 red, 2 green, 9 blue; 4 blue, 16 red, 2 green; 1 red; 2 green, 5 blue, 1 red
Game 57: 1 green, 8 blue; 1 red; 10 blue, 5 green; 3 blue, 4 green; 11 blue, 1 red; 4 blue, 3 green, 1 red
Game 58: 8 green, 5 blue; 9 blue, 8 red, 5 green; 6 red, 6 blue, 9 green; 1 green, 5 blue, 2 red; 3 red, 3 green, 2 blue; 2 green, 1 red, 1 blue
Game 59: 15 red, 4 blue, 8 green; 12 red, 6 green; 3 red
Game 60: 14 blue, 11 red; 12 blue, 6 red; 11 blue, 6 red; 5 red, 13 blue; 15 blue; 1 green, 1 blue, 16 red
Game 61: 5 red, 1 green; 4 red, 9 green; 1 blue, 6 green, 14 red
Game 62: 19 red, 1 green; 1 blue, 3 red; 15 red, 1 blue; 1 blue, 3 red; 5 red, 1 green, 1 blue
Game 63: 1 red, 3 green, 10 blue; 2 green, 1 red, 14 blue; 1 green, 5 blue, 1 red; 6 blue, 4 green, 1 red
Game 64: 5 red, 2 green; 5 green, 2 red, 2 blue; 3 red, 3 blue, 1 green; 3 blue, 3 green, 3 red; 1 green, 3 red
Game 65: 13 red, 2 green, 3 blue; 1 red, 2 blue, 1 green; 1 blue; 2 green, 1 red
Game 66: 7 red, 12 blue, 6 green; 2 red, 5 green, 11 blue; 3 green, 2 blue, 2 red; 9 blue, 1 red, 2 green
Game 67: 4 red, 3 green, 7 blue; 8 blue, 3 red; 2 red; 9 blue, 5 red, 2 green
Game 68: 12 blue; 10 green, 5 blue; 8 blue; 9 blue, 7 red, 18 green; 5 red, 12 blue, 8 green; 8 green, 13 red, 10 blue
Game 69: 1 green, 1 red; 2 red, 1 green, 3 blue; 1 red, 1 green, 4 blue; 1 green, 8 red
Game 70: 12 green, 1 blue, 4 red; 8 green, 1 red; 1 blue, 8 green; 2 green, 3 red; 5 green, 4 red; 2 blue, 12 green, 1 red
Game 71: 10 blue, 4 red, 14 green; 6 green, 7 red, 8 blue; 1 red, 1 blue, 13 green; 10 red, 6 blue, 3 green; 8 blue, 7 green, 4 red
Game 72: 1 green; 1 blue, 12 green, 14 red; 3 blue, 7 green, 8 red; 12 red, 18 green; 13 green, 11 red, 1 blue; 2 blue, 6 green, 6 red
Game 73: 17 red, 3 green, 15 blue; 15 blue, 2 red; 15 red, 7 blue, 4 green; 9 blue, 1 green, 18 red
Game 74: 10 red, 2 blue; 1 blue, 7 red; 5 blue, 2 green, 2 red; 3 blue, 15 red, 3 green; 4 blue, 3 green, 13 red
Game 75: 6 blue, 10 red; 2 green, 2 blue, 10 red; 10 green, 1 blue, 10 red; 4 blue, 6 red, 11 green
Game 76: 10 blue, 1 red, 2 green; 6 blue, 2 green, 10 red; 3 red, 15 green, 1 blue
Game 77: 5 green, 1 red; 2 blue, 1 green; 13 green, 2 red, 5 blue; 12 green, 1 blue, 2 red; 3 blue, 2 green, 2 red
Game 78: 1 green, 16 red; 6 red, 1 blue, 1 green; 13 red; 12 red, 3 green; 1 blue, 7 red
Game 79: 3 green, 7 blue; 1 red, 8 blue, 5 green; 1 red, 6 green, 7 blue; 11 green, 1 red, 7 blue; 1 blue
Game 80: 3 green, 13 red, 8 blue; 17 red, 9 blue; 7 blue, 1 green, 2 red; 8 red, 6 blue, 3 green; 1 red, 2 blue; 2 green, 4 blue, 10 red
Game 81: 3 red, 1 green, 7 blue; 2 green, 2 blue, 3 red; 3 red, 1 blue, 7 green; 6 green, 12 blue
Game 82: 11 red, 3 green, 2 blue; 3 red, 1 green, 1 blue; 16 red, 1 green
Game 83: 8 green, 3 blue, 2 red; 1 blue, 13 green, 6 red; 4 blue, 5 red, 1 green; 12 green, 4 red, 12 blue; 17 green, 7 blue, 3 red
Game 84: 2 blue, 13 red, 5 green; 3 green, 3 blue, 19 red; 2 red, 11 green, 5 blue; 3 green, 3 blue, 15 red; 7 green, 4 blue, 11 red; 1 red, 10 green
Game 85: 1 red, 3 blue, 4 green; 2 red, 11 green, 2 blue; 2 blue, 7 green, 1 red
Game 86: 3 blue, 4 green, 8 red; 4 green, 2 red; 9 red, 4 blue, 1 green; 18 red, 1 blue
Game 87: 3 red, 14 blue, 1 green; 10 blue, 1 green; 1 green, 4 red, 14 blue; 8 blue, 7 green, 4 red; 2 green, 7 red, 7 blue; 2 green, 10 blue
Game 88: 12 green, 6 red; 6 red, 3 blue, 2 green; 4 red, 4 blue, 9 green; 3 red, 4 green, 8 blue; 1 blue, 3 red
Game 89: 3 green, 3 red, 2 blue; 3 red, 2 green, 1 blue; 6 green, 4 blue, 12 red; 13 red, 14 blue, 1 green; 5 red; 10 red, 8 blue, 7 green
Game 90: 7 green, 10 blue; 6 green, 1 red, 2 blue; 6 blue; 5 green, 9 blue, 1 red; 10 blue, 1 red, 6 green
Game 91: 6 red, 2 blue; 3 blue, 3 red, 1 green; 19 blue, 7 red
Game 92: 9 green, 3 blue; 1 red, 5 green; 13 green, 3 blue, 2 red; 1 red, 3 blue, 7 green
Game 93: 11 red, 3 green, 11 blue; 7 green, 3 red, 10 blue; 11 green, 4 blue, 8 red; 14 green, 8 blue
Game 94: 7 blue; 1 green, 11 blue, 2 red; 1 green, 1 red, 19 blue; 7 green, 2 red, 10 blue
Game 95: 15 blue, 1 red, 9 green; 5 green, 1 red, 4 blue; 6 green, 17 blue; 9 blue, 11 green; 10 blue, 9 green; 9 blue, 7 green
Game 96: 7 red, 13 blue; 6 blue, 15 red, 3 green; 1 green, 1 red, 1 blue; 9 red, 2 green, 8 blue; 5 green, 8 red, 1 blue; 6 blue, 3 green, 13 red
Game 97: 19 blue, 10 red, 4 green; 8 red, 17 blue; 8 blue
Game 98: 2 blue, 2 red, 4 green; 5 green, 3 blue, 2 red; 5 green, 15 blue; 15 blue, 5 green, 1 red
Game 99: 1 blue, 2 green, 8 red; 1 blue, 7 red, 1 green; 11 red, 2 green; 1 red, 1 blue
Game 100: 8 green; 2 red, 20 green; 12 green, 1 red, 1 blue; 4 red, 1 blue; 1 blue, 6 red
	</pre>
</details>
