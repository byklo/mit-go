package main

import "fmt"

func main() {
    count := 0
    j := "aA"
    s := "aAAbbbb"

    table := make(map[rune]bool)

    for _, r := range j {
        table[r] = true
    }

    for _, r := range s {
        if _, exist := table[r]; exist {
            count++
        }
    }

    fmt.Println(count)
}
