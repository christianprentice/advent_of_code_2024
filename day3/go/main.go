package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
    "strconv"
)

func main(){
    mulCount := 0
    file , err := os.Open("../input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    mulFuncReg := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
    allMatches := []string{}

    for scanner.Scan() {
        line := scanner.Text()
        allMatches = append(allMatches, mulFuncReg.FindAllString(line, -1)...)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for _, expr := range allMatches {
        numReg := regexp.MustCompile(`\d+`)
        numbers := numReg.FindAllString(expr, -1)

        x, err := strconv.Atoi(numbers[0])
        if err != nil {
            log.Fatal(err)
        }
        y, err := strconv.Atoi(numbers[1])
        if err != nil {
            log.Fatal(err)
        }

        mulCount += x * y
    }

    // Tada!
    fmt.Println(mulCount)
}
