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

	var elves []Elf
	var elfId int = 0

	var elf Elf

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, elf)
			elf = Elf{}
			elfId++
		} else {
			elf.elfId = elfId
			calories, _ := strconv.Atoi(line)
			elf.calories += calories
		}
	}

	elves = sortelves(elves)

	println("Top 3 elves with most calories: ")

	for i := 0; i < 3; i++ {
		println(elves[i].elfId, " with ", elves[i].calories, " calories")
	}

	println("Total calories of top 3: ",
		elves[0].calories+elves[1].calories+elves[2].calories)

}

func sortelves(elves []Elf) []Elf {
	for i := 0; i < len(elves); i++ {
		for j := i + 1; j < len(elves); j++ {
			if elves[i].calories < elves[j].calories {
				elves[i], elves[j] = elves[j], elves[i]
			}
		}
	}
	return elves
}
