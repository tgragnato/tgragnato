package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	list := [][]uint{}
	for scanner.Scan() {
		sequence := []uint{}
		for _, word := range strings.Split(scanner.Text(), " ") {
			number, err := strconv.Atoi(word)
			if err != nil {
				log.Println(err.Error())
			}
			sequence = append(sequence, uint(number))
		}
		list = append(list, sequence)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	var (
		sum1 uint = 0
		sum2 uint = 0
	)

	for _, row := range list {
		ar1 := buildTriangle(row)
		ar2 := buildTriangle(row)

		ar1[len(ar1)-1] = append(ar1[len(ar1)-1], 0)
		ar2[len(ar2)-1] = append(ar2[len(ar2)-1], 0)

		for i := len(ar1) - 2; i >= 0; i-- {
			ar1[i] = append(ar1[i], ar1[i][len(ar1[i])-1]+ar1[i+1][len(ar1[i+1])-1])
		}

		for j := len(ar2) - 2; j >= 0; j-- {
			ar2[j] = append([]uint{ar2[j][0] - ar2[j+1][0]}, ar2[j]...)
		}

		sum1 += ar1[0][len(ar1[0])-1]
		sum2 += ar2[0][0]
	}

	log.Printf("count: %d, %d\n", sum1, sum2)
}

func buildTriangle(firstSeq []uint) [][]uint {
	p := [][]uint{firstSeq}
	for {
		last := len(p) - 1

		done := true
		newRow := make([]uint, len(p[last])-1)
		for i := 0; i < len(p[last])-1; i++ {
			newRow[i] = p[last][i+1] - p[last][i]
			if newRow[i] != 0 {
				done = false
			}
		}
		p = append(p, newRow)

		if done {
			return p
		}
	}
}
