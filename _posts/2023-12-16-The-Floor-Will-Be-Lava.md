---
title: The Floor Will Be Lava
description: Advent of Code 2023 [Day 16]
layout: default
lang: en
prefetch:
  - adventofcode.com
---

With the beam of light completely focused **somewhere**, the reindeer leads you deeper still into the Lava Production Facility. At some point, you realize that the steel facility walls have been replaced with cave, and the doorways are just cave, and the floor is cave, and you're pretty sure this is actually just a giant cave.

Finally, as you approach what must be the heart of the mountain, you see a bright light in a cavern up ahead. There, you discover that the beam of light you so carefully focused is emerging from the cavern wall closest to the facility and pouring all of its energy into a contraption on the opposite side.

Upon closer inspection, the contraption appears to be a flat, two-dimensional square grid containing **empty space** (`.`), **mirrors** (`/` and `\`), and splitters (`|` and `-`).

The contraption is aligned so that most of the beam bounces around the grid, but each tile on the grid converts some of the beam's light into **heat** to melt the rock in the cavern.

You note the layout of the contraption (your puzzle input). For example:

```
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
```

The beam enters in the top-left corner from the left and heading to the **right**. Then, its behavior depends on what it encounters as it moves:

- If the beam encounters **empty space** (`.`), it continues in the same direction.
- If the beam encounters a **mirror** (`/` or `\`), the beam is **reflected** 90 degrees depending on the angle of the mirror. For instance, a rightward-moving beam that encounters a `/` mirror would continue **upward** in the mirror's column, while a rightward-moving beam that encounters a `\` mirror would continue **downward** from the mirror's column.
- If the beam encounters the **pointy end of a splitter** (`|` or `-`), the beam passes through the splitter as if the splitter were **empty space**. For instance, a rightward-moving beam that encounters a `-` splitter would continue in the same direction.
- If the beam encounters the **flat side of a splitter** (`|` or `-`), the beam is **split into two beams** going in each of the two directions the splitter's pointy ends are pointing. For instance, a rightward-moving beam that encounters a `|` splitter would split into two beams: one that continues **upward** from the splitter's column and one that continues **downward** from the splitter's column.

Beams do not interact with other beams; a tile can have many beams passing through it at the same time. A tile is **energized** if that tile has at least one beam pass through it, reflect in it, or split in it.

In the above example, here is how the beam of light bounces around the contraption:

```
>|<<<\....
|v-.\^....
.v...|->>>
.v...v^.|.
.v...v^...
.v...v^..\
.v../2\\..
<->-/vv|..
.|<<<2-|.\
.v//.|.v..
```

Beams are only shown on empty tiles; arrows indicate the direction of the beams. If a tile contains beams moving in multiple directions, the number of distinct directions is shown instead. Here is the same diagram but instead only showing whether a tile is **energized** (`#`) or not (`.`):

```
######....
.#...#....
.#...#####
.#...##...
.#...##...
.#...##...
.#..####..
########..
.#######..
.#...#.#..
```

Ultimately, in this example, `46` tiles become **energized**.

The light isn't energizing enough tiles to produce lava; to debug the contraption, you need to start by analyzing the current situation. With the beam starting in the top-left heading right, **how many tiles end up being energized?**

```go
type Point struct {
	x, y int
}

type Beam struct {
	pos Point
	dir Point
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
    os.Exit(1)
	}

	grid := strings.Split(strings.TrimSpace(string(content)), "\n")
	energized := make(map[Point]bool)
	visited := make(map[string]bool)
	beams := []Beam{
		{Point{0, 0}, Point{0, 1}},
	}

	for len(beams) > 0 {
		var newBeams []Beam

		for _, beam := range beams {
			if beam.pos.x < 0 || beam.pos.x >= len(grid) ||
				beam.pos.y < 0 || beam.pos.y >= len(grid[0]) {
				continue
			}

			energized[beam.pos] = true
			key := fmt.Sprintf("%d,%d,%d,%d",
				beam.pos.x, beam.pos.y, beam.dir.x, beam.dir.y)
			if visited[key] {
				continue
			}
			visited[key] = true

			curr := grid[beam.pos.x][beam.pos.y]

			switch curr {
			case '.':
				newBeams = append(newBeams, Beam{
					Point{beam.pos.x + beam.dir.x, beam.pos.y + beam.dir.y},
					beam.dir,
				})
			case '/':
				newDir := Point{-beam.dir.y, -beam.dir.x}
				newBeams = append(newBeams, Beam{
					Point{beam.pos.x + newDir.x, beam.pos.y + newDir.y},
					newDir,
				})
			case '\\':
				newDir := Point{beam.dir.y, beam.dir.x}
				newBeams = append(newBeams, Beam{
					Point{beam.pos.x + newDir.x, beam.pos.y + newDir.y},
					newDir,
				})
			case '|':
				if beam.dir.y == 0 {
					newBeams = append(newBeams, Beam{
						Point{beam.pos.x + beam.dir.x, beam.pos.y + beam.dir.y},
						beam.dir,
					})
				} else {
					newBeams = append(newBeams,
						Beam{Point{beam.pos.x - 1, beam.pos.y}, Point{-1, 0}},
						Beam{Point{beam.pos.x + 1, beam.pos.y}, Point{1, 0}},
					)
				}
			case '-':
				if beam.dir.x == 0 {
					newBeams = append(newBeams, Beam{
						Point{beam.pos.x + beam.dir.x, beam.pos.y + beam.dir.y},
						beam.dir,
					})
				} else {
					newBeams = append(newBeams,
						Beam{Point{beam.pos.x, beam.pos.y - 1}, Point{0, -1}},
						Beam{Point{beam.pos.x, beam.pos.y + 1}, Point{0, 1}},
					)
				}
			}
		}
		beams = newBeams
	}

	fmt.Println(len(energized))
}
```

As you try to work out what might be wrong, the reindeer tugs on your shirt and leads you to a nearby control panel. There, a collection of buttons lets you align the contraption so that the beam enters from **any edge tile** and heading away from that edge. (You can choose either of two directions for the beam if it starts on a corner; for instance, if the beam starts in the bottom-right corner, it can start heading either left or upward.)

So, the beam could start on any tile in the top row (heading downward), any tile in the bottom row (heading upward), any tile in the leftmost column (heading right), or any tile in the rightmost column (heading left). To produce lava, you need to find the configuration that **energizes as many tiles as possible**.

In the above example, this can be achieved by starting the beam in the fourth tile from the left in the top row:

```
.|<2<\....
|v-v\^....
.v.v.|->>>
.v.v.v^.|.
.v.v.v^...
.v.v.v^..\
.v.v/2\\..
<-2-/vv|..
.|<<<2-|.\
.v//.|.v..
```

Using this configuration, `51` tiles are energized:

```
.#####....
.#.#.#....
.#.#.#####
.#.#.##...
.#.#.##...
.#.#.##...
.#.#####..
########..
.#######..
.#...#.#..
```

Find the initial beam configuration that energizes the largest number of tiles; **how many tiles are energized in that configuration?**

```go
type Point struct {
	x, y int
}

type Beam struct {
	pos Point
	dir Point
}

func simulateBeam(grid []string, startPos Point, startDir Point) int {
	energized := make(map[Point]bool)
	visited := make(map[string]bool)
	beams := []Beam{{startPos, startDir}}

	for len(beams) > 0 {
		var newBeams []Beam

		for _, beam := range beams {
			if beam.pos.x < 0 || beam.pos.x >= len(grid) ||
				beam.pos.y < 0 || beam.pos.y >= len(grid[0]) {
				continue
			}

			energized[beam.pos] = true
			key := fmt.Sprintf("%d,%d,%d,%d",
				beam.pos.x, beam.pos.y, beam.dir.x, beam.dir.y)
			if visited[key] {
				continue
			}
			visited[key] = true

			curr := grid[beam.pos.x][beam.pos.y]

			switch curr {
			case '.':
				newBeams = append(newBeams, Beam{
					Point{beam.pos.x + beam.dir.x, beam.pos.y + beam.dir.y},
					beam.dir,
				})
			case '/':
				newDir := Point{-beam.dir.y, -beam.dir.x}
				newBeams = append(newBeams, Beam{
					Point{beam.pos.x + newDir.x, beam.pos.y + newDir.y},
					newDir,
				})
			case '\\':
				newDir := Point{beam.dir.y, beam.dir.x}
				newBeams = append(newBeams, Beam{
					Point{beam.pos.x + newDir.x, beam.pos.y + newDir.y},
					newDir,
				})
			case '|':
				if beam.dir.y == 0 {
					newBeams = append(newBeams, Beam{
						Point{beam.pos.x + beam.dir.x, beam.pos.y + beam.dir.y},
						beam.dir,
					})
				} else {
					newBeams = append(newBeams,
						Beam{Point{beam.pos.x - 1, beam.pos.y}, Point{-1, 0}},
						Beam{Point{beam.pos.x + 1, beam.pos.y}, Point{1, 0}},
					)
				}
			case '-':
				if beam.dir.x == 0 {
					newBeams = append(newBeams, Beam{
						Point{beam.pos.x + beam.dir.x, beam.pos.y + beam.dir.y},
						beam.dir,
					})
				} else {
					newBeams = append(newBeams,
						Beam{Point{beam.pos.x, beam.pos.y - 1}, Point{0, -1}},
						Beam{Point{beam.pos.x, beam.pos.y + 1}, Point{0, 1}},
					)
				}
			}
		}
		beams = newBeams
	}

	return len(energized)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	grid := strings.Split(strings.TrimSpace(string(content)), "\n")
	rows, cols := len(grid), len(grid[0])
	maxEnergized := 0

	for y := 0; y < cols; y++ {
		energized := simulateBeam(grid, Point{0, y}, Point{1, 0})
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}

	for y := 0; y < cols; y++ {
		energized := simulateBeam(grid, Point{rows - 1, y}, Point{-1, 0})
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}

	for x := 0; x < rows; x++ {
		energized := simulateBeam(grid, Point{x, 0}, Point{0, 1})
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}

	for x := 0; x < rows; x++ {
		energized := simulateBeam(grid, Point{x, cols - 1}, Point{0, -1})
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}

	fmt.Println(maxEnergized)
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2023/day/16)

<details>
	<summary>Click to show the input</summary>
	<pre>
\..-.-..../.....|......\......................\............\.................|..............|........./.|.....
...-\...|............|......../.|........-../................|.-..............-....\.|..|.....................
.|..........\................\\|................../..........................\\../|...-...-.|..........\......
..................|\....|.....-.......................-........\.............................|/...............
.......-...........-......................./........../.......-............................\.....-.....-.../..
./.\...../....../..........\..............|............|....|........................\..../.../...|\..........
..................................../...............................................\..................-.\....
......../.........................|...............................|...............\|\........\....|....\......
.--.........................../.../.......-.\........../.\.........|...............\..|.....\-.....\....||\...
......\............|......../..\................-..............\-..-...|\.../-............/.........|...\.....
.........|.........-.........|............/............-..............\../.......-.|\\..|......./.............
..........................\........|.....\...../..-......-............../.-|.............\.\...../.......\....
.....\.................-........../........-....|.....-......|.....................................-......-\..
......|\......................................./...../........|...|....../..................-....-.-..../.....
.../......\...|./|..........\.......-./..-../.................\.....\......../........./..|......-...|.../....
--|................\.................|........../......\|......|.........................\....................
......../...........|./............../.....-.....-......../.................................-.....-...........
........................-.../..../..\|..............................|\..|.|\................/........-...\...|
\.........-....../...|.\............./......\.......\..../....\|./-|.........|......-.......\............|....
....|..........-.....|..../..........-/...-.............\..\...................\................|......|......
............-..|...........-..../...-........../.../...-....-....../\..............|........\..............--.
.....\.../......|..-..............|..-........................|......|.......|.........\............/...\...-.
|....................-.........../.....-........\.../..........|..............................................
..-.............|..../............|........................|..../.........................\..........-..-.../.
........-.|......\.........../................/..|..\................./.|....-................../...-......./.
\|..-.......\..............\.|/......-|.........../........../.-...............\....\....\..|..|..............
..........|.........\............\\.../../|...........-................|....................|....|-\/\...\....
.-........\............-.......\............-.............\.\......./...........././.|.\-...-.\./..........-/.
../.................-........//--...........|........./..../....|.....................................|....\..
.............|....|........|........./.....\..../..-..................../........|.....././...................
......................-....\..........--.....\..\........../...........................--/....................
.|.../.......-.......\.......|.............................-.............\..-.\../.......|............|......\
...../.......-......|.............../.../...........................................|......-..........-.......
.............\............................................../......................\......../.................
|....................-...-\..................\................/.../...........................................
..|....-............/....................../......../...../......./.......\.-................/................
.........../..-/...-...\..............-..../...............-|.........-............/........|.|.-...|...\.....
....../..-../.........-|...../...........................\..............-..................--...-.............
-.................-......../.....|....................\....................-../....................|./.....-..
......./...\......../.|.|.........-.................................../....|..................................
...|........\|\.........\...\...-......-...............................\.......//.|.....-.............../.....
........................../.|-...|....\..........\...................-..-............../............../.......
........../..-.......-....\.......\.............\.|................................|......../..............-..
.......-................................-..................|../...\....................................-......
...................................................................................../...\...................-
..\.-...................................|..............-.....|.........\..................../.........-....\..
...................../\......-...........--.../............|...........|.....\....//.....-........./...-|.....
......||.........|.-.........................|.........|.||..-....-......|.................\.-../.............
...|.........|.......\....\................................../.................../....\................\......
..................../|...............||.-/........|../..................................../.\......\...-......
......./........./..-............................-.................\...|...............-................-.....
.............................--...-....\......................................|..........\...-......./........
.......\...-.............../......\/.........................|........./..........|...........................
.........|.................................|.........../.|................................../...........-....|
...|...-...|...........................-....-.....................|../|........|...................../.....|..
.......-.....-..................\.-........|...-..../........\.......................\......|..........-......
|......................\..........-...............|........................../...........|....................
................|....|..................-......-/.......-....|..\.../\.|.....-..../...................../.|..\
............/........................../....-........-...............-...|...............|..\............./...
................./................\.............../...........\........\....../............\.....|........./..
..\......./.-.....................-............-./......|..........-.............\..../.......................
/..../....-........\................................\....|.............../\-.............................-|...
................\..................\.--.|......//....|..|....|....../..................|.|...|....-|-.........
...................\............\....-............\...........................................-./.............
.........................\../............|..\.........../................/..../.............-.................
..-..\.........../............|..................................../............/.....-..............|.....|..
..|/......\||.......\............/....|........|.-........./.............|.....................|./............
..........-............../...........................................|.............|..-..../.|................
..|.../..|......................./..-.-..............\../....................-..|............-...../.-........
.-....../..............\..................\................./....-...\....\...-.....\.....-.-..............\./
............./..../.\........-..\................................\....../.|.....|.......|.........|....\......
.-.....\.-......-.|.-.........../..|...........\....................................\........|.//./.......-...
......-.......-/.........................-...../........|............./..................../............-.....
.............\....|......-../...........\...................../...............................................
../.......................................-...../.|..\......|...-......../....|./..........\.........|........
....../.............\...................................\........\....-.........../..../..\..................|
.../...-.....-................/.........|....................|..............-.........\.......................
...................../..............-.......|......................./...|.....|../.|............../..........\
.././....../........-........|.../.........../|..........|....../......-.....\.............................\..
............|........|................/....\..............\............/.....\../..............|....../.......
|................\..../...........................|.......\......\.....\........../....../....................
........./...-...\......./....................-...................................-............\........-..../
............./............../.....................\.-..../.................\.........|.|............../.-...\\
.........|......-\......\.................../..|.........|......\....\..-..\-.................................
...............|../.......................-.../......\......|||................|..............................
.|...-............./...........-.......................|...............................|..............\.......
.|.............|.........................................-........|-.\.................-/.|...................
.|.......-||.........../..\....\.....................-....-....-.......-......-.........-|........|...........
.........\........-...|........../............/.................../......./........................-..........
............\.........................../..\.......-............/.......................|.........\...........
..\..\............/..................../..............-....../|.../...................|............./..../....
..\.\....\.../............/.../.........................../.-.......-.....\......|........|...../.........|.\|
.....-...........\..............-......-....-...../........-......|...../..\..|/........-............./.../...
.....\.......-........\...-.............................................|....||.........|...-......\..........
............|......................../.....-........././..............................................\.....\.
.\.........\......../........................-.............|........\../.../...-......\................|.....-
....|................|/.....................\........|........................./|.....\...-../.|......../-....
/..............-......-..............|..|.|.....................-\.|../...\..............................|.../
....\/.........|\....-..............\......-.....................\............\........................\......
......-...................\...............-/../........-.../....................\./..........-................
..........................\.......|......\.|\....../......\......../........-/............................/.|.
.......-....\...................................../..............|../.......|......./..........|--.......-....
...-...|.............................|........-.....-.|...........|.....|..............\..............-.......
-.\...........\......|.\.............................\...........|.........\.............../\.................
...../....................\../|...........................\.........|..................|........\.|..-../.....
.-....-../.................................\....../......|....../..............-...........-../...\.../.......
.\........-............................../..-...................|/..........|............................/....
../|..........|.........../...-....../.....|.-./...\....-....\|.\.....|........./.-.....\.....|./........|....
............\.....\...................-.\-.-\|...............................\...............\.....-..........
...|................-......................\....|..\.-..........|............../.....-...................../..
	</pre>
</details>
