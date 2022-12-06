package main

import (
	"bufio"
	"fmt"
	"os"
)

type stacks struct {
	cranes []string
}

func (s *stacks) push(crane string) {
	s.cranes = append(s.cranes, crane)
}

func (s *stacks) pop() string {
	crane := s.cranes[len(s.cranes)-1]
	s.cranes = s.cranes[:len(s.cranes)-1]
	return crane
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	defer file.Close()

	var stack = stacks{cranes: []string{}}

	for scanner.Scan() {
		stack.push(scanner.Text())
	}

	var cranes = [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}} // 9 stacks

	for _, crane := range stack.cranes {
		for idx, crane2 := range crane {
			for i := 0; i < 9; i++ {
				if idx == 1+i*4 {
					if string(crane2) != " " {
						cranes[i] = append(cranes[i], string(crane2))
					}
				}
			}
		}
	}

	fmt.Println(cranes)
}
