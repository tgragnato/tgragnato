---
title: If You Give A Seed A Fertilizer
description: Advent of Code 2023 [Day 5]
layout: default
lang: en
---

You take the boat and find the gardener right where you were told he would be: managing a giant "garden" that looks more to you like a farm.

"A water source? Island Island is the water source!" You point out that Snow Island isn't receiving any water.

"Oh, we had to stop the water because we ran out of sand to filter it with! Can't make snow with dirty water. Don't worry, I'm sure we'll get more sand soon; we only turned off the water a few days... weeks... oh no." His face sinks into a look of horrified realization.

"I've been so busy making sure everyone here has food that I completely forgot to check why we stopped getting more sand! There's a ferry leaving soon that is headed over in that direction - it's much faster than your boat. Could you please go check it out?"

You barely have time to agree to this request when he brings up another. "While you wait for the ferry, maybe you can help us with our food production problem. The latest Island Island Almanac just arrived and we're having trouble making sense of it."

The almanac (your puzzle input) lists all of the seeds that need to be planted. It also lists what type of soil to use with each kind of seed, what type of fertilizer to use with each kind of soil, what type of water to use with each kind of fertilizer, and so on. Every type of seed, soil, fertilizer and so on is identified with a number, but numbers are reused by each category - that is, soil 123 and fertilizer 123 aren't necessarily related to each other.

For example:

```
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
```

The almanac starts by listing which seeds need to be planted: seeds 79, 14, 55, and 13.

The rest of the almanac contains a list of maps which describe how to convert numbers from a source category into numbers in a destination category. That is, the section that starts with seed-to-soil map: describes how to convert a seed number (the source) to a soil number (the destination). This lets the gardener and his team know which soil to use with which seeds, which water to use with which fertilizer, and so on.

Rather than list every source number and its corresponding destination number one by one, the maps describe entire ranges of numbers that can be converted. Each line within a map contains three numbers: the destination range start, the source range start, and the range length.

Consider again the example seed-to-soil map:

```
50 98 2
52 50 48
```

The first line has a destination range start of 50, a source range start of 98, and a range length of 2. This line means that the source range starts at 98 and contains two values: 98 and 99. The destination range is the same length, but it starts at 50, so its two values are 50 and 51. With this information, you know that seed number 98 corresponds to soil number 50 and that seed number 99 corresponds to soil number 51.

The second line means that the source range starts at 50 and contains 48 values: 50, 51, ..., 96, 97. This corresponds to a destination range starting at 52 and also containing 48 values: 52, 53, ..., 98, 99. So, seed number 53 corresponds to soil number 55.

Any source numbers that aren't mapped correspond to the same destination number. So, seed number 10 corresponds to soil number 10.

So, the entire list of seed numbers and their corresponding soil numbers looks like this:

```
seed  soil
0     0
1     1
...   ...
48    48
49    49
50    52
51    53
...   ...
96    98
97    99
98    50
99    51
```

With this map, you can look up the soil number required for each initial seed number:

- Seed number 79 corresponds to soil number 81.
- Seed number 14 corresponds to soil number 14.
- Seed number 55 corresponds to soil number 57.
- Seed number 13 corresponds to soil number 13.

The gardener and his team want to get started as soon as possible, so they'd like to know the closest location that needs a seed. Using these maps, find the lowest location number that corresponds to any of the initial seeds. To do this, you'll need to convert each seed number through other categories until you can find its corresponding location number. In this example, the corresponding types are:

- Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
- Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
- Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
- Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.

So, the lowest location number in this example is 35.

What is the lowest location number that corresponds to any of the initial seed numbers?

