---
title: Crossed Wires
description: Advent of Code 2024 [Day 24]
layout: default
lang: en
tag: aoc24
prefetch:
  - adventofcode.com
  - deno.com
---

You and The Historians arrive at the edge of a [large grove](https://adventofcode.com/2022/day/23) somewhere in the jungle. After the last incident, the Elves installed a small device that monitors the fruit. While The Historians search the grove, one of them asks if you can take a look at the monitoring device; apparently, it's been malfunctioning recently.

The device seems to be trying to produce a number through some boolean logic gates. Each gate has two inputs and one output. The gates all operate on values that are either **true** (`1`) or **false** (`0`).

- `AND` gates output `1` if **both** inputs are `1`; if either input is `0`, these gates output `0`.
- `OR` gates output `1` if **one or both** inputs is `1`; if both inputs are `0`, these gates output `0`.
- `XOR` gates output `1` if the inputs are **different**; if the inputs are the same, these gates output `0`.

Gates wait until both inputs are received before producing output; wires can carry `0`, `1` or no value at all. There are no loops; once a gate has determined its output, the output will not change until the whole system is reset. Each wire is connected to at most one gate output, but can be connected to many gate inputs.

Rather than risk getting shocked while tinkering with the live system, you write down all of the gate connections and initial wire values (your puzzle input) so you can consider them in relative safety. For example:

```
x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -> z00
x01 XOR y01 -> z01
x02 OR y02 -> z02
```

Because gates wait for input, some wires need to start with a value (as inputs to the entire system). The first section specifies these values. For example, `x00: 1` means that the wire named x00 starts with the value `1` (as if a gate is already outputting that value onto that wire).

The second section lists all of the gates and the wires connected to them. For example, `x00 AND y00 -> z00` describes an instance of an `AND` gate which has wires `x00` and `y00` connected to its inputs and which will write its output to wire `z00`.

In this example, simulating these gates eventually causes `0` to appear on wire `z00`, `0` to appear on wire `z01`, and `1` to appear on wire `z02`.

Ultimately, the system is trying to produce a **number** by combining the bits on all wires starting with `z`. `z00` is the least significant bit, then `z01`, then `z02`, and so on.

In this example, the three output bits form the binary number `100` which is equal to the decimal number `4`.

Here's a larger example:

```
x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj
```

After waiting for values on all wires starting with `z`, the wires in this system have the following values:

```
bfw: 1
bqk: 1
djm: 1
ffh: 0
fgs: 1
frj: 1
fst: 1
gnj: 1
hwm: 1
kjc: 0
kpj: 1
kwq: 0
mjb: 1
nrd: 1
ntg: 0
pbm: 1
psh: 1
qhw: 1
rvg: 0
tgd: 0
tnw: 1
vdt: 1
wpb: 0
z00: 0
z01: 0
z02: 0
z03: 1
z04: 0
z05: 1
z06: 1
z07: 1
z08: 1
z09: 1
z10: 1
z11: 0
z12: 0
```

Combining the bits from all wires starting with `z` produces the binary number `0011111101000`. Converting this number to decimal produces **2024**.

Simulate the system of gates and wires. **What decimal number does it output on the wires starting with z?**

```ts
type Operation = "AND" | "OR" | "XOR";

interface Gate {
  input1: string;
  input2: string;
  operation: Operation;
  output: string;
}

function solve(input: string): bigint {
  const lines = input.trim().split("\n");
  const wires = new Map<string, number>();
  const gates: Gate[] = [];
  let blankLineFound = false;

  for (const line of lines) {
    if (line.trim() === "") {
      blankLineFound = true;
      continue;
    }

    if (!blankLineFound) {
      const [wire, value] = line.split(": ");
      wires.set(wire, parseInt(value));
    } else {
      const [def, output] = line.split(" -> ");
      const parts = def.split(" ");

      if (parts.length === 3) {
        gates.push({
          input1: parts[0],
          operation: parts[1] as Operation,
          input2: parts[2],
          output
        });
      }
    }
  }

  let changed = true;
  while (changed) {
    changed = false;
    for (const gate of gates) {
      const input1 = wires.get(gate.input1);
      const input2 = wires.get(gate.input2);

      if (input1 === undefined || input2 === undefined) {
        continue;
      }

      let output: number;
      switch (gate.operation) {
        case "AND":
          output = input1 & input2;
          break;
        case "OR":
          output = input1 | input2;
          break;
        case "XOR":
          output = input1 ^ input2;
          break;
      }

      if (wires.get(gate.output) !== output) {
        wires.set(gate.output, output);
        changed = true;
      }
    }
  }

  const zWires = Array.from(wires.entries())
    .filter(([wire]) => wire.startsWith("z"))
    .sort(([a], [b]) => a.localeCompare(b));

  let binaryString = "";
  for (const [_, value] of zWires.reverse()) {
    binaryString += value;
  }

  return BigInt(`0b${binaryString}`);
}

async function main() {
  const input = Deno.readTextFileSync("input.txt");
  console.log(solve(input));
}

main().catch(console.error);
```

After inspecting the monitoring device more closely, you determine that the system you're simulating is trying to add **two binary numbers**.

Specifically, it is treating the bits on wires starting with `x` as one binary number, treating the bits on wires starting with `y` as a second binary number, and then attempting to add those two numbers together. The output of this operation is produced as a binary number on the wires starting with `z`. (In all three cases, wire `00` is the least significant bit, then `01`, then `02`, and so on.)

The initial values for the wires in your puzzle input represent **just one instance** of a pair of numbers that sum to the wrong value. Ultimately, **any** two binary numbers provided as input should be handled correctly. That is, for any combination of bits on wires starting with `x` and wires starting with `y`, the sum of the two numbers those bits represent should be produced as a binary number on the wires starting with `z`.

For example, if you have an addition system with four `x` wires, four `y` wires, and five `z` wires, you should be able to supply any four-bit number on the `x` wires, any four-bit number on the `y` numbers, and eventually find the sum of those two numbers as a five-bit number on the `z` wires. One of the many ways you could provide numbers to such a system would be to pass `11` on the `x` wires (`1011` in binary) and `13` on the `y` wires (`1101` in binary):

```
x00: 1
x01: 1
x02: 0
x03: 1
y00: 1
y01: 0
y02: 1
y03: 1
```

If the system were working correctly, then after all gates are finished processing, you should find `24` (`11+13`) on the `z` wires as the five-bit binary number `11000`:

```
z00: 0
z01: 0
z02: 0
z03: 1
z04: 1
```

Unfortunately, your actual system needs to add numbers with many more bits and therefore has many more wires.

Based on forensic analysis of scuff marks and scratches on the device, you can tell that there are exactly **four** pairs of gates whose output wires have been **swapped**. (A gate can only be in at most one such pair; no gate's output was swapped multiple times.)

For example, the system below is supposed to find the bitwise `AND` of the six-bit number on `x00` through `x05` and the six-bit number on `y00` through `y05` and then write the result as a six-bit number on `z00` through `z05`:

```
x00: 0
x01: 1
x02: 0
x03: 1
x04: 0
x05: 1
y00: 0
y01: 0
y02: 1
y03: 1
y04: 0
y05: 1

x00 AND y00 -> z05
x01 AND y01 -> z02
x02 AND y02 -> z01
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z00
```

However, in this example, two pairs of gates have had their output wires swapped, causing the system to produce wrong answers. The first pair of gates with swapped outputs is `x00 AND y00 -> z05` and `x05 AND y05 -> z00`; the second pair of gates is `x01 AND y01 -> z02` and `x02 AND y02 -> z01`. Correcting these two swaps results in this system that works as intended for any set of initial values on wires that start with `x` or `y`:

```
x00 AND y00 -> z00
x01 AND y01 -> z01
x02 AND y02 -> z02
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z05
```

In this example, two pairs of gates have outputs that are involved in a swap. By sorting their output wires' names and joining them with commas, the list of wires involved in swaps is `z00,z01,z02,z05`.

Of course, your actual system is much more complex than this, and the gates that need their outputs swapped could be **anywhere**, not just attached to a wire starting with `z`. If you were to determine that you need to swap output wires `aaa` with `eee`, `ooo` with `z99`, `bbb` with `ccc`, and `aoc` with `z24`, your answer would be `aaa,aoc,bbb,ccc,eee,ooo,z24,z99`.

Your system of gates and wires has **four** pairs of gates which need their output wires swapped - **eight** wires in total. Determine which four pairs of gates need their outputs swapped so that your system correctly performs addition; **what do you get if you sort the names of the eight wires involved in a swap and then join those names with commas?**

```ts
type Operation = "AND" | "OR" | "XOR";

interface Gate {
  input1: string;
  input2: string;
  operation: Operation;
  output: string;
}

function parseInput(input: string): { wires: Map<string, number>; gates: Gate[] } {
  const lines = input.trim().split("\n");
  const wires = new Map<string, number>();
  const gates: Gate[] = [];
  let blankLineFound = false;
  for (const line of lines) {
    if (line.trim() === "") {
      blankLineFound = true;
      continue;
    }
    if (!blankLineFound) {
      const [wire, value] = line.split(": ");
      wires.set(wire, parseInt(value));
    } else {
      const [def, output] = line.split(" -> ");
      const parts = def.split(" ");
      if (parts.length === 3) {
        gates.push({
          input1: parts[0],
          operation: parts[1] as Operation,
          input2: parts[2],
          output,
        });
      }
    }
  }
  return {wires, gates};
}

function searchSwap(gates: Gate[]): string {
  const list: string[] = [];
  const highestZ = Array.from(gates)
    .filter(g => g.output.startsWith('z'))
    .sort((a, b) => b.output.localeCompare(a.output))[0].output;

  for (const gate of gates) {

    // Check for output gates with wrong operation
    if (
      gate.output.startsWith('z') &&
      gate.operation !== 'XOR' &&
      gate.output !== highestZ &&
      !list.includes(gate.output)
    ) {
      list.push(gate.output);
      console.log("Output gate with wrong operation:", gate);
    }

    // Check for XOR gates between non-input wires
    if (
      gate.operation === 'XOR' &&
      !gate.output.startsWith('x') && !gate.output.startsWith('y') && !gate.output.startsWith('z') &&
      !gate.input1.startsWith('x') && !gate.input1.startsWith('y') && !gate.input1.startsWith('z') &&
      !gate.input2.startsWith('x') && !gate.input2.startsWith('y') && !gate.input2.startsWith('z') &&
      !list.includes(gate.output)
    ) {
      list.push(gate.output);
      console.log("XOR gate between non-input wires:", gate);
    }

    // Check for AND gates not connected to x00
    if (
      gate.operation === 'AND' &&
      gate.input1 !== 'x00' &&
      gate.input2 !== 'x00' &&
      !list.includes(gate.output)
    ) {
      for (const subGate of gates) {
        if (
          (subGate.input1 === gate.output || subGate.input2 === gate.output) && 
          subGate.operation !== 'OR' &&
          !list.includes(gate.output)
        ) {
          list.push(gate.output);
          console.log("AND gate not connected to x00:", gate);
          continue;
        }
      }
    }

    // Check XOR gates connected to OR gates
    if (gate.operation === 'XOR' && !list.includes(gate.output)) {
      for (const subGate of gates) {
        if (
          (subGate.input1 === gate.output || subGate.input2 === gate.output) && 
          subGate.operation === 'OR' && !list.includes(gate.output)
        ) {
          console.log("XOR gate connected to OR gate:", gate);
          list.push(gate.output);
        }
      }
    }
  }

  return list.sort().join(",");
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const {wires, gates} = parseInput(input);
  console.log(searchSwap(gates));
}
  
main().catch(console.error);
```

## Links

- [If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2024/day/24)
- [Solve Advent of Code 2024 with Deno and Win Prizes!](https://deno.com/blog/advent-of-code-2024)

<details>
	<summary>Click to show the input</summary>
	<pre>
x00: 1
x01: 1
x02: 1
x03: 1
x04: 0
x05: 1
x06: 0
x07: 1
x08: 0
x09: 1
x10: 1
x11: 1
x12: 1
x13: 0
x14: 0
x15: 0
x16: 0
x17: 1
x18: 0
x19: 0
x20: 0
x21: 1
x22: 1
x23: 0
x24: 0
x25: 0
x26: 0
x27: 1
x28: 0
x29: 0
x30: 1
x31: 0
x32: 1
x33: 1
x34: 1
x35: 0
x36: 1
x37: 0
x38: 1
x39: 1
x40: 0
x41: 1
x42: 1
x43: 0
x44: 1
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1
y05: 0
y06: 0
y07: 0
y08: 0
y09: 0
y10: 1
y11: 0
y12: 0
y13: 0
y14: 0
y15: 1
y16: 1
y17: 1
y18: 1
y19: 0
y20: 0
y21: 1
y22: 1
y23: 0
y24: 0
y25: 0
y26: 1
y27: 1
y28: 1
y29: 1
y30: 1
y31: 0
y32: 0
y33: 1
y34: 1
y35: 1
y36: 0
y37: 1
y38: 0
y39: 0
y40: 1
y41: 1
y42: 1
y43: 1
y44: 1

ksf AND bqf -> jkw
x42 AND y42 -> vgs
wmv AND whq -> kjf
mrj AND dnc -> gvs
tjc OR hnb -> msv
kpq AND rvw -> kck
wvj OR hvw -> hqd
fqp XOR qwj -> z20
x28 XOR y28 -> qvr
x20 AND y20 -> gqm
cpd AND gcn -> sbs
y07 AND x07 -> z07
hnn OR jkw -> nff
x39 XOR y39 -> rfr
y15 XOR x15 -> mjw
wmv XOR whq -> z05
y22 XOR x22 -> kpq
ncc AND cqs -> snr
x24 XOR y24 -> vmk
csm AND cbj -> dnt
y02 XOR x02 -> gsh
x00 AND y00 -> gmk
vgr XOR gmk -> z01
dmn XOR dsj -> z19
vtd AND rfr -> sgk
hch XOR nff -> dmn
ghp OR qfs -> spk
mhf AND tgs -> fdq
x13 XOR y13 -> rhg
y24 AND x24 -> hqt
btp AND hsm -> jtg
qqj XOR tqh -> z42
x33 AND y33 -> hbn
y12 AND x12 -> jns
y01 XOR x01 -> vgr
mvb AND mqd -> stp
x10 AND y10 -> tgf
wdt AND hqd -> shh
jjc OR sgk -> rrb
gsh AND nwk -> hhn
x07 XOR y07 -> mvw
y39 AND x39 -> jjc
gcn XOR cpd -> z23
vws OR qmp -> ksf
kmq OR kqs -> rvw
hkg XOR rrb -> z40
qnm AND rfk -> z35
y11 XOR x11 -> qjj
gmf OR kdk -> vnn
sbs OR rww -> gqk
y30 AND x30 -> vqq
x06 XOR y06 -> btp
x01 AND y01 -> jvm
x05 AND y05 -> wtq
y27 AND x27 -> pft
kqq OR mst -> cmt
y17 XOR x17 -> bqf
btg XOR mjw -> z15
x16 XOR y16 -> fnb
mbw XOR kth -> z36
x41 AND y41 -> qqq
hqn AND srg -> hfm
swv OR gqm -> vmm
qdd OR jdb -> z45
mvd OR hqt -> nsp
y10 XOR x10 -> mvb
x20 XOR y20 -> qwj
dqk OR vrv -> cqs
gqk XOR vmk -> z24
y16 AND x16 -> vws
mpf AND jbj -> nnt
fnb AND fpg -> qmp
jtg OR hgj -> pmc
nwk XOR gsh -> z02
y32 AND x32 -> sdm
qtf OR gkk -> njp
qwj AND fqp -> swv
jns OR fdq -> fmh
y42 XOR x42 -> qqj
bqf XOR ksf -> z17
rfg AND nws -> cms
mdq AND cmk -> qdd
x21 XOR y21 -> tqw
y41 XOR x41 -> gnj
y33 XOR x33 -> dnc
sjn XOR btf -> z03
x23 XOR y23 -> cpd
y27 XOR x27 -> mpf
y31 XOR x31 -> kng
vgs OR rmc -> rfg
y22 AND x22 -> bhk
y35 XOR x35 -> rfk
wdt XOR hqd -> z14
qvt OR nsh -> jbj
fmk AND cmt -> ktk
x06 AND y06 -> hgj
qqq OR tmj -> tqh
vmm AND tqw -> kqs
rkh AND nsp -> gkk
wcd AND smf -> gpc
x36 XOR y36 -> kth
x34 XOR y34 -> wcd
rhg XOR fmh -> z13
x00 XOR y00 -> z00
y43 AND x43 -> tsb
fsm OR jvm -> nwk
gpt OR fsw -> vtd
btg AND mjw -> ghs
wmn OR gpc -> qnm
x28 AND y28 -> mst
y05 XOR x05 -> wmv
y18 XOR x18 -> hch
bws OR ghs -> fpg
y44 AND x44 -> jdb
kng AND mfw -> ghp
y14 AND x14 -> wqc
y32 XOR x32 -> bnm
cjr OR nbd -> mqd
y26 AND x26 -> qvt
njp XOR kwd -> z26
rvw XOR kpq -> z22
bfm XOR jnw -> z38
dsj AND dmn -> rsf
rkw XOR whp -> z08
x29 AND y29 -> rhq
y40 XOR x40 -> hkg
gqc OR rsf -> fqp
x37 AND y37 -> ttv
y15 AND x15 -> bws
rhg AND fmh -> wvj
x23 AND y23 -> rww
mvb XOR mqd -> z10
qvr XOR tfc -> z28
x40 AND y40 -> kdk
wnn OR gmt -> rkw
x13 AND y13 -> hvw
fpr XOR cnv -> z04
x18 AND y18 -> khk
rct OR cfk -> mbw
fmk XOR cmt -> z29
fpg XOR fnb -> z16
y26 XOR x26 -> kwd
x31 AND y31 -> qfs
nff AND hch -> stg
jjg OR ggr -> cnv
kth AND mbw -> dqk
btp XOR hsm -> z06
qvr AND tfc -> kqq
y35 AND x35 -> rct
whp AND rkw -> tjc
vcs XOR msv -> z09
cmk XOR mdq -> z44
rrb AND hkg -> gmf
tgs XOR mhf -> z12
y03 AND x03 -> jjg
y12 XOR x12 -> tgs
mvw AND pmc -> wnn
x37 XOR y37 -> ncc
y08 XOR x08 -> whp
y04 XOR x04 -> fpr
y34 AND x34 -> wmn
y44 XOR x44 -> cmk
x03 XOR y03 -> btf
hqn XOR srg -> z30
vnn XOR gnj -> z41
btf AND sjn -> ggr
vgr AND gmk -> fsm
nsp XOR rkh -> z25
y21 AND x21 -> kmq
khk OR stg -> z18
x19 AND y19 -> gqc
y25 AND x25 -> qtf
x14 XOR y14 -> wdt
vmm XOR tqw -> z21
x43 XOR y43 -> nws
gvs OR hbn -> smf
x17 AND y17 -> hnn
nnt OR pft -> tfc
pmc XOR mvw -> gmt
x36 AND y36 -> vrv
x09 AND y09 -> nbd
cbj XOR csm -> z11
kwd AND njp -> nsh
tsb OR cms -> mdq
vmk AND gqk -> mvd
vtd XOR rfr -> z39
dnc XOR mrj -> z33
rfg XOR nws -> z43
y30 XOR x30 -> srg
x04 AND y04 -> mjc
y38 XOR x38 -> bfm
sdm OR tmd -> mrj
ttv OR snr -> jnw
tgf OR stp -> csm
x38 AND y38 -> fsw
hfm OR vqq -> mfw
gnj AND vnn -> tmj
x08 AND y08 -> hnb
mfw XOR kng -> z31
bhk OR kck -> gcn
gfh OR mjc -> whq
shh OR wqc -> btg
mpf XOR jbj -> z27
spk AND bnm -> tmd
x09 XOR y09 -> vcs
ncc XOR cqs -> z37
wcd XOR smf -> z34
spk XOR bnm -> z32
qnm XOR rfk -> cfk
dnt OR qjj -> mhf
msv AND vcs -> cjr
x02 AND y02 -> ksr
x25 XOR y25 -> rkh
rhq OR ktk -> hqn
cnv AND fpr -> gfh
x11 AND y11 -> cbj
kjf OR wtq -> hsm
x29 XOR y29 -> fmk
bfm AND jnw -> gpt
qqj AND tqh -> rmc
hhn OR ksr -> sjn
x19 XOR y19 -> dsj
  </pre>
</details>
