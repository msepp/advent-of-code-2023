package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type partNo struct {
	mask  *mask192
	value []rune
}

func (p partNo) int() int {
	d, _ := strconv.Atoi(string(p.value))
	return d
}

type mask192 struct {
	low, mid, high uint64
}

func (m *mask192) intersect(with *mask192) bool {
	if m == nil || with == nil {
		return false
	}
	return (m.low&with.low)+(m.mid&with.mid)+(m.high&with.high) > 0
}

func (m *mask192) set(pos int) {
	if m == nil || pos < 0 {
		return
	}
	switch {
	case pos < 64:
		m.low = m.low | (1 << (63 - pos))
	case pos < 128:
		pos = pos - 64
		m.mid = m.mid | (1 << (63 - pos))
	case pos < 192:
		pos = pos - 128
		m.high = m.high | (1 << (63 - pos))
	}
}

func (m *mask192) String() string {
	return fmt.Sprintf("%064b%064b%064b\n", m.low, m.mid, m.high)[:140]
}

func isSymbol(r rune) bool {
	return r != '.' && (r < '0' || r > '9')
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func collectMatches(rows []*mask192, candidates []*partNo) int {
	sum := 0
	for _, p := range candidates {
		if rows[0].intersect(p.mask) || rows[1].intersect(p.mask) || rows[2].intersect(p.mask) {
			sum += p.int()
		}
	}
	return sum
}

func main() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var (
		partSum     int
		candidates  []*partNo
		currentLine *mask192
		lineState   = make([]*mask192, 3)
		scanner     = bufio.NewScanner(file)
	)
	for scanner.Scan() {
		currentLine = &mask192{}
		var (
			newParts    []*partNo
			currentPart *partNo
			line        = scanner.Text()
		)
		for pos, r := range line {
			if !isDigit(r) {
				if isSymbol(r) {
					currentLine.set(pos)
				}
				if currentPart != nil {
					currentPart.mask.set(pos)
					currentPart = nil
				}
				continue
			}
			if currentPart == nil {
				currentPart = &partNo{mask: &mask192{}}
				currentPart.mask.set(pos - 1)
				newParts = append(newParts, currentPart)
			}
			currentPart.mask.set(pos)
			currentPart.value = append(currentPart.value, r)
		}
		// move state forward as line is processed, current line is the next line
		// for previous candidates.
		lineState[0], lineState[1], lineState[2] = lineState[1], lineState[2], currentLine
		// collect sum from previous candidates
		partSum += collectMatches(lineState, candidates)
		// new parts become the next set of candidates
		candidates = newParts
	}
	// move the state forward again to be able to check previous candidates.
	lineState[0], lineState[1], lineState[2] = lineState[1], lineState[2], currentLine
	partSum += collectMatches(lineState, candidates)
	log.Printf("Result: %d", partSum)
}
