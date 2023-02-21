package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Usage: [https://url.com] [wordlist.txt]")
	}

	reader, err := readFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if line != "" {
					// process the final line even if it does not end with '\n'
					processLine(client, args[0], strings.TrimSpace(line))
				}
				break
			}
			log.Fatal("error reading the file:", err)
		}
		// process the line normally if it ends with '\n'
		processLine(client, args[0], strings.TrimSpace(line))
	}
}
