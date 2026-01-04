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

type Range struct {
	start int
	end int
}

func parseInput(input string) ([]Range, []int) {
	splitInput := strings.Split(input, "\n\n")
	strRanges := strings.Split(splitInput[0], "\n")
	strNums := strings.Split(strings.Trim(splitInput[1], "\n"), "\n")
	ranges := make([]Range, len(strRanges))
	for rangeId, rangeLine := range strRanges {
		splitRange := strings.Split(rangeLine, "-")
		startNum, err := strconv.Atoi(splitRange[0])
		if err != nil {
			log.Fatal(err)
		}
		endNum, err := strconv.Atoi(splitRange[1])
		if err != nil {
			log.Fatal(err)
		}
		ranges[rangeId] = Range{start: startNum, end: endNum}
	}

	numbers := make([]int, len(strNums))
	for numId, numLine := range strNums {
		num, err := strconv.Atoi(numLine)
		if err != nil {
			log.Fatal(err)
		}
		numbers[numId] = num
	}
	return ranges, numbers
}


func part1(input string) int{
	freshCount := 0
	ranges, numbers := parseInput(input)

	for _, number := range numbers {
		for _, freshRange := range ranges {
			if freshRange.start <= number && number <= freshRange.end {
				freshCount++
				break
			}
		}
	}
    return freshCount
}

func part2(input string) int{
	freshIdCount := 0
	ranges, _ := parseInput(input)
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start == ranges[j].start {
			return ranges[i].end < ranges[j].end
		}
		return ranges[i].start < ranges[j].start
	})

	lastRange := ranges[0]
	freshIdCount += lastRange.end - lastRange.start + 1
	for i := 1; i < len(ranges); i++ {
		if lastRange.end >= ranges[i].start {
			maxEnd := max(lastRange.end, ranges[i].end)
			freshIdCount += maxEnd - lastRange.end
			lastRange.end = maxEnd
		} else {
			lastRange = ranges[i]
			freshIdCount += lastRange.end - lastRange.start + 1
		}
	}
    return freshIdCount
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
