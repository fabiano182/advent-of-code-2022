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

var rucksacks []rucksack

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		var rucksack rucksack

		rucksack.compartment1 = strings.Split(line[0:len(line)/2], "")
		rucksack.compartment2 = strings.Split(line[len(line)/2:], "")

		rucksacks = append(rucksacks, rucksack)
	}

	fmt.Println("Rucksacks: ", len(rucksacks))

	var totalPriority int

	for _, rucksack := range rucksacks {
		item, _ := getItemInBothCompartments(rucksack)
		prio := getItemPriority(item)
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
