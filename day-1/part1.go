package main

import (
	"bufio"
	"os"
	"strconv"
)

type Elf struct {
	elfId    int
	calories int
}

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var elfies []Elf
	var elfId int = 0

	var elf Elf

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfies = append(elfies, elf)
			elf = Elf{}
			elfId++
		} else {
			elf.elfId = elfId
			calories, _ := strconv.Atoi(line)
			elf.calories += calories
		}
	}

	elfies = append(elfies, elf)

	elf = getElfWithMostCalories(elfies)

	println("Elf with most calories: ", elf.elfId, " with ", elf.calories, " calories")

}

func getElfWithMostCalories(elfies []Elf) Elf {
	var elf Elf
	for _, e := range elfies {
		if e.calories > elf.calories {
			elf = e
		}
	}
	return elf
}
