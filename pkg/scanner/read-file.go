package scanner

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot open the file:%v", err)
	}

	scanner := bufio.NewScanner(file)

	return scanner, nil
}
