---
title: Chronospatial Computer
description: Advent of Code 2024 [Day 17]
layout: default
lang: en
tag: aoc24
prefetch:
  - adventofcode.com
  - deno.com
  - en.wikipedia.org
---

The Historians push the button on their strange device, but this time, you all just feel like you're [falling](https://adventofcode.com/2018/day/6).

"Situation critical", the device announces in a familiar voice. "Bootstrapping process failed. Initializing debugger...."

The small handheld device suddenly unfolds into an entire computer! The Historians look around nervously before one of them tosses it to you.

This seems to be a 3-bit computer: its program is a list of 3-bit numbers (0 through 7), like `0,1,2,3`. The computer also has three **registers** named `A`, `B`, and `C`, but these registers aren't limited to 3 bits and can instead hold any integer.

The computer knows **eight instructions**, each identified by a 3-bit number (called the instruction's **opcode**). Each instruction also reads the 3-bit number after it as an input; this is called its **operand**.

A number called the **instruction pointer** identifies the position in the program from which the next opcode will be read; it starts at `0`, pointing at the first 3-bit number in the program. Except for jump instructions, the instruction pointer increases by `2` after each instruction is processed (to move past the instruction's opcode and its operand). If the computer tries to read an opcode past the end of the program, it instead **halts**.

So, the program `0,1,2,3` would run the instruction whose opcode is `0` and pass it the operand `1`, then run the instruction having opcode `2` and pass it the operand `3`, then halt.

There are two types of operands; each instruction specifies the type of its operand. The value of a **literal operand** is the operand itself. For example, the value of the literal operand `7` is the number `7`. The value of a **combo operand** can be found as follows:

- Combo operands `0` through `3` represent literal values `0` through `3`.
- Combo operand `4` represents the value of register `A`.
- Combo operand `5` represents the value of register `B`.
- Combo operand `6` represents the value of register `C`.
- Combo operand `7` is reserved and will not appear in valid programs.

The eight instructions are as follows:

The `adv` instruction (opcode `0`) performs **division**. The numerator is the value in the `A` register. The denominator is found by raising 2 to the power of the instruction's **combo** operand. (So, an operand of `2` would divide `A` by `4` (`2^2`); an operand of `5` would divide `A` by `2^B`.) The result of the division operation is **truncated** to an integer and then written to the A register.

The `bxl` instruction (opcode `1`) calculates the [bitwise XOR](https://en.wikipedia.org/wiki/Bitwise_operation#XOR) of register `B` and the instruction's **literal** operand, then stores the result in register `B`.

The `bst` instruction (opcode `2`) calculates the value of its combo operand [modulo](https://en.wikipedia.org/wiki/Modulo) `8` (thereby keeping only its lowest 3 bits), then writes that value to the B register.

The `jnz` instruction (opcode `3`) does **nothing** if the `A` register is `0`. However, if the `A` register is **not zero**, it **jumps** by setting the instruction pointer to the value of its **literal** operand; if this instruction jumps, the instruction pointer is **not** increased by `2` after this instruction.

The `bxc` instruction (opcode `4`) calculates the **bitwise XOR** of register `B` and register `C`, then stores the result in register `B`. (For legacy reasons, this instruction reads an operand but **ignores** it.)

The `out` instruction (opcode `5`) calculates the value of its **combo** operand modulo 8, then **outputs** that value. (If a program outputs multiple values, they are separated by commas.)

The `bdv` instruction (opcode `6`) works exactly like the `adv` instruction except that the result is stored in the **B register**. (The numerator is still read from the `A` register.)

The `cdv` instruction (opcode `7`) works exactly like the `adv` instruction except that the result is stored in the **C register**. (The numerator is still read from the `A` register.)

Here are some examples of instruction operation:

- If register `C` contains `9`, the program `2,6` would set register `B` to `1`.
- If register `A` contains `10`, the program `5,0,5,1,5,4` would output `0,1,2`.
- If register `A` contains `2024`, the program `0,1,5,4,3,0` would output `4,2,5,6,7,7,7,7,3,1,0` and leave `0` in register `A`.
- If register `B` contains `29`, the program `1,7` would set register `B` to `26`.
- If register `B` contains `2024` and register `C` contains `43690`, the program `4,0` would set register `B` to `44354`.

The Historians' strange device has finished initializing its debugger and is displaying some **information about the program it is trying to run** (your puzzle input). For example:

```
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
```

Your first task is to **determine what the program is trying to output**. To do this, initialize the registers to the given values, then run the given program, collecting any output produced by out instructions. (Always join the values produced by `out` instructions with commas.) After the above program halts, its final output will be `4,6,3,5,6,3,5,2,1,0`.

Using the information provided by the debugger, initialize the registers to the given values, then run the program. Once it halts, **what do you get if you use commas to join the values it output into a single string?**

```ts
class Computer {
  private registerA = 0;
  private registerB = 0;
  private registerC = 0;
  private ip = 0;
  private program: number[] = [];
  private output: number[] = [];

  setRegisters(a: number, b: number, c: number) {
    this.registerA = a;
    this.registerB = b;
    this.registerC = c;
  }

  loadProgram(program: number[]) {
    this.program = program;
    this.ip = 0;
  }

  private getComboValue(operand: number): number {
    switch (operand) {
      case 0: case 1: case 2: case 3:
        return operand;
      case 4:
        return this.registerA;
      case 5:
        return this.registerB;
      case 6:
        return this.registerC;
      default:
        throw new Error(`Invalid combo operand: ${operand}`);
    }
  }

  private adv(operand: number) {
    this.registerA = Math.floor(this.registerA / Math.pow(2, this.getComboValue(operand)));
  }

  private bxl(operand: number) {
    this.registerB ^= operand;
  }

  private bst(operand: number) {
    this.registerB = this.getComboValue(operand) % 8;
  }

  private jnz(operand: number) {
    if (this.registerA !== 0) {
      this.ip = operand;
      return true;
    }
    return false;
  }

  private bxc(_operand: number) {
    this.registerB ^= this.registerC;
  }

  private out(operand: number) {
    this.output.push(this.getComboValue(operand) % 8);
  }

  private bdv(operand: number) {
    this.registerB = Math.floor(this.registerA / Math.pow(2, this.getComboValue(operand)));
  }

  private cdv(operand: number) {
    this.registerC = Math.floor(this.registerA / Math.pow(2, this.getComboValue(operand)));
  }

  run(): string {
    this.output = [];
    
    while (this.ip < this.program.length) {
      const opcode = this.program[this.ip];
      const operand = this.program[this.ip + 1];

      switch (opcode) {
        case 0: this.adv(operand); break;
        case 1: this.bxl(operand); break;
        case 2: this.bst(operand); break;
        case 3: if (this.jnz(operand)) continue; break;
        case 4: this.bxc(operand); break;
        case 5: this.out(operand); break;
        case 6: this.bdv(operand); break;
        case 7: this.cdv(operand); break;
        default: throw new Error(`Invalid opcode: ${opcode}`);
      }

      this.ip += 2;
    }

    return this.output.join(',');
  }
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const lines = input.trim().split('\n');
  const regA = parseInt(lines[0].split(': ')[1]);
  const regB = parseInt(lines[1].split(': ')[1]);
  const regC = parseInt(lines[2].split(': ')[1]);
  const program = lines[4].split(': ')[1].split(',').map(Number);
  const computer = new Computer();
  computer.setRegisters(regA, regB, regC);
  computer.loadProgram(program);
  console.log(computer.run());
}

main().catch(console.error);
```

Digging deeper in the device's manual, you discover the problem: this program is supposed to **output another copy of the program**! Unfortunately, the value in register `A` seems to have been corrupted. You'll need to find a new value to which you can initialize register `A` so that the program's output instructions produce an exact copy of the program itself.

For example:

```
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
```

This program outputs a copy of itself if register `A` is instead initialized to `117440`. (The original initial value of register `A`, `2024`, is ignored.)

**What is the lowest positive initial value for register A that causes the program to output a copy of itself?**

```ts
class Computer {
  private registerA = 0n;
  private registerB = 0n;
  private registerC = 0n;
  private ip = 0;
  private program: number[] = [];
  private output: number[] = [];

  constructor(program: number[] = []) {
    this.program = program;
  }

  private getComboValue(operand: number): bigint {
    switch (operand) {
      case 0: case 1: case 2: case 3:
        return BigInt(operand);
      case 4:
        return this.registerA;
      case 5:
        return this.registerB;
      case 6:
        return this.registerC;
      default:
        throw new Error(`Invalid combo operand: ${operand}`);
    }
  }

  private adv(operand: number) {
    this.registerA = this.registerA / 2n ** this.getComboValue(operand);
  }

  private bxl(operand: number) {
    this.registerB ^= BigInt(operand);
  }

  private bst(operand: number) {
    this.registerB = this.getComboValue(operand) % 8n;
  }

  private jnz(operand: number) {
    if (this.registerA !== 0n) {
      this.ip = operand;
      return true;
    }
    return false;
  }

  private bxc(_operand: number) {
    this.registerB ^= this.registerC;
  }

  private out(operand: number) {
    const value = this.getComboValue(operand) % 8n;
    this.output.push(Number(value));
  }

  private bdv(operand: number) {
    this.registerB = this.registerA / 2n ** this.getComboValue(operand);
  }

  private cdv(operand: number) {
    this.registerC = this.registerA / 2n ** this.getComboValue(operand);
  }

  private run(regA: bigint): number[] {
    this.output = [];
    this.ip = 0;
    this.registerA = regA;
    this.registerB = 0n;
    this.registerC = 0n;
    
    while (this.ip < this.program.length) {
      const opcode = this.program[this.ip];
      const operand = this.program[this.ip + 1];

      switch (opcode) {
        case 0: this.adv(operand); break;
        case 1: this.bxl(operand); break;
        case 2: this.bst(operand); break;
        case 3: if (this.jnz(operand)) continue; break;
        case 4: this.bxc(operand); break;
        case 5: this.out(operand); break;
        case 6: this.bdv(operand); break;
        case 7: this.cdv(operand); break;
        default: throw new Error(`Invalid opcode: ${opcode}`);
      }

      this.ip += 2;
    }

    return this.output;
  }

  reverse(): bigint {
    const reversedProgram = [...this.program].reverse();
    
    const processDigit = (previousValues: bigint[], targetDigit: number): bigint[] => {
      return previousValues.flatMap(prev => {
        const shiftedValue = prev << 3n;
        console.log(`Processing digit ${targetDigit} with previous value ${prev} and shifted value ${shiftedValue}`);

        return Array.from({ length: 8 }, (_, i) => {
          const candidateValue = shiftedValue | BigInt(i);
          const [result] = this.run(candidateValue);
          return result === targetDigit ? candidateValue : null;
        }).filter((val): val is bigint => val !== null);

      });
    };
  
    const result = reversedProgram.reduce<bigint[]>(
      (acc: bigint[], digit: number) => processDigit(acc, digit),
      [0n]
    );
  
    return result.reduce((min, current) => current < min ? current : min, result[0]);
  }
}

async function main() {
  const input = await Deno.readTextFile("input.txt");
  const lines = input.trim().split('\n');
  const program = lines[4].split(': ')[1].split(',').map(Number);
  const computer = new Computer(program);
  console.log(computer.reverse());
}

main().catch(console.error);
```

## Links

- [If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2024/day/17)
- [Solve Advent of Code 2024 with Deno and Win Prizes!](https://deno.com/blog/advent-of-code-2024)

<details>
	<summary>Click to show the input</summary>
	<pre>
Register A: 59590048
Register B: 0
Register C: 0

Program: 2,4,1,5,7,5,0,3,1,6,4,3,5,5,3,0
	</pre>
</details>
