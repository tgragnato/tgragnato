package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tgragnato/aoc23/day-07/camelcards"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	handList := camelcards.HandList{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		cards := []rune(line[0])

		bid, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalln(err.Error())
		}

		handList.Hands = append(handList.Hands, camelcards.Hand{
			Cards: [5]rune(cards),
			Bid:   uint(bid),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	points1 := handList.GetWinnings(false)
	dump1 := handList.String()
	points2 := handList.GetWinnings(true)
	dump2 := handList.String()
	log.Printf("\n---\n%s---\n%s---\nsum: %d, %d\n", dump1, dump2, points1, points2)
}
