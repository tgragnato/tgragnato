---
title: If You Give A Seed A Fertilizer
description: Advent of Code 2023 [Day 5]
layout: default
lang: en
tag: aoc23
prefetch:
  - adventofcode.com
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

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2023/day/5)

<details>
	<summary>Click to show the input</summary>
	<pre>
seeds: 28965817 302170009 1752849261 48290258 804904201 243492043 2150339939 385349830 1267802202 350474859 2566296746 17565716 3543571814 291402104 447111316 279196488 3227221259 47952959 1828835733 9607836

seed-to-soil map:
3680121696 1920754815 614845600
1920754815 3846369604 448597692
193356576 570761634 505124585
2369352507 2535600415 31531965
2400884472 2567132380 1279237224
0 459278395 111483239
698481161 97868205 361410190
1059891351 0 15994868
111483239 15994868 81873337

soil-to-fertilizer map:
1633669237 1273301814 72865265
2398515176 2671190790 99210785
2397916384 3018946373 598792
4034325916 3061716397 20017393
3298612516 3793795301 14249501
4030007411 3051046904 2833129
1906984482 224872691 14620134
864506893 1590633724 149044542
1029530319 442871336 36727018
1921604616 770934113 68546178
3560536321 3114405501 28822192
1019762634 1263534129 9767685
3852235341 3579014714 60339892
2385228698 1577946038 12687686
2234322470 239492825 150906228
0 170310676 54562015
3208946111 3808044802 89666405
1209615399 839480291 424053838
4032840540 4041982568 1485376
2497725961 2174737461 293042810
2002543511 1346167079 231778959
3312862017 3475611771 103402943
318739997 1739678266 354749094
1013551435 3012735174 6211199
4014277153 4160859076 15730258
3589358513 3143227693 230682158
1990150794 2467780271 12392717
3051046904 3081733790 32671711
3820040671 3761600631 32194670
148429321 0 170310676
673489091 2480172988 191017802
1066257337 627576051 143358062
2790768771 2770401575 242333599
3091554979 4043467944 117391132
3416264960 3897711207 144271361
3912575233 3373909851 101701920
4072291714 3639354606 104297620
3083718615 3053880033 7836364
54562015 3019545165 13557205
1759006785 479598354 147977697
68119220 2094427360 80310101
1706534502 390399053 52472283
4054343309 3743652226 17948405

fertilizer-to-water map:
0 1095885172 129797665
2661548513 1044284418 17872363
3282164642 3678907615 214830258
1440687421 2218635146 325889720
3496994900 4208791298 25912548
3253828209 4136945159 5561683
1797056017 864689597 109403664
3259389892 4186016548 22774750
2578517508 1225682837 83031005
3193832718 3618912124 59995491
3695649169 3214450646 211194594
820325042 974093261 70191157
1284591017 1074888739 20996433
2929761569 3893737873 85668135
1305587450 2605461705 73959171
2168339930 1062156781 12731958
2465234843 2135490203 52666067
3522907448 4108091872 882860
3523790308 4255675252 39292044
2517900910 1308713842 60616598
3563082352 3176039879 38410767
3015429704 3979406008 128685864
1913427402 2131043562 4446641
2235159285 1419841841 190190495
3673886186 4108974732 21762983
891504291 1610032336 393086726
3927815169 3425645240 193266884
4121082053 3002154636 173885243
2084864581 2004107154 83475349
1917874043 557612753 69524983
890516199 2003119062 988092
766355924 2551492587 53969118
1379546621 0 61140800
3187625274 4130737715 6207444
528804063 627137736 237551861
1987399026 230411125 97465555
1906459681 2544524866 6967721
2184647884 1369330440 50511401
2425349780 2091158499 39885063
2181071888 2087582503 3575996
3144115568 4142506842 43509706
1766577141 2188156270 30478876
359533738 61140800 169270325
129797665 327876680 229736073
3906843763 4234703846 20971406
3601493119 2929761569 72393067

