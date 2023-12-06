package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type scratchCard struct {
	number  int
	matches int
}

func (c scratchCard) winner() bool {
	return c.matches > 0
}

func main() {
	t0 := time.Now()
	file, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var (
		originalCards []scratchCard
		allCopies     []scratchCard
	)
	for scanner.Scan() {
		originalCards = append(originalCards, parseLine(scanner.Text()))
	}
	winners := getWinners(originalCards)
	for len(winners) > 0 {
		var newCopies []scratchCard
		for _, v := range winners {
			sIdx := v.number
			if sIdx >= len(originalCards) {
				continue
			}
			eIdx := sIdx + v.matches
			if eIdx > len(originalCards) {
				eIdx = len(originalCards)
			}
			newCopies = append(newCopies, originalCards[sIdx:eIdx]...)
		}
		allCopies = append(allCopies, newCopies...)
		winners = getWinners(newCopies)
	}

	log.Printf("Result: %d (duration: %q)", len(allCopies)+len(originalCards), time.Since(t0))
}

func getWinners(cards []scratchCard) []scratchCard {
	var result []scratchCard
	for _, v := range cards {
		if !v.winner() {
			continue
		}
		result = append(result, v)
	}
	return result
}

func parseLine(line string) scratchCard {
	var card scratchCard

	linePcs := strings.SplitN(line, ": ", 2)
	card.number, _ = strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(linePcs[0], "Card ")))
	numberSets := strings.SplitN(linePcs[1], " | ", 2)
	winners := strToInts(numberSets[0])
	found := strToInts(numberSets[1])
	card.matches = len(matchingNumbers(winners, found))
	return card
}

func strToInts(s string) []int {
	pcs := strings.Split(s, " ")
	var ints []int
	for _, v := range pcs {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		d, _ := strconv.Atoi(v)
		ints = append(ints, d)
	}
	return ints
}

func matchingNumbers(a, b []int) []int {
	var found []int
	var needle, haystack []int
	if len(a) > len(b) {
		haystack = b
		needle = a
	} else {
		haystack = a
		needle = b
	}
	for _, n := range needle {
		if !slices.Contains(haystack, n) {
			continue
		}
		found = append(found, n)
	}
	return found
}
