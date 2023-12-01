package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"unicode"
)


func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

    fmt.Printf("Answer part 2: %d\n", part2(lines))
}


func part2(lines []string) (sum int) {
	for _, line := range lines {
		sum += run(line)
	}
	return
}

func toNumber(substring string) (int, bool) {
	switch true {
	case unicode.IsDigit(rune(substring[0])): 
		return int(substring[0] - '0'), true
	case strings.HasPrefix(substring, "zero"): return 0, true
	case strings.HasPrefix(substring, "one"): return 1, true
	case strings.HasPrefix(substring, "two"): return 2, true
	case strings.HasPrefix(substring, "three"):	return 3, true
	case strings.HasPrefix(substring, "four"): return 4, true
	case strings.HasPrefix(substring, "five"): return 5, true
	case strings.HasPrefix(substring, "six"): return 6, true
	case strings.HasPrefix(substring, "seven"): return 7, true
	case strings.HasPrefix(substring, "eight"): return 8, true
	case strings.HasPrefix(substring, "nine"): return 9, true 
	default:return 0, false
	}
}

func run(line string) int {
	firstSet := false
	first, last := 0, 0
	for i := range line {
		if n, ok := toNumber(line[i:]); ok {
			if !firstSet {
				first = n
				firstSet = true
			}
			last = n
		}
	}
	return first * 10 + last
}