water-to-light map:
2375927917 1595026882 126334140
1307603095 818620477 43777869
2050676589 1855896418 112224406
3618302244 2909504698 119958941
3078570200 3088215627 6211083
3084781283 3094426710 141266337
524666822 53020621 149058240
673725062 862398346 147671362
2364320682 2577001713 11607235
1941578413 1584221500 10805382
2162900995 2536766467 40235246
162015400 237365123 4480592
821396424 241845715 141336168
166495992 1138498800 212882164
4277433486 4220367555 17533810
3226047620 2229635217 307131250
2909428734 1968120824 34606070
1885573816 3954749082 56004597
2711875933 2868267590 41237108
0 726306378 92314099
2944034804 1721361022 134535396
1584221500 3392008740 301352316
962732592 34415039 18605582
2235705153 2101019688 128615529
4252936467 4237901365 24497019
92314099 202078861 35286262
981338174 400041457 326264921
379378156 1010069708 128429092
3591930858 3693361056 26371386
3785017329 3719732442 235016640
2502262057 4010753679 209613876
3533178870 3029463639 58751988
3738261185 2821511446 46756144
4020033969 2806534555 14976891
1952383795 2002726894 98292794
127600361 0 34415039
2753113041 3235693047 156315693
507807248 383181883 16859574
2203136241 4262398384 32568912
4035010860 2588608948 217925607

light-to-temperature map:
2137189745 1335050925 100355790
639139367 2440321747 987829
1663612748 1778059435 153830272
1122754252 1950103191 82536600
1929621334 1199531530 135519395
1286703174 2032639791 207137687
245313533 981575774 217955756
2597564380 2824691125 293777778
895004176 331442633 25226735
1493840861 236388681 616173
1494457034 764560381 107637728
1817443020 1435406715 112178314
1205290852 356669368 33552643
474799702 0 164339665
2341054397 2260378974 100255179
1043066658 2360634153 79687594
125852143 390222011 119461390
3924383937 3130691909 13614218
2467721984 3747288823 76649669
2065140729 164339665 72049016
920230911 1673437172 104622263
640127196 509683401 254876980
1024853174 1931889707 18213484
2331983314 2251307891 9071083
2237545535 237004854 94437779
3912160931 3118468903 12223006
1238843495 933716095 47859679
3326002417 3517222025 230066798
3556069215 3144306127 241363224
3797432439 4180238804 114728492
1602094762 872198109 61517986
2544371653 4127046077 53192727
2891342158 3823938492 303107585
463269289 2239777478 11530413
3194449743 3385669351 131552674
3937998155 2467721984 356969141
0 1547585029 125852143

temperature-to-humidity map:
2687600833 2313887435 187105587
3281196981 2291603041 22284394
1771250828 1899269239 314167725
784031720 478456148 306959384
2605226464 1771250828 58348072
2085418553 3793564740 111907603
1090991104 785415532 575136195
3437652344 1829598900 69670339
2874706420 2500993022 389039942
3303481375 3905472343 134170969
305575572 0 478456148
3263746362 2890032964 17450619
2527060387 2213436964 78166077
2428623843 3695128196 98436544
2663574536 4270940999 24026297
3507322683 2907483583 787644613
2197326156 4039643312 231297687
0 1360551727 305575572

humidity-to-location map:
1919184105 1156349110 51114849
4031284281 3411510751 25609498
0 171183359 79004094
1253227229 2072782209 122019778
4056893779 3437120249 136289693
3402931364 4156827458 101778985
84557792 1207463959 134801591
635909965 1371746366 266495395
4029464617 4127764171 1819664
4193183472 2857352625 101783824
1375247007 2200355685 41445634
1996492203 0 171183359
3601595563 3699895117 427869054
2218993186 1133540977 22808133
3217192942 2959136449 140385316
2987922009 4258606443 9236491
2628749093 2543337773 86365212
2167675562 369831582 51317624
3504710349 2446452559 96885214
902405360 421149206 108869392
3357578258 2811999519 45353106
1196458443 1638241761 56768786
1178674352 693035436 17784091
1970298954 530018598 26193249
2868723842 3099521765 91954544
1522874936 895259169 18612073
2841599480 4267842934 27124362
1880563756 1695010547 38620349
1011274752 913871242 137918784
219359383 556211847 136823589
1785350971 250187453 95212785
2960678386 4129583835 27243623
540622614 1977494858 95287351
356182972 710819527 184439642
1498443592 345400238 24431344
1416692641 1051790026 81750951
1541487009 1733630896 243863962
79004094 2194801987 5553698
1149193536 1342265550 29480816
2997158500 3191476309 220034442
2715114305 3573409942 126485175
2446452559 2629702985 182296534
	</pre>
</details>
