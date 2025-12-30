package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type DirectionOffset struct {
	x int
	y int
}

var neighborOffsets = []DirectionOffset{
	{x: -1, y: -1}, {x: 0, y: -1}, {x: 1, y: -1}, {x: -1, y: 0},
	{x: 1, y: 0}, {x: -1, y: 1}, {x: 0, y: 1}, {x: 1, y: 1}}

func part1(input string) int{
	availablePapers := 0
	papersGrid := strings.Split(strings.TrimSpace(input), "\n")
	for i := 0; i < len(papersGrid); i++ {
		for j := 0; j < len(papersGrid[i]); j++ {
			if papersGrid[i][j] == '.' {
				continue
			}
			
			neighborCounter := 0
			for _, neighborOff := range neighborOffsets {
				if i + neighborOff.y < 0 || i + neighborOff.y >= len(papersGrid) {
					continue
				}
				
				if j + neighborOff.x < 0 || j + neighborOff.x >= len(papersGrid[i]) {
					continue
				}

				if papersGrid[i+neighborOff.y][j+neighborOff.x] == '@'{
					neighborCounter++
					if neighborCounter > 3 {	
						break
					}
				}
			}

			if neighborCounter < 4 {
				availablePapers++
			}
		}
	}

    return availablePapers
}

func part2(input string) int{
	availablePapers := 0
	papersGrid := strings.Split(strings.TrimSpace(input), "\n")

	byteGrid := make([][]byte, len(papersGrid))
    for lineid, line := range papersGrid {
        byteGrid[lineid] = []byte(line) 
	}
	
	for {
		aboutRemoval := make(map[DirectionOffset]bool)
		for i := 0; i < len(byteGrid); i++ {
			for j := 0; j < len(byteGrid[i]); j++ {
				if byteGrid[i][j] == '.' {
					continue
				}

				neighborCounter := 0
				for _, neighborOff := range neighborOffsets {
					if i + neighborOff.y < 0 || i + neighborOff.y >= len(byteGrid) {
						continue
					}

					if j + neighborOff.x < 0 || j + neighborOff.x >= len(byteGrid[i]) {
						continue
					}

					if byteGrid[i+neighborOff.y][j+neighborOff.x] == '@'{
						neighborCounter++
						if neighborCounter > 3 {	
							break
						}
					}
				}

				if neighborCounter < 4 {
					aboutRemoval[DirectionOffset{x: j, y: i}] = true
				}
			}
		}
		if len(aboutRemoval) == 0 {
			break
		}

		for removePos := range aboutRemoval {
			byteGrid[removePos.y][removePos.x] = '.'
		}

		availablePapers += len(aboutRemoval)
	}

    return availablePapers
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
