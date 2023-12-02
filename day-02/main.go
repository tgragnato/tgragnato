package main

import (
	"bufio"
	"log"
	"os"

	"github.com/tgragnato/aoc23/day-02/configuration"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		sum1   uint
		sum2   uint
		riddle uint
		c1     = &configuration.Configuration{
			Red:   12,
			Green: 13,
			Blue:  14,
		}
		c2 = &configuration.Configuration{}
	)

	for scanner.Scan() {
		line := scanner.Text()
		riddle++
		if c1.Evaluate(line) {
			sum1 += riddle
		}
		sum2 += c2.Power(line)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}
	log.Printf("sum: %d, %d\n", sum1, sum2)
}
