package main

import (
	"bufio"
	"log"
	"os"

	"github.com/tgragnato/aoc23/day-08/wasteland"
)

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
