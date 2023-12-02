package configuration

import (
	"log"
	"strconv"
	"strings"
)

type Configuration struct {
	Red   uint
	Blue  uint
	Green uint
}

func (c *Configuration) Evaluate(line string) bool {
	line = strings.Split(line, ": ")[1]
	lastValue := uint(0)

	for _, set := range strings.Split(line, "; ") {
		for _, tuple := range strings.Split(set, ", ") {
			for _, word := range strings.Split(tuple, " ") {
				switch word {
				case "red":
					if lastValue > c.Red {
						return false
					}
				case "blue":
					if lastValue > c.Blue {
						return false
					}
				case "green":
					if lastValue > c.Green {
						return false
					}
				default:
					if value, err := strconv.Atoi(word); err == nil {
						lastValue = uint(value)
					} else {
						log.Println(err.Error())
					}
				}
			}
		}
	}

	return true
}

func (c *Configuration) Power(line string) uint {
	c.Red = 0
	c.Blue = 0
	c.Green = 0

	line = strings.Split(line, ": ")[1]
	lastValue := uint(0)

	for _, set := range strings.Split(line, "; ") {
		for _, tuple := range strings.Split(set, ", ") {
			for _, word := range strings.Split(tuple, " ") {
				switch word {
				case "red":
					if lastValue > c.Red {
						c.Red = lastValue
					}
				case "blue":
					if lastValue > c.Blue {
						c.Blue = lastValue
					}
				case "green":
					if lastValue > c.Green {
						c.Green = lastValue
					}
				default:
					if value, err := strconv.Atoi(word); err == nil {
						lastValue = uint(value)
					} else {
						log.Println(err.Error())
					}
				}
			}
		}

	}

	return c.Red * c.Blue * c.Green
}
