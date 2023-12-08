package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/tgragnato/aoc23/day-05/seed"
)

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
