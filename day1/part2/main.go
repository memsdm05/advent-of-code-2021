package main

import (
	"bufio"
	"fmt"
	"os"
)

type Window [3]int

func (w *Window) Push(val string) {
	w[0] = w[1]
	w[1] = w[2]
	fmt.Sscanf(val, "%d", &w[2])
}

func (w *Window) Sum() int {
	return w[0] + w[1] + w[2]
}

func main()  {
	f, _ := os.Open("data/day1.txt")
	scanner := bufio.NewScanner(f)

	var (
		count int
		window Window
		lastsum int
	)

	for i := 0; i < 3; i++ {
		scanner.Scan()
		window.Push(scanner.Text())
	}

	for scanner.Scan() {
		lastsum = window.Sum()
		window.Push(scanner.Text())

		if lastsum < window.Sum() {
			count++
		}
	}

	fmt.Println(count)
}
