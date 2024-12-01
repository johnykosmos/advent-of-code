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
        distance := leftList[i] - rightList[i]
        if distance < 0 {
            distance *= -1
        }
        totalDistance += distance
    }

    return totalDistance
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
}
