package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

type Point struct{ x, y int }

func findEmpty(galaxies []Point) (map[int]struct{}, map[int]struct{}) {
	eRows := make(map[int]struct{})
	eCols := make(map[int]struct{})

	for i := 0; i < 140; i++ {
		foundRow, foundCol := false, false

		for _, g := range galaxies {
			if g.x == i {
				foundRow = true
			}
			if g.y == i {
				foundCol = true
			}
		}

		if !foundRow {
			eRows[i] = struct{}{}
		}
		if !foundCol {
			eCols[i] = struct{}{}
		}
	}
	return eRows, eCols
}

func solution(galaxies []Point, Expansion uint) uint {
	eRows, eCols := findEmpty(galaxies)

	sum := uint(0)
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {

			d1 := math.Abs(float64(galaxies[i].x - galaxies[j].x))
			d2 := math.Abs(float64(galaxies[i].y - galaxies[j].y))
			distance := uint(d1 + d2)

			min := galaxies[i].x
			max := galaxies[j].x
			if max < min {
				max = galaxies[i].x
				min = galaxies[j].x
			}

			for dr := min + 1; dr < max; dr++ {
				if _, ok := eRows[dr]; ok {
					distance += Expansion - 1
				}
			}

			min = galaxies[i].y
			max = galaxies[j].y
			if max < min {
				max = galaxies[i].y
				min = galaxies[j].y
			}

			for dc := min + 1; dc < max; dc++ {
				if _, ok := eCols[dc]; ok {
					distance += Expansion - 1
				}
			}

			sum += distance
		}
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var galaxies []Point
	for x := 0; scanner.Scan(); x++ {
		line := scanner.Text()
		for y, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Point{x, y})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	log.Printf(
		"count: %d, %d\n",
		solution(galaxies, 2),
		solution(galaxies, 1000000),
	)
}
