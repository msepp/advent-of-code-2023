package main

import (
	"bufio"
	"log"
	"os"
)

type token struct {
	pattern []rune
	pos     int
	value   int32
}

func (t *token) feed(c rune) (int32, bool) {
	if t.pattern[t.pos] == c {
		t.pos++
	} else if t.pattern[0] == c {
		t.pos = 1
	} else {
		t.pos = 0
	}
	if t.pos == len(t.pattern) {
		t.pos = 0
		return t.value, true
	}
	return 0, false
}

func main() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var sum int32
	for scanner.Scan() {
		tokens := []*token{
			{pattern: []rune("zero"), value: '0'},
			{pattern: []rune("one"), value: '1'},
			{pattern: []rune("two"), value: '2'},
			{pattern: []rune("three"), value: '3'},
			{pattern: []rune("four"), value: '4'},
			{pattern: []rune("five"), value: '5'},
			{pattern: []rune("six"), value: '6'},
			{pattern: []rune("seven"), value: '7'},
			{pattern: []rune("eight"), value: '8'},
			{pattern: []rune("nine"), value: '9'},
		}

		s := scanner.Text()
		firstDigit := int32(-1)
		lastDigit := int32(-1)
		for _, c := range s {
			var wordVal int32
			// move all tokens forward. pick one if a match is found.
			for _, t := range tokens {
				if r, match := t.feed(c); match {
					wordVal = r
				}
			}
			if wordVal > 0 {
				c = wordVal
			}
			if c < '0' || c > '9' {
				continue
			}
			d := c - '0'
			if firstDigit == -1 {
				firstDigit = d * 10
			}
			lastDigit = d
		}
		sum += firstDigit + lastDigit
	}
	log.Printf("result: %d", sum)
}
