---
title: Haunted Wasteland
description: Advent of Code 2023 [Day 8]
layout: default
lang: en
prefetch:
  - adventofcode.com
---

You're still riding a camel across Desert Island when you spot a sandstorm quickly approaching. When you turn to warn the Elf, she disappears before your eyes! To be fair, she had just finished warning you about ghosts a few minutes ago.

One of the camel's pouches is labeled "maps" - sure enough, it's full of documents (your puzzle input) about how to navigate the desert. At least, you're pretty sure that's what they are; one of the documents contains a list of left/right instructions, and the rest of the documents seem to describe some kind of network of labeled nodes.

It seems like you're meant to use the left/right instructions to navigate the network. Perhaps if you have the camel follow the same instructions, you can escape the haunted wasteland!

After examining the maps for a bit, two nodes stick out: AAA and ZZZ. You feel like AAA is where you are now, and you have to follow the left/right instructions until you reach ZZZ.

This format defines each node of the network individually. For example:

```
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
```

Starting with AAA, you need to look up the next element based on the next left/right instruction in your input. In this example, start with AAA and go right (R) by choosing the right element of AAA, CCC. Then, L means to choose the left element of CCC, ZZZ. By following the left/right instructions, you reach ZZZ in 2 steps.

Of course, you might not find ZZZ right away. If you run out of left/right instructions, repeat the whole sequence of instructions as necessary: RL really means RLRLRLRLRLRLRLRL... and so on. For example, here is a situation that takes 6 steps to reach ZZZ:

```
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
```

Starting at AAA, follow the left/right instructions. How many steps are required to reach ZZZ?

```go
type Maps struct {
	Left         map[string]string
	Right        map[string]string
	Instructions []bool
}

func (m *Maps) ReachZZZ() (steps uint) {
	pointer := "AAA"
	for pointer != "ZZZ" {
		for _, step := range m.Instructions {
			if step {
				pointer = m.Right[pointer]
			} else {
				pointer = m.Left[pointer]
			}
			steps++
			if pointer == "ZZZ" {
				break
			}
		}
	}
	return
}

func (m *Maps) InitInstuctions(line string) {
	m.Instructions = []bool{}
	for _, char := range []rune(line) {
		switch char {
		case 'L':
			m.Instructions = append(m.Instructions, false)
		case 'R':
			m.Instructions = append(m.Instructions, true)
		default:
			log.Fatalln("Parsing error")
		}
	}
}

func (m *Maps) AddMap(line string) {
	splittedLine := strings.Split(line, " = (")
	splittedDestinations := strings.Split(splittedLine[1], ", ")
	m.Left[splittedLine[0]] = splittedDestinations[0]
	m.Right[splittedLine[0]] = string([]rune(splittedDestinations[1])[0:3])
}

func TestMaps_ReachZZZ(t *testing.T) {
	type fields struct {
		Left         map[string]string
		Right        map[string]string
		Instructions []bool
	}
	tests := []struct {
		name      string
		fields    fields
		wantSteps uint
	}{
		{
			name: "Example 1",
			fields: fields{
				Left: map[string]string{
					"AAA": "BBB",
					"BBB": "DDD",
					"CCC": "ZZZ",
					"DDD": "DDD",
					"EEE": "EEE",
					"GGG": "GGG",
					"ZZZ": "ZZZ",
				},
				Right: map[string]string{
					"AAA": "CCC",
					"BBB": "EEE",
					"CCC": "GGG",
					"DDD": "DDD",
					"EEE": "EEE",
					"GGG": "GGG",
					"ZZZ": "ZZZ",
				},
				Instructions: []bool{true, false},
			},
			wantSteps: 2,
		},
		{
			name: "Example 2",
			fields: fields{
				Left: map[string]string{
					"AAA": "BBB",
					"BBB": "AAA",
					"ZZZ": "ZZZ",
				},
				Right: map[string]string{
					"AAA": "BBB",
					"BBB": "ZZZ",
					"ZZZ": "ZZZ",
				},
				Instructions: []bool{false, false, true},
			},
			wantSteps: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &wasteland.Maps{
				Left:         tt.fields.Left,
				Right:        tt.fields.Right,
				Instructions: tt.fields.Instructions,
			}
			if gotSteps := m.ReachZZZ(); gotSteps != tt.wantSteps {
				t.Errorf("Maps.ReachZZZ() = %v, want %v", gotSteps, tt.wantSteps)
			}
		})
	}
}
```

