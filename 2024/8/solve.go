package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Cords struct {
    X int
    Y int
}

func checkForAntinodes(antennaGrid []string, antennaCords Cords, 
    antenna rune, antinodeMap map[Cords]bool) { 
        for rowIndex, row := range antennaGrid {
            for elementIndex, element := range row {
                if element == antenna{
                    xDistance := 2 * (antennaCords.X - elementIndex)
                    yDistance := 2 * (antennaCords.Y - rowIndex)
                    if xDistance == 0 || yDistance == 0 {
                        continue
                    }
                    antinodeCords := Cords{X: antennaCords.X - xDistance, 
                        Y: antennaCords.Y - yDistance}
                    if antinodeCords.X < 0 || antinodeCords.X >= len(row) ||
                    antinodeCords.Y < 0 || antinodeCords.Y >= len(antennaGrid) {
                        continue
                    }
                    antinodeMap[antinodeCords] = true 
                }               
            }
        }
}

func part1(input string) int{
    antennaGrid := strings.Split(strings.TrimSpace(input), "\n")
    antinodeMap := make(map[Cords]bool)
    for rowIndex, row := range antennaGrid {
        for antennaIndex, antenna := range row {
            if antenna != '.' {
                antennaCords := Cords{X: antennaIndex, Y: rowIndex}
                checkForAntinodes(antennaGrid, antennaCords, antenna, antinodeMap)     
            }
        }
    }
    return len(antinodeMap)
}

func checkForAllAntinodes(antennaGrid []string, antennaCords Cords, 
    antenna rune, antinodeMap map[Cords]bool) { 
        for rowIndex, row := range antennaGrid {
            for elementIndex, element := range row {
                if element == antenna{
                    xDistance := antennaCords.X - elementIndex
                    yDistance := antennaCords.Y - rowIndex
                    if xDistance == 0 || yDistance == 0 {
                        continue
                    }
                    currentX, currentY := antennaCords.X, antennaCords.Y
                    for {
                        currentX -= xDistance
                        currentY -= yDistance
                        if currentX < 0 || currentX >= len(row) ||
                        currentY < 0 || currentY >= len(antennaGrid) {
                            break 
                        }
                        antinodeMap[Cords{X: currentX, Y: currentY}] = true 
                    }
                }               
            }
        }
}

func part2(input string) int{
    antennaGrid := strings.Split(strings.TrimSpace(input), "\n")
    antinodeMap := make(map[Cords]bool)
    for rowIndex, row := range antennaGrid {
        for antennaIndex, antenna := range row {
            if antenna != '.' {
                antennaCords := Cords{X: antennaIndex, Y: rowIndex}
                checkForAllAntinodes(antennaGrid, antennaCords, antenna, antinodeMap)     
            }
        }
    }
    return len(antinodeMap)
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
