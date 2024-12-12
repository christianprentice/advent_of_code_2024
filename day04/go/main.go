package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isWithinBounds(r, c, rows, cols int) bool {
    return r >= 0 && r < rows && c >= 0 && c < cols
}

func checkDirection(grid [][]byte, r, c int, dir []int, word string) bool {
    rows := len(grid)
    cols := len(grid[0])
    for k := 1; k < len(word); k++ {
        nr := r + dir[0]*k
        nc := c + dir[1]*k
        if !isWithinBounds(nr, nc, rows, cols) || grid[nr][nc] != word[k] {
            return false
        }
    }
    return true
}

func findXmas(grid [][]byte) int {
    xmasCount := 0
    rows := len(grid)
    cols := len(grid[0])
    directions := [][]int{{0,1}, {1,1}, {1,0}, {1,-1}}

    //We'll only check right, down, down-right and down-left directions
    //This is so we won't have to check the whole 8 directions
    for r := 0; r < rows; r++ {
        for c :=0; c < cols; c++ {

            if grid[r][c] == 'X' {
                for _, dir := range directions {
                    if checkDirection(grid, r, c, dir, "XMAS") {
                        xmasCount++
                    }
                }
            } else if grid[r][c] == 'S' {
                for _, dir := range directions {
                    if checkDirection(grid, r, c, dir, "SAMX") {
                        xmasCount++
                    }
                }
            }

        }
    }
    return xmasCount
}

func main() {
    file , err := os.Open("../input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var grid [][]byte

    for scanner.Scan() {
        grid = append(grid, []byte(scanner.Text()))
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    xmasCount := findXmas(grid)
    fmt.Println(xmasCount)
}
