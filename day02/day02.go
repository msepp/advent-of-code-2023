package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type DiceInHand struct {
	Red   int
	Green int
	Blue  int
}

func (d DiceInHand) String() string {
	return fmt.Sprintf("red: %d, green: %d, blue: %d", d.Red, d.Green, d.Blue)
}

func ParseDiceInHand(colors []string) DiceInHand {
	var res DiceInHand
	for _, set := range colors {
		value := strings.Split(set, " ")
		count, _ := strconv.Atoi(value[0])
		switch value[1] {
		case "red":
			res.Red = count
		case "green":
			res.Green = count
		case "blue":
			res.Blue = count
		}
	}
	return res
}
