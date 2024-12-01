package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)


func parseInput(input string)([]int, []int){
    lines := strings.Split(input, "\n")
    if lines[len(lines) - 1] == ""{
        lines = lines[:len(lines) - 1] 
    }

    leftList := make([]int, len(lines)) 
    rightList := make([]int, len(lines)) 

    for lineID, line := range lines{
        sides := strings.Split(line, "   ") 
        left, err := strconv.Atoi(sides[0])
		if err != nil {
		    log.Fatal("Error: ", err)	
		}
		
		right, err := strconv.Atoi(sides[1])
		if err != nil {
			log.Fatal("Error: ", err)
		}

        leftList[lineID] = left
        rightList[lineID] = right
    }
    return leftList, rightList 
}

func part1(input string) int {
    leftList, rightList := parseInput(input)
    sort.Ints(leftList)
    sort.Ints(rightList)
    
    var totalDistance int = 0
    for i := 0; i < len(leftList); i++ {
        var distance int = leftList[i] - rightList[i]
        if distance < 0 {
            distance *= -1
        }
        totalDistance += distance
    }

    return totalDistance
}

func countOccurance(targetNumber int, rightList []int) int {
    var occurance int
    for i := 0; i < len(rightList); i++ {
        if targetNumber == rightList[i]{
            occurance++
        }
    }
    return occurance
}

func part2(input string) int {
    leftList, rightList := parseInput(input)

    var similarityScore int = 0
    for i := 0; i < len(leftList); i++ {
        var distance int = leftList[i] * countOccurance(leftList[i], rightList)
        similarityScore += distance
    }

    return similarityScore
}

func getInput() string{
    content, err := os.ReadFile("input.txt")
    if err != nil{
        log.Fatal(err)
    }
    return string(content)
}



func main(){
    var input string = getInput()
    
    fmt.Println("The solution for part 1 is:", part1(input))
    fmt.Println("The solution for part 2 is:", part2(input))
}
