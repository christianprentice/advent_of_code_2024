package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
    file , err := os.Open("../../input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    left := make([]int, 0)
    rightCounts := make(map[int]int)

    scanner := bufio.NewScanner(file)
    for i := 0; scanner.Scan(); i++ {
        line := scanner.Text()
        parts := strings.Fields(line)
        leftNum, err := strconv.Atoi(parts[0])
        if err != nil {
            log.Fatal(err)
        }
        rightNum, err := strconv.Atoi(parts[1])
        if err != nil {
            log.Fatal(err)
        }

        left = append(left, leftNum)
        rightCounts[rightNum]++
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    similarityScore := 0
    for _, leftVal := range left {
        similarityScore += leftVal * rightCounts[leftVal]
    }

    fmt.Println(similarityScore)
}

