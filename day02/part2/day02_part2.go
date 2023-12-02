package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/msepp/advent-of-code-2023/day02"
)

func main() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		s := scanner.Text()
		pcs := strings.Split(s[5:], ": ")
		res := findMinimumDice(strings.Split(pcs[1], "; "))
		sum += res.Blue * res.Green * res.Red
	}
	log.Printf("Result: %d", sum)
}

func findMinimumDice(grabs []string) day02.DiceInHand {
	smallest := day02.DiceInHand{Red: 0, Green: 0, Blue: 0}
	for _, colors := range grabs {
		grabbed := day02.ParseDiceInHand(strings.Split(colors, ", "))
		if grabbed.Red > smallest.Red {
			smallest.Red = grabbed.Red
		}
		if grabbed.Green > smallest.Green {
			smallest.Green = grabbed.Green
		}
		if grabbed.Blue > smallest.Blue {
			smallest.Blue = grabbed.Blue
		}
	}
	return smallest
}
