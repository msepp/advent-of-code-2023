package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type diceInHand struct {
	red   int
	green int
	blue  int
}

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
		sum += res.blue * res.green * res.red
	}
	log.Printf("Result: %d", sum)
}

func findMinimumDice(grabs []string) diceInHand {
	smallest := diceInHand{red: 0, green: 0, blue: 0}
	for _, colors := range grabs {
		grabbed := parseDiceInHand(strings.Split(colors, ", "))
		if grabbed.red > smallest.red {
			smallest.red = grabbed.red
		}
		if grabbed.green > smallest.green {
			smallest.green = grabbed.green
		}
		if grabbed.blue > smallest.blue {
			smallest.blue = grabbed.blue
		}
	}
	return smallest
}

func parseDiceInHand(colors []string) diceInHand {
	var res diceInHand
	for _, set := range colors {
		value := strings.Split(set, " ")
		count, _ := strconv.Atoi(value[0])
		switch value[1] {
		case "red":
			res.red = count
		case "green":
			res.green = count
		case "blue":
			res.blue = count
		}
	}
	return res
}
