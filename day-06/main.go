package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tgragnato/aoc23/day-06/wait"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		time      []uint      = []uint{}
		distance  []uint      = []uint{}
		races     []wait.Race = []wait.Race{}
		pointer               = &time
		giantRace wait.Race   = wait.Race{
			Boat: wait.Boat{
				Starting: 0,
				Increase: 1,
			},
		}
		times     []rune = []rune{}
		distances []rune = []rune{}
	)

	for scanner.Scan() {
		for _, value := range strings.Split(strings.Split(scanner.Text(), ":")[1], " ") {
			if number, err := strconv.Atoi(value); err == nil {
				*pointer = append(*pointer, uint(number))
			}
		}
		pointer = &distance
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	if len(time) != len(distance) {
		log.Fatalln("len(time) != len(distance)")
	}

	for i := 0; i < len(time); i++ {
		races = append(races, wait.Race{
			Boat: wait.Boat{
				Starting: 0,
				Increase: 1,
			},
			Time:     time[i],
			Distance: distance[i],
		})

		times = append(times, []rune(strconv.Itoa(int(time[i])))...)
		distances = append(distances, []rune(strconv.Itoa(int(distance[i])))...)
	}

	mul1 := uint(1)
	for _, race := range races {
		mul1 *= race.GetWinningTimes()
	}

	timeValue, _ := strconv.Atoi(string(times))
	giantRace.Time = uint(timeValue)
	distanceValue, _ := strconv.Atoi(string(distances))
	giantRace.Distance = uint(distanceValue)
	tim2 := giantRace.GetWinningTimes()

	log.Printf("sum: %d, %d\n", mul1, tim2)
}
