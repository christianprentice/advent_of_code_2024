package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main(){
    //mulCount := 0
    file , err := os.Open("../input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
    allMatches := []string{}

    for i := 0; scanner.Scan(); i++ {
        line := scanner.Text()
        allMatches = append(allMatches, regex.FindAllString(line, -1)...)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // Tada!
    fmt.Println("Matched expressions: ", len(allMatches))
    fmt.Println(allMatches)
}
