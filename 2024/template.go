package main

import (
    "os"
    "fmt"
    "log"
    "time"
)


func part1(input string) int{
    return 0
}

func part2(input string) int{
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
