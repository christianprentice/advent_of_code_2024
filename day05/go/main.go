package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkOrder(rules map[int][]int, updates []int) bool {
    for i := range len(updates) {
        for j := i + 1; j < len(updates); j++ {
            if requiredAfter, ok := rules[updates[i]]; ok {
                for _, after := range requiredAfter {
                    if after == updates[j] {
                        return false
                    }
                }
            }
        }
    }
    return true
}

func main() {
    file, err := os.ReadFile("../input")
    if err != nil {
        log.Fatal(err)
    }

    fileString := string(file)
    splitFile := strings.Split(fileString, "\n\n")
    rulesInput := strings.Split(splitFile[0], "\n")
    updatesInput := strings.Split(splitFile[1], "\n")

    rules := make(map[int][]int)
    for _, rule := range rulesInput {
        pages := strings.Split(rule, "|")
        page1, _ := strconv.Atoi(pages[0])
        page2, _ := strconv.Atoi(pages[1])
        rules[page2] = append(rules[page2], page1)
    }

    sumOfInorderMiddle := 0
    for _, updateStr := range updatesInput {
        pagesStr := strings.Split(updateStr, ",")
        var updates []int

        for _, pageStr := range pagesStr {
            page, _ := strconv.Atoi(pageStr)
            updates = append(updates, page)
        }

        if checkOrder(rules, updates) {
            middleIndex := len(updates) / 2
            sumOfInorderMiddle += updates[middleIndex]
        }
    }
    //Tada
    fmt.Println(sumOfInorderMiddle)
}
