---
title: Linen Layout
description: Advent of Code 2024 [Day 19]
layout: default
lang: en
tag: aoc24
prefetch:
  - adventofcode.com
  - deno.com
  - en.wikipedia.org
  - www.youtube.com
---

Today, The Historians take you up to the [hot springs](https://adventofcode.com/2023/day/12) on Gear Island! Very [suspiciously](https://www.youtube.com/watch?v=ekL881PJMjI), absolutely nothing goes wrong as they begin their careful search of the vast field of helixes.

Could this **finally** be your chance to visit the [onsen](https://en.wikipedia.org/wiki/Onsen) next door? Only one way to find out.

After a brief conversation with the reception staff at the onsen front desk, you discover that you don't have the right kind of money to pay the admission fee. However, before you can leave, the staff get your attention. Apparently, they've heard about how you helped at the hot springs, and they're willing to make a deal: if you can simply help them **arrange their towels**, they'll let you in for **free**!

Every towel at this onsen is marked with a **pattern of colored stripes**. There are only a few patterns, but for any particular pattern, the staff can get you as many towels with that pattern as you need. Each stripe can be **white** (`w`), **blue** (`u`), **black** (`b`), **red** (`r`), or **green** (`g`). So, a towel with the pattern ggr would have a green stripe, a green stripe, and then a red stripe, in that order. (You can't reverse a pattern by flipping a towel upside-down, as that would cause the onsen logo to face the wrong way.)

The Official Onsen Branding Expert has produced a list of **designs** - each a long sequence of stripe colors - that they would like to be able to display. You can use any towels you want, but all of the towels' stripes must exactly match the desired design. So, to display the design `rgrgr`, you could use two `rg` towels and then an `r` towel, an `rgr` towel and then a `gr` towel, or even a single massive `rgrgr` towel (assuming such towel patterns were actually available).

To start, collect together all of the available towel patterns and the list of desired designs (your puzzle input). For example:

```
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
```

The first line indicates the available towel patterns; in this example, the onsen has unlimited towels with a single red stripe (`r`), unlimited towels with a white stripe and then a red stripe (`wr`), and so on.

After the blank line, the remaining lines each describe a design the onsen would like to be able to display. In this example, the first design (`brwrr`) indicates that the onsen would like to be able to display a black stripe, a red stripe, a white stripe, and then two red stripes, in that order.

Not all designs will be possible with the available towels. In the above example, the designs are possible or impossible as follows:

- `brwrr` can be made with a `br` towel, then a `wr` towel, and then finally an `r` towel.
- `bggr` can be made with a `b` towel, two `g` towels, and then an `r` towel.
- `gbbr` can be made with a `gb` towel and then a `br` towel.
- `rrbgbr` can be made with `r`, `rb`, `g`, and `br`.
- `ubwu` is **impossible**.
- `bwurrg` can be made with `bwu`, `r`, `r`, and `g`.
- `brgr` can be made with `br`, `g`, and `r`.
- `bbrgwb` is **impossible**.

In this example, `6` of the eight designs are possible with the available towel patterns.

To get into the onsen as soon as possible, consult your list of towel patterns and desired designs carefully. **How many designs are possible?**

```ts
function canMakeDesign(
  design: string,
  patterns: Set<string>,
  cache: Map<string, boolean> = new Map()
): boolean {
  if (design === "") return true;
  if (cache.has(design)) return cache.get(design)!;

  for (const pattern of patterns) {
    if (design.startsWith(pattern)) {
      const remaining = design.slice(pattern.length);
      if (canMakeDesign(remaining, patterns, cache)) {
        cache.set(design, true);
        return true;
      }
    }
  }

  cache.set(design, false);
  return false;
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const [patternsStr, ...designs] = input.split("\n").filter(line => line.trim());

  const patterns = new Set(
    patternsStr.split(",")
      .map(p => p.trim())
      .sort((a, b) => b.length - a.length)
  );

  let possibleDesigns = 0;
  const cache = new Map<string, boolean>();

  for (const design of designs) {
    if (canMakeDesign(design.trim(), patterns, cache)) {
      possibleDesigns++;
    }
    if (cache.size > 10000) cache.clear();
  }

  console.log(`Number of possible designs: ${possibleDesigns}`);
}

main().catch(console.error);
```

The staff don't really like some of the towel arrangements you came up with. To avoid an endless cycle of towel rearrangement, maybe you should just give them every possible option.

Here are all of the different ways the above example's designs can be made:

`brwrr` can be made in two different ways: `b`, `r`, `wr`, `r` **or** `br`, `wr`, `r`.

`bggr` can only be made with `b`, `g`, `g`, and `r`.

gbbr can be made 4 different ways:

- `g`, `b`, `b`, `r`
- `g`, `b`, `br`
- `gb`, `b`, `r`
- `gb`, `br`

`rrbgbr` can be made 6 different ways:

- `r`, `r`, `b`, `g`, `b`, `r`
- `r`, `r`, `b`, `g`, `br`
- `r`, `r`, `b`, `gb`, `r`
- `r`, `rb`, `g`, `b`, `r`
- `r`, `rb`, `g`, `br`
- `r`, `rb`, `gb`, `r`

`bwurrg` can only be made with `bwu`, `r`, `r`, and `g`.

`brgr` can be made in two different ways: `b`, `r`, `g`, `r` **or** `br`, `g`, `r`.

`ubwu` and `bbrgwb` are still impossible.

Adding up all of the ways the towels in this example could be arranged into the desired designs yields `16` (`2 + 1 + 4 + 6 + 1 + 2`).

They'll let you into the onsen as soon as you have the list. **What do you get if you add up the number of different ways you could make each design?**

```ts
function countWays(
  design: string,
  patterns: string[],
  cache: Map<string, number> = new Map()
): number {
  if (design === "") return 1;
  if (cache.has(design)) return cache.get(design)!;

  let total = 0;
  for (const pattern of patterns) {
    if (design.startsWith(pattern)) {
      const remaining = design.slice(pattern.length);
      total += countWays(remaining, patterns, cache);
    }
  }

  cache.set(design, total);
  return total;
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const [patternsStr, ...designs] = input.split("\n").filter(line => line.trim());

  const patterns = patternsStr.split(",")
    .map(p => p.trim())
    .sort((a, b) => b.length - a.length);

  let totalWays = 0;
  const cache = new Map<string, number>();

  for (const design of designs) {
    const ways = countWays(design.trim(), patterns, cache);
    totalWays += ways;
    if (cache.size > 10000) cache.clear();
  }

  console.log(`Total number of different ways: ${totalWays}`);
}

main().catch(console.error);
```

## Links

- [If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2024/day/19)
- [Solve Advent of Code 2024 with Deno and Win Prizes!](https://deno.com/blog/advent-of-code-2024)

<details>
	<summary>Click to show the input</summary>
	<pre>
bubgb, buburr, rrrbrr, wubgbb, gggbw, wrr, rrbu, wbw, gubr, bruu, bub, rww, gguubguu, wwb, ugbrrw, burur, rgwwb, ugu, bbu, rug, uwwb, ggw, bgb, wbru, uuwrrrg, uuug, gur, urguw, ugru, urbwggg, bru, rgug, guu, g, grwr, rbuwb, bbr, wgbr, gubrw, rwrgb, rbw, buw, gbg, wrw, wrbrwg, gbrg, urgubgb, uwg, wggubgr, rgrub, bwuugrwr, gurgu, bur, wwgwwbu, bbbrgu, wbggru, wbuur, rgubg, rwww, br, wbr, wgrrggb, wwuubgg, gbu, uuu, rwbwb, uwugrgu, bugrgwwu, wgwbb, bgwwggr, wgbb, rubbr, ugbrbgub, urg, uruuwu, ggb, rgwwg, guw, uruur, wbgbr, wrbrbw, rrrggr, rgw, rgbgwbu, uru, bwgbgr, wwgg, ggbrgwr, rwgurruu, uwbr, rwur, uu, uugr, wg, wbuwuwr, bwbu, bgw, gbr, uugb, bubw, wrg, ugw, rwb, ww, rrggrb, ugww, bbw, wbb, ubbrwr, wrwr, uwbb, uwr, ugrg, ruu, gbw, gwurr, wwu, bbg, wbbw, rgww, wwbb, wuu, wwbwr, uugbwb, bbrw, wguwr, rugr, rgb, ubb, wrgr, grrrw, rrb, burwr, www, uub, wgbbwbb, rw, rbww, uwu, ugb, ugrw, ruuuubb, uwbbw, rgg, wuw, grbbw, gbru, rb, gwrwr, bgg, gwuug, rgu, bugrbb, rrrr, rubwu, bgbbu, wbgubbg, wwwuw, wwrb, uwgbg, rur, rbbu, rbr, rrrbwb, rgggugu, bgwuu, rbwgrbg, gbwwuuru, gg, wgru, grrub, ruurb, urgw, wuurgw, bbbgwww, gwb, gwwb, rbuw, ruuwwg, uuw, brggrw, wubgwu, ubw, ubgb, wbu, bubrrgww, wbg, brw, bggrrr, uwubw, wbbugu, bgu, rrbbww, ru, wbgbw, ggrwrugg, bwbwg, gbb, urwrrb, uuwuggr, gugw, urb, wbwbwrrb, grbwb, wwwgg, gwu, buwg, ug, gruwgub, gbgr, bgbbuwwr, rbb, ubuw, gww, b, wggg, wuwgub, gbbbbu, wgr, gwwbw, wurrb, gurw, rbur, brbrbrb, rgbu, ggbrb, rgggr, wwbrr, rrw, wur, urwwbgu, wubuwu, wwbu, bg, bbru, bgrguuuu, uwbru, bruburb, ugwrb, rbrwbgg, ugwru, ubu, ggwb, buu, burrbw, wr, rrrw, uuru, uurbbbb, wubrgb, gb, wbwwrug, rbu, bwrrww, uuwggrbb, gw, brg, uwrr, rbwb, rub, gggwr, gbbg, gwgb, rubggu, gbgu, wgb, brwurg, rrru, burg, rrbb, ubugub, uuwg, uwgu, grr, guwwww, gwg, grggg, ggwggbgb, wru, uurwwu, gbbbr, bbbb, brgu, buwgrr, grwgwu, buwbw, rru, gubwg, gub, wwr, gubgwg, buub, wrbbuw, urr, wu, bbgb, bw, uubb, ggrwrguw, wbgrrb, rrbuurb, ubwug, ggwww, bgrr, rgrgb, ur, brwuwbg, bubb, grruwwub, wurr, buuwbw, r, bu, brb, bgrurwb, bggrwg, rwu, wrwwbb, wbbugr, gggb, brbb, bwbuw, rwgurb, bgbuurg, buwb, uugu, ubwbggb, rruw, bggg, ubbrbw, grrr, bbruug, wgg, rwubgrgb, uubrruwr, wgbrb, bwb, gubrb, gwuurg, ggbr, wwg, ruww, bb, ggg, rgr, bww, gwburgwr, urw, gbbgb, rwg, bgr, ub, gwwug, bug, rr, rg, wrb, rrg, gggw, wwbbg, uwb, wwrw, bgbwrbg, ggrr, wug, wwuwww, bgrubw, gru, bubbgw, gwuww, gbggbb, gwbu, bgbugrb, wwrbr, ubbgr, brub, ugbwubug, uww, rugrb, rgbuu, w, wrrbw, rbg, rururwb, grbrwbwu, brurwb, ubgwrrr, ruw, bbrrrwu, ruwgw, rbrwu, rrgurg, ururbww, bbb, ugbb, bwu, wugb, uwrurwg, bwg, bggw, gug, bbgubw, uug, uur, grwwu, rrug, brgbbru, ggu, brwbu, grw, rwr, ubwbu, wbbb, wubrwg, gbgg, ugg, bbggub, gwr, ggub, wwgur, bwrw, rwbuw, wb, rbbr, urbb, gbggg, grg, ggwu, uwurgg, wbugug, ugr, rrbwr, gwuur, uwbgr, brrb, ubg, uuugwu, gr, bgbb, wggr, ggr, bwguw, ugubwbu, wgw, grb, wub, wbuww

wwwuwbugbggrwuwgwgbwwwgguggguguwbgguurbbbwgrubrgb
wbbwbrbrwwugbruubgurgwbbwwgugbbbbrbbbrwwwgrrrbgubwubwru
gwwgwbgbgbuugwurgggwrubrruuwgbwgwrgwrbwrugwwrrugrwgu
gbubwbrgugwrurggguruugbrwugrgwggrubgruuwubruuwwguggu
wwgrrgwguwugbgrbbguwguwubwuwwwuggguubuwgrwrrwwwuwbr
bbuwgwgwwwgrrgrbrwrwrbwwwbgrbrgwwwgbrgrubgwuguggwrburu
wuwbgbbrurwwwrbbubuwgwgwubbugrrurguurgguburbuugrw
gggwguwbbuwubgurrgwugwbubwbwguuwwbrwgwrgrbrggugu
gbbgggwgbbgwgguuurbbwggrrbugrggruwgrwgbgrbbruwwwrrwwgrww
grbrwggrurrwwrbuggugurburrrbrrrwwugrugrrugwruwrrgbg
urgbruwwurugbbbwgbruwgrubrrugrbwwruwuwrrubuggrwrgugbgbg
wgrbugrbbgrugrrbubgbubuwbrbuuuwuubrgwrwwrgwwburbrgwrbw
rururrurwwuurrrrgwgbugwubwwbubuuuwwbwbruuwugrbwuwbugw
uwwwwuurbwbguwruwwgggbggubbwwggbrbgrwgubrubguwwu
wgrwbgubgurbwrwwwrurwrbuggggurbrwwbrwubwrwgugurgururwu
gwwrbwwgwurgrwwguggubwwuuggrwbbggbuguwgrubbwbrwub
ubbrwwwbgbuuubggguwwbrggruwbrrwgrbggwuuuuggwu
wwwbrggrrgubwgrrbwwrrurrrugrubgrrgggbrwbrbuuwgrrrruwr
rgwbrbrugugwrbbrrbwrwrrgwggwurugurwwgugbrrbbbww
grwgggbrwurrrgbrwrrrubuwbwuwguuwbuuguuguwwbggwbguwrbb
wbgbgrbburgubbgbwubrwgrgruggwbrrguugurgrgugubgurrwgrugbwbb
rurgrbrrwrrgrgbrrubrgbwbbgwrrrbbgbgwwubgrrrugru
gwbrbguwgrgwwgurrgwggbbgrurguggbwwbbrugrww
wubgrbwuwuugbrubbgbwurbuubbwwuuuwwgrruuuuwb
buubbgggrrbrrrrwgrggbwgwgwubgggbgwbgrgbwgwwwrwrgrgbubgrbw
buuggwugrggbbubrbrubbrgbgbrwrgwwbruguugrwwbrwrggwggbgbuub
ugrbrbrururrruuwrwgwgrwgbrwbrgwwururgrgwurbbgbrubrubww
buwwgubrugubbbwugggwbwbgubuuruubrwgbgrbwwrbwgu
bbbuuuwrwrbgrugwgwggrwgbwgrgburbgbgggbbggur
bggwwwubbrbrrubggrbgbwruwgggwrwuuguwwrwwrrrurgguuruu
gguwwbbrwrwgbruggrwrbbuwruwuuugurrbgbrruwburwgbrwg
gwurruwbbggwbrgbubwwwrwgrbrbwubugbbruuwwgr
gwrrgrwbgbugbgrruwuuuugwbbuwgggwbgbuwgurgrwwuwg
rrgwwgguubguurwbrbwrwggbbrbgrrbwgwwbbgwgu
bwrwugrwwguwrwugrbuwrbgwbubggbgrbubgwwgu
wrrrurruurbbuwugrwwrwggbgwubwwuurrbbuuuuuwwrwgu
rbguuwwbgbbwgrurwwwwuwrbgurwgbwgrbbubugrbguru
gwuwwruugwggwgbbuwrgbgbwwrbwrwwwubuwggwggwwgu
ubrrbggrbwburbuubwrrurrggrwuwuwrubbggrubgwwburgbb
rruwrurbwgurwbbgruruuubrrrrgbrrbgwurgwbrrgwwur
bwwgrwwwbgbbrbbuugrwbrggugrrruggwwwwwgbrbrwwrru
rrgwgwwwrgurwuwrwburwbrwurrgwrggbgrrrwbwwrurwburgwu
bugbrbrrwbwbgbwugggrurubbwrwgwgbguugrwuwbguwruwrrruubwurbw
rurrwggubbrbrubgrgrrgrrwrruuggwwuurwuurwrbrgguuwbgw
uruwbgrwrurbwbgrubgubbruwguugwrbbgggbbbbwgu
wwubugrgwwugrwwbwruuubrgggwguwwgrwubwgrugwubrwwurruuubbw
uubgbrgubrwwbubbuguwwurrbgurbbgwbuwuwubuubwubwugw
gugbwbwgggwbbbuurrgwurugwrrgubrrubbbguurgrbgbuwgrwrbbbugwgu
rurbrwbruurrgrwbuwuguwurwubgrgbbgugubbuburgguwugugbug
gbgwrwgrbuugrgruuwrrrrbguggubgbrbwwrugbbrwuggrwwguug
rbrbrgugbwwgggurrrwrbwwwggguguubwuugrwwgubbbuwbbgbbrbuwu
rwwgrubuugwrwbgbuwuwububwrrrrrbgubbrrwggrwggu
bbugwuurgbubgwwrgbugwbgugrburrwrbgwurgrgwgu
grwgbbwugwbrwwgrwwrwwwbbggrbuugwuuurwwurgbbuwrrw
bgwrbrbrwrwugrrububwbuubwwrbrbgggrrubgwbggu
gwwuuwbgugwugguwrrgrrwwwbrwbwbrrbuwwrbugbubwurb
ubrggrrwbguuwrburwbgbrggwrbbrbwuubbbwgurggwwbgrrggrwgbrg
uwrwwwbuugbwugrgbrggbgbwruwguurwbbwbbwbbrubrwbwgwuwg
urggubbgrrbgwrrurwgurgrubuubwburgggrgbgwrbbrgbgwrrubr
uugrrrbwbruwugbbruuubrwggurwwgrgrrrbuguwwwbwwubwburuugg
buwgbbbrugrbgbuwrbbrgugbwuuwruwgurbggbbwgwgwwgbwuuwgr
uuurgwgrguuuubwubbrbbubgggbrgugwururwuuurbwwwububgru
rgugrbwuwguwgugbwgrwuubrwubbrwubwbwubugwbgubguruwuwuuuu
buwgwgbbbrubuugbwurwuububwgbuwgrugbrrwrbuwrg
uwurwrwuubbwrgwuwbruruwuwwgbgbuuurrwgwbwbgwgu
gwrbubruwwbwbwwgwbuwwggwwwbbrrgbwuwrgubbbgrrrbgur
urggwbrggguwrrguurggwrrurrugrgububgwwgwrwbuuuwwgwgrb
uugubwwrwgrgwbwrurgurrubuwgrbgrguurwbrwggugggwbrwurwrwrbwg
wbwwurwgrbbwgubwuubwugugbbgbugwugbgbwgwugrwbgr
rurwrwrwbwgrwbrugwbbugbbwwguwgbwwrwggbwubbbwggwwgbuurr
ubwgwgwwuuwggrbbrgbrbrgwwbgbrwugguwuwgwrrguuu
wwwrrwburwggbwbgrruurbrugburwbrbgurrwgggbw
bwubgwgubbwwgwwuwwugrrwgrbubbggwwburbuuurg
rrurrgrbbbbrurbuuwwrbbrwuwwurrgwbgrwwgbuwrwgu
wbbrwgugburrugbgugrrbwugwbbrwgrbguwrgrgwwrggr
rbwubrburgwrrrbgguurbrugbgrbruguwrwubwwbwwgu
wgubbwuburwuwrurgrguurwwrwrwwgrbwrrrrbbugr
wggubbubwbrrrwwuubgwwuwrgrgwuwbwbbwuguguuwurbbwugwb
rubuwuwwugrbwgrggrwguwgbwgbrrwwrgurrwrwguuuwuggwrbgrubgwrg
urwruguwgbrggrbbbwgbrgwggwbguburggrwbgrbwwrbgrbwggguwgu
rwbgbubrugubwwrwwgwwurwwgrurrbwbbwgbbugrurw
ggrwbuuwbrurbgrrwbububruwurrruwwwbugbrgwbrwgbgbbgrgub
gurububrrbwbwruwrbbguwuuurugwwrgrbrgguwbwrbggugwwuwwbwru
rrbrwrubwgwuurugguwubrgrrwbgwugburrurgwrgbuurubug
bbrwgwggrbwgrrwrrwggububwurrrggrubggrgurbbuwrwrgrgubwgrgu
ubwbuwgrwrrguwbwuubbubgwwuuwgguuggbuwrrubrrrw
bubuwugurwrwuurbbgrggrurwwbrugrbwrgrbwgggruwgrbb
gbwgrbrwggbwgwrgrwwubrugubwugwbrrwbrurrbbbwuuubgur
bwrbgbrguuwruuwurwwbuwwggrbbrgrwuwuruwrurgguubwggr
rwbubgubwuguwbbrugggrrubbuwgrguurwgrggbrggrrgrbbbbrwgubbu
rwwwgrgubbrruugwurwbgbubbbgurbwwwrurruwwrrrw
rugrbwbbruuggubwwbwbrurrwgbwwggbbrgubgubrwbrurwbwg
gbuwrbrwbwgggrbbgbwrubwbggrugubbubbgggubrbb
bbwrwbwbrurrrrrgbgwwgwbgrrruwgbbuwuugubgugguurrb
bggbuurbgwwgubrwgbgrgrrwbgbgurwrguwrwwbbwgwr
grbrgwububguggbbuwbgbrubuwwrwwwgwbbrurbrrgwgu
bwgwwgrurbbggrgwuwwwgrbbrgburgrguuuwgwuwwubbggwwuggguurwr
ubwrwgbuuwrwuurgrbbubuggrbrggwwuwurrbwrwwgrbwgu
gguwubrwwwrgrwrguwrrggubwuwwuubwbgguubbbwuurwub
grrugubwgrbbuuubuububgbwugurbwrrubgbubrbrrbbguurgbu
gwgrrrruwgwwwwwrrgurggubwuububggrrgggbugwg
rbrbguwuggwwgrgrgrwuubrbbbbugbwguuggbrggguururuu
wburrubgrrrrrwwrrwgurugwwurwwrrwugggwgrggwwubbuwugu
rrwgwwgubbrgruubuggwwuburuwuwubbbrugbrubggrr
bgbbbrrwrrgbuwbgrwubrgwwwbrugrrrwbwwrbrgrrrurubuggr
brwugwrbubuuuggggrwwwwrbwgwrggbburuwbugrwrgbrubrggrwb
wburuwwbgubggbbbrbugwbubrwwrwwwbbgguwwuurgrwuwwgu
rbubugrgwwugwgrurggwugbrurgwuwwubugwwwgbguuwrwgu
rbgbwgwbuubbbwbbuwuuwuwuwuwrbrgwuuggwrruwrbgbrgguwuurbbwr
buuwurwggwrgwwbwbugubgbrwubrrubbwbuwgrwbwbg
brwuwwrwrgbrgruuuubwwbwbwrgrrwrugwwguwwrwgwwgrgwrrrgrwwgu
rbgrwbggbwwruugrbgbububgbwbbuwbugugwrrwbgwbwguwgrrwr
uuurwurrbuwrububbrgruwrbwbwbwrrbrrbrguggburwgu
wuwwuuburrgrbbbwuguwuwgwbbbgwrurubwrruggwbwbrugbrgubgwgwgu
brbbwuguwgggrgwwrrububrbbwuburbgbbugwwwrbuwbuggu
wugwuwwrruwbgwrbwruwrrgrbggbgrbwgggurggbrubruwugwuubur
rwrwwurbguwbbbwburuuruwrgbbguwbgwgggwggbwwgrbg
uggrggwugubuuwbwwwrbgrbrbrgbbrruggbuuuwrbrwwbuggwgrw
rubwwrbwggbubuwwrrgwurruwggbbubggbgrgwwrwrwgbgggbwbbu
bgwrguwwwrrwrgurrrrbbwggguwugbbbubbguwrrurbbrrrrb
bwuuwgrbbugbwgwuwgbgguwuububgbrbubrwgrubgbbwb
ubbruwrwguurwwwgwgugwgguubrurruwrubbgrurwwugbbrrgbwwrbw
wgwwgrggrbwgruwuwgwgbubwgguwbbggrgbgrgruwgwu
gbrruwwwwbgurubrwgwggbrrwrrwrbbrwwgburwwrrb
ubugrrruwgwuwubgbuurgguwwwgubrbbrruurrrbuwwg
wwwgrbrgbguwbwurwurgrwuurwrrggrggrbwuugrwrurbruuwwwggrrgg
wgrrgggbbuubrubrrgwrurwrwbggbuwwbgbrwbwgbbruwbugur
uwugwrrrwruubruwguuugwugwugurrbgurguuwwgbwrwguwggwbgwrw
grgbugubwwgwggbrwwwwuggwwbrgbrururggurwubr
ggrggburuuuwruburbwrgggurguggbbbbwguubgugwrwggrggwugwbrw
rrrrruwgwwrruwgrgwuuuwuwrwbrbbrgrgggwgrurrruurugbrbgub
uguguggwgwrbbbgrwuuwuubgbgrbwuggbbgguwwrwwguruwbbug
gbwgbwbrugugwgbwrubururrrggrgburugwrwwwguuggug
gbrrrgrgrgwrwbgrwrbrwwwgrgwgwuwbwrbrguruburbwru
bwuuururrwgwbwwwruwrurbbgbwgubwwurubgwbubbbwbb
wgrbgubuwbrwugrggbwbruuubgburbbwbrggwrggbbgbuwwurugbbgwbbu
bgrbgwwbbuwrgrubugbbbrbugrwuwuuugwugbrruwuurbbgg
rgrwgbgrguuwbbuwwbrguwrwgrbwgbrbbbwruwgrguwrgwrwrwuuruur
bbwruuwuuwrrbbubgburbwrwwrgwuugbwgbrurrbuguuugwbrubgbwbb
gburbbrggburubugwgbgbgbrwgbrgbggwbugugbrwbwbwrwwgu
wrwgwrrrrrubgwuburgbubrrgwwuuwrrrwbbwrwbrwwgurbwuurguubb
rubwwbugwrwgrwrggwrbrwbbubwgwbbwgbgbrbrrrbubbgguuwggw
ugugbbbbwbuubgwgwgwurwwrwubgrgbubwrrbugrbgrubwgu
ggwggbuwbgrgbgrrrbrggrwruggbwbwrrwrurwrrgg
rrwurgrbrwrrrburwuwuwrrwggwbrwggrwbgggrurggwwrrbguru
wuwurbrbwwbrggbubwuggbbrrbwwbgwbgbbrrbgurwu
bbgwwbgrgurgwwbgggrrgwgwwrbugwgbrguwgrgrurgwurguwbur
wrbuwgubugwgubrggubrgggbbubbggrwggbrgrurgubburwgu
grgwubgrgwbrgurrbgbbbgbrbubwbuwgwubwwwrgrrgrur
bwruwgubgwwgwuubgugugrugbubrgbrguuggbrbbbwwrggwrbbwwb
bgrgbwrurbwwgbwwwgbgwrrrbgwrgbwburubwugwrwgu
ugbrruwgwggrbrrrbuwgwurguggrgurrugggbwburgwrrgr
bwwrgurrbubbguwgrwgwuwrbbwugrrwgbwubgbgrbguwwrburr
wgrbbbubggbuggrgubrbbwbugwgwrwrgrguuwbbrwbrrgrurubwwuw
rwrwrggugruwgrruwbrruurbgguruwwubguggbbbwugrgugbwrgbuuwr
rurwruuwubwrrrbuurwgrubbuwgugwwwggwwrbbbrwruruggbwgrggwugwgu
wuuwrrggrgwwwwgrgrguuuwubgwwwwgrguggwwrubuggggguububuggbu
gbrwgggrburrrbuubbwgwbwuwgrbwbggwgugbwwubuuubrwwgrgbbgwgu
wwgurguggguruuwwuuugurggrugruwrbugwwubwbbrggruwrwuurrgrb
rrbwrwgwbugubgurrrwugbgrbbwgbgrgbbwbgbbruugugubw
rwgugbbbrbrrgubwrrruwubuwbugubwugrubwbgubrbbbwurbrgrw
gbwrggwwbuurgburuurrgwbgwwrubgubwbrwrbguuurbggggrguwbb
rwbbwgwwubugwwbwrubbbrgbrwrbuugrrrubgruwbugrbrrbgwgu
brgwwwuwbubrrbwgrggbrugubuguubbuubbwugbbbrburugwgrw
ugbwurrrburrbrurgrbbgbrrggbbruwwrgbgugburrbuuubbbwgbbrug
ggrguuggurgrrrbrurguggugbwbwwuwrbrbgbwrbubuggg
bggruuruurrwgrwwrrggurguugrwgurgburrwwwrgugg
wwrruugrgbbgrbugbwrbbrrgubwbbuwwgrrgggrggugruubrbubgrg
ggbuuubwgrrubgwgbrrruwrbrrgrwubgguuwbrbggbrwgu
gggbuuruubrrbrwwbwgbuuuwggwrubguggbwbgruuugwgruubbbbgubu
guugwwurwruburruwuugrgruwgwugguggubgurruwggrrgwgbuwrrrub
uurwrbrugrbwugbgggugrrwguuuguguurggrrbbwgu
uburbrgwbgubggwbrrrwruwbgbbuwwrwggbwgrubwrbwgwuuwwgu
ugugrgwgrwrbuwbbbrrwguwwrbwgrwrugugwuubrwwrbwrrugwwbub
rbgbwrwurrruwbuuwrugbbwuwububrrrgbrgwbwrgwruwbuwrg
wbrruwrbubrwbwgrugbrbggwgbwubwruwrgurubrbrrrwrrgbrrr
urwguwrwbrbbggubgbbbuwwrgrbrwgrbguruwgbuwbgbugwbwurrrgw
ggwrgrgbuurbbuwwgurrrwbrwubugwrurbgwuwububburgbggbwwbwwbbwgu
urrbrwuuwurgrgbrgbruwrrwrbbruwwbbbrwwugwbrbwbwbrgugrbbwg
wrwwwuugwbwubugbugrwgubgrrgwguwgwguuuwbbbrrbu
gbgwwgugrwrwrbbbbgbgbburugurggugugwgbrbwwggrrgwrbwbggbgwb
wgburbuuwrbgruugwwwuuuugugbgrgwrbbwuwuggwuuguwubwrwrg
guwubrbbbuwrurbgwrwgwwwbgwwgwwruwuuguuwwbbgrrb
brwrwgurbuubuwrgbrbuurwbwwbgrbruwrrrrgwuwggbrwurr
uwwuuwrbggrrbbrgwrbbbgggbrrwgbgbgwbwbgubwbrggbbrw
rgbguugwwggrbrgwbbburbubwgruubwwubwbrwgrwwuuwrbu
gbgbubgugguruwbwrwbrrbwgbbugggbbwrruwbgurrguwwubwgrbwwggur
rrubgwwrugrrbwbuurbrwgbbrrbbbrrggurwuuuwrggb
rwrwwugbgugrgrggurbugwugbwgburrrgrgbbwwgwuuruggbgwwrbuurbwgu
ubwrgwubuwwbuuggbwwuurubburbuwrrbwgwwgrggrwwrr
uwgbrgburuwrwrgbgrrbuwgrbuguuwggrbbrugwgrbgggbubgbggwbwgu
uuwbbbwuruwguwrbrgguwuurbwgrgrwrbwurwuburgrrgb
wurruuubrruwrubwbwwuurubwwgrwurgwrurwrbwgu
rgggubgrwgwgubwwwuwuwurrruruwbwgrrugguwgwrguwwr
uubbruwuwubbbbrbgwwggwburgwrbubgugwgwbwbbuuuruuwgurb
rwgbuwugrwwurbuggwugruurugbguubwugbrrwbbwgbgrwrbggwbbgrbwgu
wrubwruuwgrguuubgbbugrrbgwrrrruuggwrwgwgrrww
bguggrrwwburrrwuburbgurggbbwgrwwruugbwbuwwrbwrbbbgbwgb
urwgbwbuggrwgbgguwgubugrgwbbwwgwwgrrrguwrr
ugwrrrwrwugrugbbgruwguurbggurbrbubguuubbwrbb
ubuggwrbgwggggrwwwrugbrwbrbrwrwwrrrwgburbbugbggrwwbrgrru
grbgugurggurwuggugwwrwubbgrurburbwbwuugrwrwggrruugbuwwgu
rbgurwrguwbgwbgbbrggwwrugbwugrbwugggbubburbr
ububbuggggwwgwubrwuugbgggrwgbguwurgburgwwrguwuwugb
rgrwrruuruuggrwrrrgbwbububbruuwbrgrgggwugbwrbrur
rbuwuwwrggwgwggbgbuurwbgwgbbrggguwrurwwgggwwwugbuwgwwugug
rbbbuuwbugbruugggwugrgruwbbuwubwgrbgrwrbgwwwbgrwwu
uwwwgruurbbrwbggrrbrugrwrubgwuburuubgurubub
bbgugbgruuubwrbrrwrwgbrwuwurgurggwrbgwwbgruubruw
bgwrgwwrwbubrrbrbrbwbbbruwuwwbuwwuruggbgbrbrgr
ugrgwbbrgubbbbrubrggubbwuugggugbuwbgbwguuugggrwbwbubwbuw
wrgwuwubwrgrgurrwrwbrgurbbbbwbgbubuubuugubgrbbwbggugrrwgu
uwbgrubgbggbbbgbbbwgguubguuwgrrrgwgguuubru
gwwrbuurrugwwbbgugburrggrgbggrwuwgrgububrrbgwbrugwuggrug
bbgrwbrguubggwgbrugwwubwburbbbrgrbgrruguwrrbwurur
bgwbbgrrubwbgwgbwbbggwbwuuwrurwuburgububugbuwrwbgr
uwuuurbrbbrurggwgbguuuwuwbbbuwrrgurwgbwruwwrg
wgwuwugubbubgrgubwubwwurwurbrurrrwwgwrgrbwwruuubbug
gwbwbwrrbrgrgbggubbggugrbrwbwgugbubggrbrgrwwu
wrbuurrbbuwgrbugwrrrgguwwbguwgugrrgbrgwwuuwrbwgu
gwbbwurwrgbbgwwguggbggruwgurubgrrwbrwwwwrwbwbrwuwrwguu
wugburugrgugurrrbgwwwgubuuggubbgbgrgbwbuurbrugggwur
bbwwwrrwrgrwwgggbggrrubbgggrwbrrrwwwwrbrwwr
burbgggrbwrbgbrwrwbubgwrgbwgwwbrrugrruwrgbbggrwgwgu
rbuwgggrggbrwbrgwwwgwgbbggwggbgbubrbugburwwrgwgu
wgbrbbuugguuwuwgurrggwburrugwwugugwuuuurrbwwrrrbrwg
grbbgbbuwwruwggwbbgubwuuuruugbbrwgbrrbbbwrurubwgbrbbburu
wrwgrurgbwwgrwwrwurugbrbgubgwbwbggurrurubuwuguubwbrwgu
brrurgwrbwuwbbrbgggbwwwbuuwbwrrbrwbguwuwbrugbrrggwgu
uwuuubuguuuurbuuugwbwurrrubuwguubrruwrrgugwrrwgwbrgbub
wbrwrrrrbggbubrrgwwrguguurrrwbwuugwwggrwgwgu
rgrrwugwgrgugubugrrrwburgrurguwbgrwrwgbbwwgu
wgrwbrugbwwbbubwwrgwugubrurguuwggruuuggbbgggbur
rwbwbwgruuubwbgwurrruguugurgguuuururgrurgg
gurruuwbbrrgbuuwbwwgwwwggrurwbgbwurrwwuwbugbggbrg
bbugwbuwbuwbbbwgruwggbwgggurrurwuuwgguuwbgwwuguwurruggww
brbwrbwubbuwgbwwgbbruurgwrwguguwubuwrbrguwguuguugbbwwub
rugwwuwrrbwbrrgbuuurwwgwruuubrbwbggwurwuubbuuuuguu
uggwurwbwggwwubgwbwuuburgwrubwgwubwrbwgu
bgggbwrgbugrbrbgbwbwbbwguubruugwbguuurgubuwbbrubwwgu
ubgggrbbubruurbgurubrgguwruurrbrgbgbbwgwgbgrwgg
rrbbgwbgggugbuwgrwwwuugruwwgwbwrwgwuwbbgbubuwbbwrb
uwggbrrgbrwbbbwrrwbwuugrwrrwrbwgurggburbwgu
grbbwrggrbgbwwrwugbwubuguugbburuguubbwuuwbg
brugwrwuggrwwrwbwrbbbguurbgwgwguwwuurrrugwrrgurwgu
wurruruwrguuwbrubbuwgwurrggrrwwuuuggwgbrrugb
rrbrrurguggrwruggrrwrwwrbbwwwuruwbbgrrgwgu
guugggbgwwurwwgbwbgwgggbbgrguggbuuwbwurrgbwbgrgrr
gurwurbgrrgrbguwwbwbuwrwrrurbwgrubbwgwwwrurwbrrwrr
rgrwrrggbrgwwwuwbrruggubgrguuuuburgugrurbrgwrbwurbuguww
buugwrwrrgurgrgbuwurwwbwrwrbrrrgwwbruwgwwgggrgrrruurubuu
ubgbugrrwwbwrurwuwugwurggbwbbuwrggbwrbrurrbb
bwwbwuwbwubgrbwuwwuwrbbwbubwruwwbrgwuuurwugbruwuggwbub
rggwbgubwbbbbgwggrgbubbbbuwugrrgugwbrbruwrg
gruuwbgrgrgwggwwwwrrwgbwgrguwrwbguurwrbggwgu
rwgrbbwgbrgbggrubrbwwugrwgubbubbuugbbrggrrwrurburbr
buggugrwwrwwuuubguruwbubgrwwwbubggrgbbrrwr
burbrrubrrwgurguwuuurbrrrrugrbrbwbbwbguuwubgwru
ggrubgugurgbbbbubgbugruurbwrgbuubrbbrbwguugw
rggwubbbbubuubbgwwwgbrwrwggbgrgwgwgrbrwbwuuwguwwwbbrrrb
rrwgguwbgwrrubuuggbbrwgwurubwubbuugwbgbrbbbwwurruwgbubu
rrruuuuwuubwubwrrbguwgbrurwrrgwrrbuwuwwbgwgwwgu
gwuuwrubururgbgruuwwuwwwrrrwubuwwbbuwrgugwrbwgwgrugrrrrw
rwugbruurbbwgugbrugwwbgbwuwwwwbrrbuggurwbwbggbwubbur
bgggbrurwrwgurruugrgwgwbwbwuwurwuwgrbuwbubwbbb
gbbbuwrbgguwgrguububrbubggrbgbrwggbrwuwwgugugbb
wbwbguggrrbuwrggrwugguubbbbrbbuguuurugbrrrrbbggrruubgbbrub
bggbwwurguwguwbugbururrrrbwrwburrrgubbbwwrrubuwg
rbuwrbwubrrrrwwggrwwuburuwrrbgruwwguwuuguggrwrbgbrbugu
bwbrurgugwbbgrwrrbwrrrrwbgwbbwguwubwuuwgrrwr
rwuwwbgwbwbrugbbbgrwbrgbuwuguwrwgrgbgbggbggbbwwuwbbgwgwb
bwrurugbrurgbrguuwbubururuugwuuuwugbuuwwruwuurg
wgbrurwguwrbrbbrbrbggbgbbuugwgwuwrrrrgubggrruuwuuugu
gggrgbuwbbgbgurgwbbgrbwwwwgbbwrgugbgwwwgu
bbrwguwuubguuubbwwrugubgwrbuggbrgbguwurruruwruwwgr
bbbgrgbbrbbgruwrwwrgugwuwbubrwuwrwurbwbgubwwur
gwrbruwbbbbrbwgubwwrruuuuggwrrgwrbbrgrubgubwrbggbwgu
wwwwubgrrurggwrgwbrburwbwwbbbgrggbubugbrbbbuwbwbwbbg
wbwburwgwrwwwgrbuuwwrgrbuuuwwurwruuwgrwbwgruwrruruw
gwuurrwgrbwwggwgggwwuwgurbuuugrwgrbrwbwurwwbgwgu
ubgwuggwgrbwubgurggrbrbrgrwurrrbrgwrggbrwugurbwubrwg
brbgbbgguwgbugbrrurwwugurwbbwgwgbwruwwgwbuwwgugbwgrubrbg
uwgbwgwugbrwuwggbrrgbgbrbuwrwubwwgubruwgrbggbuwwbgggrwg
uuwuuuubwburgbwrwguwbwgubuwuwbbugruwwugwuubgrruw
bbgwgwbbruwbbwggrgwgwbwubrugrrgwbgbwguwwbrgrbgwgu
wbbgbwggwruguggbbgruuuuwuuugbbbwwgbwrruwrbgrgwggubrwgb
urrrrwbgwubgubbgrrbwbruugwurwuwggwgrrgrrgugrwwrggbuurgb
rwbrggrwgwugrrbuwbrrurggwgrugbrurrrbbwwrgbwgubrguwrugrgu
urwrrgrwugbuggbbgwurgbugbrbbgrguuuuugurrurbwrbrwgwgu
wbwrwbwubgbbgugrrrwurgurgrwrugugbbrbrbbgrwgwbwgrrbbgbguwg
rwbwuwrurubwuwwgbrwugrwuurrbbubgguwbugwwrgbrrb
buguurggwrubgwurrbugugbgubrgwbggwgrgwggruur
grruwwububurbgubgbwubrurbrbwguuwgwgbbrbguwbguu
brgugwwbgggrgwbbbrugugbwgrbwuwgbbruuwuwgburwbugwwguruuuwb
wburbwruwugurugrrwrbwgbwubuubrbuugrgrwgrgrbrbru
rgwrgwwuruwguuwwbgubwgwruwuguuwggubgubugurru
rwwrrbwbuuuuwwgwuwguubbruuurgbrbrrgwgwgrbuwrwwbbubr
wrbrrgrbgggubrwwrrubbrrrbbrburrgugbwurwrggwruwubgwwug
ubwwwrgwbwburwrgwrwwwuuugwruwwwbwwuwguwgrrurb
gugwbwrgugggbbuggwwwrwuurrrwruwwrbguwrrwwrrrgwgurgrubbuw
wwubrurwwrrbwwwurrwruwwwrwruwwgrrrrbuwbgbggwwgu
gurgguwggggruububggwwbgbgggwuggbrwbrbuguwgwgwurrrugw
rrrguruggugwbwrwgbuwwurugggwwgwguwggwrrwrgwwuguwbrgrwb
bbwguuuwbugbrbuwrururbuwubwguuwrrgwwruugguurbbrwgruuurbwwb
bwrrwbubguuwggwwuwrbwwgugrbrgwrbguuuggubwurggwwg
rggwrbgrrugrbgrbwbrwbruuwbrugurrwgbwggbwrr
bwggbrbrguwurwwubuwuggwgwggrbbggurgbbbwwbgwbwg
bruwbwwuuuwrwrbwwgwbubwwuwbgrgggugbbrbbruubwrggguubwwr
uwugbgrbgburbrrubburbggbbguruwururbbwguwbururbugbugwg
gbwurgrrbbrrrbggwgbugwurbgbrbbwugrwbgwgbubw
uggrrbubburgbrgbuwbrrgggrwgbuubgrwgrruuwrggrrwuuggww
uwgrgurgbggwguuurbuuuruuwruwrwrubburruguubbwwrwgubugbbwg
brwgurruuguguururruurwuwwbgrbbbwgrgrrggbwbbwgu
wguugwurrgugwbwgwbrbgrbgubwwgbwurburubuugbubbbgwrbwgu
wbggrwwbbbgubrguwugrgbwbgbrwugwgbgwwwwuuwugubrbrbrrb
rgrgrwwbgggggwurugrrubbbuggrgbgrwrwwruguwrrurgggwguwwgbwbw
gurgrurrwbuwuwbrugrbwwrgbwruurugwwuwrrgggurrrwwwuurwggug
rwgwubgbuwuggurrugburwwgbwbwrbugrbrgbrgwgrugbgrubwgrwruu
uubrwbwgwrrruurrwbuwbgrbuuuwwwubrrubrwgwbwbuwbgbrwrw
rgwggggwruubuwurgrgbwgubgbwururugwwbbgbruuuwgrruubwbwbwr
gguwbggwwbwwubrgurwgbwurggwwuuubgrbbbrwrbububrbwrbr
ugubguwwuwuburbruubrbubwbrburrrubrruwbuwbuuwwbwggwubwwwrwgu
ururbbgbwwbwrrggwwubrurbwugwbgugugbuwbrgbbggrb
rbburwgrbugwuuuwrrggbbbwbrgbwwuuruurrwgbrrwrbburwbugwgu
gwggbrwbwguwbwuuuuguwgbbggbubwbbgugbwuuuwwbwrbwbbuwu
brbrwguwbruwwgugrurwugrgrgrwrguwugbrggurgrbuw
wubbbwuwuwwubbwbuwbubugggrbbgguubuuwwbrruuuwbbwgu
gwuubwuurwruugbgruuwrgwgrubgbbbrwuuwruubgbbbgrww
wwbbwbrbguburwbuugwrwbubgwgrgugrwwbbgwuwuwrgrgwwbggbubuwbu
gubbwugwgubbrbruuwubggwrbrbwwrguuwwuwugwuuurrrwgrbwrubwgr
grubbburrwbgrbgbwgbruwuwwwgwrwbwrggubbbwgrggbwuurwwgu
bugbbgurwwbggwrwgbgbbgbugrgugrbwrgurggwbrwgu
gwruugubwugrrbbwbgwburgwrbbgrugrrbbuugrwwugbuburggbwgu
wugruwwbwbburrgrgwubruwwbbwrgubwbrrwwgbrbgbbwgurwbrg
guggrbwubbbbgguuwrrrurbbuwgrwugbubbrurguugrgbubwubwgub
guwurburbgwbbrgbwuggwbuguwwggrururwbwwguwgwwwwwwrggbr
bwbbwrrwwubwgbwgggwbgrrbwgbbggrrwwubbwwggugurw
rgrbwuwruwbuwrwgrwwrbbgugbbbburbuguuuuuuwugurwbwgu
uwbguwrbuggwwuwbwgbgugugurguwbrubgrwrrrggguwbgwrgggbgugb
brwuuggbgrwuggbwbggggguwwguuwwrbgrrwggbrwbgbburguuwrrwuru
bgrrgugrbrrwgggbbugurugubgbrrbggbrurubuwguggbwwwwggrrugugu
wgwbubruuwrrrruwwggugburbrurbbggubbwuurrwwuwwruub
gggbwrwrbbubuuwgurwgwbrugwbwbwuurrbwuggbbgbbbbb
bubugbwubugugwrwbuurwwwgwwuugbuuwurrgbugggwuwrrubwgu
uuubbwwwbuggbbwbbugubuggggruwrguuwrruwrwwggb
rubbrwwrbwbgggburrwwwrggggubrwbgguwwrrwrrrwruwggb
gwrrwuwugrubgwwugrwrrbwwubwwrruuwwgbguugwrwugwgbb
ggwbguburwgbbgggbgbgrbwwbgubgugubbrguuwggrwwrgrurubgwrwug
bgwbruwuwuwurwbuwrrwbrwwwwbguggbugbbrbbuurugrugr
wgbbrrgrruugggggggbrrrwwgwuurrbwwbuuugrwrggbrwugrgbwb
rrrugbrurrwbwrgggwuugrurruwgrrwgugrbbbbrwgg
gbbrbwbwrgggubbrrurbggbwrrbuwwuugbwrwurwbwrbgggwruruguurrr
urgwurbrbubbgubrrrbwgbgbgggrruuggwguggubbwrrw
rwgrubbwbugrwrugugurrrbrubwbwurggrubggugrbu
brggrgrbguuguuwwwrrrrwwgwbwbwrguugubbuuuwurrr
buwbwuuuguurrwbbgbgbrgbwwburwrggwurwwrbbgbbbburgbuwgrgbwb
rrwuuruwubrgbubwrggwrguuuugwwrburgwwwwguuuurbggwbgwbw
grgggbruuwgurrbubbwrbwbugbwrbgbbguubuubwggrrgwu
rrwuurgubbuwwuuuwwrwgwrrwbbggugbugrruwbuurbbrrbgrgbruuruwb
brgwgwwguggwwuruguurwgbgrwbbuwbbwbrwwuwbgubruubbugwrguur
bbugugwrrgruwguwwrgwrggwwwwubrguwwbwguuwrbwrb
wguwwwbgbggrrrrwwwgugugrrbwrwurruurbgwuwuwwrgur
ubrbruwwbrwrurbrwgubrbggwuurbgburgurguwwrwbrgugrr
brwrwwgggubbbbbrwwwwbruwrrbbgwwbugbgwrbrgruugrrguruuuu
bbbguwbubrbwwwwgwbgguurgwrggrgguubgwwwrwrbwggrw
rgguwuruwgrwuubwwggrrrgbgbuwwrbrwwrrrgrgwwuwurbburbw
wuwuwwgwwgbwrwbbruwuurrugubwrgrwgbuwrwgubgbwbbw
bwrugbwwwggrwrguwurrbwgugrgurbrubuggugburg
rrruugguwuuwruuuwwubrrgwbgwububbbbbbwuruggu
ugrbwugbwwuuuwwurubwrbrrrgburuwrurrwuburwr
wrgrwwbwbwuwbggrwwbwwbrrrbrbbrwwgwwgggbrguubguwwrg
bbrubbrwruwggurrubrbwbgbwbbbrruwurgggrwwbbwuwuguburbb
ubgbuuburwrwbgbwrwwguwbbrrbwrrwrbggrruubrgbugrwrrrrrb
brbwuwuugwubrwrguwgbuuwgguwgubrbuwgugbbbgbwrrbgr
uuwwrruwwbuwbbbuwwwburugububwgggbbggbrgwuuugrrgr
bbbugguwubbbuubwwwurbwwgrggwgggwruuurbrrwubbbuwwu
gubrubwwwbwwggggurbbguuwwwgwbwwrgrwuwubwrwbggwbbug
rwgubgbrgbbrbruggrwrguwuuruwwgugwuuuugubrruwgbgwgu
bruuubrgugbgrgbwwuuwwgwrgrgguuwgbbgwwbuubgbbrbguburgrww
rwgbgubbwwrwwbbrwuguugrrgbburrubrbgwrrgrrrwgwubbwg
wrggbwrbuurwrgwwgrwugwubbbubrrgrgbubrbruwwgrrrrbbrgwrbw
bbuuurugrwrgubrrubgwgwuwbgwuguwgwwrgwbwgbwbg
wbbugbuuububuuruubwuwbuubugrgwugruurgubgruwburug
bgurrbgrruwwubrgrrwbwuubbwwgburwwrguwgrrwgu
gbbugwrbuwrbrrwurbwwgubuwwwbrwbbrwguggrbwrwgbgwwgbugr
bgrgwbwuugrrwgbbgruwgbwurgbbgurrgwbgurgbbruwuubrrrbuurbbrwgu
rwbwbrrbbruwbwwgggrrruubgwwrrbwrugwwrurrgurggwuwbuugbugb
uuuwuwbrgbuwgggwrrbwbubgbrruurbwububbubbrwrwuu
wbburrrwurbuubbwbugrguwurrgurrbuuwrwuwuburwugugwwwgg
rwgbrbwurggwrgguwrbrrrgbrgbgwuwrwbwwwbbwgbrbguwrurubwgwwur
wggggrbwrrwwwurbbuggubrrgbuburrbwwgrbrgwbgbw
ggbgbbrbbgubuwuuwggrubrwbwrbrwbbrggwgwbbwgbrrbwgu
ubrwgurrruuugubrbwbrwuruwruwrrwrgrwgwguuwub
uwrwwbububwgggwgwgggrrguwwuuurbgbwwbbwgbrwwuw
bwrwwurrwbbrbwbubuuwwgrbbwuurbuwrwgbuurwbgwwuruubwgb
wwggguwbbwwbbwrggwbgwbwugguugggbburbbrruwbbuguggw
bgwwuuguururgurgbrggruwrwbbbrrububuwurrwuwbgwgugubub
uwrbgbwubwrbrgwwuwgrbubggwrbbgwruwgrwubgrbbwgw
wgubruwgbgwrguggwwuwguwrurwgubbugwurwbbgrrwgrwbwgu
rbwuurguwwrgurrwwbrbwururwuwrwwggugwbwrrbggrwuurrgbrubbwgu
  </pre>
</details>
