package main

import (
	"bufio"
	"fmt"
	"os"
)

func isContained(a, b, c, d int) bool {
	if a >= c && b <= d {
		return true
	}

	return false
}

func overlaps(a, b, c, d int) bool {
	if !(a > d || b < c) {
		return true
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fullyContained, rangeOverlaps := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elf1Start, elf1End, elf2Start, elf2End := 0, 0, 0, 0
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &elf1Start, &elf1End, &elf2Start, &elf2End)
		if err != nil {
			panic(err)
		}

		if isContained(elf1Start, elf1End, elf2Start, elf2End) || isContained(elf2Start, elf2End, elf1Start, elf1End) {
			fullyContained++
		}

		if overlaps(elf1Start, elf1End, elf2Start, elf2End) {
			rangeOverlaps++
		}

	}
	fmt.Println("Fully Contains:", fullyContained)
	fmt.Println("Range Overlaps:", rangeOverlaps)

}
