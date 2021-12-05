// generated at 2021-12-04 16:23:05.369222+00:00, for day 4 part 2
package main

import (
    "advent-of-code-2021/utils"
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

// dont know why I was having so much trouble with this one

// negative values means selected
type Board [5][5]int

func (b *Board) String() string {
    var sb strings.Builder
    for _, r := range b {
        fmt.Fprintf(&sb, "%v\n", r)
    }
    return sb.String()
}

func (b *Board) HasPogMoment() bool {
    // rows
    for _, r := range b {
        if r[0] < 0 && r[1] < 0 && r[2] < 0 && r[3] < 0 && r[4] < 0 {
            return true
        }
    }

    // cols
    for i := 0; i < 5; i++ {
        if b[0][i] < 0 && b[1][i] < 0 && b[2][i] < 0 && b[3][i] < 0 && b[4][i] < 0 {
            return true
        }
    }

    return false
}

func (b *Board) Sum() int {
    sum := 0
    for r := 0; r < 5; r++ {
        for c := 0; c < 5; c++ {
            if b[r][c] > 0 {
                sum += b[r][c]
            }
        }
    }
    return sum
}

func NewBoard(s *bufio.Scanner, spots *[100][]*int) *Board {
    b := &Board{}
    r := 0
    for r < 5 && s.Scan() {
        for c, val := range strings.Fields(s.Text()) {
            fmt.Sscanf(val, "%d", &b[r][c])
            spot := &spots[b[r][c]]
            *spot = append(*spot, &b[r][c])
        }

        r++
    }
    return b
}

func main() {
    scanner, err := utils.LoadScanner("data/day4.txt")
    if err != nil { panic(err) }

    inputs := make([]int, 0)
    boards := make([]*Board, 0)
    var spots [100][]*int
    wins := 0


    for i := range spots {
        spots[i] = make([]*int, 0)
    }

    // parse input
    for _, val := range strings.Split(utils.Next(scanner), ",") {
        num, err := strconv.Atoi(val)
        if err != nil {
            panic(err)
        }
        inputs = append(inputs, num)
    }

    // parse boards
    for scanner.Scan() {
        boards = append(boards, NewBoard(scanner, &spots))
    }

    winners := make([]bool, len(boards))

    // apply inputs and delete winners
    outer: for _, input := range inputs {
        for _, spot := range spots[input] {
            *spot = -1
        }

        for i, board := range boards {
            if !winners[i] && board.HasPogMoment() {
                wins++
                winners[i] = true

                fmt.Println("board win", i)

                if wins == len(boards) {
                    fmt.Println(board.Sum() * input)
                    break outer
                }

            }
        }
    }
}

