package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func findMaxJoltage(line string, neededBatteries int) int {
	joltageByte := make([]byte, neededBatteries)
	for reserveLen := neededBatteries; reserveLen > 0; reserveLen-- {
		for targetDigit := byte('9'); targetDigit > byte('0'); targetDigit-- {
			targetId := strings.IndexByte(line, targetDigit)
			if targetId == -1 || targetId > len(line) - (reserveLen) {
				continue;
			}
			joltageByte[neededBatteries - reserveLen] = targetDigit
			line = line[targetId + 1:]
			break
		}
	}
	joltageNum, err := strconv.Atoi(string(joltageByte))
	if err != nil {
		log.Fatal(err)
	}

	return joltageNum
}

func part1(input string) int {
	joltSum := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		joltSum += findMaxJoltage(line, 2)
	}
    return joltSum
}

func part2(input string) int{
	joltSum := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		joltSum += findMaxJoltage(line, 12)
	}
    return joltSum
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
