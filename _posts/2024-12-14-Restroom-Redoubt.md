---
title: Restroom Redoubt
description: Advent of Code 2024 [Day 14]
layout: default
lang: en
tag: aoc24
prefetch:
  - adventofcode.com
  - deno.com
---

One of The Historians needs to use the bathroom; fortunately, you know there's a bathroom near an unvisited location on their list, and so you're all quickly teleported directly to the lobby of Easter Bunny Headquarters.

Unfortunately, EBHQ seems to have "improved" bathroom security **again** after your last [visit](https://adventofcode.com/2016/day/2). The area outside the bathroom is swarming with robots!

To get The Historian safely to the bathroom, you'll need a way to predict where the robots will be in the future. Fortunately, they all seem to be moving on the tile floor in predictable **straight lines**.

You make a list (your puzzle input) of all of the robots' current **positions** (`p`) and **velocities** (`v`), one robot per line. For example:

```ts
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
```

Each robot's position is given as `p=x,y` where `x` represents the number of tiles the robot is from the left wall and `y` represents the number of tiles from the top wall (when viewed from above). So, a position of `p=0,0` means the robot is all the way in the top-left corner.

Each robot's velocity is given as `v=x,y` where `x` and `y` are given in **tiles per second**. Positive `x` means the robot is moving to the **right**, and positive `y` means the robot is moving **down**. So, a velocity of `v=1,-2` means that each second, the robot moves `1` tile to the right and `2` tiles up.

The robots outside the actual bathroom are in a space which is `101` tiles wide and `103` tiles tall (when viewed from above). However, in this example, the robots are in a space which is only `11` tiles wide and `7` tiles tall.

The robots are good at navigating over/under each other (due to a combination of springs, extendable legs, and quadcopters), so they can share the same tile and don't interact with each other. Visually, the number of robots on each tile in this example looks like this:

```
1.12.......
...........
...........
......11.11
1.1........
.........1.
.......1...
```

These robots have a unique feature for maximum bathroom security: they can **teleport**. When a robot would run into an edge of the space they're in, they instead **teleport to the other side**, effectively wrapping around the edges. Here is what robot `p=2,4 v=2,-3` does for the first few seconds:

```
Initial state:
...........
...........
...........
...........
..1........
...........
...........

After 1 second:
...........
....1......
...........
...........
...........
...........
...........

After 2 seconds:
...........
...........
...........
...........
...........
......1....
...........

After 3 seconds:
...........
...........
........1..
...........
...........
...........
...........

After 4 seconds:
...........
...........
...........
...........
...........
...........
..........1

After 5 seconds:
...........
...........
...........
.1.........
...........
...........
...........
```

The Historian can't wait much longer, so you don't have to simulate the robots for very long. Where will the robots be after `100` seconds?

In the above example, the number of robots on each tile after 100 seconds has elapsed looks like this:

```
......2..1.
...........
1..........
.11........
.....1.....
...12......
.1....1....
```

To determine the safest area, count the **number of robots in each quadrant** after 100 seconds. Robots that are exactly in the middle (horizontally or vertically) don't count as being in any quadrant, so the only relevant robots are:

```
..... 2..1.
..... .....
1.... .....
           
..... .....
...12 .....
.1... 1....
```

In this example, the quadrants contain `1`, `3`, `4`, and `1` robot. Multiplying these together gives a total **safety factor** of `12`.

Predict the motion of the robots in your list within a space which is `101` tiles wide and `103` tiles tall. **What will the safety factor be after exactly 100 seconds have elapsed?**

```ts
class Robot {
  x: number;
  y: number;
  vx: number;
  vy: number;

  constructor(pos: string, vel: string) {
    const [x, y] = pos.split(",").map(Number);
    const [vx, vy] = vel.split(",").map(Number);
    this.x = x;
    this.y = y;
    this.vx = vx;
    this.vy = vy;
  }

  move(width: number, height: number, seconds: number) {
    this.x = ((this.x + this.vx * seconds) % width + width) % width;
    this.y = ((this.y + this.vy * seconds) % height + height) % height;
  }
}

function parseInput(input: string): Robot[] {
  return input.trim().split("\n").map(line => {
    const [pos, vel] = line.split(" ");
    return new Robot(
      pos.replace("p=", ""),
      vel.replace("v=", "")
    );
  });
}

function simulate(robots: Robot[], width: number, height: number, seconds: number) {
  robots.forEach(robot => robot.move(width, height, seconds));
}

function countQuadrants(robots: Robot[], width: number, height: number): number {
  const midX = Math.floor(width / 2);
  const midY = Math.floor(height / 2);
  const quadrants = [0, 0, 0, 0];

  robots.forEach(robot => {
    if (robot.x === midX || robot.y === midY) {
      return;
    }

    if (robot.x < midX && robot.y < midY) {
      quadrants[0]++;
    } else if (robot.x >= midX && robot.y < midY) {
      quadrants[1]++;
    } else if (robot.x < midX && robot.y >= midY) {
      quadrants[2]++;
    } else {
      quadrants[3]++;
    }
  });

  return quadrants.reduce((acc, count) => acc * count, 1);
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const robots = parseInput(input);
  const width = 101;
  const height = 103;

  simulate(robots, width, height, 100);
  const safetyFactor = countQuadrants(robots, width, height);
  console.log(`Safety factor: ${safetyFactor}`);
}

main().catch((err) => console.error(err));
```

During the bathroom break, someone notices that these robots seem awfully similar to ones built and used at the North Pole. If they're the same type of robots, they should have a hard-coded Easter egg: very rarely, most of the robots should arrange themselves into **a picture of a Christmas tree**.

**What is the fewest number of seconds that must elapse for the robots to display the Easter egg?**

```ts
class Robot {
  x: number;
  y: number;
  vx: number;
  vy: number;

  constructor(pos: string, vel: string) {
    const [x, y] = pos.split(",").map(Number);
    const [vx, vy] = vel.split(",").map(Number);
    this.x = x;
    this.y = y;
    this.vx = vx;
    this.vy = vy;
  }

  move(width: number, height: number, seconds: number) {
    this.x = ((this.x + this.vx * seconds) % width + width) % width;
    this.y = ((this.y + this.vy * seconds) % height + height) % height;
  }

  getPosition(): [number, number] {
    return [this.x, this.y];
  }
}

function parseInput(input: string): Robot[] {
  return input.trim().split("\n").map(line => {
    const [pos, vel] = line.split(" ");
    return new Robot(
      pos.replace("p=", ""),
      vel.replace("v=", "")
    );
  });
}

function isChristmasTreePattern(positions: Set<string>, width: number, height: number): boolean {
  const grid = Array.from({ length: height }, () => 
    Array.from({ length: width }, () => false)
  );
  
  for (const pos of positions) {
    const [x, y] = pos.split(',').map(Number);
    grid[y][x] = true;
  }

  let density = 0;
  for (let y = 0; y < height - 1; y++) {
    for (let x = 0; x < width - 1; x++) {
      grid[y][x] &&
      (
        (y > 0 && grid[y-1][x]) ||
        (y < height - 1 && grid[y+1][x]) ||
        (x > 0 && grid[y][x-1]) ||
        (x < width - 1 && grid[y][x+1]) ||
        (x < width - 1 && y < height - 1 && grid[y+1][x+1]) ||
        (x > 0 && y < height - 1 && grid[y+1][x-1]) ||
        (x < width - 1 && y > 0 && grid[y-1][x+1]) ||
        (x > 0 && y > 0 && grid[y-1][x-1])
      ) &&
      density++;
    }
  }

  if (density >= 0.6 * positions.size) {
    return true;
  } else {
    return false;
  }
}

async function findChristmasTree(robots: Robot[], width: number, height: number, maxSeconds: number): Promise<number> {
  for (let seconds = 1; seconds <= maxSeconds; seconds++) {
    robots.forEach(robot => robot.move(width, height, 1));

    const positions = new Set(
      robots.map(robot => robot.getPosition().join(','))
    );

    if (isChristmasTreePattern(positions, width, height)) {
      const matrix = Array.from({ length: height }, () => 
        Array.from({ length: width }, () => ' ')
      );
      positions.forEach(pos => {
        const [x, y] = pos.split(',').map(Number);
        matrix[y][x] = 'X';
      });
      console.log(matrix.map(row => row.join('')).join('\n'));
      return seconds;
    }
  }

  return -1;
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const robots = parseInput(input);
  const result = await findChristmasTree(robots, 101, 103, 10000);
  
  if (result === -1) {
    console.log("Christmas tree pattern not found");
  } else {
    console.log(`Christmas tree appears after ${result} seconds`);
  }
}

main().catch((err) => console.error(err));
```

## Links

- [If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2024/day/14)
- [Solve Advent of Code 2024 with Deno and Win Prizes!](https://deno.com/blog/advent-of-code-2024)

<details>
	<summary>Click to show the input</summary>
	<pre>
p=10,22 v=50,-21
p=35,60 v=-8,52
p=79,50 v=82,9
p=52,1 v=-28,-55
p=85,67 v=-51,-18
p=72,34 v=82,13
p=57,0 v=67,75
p=7,52 v=85,67
p=81,19 v=-73,-21
p=46,32 v=-82,29
p=68,101 v=-73,30
p=23,42 v=-36,-63
p=25,93 v=40,-16
p=30,93 v=-76,-55
p=22,58 v=72,-9
p=25,84 v=18,-87
p=38,80 v=39,-42
p=35,71 v=-55,-53
p=88,60 v=-32,-29
p=86,86 v=-16,98
p=30,102 v=66,35
p=64,88 v=54,75
p=24,4 v=18,93
p=4,6 v=4,-93
p=57,70 v=-7,-42
p=23,38 v=-37,56
p=75,84 v=-89,-21
p=10,83 v=30,-33
p=77,66 v=-73,-2
p=4,15 v=84,2
p=44,34 v=-48,31
p=15,78 v=-4,50
p=31,60 v=-26,-66
p=62,50 v=34,-94
p=98,25 v=37,49
p=74,81 v=1,59
p=94,81 v=-38,-60
p=70,20 v=-54,-60
p=89,51 v=-85,-32
p=49,51 v=-48,76
p=53,66 v=33,83
p=61,29 v=-47,-63
p=49,43 v=94,-7
p=17,23 v=-49,-30
p=15,55 v=24,-63
p=82,3 v=-13,77
p=77,44 v=54,20
p=8,43 v=30,-76
p=70,72 v=-34,97
p=58,10 v=96,45
p=67,95 v=-20,39
p=92,81 v=-83,-26
p=44,98 v=64,-41
p=1,99 v=71,73
p=81,25 v=-59,44
p=16,31 v=91,-32
p=48,44 v=-41,-63
p=88,11 v=-72,37
p=31,89 v=-77,-75
p=42,41 v=-40,33
p=61,41 v=-73,-16
p=96,7 v=-71,-37
p=16,35 v=-43,-43
p=80,79 v=-59,70
p=20,35 v=-70,-5
p=41,47 v=-50,-70
p=86,37 v=11,-72
p=61,97 v=54,-73
p=82,4 v=90,-3
p=24,35 v=38,80
p=52,93 v=-81,2
p=90,37 v=90,-43
p=18,28 v=31,11
p=33,58 v=-2,56
p=100,102 v=-32,19
p=97,51 v=-45,-15
p=68,17 v=76,-93
p=100,13 v=83,-46
p=88,8 v=-12,75
p=71,14 v=-6,-30
p=83,10 v=15,-37
p=81,84 v=35,12
p=98,90 v=-52,43
p=91,60 v=65,82
p=3,72 v=-24,-40
p=84,49 v=-58,22
p=82,75 v=-99,27
p=4,73 v=46,79
p=15,32 v=51,24
p=6,32 v=-44,-81
p=21,101 v=-36,28
p=21,100 v=51,-73
p=50,32 v=-92,-98
p=1,58 v=-24,-2
p=52,96 v=-95,-64
p=83,28 v=-60,11
p=85,101 v=-59,-26
p=19,2 v=-31,-97
p=50,32 v=-68,31
p=13,66 v=43,36
p=100,15 v=-30,-14
p=84,78 v=-58,88
p=65,50 v=-53,-21
p=50,75 v=26,61
p=89,60 v=-93,47
p=15,39 v=38,4
p=15,50 v=-90,-27
p=63,66 v=29,60
p=17,79 v=-50,-89
p=21,39 v=-83,60
p=62,35 v=6,-51
p=99,44 v=9,8
p=32,4 v=-43,-35
p=57,2 v=-81,39
p=4,7 v=-44,-10
p=93,12 v=56,-19
p=33,79 v=45,-24
p=66,72 v=86,-41
p=84,83 v=-28,39
p=11,30 v=56,-88
p=51,19 v=-54,-89
p=0,36 v=-32,-28
p=69,39 v=41,20
p=75,88 v=8,79
p=5,12 v=-64,93
p=38,9 v=-62,91
p=25,100 v=64,-15
p=97,60 v=-12,-69
p=57,76 v=-7,-13
p=42,48 v=-69,32
p=31,35 v=-81,65
p=28,28 v=46,85
p=77,84 v=-73,68
p=57,18 v=-54,44
p=48,98 v=59,17
p=28,86 v=92,-17
p=1,93 v=90,-53
p=0,31 v=84,-32
p=87,34 v=9,-65
p=20,20 v=44,-25
p=65,93 v=-96,-54
p=45,75 v=-82,-87
p=48,58 v=26,65
p=87,76 v=-86,16
p=31,79 v=52,34
p=100,101 v=58,-36
p=76,20 v=59,-76
p=100,91 v=-85,-4
p=14,13 v=92,-46
p=72,66 v=-60,-65
p=74,41 v=-60,96
p=73,17 v=48,80
p=48,50 v=-48,94
p=76,50 v=-81,12
p=51,102 v=80,1
p=44,12 v=93,-86
p=49,16 v=40,-43
p=75,97 v=2,-24
p=54,35 v=13,-34
p=50,4 v=-75,75
p=28,87 v=77,81
p=50,38 v=13,60
p=51,43 v=-27,22
p=71,52 v=-13,29
p=37,74 v=-29,-40
p=67,82 v=-13,61
p=4,21 v=-97,-50
p=36,2 v=25,-66
p=17,67 v=-30,-87
p=90,29 v=96,15
p=90,96 v=-52,39
p=20,9 v=78,17
p=70,74 v=-32,-93
p=77,60 v=-83,51
p=74,97 v=92,1
p=75,82 v=28,61
p=40,4 v=-58,-94
p=64,26 v=6,-86
p=71,14 v=-81,-93
p=11,96 v=24,70
p=97,26 v=97,-61
p=77,53 v=-44,55
p=1,93 v=83,46
p=11,59 v=90,54
p=3,2 v=-85,-8
p=18,90 v=79,-60
p=82,31 v=96,-34
p=41,22 v=-9,-72
p=52,17 v=86,-95
p=48,39 v=-41,-63
p=3,61 v=-16,-35
p=6,26 v=50,42
p=73,91 v=-53,99
p=55,89 v=-54,-35
p=24,22 v=17,-66
p=80,84 v=-97,53
p=20,73 v=69,-40
p=26,100 v=-83,19
p=68,87 v=-40,88
p=80,45 v=-15,-4
p=2,80 v=-98,23
p=30,71 v=-62,9
p=2,89 v=-50,88
p=19,99 v=21,-1
p=11,81 v=-3,-13
p=70,59 v=-32,34
p=84,68 v=89,-22
p=48,35 v=-48,-79
p=78,80 v=-60,34
p=10,50 v=77,-31
p=94,79 v=23,90
p=46,47 v=-68,2
p=20,52 v=30,-56
p=79,36 v=2,-97
p=75,21 v=-74,-22
p=54,53 v=4,-75
p=100,76 v=69,43
p=95,33 v=-64,22
p=26,61 v=38,63
p=13,20 v=43,-59
p=20,36 v=72,47
p=78,98 v=49,-24
p=83,76 v=-12,-22
p=79,98 v=-73,77
p=31,22 v=-96,33
p=61,22 v=27,-90
p=83,41 v=22,40
p=98,15 v=76,-37
p=23,46 v=-49,43
p=2,87 v=-52,-8
p=74,2 v=74,-57
p=19,36 v=51,36
p=30,93 v=-3,-93
p=75,21 v=90,-28
p=2,32 v=30,-12
p=69,88 v=-40,-90
p=1,77 v=-81,27
p=1,70 v=-3,56
p=60,70 v=3,-13
p=54,14 v=67,8
p=1,61 v=64,-47
p=50,99 v=-40,72
p=62,43 v=74,76
p=63,30 v=14,98
p=10,3 v=24,-46
p=25,64 v=-83,-87
p=8,31 v=-70,54
p=17,18 v=-80,-95
p=100,53 v=42,-6
p=61,101 v=28,-44
p=91,34 v=-6,23
p=43,28 v=-31,62
p=52,88 v=75,6
p=45,66 v=-28,-96
p=29,40 v=5,78
p=21,91 v=-50,97
p=72,6 v=48,35
p=4,11 v=97,48
p=66,89 v=-72,4
p=63,70 v=-47,56
p=28,75 v=10,-78
p=94,47 v=-65,-27
p=14,74 v=51,25
p=34,65 v=5,-29
p=61,65 v=-73,47
p=88,17 v=-12,-28
p=46,43 v=-76,40
p=78,76 v=21,34
p=54,32 v=-41,-7
p=48,92 v=-41,-35
p=89,59 v=29,9
p=89,66 v=36,-58
p=38,83 v=-19,16
p=84,4 v=38,66
p=39,101 v=88,54
p=6,32 v=82,50
p=30,79 v=-63,-80
p=54,35 v=41,-16
p=33,102 v=-49,35
p=23,37 v=-91,-61
p=77,30 v=-35,-69
p=79,4 v=35,-57
p=45,29 v=-95,-23
p=49,57 v=87,71
p=53,31 v=87,-7
p=10,78 v=-34,35
p=10,16 v=-37,26
p=46,15 v=-41,-95
p=2,26 v=49,-30
p=15,68 v=-50,52
p=69,40 v=-47,-90
p=53,39 v=14,9
p=23,61 v=11,83
p=44,33 v=-27,-79
p=81,2 v=-40,-55
p=89,42 v=-72,31
p=18,32 v=-39,19
p=97,91 v=-76,52
p=51,80 v=60,-64
p=48,36 v=53,98
p=86,91 v=97,1
p=48,88 v=-22,-42
p=100,17 v=-24,-90
p=40,52 v=-36,-85
p=65,52 v=-19,-25
p=0,45 v=82,-84
p=12,102 v=78,-73
p=67,19 v=87,-43
p=24,78 v=84,93
p=68,52 v=-14,71
p=57,10 v=-7,6
p=4,10 v=87,62
p=64,38 v=41,51
p=23,56 v=12,-38
p=53,25 v=46,-54
p=16,90 v=-84,59
p=23,97 v=-36,-35
p=35,78 v=85,92
p=53,31 v=-53,19
p=61,84 v=54,41
p=86,93 v=76,10
p=78,80 v=-46,-24
p=2,6 v=43,46
p=58,82 v=20,97
p=46,20 v=-75,15
p=60,85 v=30,-54
p=57,29 v=-7,-96
p=75,98 v=62,79
p=64,44 v=54,24
p=42,51 v=-76,19
p=37,66 v=-70,-9
p=85,16 v=-58,60
p=8,70 v=-57,-69
p=36,55 v=-73,-96
p=67,93 v=88,30
p=18,48 v=54,-94
p=53,14 v=20,-93
p=28,37 v=24,-10
p=64,10 v=-94,-94
p=10,19 v=-41,-19
p=6,80 v=-64,-98
p=30,66 v=52,54
p=35,24 v=5,-3
p=14,1 v=18,-60
p=54,0 v=-61,-17
p=27,2 v=38,60
p=49,88 v=48,-17
p=63,1 v=-7,-94
p=98,75 v=-78,43
p=88,97 v=-25,10
p=30,75 v=-68,26
p=65,57 v=68,18
p=19,30 v=3,49
p=52,47 v=-50,94
p=16,8 v=27,-52
p=27,64 v=-84,-29
p=85,30 v=49,31
p=94,20 v=10,80
p=57,84 v=-84,32
p=45,93 v=-2,-46
p=94,68 v=-79,83
p=61,27 v=-73,2
p=15,64 v=99,-4
p=28,82 v=46,-51
p=88,14 v=81,74
p=2,90 v=67,48
p=8,66 v=64,-31
p=4,5 v=-97,48
p=39,60 v=5,74
p=62,11 v=81,-75
p=82,54 v=35,-84
p=89,72 v=-88,-24
p=42,45 v=73,67
p=77,78 v=89,90
p=44,5 v=80,28
p=58,28 v=-31,-31
p=91,68 v=-73,-22
p=68,48 v=-76,80
p=22,26 v=24,6
p=34,6 v=-69,-83
p=64,81 v=-13,-8
p=15,16 v=-16,64
p=42,49 v=93,-56
p=87,38 v=-59,58
p=95,34 v=-92,-79
p=66,101 v=7,-19
p=63,38 v=82,1
p=95,53 v=63,-27
p=95,2 v=-65,-48
p=16,73 v=-64,-78
p=13,60 v=-58,27
p=69,55 v=34,-11
p=32,20 v=10,-71
p=52,100 v=-95,-84
p=50,16 v=-21,35
p=42,44 v=73,85
p=87,33 v=-93,74
p=60,75 v=40,34
p=7,1 v=27,-13
p=12,63 v=29,-63
p=18,56 v=-90,-18
p=30,49 v=-56,-92
p=49,8 v=-63,49
p=80,15 v=64,-79
p=36,66 v=25,-24
p=76,37 v=-80,22
p=76,68 v=29,-31
p=90,2 v=-92,51
p=14,42 v=72,49
p=57,89 v=-94,-53
p=48,63 v=29,-14
p=80,2 v=-86,19
p=62,45 v=-53,36
p=40,36 v=67,67
p=45,96 v=62,-66
p=42,61 v=-81,18
p=54,69 v=-27,-79
p=72,17 v=55,-97
p=92,93 v=-3,3
p=11,97 v=10,59
p=33,61 v=5,-65
p=26,31 v=-48,34
p=1,32 v=81,-18
p=68,70 v=27,-54
p=89,10 v=-99,-12
p=32,102 v=65,-64
p=52,89 v=-41,-62
p=100,10 v=-11,-59
p=93,63 v=83,34
p=16,67 v=-23,43
p=58,35 v=13,-10
p=24,31 v=85,-81
p=17,94 v=58,10
p=40,55 v=-22,-65
p=47,12 v=-10,54
p=52,29 v=-68,-25
p=11,83 v=-71,-80
p=62,29 v=-46,-25
p=73,95 v=67,47
p=74,25 v=81,14
p=99,45 v=11,-13
p=31,13 v=76,26
p=27,90 v=17,-42
p=68,80 v=89,-95
p=14,54 v=57,94
p=97,56 v=29,-9
p=65,8 v=-68,57
p=86,76 v=82,83
p=81,43 v=96,-54
p=86,17 v=56,26
p=41,92 v=-82,-10
p=49,41 v=60,58
p=49,72 v=93,3
p=50,13 v=-7,17
p=46,85 v=-5,81
p=7,40 v=-51,78
p=88,91 v=62,59
p=97,14 v=93,86
p=62,80 v=98,62
p=22,50 v=51,-65
p=14,34 v=77,38
p=34,86 v=-54,-6
p=2,86 v=-85,25
p=67,28 v=-15,51
p=90,87 v=19,-48
p=13,25 v=45,-5
p=59,62 v=13,-67
p=53,40 v=13,11
p=48,52 v=37,74
p=82,62 v=-66,-47
p=74,64 v=-25,-25
p=72,28 v=-68,-37
p=51,8 v=-48,73
p=46,32 v=-89,75
p=100,15 v=-34,58
p=34,39 v=-8,71
p=78,70 v=47,-22
p=31,50 v=-49,2
p=15,84 v=64,-98
p=76,92 v=75,-6
p=40,18 v=54,-96
p=81,3 v=1,-96
p=9,95 v=-85,-89
p=64,18 v=34,87
p=72,100 v=48,12
p=39,77 v=40,-53
p=74,11 v=95,-77
p=29,20 v=72,51
p=96,91 v=-19,-75
p=11,68 v=51,-58
p=64,9 v=67,28
p=73,94 v=-46,30
p=77,39 v=15,49
p=100,20 v=-85,-77
p=92,80 v=-5,-64
p=80,11 v=-62,-41
p=78,47 v=-54,74
p=32,76 v=24,-12
p=13,22 v=-10,-52
p=30,78 v=-49,-53
  </pre>
</details>

<details>
	<summary>Click to show the easter egg</summary>
	<pre>
                                X                                                               X    
                                                       X               X                 X     X     
 X                                                                                         X         
                                                                          X                          
                                                              X                                      
                               X      X                        X                                     
                 X                    X                                         X                    
                       X                                                                             
                X                  X                        X                                        
      X                                                                              X               
                                                                            X                        
                                                                                                 X   
                  X                    XX                                X                           
               X                                           X                   X                     
                                                                                                    X
                    X                                                                                
                                                             X                                       
                                                                                                     
                                      X                            X                                 
                                                X                                                    
                                                                 X          X        X           X   
                                              X                                                      
                                                                                                     
  X                                      X                     X             X                       
                                                                                                     
    X                              X               X    X                                            
                                             X           X                                           
                                                                                              X      
                                                              X     X                                
       X                                                                                             
                                                                                                     
               X                                                                                     
                                  X          X                                                       
                                                           X                                         
                                                                   X                                 
                                                                                X                    
                                                                                                   X 
                                             XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                         
                                             X                             X                         
                                             X                             X  X                      
                                             X                             X                        X
                        X                    X                             X                         
       X    X       X   X                    X              X              X    X                    
X                                            X             XXX             X                    X    
                     X       X               X            XXXXX            X                         
       X                                     X           XXXXXXX           X                         
                X                   X        X          XXXXXXXXX          X                         
                                             X            XXXXX            X                       X 
                                             X           XXXXXXX           X                         
                                             X          XXXXXXXXX          X                   X     
               X                             X         XXXXXXXXXXX         X                         
       X                                     X        XXXXXXXXXXXXX        X              X          
                                             X          XXXXXXXXX          X    X                    
  X                   X             X        X         XXXXXXXXXXX         X                         
    X                                        X        XXXXXXXXXXXXX        X                         
                                             X       XXXXXXXXXXXXXXX       X                         
                                             X      XXXXXXXXXXXXXXXXX      X                         
                                             X        XXXXXXXXXXXXX        X                         
                   X                X        X       XXXXXXXXXXXXXXX       X    X                    
                                             X      XXXXXXXXXXXXXXXXX      X                    X    
   X                                         X     XXXXXXXXXXXXXXXXXXX     X                         
                                             X    XXXXXXXXXXXXXXXXXXXXX    X    X                    
     X                                       X             XXX             X                         
                                             X             XXX             X                         
                                             X             XXX             X                         
                                             X                             X                         
                                             X                             X                         
                    X                        X                             X                         
                    X    X          X        X                             X X                       
                   X                         XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX         X               
                                                                                                     
    X                                     X                             X                            
                                       X                                                             
                                      X                       X                                      
                                                                                    X                
               X                                                                                     
                                                                                                 X   
                                                                                  X                  
                                                                                                     
                                                                                                     
                                                                                                     
                                             X     X                                                 
                                                                   X        X     X                  
         X                                  X  X                 X                                   
      X                                 X                           X                                
                                                                        X                         X  
                                                                                                     
                                X                        X                                           
                                                                 X          X                        
                 X              X                   X                                                
                                                                                                     
                                                                                                     
                                                                                                     
                                                                  X                 X                
                                                                                                     
                                             X                                                       
                                                     X                        X                      
              X                                                                                      
                              X                        X                           X                 
                         X                                                                           
                                                                                                     
                                                                                         X           
                                                                   X   XX                            
  </pre>
</details>