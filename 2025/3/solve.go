package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func part1(input string) int{
	joltSum := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		for firstDigit := byte('9'); firstDigit > byte('0'); firstDigit-- {
			firstId := strings.IndexByte(line, firstDigit)
			if firstId == -1 || firstId == len(line) - 1 {
				continue;
			}
			
			lineRest := line[firstId + 1:]
			secondDigit := byte('0')
			for i := 0; i < len(lineRest); i++ {	
				if secondDigit < lineRest[i] {
					secondDigit = byte(lineRest[i])
				}
			}
			joltageStr := string([]byte{firstDigit, secondDigit})
			joltageNum, err := strconv.Atoi(joltageStr)
			if err != nil {
				log.Fatal(err)
			}
			joltSum += joltageNum
			break
		}
	}

    return joltSum
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
