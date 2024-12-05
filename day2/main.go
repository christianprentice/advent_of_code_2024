package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Some crazy function I found on https://zenn.dev/nnabeyang/articles/22c0c1c3665646
const intSize = 32 << (^uint(0) >> 63)
func abs(v int) int {
    y := v >> (intSize - 1)
    return (v ^ y) - y
}

func checkSafe(tail byte, head byte, isAscending bool) bool {
    //casting is needed to use the `abs()` function
    tailValue := int(tail)
    headValue := int(head)
    if isAscending && tailValue < headValue {
        return false
    } else if abs(tailValue - headValue) > 3 {
        return false
    }else if tailValue == headValue {
        return false
    } else {
        return true
    }
}

func main(){
    safeCount := 0
    file , err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for i := 0; scanner.Scan() && i < 1000; i++ {
        line := scanner.Text()
        // check head and trailing tail if ascending
        isAscending := line[0] > line[2]
        for j := 2; scanner.Scan() && j < 5; j += 2 {
            if !checkSafe(line[j], line[j+2], isAscending) {
                break
            }
        }
        safeCount++
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // Tada!
    fmt.Println(safeCount)
}

