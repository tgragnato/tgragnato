---
title: Resonant Collinearity
description: Advent of Code 2024 [Day 8]
layout: default
lang: en
tag: aoc24
prefetch:
  - adventofcode.com
  - deno.com
---

You find yourselves on the [roof](https://adventofcode.com/2016/day/25) of a top-secret Easter Bunny installation.

While The Historians do their thing, you take a look at the familiar **huge antenna**. Much to your surprise, it seems to have been reconfigured to emit a signal that makes people 0.1% more likely to buy Easter Bunny brand Imitation Mediocre Chocolate as a Christmas gift! Unthinkable!

Scanning across the city, you find that there are actually many such antennas. Each antenna is tuned to a specific **frequency** indicated by a single lowercase letter, uppercase letter, or digit. You create a map (your puzzle input) of these antennas. For example:

```
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
```

The signal only applies its nefarious effect at specific **antinodes** based on the resonant frequencies of the antennas. In particular, an antinode occurs at any point that is perfectly in line with two antennas of the same frequency - but only when one of the antennas is twice as far away as the other. This means that for any pair of antennas with the same frequency, there are two antinodes, one on either side of them.

So, for these two antennas with frequency `a`, they create the two antinodes marked with `#`:

```
..........
...#......
..........
....a.....
..........
.....a....
..........
......#...
..........
..........
```

Adding a third antenna with the same frequency creates several more antinodes. It would ideally add four antinodes, but two are off the right side of the map, so instead it adds only two:

```
..........
...#......
#.........
....a.....
........a.
.....a....
..#.......
......#...
..........
..........
```

Antennas with different frequencies don't create antinodes; A and a count as different frequencies. However, antinodes **can** occur at locations that contain antennas. In this diagram, the lone antenna with frequency capital A creates no antinodes but has a lowercase-a-frequency antinode at its location:

```
..........
...#......
#.........
....a.....
........a.
.....a....
..#.......
......A...
..........
..........
```

The first example has antennas with two different frequencies, so the antinodes they create look like this, plus an antinode overlapping the topmost A-frequency antenna:

```
......#....#
...#....0...
....#0....#.
..#....0....
....0....#..
.#....A.....
...#........
#......#....
........A...
.........A..
..........#.
..........#.
```

Because the topmost `A`-frequency antenna overlaps with a `0`-frequency antinode, there are `14` total unique locations that contain an antinode within the bounds of the map.

Calculate the impact of the signal. **How many unique locations within the bounds of the map contain an antinode?**

```ts
interface Point {
  x: number;
  y: number;
}

function findAntennas(grid: string[]): Map<string, Point[]> {
  const antennas = new Map<string, Point[]>();
  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      const char = grid[y][x];
      if (char !== '.') {
        if (!antennas.has(char)) {
          antennas.set(char, []);
        }
        antennas.get(char)!.push({ x, y });
      }
    }
  }
  return antennas;
}

function calculateAntinodes(a1: Point, a2: Point, width: number, height: number): Point[] {
  const antinodes: Point[] = [];
  const x1 = a1.x + 2*(a2.x - a1.x);
  const y1 = a1.y + 2*(a2.y - a1.y);
  const x2 = a2.x + 2*(a1.x - a2.x);
  const y2 = a2.y + 2*(a1.y - a2.y);

  if (x1 >= 0 && x1 < width && y1 >= 0 && y1 < height) {
    antinodes.push({ x: x1, y: y1 });
  }

  if (x2 >= 0 && x2 < width && y2 >= 0 && y2 < height) {
    antinodes.push({ x: x2, y: y2 });
  }

  return antinodes;
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const grid = input.trim().split("\n");
  const antennas = findAntennas(grid);
  const height = grid.length;
  const width = grid[0].length;

  let antinodes: boolean[] = Array(width * height).fill(false);
  for (const [_, positions] of antennas) {
    for (let i = 0; i < positions.length; i++) {
      for (let j = i + 1; j < positions.length; j++) {
        for (const node of calculateAntinodes(positions[i], positions[j], width, height)) {
          antinodes[node.y * height + node.x] = true;
        }
      }
    }
  }

  const uniqueLocations = antinodes.filter((k) => k).length;
  console.log(`Number of unique antinode locations: ${uniqueLocations}`);
}

main().catch((err) => console.error(err));
```

Watching over your shoulder as you work, one of The Historians asks if you took the effects of resonant harmonics into your calculations.

Whoops!

After updating your model, it turns out that an antinode occurs at **any grid position** exactly in line with at least two antennas of the same frequency, regardless of distance. This means that some of the new antinodes will occur at the position of each antenna (unless that antenna is the only one of its frequency).

So, these three `T`-frequency antennas now create many antinodes:

```
T....#....
...T......
.T....#...
.........#
..#.......
..........
...#......
..........
....#.....
..........
```

In fact, the three `T`-frequency antennas are all exactly in line with two antennas, so they are all also antinodes! This brings the total number of antinodes in the above example to `9`.

The original example now has `34` antinodes, including the antinodes that appear on every antenna:

```
##....#....#
.#.#....0...
..#.#0....#.
..##...0....
....0....#..
.#...#A....#
...#..#.....
#....#.#....
..#.....A...
....#....A..
.#........#.
...#......##
```

Calculate the impact of the signal using this updated model. **How many unique locations within the bounds of the map contain an antinode?**

```ts
interface Point {
  x: number;
  y: number;
}

function findAntennas(grid: string[]): Map<string, Point[]> {
  const antennas = new Map<string, Point[]>();
  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      const char = grid[y][x];
      if (char !== '.') {
        if (!antennas.has(char)) {
          antennas.set(char, []);
        }
        antennas.get(char)!.push({ x, y });
      }
    }
  }
  return antennas;
}

function calculateAntinodes(a1: Point, a2: Point, width: number, height: number): Point[] {
  const antinodes: Point[] = [];

  for (let k = 1; k < width * height; k++) {
    const x1 = a1.x + k*(a2.x - a1.x);
    const y1 = a1.y + k*(a2.y - a1.y);
    const x2 = a2.x + k*(a1.x - a2.x);
    const y2 = a2.y + k*(a1.y - a2.y);

    if (x1 >= 0 && x1 < width && y1 >= 0 && y1 < height) {
      antinodes.push({ x: x1, y: y1 });
    }

    if (x2 >= 0 && x2 < width && y2 >= 0 && y2 < height) {
      antinodes.push({ x: x2, y: y2 });
    }
  }

  return antinodes;
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const grid = input.trim().split("\n");
  const antennas = findAntennas(grid);
  const height = grid.length;
  const width = grid[0].length;

  let antinodes: boolean[] = Array(width * height).fill(false);
  for (const [_, positions] of antennas) {
    for (let i = 0; i < positions.length; i++) {
      for (let j = i + 1; j < positions.length; j++) {
        for (const node of calculateAntinodes(positions[i], positions[j], width, height)) {
          antinodes[node.y * height + node.x] = true;
        }
      }
    }
  }

  for (let y = 0; y < height; y++) {
    let row = "";
    for (let x = 0; x < width; x++) {
      row += antinodes[y * height + x] ? "O" : ".";
    }
    console.log(row);
  }

  const uniqueLocations = antinodes.filter((k) => k).length;
  console.log(`Number of unique antinode locations: ${uniqueLocations}`);
}

main().catch((err) => console.error(err));
```

## Links

- [If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2024/day/8)
- [Solve Advent of Code 2024 with Deno and Win Prizes!](https://deno.com/blog/advent-of-code-2024)

<details>
	<summary>Click to show the input</summary>
	<pre>
............................................e.....
.................................................e
......O...................................Y.......
................3................E..........Y.....
.....O1............................e....j.........
......................6...........................
.....8......Z..........6..........................
...............3.............................u..j.
.E...........A............b...5...................
.........1.O.Z....................................
........G...0.E..........1..6.....................
......8................A..............g.B.........
..............3..............b...u................
........Z......8..b.........u....BO..........n....
....8....Z.............3.....................B....
...........................................Y......
...................G..............................
...0...............................j.......4......
.....0................A.................4......n..
..0..............x................n.e.............
.............................................4.Y..
.G.......................b................Q.......
.............x......................M.a...m.......
..E...........G.....................a.............
.................9.......Q..............7.n.......
...........................5......m....a..........
.........................5........................
.....X...J......5...............................M.
..............X..........................M........
........................W......o4...7........g.M..
..................................N............j..
..........................N..Q...............q....
.......J..............x....N.......a..............
....................x........N......U.............
.....2......J.....................w...............
...............6...................7.m........z...
.....................W..z..7.m.......o........gU..
........y.........................................
............y.........W.......Q...................
....2.......................................q.....
.y.....................q................o..z.....g
J..........9........................o.w........z..
..................................................
.............................................U....
....u..............X..........................q...
.....................................w............
..........9.......................................
......9..........2.y......................A.......
.......................................w..........
......................X...........................
	</pre>
</details>