```go
type Map struct {
	Destination uint
	Source      uint
	Range       uint
}

func (m *Map) GetDestination(source uint) (uint, error) {
	if source < m.Source || source >= m.Source+m.Range {
		return 0, errors.New("oob")
	}

	return source + m.Destination - m.Source, nil
}

func (m *Map) GetSource(destination uint) (uint, error) {
	if destination < m.Destination || destination >= m.Destination+m.Range {
		return 0, errors.New("oob")
	}

	return destination - m.Destination + m.Source, nil
}

type MapSet struct {
	Set []Map
}

func (s *MapSet) GetVertices(destinations []uint) []uint {
	vertices := []uint{}
	for _, mapItem := range s.Set {
		vertices = append(vertices, mapItem.Source)
	}
	for _, destination := range destinations {
		vertices = append(vertices, s.GetSource(destination))
	}
	return vertices
}

func (s *MapSet) GetDestination(source uint) uint {
	for i := range s.Set {
		if val, err := s.Set[i].GetDestination(source); err == nil {
			return val
		}
	}

	return source
}

func (s *MapSet) GetSource(destination uint) uint {
	for i := range s.Set {
		if val, err := s.Set[i].GetSource(destination); err == nil {
			return val
		}
	}

	return 0
}

type RecursiveMapSet struct {
	RecursiveSet []MapSet
	Seeds        map[uint]uint
	sync.Mutex
}

func (r *RecursiveMapSet) GetVertices() []uint {
	vertices := []uint{}

	for i := len(r.RecursiveSet) - 1; i >= 0; i-- {
		vertices = r.RecursiveSet[i].GetVertices(vertices)
	}

	return vertices
}

func (r *RecursiveMapSet) GetDestination(source uint) uint {
	for i := 0; i < len(r.RecursiveSet); i++ {
		source = r.RecursiveSet[i].GetDestination(source)
	}

	return source
}

func (r *RecursiveMapSet) SetSeeds(source []uint) {
	for _, seed := range source {
		r.Seeds[seed] = 0
	}
}

func (r *RecursiveMapSet) GetLowestLocation() uint {
	for seed := range r.Seeds {
		r.Seeds[seed] = r.GetDestination(seed)
	}

	seeds := make([]uint, 0, len(r.Seeds))
	for seed := range r.Seeds {
		seeds = append(seeds, seed)
	}

	sort.Slice(seeds, func(i, j int) bool {
		return r.Seeds[seeds[i]] < r.Seeds[seeds[j]]
	})

	return r.Seeds[seeds[0]]
}

func ParseSeeds(line string) []uint {
	line = strings.Split(line, ": ")[1]
	var parsed []uint = []uint{}

	for _, value := range strings.Split(line, " ") {
		if intvalue, err := strconv.Atoi(value); err == nil {
			parsed = append(parsed, uint(intvalue))
		}
	}

	return parsed
}

func ParseMap(line string) (*Map, error) {
	splitLine := strings.Split(line, " ")

	destination, err := strconv.Atoi(splitLine[0])
	if err != nil {
		return nil, err
	}
	source, err := strconv.Atoi(splitLine[1])
	if err != nil {
		return nil, err
	}
	intrange, err := strconv.Atoi(splitLine[2])
	if err != nil {
		return nil, err
	}

	return &Map{
		Destination: uint(destination),
		Source:      uint(source),
		Range:       uint(intrange),
	}, nil
}

func TestMap_GetDestination(t *testing.T) {
	type fields struct {
		Destination uint
		Source      uint
		Range       uint
	}
	tests := []struct {
		name    string
		fields  fields
		source  uint
		want    uint
		wantErr bool
	}{
		{
			name: "Case 0",
			fields: fields{
				Destination: 52,
				Source:      50,
				Range:       48,
			},
			source:  0,
			want:    0,
			wantErr: true,
		},
		{
			name: "Case 49",
			fields: fields{
				Destination: 52,
				Source:      50,
				Range:       48,
			},
			source:  49,
			want:    0,
			wantErr: true,
		},
		{
			name: "Case 50",
			fields: fields{
				Destination: 52,
				Source:      50,
				Range:       48,
			},
			source:  50,
			want:    52,
			wantErr: false,
		},
		{
			name: "Case 97",
			fields: fields{
				Destination: 52,
				Source:      50,
				Range:       48,
			},
			source:  97,
			want:    99,
			wantErr: false,
		},
		{
			name: "Case 98",
			fields: fields{
				Destination: 52,
				Source:      50,
				Range:       48,
			},
			source:  98,
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &seed.Map{
				Destination: tt.fields.Destination,
				Source:      tt.fields.Source,
				Range:       tt.fields.Range,
			}
			got, err := m.GetDestination(tt.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Map.GetDestination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Map.GetDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}

var testSet = seed.MapSet{
	[]seed.Map{
		{
			Destination: 50,
			Source:      98,
			Range:       2,
		},
		{
			Destination: 52,
			Source:      50,
			Range:       48,
		},
	},
}

func TestMapSet_GetDestination(t *testing.T) {
	tests := []struct {
		name   string
		Set    seed.MapSet
		source uint
		want   uint
	}{
		{
			name:   "Case 0",
			Set:    testSet,
			source: 0,
			want:   0,
		},
		{
			name:   "Case 49",
			Set:    testSet,
			source: 49,
			want:   49,
		},
		{
			name:   "Case 50",
			Set:    testSet,
			source: 50,
			want:   52,
		},
		{
			name:   "Case 97",
			Set:    testSet,
			source: 97,
			want:   99,
		},
		{
			name:   "Case 98",
			Set:    testSet,
			source: 98,
			want:   50,
		},
		{
			name:   "Case 99",
			Set:    testSet,
			source: 99,
			want:   51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Set.GetDestination(tt.source); got != tt.want {
				t.Errorf("MapSet.GetDestination() = %v, %v => %v", got, tt.source, tt.want)
			}
		})
	}
}

func TestParseSeeds(t *testing.T) {
	t.Run("seeds: 79 14 55 13", func(t *testing.T) {
		if got := seed.ParseSeeds("seeds: 79 14 55 13"); !reflect.DeepEqual(got, []uint{79, 14, 55, 13}) {
			t.Errorf("ParseSeeds(seeds: 79 14 55 13) = %v", got)
		}
	})
}

func TestRecursiveMapSet_GetLowestLocation(t *testing.T) {
	t.Run("Simplest example given", func(t *testing.T) {
		r := &seed.RecursiveMapSet{
			RecursiveSet: []seed.MapSet{testSet},
			Seeds:        map[uint]uint{79: 0, 14: 0, 55: 0, 13: 0},
		}
		if got := r.GetLowestLocation(); got != 13 {
			t.Errorf("RecursiveMapSet.GetDestination() = %v, want 13", got)
		}
	})
}
```

