package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Problem struct {
	numbers []int
	operation string
}

func parseInput(input string) []Problem {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var problems []Problem
	for lineId, line := range lines[:len(lines) - 1] {
		row := strings.Fields(line)
		if lineId == 0 {
			problems = make([]Problem, len(row))
			for i := range problems {
				problems[i].numbers = make([]int, len(lines) - 1)
			}
		}
		for probId, elem := range row {
			num, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatal(err)
			}
			problems[probId].numbers[lineId] = num	
		}
	}

	for opId, opSign := range strings.Fields(lines[len(lines) - 1]) {
		problems[opId].operation = opSign
	}

	return problems
}

func part1(input string) int{
	totalSum := 0
	problems := parseInput(input)
	
	for _, problem := range problems {
		total := problem.numbers[0]
		for _, num := range problem.numbers[1:] {
			switch (problem.operation) {
			case "*":
				total *= num	
			case "+":
				total += num
			}
		}
		totalSum += total
	}

    return totalSum
}

func part2(input string) int{
    return 0
}

func getInput() string{
    content, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}


func main(){
    var input string = getInput()
    
    start := time.Now()
    fmt.Printf("The solution for part 1 is: %d\nIn time: %s\n", part1(input), time.Since(start))
    start = time.Now()
    fmt.Printf("The solution for part 2 is: %d\nIn time: %s\n", part2(input), time.Since(start))
}
