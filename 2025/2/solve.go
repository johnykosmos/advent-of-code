package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func customPow10(exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= 10
	}
	return result
}

func parseInput(input string) [][]string {
	lines := strings.Split(strings.TrimSpace(input), ",")
	idPairs := make([][]string, len(lines))
	for lineid, line := range lines {
		idPair := strings.Split(line, "-")
		if idPair[0][0] != '0' || idPair[1][0] != '0' {
			idPairs[lineid] = idPair
		}
	}
	return idPairs
}

func part1(input string) int {
	idPairs := parseInput(input)
	pairSum := 0
	for _, pair := range idPairs {
		patternRangeMin := int(math.Round(float64(len(pair[0])) / 2))
		patternRangeMax := int(math.Floor(float64(len(pair[1])) / 2))

		minLimit, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Fatal(err)
		}

		maxLimit, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatal(err)
		}

		for patLen := patternRangeMin; patLen <= patternRangeMax; patLen++ {
			var startNum int
			var lastNum int

			if patLen == patternRangeMin && len(pair[0])%2 == 0 {
				atoiResult, err := strconv.Atoi(pair[0][:patternRangeMin])
				if err != nil {
					log.Fatal(err)
				}
				startNum = atoiResult
			} else {
				startNum = customPow10(patLen - 1)
			}

			lastNum = customPow10(patLen) - 1

			for i := startNum; i <= lastNum; i++ {
				realid := i * customPow10(patLen) + i

				if realid < minLimit {
					continue
				}
				if realid > maxLimit {
					break
				}
				//fmt.Println(realid, maxLimit, startNum)
				pairSum += realid
			}
		}
	}
	return pairSum
}

func part2(input string) int{
	idPairs := parseInput(input)
	pairSum := 0
	for _, pair := range idPairs {
		patternRangeMax := int(math.Floor(float64(len(pair[1])) / 2))
		minLimit, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Fatal(err)
		}

		maxLimit, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatal(err)
		}

		invalidIdMap := make(map[int]bool)
		//fmt.Println(minLimit, maxLimit)
		for patLen := 1; patLen <= patternRangeMax; patLen++ {
			var startNum int
			var lastNum int

			startNum = customPow10(patLen - 1)
			lastNum = customPow10(patLen) - 1
		
			for i := startNum; i <= lastNum; i++ {
				realid := i
				for j := 1; ; j++ {
					realid += i * customPow10(patLen * j)
					if realid < minLimit {
						continue
					}
					if realid > maxLimit {
						break
					}
					invalidIdMap[realid] = true
				}
			}
		}
		for key := range invalidIdMap {
			pairSum += key
		}
	}

	return pairSum
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
