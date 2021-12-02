package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	f, _ := os.Open("data/day2.txt")
	scanner := bufio.NewScanner(f)
	var (
		dir string
		mag int

		depth int
		dist int
		aim int
	)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %d", &dir, &mag)
		switch dir {
		case "forward":
			dist += mag
			depth += aim * mag
		case "down":
			aim += mag
		case "up":
			aim -= mag
		}
	}

	fmt.Println(depth * dist)
}