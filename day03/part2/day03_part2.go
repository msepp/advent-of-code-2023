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

func collectMatches(rows []*mask192, candidates []*partNo) []*partNo {
	var matches []*partNo
	for i, p := range candidates {
		if rows[0].intersect(p.mask) || rows[1].intersect(p.mask) || rows[2].intersect(p.mask) {
			matches = append(matches, candidates[i])
		}
	}
	return matches
}

func calculateGearRatio(lineNo int, gear *mask192, candidates [][]*partNo) int {
	var (
		gearParts      []*partNo
		gearCandidates []*partNo
	)
	if lineNo > 0 {
		gearCandidates = append(gearCandidates, candidates[lineNo-1]...)
	}
	gearCandidates = append(gearCandidates, candidates[lineNo]...)
	if lineNo+1 < len(candidates) {
		gearCandidates = append(gearCandidates, candidates[lineNo+1]...)
	}
	for i := range gearCandidates {
		candidate := gearCandidates[i]
		if candidate.mask.intersect(gear) {
			gearParts = append(gearParts, candidate)
		}
		if len(gearParts) > 2 {
			return 0
		}
	}
	if len(gearParts) != 2 {
		return 0
	}
	return gearParts[0].int() * gearParts[1].int()
}

func main() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var (
		partsByLine [][]*partNo
		gearsByLine [][]*mask192
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
			lineGears   []*mask192
			line        = scanner.Text()
		)
		for pos, r := range line {
			if !isDigit(r) {
				if isSymbol(r) {
					currentLine.set(pos)
					if r == '*' {
						gear := &mask192{}
						gear.set(pos)
						lineGears = append(lineGears, gear)
					}
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
		// record all gearsByLine
		gearsByLine = append(gearsByLine, lineGears)
		// move state forward as line is processed, current line is the next line
		// for previous candidates.
		lineState[0], lineState[1], lineState[2] = lineState[1], lineState[2], currentLine
		// collect all partsByLine
		partsByLine = append(partsByLine, collectMatches(lineState, candidates))
		// new partsByLine become the next set of candidates
		candidates = newParts
	}
	// move the state forward again to be able to check previous candidates.
	lineState[0], lineState[1], lineState[2] = lineState[1], lineState[2], currentLine
	// collect all partsByLine from last line as well
	partsByLine = append(partsByLine, collectMatches(lineState, candidates))
	// must shift the parts by one to get rid of first empty row (due to deferred
	// processing)
	partsByLine = partsByLine[1:]
	// Now have parts per line, we can find the gears per line and see if there's
	// partsByLine on previous, current and next line that match. If there's exactly two,
	// keep the gear part sum.
	sum := 0
	for lineNo, lineGears := range gearsByLine {
		for _, gear := range lineGears {
			sum += calculateGearRatio(lineNo, gear, partsByLine)
		}
	}
	log.Printf("Result: %d", sum)
}
