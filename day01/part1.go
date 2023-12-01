package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
)


func main() {
	file, _ := os.Open("input")

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

    fmt.Printf("Answer part 1: %d\n", sum(run(lines)))
}

func sum(arr []int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}

func run(lines []string) []int {
	numbers := []int{}

	for _, line := range lines {
		numbers = append(numbers, runOne(line))
	}
	return numbers
}

func runOne(line string) int {
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
