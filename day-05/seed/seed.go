package seed

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"sync"
)

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
