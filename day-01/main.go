package main

import (
	"bufio"
	"log"
	"os"

	"github.com/tgragnato/aoc23/day-01/extraction"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum1 := uint(0)
	sum2 := uint(0)
	for scanner.Scan() {
		line := scanner.Text()
		sum1 += extraction.ExtractValues1(line)
		sum2 += extraction.ExtractValues2(line)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}
	log.Printf("sum: %d, %d\n", sum1, sum2)
}
