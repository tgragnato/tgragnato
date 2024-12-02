---
title: Parabolic Reflector Dish
description: Advent of Code 2023 [Day 14]
layout: default
lang: en
prefetch:
  - adventofcode.com
  - en.wikipedia.org
---

You reach the place where all of the mirrors were pointing: a massive [parabolic reflector dish](https://en.wikipedia.org/wiki/Parabolic_reflector) attached to the side of another large mountain.

The dish is made up of many small mirrors, but while the mirrors themselves are roughly in the shape of a parabolic reflector dish, each individual mirror seems to be pointing in slightly the wrong direction. If the dish is meant to focus light, all it's doing right now is sending it in a vague direction.

This system must be what provides the energy for the lava! If you focus the reflector dish, maybe you can go where it's pointing and use the light to fix the lava production.

Upon closer inspection, the individual mirrors each appear to be connected via an elaborate system of ropes and pulleys to a large metal platform below the dish. The platform is covered in large rocks of various shapes. Depending on their position, the weight of the rocks deforms the platform, and the shape of the platform controls which ropes move and ultimately the focus of the dish.

In short: if you move the rocks, you can focus the dish. The platform even has a control panel on the side that lets you tilt it in one of four directions! The rounded rocks (`O`) will roll when the platform is tilted, while the cube-shaped rocks (`#`) will stay in place. You note the positions of all of the empty spaces (`.`) and rocks (your puzzle input). For example:

```
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
```

Start by tilting the lever so all of the rocks will slide **north** as far as they will go:

```
OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....
```

You notice that the support beams along the north side of the platform are **damaged**; to ensure the platform doesn't collapse, you should calculate the **total load** on the north support beams.

The amount of load caused by a single rounded rock (`O`) is equal to the number of rows from the rock to the south edge of the platform, including the row the rock is on. (Cube-shaped rocks (`#`) don't contribute to load.) So, the amount of load caused by each rock in each row is as follows:

```
OOOO.#.O.. 10
OO..#....#  9
OO..O##..O  8
O..#.OO...  7
........#.  6
..#....#.#  5
..O..#.O.O  4
..O.......  3
#....###..  2
#....#....  1
```

The total load is the sum of the load caused by all of the **rounded rocks**. In this example, the total load is `136`.

Tilt the platform so that the rounded rocks all roll north. Afterward, **what is the total load on the north support beams?**

```go
func tiltNorth(grid [][]rune) {
	rows := len(grid)
	cols := len(grid[0])

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if grid[row][col] == 'O' {
				newRow := row
				for newRow > 0 && grid[newRow-1][col] == '.' {
					grid[newRow-1][col] = 'O'
					grid[newRow][col] = '.'
					newRow--
				}
			}
		}
	}
}

func calculateLoad(grid [][]rune) int {
	total := 0
	rows := len(grid)

	for row := 0; row < rows; row++ {
		for _, ch := range grid[row] {
			if ch == 'O' {
				total += rows - row
			}
		}
	}
	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	tiltNorth(grid)
	fmt.Printf("Total load on north support beams: %d\n", calculateLoad(grid))
}
```

The parabolic reflector dish deforms, but not in a way that focuses the beam. To do that, you'll need to move the rocks to the edges of the platform. Fortunately, a button on the side of the control panel labeled "**spin cycle**" attempts to do just that!

Each **cycle** tilts the platform four times so that the rounded rocks roll **north**, then **west**, then **south**, then **east**. After each tilt, the rounded rocks roll as far as they can before the platform tilts in the next direction. After one cycle, the platform will have finished rolling the rounded rocks in those four directions in that order.

Here's what happens in the example above after each of the first few cycles:

```
After 1 cycle:
.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....

After 2 cycles:
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O

After 3 cycles:
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O
```

This process should work if you leave it running long enough, but you're still worried about the north support beams. To make sure they'll survive for a while, you need to calculate the **total load** on the north support beams after `1000000000` cycles.

In the above example, after `1000000000` cycles, the total load on the north support beams is `64`.

Run the spin cycle for `1000000000` cycles. Afterward, **what is the total load on the north support beams?**

```go
type Platform struct {
	grid [][]rune
	rows int
	cols int
}

func newPlatform(input []string) *Platform {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}
	return &Platform{
		grid: grid,
		rows: len(grid),
		cols: len(grid[0]),
	}
}

func (p *Platform) hash() string {
	var sb strings.Builder
	for _, row := range p.grid {
		sb.WriteString(string(row))
	}
	return sb.String()
}

func (p *Platform) tiltNorth() {
	for col := 0; col < p.cols; col++ {
		for row := 0; row < p.rows; row++ {
			if p.grid[row][col] == 'O' {
				newRow := row
				for newRow > 0 && p.grid[newRow-1][col] == '.' {
					p.grid[newRow-1][col] = 'O'
					p.grid[newRow][col] = '.'
					newRow--
				}
			}
		}
	}
}

func (p *Platform) tiltSouth() {
	for col := 0; col < p.cols; col++ {
		for row := p.rows - 1; row >= 0; row-- {
			if p.grid[row][col] == 'O' {
				newRow := row
				for newRow < p.rows-1 && p.grid[newRow+1][col] == '.' {
					p.grid[newRow+1][col] = 'O'
					p.grid[newRow][col] = '.'
					newRow++
				}
			}
		}
	}
}

func (p *Platform) tiltWest() {
	for row := 0; row < p.rows; row++ {
		for col := 0; col < p.cols; col++ {
			if p.grid[row][col] == 'O' {
				newCol := col
				for newCol > 0 && p.grid[row][newCol-1] == '.' {
					p.grid[row][newCol-1] = 'O'
					p.grid[row][newCol] = '.'
					newCol--
				}
			}
		}
	}
}

func (p *Platform) tiltEast() {
	for row := 0; row < p.rows; row++ {
		for col := p.cols - 1; col >= 0; col-- {
			if p.grid[row][col] == 'O' {
				newCol := col
				for newCol < p.cols-1 && p.grid[row][newCol+1] == '.' {
					p.grid[row][newCol+1] = 'O'
					p.grid[row][newCol] = '.'
					newCol++
				}
			}
		}
	}
}

func (p *Platform) cycle() {
	p.tiltNorth()
	p.tiltWest()
	p.tiltSouth()
	p.tiltEast()
}

func (p *Platform) calculateLoad() int {
	total := 0
	for row := 0; row < p.rows; row++ {
		for col := 0; col < p.cols; col++ {
			if p.grid[row][col] == 'O' {
				total += p.rows - row
			}
		}
	}
	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	platform := newPlatform(input)
	seen := make(map[string]int)
	target := 1000000000

	cycle := 0
	for cycle < target {
		platform.cycle()
		cycle++

		hash := platform.hash()
		if prev, exists := seen[hash]; exists {
			cycleLength := cycle - prev
			remaining := target - cycle
			skipCycles := (remaining / cycleLength) * cycleLength
			cycle += skipCycles
			for cycle < target {
				platform.cycle()
				cycle++
			}
			break
		}
		seen[hash] = cycle
	}

	fmt.Printf("Total load after 1000000000 cycles: %d\n", platform.calculateLoad())
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2023/day/14)

<details>
	<summary>Click to show the input</summary>
	<pre>
O...#.O......#...##....O#.#.OO...O.#OO#.....#O.O.....#......#.OO...O...O..O#O#....#...O..O...O.....O
.OO......O.OO.O.#.#O.O...#.##OO.....#.#O#...OOOOO..##O#O..O#O.##O...#O....O#.#.OO....#.........O.O.O
......O.O.O.....O.#.#.#OO#O#.O...O.#.......#O....#.O.#....OOO...#....O#O.......#......O...O##...O...
O...##O..#....OO.OO....O...O...O#..##.OOO..O..#OO....###....OO.O.O...........O...O..O.....#.#..O.O#.
............#.#.##O.O.#.#OO#........O..O.O..##....O.#..#O..O...O#.#.#.........O...........OO......O.
#.OO........#..#O.#......#....#....#......OO..OO...O#...#...#..##..O#.O......#.#..O..#O#...#..#.O#.O
O.O.##...........O...O..O...O.OO.......##....#...O.....##...#.....OO...O.O#...OO..........O...O..#O.
.O##....#.O#O....O.O.#.#O.O#..OO..##..#OO.#...#......#.#O#.....O...O#.O...#..O.#.#......#.##.O....O.
.........#..#.O..O..O.O...#.#.O...#.OOO.##..#.O.....OO#....O...OOO#.O#....##O..#.O.O.....O.O.O.O..O.
......#O.....OOO###.O..O####O.#.O...............#..O......##O...#.#..###...O....#O...OO.OO.O......O.
#.O....#.O..#O...#...#...O....OOO.#O.#.#..#.....#.#..OO..#OO.......#OO.OOO....O.#O.#.O......####..O.
OOOO...#.OO.#....O...O.O...##....#OO...O.##.O.....###..O........OO..O.O.......#.#...O..#..O.#..O....
OOO..##....O...O.....O#..O.O.#..O..#.#.#..O........O..#OOO.....#....O#OO.#.#.....O..#..O...O..O.#OO#
...##.#.##.#O....O#..O..O.........#.#.#O...O.......O...O#..OOOO....O...O.OO#.O.##...#....O..O.......
....##.###OOO.O...#.......OO....OOO....O.OO#..O.#..OO....OO......O..OO...OOO.OO#...#..OOO..O.###O...
O###O.....O.O.#O#.O.O#O.O###.....O..O........#....#..O#..#.#..O...........#......#OO........O..O####
.#.#.OOO.O.......#O..O.#..O.OO.OO.O..O.#.#...O.O.OO..##..O...#...#O#..O.........O.O..O.O#....O..#O..
.#..#O.#O..#.O.O.O...#OO....#...O....O..O.O...#..#.##O.O..O..#O.....OO##...O#.......O#O.#..........#
#O...#.#OO.O...#O...O...#.#.O#..#.O#.###.#......O..O....#..#.OO..#....O..O.O....#.O.O.O.#..O..OOO..O
O.....#...#.#OOO.OOOO..#O.O.#...O.##..OO..#O...#.....OO#..#.#...O##OO...O.....O..#...#...........O.O
.##.O#.O..#O#.O##.O.#.........##.O......#OO...OO.......#...#O.O#...O..#.#...O..#...#O...O.#....OO.OO
.O#OO.O..#...#..O..O...##.O....###O.....O...#.#OOO...OO...O.......#.#.#.OO....#OO.#..O..........OO..
.....#OO.O..O.O.O...O.O#.#.O...O..#...O##.......#......#O...#.#.#....#..O.O..#..........O..#.#..#O..
...O.O.O.OO...O#..O..OO..O#.#.O.#OO..O.O..OO..O#...#.O#...O.......#..#..O..##..O..#O......O.OO...O.O
#.O....O.O.#...O.....O.....O#......#.#..OO...#.....O.#.#.O...#.#....#....O#.O.#......#OOO.O...O#....
...OOO......#..O..O#.......#....O#OO..#O.#.#...#...O.O..OO..#.O...O..#.....OO..O..#...#.O...O#..O.OO
#OO.OO#.OOO...##............O....O.O..#.##.#.O.........##...#O.O..#................#..O.O.O#O.O#..##
.O.O.O##..#O.O#O...##.#OO#...#.OO#..O...O#...O....O#..OO....O....#...O...#OO.OO#O....#..O..#OO....O.
.OO.#.......#...O...OO..#..OO.O.O.OO#O.O.O...O....O#..#.O.........O..#.O...#.#...O.O......O........O
.O....##..#....#.....O.#...O.O.#....#....OOOO...O.O###.O........O.....O.O..#.##OOO.O.......#.#.O..O.
.#.....O...#.O...#........#...#....#OO.O...O#O........OOO...#.O....OOO#.O.....OO....#.#...#..OO..#..
OO.OO...O..#..#O....O.O....OOOOO.#O.#O.O...#.O....O#.O.O.#O.#...#..O...O.#O.O.#..O....OOO......#.O#.
.O##..OO.....#.OO..#..O..#...#.#........O..OO...#O..O....#..O..O#.....O#..##O..O#.##.#...#.O..O.O.O.
..#...#O.....OO#.....O..#.#O.#.#..#..O.....#...OO.#.....O..O.O##.O.....OOOO.#O.....O.#..O........#.#
.#.#O..O.O..#..O##.....O.OO...##.........O...####..O..O....#O..O.....##........O.#..#.#OO.#..O.##.O#
#...#.O....OO......#O...#O..O#..#O#O.#..#.....OO....OO#..#O....O.O..#.O.#.....O#.O.OO..O#.O.O..O.O#.
OO....OO......#O........#O...#.#..O##....O#..#.#O.O.#....O..O..O.O....O.......#..OOO#.........O....O
O.....#O.OO##......#O##..O..O...##..O..........O..OO.O..#O....O..OOO#OO.##..#....O#..#...O.O.O...O..
.............#.#.#.O...#.#..#......O.#.O..#....#..##....#.##.##.#.O...O.O..#.#.#...OO....O..#O..#..O
.O.#.....O.O..O..##O...OO.#.O....#..#.#...O.O#...#O.O.OO...O..#..OO...O...O...O.....#.O..O#..#.#..OO
.....O..OO......#.#.#.O.##.OO...O..O.#..#...#OO....OOO......O.OO.#.O..O#O...O.#......OO##O.O.#O..O..
O##O.#.#..#..OO.O.........#O.#..O.#.##.O#..OOOO..#....O...O.........O..O..###...O...O.#.....O..##.O.
..O...#O#O.O...O.#..#O.......OO...........#...#..##.O......#..O.O.#OO.OO.#O....#.O.....O..#..#.#..#.
O#..O..##O#....O..OOO..OO...#O..O.O...O#O...OO.O##.O#O..#......OO##O.##...OOO.O..#....O##O..#......#
..#O...O..O.O...#..O.O...#O.....O.O......##OO#..O......#..O.#.O......OO.#O........O.....OO...O...O.O
........OO.#.OO..O......#...#...#..OOOO.OO.#....O.OO...O.##.....O.O....#O.#.OO..##.##..#.#O#.OO...O#
....#......#O.OO.....OOO....OO..O.O.O....##..O..O.O#.O....O.#O.#O.....O..O........#O.....O.O.....O#.
#.O#.O#......##.....O.#O...###...##....#O#.#.#....OO...O..#.O....O..#..#..O..........O......#.O.O...
.O.#..O#.OO....O.#...#O....#....OO......#O.O..#O..........O.#..#...#OOO........#.#.........O.#..O...
#OO.OO..O..O.O.#O.O.....O.##..#.O#..#.O....#........#.OOO#.#.....#.#.......O#..##..#.OO....O#O#.....
.......#O.O....#.OOO..OO...O...#....O..#.#....OO..O....O.#.O.#O#..O#.#.O..O......##...OO..#....O.#..
.O..O#.#OO.O.O.O.##...O..O.O..O.##O...#.O..O..O..O..O.....O...#..OO#O..O###...#...#....O##.OO...#...
#O....#...##..O..#...#OO.O.O.#......O...O.#.#O..O...O....OOO..O.O..O...O.OO...O...#..O....O#.###....
....#..O.....##.O.......##.##..OOO.O.#.#.....#....O..#..O....#....O...#.O..O#..O##...#.##....O.#OO##
O....O...OO.O........#..O...#..#.....O......O.O.OO.OO..#....O..O.............##.O.....##O#.OOO......
........#.O.####O.OO.#.O..O.O.#.#....OO#.......O..#....##.OOO..O.O.O..#.O.#...O.O.O..O.##O..O#....O.
...#....#.....O....O.O.O##.O..#O.OO...#.##.O..###OO##.O......#..##.OO...O.......O.......#..O.#.O..O.
.O.O...O......OO.#....#O..O.O.O..OO.O..O##O...O....O..OO.#..O....O..#.O..OO.OO....O#O..#.O..#O...#.#
..O#O.#O..#...##..#.O...OO..O#.......#..O........#.OO#..O#.#.#O....##...O.........O......OO#.....O.#
....OO.......###..O#.....O##..#O..#..#.O.....##.......##..O#O....#....#O.O..#..#..O.###O..#.O#......
O..##..O...#..O...O#OOO....#.O##O.#........#.O..O.#..#.OO.#O.#.....##.O..O.O....#..#..OO....#..#..##
.#...O.OO#O.O..O.##O..O#...##...O....OO.O#....#...#O.O#....O..........#....#............##.......#..
.O.#........##O...#O..O..O....OO#O....#.O..OO.##..O..OO#.###OO.....#..#.......O.#.....OO........OO.#
.OO...O.#..##.....OO..O.O.O.#.###.##.O.#..O..#..O..O...##.#......O....O.O.........O..#.O.O#O.#O...#.
#..#.#..O...##.....OO#O.OOO.....#...#.....##.#...O....O....#..OO#.#.OO....#.....O.O...O#.#..#..O#O.O
.O#.O.#..O..O#...OO...O...O#...O..O..#.O..#...O#.#....O......O#...O...OO.O##.#O....O....#OO....#O...
OO..OOO.O......OO...OO#O...#.#.#O.OOO#....#.OO#.#.O.#......O.O.OO#..O..#..O.O#.O#.O....O..O...O.....
...OO....O...#OO#..#..O.#O#....O.....OO.O.#.OO##.....#..O#.OO..#.#.O#..#..O.#..O....#....O.#O...OO#O
..#.......OO..##.O#...O#....O.#.O..##..#O.......#O#......O..O.........O#..#OO..O..#O.#.#.O.O#.#OO#.O
...#O#...O...OO...OO..OO...#OO...O...##.O....##.#O##.O....#..O#..O.....O.......###.#.OO#...O......#.
.O.OO#O##O##...#..#....#..#O...O#..O.O##......#........#O.........#.O....O..#..O.....O..#......O.#.#
..#....#.O.O.O..O.....O.O#O.O..#......O............O.........O#.O....#.O#..O..O.#OO.....#......O...#
OO...#..O..O##.#.#O.O.#.#....O....O...O.#..O##...O...#.O...#.O.OO...#........O#O..O.O.O..O.....#.O..
.OOO....#...O..O.....O..O...O...O..O...O#...#O....O.O...#..OO##.O.#..O..........O#..#..#..OO#.....#.
O..#..O#....#O..##O..#.#.......OO..O..#...O....O.#....O..O.##........O...O..OO...#....OO#..O.#.#.#..
O#O.O.O#.#OO#..OO.O.#......#.....#.......##.O..#.O..#..OO..O.#.#.O.O.O#.#......O....####OO#.#O#OO.#.
.#.....O#..O#O#..OO...#.O.#....#.#..##....O##....#.........##OO...##...O......#.OO...#O....O.O.O#O..
O......#.O....O...O.....OO.#....#O.##....#.O.O...O..#...O....#.O....O.O....#.OO.#.......#...#..#...O
...#OO..#......#..OO.O.O.##.O....O.##O##.OO#O.O.O..O.#O#O.OO...........O...###.#......O..#..#O#..OOO
O..O.#..O.OO.O...OO.OO..OOOO....O...#....##.OO.O.....#.......O.O........O.#.......#.......#..#.O..O.
.OO.##...#O..O#....O....#..O#....#.....O...O#.....OOO#.....................#..#..#..#O.......O.#OO..
.#OO..O....O..#.....#O...OO.#....#.#.O...O.#........O...O..O..OO#..#.....#.......OOOOO.OO....O....O.
.O#...#.......O.O...O.O.O#...OO.O#...O#O.O.#......#......O..O##.##.....O.....##...........O#...O....
.#.O.O#.....O.#..OO......#O.#..O..O.O.O..#.#...O##.#O..OOO......#...#..OO...O.#.....#O.OOO...#.O..##
.........O..##OO.#OOOO.O.O..OOO...OO.#.....#...O.#O##..O#..#.#..#...#..O.O...#OO.O...O...#O...#..#..
.#.#O.#...#OO.O.........O#..O...OO#.....O....OOOO#.....O.O....#...O...........O#.#.O..O.#.O.#.O.#.#.
......O.#...O.....O.O..O#...O.#........#.O...O.O....#..O#....O.#.O.O....O...#OO...#........OO..O.O.O
....O.O..OO.....#...OOOO.O#............O#OO....OOOOOO..#..#O#..#.#.O.OO.#O.....O.#......##..O#..#O..
..#......#...#O.#..#.#....OOO#...O.#....O..#O.....O.....O.O......O.......O.O..O..O.OO.O.#OO..O#...#.
.O.#.O.OO..O#.O#OO..........O.......O..#..#.#..O...#...##.....####...#........##O.......#....#...O##
........#.#.O.#O.......#O...O...#.#O..#OOO..#O..###...O.#.#O...O.O#..O..O...O.O....###.OO.#OO..O.#.O
..#O....#..O..####.O#...O....O.#O.##.O....O.O......O#O...OO###...O..#........#.#O.O#OO.......#.....O
....O.....#.#.O.#..O.#O.#.#O.#.......OO.....OOO...O.O..#.#....#..O.O..#.OO.O..O..OO##.....#.O..OO#O.
O#O..OO.O....OO.#O......O......#...#.....O.#OO.#.....OO..O#.O..O.O...#....O.O..OO.O....#....OO..O...
...O.....##....#....#....#......O#O..OO.....OO#O.OO......#OO.......O.#..###OO...#.O..OO.#.O.O#OO..#.
.O.#.O..#O.....#.#O....##OO.#O..##..##.##...#OO.O.#O#..O..O.O..O..##O.O.....O.OOO..#..#O..O#..OO...O
..##O#O.O.#.........O....O.OO#O#O.....O.....O#O.OO.O......O.O.OO......#.....O.O.O.O.O.....#.O.......
.....#OO.#O........O#....O.#OO.OO........#.O....O.OOOOOO.#.#O......O.#...O#.O.O..OO.....O.OO...#.O.#
.....#................OO..OOOO..#.....O.###O.O...O..#O..........OO.#...OO.O#...O......O..OOOO.#..###
O.....O...#O.#..........#O......O..OOO#....O.OOOO....O..#..#O....O.....O#..O..#.O.O....#OO.#....#O..
	</pre>
</details>
