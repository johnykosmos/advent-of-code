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

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
    fmt.Println(antinodeMap)
    return len(antinodeMap)
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
