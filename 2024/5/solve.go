package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
    "time"
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

func sumMiddles(validUpdates [][]string) int{
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
    return sumMiddles(validUpdates) 
}

func part2(input string) int{
    var invalidUpdates [][]string
    rules, updateOrder := parseInput(input)
    ruleMap := getHierarchy(rules)
    for _, order := range updateOrder {
        splitOrder := strings.Split(order, ",") 
        if !isOrderValid(splitOrder, ruleMap) {
            invalidUpdates = append(invalidUpdates, splitOrder)
        }
    }

    for _, order := range invalidUpdates {
        for !isOrderValid(order, ruleMap){
            for i := 0; i < len(order) - 1; i++ {
                valid := false
                for _, rule := range ruleMap[order[i]] {
                    if order[i+1] == rule {
                        valid = true
                        break
                    }
                }
                if !valid {
                    tmp := order[i]
                    order[i] = order[i+1]
                    order[i+1] = tmp        
                }
            }
        }
    }
    return sumMiddles(invalidUpdates) 
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
