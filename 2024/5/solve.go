package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func parseInput(input string) ([]string, []string){
    splitInput := strings.Split(input, "\n\n") 
    rules := strings.Split(strings.TrimSpace(splitInput[0]), "\n") 
    order := strings.Split(strings.TrimSpace(splitInput[1]), "\n")

    return rules, order 
}

func getHierarchy(rules []string) map[string][]string {
    ruleMap := make(map[string][]string)
    for _, rule := range rules {
        splitRule := strings.Split(rule, "|")
        ruleMap[splitRule[0]] = append(ruleMap[splitRule[0]], splitRule[1])
    }
    return ruleMap
}

func isOrderValid(splitOrder []string, ruleMap map[string][]string) bool {
    for i := 0; i < len(splitOrder) - 1; i++ {
        valid := false
        for _, rule := range ruleMap[splitOrder[i]] {
            if splitOrder[i+1] == rule {
                valid = true
                break
            }
        }
        if !valid {
            return false
        }
    }
    return true
}
func part1(input string) int{
    var validUpdates [][]string
    rules, updateOrder := parseInput(input)
    ruleMap := getHierarchy(rules)
    for _, order := range updateOrder {
        splitOrder := strings.Split(order, ",") 
        if isOrderValid(splitOrder, ruleMap) {
            validUpdates = append(validUpdates, splitOrder)
        }
    }

    middleSum := 0
    for _, update := range validUpdates {
        middleNum, err := strconv.Atoi(update[len(update)/2])
        if err != nil {
            log.Fatalf("Could not convert %s to int", update[len(update)/2])
        }
        middleSum += middleNum
    }

    return middleSum 
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
    
    fmt.Println("The solution for part 1 is:", part1(input))
    fmt.Println("The solution for part 2 is:", part2(input))
}
