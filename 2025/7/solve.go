package main

import (
    "os"
    "fmt"
    "log"
    "time"
	"strings"
)

type Coords struct {
	x int
	y int
}

var visited map[Coords]bool 

func countSplits(grid []string, split Coords) int {
	beamSplits := 0
	if split.x < 0 || split.x == len(grid[split.y]){
		return beamSplits
	}
	
	for split.y < len(grid) {
		if grid[split.y][split.x] != '.' {
			if visited[split] == true {
				return beamSplits
			}
			visited[split] = true
			beamSplits = countSplits(grid, Coords{split.x - 1, split.y}) + countSplits(grid, Coords{split.x + 1, split.y}) + 1
			break
		}
		split.y++
	}

	return beamSplits
}

func part1(input string) int {
	grid := strings.Split(strings.TrimSpace(input), "\n")
	beamStart := strings.Index(grid[0], "S")
	visited = make(map[Coords]bool)
	beamSplits := countSplits(grid[1:], Coords{beamStart, 0})
    return beamSplits
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
