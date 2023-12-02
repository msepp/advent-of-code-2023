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
	limits := diceInHand{red: 12, green: 13, blue: 14}
	file, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		s := scanner.Text()
		pcs := strings.Split(s[5:], ": ")
		if valid := validateGame(strings.Split(pcs[1], "; "), limits); !valid {
			log.Printf("Game %s is NOT valid", pcs[0])
			continue
		}
		log.Printf("Game %s is VALID", pcs[0])
		gameID, _ := strconv.Atoi(pcs[0])
		sum += gameID
	}
	log.Printf("Result: %d", sum)
}

func validateGame(grabs []string, limit diceInHand) bool {
	for _, colors := range grabs {
		hand := parseDiceInHand(strings.Split(colors, ", "))
		if hand.red > limit.red || hand.green > limit.green || hand.blue > limit.blue {
			return false
		}
	}
	return true
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
