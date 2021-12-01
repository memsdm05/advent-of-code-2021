package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("data/day1.txt")
	scanner := bufio.NewScanner(f)

	var (
		count int
		val int
		lastval int
	)

	for scanner.Scan() {
		lastval = val
		fmt.Sscanf(scanner.Text(), "%d", &val)

		if lastval < val {
			count++
		}
	}

	count--
	fmt.Println(count)
}
