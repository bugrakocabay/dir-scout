package utils

import (
	"bufio"
	"os"
)

func GetTotalLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err = scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
