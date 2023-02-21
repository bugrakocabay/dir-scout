package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string) (*bufio.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot open the file:%v", err)
	}

	// create a buffered reader with a 64KB buffer
	return bufio.NewReaderSize(file, 64*1024), nil
}
