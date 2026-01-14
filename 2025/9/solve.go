package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Tile struct {
	x, y int
}

func getAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func (t *Tile)countArea(otherT Tile) int {
	return getAbs(t.x - otherT.x + 1) * getAbs(t.y - otherT.y + 1)
}

func parseInput(input string) []Tile {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	tiles := make([]Tile, len(lines))
	for lineId, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		tiles[lineId] = Tile{x: x, y:y}
	}
	return tiles
}

func part1(input string) int{
	tiles := parseInput(input)
	max := 0
	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			area := tiles[i].countArea(tiles[j])
			if max < area {
				max = area
			}
		}
	}

    return max
}

func part2(input string) int{
	tiles := parseInput(input)
	fmt.Println(tiles)

    return 0
}

func getInput() string{
    content, err := os.ReadFile("sample-input.txt")
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
