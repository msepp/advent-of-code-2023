package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		winners, found := parseLine(scanner.Text())
		matches := matchingNumbers(winners, found)
		if len(matches) == 0 {
			continue
		}
		points := 1 << (len(matches) - 1)
		fmt.Printf("win: %v\ngot: %v\nmatches: %v\npoints: %d\n", winners, found, matches, points)
		sum += points
	}
	log.Printf("Result: %d", sum)
}

func parseLine(line string) ([]int, []int) {
	pcs := strings.SplitN(strings.SplitN(line, ": ", 2)[1], " | ", 2)
	winners := strToInts(pcs[0])
	found := strToInts(pcs[1])
	return winners, found
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
