---
title: Pulse Propagation
description: Advent of Code 2023 [Day 20]
layout: default
lang: en
tag: aoc23
prefetch:
  - adventofcode.com
---

With your help, the Elves manage to find the right parts and fix all of the machines. Now, they just need to send the command to boot up the machines and get the sand flowing again.

The machines are far apart and wired together with long **cables**. The cables don't connect to the machines directly, but rather to communication **modules** attached to the machines that perform various initialization tasks and also act as communication relays.

Modules communicate using **pulses**. Each pulse is either a **high pulse** or a **low pulse**. When a module sends a pulse, it sends that type of pulse to each module in its list of **destination modules**.

There are several different types of modules:

**Flip-flop** modules (prefix `%`) are either **on** or **off**; they are initially **off**. If a flip-flop module receives a high pulse, it is ignored and nothing happens. However, if a flip-flop module receives a low pulse, it **flips between on and off**. If it was off, it turns on and sends a high pulse. If it was on, it turns off and sends a low pulse.

**Conjunction** modules (prefix `&`) **remember** the type of the most recent pulse received from **each** of their connected input modules; they initially default to remembering a **low pulse** for each input. When a pulse is received, the conjunction module first updates its memory for that input. Then, if it remembers **high pulses** for all inputs, it sends a **low pulse**; otherwise, it sends a **high pulse**.

There is a single **broadcast module** (named `broadcaster`). When it receives a pulse, it sends the same pulse to all of its destination modules.

Here at Desert Machine Headquarters, there is a module with a single button on it called, aptly, the **button module**. When you push the button, a single **low pulse** is sent directly to the `broadcaster` module.

After pushing the button, you must wait until all pulses have been delivered and fully handled before pushing it again. Never push the button if modules are still processing pulses.

Pulses are always processed **in the order they are sent**. So, if a pulse is sent to modules `a`, `b`, and `c`, and then module `a` processes its pulse and sends more pulses, the pulses sent to modules `b` and `c` would have to be handled first.

The module configuration (your puzzle input) lists each module. The name of the module is preceded by a symbol identifying its type, if any. The name is then followed by an arrow and a list of its destination modules. For example:

```
broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
```

In this module configuration, the broadcaster has three destination modules named `a`, `b`, and `c`. Each of these modules is a flip-flop module (as indicated by the `%` prefix). `a` outputs to `b` which outputs to `c` which outputs to another module named `inv`. `inv` is a conjunction module (as indicated by the `&` prefix) which, because it has only one input, acts like an inverter (it sends the opposite of the pulse type it receives); it outputs to `a`.

By pushing the button once, the following pulses are sent:

```
button -low-> broadcaster
broadcaster -low-> a
broadcaster -low-> b
broadcaster -low-> c
a -high-> b
b -high-> c
c -high-> inv
inv -low-> a
a -low-> b
b -low-> c
c -low-> inv
inv -high-> a
```

After this sequence, the flip-flop modules all end up **off**, so pushing the button again repeats the same sequence.

Here's a more interesting example:

```
broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
```

This module configuration includes the `broadcaster`, two flip-flops (named `a` and `b`), a single-input conjunction module (`inv`), a multi-input conjunction module (`con`), and an untyped module named `output` (for testing purposes). The multi-input conjunction module `con` watches the two flip-flop modules and, if they're both on, sends a **low pulse** to the `output` module.

Here's what happens if you push the button once:

```
button -low-> broadcaster
broadcaster -low-> a
a -high-> inv
a -high-> con
inv -low-> b
con -high-> output
b -high-> con
con -low-> output
```

Both flip-flops turn on and a low pulse is sent to `output`! However, now that both flip-flops are on and `con` remembers a high pulse from each of its two inputs, pushing the button a second time does something different:

```
button -low-> broadcaster
broadcaster -low-> a
a -low-> inv
a -low-> con
inv -high-> b
con -high-> output
```

Flip-flop `a` turns off! Now, `con` remembers a low pulse from module `a`, and so it sends only a high pulse to `output`.

Push the button a third time:

```
button -low-> broadcaster
broadcaster -low-> a
a -high-> inv
a -high-> con
inv -low-> b
con -low-> output
b -low-> con
con -high-> output
```

