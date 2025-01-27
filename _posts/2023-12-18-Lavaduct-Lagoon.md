---
title: Lavaduct Lagoon
description: Advent of Code 2023 [Day 18]
layout: default
lang: en
tag: aoc23
prefetch:
  - adventofcode.com
  - en.wikipedia.org
---

Thanks to your efforts, the machine parts factory is one of the first factories up and running since the lavafall came back. However, to catch up with the large backlog of parts requests, the factory will also need a **large supply of lava** for a while; the Elves have already started creating a large lagoon nearby for this purpose.

However, they aren't sure the lagoon will be big enough; they've asked you to take a look at the **dig plan** (your puzzle input). For example:

```
R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
```

The digger starts in a 1 meter cube hole in the ground. They then dig the specified number of meters **up** (`U`), **down** (`D`), **left** (`L`), or **right** (`R`), clearing full 1 meter cubes as they go. The directions are given as seen from above, so if "up" were north, then "right" would be east, and so on. Each trench is also listed with **the color that the edge of the trench should be painted** as an [RGB hexadecimal color code](https://en.wikipedia.org/wiki/RGB_color_model#Numeric_representations).

When viewed from above, the above example dig plan would result in the following loop of **trench** (`#`) having been dug out from otherwise **ground-level terrain** (`.`):

```
#######
#.....#
###...#
..#...#
..#...#
###.###
#...#..
##..###
.#....#
.######
```

At this point, the trench could contain 38 cubic meters of lava. However, this is just the edge of the lagoon; the next step is to **dig out the interior** so that it is one meter deep as well:

```
#######
#######
#######
..#####
..#####
#######
#####..
#######
.######
.######
```

Now, the lagoon can contain a much more respectable `62` cubic meters of lava. While the interior is dug out, the edges are also painted according to the color codes in the dig plan.

The Elves are concerned the lagoon won't be large enough; if they follow their dig plan, **how many cubic meters of lava could it hold?**

```go
type Point struct {
	x, y int
}

type Instruction struct {
	dir   string
	steps int
	color string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var instructions []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var inst Instruction
		fmt.Sscanf(line, "%s %d (#%s)", &inst.dir, &inst.steps, &inst.color)
		instructions = append(instructions, inst)
	}

	vertices := []Point{ {0, 0} }
	boundaries := make(map[Point]bool)
	currentPos := Point{0, 0}
	boundaries[currentPos] = true

	perimeter := 0
	for _, inst := range instructions {
		dx, dy := 0, 0
		switch inst.dir {
		case "U":
			dy = -1
		case "D":
			dy = 1
		case "L":
			dx = -1
		case "R":
			dx = 1
		}

		for i := 0; i < inst.steps; i++ {
			currentPos.x += dx
			currentPos.y += dy
			boundaries[currentPos] = true
		}
		vertices = append(vertices, currentPos)
		perimeter += inst.steps
	}

	area := 0
	for i := 0; i < len(vertices)-1; i++ {
		area += vertices[i].x*vertices[i+1].y - vertices[i+1].x*vertices[i].y
	}
	area += vertices[len(vertices)-1].x*vertices[0].y - vertices[0].x*vertices[len(vertices)-1].y
	if area < 0 {
		area = -area
	}
	area /= 2

	interior := area - perimeter/2 + 1
	total := interior + perimeter

	fmt.Printf("The lagoon can hold %d cubic meters of lava\n", total)
}
```

The Elves were right to be concerned; the planned lagoon would be **much too small**.

After a few minutes, someone realizes what happened; someone **swapped the color and instruction parameters** when producing the dig plan. They don't have time to fix the bug; one of them asks if you can **extract the correct instructions** from the hexadecimal codes.

Each hexadecimal code is **six hexadecimal digits** long. The first five hexadecimal digits encode the **distance in meters** as a five-digit hexadecimal number. The last hexadecimal digit encodes the **direction to dig**: `0` means `R`, `1` means `D`, `2` means `L`, and `3` means `U`.

So, in the above example, the hexadecimal codes can be converted into the true instructions:

- `#70c710` = `R 461937`
- `#0dc571` = `D 56407`
- `#5713f0` = `R 356671`
- `#d2c081` = `D 863240`
- `#59c680` = `R 367720`
- `#411b91` = `D 266681`
- `#8ceee2` = `L 577262`
- `#caa173` = `U 829975`
- `#1b58a2` = `L 112010`
- `#caa171` = `D 829975`
- `#7807d2` = `L 491645`
- `#a77fa3` = `U 686074`
- `#015232` = `L 5411`
- `#7a21e3` = `U 500254`

Digging out this loop and its interior produces a lagoon that can hold an impressive `952408144115` cubic meters of lava.

Convert the hexadecimal color codes into the correct instructions; if the Elves follow this new dig plan, **how many cubic meters of lava could the lagoon hold?**

```go
type Point struct {
	x, y int64
}

func hexToInstruction(hex string) (dir string, steps int64) {
	hex = strings.Trim(hex, "(#)")

	dirCode := hex[len(hex)-1]
	switch dirCode {
	case '0':
		dir = "R"
	case '1':
		dir = "D"
	case '2':
		dir = "L"
	case '3':
		dir = "U"
	}

	distance := hex[:5]
	steps, _ = strconv.ParseInt(distance, 16, 64)
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	vertices := []Point{ {0, 0} }
	currentPos := Point{0, 0}
	var perimeter int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		color := parts[2]

		dir, steps := hexToInstruction(color)

		switch dir {
		case "U":
			currentPos.y -= steps
		case "D":
			currentPos.y += steps
		case "L":
			currentPos.x -= steps
		case "R":
			currentPos.x += steps
		}

		vertices = append(vertices, currentPos)
		perimeter += steps
	}

	var area int64
	for i := 0; i < len(vertices)-1; i++ {
		area += vertices[i].x*vertices[i+1].y - vertices[i+1].x*vertices[i].y
	}
	area += vertices[len(vertices)-1].x*vertices[0].y - vertices[0].x*vertices[len(vertices)-1].y
	if area < 0 {
		area = -area
	}
	area = area / 2

	interior := area - perimeter/2 + 1
	total := interior + perimeter

	fmt.Printf("The lagoon can hold %d cubic meters of lava\n", total)
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2023/day/18)

<details>
	<summary>Click to show the input</summary>
	<pre>
R 4 (#9505a2)
U 3 (#984f41)
R 5 (#763772)
U 4 (#5e0883)
R 8 (#6dff52)
D 4 (#5e0881)
R 9 (#220312)
U 7 (#984f43)
L 8 (#1eacc2)
U 8 (#0f1673)
L 2 (#7cab42)
U 9 (#7dae23)
L 5 (#0d4642)
D 13 (#2be373)
L 4 (#225752)
D 4 (#3f8d43)
L 7 (#2930b2)
U 10 (#67f943)
L 6 (#25a0a0)
D 12 (#69b2d3)
L 3 (#8be9f0)
U 4 (#207593)
L 3 (#b18a92)
U 5 (#63c483)
L 9 (#613142)
U 2 (#a6aff3)
L 4 (#7a5082)
U 4 (#2beb03)
R 11 (#050d70)
U 2 (#665ed3)
R 2 (#5422f2)
U 6 (#264a73)
L 7 (#5422f0)
U 3 (#47fcc3)
L 6 (#050d72)
U 8 (#0d1f63)
L 7 (#2828e2)
U 5 (#440971)
L 6 (#547822)
D 8 (#9961b1)
L 5 (#547820)
D 6 (#304541)
L 6 (#8bd7f2)
D 3 (#4eb533)
R 11 (#5c4602)
D 4 (#096b13)
L 11 (#260da0)
D 4 (#400e13)
L 8 (#5f2bb0)
D 5 (#760133)
L 7 (#6a72f0)
D 9 (#33bdc3)
L 3 (#80d940)
U 4 (#155633)
L 4 (#8f4e82)
U 8 (#85bc63)
L 3 (#26b562)
U 4 (#66fa93)
L 7 (#b26d12)
U 11 (#67b8f3)
L 2 (#2bb2f0)
U 3 (#b531a3)
R 4 (#4013d0)
U 6 (#2bd863)
R 2 (#4a6a80)
U 12 (#0f7c13)
R 5 (#6cc370)
U 5 (#109003)
R 4 (#457c40)
U 9 (#935793)
R 4 (#1edea0)
U 4 (#5bc2b3)
R 3 (#4e6750)
U 6 (#947b93)
R 6 (#7fafb0)
U 5 (#8f4c91)
R 2 (#0fa3b0)
U 6 (#21c311)
L 12 (#485f20)
U 3 (#3f2ea1)
L 3 (#8688c0)
U 4 (#0175f3)
R 4 (#7283b0)
U 6 (#64da33)
R 7 (#23e570)
U 7 (#869213)
R 4 (#12f200)
D 8 (#42f3b1)
R 5 (#854f90)
D 3 (#34eaf1)
R 9 (#854f92)
D 6 (#695841)
R 12 (#05b840)
D 6 (#0a3561)
L 3 (#901430)
D 3 (#5ece03)
L 8 (#5f85b0)
D 8 (#3856f3)
R 9 (#4e6aa0)
D 9 (#1cba43)
L 9 (#4b8f20)
D 5 (#730143)
R 11 (#726f40)
D 7 (#2bc223)
R 6 (#2617e0)
U 6 (#2c3d13)
R 4 (#6482c0)
U 5 (#2d55f3)
R 4 (#6e4650)
U 8 (#2d55f1)
R 6 (#0138d0)
U 2 (#243eb3)
R 7 (#17bd10)
U 5 (#7c0aa3)
R 4 (#31a982)
U 7 (#5ae133)
R 4 (#b5ad22)
U 9 (#6a6203)
R 5 (#046852)
U 4 (#12d6b3)
R 6 (#0d09e2)
U 7 (#526203)
R 11 (#3f8962)
U 5 (#4245e3)
L 4 (#699702)
U 6 (#17dce3)
L 11 (#44f9a2)
U 3 (#0c7143)
L 4 (#09aec2)
D 4 (#746e03)
L 6 (#49edf0)
U 4 (#839f83)
L 5 (#5e9ce0)
U 5 (#839f81)
R 9 (#5c47d0)
U 3 (#17f3f3)
R 2 (#3180a0)
U 11 (#7a7973)
R 3 (#7d7750)
U 7 (#9ed523)
R 3 (#6b5450)
D 9 (#523bb3)
R 5 (#9a1b80)
D 3 (#4323e3)
R 3 (#299430)
U 2 (#97c0e3)
R 4 (#6e8450)
U 10 (#448eb1)
R 5 (#63bda2)
U 3 (#702a61)
R 3 (#2c2690)
U 12 (#7379e1)
R 4 (#2c2692)
U 7 (#0c4041)
R 3 (#63bda0)
U 7 (#45cc21)
L 4 (#4cd2e0)
U 4 (#2ae8f1)
L 7 (#738540)
D 11 (#a34001)
L 6 (#738542)
D 2 (#16d291)
L 4 (#05ba40)
D 5 (#8872a1)
L 9 (#186db0)
D 3 (#8e7ce1)
L 5 (#9c6d10)
U 6 (#8e7ce3)
L 9 (#56a350)
U 2 (#9df721)
L 6 (#4e9042)
U 3 (#2659a1)
R 4 (#3bd022)
U 5 (#2659a3)
R 11 (#811db2)
U 5 (#ad2b91)
L 9 (#05ba42)
U 5 (#72a9c1)
R 8 (#211510)
U 5 (#13cb31)
R 6 (#6c2480)
D 7 (#13cb33)
R 6 (#2f0f60)
U 7 (#4683f1)
R 4 (#0054d0)
U 4 (#8bf741)
L 13 (#7fb5c0)
U 2 (#149691)
L 3 (#196d60)
U 6 (#5fef73)
R 3 (#0cf060)
U 5 (#52b473)
R 6 (#022db0)
U 9 (#346de3)
R 6 (#1f0e40)
U 5 (#251e63)
R 10 (#777a02)
U 4 (#6f8d93)
R 6 (#777a00)
U 12 (#3cdb63)
L 3 (#426ac2)
U 8 (#1cd103)
L 10 (#83a360)
U 5 (#07e093)
L 3 (#2a5f40)
U 6 (#22d423)
R 2 (#448862)
U 6 (#8da513)
R 7 (#697a42)
U 2 (#6478e3)
R 3 (#426ac0)
U 4 (#8c62b3)
R 11 (#908510)
U 2 (#625113)
R 2 (#7dcc10)
U 4 (#315be3)
L 10 (#242e40)
U 2 (#352c83)
L 3 (#68e2f0)
U 4 (#658603)
R 3 (#4aa570)
U 4 (#387113)
L 3 (#3616a0)
U 4 (#3bcef3)
L 6 (#3616a2)
U 12 (#68c1b3)
L 2 (#4aa572)
U 8 (#028b23)
L 10 (#4126a0)
U 3 (#14e773)
R 10 (#763032)
U 2 (#86bf23)
R 4 (#3f43a2)
U 5 (#468f11)
R 7 (#0ec812)
U 3 (#99ca31)
L 7 (#4bc6a0)
U 8 (#67a621)
L 3 (#4bc6a2)
U 3 (#53c3a1)
L 10 (#6d0ad0)
U 6 (#218a01)
R 12 (#6d0ad2)
U 3 (#4e8fe1)
R 8 (#52f3a2)
U 7 (#2ae183)
R 8 (#420262)
U 3 (#a3eba3)
R 3 (#6f0962)
U 11 (#9018d3)
R 6 (#444152)
U 7 (#acf6f3)
R 5 (#4caec2)
D 10 (#47f963)
R 3 (#2094f2)
U 13 (#02f073)
R 3 (#2cb412)
U 7 (#29d5b3)
R 5 (#999602)
D 2 (#52abf3)
R 4 (#7fbe72)
D 9 (#5210d3)
R 4 (#b456f2)
D 9 (#378693)
R 5 (#b456f0)
U 10 (#58bea3)
R 5 (#111ef2)
D 7 (#1eb663)
R 3 (#873e92)
D 2 (#02cbb3)
R 3 (#8756c2)
D 12 (#682d83)
R 3 (#64e760)
U 7 (#6478a3)
R 8 (#7a9d20)
D 4 (#6ec453)
R 8 (#8aec52)
D 3 (#858fa3)
R 5 (#8aec50)
D 6 (#5fe253)
L 7 (#804c10)
D 4 (#5721c1)
L 3 (#5a26e0)
D 5 (#344fc1)
L 2 (#260642)
D 3 (#947521)
L 10 (#260640)
D 4 (#10be21)
L 6 (#28aba0)
U 7 (#239181)
R 4 (#a84ed0)
U 5 (#594d81)
L 4 (#523490)
U 4 (#5d20e3)
L 3 (#621ce0)
D 3 (#5d20e1)
L 9 (#45f630)
U 5 (#07f8b1)
L 2 (#9bceb0)
U 4 (#541ae1)
L 9 (#56a000)
D 7 (#5cdf31)
L 9 (#2f94d0)
D 2 (#1705b3)
L 7 (#7f1d70)
D 7 (#99f463)
R 6 (#8a9190)
D 5 (#9f9111)
R 5 (#02d950)
D 4 (#1b0841)
R 7 (#199fc0)
D 2 (#739481)
R 8 (#199fc2)
U 4 (#4b47c1)
R 9 (#7014b0)
D 4 (#24c0a1)
R 15 (#30ae00)
D 4 (#8e1d01)
L 8 (#85f080)
D 3 (#2f10a1)
L 3 (#46b8e2)
D 4 (#830d61)
L 12 (#04b302)
D 2 (#9d5ae1)
L 3 (#018932)
D 5 (#04ada3)
L 6 (#2fe952)
D 5 (#3ff843)
L 7 (#8e4dd2)
D 6 (#78d433)
L 12 (#127612)
D 6 (#3f20f1)
R 12 (#994c62)
D 4 (#5cb071)
R 3 (#994c60)
D 3 (#21a8b1)
R 8 (#533c92)
D 3 (#58d461)
R 7 (#316010)
D 5 (#3772e1)
R 4 (#3efac2)
D 3 (#185d41)
R 10 (#2a5450)
D 4 (#66b841)
R 3 (#2a5452)
D 7 (#4c84c1)
R 4 (#3efac0)
D 3 (#0544d1)
R 4 (#959a70)
D 6 (#07b843)
R 2 (#789922)
D 7 (#97a0f3)
R 5 (#789920)
D 3 (#68f8c3)
R 3 (#632b70)
D 7 (#0628c1)
R 3 (#46b8e0)
D 9 (#0a5911)
R 3 (#0caab0)
U 6 (#247553)
R 3 (#451e30)
U 9 (#53d613)
R 5 (#85d060)
U 7 (#1e3623)
R 5 (#044a50)
D 9 (#558213)
R 7 (#481c60)
D 5 (#49bbd3)
L 7 (#481c62)
D 6 (#830fb3)
R 2 (#6b1ef0)
D 2 (#0e5861)
R 9 (#070aa0)
D 6 (#693991)
R 4 (#070aa2)
U 4 (#4c5121)
R 9 (#205ad0)
U 5 (#0c67d1)
R 7 (#062fc0)
U 6 (#16c763)
R 8 (#669720)
U 5 (#4e9843)
R 6 (#05bc32)
U 4 (#9513d3)
R 3 (#05bc30)
U 2 (#928d03)
R 4 (#a01410)
D 11 (#7b2af1)
R 6 (#4c41d0)
D 13 (#52bb51)
R 3 (#9001e0)
D 2 (#0a18d1)
R 4 (#968322)
D 5 (#509641)
L 13 (#1abda2)
D 6 (#540201)
L 9 (#2b02f2)
D 2 (#106921)
L 5 (#3a2dc0)
D 6 (#319fd3)
R 9 (#56ac20)
D 3 (#319fd1)
R 5 (#400070)
D 7 (#735a51)
R 10 (#20c2b2)
D 2 (#a17341)
R 3 (#005a72)
D 5 (#224c51)
L 7 (#022422)
D 3 (#3a9061)
L 9 (#402e82)
D 8 (#4d0491)
L 5 (#6e9c32)
D 2 (#0b58d1)
L 15 (#2bb0f0)
U 4 (#8a1051)
L 9 (#2bb0f2)
U 5 (#84c961)
L 3 (#39dea2)
U 13 (#75c911)
L 4 (#3cb3c2)
D 4 (#2466e1)
L 8 (#3cb3c0)
D 8 (#685971)
L 4 (#6adcc2)
D 6 (#1e7673)
L 5 (#526972)
D 4 (#1b2ad3)
L 7 (#853862)
D 9 (#696ea3)
R 10 (#0d3450)
D 4 (#4633b1)
R 3 (#7c6780)
D 9 (#4633b3)
R 5 (#4e0600)
D 8 (#5f7983)
R 7 (#8861d2)
U 5 (#79c723)
R 5 (#4b16b2)
U 4 (#b58831)
R 3 (#06a1b2)
U 9 (#59fd71)
R 7 (#6fa5e2)
U 6 (#873f01)
R 3 (#6fa5e0)
U 3 (#61b3c1)
R 5 (#06a1b0)
D 7 (#6931a1)
R 9 (#ab7272)
U 7 (#640df3)
R 10 (#00b292)
D 7 (#8488f3)
R 6 (#8572c2)
D 8 (#105483)
L 7 (#5f6022)
U 5 (#3fadd3)
L 2 (#5f6020)
U 7 (#7ee933)
L 3 (#333b12)
D 10 (#221823)
L 3 (#a61062)
D 2 (#880f83)
L 5 (#0c2a12)
D 9 (#79c721)
L 5 (#7ccff2)
D 3 (#35a183)
R 8 (#14e432)
U 6 (#557103)
R 5 (#b5e082)
U 2 (#6bee81)
R 8 (#261612)
D 2 (#1f2401)
R 4 (#af2ae2)
D 11 (#40ed21)
L 6 (#1c2ef2)
D 5 (#293701)
L 7 (#3699c0)
D 6 (#5708b1)
L 3 (#8dd6c0)
D 12 (#2894f1)
L 5 (#513a12)
D 7 (#198a23)
R 12 (#20de32)
D 2 (#198a21)
R 9 (#525842)
D 4 (#273111)
L 10 (#3dd5a0)
D 5 (#14f9b1)
L 11 (#269190)
D 5 (#3d5191)
L 10 (#18a862)
U 3 (#71c981)
L 5 (#18a860)
U 4 (#1c7671)
R 8 (#6c8330)
U 11 (#5356c3)
L 8 (#41d400)
U 9 (#6b4d03)
R 5 (#41d402)
U 14 (#0cedc3)
L 5 (#686260)
D 5 (#717c11)
L 7 (#2afee0)
D 8 (#103261)
L 5 (#95f2c0)
D 9 (#81ae73)
L 8 (#88ebb0)
D 10 (#0ccc91)
L 7 (#5866b0)
D 9 (#304491)
L 3 (#2e6b60)
D 7 (#67abd1)
L 5 (#2e6b62)
D 9 (#2b3131)
L 2 (#5866b2)
D 11 (#375731)
L 8 (#5855c0)
D 7 (#33fcb1)
L 4 (#4b9132)
D 6 (#5bc731)
R 5 (#4b9130)
D 5 (#36f2b1)
R 7 (#68cbf2)
D 4 (#2e6191)
R 8 (#a1b162)
D 6 (#585861)
R 4 (#004682)
U 5 (#23f741)
R 6 (#abffe2)
D 5 (#46fcd3)
R 10 (#8cc712)
D 6 (#46fcd1)
R 6 (#31aad2)
U 3 (#4ed551)
R 3 (#9a6852)
U 5 (#7b9781)
R 9 (#5d9402)
D 8 (#20bc71)
R 10 (#577972)
U 6 (#624ef1)
R 7 (#2e7e40)
D 8 (#0e1851)
R 4 (#868f30)
D 6 (#9ac001)
R 3 (#91bd02)
D 7 (#8455f1)
R 6 (#3dc142)
D 4 (#6bf421)
L 9 (#262a72)
D 2 (#6c0b41)
L 2 (#2def60)
D 3 (#2e8621)
L 7 (#2def62)
D 7 (#5cd0a1)
L 3 (#2eabf2)
D 9 (#87f891)
L 7 (#7801e2)
U 4 (#2d78c1)
L 4 (#884eb2)
U 8 (#5f91d1)
L 8 (#b069c2)
U 3 (#706ec3)
L 10 (#14d6a0)
U 2 (#2e1ee3)
L 5 (#6fcd20)
U 2 (#2e1ee1)
L 4 (#6a8280)
U 9 (#10aed3)
L 6 (#28d1e0)
U 8 (#9ef053)
L 9 (#0e4470)
U 3 (#9fcdf3)
L 6 (#6be0a0)
U 4 (#4b5df3)
L 13 (#3d1b92)
D 5 (#0a7c03)
L 4 (#4225d2)
D 4 (#1a1661)
L 12 (#012740)
D 5 (#3ce951)
L 3 (#52df42)
D 7 (#ace6c1)
L 8 (#52df40)
D 4 (#255791)
R 7 (#012742)
D 3 (#4a15d1)
R 10 (#318982)
D 3 (#510b33)
L 14 (#7cfeb2)
D 2 (#510b31)
L 3 (#13c862)
D 5 (#126f93)
L 8 (#204922)
D 8 (#7b3623)
L 5 (#6c3780)
U 4 (#0c8663)
L 9 (#6c3782)
U 3 (#794583)
L 11 (#204920)
U 8 (#5fe243)
L 3 (#508b42)
U 4 (#62a383)
R 10 (#67e9b2)
U 8 (#94a731)
R 6 (#8e9b22)
U 4 (#858011)
L 10 (#8e9b20)
U 7 (#4c8cd1)
L 6 (#8f5152)
U 8 (#483551)
L 4 (#42e142)
D 12 (#40e3a1)
L 4 (#0b5b92)
D 9 (#0d1471)
R 4 (#37cb02)
D 5 (#b1c241)
L 8 (#67aa22)
D 5 (#799ad1)
L 5 (#3aa9f2)
D 7 (#473731)
L 4 (#60de22)
D 6 (#65e001)
L 2 (#80c4d0)
D 9 (#2fbf71)
L 3 (#826d60)
U 4 (#848001)
L 3 (#687032)
U 11 (#9768d1)
L 4 (#5b7e22)
D 5 (#25fb53)
L 7 (#ade7d2)
U 2 (#25fb51)
L 4 (#394662)
U 11 (#840763)
	</pre>
</details>
