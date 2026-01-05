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

func parseInputP1(input string) []Problem {
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
	problems := parseInputP1(input)
	
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

func parseInputP2(input string) []Problem {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines) - 1] // I had to do it because I couldn't figure out
	var columnWidth []int		   // how to get rid of the last line which was blank.
	var problems []Problem

	// The first thing I needed to figure out was how many numbers
	// are in the problem. Program iterates through the last line
	// with operation signs and counts every slot before the next sign.
	count := 0
	operationLine := lines[len(lines) - 1]
	for id := range  operationLine{
		if 	id + 1 == len(operationLine) {
			columnWidth = append(columnWidth, count + 1)
			problems = append(problems, Problem{numbers: make([]int, count + 1),
				operation: string(operationLine[id - count])})
		} else if operationLine[id + 1] != ' ' {
			columnWidth = append(columnWidth, count)
			problems = append(problems, Problem{numbers: make([]int, count),
				operation: string(operationLine[id - count])})
			count = 0
		} else {
			count++
		}
	}

	// Given all that I could just iterate with moving pivot through
	// every problem (as I know now how long is it) and get the vertical numbers 
	// with nested loops.
	currId := 0
	for probId, colLen := range columnWidth {
		for i := range colLen {
			num := ""
			for j := 0; j < len(lines) - 1; j++ {
				if lines[j][currId] != ' ' {
					num += string(lines[j][currId])
				} else if num != "" {
					break
				}
			}
			if num != "" {
				intNum, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}
				problems[probId].numbers[i] = intNum
			}
			currId++
		}
		currId++
	}
	return problems
}

func part2(input string) int{
	totalSum := 0
	problems := parseInputP2(input)

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