This time, flip-flop `a` turns on, then flip-flop `b` turns off. However, before `b` can turn off, the pulse sent to `con` is handled first, so **it briefly remembers all high pulses** for its inputs and sends a low pulse to `output`. After that, flip-flop `b` turns off, which causes `con` to update its state and send a high pulse to `output`.

Finally, with `a` on and `b` off, push the button a fourth time:

```
button -low-> broadcaster
broadcaster -low-> a
a -low-> inv
a -low-> con
inv -high-> b
con -high-> output
```

This completes the cycle: `a` turns off, causing `con` to remember only low pulses and restoring all modules to their original states.

To get the cables warmed up, the Elves have pushed the button `1000` times. How many pulses got sent as a result (including the pulses sent by the button itself)?

In the first example, the same thing happens every time the button is pushed: `8` low pulses and `4` high pulses are sent. So, after pushing the button `1000` times, `8000` low pulses and `4000` high pulses are sent. Multiplying these together gives `32000000`.

In the second example, after pushing the button `1000` times, `4250` low pulses and `2750` high pulses are sent. Multiplying these together gives `11687500`.

Consult your module configuration; determine the number of low pulses and high pulses that would be sent after pushing the button `1000` times, waiting for all pulses to be fully handled after each push of the button. **What do you get if you multiply the total number of low pulses sent by the total number of high pulses sent?**

```go
type ModuleType int

const (
	Broadcaster ModuleType = iota
	FlipFlop
	Conjunction
)

type Pulse struct {
	source string
	dest   string
	high   bool
}

type Module struct {
	name         string
	mtype        ModuleType
	destinations []string
	state        bool
	memory       map[string]bool
}

func main() {
	modules := make(map[string]*Module)
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		name := parts[0]
		dests := strings.Split(parts[1], ", ")

		var m Module
		m.destinations = dests

		if name == "broadcaster" {
			m.name = name
			m.mtype = Broadcaster
		} else if name[0] == '%' {
			m.name = name[1:]
			m.mtype = FlipFlop
			m.state = false
		} else if name[0] == '&' {
			m.name = name[1:]
			m.mtype = Conjunction
			m.memory = make(map[string]bool)
		}

		modules[m.name] = &m
	}

	for name, module := range modules {
		for _, dest := range module.destinations {
			if m, exists := modules[dest]; exists && m.mtype == Conjunction {
				m.memory[name] = false
			}
		}
	}

	lowCount, highCount := 0, 0
	for i := 0; i < 1000; i++ {
		queue := []Pulse{ {source: "button", dest: "broadcaster", high: false} }

		for len(queue) > 0 {
			pulse := queue[0]
			queue = queue[1:]

			if pulse.high {
				highCount++
			} else {
				lowCount++
			}

			module, exists := modules[pulse.dest]
			if !exists {
				continue
			}

			switch module.mtype {
			case Broadcaster:
				for _, dest := range module.destinations {
					queue = append(queue, Pulse{source: module.name, dest: dest, high: pulse.high})
				}

			case FlipFlop:
				if !pulse.high {
					module.state = !module.state
					for _, dest := range module.destinations {
						queue = append(queue, Pulse{source: module.name, dest: dest, high: module.state})
					}
				}

			case Conjunction:
				module.memory[pulse.source] = pulse.high
				allHigh := true
				for _, high := range module.memory {
					if !high {
						allHigh = false
						break
					}
				}
				for _, dest := range module.destinations {
					queue = append(queue, Pulse{source: module.name, dest: dest, high: !allHigh})
				}
			}
		}
	}

	fmt.Printf("Result: %d\n", lowCount*highCount)
}
```

The final machine responsible for moving the sand down to Island Island has a module attached named `rx`. The machine turns on when a **single low pulse** is sent to `rx`.

Reset all modules to their default states. Waiting for all pulses to be fully handled after each button press, **what is the fewest number of button presses required to deliver a single low pulse to the module named rx?**

