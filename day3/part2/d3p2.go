// generated at 2021-12-03 13:18:17.319924+00:00, for day 3 part 2
package main

import (
    "advent-of-code-2021/utils"
    "fmt"
    "strconv"
)

// THIS IS SO CRINGE

const SIZE = 12

func at(code, idx uint) bool {
    return (code >> (SIZE - idx - 1)) & 1 == 1
}

func most(codes []uint, pos uint, defbit bool) bool {
    counter := 0

    for _, code := range codes {
        if at(code, pos) {
            counter++
        } else {
            counter--
        }
    }

    if counter == 0 {
        return defbit
    }
    return counter > 0 == defbit
}

func find(c []uint, defbit bool) uint {
    codes := c
    newcodes := make([]uint, 0)

    for i := uint(0); i < SIZE && len(codes) > 1; i++ {
        cbit := most(codes, i, defbit)

        for j := 0; j < len(codes); j++ {
            if at(codes[j], i) == cbit {
                newcodes = append(newcodes, codes[j])
            }
        }
        codes = newcodes
        newcodes = nil
    }

    return codes[0]
}

func main() {
    scanner, err := utils.LoadScanner("data/day3.txt")
    if err != nil { panic(err) }

    codes := make([]uint, 0)

    for scanner.Scan() {
        v, _ := strconv.ParseUint(scanner.Text(), 2, 16)
        codes = append(codes, uint(v))
    }

    fmt.Println(find(codes, true) * find(codes, false))
}

