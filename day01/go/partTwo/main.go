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

    left := make(map[int]int)
    right := make(map[int]int)

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

        left[leftNum] = left[leftNum] + 1
        right[rightNum] = right[rightNum] + 1
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    similarityScore := 0
    for leftVal, count := range left {
        similarityScore += leftVal * count
    }

    fmt.Println(similarityScore)
}

