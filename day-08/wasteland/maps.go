package wasteland

import (
	"log"
	"strings"
)

type Maps struct {
	Left         map[string]string
	Right        map[string]string
	Instructions []bool
}

func (m *Maps) ReachZZZ() (steps uint) {
	pointer := "AAA"
	for pointer != "ZZZ" {
		for _, step := range m.Instructions {
			if step {
				pointer = m.Right[pointer]
			} else {
				pointer = m.Left[pointer]
			}
			steps++
			if pointer == "ZZZ" {
				break
			}
		}
	}
	return
}

func (m *Maps) InitInstuctions(line string) {
	m.Instructions = []bool{}
	for _, char := range []rune(line) {
		switch char {
		case 'L':
			m.Instructions = append(m.Instructions, false)
		case 'R':
			m.Instructions = append(m.Instructions, true)
		default:
			log.Fatalln("Parsing error")
		}
	}
}

func (m *Maps) AddMap(line string) {
	splittedLine := strings.Split(line, " = (")
	splittedDestinations := strings.Split(splittedLine[1], ", ")
	m.Left[splittedLine[0]] = splittedDestinations[0]
	m.Right[splittedLine[0]] = string([]rune(splittedDestinations[1])[0:3])
}

func (m *Maps) ReachXXZ() (steps int) {
	pointer := []string{}
	for key := range m.Right {
		if []rune(key)[2] == 'A' {
			pointer = append(pointer, key)
		}
	}

	found := false
	for !found {

		found = true
		step := m.Instructions[steps%len(m.Instructions)]

		for i := 0; i < len(pointer); i++ {
			if step {
				pointer[i] = m.Right[pointer[i]]
			} else {
				pointer[i] = m.Left[pointer[i]]
			}
		}

		steps++

		for _, ghost := range pointer {
			found = found && []rune(ghost)[2] == 'Z'
		}
	}

	return
}
