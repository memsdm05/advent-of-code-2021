// generated at 2021-12-03 13:18:17.319924+00:00, for day 3 part 1
package main

import (
    "advent-of-code-2021/utils"
    "fmt"
)

func main() {
    scanner, err := utils.LoadScanner("data/day3.txt")
    if err != nil { panic(err) }

    var (
    	counts [12]int
        gamma uint
        epsilon uint
    )

    for scanner.Scan() {
        for i, c := range scanner.Text() {
            if c == '1' {
                counts[i]++
            } else {
                counts[i]--
            }
        }
    }

    //fmt.Println(counts)

    for _, v := range counts {
        gamma <<= 1
        if v > 0 { gamma |= 1 }
    }

    for _, v := range counts {
        epsilon <<= 1
        if v < 0 { epsilon |= 1 }
    }

    fmt.Println(gamma * epsilon)
}

