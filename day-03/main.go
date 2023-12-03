package main

import (
	"bufio"
	"log"
	"os"

	"github.com/tgragnato/aoc23/day-03/engine"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		engineSchematic *(engine.Schematic) = &engine.Schematic{
			Schema: [][]rune{},
		}
	)

	for scanner.Scan() {
		runes := []rune(scanner.Text())
		runes = append(runes, '.')
		engineSchematic.Schema = append(engineSchematic.Schema, runes)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}
	log.Printf(
		"sum: %d, %d\n",
		engineSchematic.SumPartNumber(),
		engineSchematic.SumGearRatio(),
	)
}
