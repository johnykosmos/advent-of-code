package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var DIRECTIONS = [8][2]int{
    {0, 1},
    {1, 0},
    {0, -1},
    {-1, 0},
    {1, 1},
    {1, -1},
    {-1, -1},
    {-1, 1}, 
}

func lookForPattern(dx int, dy int, xOffset int, yOffset int,
                    lines []string, pattern string) int {
    for i := 0; i < len(pattern); i++{
        x := xOffset + i*dx
		y := yOffset + i*dy

		if y < 0 || y >= len(lines) || x < 0 || x >= len(lines[y]) {
			return 0
		}

		if lines[y][x] != pattern[i] {
			return 0
		}    
    }
    return 1
}

func part1(input string) int{
    pattern := "XMAS"
    lines := strings.Split(input, "\n")
    xmasCount := 0
    for lineIndex, line := range lines {
        for i := 0; i < len(line); i++ {
            for _, direction := range DIRECTIONS {
                xmasCount += lookForPattern(direction[0], direction[1], i, lineIndex, lines, pattern)
            }
        }
    }
    return xmasCount
}

func part2(input string) int{
    pattern := "MAS"
    lines := strings.Split(input, "\n")
    xmasCount := 0
    for lineIndex, line := range lines {
        for i := 0; i < len(line); i++ {
            masCount := 0
            masCount += lookForPattern(1, 1, i, lineIndex, lines, pattern)
            masCount += lookForPattern(-1, -1, i + 2, lineIndex + 2, lines, pattern)
            if masCount < 1 {
                continue 
            }

            masCount += lookForPattern(-1, 1, i + 2, lineIndex, lines, pattern)
            masCount += lookForPattern(1, -1, i, lineIndex + 2, lines, pattern)
            if masCount >= 2 {
                xmasCount++
            }
        }
    }
    return xmasCount
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
    
    fmt.Println("The solution for part 1 is:", part1(input))
    fmt.Println("The solution for part 2 is:", part2(input))
}
