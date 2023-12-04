package main

import (
	"bufio"
	"log"
	"os"

	"github.com/tgragnato/aoc23/day-04/scratchcards"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sl := &scratchcards.ScratchcardList{}
	sl.Populate(lines)

	log.Printf("sum: %d, %d\n", sl.Points(), sl.Matches())
}
