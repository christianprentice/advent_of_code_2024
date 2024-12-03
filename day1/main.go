package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main(){
    file , err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    left := make([]int, 1000)
    right := make([]int, 1000)

    scanner := bufio.NewScanner(file)

    for i := 0; scanner.Scan() && i < 1000; i++ {
        line := scanner.Text()

        leftNum, err := strconv.Atoi(line[:5])
        if err != nil {
            log.Fatal(err)
        }

        rightNum, err := strconv.Atoi(line[8:])
        if err != nil {
            log.Fatal(err)
        }

        left[i] = leftNum
        right[i] = rightNum
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    sort.Ints(left)
    sort.Ints(right)

    fmt.Println("Left:")
    for _, num := range left {
        fmt.Println(num)
    }

    fmt.Println("Right:")
    for _, num := range right {
        fmt.Println(num)
    }
}
