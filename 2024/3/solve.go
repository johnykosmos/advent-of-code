package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func getMulList(input string) []string {
    var mulList []string
    pattern := "mul("
    startIndex := 0

    for {
        patternIndex := strings.Index(input[startIndex:], pattern)
        if patternIndex == -1 {
            break
        }
        patternIndex += startIndex
        startIndex = patternIndex + len(pattern)
        
        digitCounter := 0
        for i := startIndex; i < len(input); i++ {
            if digitCounter > 3{
                break
            }

            if input[i] == ',' {
                digitCounter = 0
            }else if input[i] >= '0' && input[i] <= '9' {
                digitCounter++
            }else if input[i] == ')' {
                mulList = append(mulList, input[startIndex:i])        
                break
            }else{
                break
            }
        }
        
    }
    return mulList 
}


func part1(input string) int{ 
    mulList := getMulList(input)
    mulSum := 0
    for _, mul := range mulList {
        splitMul := strings.Split(mul, ",")
        num1, err := strconv.Atoi(splitMul[0]) 
        if err != nil {
            log.Fatalf("Could not convert %s to integer!", splitMul[0])
        }
        num2, err := strconv.Atoi(splitMul[1])
        if err != nil {
            log.Fatalf("Could not convert %s to integer!", splitMul[1])
        }

        mulSum += num1 * num2 
    }

    return mulSum 
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
