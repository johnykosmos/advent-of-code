package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)


func parseInput(input string) [][]int {
    lines := strings.Split(input, "\n")
    numbers := make([][]int, len(lines))

    for lineIndex, line := range lines {
        splitLine := strings.Split(line, " ") 
        parsedLine := make([]int, len(splitLine))
        for numIndex, num := range splitLine {
            parsedNum, err := strconv.Atoi(num) 
            if err != nil {
                log.Fatalf("Error: Cannot convert %s at line %d to int!", num, lineIndex)
            }

            parsedLine[numIndex] = parsedNum
        }
        numbers[lineIndex] = parsedLine
    }
    
    return numbers
}

func isReportSafe(report []int) bool {
    lastDistance := 0 
    for i := 0; i < len(report) - 1; i++ {
        distance := report[i] - report[i + 1]
        if distance * lastDistance < 0 {
            return false
        }

        if distance != 0 {
            lastDistance = distance
        }
        
        if distance < 0{
            distance *= -1
        }

        if distance == 0 || distance > 3 {
            return false
        }
        
    }
    return true 
}

func part1(input string) int{
    parsedReports := parseInput(input)

    var safeReports int = 0
    for _, report := range parsedReports {
        if isReportSafe(report){
            safeReports++
        }
    }

    return safeReports
}

func parseUnsafety(report []int, levelToRemove int) []int{
    var parsedReport []int
    for i := 0; i < len(report); i++{
        if i == levelToRemove {
            continue 
        }
        parsedReport = append(parsedReport, report[i]) 
    }
    return parsedReport
}

func part2(input string) int{
    parsedReports := parseInput(input)

    var safeReports int = 0
    for _, report := range parsedReports {
        isSafe := true
        if !isReportSafe(report){
            isSafe = false
            for i := 0; i < len(report); i++ {
                parsedUnsafeReport := parseUnsafety(report, i)
                if isReportSafe(parsedUnsafeReport){
                    isSafe = true
                    break
                }            
            }
        }
        if isSafe {
            safeReports++
        }
    }

    return safeReports
}

func getInput() string{
    content, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    return strings.TrimSpace(string(content))
}


func main(){
    var input string = getInput()
    fmt.Println("The solution for part 1 is:", part1(input))
    fmt.Println("The solution for part 2 is:", part2(input))
}
