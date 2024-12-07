package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Direction struct {
    x int
    y int
}

type Guard struct{
    x int
    y int
    direction Direction 
}

var GUARD_DIRECTIONS = map[string]Direction{
    "^": {0 ,-1}, 
    ">": {1, 0}, 
    "v": {0, 1}, 
    "<": {-1, 0},
}

func (guard *Guard) rotate(){
    newdx, newdy := -guard.direction.y, guard.direction.x 
    guard.direction.x, guard.direction.y = newdx, newdy
}

func (guard *Guard) move(mapGrid []string, xCounter *int) bool {
    if mapGrid[guard.y][guard.x] != 'X' && mapGrid[guard.y][guard.x] != 'O' {
        row := []byte(mapGrid[guard.y])
        row[guard.x] = 'X'
        mapGrid[guard.y] = string(row)
        *xCounter++
    }
    newX, newY := guard.x + guard.direction.x, guard.y + guard.direction.y
    if newX < 0 || newX >= len(mapGrid[guard.y]) ||
    newY < 0 || newY >= len(mapGrid) {
        return false
    }
    if mapGrid[newY][newX] != '.' && mapGrid[newY][newX] != 'X' {
        guard.rotate()
        return true 
    }
    guard.x, guard.y = newX, newY
    return true
}

func findGuardPosition(mapGrid []string) (Guard, error) {
    for rowIndex, row := range mapGrid {
        for direction := range GUARD_DIRECTIONS {
            guardPosition := strings.Index(row, direction)
            if guardPosition != -1 {
                return Guard{x: guardPosition, y: rowIndex, 
                direction: GUARD_DIRECTIONS[direction]}, nil
            }
        }
    }
    return Guard{}, errors.New("Could not find a guard!") 
}

func part1(input string) int{
    xCounter := 0
    mapGrid := strings.Split(strings.TrimSpace(input), "\n")
    guard, err := findGuardPosition(mapGrid)
    if err != nil {
        log.Fatal(err)
    }
    for guard.move(mapGrid, &xCounter){}

    return xCounter 
}


func part2(input string) int {
    xCounter := 0
    oCounter := 0
    mapGrid := strings.Split(strings.TrimSpace(input), "\n")
    obstructionMap := make(map[Direction]bool)

    guard, err := findGuardPosition(mapGrid)
    if err != nil {
        log.Fatal(err)
    }

    obstructionMap[Direction{guard.x, guard.y}] = true
    mapGridCopy := append([]string{}, mapGrid...)
    for guard.move(mapGrid, &xCounter) {
        obstruction := Direction{guard.x, guard.y}
        if obstructionMap[obstruction] {
            continue
        }
        obstructionMap[obstruction] = true

        row := []byte(mapGridCopy[guard.y])
        row[guard.x] = 'O'
        mapGridCopy[guard.y] = string(row)

        guardCopy := guard
        guardCopy.x, guardCopy.y = guard.x-guard.direction.x, guard.y-guard.direction.y
        visited := make(map[Guard]bool)

        for guardCopy.move(mapGridCopy, &xCounter) {
            if visited[guardCopy] {
                oCounter++
                break
            }
            visited[guardCopy] = true
        }

        row[guard.x] = '.'
        mapGridCopy[guard.y] = string(row)

    }

    return oCounter
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
