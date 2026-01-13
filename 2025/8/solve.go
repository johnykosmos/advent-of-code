package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const n int = 10

type Coords struct {
	x, y, z int
}

type Connection struct {
	box1, box2 int
	distance float64
}

func (cords *Coords)calculateDistance(otherCords Coords) float64 {
	return math.Sqrt(math.Pow(float64(cords.x - otherCords.x), 2) + 
	math.Pow(float64(cords.y - otherCords.y), 2) + 
	math.Pow(float64(cords.z - otherCords.z), 2))
}

func parseInput(input string) []Coords {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	coordinates := make([]Coords, len(lines))
	for lineId, line := range lines {
		strCoords := strings.Split(line, ",")
		intCoords := make([]int, len(strCoords))
		for coordId, coord := range strCoords {
			intCoord, err := strconv.Atoi(coord)
			if err != nil {
				log.Fatal(err)
			}		
			intCoords[coordId] = intCoord
		}
		coordinates[lineId].x = intCoords[0]
		coordinates[lineId].y = intCoords[1]
		coordinates[lineId].z = intCoords[2]
	}
	return coordinates
}

func part1(input string) int{
	boxesCords := parseInput(input)	

	// First we have to get all possible connections
	var allConnections []Connection 
	for i := 0; i < len(boxesCords); i++ {
		for j:= i + 1; j < len(boxesCords); j++ {
			allConnections = append(allConnections, 
			Connection{box1: i, box2: j, distance: boxesCords[i].calculateDistance(boxesCords[j])})	
		}
	}

	// Then we sort it and then include only first n elements 
	// where n is described in the problem
	sort.Slice(allConnections, func(i int, j int) bool {
		return allConnections[i].distance < allConnections[j].distance
	})

	// Each element's ids will correspond directely to each box's ids;
	// and element stored would correspond to where it is connected 
	boxesConnections := make([]int, len(boxesCords))
	for i := range len(boxesConnections) {
		boxesConnections[i] = i
	}

	// Through n iterations we have to merge all the connections to
	// get the biggest groups of elements that are connected to each other.
	// It will be done by checking for the roots of each box.
	for i := range n {
		connection := allConnections[i]	
		box1Root := connection.box1
		box2Root := connection.box2

		// Check for the first box's root
		for box1Root != boxesConnections[box1Root] {
			box1Root = boxesConnections[box1Root]
		}

		// Check for the second box's root
		for box2Root != boxesConnections[box2Root] {
			box2Root = boxesConnections[box2Root]
		}

		// If these roots aren't equal we have to assign other box's
		// root to the first box which will eventually lead to merging of 
		// the groups, but for now it will be just an array of pointers to 
		// different boxes
		if box1Root != box2Root {
			boxesConnections[box1Root] = box2Root
		}
	}
	
	
	// Now we have to count all the boxes connected in every group by
	// mapping it to it's root
	circuitMap := make(map[int]int)
	for _, connection := range boxesConnections {
		for connection != boxesConnections[connection] {
			connection = boxesConnections[connection]
		}
		circuitMap[connection]++
	}

	circuits := make([]int, len(circuitMap))
	i := 0
	for id := range circuitMap {
		circuits[i] = circuitMap[id]
		i++
	}
	sort.Slice(circuits, func(i int, j int) bool {
		return circuits[i] > circuits[j]
	})	

	return circuits[0] * circuits[1] * circuits[2]
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
