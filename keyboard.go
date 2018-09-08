/*
500. Keyboard Row

Given a List of words, return the words that can be typed using letters of alphabet on
only one row's of American keyboard like the image below.

Example 1:

Input: ["Hello", "Alaska", "Dad", "Peace"]
Output: ["Alaska", "Dad"]

Note:
You may use one character in the keyboard more than once.
You may assume the input string will only contain letters of alphabet.
*/

package main

import "fmt"
import "os"
import "strings"

func sliceContains(s []string, x string) bool {
    for _, r := range s {
        if r == x {
            return true
        }
    }
    return false
}

func isOneRow(s string) bool {
    top := []string {"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
    mid := []string {"a", "s", "d", "f", "g", "h", "j", "k", "l"}
    bot := []string {"z", "x", "c", "v", "b", "n", "m"}

    var row []string

    switch {
        case sliceContains(top, strings.ToLower(string(s[0]))):
            row = top
        case sliceContains(mid, strings.ToLower(string(s[0]))):
            row = mid
        case sliceContains(bot, strings.ToLower(string(s[0]))):
            row = bot
        default:
            fmt.Printf("%s is not an alphabet\n", string(s[0]))
            os.Exit(0)
    }

    for _, c := range s[1:] {
        if ! sliceContains(row, strings.ToLower(string(c))) {
            return false
        }
    }

    return true
}

func main() {
    queries := []string {"Alaska", "johnny", "TemMpSet", "you"}
    hits := make([]string, 0)

    for _, q := range queries {
        if isOneRow(q) {
            hits = append(hits, q)
        }
    }

    fmt.Println(queries)
    fmt.Println(hits)
}
