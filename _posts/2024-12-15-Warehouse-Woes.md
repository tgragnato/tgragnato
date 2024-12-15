---
title: Warehouse Woes
description: Advent of Code 2024 [Day 15]
layout: default
lang: en
prefetch:
  - adventofcode.com
  - deno.com
---

You appear back inside your own mini submarine! Each Historian drives their mini submarine in a different direction; maybe the Chief has his own submarine down here somewhere as well?

You look up to see a vast school of [lanternfish](https://adventofcode.com/2021/day/6) swimming past you. On closer inspection, they seem quite anxious, so you drive your mini submarine over to see if you can help.

Because lanternfish populations grow rapidly, they need a lot of food, and that food needs to be stored somewhere. That's why these lanternfish have built elaborate warehouse complexes operated by robots!

These lanternfish seem so anxious because they have lost control of the robot that operates one of their most important warehouses! It is currently running amok, pushing around boxes in the warehouse with no regard for lanternfish logistics **or** lanternfish inventory management strategies.

Right now, none of the lanternfish are brave enough to swim up to an unpredictable robot so they could shut it off. However, if you could anticipate the robot's movements, maybe they could find a safe option.

The lanternfish already have a map of the warehouse and a list of movements the robot will **attempt** to make (your puzzle input). The problem is that the movements will sometimes fail as boxes are shifted around, making the actual movements of the robot difficult to predict.

For example:

```
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
```

As the robot (`@`) attempts to move, if there are any boxes (`O`) in the way, the robot will also attempt to push those boxes. However, if this action would cause the robot or a box to move into a wall (`#`), nothing moves instead, including the robot. The initial positions of these are shown on the map at the top of the document the lanternfish gave you.

The rest of the document describes the **moves** (`^` for up, `v` for down, `<` for left, `>` for right) that the robot will attempt to make, in order. (The moves form a single giant sequence; they are broken into multiple lines just to make copy-pasting easier. Newlines within the move sequence should be ignored.)

Here is a smaller example to get started:

```
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<
```

Were the robot to attempt the given sequence of moves, it would push around the boxes as follows:

```
Initial state:
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move <:
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move ^:
########
#.@O.O.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move ^:
########
#.@O.O.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move >:
########
#..@OO.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move >:
########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move >:
########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

Move v:
########
#....OO#
##..@..#
#...O..#
#.#.O..#
#...O..#
#...O..#
########

Move v:
########
#....OO#
##..@..#
#...O..#
#.#.O..#
#...O..#
#...O..#
########

Move <:
########
#....OO#
##.@...#
#...O..#
#.#.O..#
#...O..#
#...O..#
########

Move v:
########
#....OO#
##.....#
#..@O..#
#.#.O..#
#...O..#
#...O..#
########

Move >:
########
#....OO#
##.....#
#...@O.#
#.#.O..#
#...O..#
#...O..#
########

Move >:
########
#....OO#
##.....#
#....@O#
#.#.O..#
#...O..#
#...O..#
########

Move v:
########
#....OO#
##.....#
#.....O#
#.#.O@.#
#...O..#
#...O..#
########

Move <:
########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########

Move <:
########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########
```

The larger example has many more moves; after the robot has finished those moves, the warehouse would look like this:

```
##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########
```

The lanternfish use their own custom Goods Positioning System (GPS for short) to track the locations of the boxes. The **GPS coordinate** of a box is equal to 100 times its distance from the top edge of the map plus its distance from the left edge of the map. (This process does not stop at wall tiles; measure all the way to the edges of the map.)

So, the box shown below has a distance of `1` from the top edge of the map and `4` from the left edge of the map, resulting in a GPS coordinate of `100 * 1 + 4 = 104`.

```
#######
#...O..
#......
```

The lanternfish would like to know **the sum of all boxes' GPS coordinates** after the robot finishes moving. In the larger example, the sum of all boxes' GPS coordinates is `10092`. In the smaller example, the sum is `2028`.

Predict the motion of the robot and boxes in the warehouse. After the robot is finished moving, **what is the sum of all boxes' GPS coordinates?**

```ts
type Position = { x: number; y: number };
type Boxes = Set<string>;

class RobotWarehouse {
  private robotPos: Position;
  private boxes: Boxes;
  private map: string[][];
  private moves: string[];

  constructor(mapStr: string, movesStr: string) {
    this.map = mapStr.split("\n").map(line => line.split(""));
    this.moves = movesStr.replace(/\n/g, "").split("");
    this.robotPos = { x: 0, y: 0 };
    this.boxes = new Set();
    
    for (let y = 0; y < this.map.length; y++) {
      for (let x = 0; x < this.map[y].length; x++) {
        if (this.map[y][x] === "@") {
          this.robotPos = { x, y };
          this.map[y][x] = ".";
        } else if (this.map[y][x] === "O") {
          this.boxes.add(`${x},${y}`);
          this.map[y][x] = ".";
        }
      }
    }
  }

  private isWall(pos: Position): boolean {
    if (
      pos.y < 0 || pos.y >= this.map.length ||
      pos.x < 0 || pos.x >= this.map[pos.y].length
    ) {
      return true;
    }
    return this.map[pos.y][pos.x] === "#";
  }
  
  private hasBox(pos: Position): boolean {
    return this.boxes.has(`${pos.x},${pos.y}`);
  }

  private getBoxLineCount(start: Position, dx: number, dy: number): number {
    let count = 0;
    let pos = { ...start };
    
    while (this.hasBox(pos)) {
      count++;
      pos.x += dx;
      pos.y += dy;
    }

    if (this.isWall(pos)) {
      return 0;
    }
    return count;
  }

  public runRobot() {
    for (const move of this.moves) {
      const newPos = { ...this.robotPos };
    
      switch (move) {
        case "^": newPos.y--; break;
        case "v": newPos.y++; break;
        case "<": newPos.x--; break;
        case ">": newPos.x++; break;
      }
    
      if (this.isWall(newPos)) continue;
    
      if (this.hasBox(newPos)) {
        const dx = newPos.x - this.robotPos.x;
        const dy = newPos.y - this.robotPos.y;
        const boxNewPos = { x: newPos.x + dx, y: newPos.y + dy };
        const boxCount = this.getBoxLineCount(newPos, dx, dy);
    
        if (this.isWall(boxNewPos) || boxCount == 0) continue;

        for (let i = boxCount - 1; i >= 0; i--) {
          const fromPos = {
            x: newPos.x + (i * dx),
            y: newPos.y + (i * dy)
          };
          const toPos = {
            x: fromPos.x + dx,
            y: fromPos.y + dy
          };
          this.boxes.delete(`${fromPos.x},${fromPos.y}`);
          this.boxes.add(`${toPos.x},${toPos.y}`);
        }
      }
    
      this.robotPos = newPos;
    }
  }

  public getSum(): number {
    let sum = 0;
    this.boxes.forEach(pos => {
      const [x, y] = pos.split(",").map(Number);
      sum += y * 100 + x;
    });
    return sum;
  }

}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const [mapStr, movesStr] = input.trim().split("\n\n");
  const robot = new RobotWarehouse(mapStr, movesStr);
  robot.runRobot();
  console.log(robot.getSum());
}

main().catch((err) => console.error(err));
```

The lanternfish use your information to find a safe moment to swim in and turn off the malfunctioning robot! Just as they start preparing a festival in your honor, reports start coming in that a **second** warehouse's robot is also malfunctioning.

This warehouse's layout is surprisingly similar to the one you just helped. There is one key difference: everything except the robot is **twice as wide**! The robot's list of movements doesn't change.

To get the wider warehouse's map, start with your original map and, for each tile, make the following changes:

- If the tile is `#`, the new map contains `##` instead.
- If the tile is `O`, the new map contains `[]` instead.
- If the tile is `.`, the new map contains `..` instead.
- If the tile is `@`, the new map contains `@.` instead.

This will produce a new warehouse map which is twice as wide and with wide boxes that are represented by `[]`. (The robot does not change size.)

The larger example from before would now look like this:

```
####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##....[]@.....[]..##
##[]##....[]......##
##[]....[]....[]..##
##..[][]..[]..[][]##
##........[]......##
####################
```

Because boxes are now twice as wide but the robot is still the same size and speed, boxes can be aligned such that they directly push two other boxes at once. For example, consider this situation:

```
#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^
```

After appropriately resizing this map, the robot would push around these boxes as follows:

```
Initial state:
##############
##......##..##
##..........##
##....[][]@.##
##....[]....##
##..........##
##############

Move <:
##############
##......##..##
##..........##
##...[][]@..##
##....[]....##
##..........##
##############

Move v:
##############
##......##..##
##..........##
##...[][]...##
##....[].@..##
##..........##
##############

Move v:
##############
##......##..##
##..........##
##...[][]...##
##....[]....##
##.......@..##
##############

Move <:
##############
##......##..##
##..........##
##...[][]...##
##....[]....##
##......@...##
##############

Move <:
##############
##......##..##
##..........##
##...[][]...##
##....[]....##
##.....@....##
##############

Move ^:
##############
##......##..##
##...[][]...##
##....[]....##
##.....@....##
##..........##
##############

Move ^:
##############
##......##..##
##...[][]...##
##....[]....##
##.....@....##
##..........##
##############

Move <:
##############
##......##..##
##...[][]...##
##....[]....##
##....@.....##
##..........##
##############

Move <:
##############
##......##..##
##...[][]...##
##....[]....##
##...@......##
##..........##
##############

Move ^:
##############
##......##..##
##...[][]...##
##...@[]....##
##..........##
##..........##
##############

Move ^:
##############
##...[].##..##
##...@.[]...##
##....[]....##
##..........##
##..........##
##############
```

This warehouse also uses GPS to locate the boxes. For these larger boxes, distances are measured from the edge of the map to the closest edge of the box in question. So, the box shown below has a distance of `1` from the top edge of the map and `5` from the left edge of the map, resulting in a GPS coordinate of `100 * 1 + 5 = 105`.

```
##########
##...[]...
##........
```

In the scaled-up version of the larger example from above, after the robot has finished all of its moves, the warehouse would look like this:

```
####################
##[].......[].[][]##
##[]...........[].##
##[]........[][][]##
##[]......[]....[]##
##..##......[]....##
##..[]............##
##..@......[].[][]##
##......[][]..[]..##
####################
```

The sum of these boxes' GPS coordinates is `9021`.

Predict the motion of the robot and boxes in this new, scaled-up warehouse. **What is the sum of all boxes' final GPS coordinates?**

```ts
type Position = { x: number; y: number };
type Boxes = { left: Set<string>; right: Set<string> };

class RobotWarehouse {
  private robotPos: Position;
  private boxes: Boxes;
  private map: string[][];
  private moves: string[];

  constructor(mapStr: string, movesStr: string, isWide = false) {
    mapStr = mapStr.split("\n").map(line =>
      line.split("").map(char => {
        switch (char) {
          case "#": return "##";
          case "O": return "[]";
          case ".": return "..";
          case "@": return "@.";
          default: return char + char;
        }
      }).join("")
    ).join("\n");

    this.map = mapStr.split("\n").map(line => line.split(""));
    this.moves = movesStr.replace(/\n/g, "").split("");
    this.robotPos = { x: 0, y: 0 };
    this.boxes = { left: new Set(), right: new Set() };

    for (let y = 0; y < this.map.length; y++) {
      for (let x = 0; x < this.map[y].length; x++) {
        switch (this.map[y][x]) {
          case "@": 
            this.robotPos = { x, y };
            this.map[y][x] = ".";
            break;
          case "[":
            this.boxes.left.add(`${x},${y}`);
            this.map[y][x] = ".";
            break;
          case "]":
            this.boxes.right.add(`${x},${y}`);
            this.map[y][x] = ".";
            break;
        }
      }
    }

    this.print("");
  }

  private isWall(pos: Position): boolean {
    return this.map[pos.y]?.[pos.x] === "#";
  }

  private moveVerticalBoxes(start: Position, up: boolean) {
    const dy = up ? -1 : 1;
    const connectedBoxes = new Set<string>();
    const toCheck: Position[] = [start];

    while (toCheck.length > 0) {
      const pos = toCheck.pop()!;
      if (this.isWall(pos)) return;
      const key = `${pos.x},${pos.y}`;
      if (connectedBoxes.has(key)) continue;
      if (this.boxes.left.has(key)) {
        connectedBoxes.add(key);
        toCheck.push(
            {x: pos.x + 1, y: pos.y},
            {x: pos.x, y: pos.y + dy}
        );
      }
      if (this.boxes.right.has(key)) {
        connectedBoxes.add(key);
        toCheck.push(
            {x: pos.x - 1, y: pos.y},
            {x: pos.x, y: pos.y + dy}
        );
      }
    }

    if (connectedBoxes.size === 0) {
      console.log("Logic error in move vertical boxes -- no connected boxes")
      return;
    }

    const moves = Array.from(connectedBoxes)
        .map(pos => {
            const [x, y] = pos.split(',').map(Number);
            return { pos, x, y };
        })
        .sort((a, b) => up ? a.y - b.y : b.y - a.y);
    const boxMoves = new Map<string, string>();
    for (const {pos, x, y} of moves) {
        const targetY = y + dy;
        const targetPos = `${x},${targetY}`;
        
        if (this.boxes.left.has(pos)) {
            boxMoves.set(pos, targetPos);
        }
        if (this.boxes.right.has(pos)) {
            boxMoves.set(pos, targetPos);
        }
    }
    for (const [fromPos, toPos] of boxMoves) {
        if (this.boxes.left.has(fromPos)) {
            this.boxes.left.delete(fromPos);
            this.boxes.left.add(toPos);
        }
        if (this.boxes.right.has(fromPos)) {
            this.boxes.right.delete(fromPos);
            this.boxes.right.add(toPos);
        }
    }

    this.robotPos = start;
  }

  private moveHorizontalBoxes(start: Position, left: boolean) {
    const dx = left ? -1 : 1;

    let boxCount = 0;
    let pos = { ...start };
    while (
      this.boxes.left.has(`${pos.x},${pos.y}`) ||
      this.boxes.right.has(`${pos.x},${pos.y}`)
    ) {
      boxCount++;
      pos.x += dx;
    }
    if (this.isWall(pos) || boxCount === 0) {
      return;
    }

    for (let i = boxCount - 1; i >= 0; i--) {
      const fromPos = {
        x: start.x + (i * dx),
        y: start.y
      };
      const toPos = {
        x: fromPos.x + dx,
        y: fromPos.y
      };
      if (this.boxes.right.has(`${fromPos.x},${fromPos.y}`)) {
        this.boxes.right.delete(`${fromPos.x},${fromPos.y}`);
        this.boxes.right.add(`${toPos.x},${toPos.y}`);
      } else if (this.boxes.left.has(`${fromPos.x},${fromPos.y}`)) {
        this.boxes.left.delete(`${fromPos.x},${fromPos.y}`);
        this.boxes.left.add(`${toPos.x},${toPos.y}`);
      } else {
        console.log("Logic error in move horizontal boxes")
      }
    }
    this.robotPos = start;
  }

  public runRobot() {
    for (const move of this.moves) {
      const newPos = { ...this.robotPos };
    
      switch (move) {
        case "^": newPos.y--; break;
        case "v": newPos.y++; break;
        case "<": newPos.x--; break;
        case ">": newPos.x++; break;
      }

      if (this.isWall(newPos)) continue;

      if (
        !this.boxes.left.has(`${newPos.x},${newPos.y}`) &&
        !this.boxes.right.has(`${newPos.x},${newPos.y}`)
      ) {
        this.robotPos = newPos;
        this.print(move);
        continue;
      }

      switch (move) {
        case "^": 
          this.moveVerticalBoxes(newPos, true);
          break;
        case "v":
          this.moveVerticalBoxes(newPos, false);
          break;
        case "<":
          this.moveHorizontalBoxes(newPos, true);
          break;
        case ">": 
          this.moveHorizontalBoxes(newPos, false);
          break;
      }

      this.print(move);
    }
  }

  public print(move: string) {
    let matrix = "";
    for (let y = 0; y < this.map.length; y++) {
      for (let x = 0; x < this.map[y].length; x++) {
        if (this.robotPos.x === x && this.robotPos.y === y) {
          matrix += "@";
        } else
        if (this.boxes.left.has(`${x},${y}`)) {
          matrix += "[";
        } else if (this.boxes.right.has(`${x},${y}`)) {
          matrix += "]";
        } else {
          matrix += this.map[y][x];
        }
      }
      matrix += " " + move + "\n";
    }
    console.log(matrix);
  }

  public getSum(): number {
    let sum = 0;
    this.boxes.left.forEach(pos => {
      const [x, y] = pos.split(",").map(Number);
      sum += y * 100 + x;
    });
    return sum;
  }
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const [mapStr, movesStr] = input.trim().split("\n\n");
  const robot = new RobotWarehouse(mapStr, movesStr, true);
  robot.runRobot();
  console.log(robot.getSum());
}

main().catch(console.error);
```

## Links

- [If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2024/day/15)
- [Solve Advent of Code 2024 with Deno and Win Prizes!](https://deno.com/blog/advent-of-code-2024)

<details>
	<summary>Click to show the input</summary>
	<pre>
##################################################
#..OO......OOO.......#O.#O#..O#.O.O.#..O#.O....OO#
#...OO..O..#...O.OO...O...OOOOOO#O...#...O..O#.OO#
#...OO#O....O.......O.......OO..O..O..O#OOOOO.#.##
#O....#...O..#...#OO.O...#O..O.O#O.O....O.....O..#
#OO..O....O........OO..#O...##...OO.........OOOO.#
#.O..O.#..OO#......O.......O..O.......O.O.#..#..##
#.........#.OOOOO....O...O..O.OO....O....##..#.O.#
#.O.O.OO.#..O..O...#O.......OOO.OO...#O.#.#O.O...#
#...#...OO...#OO..O..O#O.OO......O.O...O....O..O.#
#OO.O...O.O#.#.OO....OO...O#O.O.O....O...##...O.##
#.#..O.O#....O...OO.OO..OO.......O#OO.....OOO.#..#
#.......OOO...O..O.O.O..O..#O.O#OO..O..OO.O......#
#....O.O......O...#O.O#O..O..OO.OO...OO..OO...#.O#
#O....O.O.O...OO.O..O....O.O..OO..#OOO.OO.O......#
#O..#.........O..O.O..#..#.....#O.O....#.O...#.O.#
#OO#O..#....O.O...OOO..#..O.#...O..O...O..O.....O#
#O...O.#..O....O....O.........O....#.............#
##.O.....O...O...#O#O...OOOO...O.O....O..OO.....O#
##...O..#.O.#..OO.....#...OO.#....O.......O..O...#
#OOO#...OOOO....#.#.......OO..O.O...O.O..O..O...O#
#......O.O..O...O.#.OO..O....#...OOO........OO#..#
#.OOO.O###..O.....#.OO...OO#O.O.....O#.OO#.....O.#
#.#O.......O#..O.O.........#O...OO..#O...OO..OO.O#
#..O...O......O.OO...O#.@.........OO..#.........O#
#...##.O.OOOO.....O..OO.......O.#O.....O#OO.OOO..#
#OOO#O..##........O......OO....O..##.O..O...O...##
#OO......##.O.O#.#.....O.O.O.#O.O......O.O##.....#
#.O.O.O..O....O.O.....O...#.......#.O###.#O..OO..#
#.OO.OO.....O#.O..O#OO..OOO..OO...O...#O#O.##.#O.#
#O...#.O#...#..O.O....O....OO.....O...OOO....O..O#
#O.......OO..#O.OOOOO#O...OO.OO....OO........#..O#
#.O...O....#.#.OO...O...O.OO...OO..O.O...OO.....O#
#O.OOOOO....O..O..O....#.O........O....O.OO..O...#
#..#...O#...O...........OOO...O.....O.....#......#
#OO.O...........#.....#.#OOO..O.....O....#O.O....#
#.#...OOO..O.#.O.......##O..OO.O...O.#O.###O.....#
#O...OO...O.#.O......##O......O..O......O.....O..#
#.....#O..O......OOO...O.....O.......O##...#O....#
#......O........#O..#...OOO...O...O....#.......O.#
#O...O.O..O.O#...O##.OOOO.O.O.O.........O....#..O#
##OO.......#OO#.O...O..O.O#....O........#.OO...#.#
#.O...#O...O...........OOO...#..O......#.OO..OO.##
#O.#......#......O#...OO...O....#..#O...#.#...#..#
#.OOOO....O........O.....O..#.O..O#O.......O..O.O#
#O...O...#..O...OO...#O#.OO..O....#OOO......O.#..#
#...#..O..O..OO...#O........#.#..OOO#....O.O.#O..#
#O.OOO...O..O..#.O..O..O.O.O.O....OO.O.O.#.....#O#
#OO..#O.O.#.#.O.....O#.....OO.O...OO.....O...OO..#
##################################################

v>>^^<v<^>^v^>>>^v^^<v^<<>v<^>v^><vv>^v^^>>>v<<>>>>>v>^>^>><v>v^vv><>^^>v^<>>^><^^^v^<>v<^^v>^v<<v^<^<^^<<<^<^>^^v<>v<<^<>v<^>v^>^^>><^<<v>v<<v>><<v^^vv<<^>^><vvv<>^><>>><^>v<><vv<<>>v^>v<^><>v<<>><<<v<<>v<<v^^^^>>v^^<v^^>><>^v>vv>v<>^<v^^>^><v^<<>vv>vv>><>^>^>^<v^<<>v>v>v<>><<^>v^<^<v<v<vv<>v<<^>v>>^v^<v<^v^^>><v>vv<^>v^>><>>^>v^v><v>^>><<>^<^>><<><<v<>>>^<^><v>>^^^^vv^vvv>^vv^^v^^<v><vv>^^>v^vvvv^>^^^v<^v^<<vv<^>><^^^<<v>>><><<>^v<<<><^v<>^>^v>^^vv<<><>^<>>^v<^v><v<vv>>v^^><^>v<>^<^^>v<>><vvv<^<<<><vv^<<>><^v<<v<><>vv<v><^v<<^<<v>>vv<<>>>vv^<<<<<^^<><^>vvv>>><<^v<^v^vv^^<v^vv<<>v>v^v<^<><<<<><^>><><v>v><<v<^>^^vvv^^^<^^<>^>v<<<>>v>^^<<<>>>>>>^>v^v^>^>><^><v^v<<v<^>^v^v^v<^<^^<<>^^>^><^<>>>v<^^v<>>v>^<^>>>>^<vv^<^>^>>^v^v^><<^>><>v<<^>v>>v<><>vv>vv><^vv<v<v^>>^v^v<>><^><>^^<v>^^><>^>vvvvv^>v>^><<>><^>v><^<<^v<<vv><^<v^v>^^v^v^>>>^v<<^^>v><^^v<><>v>^^<>^^<^<^<<vvvv<v>v>^v<vv<v><v>^v<^<<<><<<vv><><>v<vvv^<vv>^<vv<v^<><^vv<<^^>>vv^^><vv<v>^<v<<<<>vv<^<<^^>vvv^v^><^v>v^<<^v^><v<<v<v<^^<vv
>v>v<vv>v>^>^>vvv><>v<>>^v>^vv<<^<<v<><vvvvv><^^^v<^v^><<>v>v<>><>v<><<<>vvv^>>^^<^v>^^v^<<>v>v><v><v>>^vv<<><<^^v<<<^vv<^>>>^^<<>^v^^><>>v<>><^v^v><<v><>><^>>^^^><v^>>vv^>><<<^><v>>v<<<<><>vv<>>>><<<^<^vvv><^v<<^v<v^v>>^<^^vv<<^v^>>v>vv^<><^<^><^<^<><>^v>vv<<v><<v>><vv^>^<^v<vv<<^v><^^>v<<<<<>^^>>^>v>^^<v<<^^^^v^>v<>vv<<>v^>><<v<^v<>vvv^v^<>>v<vv^<v^<v^<<^<<^^>vv>v>v>^vv^^v<>v<v><^v^<v>v^><<^vvv>^>v>v<vv<>^vv^>v<<><<>^^v>>v<>^<<<^<<>v<<>>v>><<^^>v<<<vv>>><^<>>>v><<<><^>^v^^^^^<><>^>^<^^^<^<<^^^>vv<><v>^v>>^^v^v^v>><<<<<^v^<<v<v^<^^^<<>>^^v<^>>>>vv>v^>^<<<>>>^><v<^>vv<vv<^>^vv>>vv^>v^^>^v^<<^<^v>>>^^>^^v^>vv^<><<>^vv<<<v>>v>>v^<vvv<vv>^v^^<<>><v<>^<v>>>>>^<<<^><<v^<<<<>^<^^^v<<v^v<^<><v^v^vvv>^>v<<^^^<^>vv^^>^v><^>^v<<^^>vv>>><^><vvv>v<v<<<>^v<^v>vv^>^v^<^<v^^>><<>^><>v>^^<<<>v>v^><>^^^>><><v>>v^<vvvv^v^>>>>>>>>v<v<<^vv><^<^<>^^^>^<^><v<^^<^v^v^v<v>^^<v><^>^^>v^^vv>^><v>v>^^<v^>^v<v<v^<>^^<><^^<<>^<>v<^^>><^v<>>>^><<^<<^^^v>v>^^^v<<^^>^vv>^v><<>>v<v>^v>v<>v><>><v>vv^v^<>^^v>^^>^v<^>v<>
^<><v<^<<<^v>^vv<<>^v^v^<<<^>>>>^^>^<>v><><^^<v^v^^v^><<^^<>><^v^v>><<<^<<<v<>^v>>>v>><^^v<vvvv>>><v<>^v<<v^>v>v>v^<v>^v>>vvv^^<<<<^><>^>>v>vv><<<<<<><><^><v>v<v^>^v<<^^>^<<^<^v<^><>>>>v^>>><<^^v><vv<<v^>^^^>v<v>><v^^^^<>v<<>v<^<><<vv^>^<><^v>v<v^<v^^^>v^><^<>^^v>v>><v<<v>^<v^<><<<<v>>^v<>vv<>^<>v^v^>^^^<>^<>>^^v>v^^^>v^>v<vvv><<>>^^v>^v>>v><vv>^v^<<>^v^>^^^<>vv<>><><<>>^<>^v<><><<v>>^>vv^<^v^<^><^>^^^v<^v>>^<>>^^<><>>v^<vv^<^<>^<<>vv^<<>v^>>vv<<<^v^<<^v^^v<>v^<vv^<>>>vv<><v<<><<vv><<>v^vv^^<v<^<^^v>>v<vvv>^v<><<><v^>v^^v<^<v<v>v>^^vvv>><^^>>v<<<^v^<^<v>^^<^<^>^>^>><><<^vv^<vv><v><<>v>v^<^v^>>^v<<><<v>v^>^^>v>vv<v><^v<v<<>^>><v<v^vvv><^^^^>^v<v^<<<>vv^<>>>>><^v<<v<<v<^><<v>^><><>^>vv<<>^vv<<v><<>^>vvv><<^v><^<>>><><^<>>^>>^v<<<>v<^v<>vvv<><>>>^><v^<>v<<v^>>v^^<v>vv<>v<v>><>>^v>><<>v>v^v>^>>>vv^^v<>v^>>vv^<^v>v<v>v<<>v>^v^v^^>><v^^^^>v>>v<^<><vvvvv<<^^v>>>v<^<^<>^>v^>v><v<^v<<<>^^vv^^^<><^v^v>vv^<>^>>^^v^>vv><^^>^<^>^>v^>^>><v^v<v^<<v>v<^v^v^<v><v>>v<<<^<>>v<><<v<vv^v>^<^vv^v<vv>v><vv<v
<<>^vvvvvv>^<<v^>>v^><>vv<vv>^<<^v<>v><>^>>>vvvv<v>^>^<>^<>>>v<v>>v<<>>v>^>^^^vv^v><v>^<v>v<v>><>v><<v>^<v><<vv^^v>>v<<>>^<^><v^v>>v^^<^^^v<<<<^><vvv>v<^v><>>><^>><v^>>^<^><><^<<><<>^<v><<v<^<^vvvv><^vv<v<v<<><v^^>^^^v<>^>^v^<<^><v>>v>^^^^<v<<vv><>>v^>^^>^^<<vvv<v>v^vv^>v<^^^<^^v<^v^>^<^>^v<vvvvv><^^v>v^><^>^v^<>>^>>^^>^<vv>v^>^v<<^<<><v<>v>>v<><^>>>v^vv^<^v^^>>^<v<v<^>^^^>>>v^^v<^<v<^^<><<v^>^<<^<vvvv<vv<vv>^>vv^<>^>>v><<<<v>^^>^>^><vv>^<v<^vv<<^v>^v^<>>>^v<<<v>>><vv<^>v<^vv<<^><v<><<>^<>vv<v^v>>v>vv<^><>v>>^v<<v>^>v>>><<<<<v<<v<v<<^^<^><<^><v^vv<<<^v>v<v>>>v<^^<<><>v><^^v^^<>>v<v>^><<>v^>>v><^>^<>vv><v^vv><>^>>v^>>v<vv^<<>>><^^v>^<<<<^<>^><<v^^^>v^v^^vv>>^>v>^^vv<>^<<^^v<v>v^<^<>^<v>^v^<>vvv^^>v<v^<^^^v<>v<^<^<<v^><<v>vv<>>^>>v^vv<^<^^^^v^v>^<<<v<>v>^v>v>vv<>^<<>v>v^v^<^>v>>v<^v<^<><^v^^vv^>^<v<>^<^^<^v>>^<><v<v^>>v>>>v>vv>^^vvv>>^^<>>>^><>^^^<<<^^<<<<>><^>>vv>>>^^^><<><<<>^v>v^<^^^<>v<^<vvv>^<<<^^<>^v><>^^^>><<v><vv^<<^<<vv<><>v<<>^><<^^<^^v>>>>vv><<^<^<v><v^>>><>v<v><<<v<>v>v>vvvv>
<vv>v^vvvv>vv^^^<v^vv<<^v<<><^v^v^>^^<^^^<^<<>><<>^>^<<<<v^^>><<^v^><^v<>^>>^<<^<>v<>v>^>><^vv>v^v<><v^><<<v^^<>>v^<<<v^vv>><<^>v><<^><<vv^>v^^<v><<<^<vv<^<<<^>><>^<v^^^^>^v<v>v^v<^<^v<>v><^>^>><<<<>v^v><>^^><>^^^<>^<>v>^vv^^<^<<v<<^^^^v>v^>^>v^><<^^v^^>^>vv>vv<^v<>^v<^^><>>^>v^^v^^>>^^<><v><v<>v>^>^>><v<<<^><<><^<>v>>>>>>>^v^>vv^^<^<>^v>v>^^<v<<<^>><^vv>>^vv<^<^vvv^<<>^^>^^v^>^<<<>^^^v<<>>^<><^>v>^<<^v>>v^^>>><^vv<^^>^v>>^v^>v><v^^<><><<^v<^><^<<v<<>><>v^<>v<<v^>v>^v>v^^>v>>>><vv>^>^><<v^><v>^><>v<v^<^^<^v^^<>>^^v<v^^<>><^><^v<>vv^><>>>><<>^v^<<<^^<>><v^<v>v^>>^>><v><><^v^>>v<<>>v><v^<><<<^^^vv<<<<<><^^>v<^^<v<>>v^<<<^v>v<^^^vvv^^v^<^>vv^<><>>>v^<^^>^><>>>vv<^^<>^><<vv<>>>^<v^<vvv<<<^^<^<vvv><>v>>v^v<^<<>^vv>v>v^vvv<^^^v<^^<^>v^vvv>v>v<<>><^^>v<><^>vvv^>^v<v<<v>v>v<^^v>v^vv^<^<>^>^<<^v<>><<<^<^^<<>^^^>v>><v>>^vv>v^v<>^v^<<><v>^^v^^v^^><><v^><^v>>>v>^<^^v^<^<^>><>>><>^v<<<^>^^<vv^<^^<>><<^<<v^>^^^<<>>>^<>^<^^<v^>vv<v^v<^vvv<>v<v<>><^v^^<^v^v^^v><^v^v^<>>v<>vv^v^vv^^<^>^>vv^vvv<>v^vv^v<
<>v>^<<<<>^v<>^>vv^<vvv^v<<v^>><<v^v^vv^^^^v>vvv><v<^^^>v^>^><^<v>>vv><^^>>v^<^v^<v>v>vv><v><<vvvv>vvv<^^v<^v<>><^^v<<v<^<^^v<<^v<>vvv<v<^>>^v>v^^>^>^^>v<<>^^v<v^^>^<v^<>^v^<v><<v>^v><vv^v^<v^v^><<>>vv>v<^>v>v^<>>v>>>^^>><>^^vv^<v^^v<^<^^^v^^><>v><v^<v<v^v<^^>^v<v^^><><<vv<v<><<>><^v>>v<v^>>><vvvv^>v>v<<v<<^<vv^^^vv<vv<v^^^<>><^<v<<v^><<<>^<>v^>><^v^>vv^><^^>^<^vv<v>^<^^<>>v^<^^<><<vv^<^<<<>><v^<<^^<<>v>><<v<>v^><^><^<>><^<v<^^v>vv^><^>><^<v><<^<><>^^>^^<<^^>vv^^^vv<<>^^^<<><^>>^<^<v<^^<>^<^<><><>v^^vv>>^><>vv>v><^^v>^><<^v>vv<^>^^<v><v^^^><^^>vv^<>v^<v<^^vv^v>>>^>^>>v^^^<v^v>>^^v<<v^>>>v<v>^>v>^^^^^>v<vvv<<<vv>^^<^^>v>^v>v<vv^<<>v<>v<^v><<>>^v^vvv^^><<v<v>>^^>v<vvv>^>^<><vv>><<<<v<vv^<<<<^v^<>v<>v<^^v^<>^>><>><vv^<<<^v<^<>v<><<<^>^^^>^>^v><^>v<v^^>v>v><<^^^<>^<<><>v<>v>^<vv<^<vv<^^vv<<>>>vv>^v^^v^>v<^<^>^v^>^v>>^vv<<vv<<<^^<^<>>vv<>><><>^v<^v<^vvv><^v<v^vv><<<v^^^>v^v<<<^<<><^^v^<<>^>><^v>^v^^^^v^v>>>^<v><vv^v^^<><<<v^>>^>^^vv^v><<<>v>>>^>>^>>^<v<<<vv>^^^<v>^^v><v<><v^v<><>><<^^v<><^>
v<><>>>v^<>^>v<<^vv<<v>^<>^v^<v>^<vvv<v^><>^<v><v<>v^v^^<>v<><v^>vv^^vv<^><><^>vv^>^v>>^<^v<vv^>v<<v<><v^>>^^^>>v>><^>vvv><<^><>v>v^v^v>^><<^^vv<v^v>>>^<v>><<v^<v^^^v^<v^^v<v^<<^v^<>>^>^^><v<>><><<<^<<<^v>v>>v^v^v<vvv>^^>vv<^<^^<^^^<>v>^>v<^><<<<><><<v>^><<v>vv>>><<v<<v>^><v^v^vv^^>>^v<>>v^^<<v<<<>>>^vvv>v^><^vv<<^v<>vv><vv^^>>>^<<v^vvv><^^v<>v>^><<>>v>><<v^v^<^<>v<vv<v^>vv^<^^>>^<^><<v^<<^>^>>v^^<<><vvv<vv^<^<>^v^<v>^<vvv>^<>^>v^^^^vv<^^<>vv^>v><^<^v>vv>vv><v<^<^v^vv^<v>v^v^v^v>><^<><>><>>><^v>>>^^>>v<><<>><^<><<vv>v>>v<^^^^<><<>vv<^<<<v<<<v^v^><v^<<>^vv^<>v<^v^v>^>>vv>vv<vv><^<><^<^^>>><>>^<<^^v>>vvv>v<>^^><>v^v<<^^vvv<<>^v<^^^<^v>v<v^v^>^><<^^<v<>^><v^<^<^<>vv>v><<v>^vv<^<<vv<<<>vv^vv<vv>>><v<v^>>^vv>v^^>v<^^<^vv<>><><^v>v<<<^<>>v<^^<<<>>^vv^v><><v><<^<<^v^^^v^v<v<vvv>>^<v<^>>>^<vv>^>>v^^v<^>^<vvvv>>^>vv<><<<v^><^^v>v^>^^>>>^v<>>^^^v<<v><<vvvv^<^vv>^<^>^vv<><^><vvv>v^^^>><>v^v>^vv>>><<>v>>>vv<v<v^<vv>v^>>v^v<^^<v>>^<v^v>^>v^v><<<><<<^>^<v^>v^v^^>>v<<>><<v>vv^<^^><^vv<v<<>v^^v^><<>v>
<<^vv>^^^>>>v^v^<v<>>><v^^>v>v^v<v>><vv<v^^<>vv<<>><^<>>v<><vv<>>^<^><<v<<<>>vvv><v^^v<vvvv<>^^^v<^<v^<<<<><v>vv^<>v>>v><<>v^><<><v^>>^^<^<v^^<v^>>v>^<v^^<vv><<^v<><v<>^vv<><v<^v><<^^^v>^^>v^><v^^<v^^^>^^<<>>^><>^<^v<>^^v<>v<<v^>^^<^v>v>^^v><<><^^<^v^^^v<^<<>^^>v>v^<<>><<<v>>vv>><v<<>v<v^^^>^^v><^^<<^>^v>^<^>v<<^vvvvv<^^v<>^>>^><v^<^^^>^^<vv<v>vv<<<<v^^<>vv>vv><<><v>>v>^>^v^vv^>>^^v^><<>>>^><<^v><<<<<^^>vv>v<^v<><<^>^^^^>v<>>v><<v>><>>vvv^<<v^<v><>^>v>v^<^<^>^^^v<<^vvv<>vvv^<<v^v^><<<^^^<^^^v>^>^v>vv>v^><<v<<<<^^>>v<^v<<<>>v<^v>^vv^v<<v<<v><v><<^<^v^^^vv>^^v^<>^>>^<^<^>><<^><>>v<v^^v<<^<<<>>^v^vv><>>^<v>><v^<^>>v>^^v^<^>>^>vvvv>><<>>><v>v<<^>v<>^<v>^v>>v<>v<vv>^v>^>>>vv^v^^^^^<vv<v<<v>>v>><v<^<<v^>v<v^<<<v<<>^^>^v^^vv^<^<^v><<>^>^v<<^^>>^^v>v<^v>>^v<^^^>^><^>^<><v><>>^>vv>>>v<v^^>^v^v^><^><v^^^<v><><^<>^>v>><<<^^v<><<>v^v<^>vvv<>^>vv>>><><v<<^<^<<^v><v^<v>^><^^^^v^>^^<<v^<<v^v^><>^>^>>v^<^>v<>v<^^<>vvv<><>v<^<^<<^^>^<><^>v><vvv^^v<^>vvvv>v>^vv^^v>^<vv^<^><v>^v><><><v<<v<v^<^v>^^<>^>>>>
^v<<^>>><v><v<^>^^<v<vv^<v<<vvvv>>^v<v<^<>v>><>v>^><vvvv<>^^<<><>>vv<<^^^<<^^^>><^<<>vv<v^>^>vvvvv<^<^v><v^<<<>>vv^^^>^<<<vv^<^<^v>>^v>v<><^v>^^vv>^<><^^^v<><<vvv><<^>^^<<^^>><<^^<^^<><<>>^<<vv<<>^^^^v><>v^>>^>^<>v^>>><>^<<vv^^>^v>^v^>v^>^<^v^^^>>^vv<>>>^vvv>>^vvv>^<vv^v^^>>><v^v^v<><v^v<<>>>^^>^><vvv^^v>vv<>v<vvv>^<<><^>>^<v><^^>^^<<v^^v<>^^>v<>>><^v^^>>^^<v<<<<^><>v<<>><><^>>^<vvv>^v^vv^vv<^<^>^>v>><v>^^<<v>>><>v^v^>vv>v<>v>v^^<<^<^vv<>^<><vv>^^vv><v>>^^^<<<<>vvv^<<<>><><^^^vv^vv^<>^><<<^vvv<<><>^v<>>^<>>v^^>v<<<^vv<<><v^>>vv<vvv<>vv><^<>^v<>>^v>^<^^vvv<><<>^>v^^^^<>v^^<vv^^^<^^v>>>vvv^^^<^>>^v>>v><v>^v>^<^<^v^<>^^v^v^^^v^<^^<^>>v<^<>^>v^v>>>><>vv^vvv>^>v<<<<><v><>><>>><<v^<vvv>^^vv><^^>>^vv^v^<^^<vvv<<vv^^>^>v^>>>^^>v<>^>><><v<v<>^^v<^vv<vvv^vv^><^<^^>><<<^><v<^>>>v>>>^v<^<>>v>^v^<<v<>^v^v>v>v>>><v>v>v<><<><^v<^v^^>>^vv^^^v>>^<^v><^><<^^<v<>^^^<<v>vv<<vv>^><v<v>>>v^>^><vv^v<v^>^^<vv>>>^vv<v<^v<v<vv><^<vv<<>v<>>v^v^<><>v^^>v>><^>v>v^<>v>^v^<<>>>^^v^^>v<<>v>^>^v^vvv^v<v<^^>v<^^<vv>v>^
^^v^vv^>>>v<<^^<v<^>^^<vv>><^>>v>>v<^>>^^>vvv>^>>vvvv^<<v<<^>>^^^^^^v^<^<>v<^>^<^<v>>>v>>^>^vv<>v<<v^^<v><^><^<^v^<^v>vv^^<>>>>v^<<vv^^>v<^^v<<^^<>^<vvv>>^>v><v<^v<<vv^<v^v<v<vv^<>>>v^^v<<vvvv>v^v><<><^^v^><v^<>>v>>>^<v<vv^v^<v<^<<<^<^v>^<<^<v>^<v^vv><>>v<<<^>^v>^vv^v^v<vvvvvvv^<^^v^v><>^<v><^vv^>>v>>>v<v^^>><^><<>>>^>^v^v>>^<v>v^<>v><^<v^^^<^^<v>v>vv>>>>v<v>^v>^>>>v>^<><<vv<^>^v<v>vv>^><^<>^>^>^>^>>^<vv<v<>>vv>>vv^<<<<<<vv>><><^>><>^><>v>v<^^<v^>vv<>>^^<^<<<<<^v><>^v^>^>>v<v^><^>^v>vv>>^><<v^^^><><^v<v>vv^v>^>vv^v<<>v<<>^<<>^<>>><^^<v>v<v<>vvvv<<<<v<>^<<<<>><vv>^^<^^^v>^^v<v<^>v<<>^vv^^<>v<^v<<>><^<>>>^<v^>>>>^<><<<vv^v<<v<>>>v^^v<v>^^^v^^<><v^vv><^>><^^<<^<^<<><>^vv>>v<>^v<<>vvv^vvv^v<v^<>>><<<v<><<>vvv>^^v^<>v<vv<v><v><>vv<^>^vv>^v^^v^>>>^>^<vv<v>^><>>v<^^><<<<^vv^><<<^>v^<<^>v>><<<>^>>^><<><^v<><^vv^vv<>v^v<v>^>>^vvv^^<>vvv<v^v^<vv^^>>>v<v<v>><v^>vv>^^v<v^^^v<<v<^>^vvvv^><vvv<><^<>^v<<^>^<^><>^^v^^>v<>>v^<>>><><<^^vv>vvv^<>^<v<>><^^^<><<^^v^<v<<^><v<^vv<>^>>^><^v><>v<><<<>v>v^v>^<<
vvv<><v<<v>vv^<><^<>^v^v<<v<^^^^^<>^vv>^<<>^v><^v<^v>^^v^^^^^>v<^<>^>vv^^><<><^<vvv<><v<v^>vvv^v>^^><>^><vv<^<>vv^<<>>>^<v^<vv^v<v<<>>^><v<<^v<^<vv^>^>^><>^<v^^<v^vv<v^>^<v><>^>^v^<^<vv>v>v>>^<^>><^^>^v>^v>^^v<<^^^><>v^^<<<^>^<v<vvv>v<>>^^<v>>><>v>>^>>vv^v>v>>>v<>>^^^<v>v^><^<<vv<>>^><<^v><^>v^<<vv<v^^^>vvvv<^^v<v>>v><vv<>>v^<^<v^>v>v<>^vvv>v<vvv<^^v<v^<<<^v<v<>^^^^<<^<>^><^^^>>^<>v<<>>><^<><<<<vvv><<v<^^vv<vvvv<^^^>^>>><>><<v>^vv>v>^<>>v<v^<vv><>>^vvvv><v<vv<v^>v^^^^<><>^^v>v^vv^<>^<>^><>><v^>><<vv^v^>^>vv<>vvv<<v^^>^>^^<^v>^<v><v>>vv^^v>v>v^>>^^v>^<v<>^>><vv^v^v^<<vvv<>><>><vv>v<>v^>^><<<<>^v^^<^^vv>^<^v>^v>v<<^^<^^^>^^^><^v^<v<^<v<v<^>>v><vv^<v<vv^^<>^^vvv^v^vv^^v^<vv<<<<>>v>^^>v^^<<v<><<^>>^><^v>v^>v<<<^>^v^^^^<^<>^<>>>^^<<^>v<^^<<><vv^^><v^v<^<<>><<<<^<>>^^^>^^>>>vvv^><^<^^^v>vv^^>v<v<^>v^<v><>>^<>>v<^v<vv^<><v><^v<^vv<<^<<^v><^v<v^><<v<^v>><^^v^<^v<^^^<v>><v<^<<^<<^>>>>v<>>><><v>^<v>>>v><<<>>v<<>>>>v>v>^>^>^<<>><>>>v^<vv>vv<^v<<<>v<>>^>^^>>v><^>>>^vvv<v^^vv^vvv<>^v^vv><>vvv^v><^^
v^v<^<>^^^vv<>^vv^><>^><v^^v>v>>^>^>vv><><v^<^v><<^><v>><v><v<^^<^v<><v<>v>^><>v^<v<<^^^>>^^^>><<v^v<><<<<<^v^^<>vv<>^vv<v>>^>^v^<^v^^vv<vv^<<<v^vv<^v^<^<vv<<>>><<>>vvv<v^^<^^^>><><^^^>>^><<<<<<v<<><v<^>^<>><v<<>><<v^^vv^<^vvv<>^>v<<vvv>^>v<>><^^<>v>^^<v<>v<><>vv><v^^>^<^<<v>>v^>vvv>>vv^>>>>^<<>vv<v<<>^>v^^>v^v<v^<<>>>v^v>v>v>vv<<<v><^^v^^<^>^>><v<^^>v>v><<><v<>>^v^>^>^<<vvv^><^<<^<>>^<<>vv>>><>>>^<^<v<>vv^>^v><>v^>vv^><<v<<>^v^>>^^^v^>v^^v>><^<vvv><^^v^><><>>^<vv><<v^<>^<>>>>^v^^^>^>^>vv>^<><<<>^<^v>^v^v^v>>>>^v>>v<^vv>^^<>v^v<vv><v>vvvv>>>^<^<>>>^v<v>>>^^^<vvv>>><<^vv<<<>vvv^>^vv>>><><vv<^^<<>^v>><v<v>>^vvv^v^>^v<<^^<^^v^v<^>v^^v^v^<^v<>v<>^<>>^v<^^^^^<<>^<v^<>>^v<^<^>vv>>^>vvv>^^<v><><<v>v>><>>><^v^<>^v^vvv><v^vv^vv^<^^>>v>^^<<^v<<<^^v<^>>vvv>v<<<>>><>>^v<>>>^^^>>v<vv^<<><v>v><vv<<><>^^vv^^<<^>>v^>^^^>>>>^v^><v^<^^v^^^<^v>vv<<v>v>>^v^>^^>><vvv<<v^><^v><>^<<^>>^><v<><<v^>v^vv<v^^^^^<>><^^<v<^<>vv<^<^v>^^v><><v>^vv>vv<^>>^>v<^vv<vv<<<<>^<v>^<<^>^^^<<>^><>v^^<^>>^^v<v^><^>v><>^<>v^^<<^
>^><^vv<v>v^^^<>>^v^>v>^v^<<v^>>>v<v<vv^<<><<^^v<<^v>^v<^<<<<v<<<>v>^>>^>>^>><^>>v><<v^<>v><<v>^><>v<<v>^v>^>^>^^^>>v>v>^^<^vvv^<^^><v^><>><>v^>>>^>^^>>><>^v<v^<v<^^^v>v>v<v>>vv<v^^<^v>^^^><>^>vvv^^<>^<^<vvv^v<<>>vv><v<vv>vv><<<<>v^^>vv>v<^><>v>>^>>^<v<v>>^^v>^v<^>^v>>^v><>^v>>v^v^^><<><^>><v<>^v^vvv<<v>v><<^>v<>^<<<<vv^>><v^^<<vv<^^>v>><v<><^>^>>v^<vv>>vv^>^<^<^>^v^<<^^^v>>><vvv>^><vv^vv>v<v>^^vvvvv<>^^^^v^v>vv^^vv<^<><v<<v>><>v^v<^>^vvv^<^<^^>v^<^v>><>v^^^>v^^<>>>^<^<>^^<<^v^>v<v<v^>>^^^>>^^<<<>^>^<^^>><>^<^v><v<>><>^>^<>^^^<><v><^<^v^^>^vv^^<>v^>^>v><<v<>^^v^<<^^<^v^v<<v<^>><>vv^v>^<>^^v>>vvvv<v>vv<<><>^^^><^^v^><vvv>v<<vv<<^<vvvvvv>v>vv><^<>v>>>^v<>><^><>v^<>v>^>v^<^^^<v^<v>v^<v<<v<v<^<v<^>^^<<<^v><v<<>^<<>v>^^^v<>v<^v><^^^>>>>v>vv<<^>v<>^v<^v^><<^^><^>v^v^^v<<<><<v>vv<v>^<vv^>v^^^>^v<><^<v^>v^<vv^<^<<<^v<v^^>>>v<v<^><^<><v><<^>^v<<<>>>v>v>^<<><^<v>><v><vv><><<>v<v>v^>^<v^<^^^^^>>>^v<^^>vv>vv<<^^>><<v^><^v>v>v<<<>>v>v<<<vv<vv>^v>^^vv>v^v>v>^^^>^><><v>vv>v^>>^^>v^^>>v>><>v>v><^<<vvv
<v>v>><^<v^^><^<vv^>>v>v^<vv<<vv><vv^>>v<^><v^<v^^<><<<v><v<v^<^>v^>>^^>v^v>^>v>v>>vvvvv>>>^v<><<^^vv>^>v<v^vv>v>v<vv>><v<vv<vv>vvv<<<^<vv<<^<>v<^>v>v>^>v>^v>v<^<<v>v><<>^vv<^v<^^>>^^>v^>vvv>>v^>^^v^><^vv^^v>>>^^v>^^^vv^>>v>^<>v<>v><<>^v<^v^<>vv<^><v^>v^<v<>>>v><v<v<v^vv><^>><v>^<><<^<^^^<<v><<<vv<<><><>v>>><v<<<^<><^>v>^><<<^>v<><<>v<><^><<^^>vv<<vv^>v<>^^v>>v<vv<^^<>v>v>^<^^v>^^>><^^^^vv<^v^^>v>^^>><<v><^v<>vv<^vv>^<<^<^^<v>^^<>v<<<v<<^^^v<v>v^>>><<v>v<<v^>^<><v^>^><<<^^<^^^v<v^<<<>>>^^^<^<vv<^^v<^^<vv>^>v<><^<^><>>>><^^v>^v<<^>><^vv>^<><><^^vvv<>^<>v><>v^v^^v>v>v><<<^^<<^<>><<v<<>>v^^>^^<^v<^^>^^^<^v^^<><^>vv<<^<v^vvvv<^v<vv>^<<><^^^v^<>v>v<^v<v<v>v>^>^<<>>^v^><>v^^>^>>>>>v>v>^<^^>>>^^<<>v^vv<^^^v^^v^v<v>v<<^>v^<>>v^><>>^^>v>vv>>^<v<v^vv><vv><vv><<<<v^vv<^>^<v^>^>><^^><v^>><v^>>^v^>><v^>>^^>v<>><>>>^^vv^^>^>><><<<v>>>^^^^>^vv^^<^v^^^^><v^^^^<<<v<^<>vv<<^>vv<<^>^v<vv>vv<^<><<<v^^v>>>>><v^<^v>v><^<vv<<>>v>v<v<<v>v^><>^>^^<<v><><^>v>v>^^<^^^^<v^^<<><<>>^^><v<^v<v><^<>>>>>v^v><<>v>^vv<^
v<vvv>>v^<^<<^^>>^>v^<<v>v><^v>v^vv<<<v<<<<^>^^^^^vvv><v^<>v>v<^><^<<>^^v<vv>>vv><^v^vv^v^v^^v>>v^<^^<<<^>vv<v<v^><^<<v<^v<><<^>><^^<>>>><v>^^^^^v<>v>vv>^<v<>^^vv<<>^<>^<>>^<>><vv><^<v<v^v>><^>>v<<vvv>>vv<v^>^<^^<>>v^<<^<<>^>vv^>^v<v>v>v^<v<^^><>v<^^v<^^^^<^<^v<v^v^^^<>vv^<<<^^>><<>^>v<<v<v^v>^v^<v>v>>^v<v><>^^<v<<v^>vv^<^>v<^v>v><v<<>^vv<>><v>><>^>v<<^vv><^>>><>vv^^^^<v<<<v>^<>><<<^v>^v><v<^<>^><>>>>v<>vvvv^v>^^<^<^>v>v<><>^v^>vv>^>^><v^>>><vv<^<vv><>>>v<>v^<vv^>v<^><>^<>^^<>^>><v^vv<v><^^<>v^>>vv^<<vv^v>^>><v<<v>v>^^^>v<<<>vv>v<^v<<<><v<^vv<vvv^v<^^^<<<vv<^>><<^>v>v^v^^v^v>v<<<<^>>v<>v^v><^<^^v<vv<v<>>vvvv^>>>v>v>v><vv<<<<>>^vv<^<><v<^^<>><^<>^^>vv^><v^vv<>>v^^v><<^><^^<<^^^^v<^v<v<<v^>vv<<><v<>>^^v>v>>^^^<<>v<v<<>^<v<^^>^^^<<^^<v>><^<^<><^<^v^>vv<>v>^<>^<^v>>v<>>>>v>vvvv^<<<^^><>^v^><>^>>^<<>^v^>>v^v>^^<v>v<<^>>^>^>v>^<^>><<^>^v>>>v<>vv<^>^^v><<^v<<^v>vvv>^v>^<v^^<^v>^<^^<v<<v<^<<^^<v^>>v<>^<<^^<><>v<v><<<^>><<<^<^v><><>v>^>^v>v<^vv<v^>^^^^>^^><^v>^><<v^<v^vv><^<v<<>^^<<>^<>v^>><^>^
>v^vv>vv<^^^^><v<v<^<^v^><v><vvv^><v<v^v><v^^^v^>vv^^<><v<<^vvv<<>^^vv>vv>^^<>>>vv^>^>>>><<<>v>>v^<<^^^v><^^v>vv>^^>v>>v<^<^<v^<>v^^^v^^^vv><>^^>v^<^^<<<^^^>>>v^^v<>>^v>vv><^^vvv<>v^<<^<^<<^<v><<<<v^>><v>^>><^<>>v^><v<^>vvv<<<v>v<v>v<v^>vvv^><v<^^>><^<^>>>>>^^v><v>^^vv^<v<^<>v>^<<v>^v^v^>^>v<>^<^>v>^^>v^^>>^<<<^v<>^<><^v<>v^<^v^v<^v>>>>>^<>^^<v<v><>^<<v>><><<<v<>^^>^<>><^<^>^<<<<>^>^v><<>>><<>>^><<^v<^v><>><v^<<v<>>><>^>v^^<v^^>^<<<>>^^>^><>v<<v<>v<^^>>v><>^<^^>vv<v>^<>v>><^vvv^^v^^>^<<>>^<v<^>^^v^v<v^>v<^>^v^vv^^vv<v^>^^v<^<v<v^<<>>>v<^><>><^vv<<v<^<^v><><^<<v>>^^^v>v<><vv^<v<^v>>>^^^^^^v>><v^>v<<>v<<<><vvv^^v^v<v^v>>^<v<<>>v>^<<^^>^v<v^<^v>^><><<^vv^<<<>>^>>^^^^vvvv<>^>v<v^>^vv<vv^<<v^vv<v>>>v<<v><vv<<>^><><<v><>v<<>v>>^^<<^<>^<^^>v^<^v^^^^^^^v^^>v<>>^<><<^^>vv^^v<^^v<^^>>v<^^<v<^>><>^v<^<v<^<^<^<^>^<^v>>v<>><<^<>^^>v^>>vvv^<^><<^v>^v^^vvvv>^>v<<v<v>>><v>^^>>>>^><^<v>v^v^>v<v>^^^vv^>vv>v<^v<><^<vvv>><>>>vv><v><v<v^><><^><<vv>^<>v>><><vv><><^>^<v^v>^<v^v<^>vvvv<^^>^>>v<<^<^<^<v^v<^^<v
>vvv^vv>v<>>v<<<v<<<>><v<v<^^v^<>^v^v<>>>^>^>^v<^<<<v<^>>vvv<vv>v>>^<v>v><v>^v>^v<><v^v<^^vvv^v<>><^>v<>v^<<<>v>v><<^>vv^><^v><<^^^><^<^>>>>v<>vvv^<v><>>>>>^<>^^>vv<>><v^^<^<<>>><<>^<<^>>v^v>><v<v^<^^<><><^v<vvv^v>>>^^^>><<<<vvv^v>><>>>v>^v>>^><>^<<v<^>^<>>vv<v^v<>^>^v^<>^<^^><<<v>>^v<v>><v<v^>^vv>>>^vv^v<>^^^v<v<<^>vvv^>><^v>^><^>>v<^^v>^<^<^^v<^^<<^><v<v><^^<^<v^>vv<<><v^^<<^>v^>v<><>^^v^<vv^^<>v^<>>^><vv>><^><^^>v>>^<^<>^v>><^v^v>>>v^>>>^v<^<>>^vv^<>><>v<^<^<^^><^<vv^<^>><>vv^v<<>v>v>vv>^<^<^>^<>^^<>v><><><^<v><v^^<^v^<<>^vv<v^v>^>v><>^>^v<^>>vv><^<^>>^><^<>><^v^^v><^^>v<>>vv<v<v^><^^<v<>>^<>^<v<<^<v>vv<^^^>>><<^<v>><v^^>vv>^<^^<^v^<^<v<<>v^^^>vvvv^><>>v<v><>^>^v<<><<>^<vv<<<v<>><v^>vvv^^<<^>>>vvvv>v><^<>v<<^v^^>^>v^^vvv^v<><>><v<v>><^><^v>^^^<<^<v<<^^<^v<^<<vvv>>>>vv><<v>^>vvv^^^<>v^v^vvv<>>^^>>^^<<^^<<<<^^^>>>>^<^<<^><>^<^v^>><^<^><v<^<>>vvvv><>><^^>v^v<v^>>v<v>v>^^v^<>><><<v>>^v^^v<<^>^v^^v<<>^>^>>^<<v><><><<>^vvv<>>^v>^^>^<>v<v<^><^^<>^>^^v<>><<^>vv<v<^^<><v<^<<v>><><v<<>^>^^^>^
><>>v<<<<<^<v<vv^>^>v<v^<><^>^<>><^>^^>vvvv<>^>^^^vv>v<>v<<>>^<>v^v^>^vvvvv<^v<^^^<vv>^v>^>vv<^^>^>><>v<v>>v<v<><<vv^^>^^v<^>v>^<<<v^vv<<v^>v^>>^^^>>^<^>v^<^<<>>^>>v>vv>^<v<>^><v^<>>><<>><^>^vv^>^v^v^<<<<>^v><^>><>>v^^^>v<<>>^<>^>^<^v^v<^vv>^>v<v^>v<v>^^>>vv^^^<<v^<v^<^v<^v>^>>v^^vv<v^>>^^^v><^^<><^v<v><v<<^<^>^<><<<v>v>>v>v^><v>^><^><>><>>>vvv><>^^>>>><<<^v^><>>v><^<>>><v<v>><^v^^^>vv><v>vv<<^^v^>^^^>>>>v<>><^>><^<^vv^><<^v^v<><vv<^v^<v^<^<^<><<<v<^<vv<<>^<>v^>><v^vvvvv^>^v^^<^^^><<vv^<v^v^>^vv^<<vv^^v^<<<<><<vvvv<>^>><>^>><<>><v^^<<>^><>v^<^><^><vv^^>v^v>vv>^^vv^^^<<v>>>v<v<^>^>^v<^vv^<>>^v<^^<v^>^<^<<>^>^^vvvv<<v>v<<^<<>v^^^^^v<vv>v>^v>v>^^>v^v^<^v<v>^vv^v^^v>^<vv<^>vv^v>>^>^<^>>v>v<<^^^>^^v^><v<>^vv>^vv^vv>v^><<><^^<<^v>^^v>>>>>^v<v<^v<^v^>v<<v<<>v><v<><vvv<<vv>vv<^v>v<^v<>><<^^vvv<>^>v^>vvv<v^v<v<v<>>v<>><<<v^^<^^<<<<^v><<>^v^<v><v<^^vv^^^>><<><^v^v<<<<v>^<>^<>><^v><^><^^><>><<^<<^>vv<<<v>v<^^>^^^<v<>>v<v<^v<^>>v><v<^>v^^>>v<^v^>>^<<>vv^^^<v>>>^v<^v>>^^^^vvv><^<^^<^<v^>v<v>^>v<^^<
^^v<v<<vv<^v<^^^^>^<><v<^>><>v>v<>>v>^<>>^^<<vv<^v^^^^<^<>^^^^v<>v<<^^^v><^<><v<vv>>^<^^<v<^>><<vv^<v<<><><v^v^^><^<<><^v>>v>^v<^<v^^v><<>>^<vvv<vv^^<vvvv>v^<<><>>><>^<^v^v^v^^v^^<<^<<>><vv>v^v^v>v^<v><v^>>>^v<^v><>v>>>v><v^v^>><>>v><>^<v>>^v>^<v<v><<<<^^<<^>>^^>>^v<^>^><>><^<<v<^>>v^<><<><v><^<^^^<vvv<^<<<<>>^vv<<vvvvv^>><>v^<v>vv>>v^>^^v<v^<v^vvv<>^>v<><^>^><<<<>>^>^<^v^^^>v><^<>v><<<v^<>>v^<v>^<<v<<<v<>^><v<vv>^>vv>v<^^><^<<vv^>><^<vv<^v>>^>^<<v^^<><v><>^>vv^^<vvv^v^^<><^^>>v<<<>v^>^^<<^<><v>>>^^^><v^v><>v<>v^v<vvv^v<<vvv^>vv<<>v<>v^^<vv<^<>>><vv>>v^<>^^<<>v>v^<>^<<v^<>v^>v^<>vv<>^<<<>>^vvv>>v<v<>v<<><<<><v<<v^v>>v>>^v^^v<vv<v>>v<vv<v<>v^^^>v^^v^^><^<^<^<v^<>v^^<v^><^vv^^vv^<vvvv^^^v<<^>v>><^v^<^vvv<v^>^<v<vv^v^v>^><^^v<>^v^>v<^<<^<^v><^vvv^>v<v<>^v<^^v<>>v<^v>^vv>v<<<v>>^>v^<<<>^<^^^v>^^vv>v<vv>>v^><>v<<>><vvv^v^^vv<<^^v>><<v>v^<>><<<^^>vv^v<>v^^v>vvv>^><>v<<<>><>^^><<<>^>>^^<^v^^v><^>^><<v^v<^^>>v>^v>v>>^<<^v><<>v^^v^<^><><<^^>^^<^<^^<<<>^vv^>><<v<>^^vv^>v>v><^^>v><^<^^>^^^<>^v>vv
vvv^>>vv^v>>>>^<<v>><vv^^vv<^^^><^><>v<>^>v<>^v>>v<v<>vv^v<><<^^^v^^<<^<<><^^><^<vv><>^^v^>^>^>^<<><<><<v<>v><v<vv>>>>v<^v<v><>vvvvvv^<v^v><^<><v<^<^^vv>>v><<^<v<<>^>>><v<^>^^^^>v<<<<>>v<<vv^<vvv><>^<>v<>><<>v^^<^>^>v><v>^<<v<^v>v^^<v<<v>vv>^<<>>^<^>^^<^>v<>>v>v><^vvv<^<>v>>^^>^v>vv>^v^<^^<<>^<^<v<v>>^><><><>v^><>v^^vvvv^<^^^^<^>^v>^vv>>v<<v^^<>><>vv^>^v><v<>^<v^<>v<<^v<<v^v<>vv<v>v<>^vv^v>^<^>v^v^<<<<v^>vv^v<<v^><^><^^v>^v<^^<>v<>>>^^<<<>>>v>^^^^^<>v^vv>^<^<^^^vv>^<v^>v>>>>^>^<^^vv^^vv^v>vvv>>>><<v<>vv^vv<<><vvv^>>>^^^v><v<><v><^<>^v>^>>^^>vv<<v><>>v^^^^<v>v<><v^^^^<>^>^v>^><><^>>v<^>^vv>^>v^<^^>v>^<v><vv>^^v><<><^><>>v<^><v^^v<^>^^<<^><><v><>>>^<vv<<>v^<v^^>>vv>>v<^>v>>>><^><^>v^<v<^<v^>^><v>^v>>^>>v^vv<<vv<<>vv^^>>v>vv<><>vv^<<vv>^^^^v^^<vv><vv>>><^^<^^<><^v<v>v>^^^v^>v<><vv>^<>^>v>>^v>v^>v<<<<vv^>v>^v<v>v>>vv<<^><><<<v^^>^><<<^v<<^^^<<v^v<v^v>>v<v>vv^><v^^<>^^<^>^vvv>v>>vvv<>v>>^>><>^v>v<^<^^><^>v<>v>v<<>^^^>>^>>>^^v^<<^>^^<^>>>^^^^^>^v<^v>^^<<v>>^<^<v>v<>>>^v^<v>v^^<^><<v^v^v^<<v>
	</pre>
</details>
