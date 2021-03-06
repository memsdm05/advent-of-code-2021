package utils

import (
	"bufio"
	"os"
)

func LoadScanner(path string) (*bufio.Scanner, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(f), nil
}

func Next(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}