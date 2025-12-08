package main

import (
    "os"
    "fmt"
    "log"
    "time"
	"strings"
	"strconv"
	"math"
)

const MAXDIAL int = 100 

func parseInput(input string) []int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rotationTable := make([]int, len(lines))
	for lineid, line := range lines {
		direction := line[0]	
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		
		if direction == 'L' {
			rotation *= -1
		}
		rotationTable[lineid] = rotation
	}
	return rotationTable
}

func part1(input string) int{
	rotationTable := parseInput(input)
	dial := 50
	count := 0
	for _, rotation := range rotationTable {
		rotatedDial := dial + rotation
		dial = ((rotatedDial % MAXDIAL) + MAXDIAL) % MAXDIAL
		if (dial == 0) {
			count++
		}
	}
    return count
}

func part2(input string) int{
	rotationTable := parseInput(input)
	dial := 50
    count := 0
	for _, rotation := range rotationTable {
		oldPos := float64(dial)
		dial += rotation
		newPos := float64(dial)

		var hits float64
		if rotation > 0 {
			hits = math.Floor(newPos/float64(MAXDIAL)) - math.Floor(oldPos/float64(MAXDIAL))
		} else {
			hits = math.Floor((oldPos-1)/float64(MAXDIAL)) - math.Floor((newPos-1)/float64(MAXDIAL))
		}

		count += int(math.Abs(hits))
		dial = ((dial % MAXDIAL) + MAXDIAL) % MAXDIAL
	}
    return count
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
