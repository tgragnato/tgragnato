package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func toList(line string) (groups []int) {
	for _, val := range strings.Split(line, ",") {
		n, _ := strconv.ParseUint(val, 10, 8)
		groups = append(groups, int(n))
	}
	return
}

type Store map[string]uint

func (s Store) salva(line string, num []int) uint {
	line = strings.Trim(line, ".")

	key := line + fmt.Sprintf("%v", num)
	if val, ok := s[key]; ok {
		return val
	}

	s[key] = s.disposizioni(line, num, true)
	return s[key]
}

func (s Store) disposizioni(line string, num []int, variant bool) (sum uint) {
	line = strings.Trim(line, ".")

	if line == "" && len(num) == 0 {
		return 1
	}

	if line == "" {
		return 0
	}

	if len(num) == 0 && strings.Contains(line, "#") {
		return 0
	}

	if len(num) == 0 {
		return 1
	}

	if line[0] == '?' {
		if variant {
			sum += s.salva(line[1:], num)
		} else {
			sum += s.disposizioni(line[1:], num, false)
		}
		line = "#" + line[1:]
	}

	if len(line) < num[0] || strings.ContainsRune(line[:num[0]], '.') {
		return
	}

	if len(line) > num[0] {
		switch line[num[0]] {
		case '#':
			return
		case '?':
			line = line[:num[0]] + "." + line[num[0]+1:]
		}
	}

	if variant {
		return sum + s.salva(line[num[0]:], num[1:])
	}
	return sum + s.disposizioni(line[num[0]:], num[1:], false)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		sum1 uint = 0
		sum2 uint = 0
	)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		sum1 += Store{}.disposizioni(fields[0], toList(fields[1]), false)

		conditions := strings.Repeat(fields[0]+"?", 5)
		conditions = conditions[:len(conditions)-1]

		groupsStr := strings.Repeat(fields[1]+",", 5)
		groupsStr = groupsStr[:len(groupsStr)-1]

		sum2 += Store{}.salva(conditions, toList(groupsStr))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	log.Printf(
		"sum: %d, %d\n",
		sum1,
		sum2,
	)
}
