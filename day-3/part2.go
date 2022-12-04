package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type rucksack struct {
	compartment1 []string
	compartment2 []string
}

type groups struct {
	rucksacks []string
}

var rucksacks []rucksack

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	idx := 0
	var lines []string
	var group []groups

	for scanner.Scan() {
		line := scanner.Text()

		if idx%3 == 0 {
			lines = []string{}
		}

		lines = append(lines, line)

		if idx%3 == 2 {
			var g groups
			g.rucksacks = lines
			group = append(group, g)
		}

		idx++
	}

	fmt.Println("Groups: ", len(group))

	var totalPriority int

	for _, group := range group {
		item, _ := getItemInAllRucksacks(group.rucksacks)
		prio := getItemPriority(item)
		println("Item: ", item, " Priority: ", prio)
		totalPriority += int(prio)
	}

	fmt.Println("Total priority: ", totalPriority)

}

func getItemPriority(item string) byte {
	ascii := []byte(item)[0]

	if ascii >= 65 && ascii <= 90 {
		return ascii - 64 + 26
	}

	return ascii - 96
}

func getItemInBothCompartments(rucksack rucksack) (string, error) {
	for _, item1 := range rucksack.compartment1 {
		for _, item2 := range rucksack.compartment2 {
			if item1 == item2 {
				return item1, nil
			}
		}
	}

	return "", errors.New("No item in both compartments")
}

func getItemInAllRucksacks(rucksacks []string) (string, error) {
	for _, item1 := range strings.Split(rucksacks[0], "") {
		for _, item2 := range strings.Split(rucksacks[1], "") {
			for _, item3 := range strings.Split(rucksacks[2], "") {
				if item1 == item2 && item2 == item3 {
					return item1, nil
				}
			}
		}
	}

	return "", errors.New("No item in all rucksacks")
}
