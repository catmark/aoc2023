package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
    fmt.Printf("Answer part 1: %d\n", part1(lines))
}

func part1(lines []string) (sum int) {
	for _, line := range lines {
		sum += run(line)
	}
	return
}

func run(line string) int {
	firstSet := false
	first, last := 0, 0
	for _, c := range line {
		if unicode.IsDigit(c) {
			if !firstSet {
				first = int(c - '0')
			}
			firstSet = true
			last = int(c - '0')
		}
	}
	return first * 10 + last
}
