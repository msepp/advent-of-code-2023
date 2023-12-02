package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var sum int32
	for scanner.Scan() {
		firstDigit := int32(-1)
		lastDigit := int32(-1)
		s := scanner.Text()
		for _, c := range s {
			if c < '0' || c > '9' {
				continue
			}
			d := c - '0'
			if firstDigit == -1 {
				firstDigit = d * 10
			}
			lastDigit = d
			log.Printf("replaced: %q: %d", s, firstDigit+lastDigit)
		}
		sum += firstDigit + lastDigit
	}
	log.Printf("result: %d", sum)
}