The sandstorm is upon you and you aren't any closer to escaping the wasteland. You had the camel follow the instructions, but you've barely left your starting position. It's going to take significantly more steps to escape!

What if the map isn't for people - what if the map is for ghosts? Are ghosts even bound by the laws of spacetime? Only one way to find out.

After examining the maps a bit longer, your attention is drawn to a curious fact: the number of nodes with names ending in A is equal to the number ending in Z! If you were a ghost, you'd probably just start at every node that ends with A and follow all of the paths at the same time until they all simultaneously end up at nodes that end with Z.

For example:

```
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
```

Here, there are two starting nodes, 11A and 22A (because they both end with A). As you follow each left/right instruction, use that instruction to simultaneously navigate away from both nodes you're currently on. Repeat this process until all of the nodes you're currently on end with Z. (If only some of the nodes you're on end with Z, they act like any other node and you continue as normal.) In this example, you would proceed as follows:

- Step 0: You are at 11A and 22A.
- Step 1: You choose all of the left paths, leading you to 11B and 22B.
- Step 2: You choose all of the right paths, leading you to 11Z and 22C.
- Step 3: You choose all of the left paths, leading you to 11B and 22Z.
- Step 4: You choose all of the right paths, leading you to 11Z and 22B.
- Step 5: You choose all of the left paths, leading you to 11B and 22C.
- Step 6: You choose all of the right paths, leading you to 11Z and 22Z.

So, in this example, you end up entirely on nodes that end in Z after 6 steps.

Simultaneously start on every node that ends with A. How many steps does it take before you're only on nodes that end with Z?

```go
func (m *Maps) ReachXXZ() (steps int) {
	pointer := []string{}
	for key := range m.Right {
		if []rune(key)[2] == 'A' {
			pointer = append(pointer, key)
		}
	}

	found := false
	for !found {

		found = true
		step := m.Instructions[steps%len(m.Instructions)]

		for i := 0; i < len(pointer); i++ {
			if step {
				pointer[i] = m.Right[pointer[i]]
			} else {
				pointer[i] = m.Left[pointer[i]]
			}
		}

		steps++

		for _, ghost := range pointer {
			found = found && []rune(ghost)[2] == 'Z'
		}
	}

	return
}

func TestMaps_ReachXXZ(t *testing.T) {
	t.Run("Example 3", func(t *testing.T) {
		m := &wasteland.Maps{
			Left: map[string]string{
				"11A": "11B",
				"11B": "XXX",
				"11Z": "11B",
				"22A": "22B",
				"22B": "22C",
				"22C": "22Z",
				"22Z": "22B",
				"XXX": "XXX",
			},
			Right: map[string]string{
				"11A": "XXX",
				"11B": "11Z",
				"11Z": "XXX",
				"22A": "XXX",
				"22B": "22C",
				"22C": "22Z",
				"22Z": "22B",
				"XXX": "XXX",
			},
			Instructions: []bool{false, true},
		}
		if gotSteps := m.ReachXXZ(); gotSteps != 6 {
			t.Errorf("Maps.ReachXXZ() = %v, want 6", gotSteps)
		}
	})
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	foundInstructions := false
	maps := wasteland.Maps{
		Left:  map[string]string{},
		Right: map[string]string{},
	}

	for scanner.Scan() {
		line := scanner.Text()

		if !foundInstructions {
			maps.InitInstuctions(line)
			foundInstructions = true
			continue
		}

		if line != "" {
			maps.AddMap(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	log.Printf("count: %d, %d\n", maps.ReachZZZ(), maps.ReachXXZ())
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/)

- [input.txt](/documents/2023-12-08-input.txt)
- [Challenge](https://adventofcode.com/2023/day/8)
