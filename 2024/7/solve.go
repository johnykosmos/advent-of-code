package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Equation struct {
    Total int
    Numbers []int
};


func parseInput(input string) []Equation{
    lines := strings.Split(input, "\n")
    equations := make([]Equation, len(lines))
    for lineIndex, line := range lines {
        splitLine := strings.Split(line, ": ")
        total, err := strconv.Atoi(splitLine[0])
        if err != nil {
            log.Fatal("Could not convert str to int!")
        }
        
        numbers := strings.Split(splitLine[1], " ")
        equationNumbers := make([]int, len(numbers))
        for numIndex, number := range numbers {
            intNum, err := strconv.Atoi(number)
            if err != nil {
                log.Fatal("Could not convert str to int!")
            }
            equationNumbers[numIndex] = intNum
        }

        equations[lineIndex] = Equation{Total: total, Numbers: equationNumbers}
    }
    
    return equations  
}

func isEquationValid(equation Equation, total int, index int, concat bool) bool {
        if index == len(equation.Numbers) - 1 {
        return total == equation.Total
    }
    
    if isEquationValid(equation, total + equation.Numbers[index+1], index+1, concat) ||
        isEquationValid(equation, total * equation.Numbers[index+1], index+1, concat){
        return true
    } else if concat{
        concatTotal, err := strconv.Atoi((strconv.Itoa(total) + strconv.Itoa(equation.Numbers[index+1])))
        if err != nil {
            log.Fatal("Could not convert str to int!")
        }
        
        if isEquationValid(equation, concatTotal, index+1, concat) {
            return true
        }
    }

    return false
}
func part1(input string) int{
    sum := 0
    equations := parseInput(input)
    for _, equation := range equations {
        if isEquationValid(equation, equation.Numbers[0], 0, false) {
            sum += equation.Total;
        }
    }
    return sum 
}


func part2(input string) int{
    sum := 0
    equations := parseInput(input)
    for _, equation := range equations {
        if isEquationValid(equation, equation.Numbers[0], 0, true) {
            sum += equation.Total;
        }
    }
    return sum 
}

func getInput() string{
    content, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    return strings.TrimSpace(string(content))
}


func main(){
    var input string = getInput()
    
    start := time.Now()
    fmt.Printf("The solution for part 1 is: %d\nIn time: %s\n", part1(input), time.Since(start))
    start = time.Now()
    fmt.Printf("The solution for part 2 is: %d\nIn time: %s\n", part2(input), time.Since(start))
}
