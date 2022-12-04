package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var shapeBonus = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var score = map[string]int{
	"win":  6,
	"draw": 3,
	"loss": 0,
}

type players struct {
	elves string
	me    string
}

type round struct {
	winner     string
	elvesScore int
	meScore    int
}

var rounds []round

// Then, a winner for that round is selected:
// Rock defeats Scissors
// Scissors defeats Paper
// and Paper defeats Rock.
// If both players choose the same shape, the round instead ends in a draw.

var scoreMap = map[players]round{
	{"rock", "rock"}:         {winner: "draw", elvesScore: score["draw"] + shapeBonus["rock"], meScore: score["draw"] + shapeBonus["rock"]},
	{"rock", "paper"}:        {winner: "me", elvesScore: score["loss"] + shapeBonus["rock"], meScore: score["win"] + shapeBonus["paper"]},
	{"rock", "scissors"}:     {winner: "elves", elvesScore: score["win"] + shapeBonus["rock"], meScore: score["loss"] + shapeBonus["scissors"]},
	{"paper", "rock"}:        {winner: "elves", elvesScore: score["win"] + shapeBonus["paper"], meScore: score["loss"] + shapeBonus["rock"]},
	{"paper", "paper"}:       {winner: "draw", elvesScore: score["draw"] + shapeBonus["paper"], meScore: score["draw"] + shapeBonus["paper"]},
	{"paper", "scissors"}:    {winner: "me", elvesScore: score["loss"] + shapeBonus["scissors"], meScore: score["win"] + shapeBonus["scissors"]},
	{"scissors", "rock"}:     {winner: "me", elvesScore: score["loss"] + shapeBonus["scissors"], meScore: score["win"] + shapeBonus["rock"]},
	{"scissors", "paper"}:    {winner: "elves", elvesScore: score["win"] + shapeBonus["scissors"], meScore: score["loss"] + shapeBonus["paper"]},
	{"scissors", "scissors"}: {winner: "draw", elvesScore: score["draw"] + shapeBonus["scissors"], meScore: score["draw"] + shapeBonus["scissors"]},
}

func mapShapes(shape string) (string, error) {
	if shape == "A" || shape == "X" {
		return "rock", nil
	}
	if shape == "B" || shape == "Y" {
		return "paper", nil
	}
	if shape == "C" || shape == "Z" {
		return "scissors", nil
	}

	return "Error", errors.New("Invalid shape")

}

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var elves string
	var me string

	for scanner.Scan() {
		line := scanner.Text()
		options := strings.Split(line, " ")

		elves, _ = mapShapes(string(options[0]))
		me, _ = mapShapes(string(options[1]))

		result := scoreMap[players{elves, me}]

		rounds = append(rounds, result)

	}

	var totalScore int

	for _, round := range rounds {
		totalScore += round.meScore
	}

	println("My score: ", totalScore)

	for _, round := range rounds {
		totalScore += round.elvesScore
	}

	println("Elves score: ", totalScore)
}