```go
type ModuleType int

const (
	Broadcaster ModuleType = iota
	FlipFlop
	Conjunction
)

type Pulse struct {
	source string
	dest   string
	high   bool
}

type Module struct {
	name         string
	mtype        ModuleType
	destinations []string
	state        bool
	memory       map[string]bool
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	modules := make(map[string]*Module)
	initialStates := make(map[string]*Module)
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		name := parts[0]
		dests := strings.Split(parts[1], ", ")

		var m Module
		m.destinations = dests

		if name == "broadcaster" {
			m.name = name
			m.mtype = Broadcaster
		} else if name[0] == '%' {
			m.name = name[1:]
			m.mtype = FlipFlop
			m.state = false
		} else if name[0] == '&' {
			m.name = name[1:]
			m.mtype = Conjunction
			m.memory = make(map[string]bool)
		}

		modules[m.name] = &m
		initialStates[m.name] = &Module{
			name:         m.name,
			mtype:        m.mtype,
			destinations: append([]string{}, m.destinations...),
			state:        m.state,
			memory:       make(map[string]bool),
		}
		if m.memory != nil {
			for k, v := range m.memory {
				initialStates[m.name].memory[k] = v
			}
		}
	}

	for name, module := range modules {
		for _, dest := range module.destinations {
			if m, exists := modules[dest]; exists && m.mtype == Conjunction {
				m.memory[name] = false
			}
		}
	}

	lk := 0
	zv := 0
	sp := 0
	xt := 0
	buttonPresses := 0
	for {
		buttonPresses++
		queue := []Pulse{ {source: "button", dest: "broadcaster", high: false} }

		for len(queue) > 0 {
			pulse := queue[0]
			queue = queue[1:]

			if pulse.dest == "lk" && !pulse.high && lk == 0 {
				lk = buttonPresses
			}
			if pulse.dest == "zv" && !pulse.high && zv == 0 {
				zv = buttonPresses
			}
			if pulse.dest == "sp" && !pulse.high && sp == 0 {
				sp = buttonPresses
			}
			if pulse.dest == "xt" && !pulse.high && xt == 0 {
				xt = buttonPresses
			}
			if lk != 0 && zv != 0 && sp != 0 && xt != 0 {
				fmt.Println(lk, zv, sp, xt, LCM(lk, zv, sp, xt))
				return
			}

			module, exists := modules[pulse.dest]
			if !exists {
				continue
			}

			switch module.mtype {
			case Broadcaster:
				for _, dest := range module.destinations {
					queue = append(queue, Pulse{source: module.name, dest: dest, high: pulse.high})
				}

			case FlipFlop:
				if !pulse.high {
					module.state = !module.state
					for _, dest := range module.destinations {
						queue = append(queue, Pulse{source: module.name, dest: dest, high: module.state})
					}
				}

			case Conjunction:
				module.memory[pulse.source] = pulse.high
				allHigh := true
				for _, high := range module.memory {
					if !high {
						allHigh = false
						break
					}
				}
				for _, dest := range module.destinations {
					queue = append(queue, Pulse{source: module.name, dest: dest, high: !allHigh})
				}
			}
		}
	}
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2023/day/20)

<details>
	<summary>Click to show the input</summary>
	<pre>
%hs -> sl
&dg -> rx
%vp -> fd, dv
%kz -> jc, mc
%nv -> dv
%hx -> gf
%mm -> vh
%fd -> td
&dv -> hx, bl, rc, fd, xt
%hg -> xq
%td -> dv, hx
%bl -> jt
%br -> jq
%qh -> ln
&xq -> zl, cx, qh, hs, nt, sp
%sg -> vv, tr
%dm -> bl, dv
%gt -> xq, hg
%ln -> mq, xq
%mc -> xv, jc
%tx -> rv, jc
&lk -> dg
%mg -> hl, jc
&vv -> zv, br, kx, mm, tr
%nt -> xq, cx
&zv -> dg
%cd -> jc, ps
%rc -> rm, dv
%nj -> pt, xq
broadcaster -> nt, kx, rc, mg
%gf -> dc, dv
%rm -> dm, dv
%xx -> vv, cz
%jt -> dv, vp
%zl -> nj
&sp -> dg
%xc -> jc, kz
&xt -> dg
%tp -> jc
%lc -> vv, vn
%vh -> xx, vv
%mq -> hs, xq
%cc -> vv
%vn -> vv, cc
%tr -> br
%hl -> qb, jc
%dc -> dv, nv
%jq -> mm, vv
%kx -> vv, sg
%cx -> qh
%sl -> zl, xq
%cz -> lc, vv
%qb -> jc, cd
&jc -> ps, xv, lk, mg
%xv -> tx
%pt -> xq, gt
%rv -> jc, tp
%ps -> xc
	</pre>
</details>
