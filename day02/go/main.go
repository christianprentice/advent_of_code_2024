package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Some crazy function I found on https://zenn.dev/nnabeyang/articles/22c0c1c3665646
const intSize = 32 << (^uint(0) >> 63)
func abs(v int) int {
    y := v >> (intSize - 1)
    return (v ^ y) - y
}

func checkSafe(tail int, head int, isAscending bool) bool {
    //casting is needed to use the `abs()` function
    fmt.Println("Checking if safe for: ", tail, head, isAscending)
    if isAscending && (tail > head) {
        fmt.Println("1st condition checked")
        return false
    } else if !isAscending && (tail < head) {
        fmt.Println("2nd condition checked")
        return false
    } else if abs(tail - head) > 3 {
        fmt.Println("3rd condition checked")
        return false
    }else if tail == head {
        fmt.Println("4th condition checked")
        return false
    } else {
        fmt.Println("Safe")
        return true
    }
}

func main(){
    safeCount := 0
    file , err := os.Open("../input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    OUTER:
    for i := 0; scanner.Scan() && i < 1000; i++ {
        line := scanner.Text()
        lineList := strings.Fields(line)
        fmt.Println(lineList)
        intList := make([]int, len(lineList))
        for j, str := range lineList {
            num, err := strconv.Atoi(str)
            if err != nil {
                fmt.Println("Error converting string to integer:", err)
                continue
            }
            intList[j] = num
        }
        for j := 0; j < len(intList)-1; j++ {
            fmt.Println("Checking: lineList", intList, "Values: ", intList[j], intList[j+1])
            // check head and trailing tail if ascending
            if !checkSafe(intList[j], intList[j+1], intList[0] < intList[1]) {
                // nice
                continue OUTER
            }
        }
        // countup if safe
        safeCount++
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // Tada!
    fmt.Println(safeCount)
}

