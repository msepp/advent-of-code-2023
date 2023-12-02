package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/msepp/advent-of-code-2023/day02"
)

func main() {
	limits := day02.DiceInHand{Red: 12, Green: 13, Blue: 14}
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

func validateGame(grabs []string, limit day02.DiceInHand) bool {
	for _, colors := range grabs {
		hand := day02.ParseDiceInHand(strings.Split(colors, ", "))
		if hand.Red > limit.Red || hand.Green > limit.Green || hand.Blue > limit.Blue {
			return false
		}
	}
	return true
}
