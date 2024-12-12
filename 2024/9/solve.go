package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func parseInput(input string) []int {
    var blocks []int
    fileIndex := 0
    for i := 0; i < len(input); i++ {
        if i != 0 && i % 2 == 0 {
            fileIndex = i / 2
        }else if i % 2 == 1 {
            fileIndex = -1
        }

        count := int(input[i] - '0')
        for j := 0; j < count; j++ {
            blocks = append(blocks, fileIndex)
        }
    }
    return blocks
}


func compactFiles(spaceLayout []int) {
    lastElement := len(spaceLayout) - 1
    for i := 0; i != lastElement; i++ {
        if spaceLayout[i] == -1 {
            for spaceLayout[lastElement] == -1 {
                lastElement-- 
            }
            if lastElement > i && spaceLayout[lastElement] != -1 {
                tmp := spaceLayout[i]
                spaceLayout[i] = spaceLayout[lastElement]
                spaceLayout[lastElement] = tmp
                lastElement-- 
            } else {
                break
            }
        }
    }
}

func part1(input string) int{
    filesystemSum := 0
    spaceLayout := parseInput(input)
    compactFiles(spaceLayout)
    for position, fileIndex := range spaceLayout {
        if fileIndex == -1 {
            break
        }
        filesystemSum += position * fileIndex 
    }

    return filesystemSum 
}

func part2(input string) int{
    return 0
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