Everyone will starve if you only plant such a small number of seeds. Re-reading the almanac, it looks like the seeds: line actually describes ranges of seed numbers.

The values on the initial seeds: line come in pairs. Within each pair, the first value is the start of the range and the second value is the length of the range. So, in the first line of the example above:

seeds: 79 14 55 13

This line describes two ranges of seed numbers to be planted in the garden. The first range starts with seed number 79 and contains 14 values: 79, 80, ..., 91, 92. The second range starts with seed number 55 and contains 13 values: 55, 56, ..., 66, 67.

Now, rather than considering four seed numbers, you need to consider a total of 27 seed numbers.

In the above example, the lowest location number can be obtained from seed number 82, which corresponds to soil 84, fertilizer 84, water 84, light 77, temperature 45, humidity 46, and location 46. So, the lowest location number is 46.

Consider all of the initial seed numbers listed in the ranges on the first line of the almanac. What is the lowest location number that corresponds to any of the initial seed numbers?

```go
func (r *RecursiveMapSet) MangleSeeds() {
	seedSlice := []uint{}
	for seed := range r.Seeds {
		seedSlice = append(seedSlice, seed)
	}

	vertices := r.GetVertices()
	mangled := map[uint]uint{}

	for i := 0; i < len(seedSlice)-1; i += 2 {

		for _, vertex := range vertices {
			if vertex >= seedSlice[i] && vertex < seedSlice[i]+seedSlice[i+1] {
				destination := r.GetDestination(vertex)
				if destination == 0 || vertex == 0 {
					continue
				}
				mangled[vertex] = destination
			}
		}

		mangled[seedSlice[i]] = r.GetDestination(seedSlice[i])
		mangled[seedSlice[i]+seedSlice[i+1]] = r.GetDestination(seedSlice[i] + seedSlice[i+1])
	}

	r.Seeds = mangled
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		rms seed.RecursiveMapSet = seed.RecursiveMapSet{
			Seeds: map[uint]uint{},
		}
		ms seed.MapSet = seed.MapSet{}
	)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			rms.SetSeeds(seed.ParseSeeds(line))
			continue
		}

		if line == "" && len(ms.Set) != 0 {
			rms.RecursiveSet = append(rms.RecursiveSet, ms)
			continue
		}

		if mapItem, err := seed.ParseMap(line); err == nil {
			ms.Set = append(ms.Set, *mapItem)
		} else {
			ms = seed.MapSet{}
		}
	}

	rms.RecursiveSet = append(rms.RecursiveSet, ms)

	lowestLocation1 := rms.GetLowestLocation()
	rms.MangleSeeds()
	lowestLocation2 := rms.GetLowestLocation()

	log.Printf("sum: %d, %d\n", lowestLocation1, lowestLocation2)
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/)

- [input.txt](/documents/2023-12-05-input.txt)
- [Challenge](https://adventofcode.com/2023/day/5)
