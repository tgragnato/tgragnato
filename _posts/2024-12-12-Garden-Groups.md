---
title: Garden Groups
description: Advent of Code 2024 [Day 12]
layout: default
lang: en
tag: aoc24
prefetch:
  - adventofcode.com
  - deno.com
---

Why not search for the Chief Historian near the [gardener](https://adventofcode.com/2023/day/5) and his [massive farm](https://adventofcode.com/2023/day/21)? There's plenty of food, so The Historians grab something to eat while they search.

You're about to settle near a complex arrangement of garden plots when some Elves ask if you can lend a hand. They'd like to set up fences around each region of garden plots, but they can't figure out how much fence they need to order or how much it will cost. They hand you a map (your puzzle input) of the garden plots.

Each garden plot grows only a single type of plant and is indicated by a single letter on your map. When multiple garden plots are growing the same type of plant and are touching (horizontally or vertically), they form a **region**. For example:

```
AAAA
BBCD
BBCC
EEEC
```

This 4x4 arrangement includes garden plots growing five different types of plants (labeled `A`, `B`, `C`, `D`, and `E`), each grouped into their own region.

In order to accurately calculate the cost of the fence around a single region, you need to know that region's **area** and **perimeter**.

The **area** of a region is simply the number of garden plots the region contains. The above map's type `A`, `B`, and `C` plants are each in a region of area `4`. The type `E` plants are in a region of area `3`; the type `D` plants are in a region of area `1`.

Each garden plot is a square and so has **four sides**. The **perimeter** of a region is the number of sides of garden plots in the region that do not touch another garden plot in the same region. The type `A` and `C` plants are each in a region with perimeter `10`. The type `B` and `E` plants are each in a region with perimeter `8`. The lone `D` plot forms its own region with perimeter `4`.

Visually indicating the sides of plots in each region that contribute to the perimeter using `-` and `|`, the above map's regions' perimeters are measured as follows:

```
+-+-+-+-+
|A A A A|
+-+-+-+-+     +-+
              |D|
+-+-+   +-+   +-+
|B B|   |C|
+   +   + +-+
|B B|   |C C|
+-+-+   +-+ +
          |C|
+-+-+-+   +-+
|E E E|
+-+-+-+
```

Plants of the same type can appear in multiple separate regions, and regions can even appear within other regions. For example:

```
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
```

The above map contains **five** regions, one containing all of the `O` garden plots, and the other four each containing a single `X` plot.

The four `X` regions each have area `1` and perimeter `4`. The region containing `21` type `O` plants is more complicated; in addition to its outer edge contributing a perimeter of `20`, its boundary with each `X` region contributes an additional `4` to its perimeter, for a total perimeter of `36`.

Due to "modern" business practices, the **price** of fence required for a region is found by **multiplying** that region's area by its perimeter. The **total price** of fencing all regions on a map is found by adding together the price of fence for every region on the map.

In the first example, region `A` has price `4 * 10 = 40`, region `B` has price `4 * 8 = 32`, region `C` has price `4 * 10 = 40`, region `D` has price `1 * 4 = 4`, and region `E` has price `3 * 8 = 24`. So, the total price for the first example is `140`.

In the second example, the region with all of the `O` plants has price `21 * 36 = 756`, and each of the four smaller `X` regions has price `1 * 4 = 4`, for a total price of `772` (`756 + 4 + 4 + 4 + 4`).

Here's a larger example:

```
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
```

It contains:

- A region of `R` plants with price `12 * 18 = 216`.
- A region of `I` plants with price `4 * 8 = 32`.
- A region of `C` plants with price `14 * 28 = 392`.
- A region of `F` plants with price `10 * 18 = 180`.
- A region of `V` plants with price `13 * 20 = 260`.
- A region of `J` plants with price `11 * 20 = 220`.
- A region of `C` plants with price `1 * 4 = 4`.
- A region of `E` plants with price `13 * 18 = 234`.
- A region of `I` plants with price `14 * 22 = 308`.
- A region of `M` plants with price `5 * 12 = 60`.
- A region of `S` plants with price `3 * 8 = 24`.

So, it has a total price of `1930`.

**What is the total price of fencing all regions on your map?**

```ts
type Point = { row: number; col: number };

function findRegions(grid: string[][]): Map<string, Point[]> {
  const height = grid.length;
  const width = grid[0].length;
  const visited = new Set<string>();
  const regions = new Map<string, Point[]>();

  function isValid(row: number, col: number): boolean {
    return row >= 0 && row < height && col >= 0 && col < width;
  }

  function floodFill(row: number, col: number, type: string): Point[] {
    const region: Point[] = [];
    const queue: Point[] = [{ row, col }];
    
    while (queue.length > 0) {
      const current = queue.pop()!;
      const key = `${current.row},${current.col}`;
      
      if (visited.has(key)) continue;
      if (!isValid(current.row, current.col)) continue;
      if (grid[current.row][current.col] !== type) continue;
      
      visited.add(key);
      region.push(current);

      const directions = [[-1,0], [1,0], [0,-1], [0,1]];
      for (const [dr, dc] of directions) {
        queue.push({ row: current.row + dr, col: current.col + dc });
      }
    }
    return region;
  }

  for (let row = 0; row < height; row++) {
    for (let col = 0; col < width; col++) {
      if (!visited.has(`${row},${col}`)) {
        const type = grid[row][col];
        const region = floodFill(row, col, type);
        if (region.length > 0) {
          const key = `${type}_${regions.size}`;
          regions.set(key, region);
        }
      }
    }
  }
  
  return regions;
}

function calculatePerimeter(grid: string[][], region: Point[]): number {
  let perimeter = 0;
  const type = grid[region[0].row][region[0].col];
  
  for (const point of region) {
    const directions = [[-1,0], [1,0], [0,-1], [0,1]];
    for (const [dr, dc] of directions) {
      const newRow = point.row + dr;
      const newCol = point.col + dc;
      
      if (!grid[newRow]?.[newCol] || grid[newRow][newCol] !== type) {
        perimeter++;
      }
    }
  }
  
  return perimeter;
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const grid = input.trim().split("\n").map(line => line.split(""));

  const regions = findRegions(grid);
  let totalPrice = 0;
  
  for (const [_, region] of regions) {
    const area = region.length;
    const perimeter = calculatePerimeter(grid, region);
    const price = area * perimeter;
    totalPrice += price;
  }
  
  console.log(`Total price: ${totalPrice}`);
}

main().catch((err) => console.error(err));
```

Fortunately, the Elves are trying to order so much fence that they qualify for a **bulk discount**!

Under the bulk discount, instead of using the perimeter to calculate the price, you need to use the **number of sides** each region has. Each straight section of fence counts as a side, regardless of how long it is.

Consider this example again:

```
AAAA
BBCD
BBCC
EEEC
```

The region containing type `A` plants has `4` sides, as does each of the regions containing plants of type `B`, `D`, and `E`. However, the more complex region containing the plants of type `C` has `8` sides!

Using the new method of calculating the per-region price by multiplying the region's area by its number of sides, regions `A` through `E` have prices `16`, `16`, `32`, `4`, and `12`, respectively, for a total price of `80`.

The second example above (full of type `X` and `O` plants) would have a total price of `436`.

Here's a map that includes an E-shaped region full of type `E` plants:

```
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
```

The E-shaped region has an area of `17` and `12` sides for a price of `204`. Including the two regions full of type `X` plants, this map has a total price of `236`.

This map has a total price of `368`:

```
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
```

It includes two regions full of type `B` plants (each with `4` sides) and a single region full of type `A` plants (with `4` sides on the outside and `8` more sides on the inside, a total of `12` sides). Be especially careful when counting the fence around regions like the one full of type `A` plants; in particular, each section of fence has an in-side and an out-side, so the fence does not connect across the middle of the region (where the two `B` regions touch diagonally). (The Elves would have used the Möbius Fencing Company instead, but their contract terms were too one-sided.)

The larger example from before now has the following updated prices:

- A region of `R` plants with price `12 * 10 = 120`.
- A region of `I` plants with price `4 * 4 = 16`.
- A region of `C` plants with price `14 * 22 = 308`.
- A region of `F` plants with price `10 * 12 = 120`.
- A region of `V` plants with price `13 * 10 = 130`.
- A region of `J` plants with price `11 * 12 = 132`.
- A region of `C` plants with price `1 * 4 = 4`.
- A region of `E` plants with price `13 * 8 = 104`.
- A region of `I` plants with price `14 * 16 = 224`.
- A region of `M` plants with price `5 * 6 = 30`.
- A region of `S` plants with price `3 * 6 = 18`.

Adding these together produces its new total price of `1206`.

**What is the new total price of fencing all regions on your map?**

```ts
type Point = { row: number; col: number };

function findRegions(grid: string[][]): Map<string, Point[]> {
  const height = grid.length;
  const width = grid[0].length;
  const visited = new Set<string>();
  const regions = new Map<string, Point[]>();

  function isValid(row: number, col: number): boolean {
    return row >= 0 && row < height && col >= 0 && col < width;
  }

  function floodFill(row: number, col: number, type: string): Point[] {
    const region: Point[] = [];
    const queue: Point[] = [{ row, col }];
    
    while (queue.length > 0) {
      const current = queue.pop()!;
      const key = `${current.row},${current.col}`;
      
      if (visited.has(key)) continue;
      if (!isValid(current.row, current.col)) continue;
      if (grid[current.row][current.col] !== type) continue;
      
      visited.add(key);
      region.push(current);

      const directions = [[-1,0], [1,0], [0,-1], [0,1]];
      for (const [dr, dc] of directions) {
        queue.push({ row: current.row + dr, col: current.col + dc });
      }
    }
    return region;
  }

  for (let row = 0; row < height; row++) {
    for (let col = 0; col < width; col++) {
      if (!visited.has(`${row},${col}`)) {
        const type = grid[row][col];
        const region = floodFill(row, col, type);
        if (region.length > 0) {
          const key = `${type}_${regions.size}`;
          regions.set(key, region);
        }
      }
    }
  }
  
  return regions;
}

function calculateSides(region: Point[]) {
  const cellSet = new Set(region.map((p: Point) => `${p.row},${p.col}`));
  const leftEdges: string[] = [];
  const rightEdges: string[] = [];
  const topEdges: string[] = [];
  const bottomEdges: string[] = [];

  region.forEach(({ row, col }) => {
    const key = `${row},${col}`;
    if (!cellSet.has(`${row - 1},${col}`)) leftEdges.push(key);
    if (!cellSet.has(`${row + 1},${col}`)) rightEdges.push(key);
    if (!cellSet.has(`${row},${col - 1}`)) topEdges.push(key);
    if (!cellSet.has(`${row},${col + 1}`)) bottomEdges.push(key);
  });

  function parseKey(k: string) {
    const [X, Y] = k.split(",").map(Number);
    return [X, Y] as [number, number];
  }

  function createUnionFind(elements: string[]) {
    const parent: Record<string, string> = {};
    const size: Record<string, number> = {};
    elements.forEach((e) => {
      parent[e] = e;
      size[e] = 1;
    });
  
    function find(element: string) {
      let p = element;
      while (p !== parent[p]) {
        parent[p] = parent[parent[p]];
        p = parent[p];
      }
      return p;
    }
  
    function union(a: string, b: string) {
      const rootA = find(a);
      const rootB = find(b);
      if (rootA !== rootB) {
        if (size[rootA] < size[rootB]) {
          parent[rootA] = rootB;
          size[rootB] += size[rootA];
        } else {
          parent[rootB] = rootA;
          size[rootA] += size[rootB];
        }
      }
    }
  
    function countComponents() {
      const roots = elements.map(find);
      return new Set(roots).size;
    }
  
    return { union, countComponents };
  }

  function unionLineSegments(elements: string[], vertical: boolean) {
    if (!elements.length) return 0;
    const uf = createUnionFind(elements);
    const positions = new Set(elements);
    elements.forEach((element) => {
      const [X, Y] = parseKey(element);
      if (vertical) {
        const up = `${X},${Y - 1}`;
        const down = `${X},${Y + 1}`;
        if (positions.has(up)) uf.union(element, up);
        if (positions.has(down)) uf.union(element, down);
      } else {
        const left = `${X - 1},${Y}`;
        const right = `${X + 1},${Y}`;
        if (positions.has(left)) uf.union(element, left);
        if (positions.has(right)) uf.union(element, right);
      }
    });
    return uf.countComponents();
  }

  return (
    unionLineSegments(leftEdges, true) +
    unionLineSegments(rightEdges, true) +
    unionLineSegments(topEdges, false) +
    unionLineSegments(bottomEdges, false)
  );
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const grid = input.trim().split("\n").map(line => line.split(""));

  const regions = findRegions(grid);
  let totalPrice = 0;
  
  for (const [char, region] of regions) {
    const area = region.length;
    const sides = calculateSides(region);
    const price = area * sides;
    totalPrice += price;
    console.log(`Char: ${char}, Area: ${area}, Sides: ${sides}, Price: ${price}`);
  }
  
  console.log(`Total price: ${totalPrice}`);
}

main().catch((err) => console.error(err));
```

## Links

- [If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2024/day/12)
- [Solve Advent of Code 2024 with Deno and Win Prizes!](https://deno.com/blog/advent-of-code-2024)

<details>
	<summary>Click to show the input</summary>
	<pre>
YYYYYYYYYEJJEEEEEEEEEEEEEEGGGGGGGGGGGGGGGGGGCCCCCCCCCCWWCCLLLKKKKJKKKKKKFFFFFFFBBBBBBBBBBBAEEEEEEEEEEEEPPPPPPPPPOOOOZYYYZZZZZZAAUUUUUUUUUUUU
YYYYYYYYNEEEEEEEEEEEEEEEEEGGGGGGGGGGGGGGGGGCCCCCCCCCCCCCCCLLLLKKKKKKKKKKKFFFFFBBFBOBBBBBBBBBEEEEEEEEEEEPPPPPPOOOOOOOZZZZZZZZZZZAAUKUUUUUUUUU
YYYYYYNNNNYEEEEEEEEEEEEEEEGGGGGGGGGGGGGGGGGGCCCCCCCCCCCCCCCCLKKKKKKKKKKKKFFFFFFFFBBBBBBBBBBBBEEEEEEEEPPPPPPPOOOOOOOOOZZZZZZZZZUUUUUUUUUUUUUU
YYYNNNNNNNEEEEEEEEEEEEEEEEEGGGGGGGGGGGGGGGVVVVCCCCCCCCCRREEEEKKKKKKKKKKKKFFFFFFXFBBBBBBMMBBBBBEEEEEFEPPPPPPOOOOOOOOOOOZZZZZZZZZZUUUUUUUUUUUU
YYYNNNNNNWWEEEEEEEEEEEEEEEEGGGGGGGGGGGGGGGWVVVVVCCCCCCRRREEMEEKKKKKKKKKKKFFFFFXXFFBBBMBMMEEEEEEEEEEEPPPPPPPOOOOOOOOOOOZZZZZZZZZUUUUUUUUUUUUU
YYNNNNNNNNEEEEEEEEEEEEEEEEEEIIIIGGGGGGGGGWWWVVVVVCCCCCCEEEEMEKKKKKKKKKKKKKFFFFXXXBBBMMMMMEEEEEEEEEEEPPPPPPPPOOOOOOOOOOOOZZZZUUZUUUUUUUUUUUUU
YNNNNNNNNNNEEEEEEEEEEEEEEEEEIIIIGGGGGGGWWWWVVVVWWWWCWCCWEEMMMKKKKKKKKKKKKKFFFFXXXXXXMMMMEEEEEEEEEEEPPPPHHHHPPOOOOOOOOOOOZZZUUUUUUUUUUUUUUUUN
NNNNNNNNNNNNEEKEEKKKVEEEEEEIIIIIGIGGVGGGWCWWWWVVVWWWWCWWMMMMKKKKKKKKKKKKKKXXXXXXXXXXMMMMMEEEEEEEEEEPPHHHHHHPPOOOOOOOOOOOZZZUUUUUUUUUUUUUUUUN
NNNNNNNNNNNNAKKKKKKKKEEEEEEIIIIIIIGVVGGGWCCWWWWWWWWWWWWWWMMMMMMBBBKKKKBKKKBBXXXXXXXXMMXMMMEMMEEEEEQCPCHHHHHHPPOOOVVVOOOOZOOOOUUUUUUUUUUUUUNN
JJNNNNNNNNNNNKKKKKKKKEEEETTTTTTTTTGGGGGGCCCWWWWMWMMWWWWWWWMMMMMMMBBBKBBBKBBBBXXXXXXXXXXMMMMMMEEEQQQCCCHHHHHHHPVOOVVVOVOOOOOOOOOUUUUUUUUUUNNN
JJNNJNNNNNNJKKKKKKKKKKEEETTTTTTTTTTTIGGGCCCCWWWMMMMMWWWWWWWWWWMMKVBBBBNBBBBBBXXXXXXXXXXMMMMMMEEQQYQCCCHHHHHHHVVOOVVVVVOOOOOOOOOOUUQQUUUUNNNN
JJNNJNNNJJJJJKKKKKKKKKKKZTTTTTTTTTTTTTCGCCCCWWWMMMMMWWWWWWWWWWMVKVBBBNNNNBBBBXXXXXXBXNMMMMMMMMHQQQQCCCCHHHHHVVVHVVVVVVVVOOOFFFOOQQQQUNNNNNNN
JJJJJNNNJJJJJJKKKKKKKKKZZTTTTTTTTTTTTTCCCCCCCWMMMMWWWWWWWVWWVVVVVVBBBNXNNNNBBBBXXBBBMMMMMMMMMQQQQQQCCCCHHHHHHVVHVVVVVVVVOFFFFFOOQQQQUQNNNNNN
JJJJJNIJJJJJKKKKKKKKKKKKZZZZITTTTTTTTTCCCCCCCCMMMMWWWWWWVVVVVVVVVVVBVNNNNNNNBBBBXBBBBBMMMMMMMKQQQQQQCQCHHHHHHVVVVVVVVVVVVFFFFFOQQQQQQQQNNNNN
JJJJJJJJJJJJJKKKKKKKKKKKZLZZZTTTTTTTTTCCCCCCCCMMMMMMMWVWVVVVVVVVVVVVVNNNNNNNBBBBXBBBBBMMMMMMMQQQQQQQQQHHHHHMHHZVVVVVVVVVVVFFQOOOQQQQQQNNNNNN
JJJJJJJJJJCCCKVKKKKKKKZZZLLLLZZCTTTTCCCCCCCCCCMMMMMMMMVVVVVVVVVVVVVVNNNNNNRRRRRBBBBBMMRRRMMMMMQQQQQQQQMMMXMMMHVVVVVVVVVVVVQQQQQOOOQQQQNNNNNN
JJJJJJJJJJCCCCKKKKKKKZZZLLLLCCCCTTTTTTTTTCCLLMMMMMMMMMTVVVVVVVVVVVVVNNNNNNRRRRRRRBBBMMMMRMMMMQQQQQQQQQMMMMMMMEMVVVVVVVVVVVQQQHQQQQQQQQNNNNNN
JJJJJJJJJJCCCCIKKKKKKZLLLLLDLCCCTTTTTTTTTCLLLMMMMMMMMTTTTVVVVVVVVVVVNNNNNNNRRRRRRRBBMMMMMMMMQQQQQQQQQQMMMMMMMMMMVVVVVVVVVVVVQHHHQQQQNNNNNNNN
JJJJJJJJJCCCCCIKKKKKZZLLLLLLLCCCTTTTTTTTTCCLLLLLMMMEMTTTIIIIVVVVVVVVNNNNNNNNRRRRRRRRMMMMMMMMMQQQQQQQQMMMMMMMMMMMVVVVVVVVVVVVQHHHHQQQQNNNNNNN
JJJJJJDJCCCCCCCCKKKKKCLLLLLLLCCCTTTTTTTTTLLLLLLEEMMEITIIIIIIIIIVVVVVNNNNNNNRRRRRRRRJJJMMMMMMMBBQQQQQQMMMMMMMMMVVVVVVVVVVZIHHHHHHHQQQNNKKNNNN
JJJJJJDJCCCCCCCCCCCCCCLLLLSLLZZZZZTTTTTTTLLLLLLLEEEEIIIIIIIIIIIEVVNNNNNNNNRRRRRRRRRRMMMMMMMMBBBQBPQQQMMMMMMMMMVVVVVZVVVVZIHHHHHHHHHNNNNNNNNN
DJJJDDDCCCCCCCCCCCCCCLLLLLSSSZZZZZTTTTTTTLLLLLLLLEEIIIIIIIIIIIIIINNNNNNNNNRRRRRRRRRRMMMMMMMMMBBBBBBMMMMMMMMMMMVVVVVZZZVVZIIIIHHHHHHNNNNNNNNN
DDDJDDCCCCCCCCCCCCCCCLLLLLLSSZZZZZTTTTTTTLLLLLLIEEIIIIIIIIIIIIIIIINNNNNNNCNRRRRRRQQRMMMMMMMMMBBBBBBMMMMMMMMMMMMVVVVZZZZZZIIIIHHHHHHNNNNNNNNN
DDDDDDCCCCCCCCCCCCCCCLLLLSSSSZZZZZTTTTTTTLLLLLLIIIIIIIIIIIIIIIIIPIPNNCNNNCCCCRRRRRMMMMMMMMMMMBBBBBBMMMMMMMMMMMVVVVVZZZZZZIIIHHHHCNNNNNNNNNNN
IDDDDDCCCCCCCCCCCCCCLLNNLSSWSSZZZZTTTTTTTLLLLLIIIIIIIIIIQIQFQQIPPPPPPCNCCCCGCRRRMMMMMMMMMMMMBBBBBBBBBBMMMMMMMMMMMZZZZZZZZIIIHHHHHCNNNNNNNNNN
IIIDDIIYCCCCCCCCCCCLLLLNNSSWWWWZWWWEELLLLLLLLLIIIIIIIIIIQQQQQQPPPPPPCCCCCCCCCRRMMMMMMMMMMMMMBTBBBBBBBBMMMMMMMMMBBBZZZZZZZZIHHHNNNNNNNNNNNNNN
IIIDIIIICCCCCCCCLLLLLLLNNSSWWWWWWWWEEEEELLLLLIIDDDIIIIQQQQQQQQQPPPPPCCCCCCCCCRMMMMMMMMMMMMMMBBBBBBBBZZTMMMMMMMBBBBBBZZZZZZZNNNNNNNNNNNNNNNNN
IIIIIIIIIKCKCCCCXLLLLLNNNSSWWWWWWEEEEEEELLLLLIIDDDIIIQQQQQQQQQPPPPPPPPPCCCCCCRMMMMMMMMMMCCCMTBBBBBZZZZZPMMMMKMBBBBBBZZZCZZZNNNNNNNNNNNNNNNXX
IIIIIIIIIKCCCCCCXLLLLLNNNNNWWWWWWWEEEEELLELLLAIDDDDDDQQQQQQPPPPPPPPPPCCCCCCCCCCCMMMMMMMMMMMMMWBBBZZZZZLLLMMKKSBBSBBBZZZXZZZNNNNNNNNNNNNNNNXX
IIIIIIIIIXCCCCXXXXXLLLNNNYNVWWWWWEEEEEELEELLLLDDDDDDDQQQQQQPPPPPPPPPCCCCCCCCCCCCCMMMMMMMMMMMMWWZZZZZZZZLLULLKSSSSBSBXNNXXRXXNNNNNNNNNNNXXXXX
IIIIIIIIIXXCCXXXXXXXNNNNNYNVBBBWWBEEEEEEELLDDDDDDDDDDQQQQQQQQQPPPPPPPCCCCCCCCCCCCMMMMFFMMMMNMZZZZZZZZZLLLULLKSSSSSSXXXNXXXXXXNNNNNNNNNNXXXXX
IIIIIIIXXXXXXXXXXXNAANNNNYNNBBBBBBBEJJEAEDDDDDDDDDQQQQQQQQQQPQPPPPPPPCCCCCCCCCCCTFFFFFFFMMMMMMMUZZZZZZLLLLLLLSSSSSSSXXXXXXXNNNNNNNNNNNNXXXXX
IIIIIIIXXXXXXXXXXXNAANNNNNNNBBBBBBBJJJDDEDDDDDDDBBIIIQQQQQQQPPPPPPPCCCCCCCCCCCCUFYFFFFFNFMMMMEUUZZZZZAALLLLLLSSSSSSXXXXXXXXNNNNNNNNNNGGGXXXX
IIIIIIIIIXXXXXXXXNNNNNNNNNDNBBBBBBBBJJDDDDDDDDDBBIIIQQQQQQQQPQPPRPPCCCCCCCCCCCCFFFFFFFFFFFFMEEUUZZQZQLLLLQSSLSSSSSSSXXXXXXXNNNNNNNNGGVGGGXXX
IIIIIIIXXXXXXXXXXXXNNNNNNDDDBBBBBBBJJJDDDDDDDBBBBBBBQQQQQQQQQQPCRRRCCCCCCCCCCYYYYFFFFFFFFFFUUUUUUZQQQLLLLQSSSSSSSSSSXXXXXXNNNNNNNNNGGGGGXXGG
IIIIIIIIXXXXXXXXXXXNNNNNDDBBBBBBBBBBBDDDDDDDBBBBBBBBQQQQQQQQQQPCRRCCCCCCCCCCCYYYFFFFFFFFFFTTUUUUUQQQQQQLLQQQSSSSSSSXXXXXXXXNNNNNNNNGGGGGGGGG
IIIIIIXXXXXXXXXXXXXXNNNNDDBBBBBBBBBJJDDDDDDDBBBBBBBBBBQQQQQQCCCCCRCCCCCCCCCYYYYYYYFFFFFFFZTTTUUUQQQQQQQQQQQQQSSSSSXXXXXXXXNNNNNNNNNNGGGGGGGG
IIIIIIIXXXXXXXXXXXXNNNNNNNNVVBBBBBBJJJJDDDDDBBBBBBBBBBBQQQQQQQCCCCCCCCCYCCYYYYYYYYYFFFFFFZTTTZQQQQQQQSSSQQQQSSSSSSXXXXXXXXXNNNNNNGGGOGGGGGGG
IIINNNNZCXCXXCCXXXXNNNNVVVVVVBBBJJBBJJJDDKDKBBBBBBBTTBBQQQUQQQCCCCCCCYYYYCYYYYYYYYYFFFZZZZZZZZZQQQQSSSSSQQQQQQXXXXXXXXXXXXXNNNNNNGGGGGGGGGGG
NNNNNNNZCCCXCCXXXXXXNNNOVVVVVBBBJJBBJJJJDKKKBBFTTTTTTBBQQQQQQCCCCCCCCYYYYYYYYYYYYYYZQQZZZZZZZZZZSSSSSSSSQSQQQQXXXXXXXXXXKXXNNNNNNNGGGGGGGGGG
CCNNNNNNCCCCCXXXXFFXXNOOOOOVOOBBJJJJJJJJDKKKKKFFTTUTTBTQHQQHQCCCCCCCHYYYYYYYYYYYYZYZZZZZZZZZZZZZSSSSSSSSSSQQXXXXXXXXKXXKKKXNNNNNNNGGGGGGGGGG
CCCCNNNCCCCCXXXXFFFXXXOOOOOVOOJJJJJJJJJJKKKKKKFFFFFFFTTHHQHHQCCCCCCCHYYYYYYYYYYYYZZZZZZZZZZZZZZZSSSHSSSSSSSSFXXXXXXXKKKKKKKNNNNNNNNGGGGGGGGG
CCCCCCNCCCRCXXFXFFFOOOOOOOOOOOJAAAAAAJKKKKKKTTFFFFFFFTTHHHHHCCCCCCCCCYYYYYYYYYYYYYZZZZZZZZZZZZZZSSSSSSSSSSSSFFXXXSXXKKKKKKKNNNNNNNGGGGGGGGGY
CCCCCCCCCCRCCXFFFFFFFOOOOOOOOOOOAAAAAKKKKFFFFFFFFFFFFTTHHHHHHCCCHCCHHHQQYYYYYYYZZZZZZZEEEEEEEZZZZZZSSSSSSSSFFFSXSSXXKKKKKKKKKNNNNNNGGGGGGGGY
CCCCPCCCCCCXXXFFFFFFFOOOOOOOOOOOOAAAKKKKKFFFFFFFFFFFFTTTHHHHHCCCHCHHHHHQQYYYYYYYYYZHZZEEEEEEEZZZZSSSSSSSSSSSSFSSCSXHKKKKKKKKNNNNCCNGGGGGGGGG
CCCCCCCCCCCCXXFFFFTTFOOOOOOOOOOOAAAKKKKKKFFFFFFFTTTTTTTTHHHHHHHHHHHHHHHQQYYYYYYYYYZZZEEEEEEEEEZZZSSSSSSSSSSOSSSSSSSHHHKKKKKKKNCCCCJGGGLGGGGG
SCMMMCCCCCCVVFFFFFFTOOOOOOOOOOAAAAAKKKKKKFFFFFFFTTTTTTTTHHHHHHHHHHHHHHQQRRYYYYYBYYZEEEEEEEEEEEZZZSSSSSSSSSSOSSSSSSSSSKKKKKKRKKKSCCCCCLLLGGLL
SCMMMMMTCTTTTTTFFFTTOOOOOOOOOOAAAAAKKKKKKFFFFFFFTTTTTTTHHHHHHHHXHHHNNNQQRRRRUNNBYMMEEEEEEEEEEEZCCCCCSSSSSSSOSSSSSSSSKKKKKKKRKKKCCCCXCLLLLLLI
SSITMTTTTHTTTTTEETTTOOOOOOIOAAAAAAKKKKKKKFFFFFFFTTTUTTTTHHHHHHHHHHHHNNNRRRREEEEEEEEEEEEEEEEEEECCCCSSSSSSTSSSSSSSSSSSSKQQKKKKKSSCXCCJLLLLLLLL
IIITTTTTTTTTTTTTTTTOOOOOOOOOAQAQQAKKKKKKKKKKKKKTTKUUDTTTTHUHSHHHHHHNNNRRRRREEEEEEEEEEEEEEEEEESSCCEEESSSSTTSSSSSSSSSSSSSSCKEEJJJJJJJJULLLLLUU
IIIIMTTTTTTTTTTTTTTCOOXXOOOOQQAQQKKKKKKKKKKKKKKKKKUUDTTTUUUEUMMHHHHHNPRRRRREEEEEEEEEEEEEEEEEEISCCEEEESMMSSSSSSSSSSSSSSSSJJJEJJJJJJJJULLLLUUU
IIIIMMTTTTTTTTTTTTDOOOXXOOOPQQQQQQYYYNKKKKKKKNNKKKUUUUUUUUUUUMMHHHHPPPRRPRREEEEEEEEEEEEEESSSSSSCCEEEEEEESSSSSSSSSSSSSSSJJJJJJJJJJJJUUUUUUUUU
IIIMMMTRTTTTTTTTTTTEOOOOOOOIQQQQQQYYYNNNNNNNKNNKKKUUUUUUUHUUUMMMMPPPPPPPPFROSXXXXXXEEEEEEXSSSSSXCEEEEEEEBBSSSSSSSSSSSSSFJJJJJJJJJJUUUUUUUUUU
IIITMMMTTTTTTTTTTTTZZZZDDIIIIQIQQQYNNNZNNNNNNNNNKKKUUUUUUUUMMMMMPPPPPPPPFFRRSXXXXXXEEEEEESSSSSXXXXXEEEEEEQSSEEESESSSSSSSRJJJJJJJTNUUNUUUUUUU
IIITTTTTTLTTTTTTTTTZZZZIDIIIIIIQQYYNNNNNNNNNNNNNKKKUUUUUUUUUMMMMPPPPPPPPPSSSSSXXXXXXXXXXXXXSSSXXXXXXEEEEQQQEEEEEESSSSSJSSJJJJJJJTNNNNUUUUUUU
IIIIIIIILLLLTTTTTTTZZZZIIIIIIIIQIAYYNNNNNNNNNNNNNKNUUUUUUUUUMMMPPPPPPPPPPSSSSSSXXXXXXXXXXXXXSXXXXXXXXEEEQQQEEEEEESSSSSSSSJJJJJJJJNANNUUUUUUU
IIIIIIILLLLLZTTTTZZZZZIIIIIIIIIIIAANNNNNNNNNNNNNNNNUUUUUUUUUUMMMPPPPPPPPVVSSSSXDDDDXXXXXXXXXXXXXXXXXEEEEQQQQQEEEYSSSSCSDDDJJJJJJJNNNVVUUUUUU
IIIIIIILLLKLZZZZTZZZZZIIIIIIIIIIIAANNNNNNNNNNNNNNNNUUUUUUUUMMMMMPPPPPPPPPVSSSSXDDDDXXXXXXXXXXXXXXXXHHQCQQQQQQQQQYSSSSSDDDFFJJJJNNNNNVVVVVUUU
IIIIIIILSLKLZZZZTZZZZZZZIIIIIIIIAAANNNNNNNNNNNNNNUUUUUUUUUUUUMMMPPPPPPPPPSSSSSXDDDDXXXXXXXXXXXXXXKKKHQQQQQQQQQQQSSVSDDDDDDJJJJJJNNNNVVVVVVUU
IIIIIIILLLLLZZZZZZZZZZLZLIIIIIIZQAANNNNNNNNNNNNNUUUUUUUUUUUMMMMPPPPPPPPPPPSGGSGDDDDXXXXXXXXXXXXXXKHHHQQQQQQQQQQQQSVDDDDDDDDJRRRNNNNNFVVVVVVV
IIIIIILLLLLLZZZZZLZLLZLLLLIIIIZZZTTTTTNNNNNNNNNUUJUUUUUUUUUUNNMPPPPPPPPPPPSGGGGDDDDXXXXXXXXXXXXXXXHHHHHHQQQQQQQQQVVDDDDDDDRJRRRRNNNLVVVVVVVV
IIIIIILLLLLLLLZZLLLLLLLLLLIIIZZZZTTTTTZNNNNYNNJJUJUJUUUUUUUNNNPPPPPPPPPPPFCCCGGDDDDGXXXXXXXXXXHHHHHHHHHHHHQQQQQQQQQDDDDDDDRRRRNNNNNVVVVVVVVV
RRIPLILLLLLLLLLLLLLLLLLLLLLLIZTTTTTTTTTTZYYYYNJJJJUJJUUUUNNNNNNZPPPPPPPPDDDDDDDDDDDGXXXXXXXXXXXXHHHHHHHHHQQQQQQQQQDDDDDDDDRRRRRNNNRVVVVVVVVV
RRPPLLLLLLLLLZLLLLLLLLLLLLLLIZTTTTZTTTTTZYYYYJJJJJJJNJUUJNNNNNNNRRPPPPCCDDDDDDDDDDDGXCOXXXXXXXXDDHHHHHHHHHQQQQQQQQDDWDDDDRRRRRRRRRRVVVVVVVVV
RPPPLLLLLLLLLZLLLLLLLLLLLLLIIATTTTZTTTTTZZZZYJJJJJJJJJJJJJNNNNNNNRPPPCCCDDDDDDDDDDDGOOOXXXXXXXXXDDHHHHHHHHHQQQQQQQFDWWDDDRRRRRRRRRVVVVVVVVVV
PPPPELLLLLZZZZLLLLLLLLLLLLLLAATTTTZTTTTTTTTPJJJJJJJJJJJJJJJNNNNRRRRPCCCCDDDDDDDDDDDDVOOXXOIOOXXXDDHHHHHHHHQQQQQQQQWWWWVVDRRRRRRRRRRVVVVVVVVV
PWPPPPLLLLLLLZLLLLLLLLLLLLLLAATTTTZZZZTTTTTXXXJJJJJJJJJJJJSNNMMMMRRRRCCCDDDDDDDDDDDDOOOOOOOOOAXXKYHYYYHHHQQQQQQQQQWWWWWWWRRRRRIRRRRVVVVVVVVV
PPPPPYLLLLLLLLLOLLLLLLLLLLLAAATTTTZZZZTTTTTXXXXXJJJJJJJNNJNNNMMMMMRRRRRRDDDDDDDDDDDDOMOOOOOOOAXXKYYYYYYYYQQQQQQHWWWWWWWWWRIIIIIIRZVVVVVVVVVV
PPPPPPPPSSLLLLBBGGLLLLLLLTTTTTTTTTTZZZTTTTTXXXXXJJJJJJJNNNNNNMMMKKMRRRRRDOOOOOOOOODDDDDDDOOOOXXXKKYYYYYYMMQQQQHHWWWWWWWWIIIIITIZZZVVVJVVVVVV
PPPPPPPPPSSLBBBGGGGLLLLLATTTTTTTTTTZZZZZXXXXXXXXXJSJJJJJJJNNNMMMMMMRRRRRDOOOOOOOOODDDDDDDHOOOPPKKKYYYYYYYYQQQQHHHWWWWWWIIIIIIIIWZZZZZVVVVVVV
PPPPPPPPSSBBBBBBBBBBLLLLATTTTTTTTTTZZZZUUXXXXXXXJJJJJWJJJJBNNNNMMMMMRRROOOOOOOOOOODDDDDDDHOPPPPKKKKYYYYYYLQHHHHOHHIIWWIIIIIIIIIIIZZZVVVVVVVV
PPPPPPPPPBBBBBBBBBBAALAAATTTTTTTTTTZZZZZXXXFFFJJJJJJJJJJJJBNNBMMMMMMRRROOOOOOOOOOODDDDDDDOPPPPKKKKAYYYYYLLQHHOOOHHHIIIIIIIIIIIIIIZZZZVVVVVVV
PPPPPPPPPPBBBBBBBBAAAAAANTTTTTTTTTTZUZUXXXYBBFJJBBJJJBBBMMBLLDRRRRRRRRROOOOOOOOOOODDDDDDDIIIIIIIKKAYYKYYYLQAOOOOOOIIIIIIIIIIIIIZZZZZZZVVVVVR
PPPPBBBBPBBBBBBBBBBAAAAANTTTTTTTTTTZUUUXXXYBBBBBBBBJJBBBBBBLLDDRRROOOOOOOOOOOOOOOOOOOOHIIIIIIIIIAAAAYYYYYAAAPPOOOIIIIIIIIIIIIIIIZZZZZZZZVVVR
PPPPPTBBBBBBBBBBBBBAAAANNTTTTTTTTTTZZUUXXUBBBBBBBBBBBBBBBBDDDDDRDROOOOOOOOOOOOOOOOOAOOHIIIIIIIAAAAAAAAAAYAAAAAWWOWIIIIIIIIIIIOOIZZZNZZZZVVRR
PPPPPPBBBBBBBBBBBBAAAANNNNNNNNNNNNZZZUUUUUBBBMBBBBBBBBBBBBDDDDDDDROOOOOOOOOOOOPOOOOOOOHIIIIIIAACCCAAAAAAAAAAAAWWWWWBIIIIIIIIIOOUUZZZZZRZRRRR
PPPPPBBBBBBBBBBBBAAAANNNNNNNNNNNNNNUUUUUUUUBBMMBBBBBBBBBBBDOOOOOOOOOOOOOOOOOOOPOPOOOOOOIIIICICACUUUUAAAAAAAAAAWWWEWBIIIIIIIIIIUUUZUUZRRRRRRR
VPPPBBBBBBBBBBBBAAAAAARRRRRNNNNNNNNUUUUUUUUUMMMMBBBBBBBBDDDOOOOOOOOOOOOOOOOOOOPPPPOOOIIIIICCCCCCCUUUUAAAAAAAWWWWWEWBIIIIIIUIIIUUUUUUUURRRRRR
VPVPUBBBBBBBBBBBBAARRRRRRRRRNNNNNNNUUUUUUUUMAAMMBBBBBBDDDDDOOOOOOOOOOOOOOOOOPPPPPPOOOOIIIIICCCCCUUUMUUUAAAAVWWWEWEEIIIIITUUUUUUUUUUUUURRRRRR
VVVVBBBBBBBDDDYBAAARRRRRRRRRNUNNUUUUUUUUUUUMMMMMBBBBBBBDDDDOOOOOOOOORKKKKRRPPPPPPLPPOOIIIIICCCCCCCUUUUUAAWSSSSSEEEEEEEEETTTUUUUUUUUUUURRRRRR
VVVVBBBBBBDDDDDBOAARRRRRRRRRUUUUUUUUUUUUUUMMMMMMBBBBBBBDDDDDDDKNDRRRRKKKKKRRRPPPPLPPIIIIIIICCCCCCCUUUUUUUUSSSSSSEEEEEEETTTTEUUUUUUUUUURRRRRR
VVVVVVVNDDDDDDDDOOOORRRRRRRRRUUUUUUUUUUUUUUUMZXXBBBBBHBDDDDDDHKKKKKKKKKKKKRRRRPPPPPPIIIIIIICCCCCCUUUUUUUUUSSSSSSSSSEEETTTTTUUUUUUDDUUURRRRRR
VVVVVVVDDDDDDDDDOOOORRRRRRRUUUUUUUUUUUUUUUUUZZZBBBBBBBDDDDDDDDKJJKKKKKKKKKKKKKPPPPPPIIIIIICCCCCCCUUUUUUUNNSSSSSSSMMEMETTTTTTUUQQDDDDQUURRRRR
VVVVVVVVDDDDDDOOOOOORRRRRRUUUUUUUUUUUUUUUUUUUZZBBBUUUUDDDDDDDKKJJKKKKKQKKKKKKGPPPPPPPIIIIIIICCCCCUUUUUUUNSSSSSSSSMMEMEETTTTTUUQQQQDDQQQPPPPP
VVVVVVVHHDDDDDDOOOOOORRRRRRRJUUUUUUUUUUUUUUZZZZZXBZZZDDDDDDDDJJJJJKKKKKKKKKKKGFPVPPPPIIIIICCCCCQCUUUUUUUUTUSSSSSMMMMMMETTTTTUQQQQQDDQQQPPPPP
VVVVVVHHHDDDDDDOOOOOOORRRRRRJUSJUUJUUUUUHUZZZZZZXZZAAADDDDGGDJJJJKKKKKKKKKKKKKFPPFPPIIIIIIIIIIIIUUUUUUUUUUUUSSSSSSMMMUCCTCTTTTQAQQQQQQQPPPPZ
VVVVVVHHHDSDDOOOOOOOOORRRRRRJUSJJJJZUZZUZZZZZZZZZZZAADDDDDGGDJJJJKKKKKKKIKKFFFFPFFIIIIIIIIIIIIIUUUUUUUUUUUUUSSSSSSMUUUCCCCCCCCAAAQQQQQQQPPPP
VVVVVVHHNNNODOOOOOOOOORRRRRRJJJJJJAZZZZZZZZZZZZZZZDDDDDDDDGGGJJJKKKKKKKKKKKFFFFFFFIFIIIIIIIIIIIIUUUUUUUUUYUNSSSSSSMSUUCCCCCCCCAAAAAAAAQQQQQQ
VVNNNVHHNNNOOOOOOOOOOOOJAARJJJJGJJJZZZZZZZZZZPPPZZDDDDDDDDDGGJJJCCKKKKKKKKFFFFFFFFFFFIIIIIIIIIIUUUUUUUUUUUNNSSSSSSSSSCCCCCCCAAAAAAAAAAAQQQQQ
VVNNNNNNNNNOOOOOOOOOOOOJJJJJJJJJJJJZZZZZZZZZZPPPZZZDDDDDDDDDDDDCCCCKKKKKCFFFFFFFFFFFGGIIIIIIIIIUUUUNNUNNNNNSSSSSSSSSSSCCCCCCAAAAAAAAAAQQQQQQ
TTTNNNNNNNNOOOOOOOOOOOOJJJJJJJJJJJJZZZZZZZZPPPPPPPZDDDDDDDDDDDDDCCKKKKKCCCFFFFFFFFTFRRRRRIIIIAAAXUUNNNNNNNSSSSSSSSSSSZCCCCCCAAAAAAAAAZEQQQQQ
TNNNNNNNNNNNNOOOOOOOOOOJJJJJJJJJJJJZZZZZZZZPPPPPPPPDDDDDDDDDDDDCCCKCCCCCCCCCFFFFFFRRRRRRRRIIIAAAUUUUNNNNNISSSSSSSSSSZZCVCCCCAAAAAAAAAZZQQQQQ
TNNNNNNNNNNNNOOOOOOOOOOJJJJJJJJJJJJJJZZZZPPPPPPPPPPDDDDDDDDDDDDDCCCCCCCCCCCCFFFXFFIIRRRRRRAIIAAAFFUNNNIINISSSSSSSSSSSSSSCCCCCAAAAAAAZZZZQQQQ
TNNNNNNNNNNNOOOPOOOOOOOLJJJJJJJJJJJJJZZZZPPPPPPPPPPDHDDDDDDDDYDDDCCCCCCCCCCCXSFXXXXIRRRRRRAAAAAAAANNNIIIIIISSSSSSSSSSSSSSSSCOAAAAAAAZZAZQQQQ
NNNNNNNNNNNUUOPPPPPPOOOJJJJJJJJJJJJJZZZTTTPPPPPPPPPPHDDDDDDDDYDDDDCCCCCCCCCXXXXXXXXXRRRRRRRRAAAAAANNNIIIIIISSISSSSSSSSSSSSSSOAAAAAAAAAAAQQQQ
NNNNNNNNNNNUUPPPPPPPPOOJYYJJJJJJJJJJZZZTTTTTTTTTPPNNDDDDDDDDDDDDDCCCCCCCCCCXXXXXXXXXXXRRRRRAAAAAAANNNIIIIIIIIISSSSSSSSSSSSRRAAAAAAFFAAAAAQQQ
NNNNNNUNNNPUUUUPPPPPPPPPJJJJJJJJJJJJZZZTTTTTTTTTPPPNNDDDDDDDDDDDDCCCCCCCCCCXXXXXXXXXXXRRRRRAAAAAANNNNIIIIIIIIISSSSSSFFRRSRRRAAAEFFFFAAAAAAQQ
NNNNNNUUPPPUUUPPPPPPPPPBIBJJJJJJTTTTTTTTTTTTTTTTPPPNNNYDDDDDDDDDDCCCCCCCCCCCXXXXXXXXXXRRRRRAAARJNNNNNIIIIIIIISSSSSSSFFRRRRRRRREEFFFFTAFFVRQQ
NCNNNNUUPPPPUUUPPPPPPPBBBBCJBJJJTTTTTTTTTTTTTTTTPPPPYNYYDVDDDDDDDDCCCCCCCCCCXXXXXXXXXXRRRRRRRRRJJJNJJIIIIIIIISSSFFFFFFFFRRRRRRREFFFFFFFFRRRR
CCCNUNUUPPPPUUPPPPPPBPBBBBBBBJBBTTTTTTTTTTTTTTTTPPPYYYYYVVDDDDDDDDCCCCCCCCCCBXXXXXXXRRRRRRRRRRRJJJJJJJIIIISSSSSSFFFFFRRRRRRRRXFFFFFFFFFFRRRR
CCCCUUUUPPPPUUBPPPPPBBBBBBBBBBBBTTTTTTTTTTTTTTTTPPPYYYYYDDDDDDDEDCCCCCCCCCKKKKKKKKKKRRRRRRRRJJJJJJJJJJIJJJJSSSSSFFFFFRRRRRRRRXFFFFFFFGFRRRRR
CCCUUUUUUPPPPUUPPPPPTTTBBBBBBBBBTTTTTTTTTTTTTTTTPPPYYYYYDDDDDDDDDCCCRCCCCCKKKKKKKKKKKKKKKKKKKJJJJJJJJJJJJJOSJSSSFFFFFFRRRRRRRRRFFFFFRGRRRRRR
CCCUUUUUPPUUPUIPPPPPTTBBBBBBBBBYTTTTTTTTTTTTTTTTBBPPPYYYYYDDDDDDDCCCCCCCCCKKKKKKKKKKKKKKKKKKKJJJJJJJJJJJJJJJJSSSFFFFRRRRRRRRRRRQQFQQRRRRRRRR
CCUUUUUUUUUUUNGGTPPPTTBBBBBBBBBYTTTTTTTTTTTTTTTTGBGYYYYYYYYDDDDDDDCCCCCCBCKKKKKKKKKKKKKKKKKKKJJJJJJJJJJJJJJJJSSFFFFFRRRRRRRRRRRQQFQQRRRRRRRR
UUUUUUUUUUUUUGGGTPTTTTTBBBBBBBBBYYYYTTTTTTGGGGGGGGGGGYYYYYYDDDDDDNCCCCCCBBKKKKKKKKKKKKKKKKKKKJJJJJJJJJJJJJJJJJFFJFFFRRRRRRRRRRQQQQQQQRRRRRRR
UUUUUUUUUUUUUGGGTTTTSTBBBBBBBBBYYYYYTTTTTTGGGGGGGGGGGYYYYYYYDGDDDDFCCCCBBBKKKKKKKKKKKKKKKKKKKJJJJJJJJJJJJYJJJJJJJFRRRRRRRRHWRRRQQQQQQRRRRRRR
UUUUUUUUUGUUGGGGGGGSSTSSBBBBBBBBYYYYYYYYYWGGGGGGGGGGYYYYYYYGGGGGGGFCCCCBBBKKKKKKKKKKKKKKKKKKKJJJJJJJJJJJYYJJJJJJJJJRRRRRFRHWWWWQQQQQQRRRRRRR
UUUUUUUUUUAUGGGGGGSSSSSSSBBBBBBYYYYYYYYYWWGGGGGGGGGYYYYYYYYGGGFGFFFFFCBBBBKKKKKKKKKKKKKKKKKKKUJJJJJJJJJJYYJJJJJJJJJJRRRRRWWWWWWWQQQQQRLLRRRR
GUUUUUUUUUGGGGGGGGSSSSSSSSBBBBBYYYYYYYWYWWWGGGGGGGPPGJYYYYYYFFFFFFFFFFFFBBBBBBBBBBBKKKKKKKKKKUJJJJJJJYYYYYYJJZJJJJJJJRJRBWWWWWWWWWTQQLLLLRRR
GUUUUUUUUUUGGGGGGGGSSSSSSSSRYYYYYYYYYYWWWWWWGGGGGGGPGGYYYYYFFFFFFIFFFFFBBBBBBBBBBBCKKKKKKKKKKUUJJJJJYYYYYYYJJZJJJJJJJJJWWWWWWWWWWWQQQLLLLLLR
GGGUUUUUUUGGBBBCCBBSSSSSSSSRYYYYYYYMMWWWWWWWGGGGGGGGGYYYYYYFFFFFFFFFFFFFFFBBBBMMBBCKKKKKKKKKKUUUJJJJJJJYYZZZDZZJJJJJJJWWJWMWWWWWWWLLLLLLLLLL
GGGGUUUUGGGGBBCCCBBPPSSSSSSSSSOOMMYMWWWKKWWWGGGGGGGGEEEEYYYDFDDFFFFFFFFFFFBBBBBMBBCCCKKKKKKKKUUUJPPJJYYYYYZZZZJJJJJJJJJWJWWWWWWFWWWLLLLLLLLL
GGGGGUUUGGGGBBBCBBBBBBRSSSSSSBMMMMMMMWWKKKKWGGGGGGGEEEEEYYYDEDMMMFFJJFFCMMBBBMMMCCCCHCUUUUUUUUUUUUPPPPPYYYYZZZZZJJJJJJJJJJWCWWWWWLLLLLLLLXLX
PGGGGGGGGBBBBBBBBBBBBBBSSSSSSBBBMMMMMWKKKKKKGGGGGGGGEEEEEDRDDDMDDDDJJJJJMMMMMMMMMCMCHCUUUUUUUUUUUUPEPEYYYYYZZZZJJJJJJJJJJJWCCWWWWTTLLLLLXXXX
PGGGGGGDGGBBBBBBBBBBBBBIISSSSSBBMMMMMKKKKKKKKKGGGGGBBBBEBDDDDDDDDDDJJJJMMMMMMMMMMMMHHCUUUUUUUUUUUUPEEEEYYYYZZZZJJJJJJJJJJJCCCWWWWTTLLXXLXXXX
PPPGGDDDDBBBBBBBBBBBBNBISSSSSSBBBMMMMKKKKKKKKKKGGLBBBBBBBDDDDDDDDDDJJJJMMMMMMMMMMHHHHHHHHHUUUUUUUPPPEEEEEYZZZZZJJJJJJJCCCJCCCWWCCTTTLLXXXXXX
PPPGGDDDDDDBBBBBBBBNNNBBSSSSSSSBBBMMMMKKKKKKKKGGABBBBBBBBDDDDDDDDJJJJJJJJMMMMMMMMMHHHHHHHHHUUUUUUUPEEEEECZZZZZZZZJJJJCCRCCCCCCCCTTTTLLLXXXXX
PPPPGDDDDDBBBBBBBBBNNNNNNMSSSSSSBBMBBBBKKKKKKKKKKKBBBBBBBBBYDDDDDJJJJJJYYYYMMMMMMMHHHHHHHHHUEEUUUZEEEEEECZZZZZZZZZJJJCCCCCCCCTTTTTTTLXXXXXXX
PPPBPRDDDDKBBBBBBBBBNNNNMMSSSSSBBBBBBBBKKKKKKKKKKKBBBBBBBBBDDDDDDJJEEEEYYYYMMMMMMMHHHHHHHHUUEEEEEEEEEEEZZPZZZZZZZZJJAECCCCCCCHHHTTTTWXWXXXXX
PPPPPRRDDVBBBBBBBMMMMNNNMMMSSSZZBBBBBBBKKKKKKKKKKBBBBBBBBBBDDDDDDJEEEEEEEEYMAMMMMMMMMHPHHHUUUEEEEEEEEEEZZZZZZZZZZZZJAACCCCCHHHHTTTTTWWWWWWWX
PPPPPPPPVVVBBBBBBNNMMMMMMSSSZZZZZBBBBBKKKKKKKKKKKBBBBBBBBBBBDDDDDJJEEEEEEEEMAMMMMMMMMMMMMHHUUEEEEEEEEEZZZZZZZZZZZZZAAAACCCCCHHHTTTTTMWWWWWWX
PPPPPPPPPVVBBBBBBMMMMMMMMMMZZZZZZZZZBBKKKKKKKKKKBBBBBBBBBBBBDGDQEEEEEEEEEEEEEOMMMMMMMMUCUUUUUUEEEEEEEEEZZZZZZZZZZZZBAAAAACCCCAHHHHWMMWWWWWWX
PPPPPPPPNDDDBBBBMUMMMMMMMMMMMZZZZZZBBZKKKKKKKKKKBBBBBBBBBBBLDBUEEEEEEEEEEDDDDDDDDDDDDMUUUUUUUEEEEEEEEEEZZZZZZZZZZZAAAAAAAACAZAHHHWWWWWWWWWWW
PPPPPPPPDDDDGBBBMMMMMMMMMMMMMMZZZZZZZZKKKKKKKKKKBBBBBBBBBBBBBBBEEEEEEEEEEDDDDDDDDDDDDJUYUUUUUEEEEEEEZZZZZZZZZZZZWWAAAAAAAAAAAAAAWWWWWWWWWWWW
PPPPDDDDDDDGGBBMMMMMMMMMMMMMMMZZZZZZZZZZKKDKKKKKBBIBIIBBBBBBBQEEEEEEEEEEEDDDDDDDDDDDDJJYUUUUUEEEEEEEZHZZZZZZZZZZZAAAAAAAAAAAAAAAWWWWWWWWWWWW
PDDDDDDDDDDDDDDFFMMMMMMMMMMMMZZZZZZZZZZZKKKKKKKDDIIIIIIBBBBBBBEEEPPEEEEEEEEEEEEDDDDDDYYYUUUUUUEEEEEEZZZZZZZZZZZZZAAAAAAAAAAAAAAAGWWWWWWWWWWW
PDDDDDDDDDDDDFFFFMMAAAMMMMPMZZZZZZZZZZSZZWKDDDDDDIIIIIIIIBBBBBHEEPPPEEEEEEEEEEEDDDDDDYYYUUUUUEEEEEEEZZZZZHHHHHHHHHHAAAAAAAAAAAAAWWWWWWWWWWWW
DDDDDDDDDDDDDFFFFFFAAAMMMMMMZZZZZZZZZZZZZZZZDDDDDIIIIIIIIIBBLBPPVPPPEPEEEEEEEEEDDDDDDYYYYYUUUUUHEEEEZZZZZHHHHHHHHHHHAAAAAAAAAAAAWWWWWWWWWWWW
DDDDDDDDDDDDXXFFFFFAAAMMMMMMZZZZZZZZZZZZZZDDDDDDDIIIIIIIIILLLPPPPPPPEPPPPEEEEEEDDDDDDYYYYYUUUHHHHHZZZZZZZHHHHHHHHHVVVAAAAAAAAAAAWWWWWWWWWWWW
DDDDDDDDDDDSXXFFFFFAAAMMMAZZZZZZZZZZZZZZOZIIDZZZZZIIIIIIIILLLPPPPPPPPPPPPPPPEELDDDDDDYYYYYUUUHHHHHZZZZZZZHHHHHHGHVVAAAAAAAAAAAAAAWWWWWWWWWWW
DDDDDDDDDDXXXXFFFFAAAAAMAAAZZZZZZZZZZZZZOOIIDDZZIIIIIIIIIILLLPPPPPPPPPPPPPPPGLLLGDDDDYYYYYUHHHHHHHZZZZZHHHHHHHHHHVVVAAAAAAAAAWWAFWWWWWWWWWWW
DDDDDDDDDDDXXXFFFAAAAAAAAAAZZZZZZZZZZRRIIIIIIIZZZZIIIIIIIILLZPPPPPPPPPPPPPGGGGLLGDDDDYYYYYUUUUUHHHZZZZZHIHHHHHVVVVVAAAAAAAAAAAAFFWWWWWWWQWQW
DDDDDDDDDDDXXFFFFAAAAAAAAAAAAAAZZZZZZIIIIIIIIIIIZZIZIIIIIILZZZPPPPPPPPPPJPGGJGLGGGGGYYYYYYUUUUUIIIPIIIHHIIHHHVVVVVUUUUAAAFAAAFFFFWWWWWWQQQQQ
NDDDDDDDDDDXXFFFFAAAAAAAAAAAAAAZZZZIIIIIIIIIIIIIIZIZZIIIIIZZZXPPPPPPPPPJJPDGGGGGGGGGNGYYYYYUUUUUIIIIIIHHIIIIHVVVVUUUUUUUAFAAAFFFFFFWQWQQQQQQ
DDDDDDFDDFFFXFFFAAAAAAAAAAAAAAAXZXZZIIIIIIIIIIIIZZZZZIIIIZZZZPPPPPPPPPPPPPGGGGGGGGGGGGYYYYYUUJJUIIIIIIIIIIIVVVVVUUUUUUUUAFFFFFFFFFQQQQQQQQQQ
VVDDDDDFFFFFXFFFFAAAAAAAAAAAAXXXXXXZIIIIIIIIIIIIIZZZZZZZZZZZZPPPPPPPPPPDGGGGGGGGGGGGGGYYYYYUJJJJJJIIIIIIIIIIIIVVUUUUUUUUUUUFFFFFFFFQQQQQFFQQ
VVDDDFFFFFFFFFFFFAAAAAAAAAAAXXXXXXIIIIIIIIIIIIIIZZZZZZZZZZZZZPPPPPPPPPPPDGGGGGGGGGGGGGGGYYYUUUJJJJIIIIIIIIIIIVVVUUUUUUUUUUUUFFFFFFFQQQFFFFFQ
VVFFFFFFFFFFFFFFFAAAAAAAAAAAXXXXXXIPIIIIIIIIIIZZZZZZZZZZZZZZZDDDDDDPPPPDDDGGGGGGGGGGGGGGYYYYYJJJJIIIIIIIIIIIVVUUUUUUUUUUUUUUUFFFFFFQFFFFFFFF
FVFFFFFFFFFFFFFFFAAAAAAAAAAXXXXXXXIIIIIIIIIIIIZZZZZZZZZZZZZZZDDDDDDDPDDDDGGGGGGGGGGGGGGGYYYYJJJJJJJIIIIIIIIIVUUUUUUUUUUUUUUUUFFFFFFFFFFFFFFF
FFFFFFFFFFFFFFFFFFAAAAATTAXXXXXXIIIIIIIIIIIIIIIZZZZZZZZZZZZZZDDDDDDDDDDDDDGGGGGGGGGGGGGGJJJJJJJJJJJJIIIIIIIIVVVUUUUUUUUUUUUUFFFFFFFFFFFFFFFF
	</pre>
</details>
