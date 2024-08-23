---
title: Hot Springs
description: Advent of Code 2023 [Day 12]
layout: default
lang: en
prefetch:
  - adventofcode.com
---

You finally reach the hot springs! You can see steam rising from secluded areas attached to the primary, ornate building.

As you turn to enter, the researcher stops you. "Wait - I thought you were looking for the hot springs, weren't you?" You indicate that this definitely looks like hot springs to you.

"Oh, sorry, common mistake! This is actually the onsen! The hot springs are next door."

You look in the direction the researcher is pointing and suddenly notice the massive metal helixes towering overhead. "This way!"

It only takes you a few more steps to reach the main gate of the massive fenced-off area containing the springs. You go through the gate and into a small administrative building.

"Hello! What brings you to the hot springs today? Sorry they're not very hot right now; we're having a lava shortage at the moment." You ask about the missing machine parts for Desert Island.

"Oh, all of Gear Island is currently offline! Nothing is being manufactured at the moment, not until we get more lava to heat our forges. And our springs. The springs aren't very springy unless they're hot!"

"Say, could you go up and see why the lava stopped flowing? The springs are too cold for normal operation, but we should be able to find one springy enough to launch you up there!"

There's just one problem - many of the springs have fallen into disrepair, so they're not actually sure which springs would even be safe to use! Worse yet, their condition records of which springs are damaged (your puzzle input) are also damaged! You'll need to help them repair the damaged records.

In the giant field just outside, the springs are arranged into rows. For each row, the condition records show every spring and whether it is operational (.) or damaged (#). This is the part of the condition records that is itself damaged; for some springs, it is simply unknown (?) whether the spring is operational or damaged.

However, the engineer that produced the condition records also duplicated some of this information in a different format! After the list of springs for a given row, the size of each contiguous group of damaged springs is listed in the order those groups appear in the row. This list always accounts for every damaged spring, and each number is the entire size of its contiguous group (that is, groups are always separated by at least one operational spring: #### would always be 4, never 2,2).

So, condition records with no unknown spring conditions might look like this:

```
#.#.### 1,1,3
.#...#....###. 1,1,3
.#.###.#.###### 1,3,1,6
####.#...#... 4,1,1
#....######..#####. 1,6,5
.###.##....# 3,2,1
```

However, the condition records are partially damaged; some of the springs' conditions are actually unknown (?). For example:

```
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
```

Equipped with this information, it is your job to figure out how many different arrangements of operational and broken springs fit the given criteria in each row.

In the first line (???.### 1,1,3), there is exactly one way separate groups of one, one, and three broken springs (in that order) can appear in that row: the first three unknown springs must be broken, then operational, then broken (#.#), making the whole row #.#.###.

The second line is more interesting: .??..??...?##. 1,1,3 could be a total of four different arrangements. The last ? must always be broken (to satisfy the final contiguous group of three broken springs), and each ?? must hide exactly one of the two broken springs. (Neither ?? could be both broken springs or they would form a single contiguous group of two; if that were true, the numbers afterward would have been 2,3 instead.) Since each ?? can either be #. or .#, there are four possible arrangements of springs.

The last line is actually consistent with ten different arrangements! Because the first number is 3, the first and second ? must both be . (if either were #, the first number would have to be 4 or higher). However, the remaining run of unknown spring conditions have many different ways they could hold groups of two and one broken springs:

```
?###???????? 3,2,1
.###.##.#...
.###.##..#..
.###.##...#.
.###.##....#
.###..##.#..
.###..##..#.
.###..##...#
.###...##.#.
.###...##..#
.###....##.#
```

In this example, the number of possible arrangements for each row is:

- ???.### 1,1,3 - 1 arrangement
- .??..??...?##. 1,1,3 - 4 arrangements
- ?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement
- ????.#...#... 4,1,1 - 1 arrangement
- ????.######..#####. 1,6,5 - 4 arrangements
- ?###???????? 3,2,1 - 10 arrangements

Adding all of the possible arrangement counts together produces a total of 21 arrangements.

For each row, count all of the different arrangements of operational and broken springs that meet the given criteria. What is the sum of those counts?

As you look out at the field of springs, you feel like there are way more springs than the condition records list. When you examine the records, you discover that they were actually folded up this whole time!

To unfold the records, on each row, replace the list of spring conditions with five copies of itself (separated by ?) and replace the list of contiguous groups of damaged springs with five copies of itself (separated by ,).

So, this row:

`.# 1`

Would become:

`.#?.#?.#?.#?.# 1,1,1,1,1`

The first line of the above example would become:

`???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3`

In the above example, after unfolding, the number of possible arrangements for some rows is now much larger:

- ???.### 1,1,3 - 1 arrangement
- .??..??...?##. 1,1,3 - 16384 arrangements
- ?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement
- ????.#...#... 4,1,1 - 16 arrangements
- ????.######..#####. 1,6,5 - 2500 arrangements
- ?###???????? 3,2,1 - 506250 arrangements

After unfolding, adding all of the possible arrangement counts together produces 525152.

Unfold your condition records; what is the new sum of possible arrangement counts?

```go
func toList(line string) (groups []int) {
	for _, val := range strings.Split(line, ",") {
		n, _ := strconv.ParseUint(val, 10, 8)
		groups = append(groups, int(n))
	}
	return
}

type Store map[string]uint

func (s Store) salva(line string, num []int) uint {
	line = strings.Trim(line, ".")

	key := line + fmt.Sprintf("%v", num)
	if val, ok := s[key]; ok {
		return val
	}

	s[key] = s.disposizioni(line, num, true)
	return s[key]
}

func (s Store) disposizioni(line string, num []int, variant bool) (sum uint) {
	line = strings.Trim(line, ".")

	if line == "" && len(num) == 0 {
		return 1
	}

	if line == "" {
		return 0
	}

	if len(num) == 0 && strings.Contains(line, "#") {
		return 0
	}

	if len(num) == 0 {
		return 1
	}

	if line[0] == '?' {
		if variant {
			sum += s.salva(line[1:], num)
		} else {
			sum += s.disposizioni(line[1:], num, false)
		}
		line = "#" + line[1:]
	}

	if len(line) < num[0] || strings.ContainsRune(line[:num[0]], '.') {
		return
	}

	if len(line) > num[0] {
		switch line[num[0]] {
		case '#':
			return
		case '?':
			line = line[:num[0]] + "." + line[num[0]+1:]
		}
	}

	if variant {
		return sum + s.salva(line[num[0]:], num[1:])
	}
	return sum + s.disposizioni(line[num[0]:], num[1:], false)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		sum1 uint = 0
		sum2 uint = 0
	)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		sum1 += Store{}.disposizioni(fields[0], toList(fields[1]), false)

		conditions := strings.Repeat(fields[0]+"?", 5)
		conditions = conditions[:len(conditions)-1]

		groupsStr := strings.Repeat(fields[1]+",", 5)
		groupsStr = groupsStr[:len(groupsStr)-1]

		sum2 += Store{}.salva(conditions, toList(groupsStr))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	log.Printf(
		"sum: %d, %d\n",
		sum1,
		sum2,
	)
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/)

- [input.txt](/documents/2023-12-12-input.txt)
- [Challenge](https://adventofcode.com/2023/day/12)
