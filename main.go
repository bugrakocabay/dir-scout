package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Usage: [https://url.com] [wordlist.txt]")
	}

	// read .txt file
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal("error opening the file:", err)
	}
	defer file.Close()

	// create a buffered reader with a 64KB buffer
	reader := bufio.NewReaderSize(file, 64*1024)
	client := &http.Client{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if line != "" {
					// process the final line even if it does not end with '\n'
					fullUrl := fmt.Sprintf("%s/%s", args[0], line)
					req, err := http.NewRequest("GET", strings.TrimSpace(fullUrl), nil)
					if err != nil {
						panic(err)
					}
					resp, err := client.Do(req)
					if err != nil {
						panic(err)
					}
					if resp.StatusCode >= 400 {
						continue
					}
					fmt.Printf("/%s: %d\n", strings.TrimSpace(line), resp.StatusCode)
				}
				break
			}
			log.Fatal("error reading the file:", err)
		}
		// process the line normally if it ends with '\n'
		fullUrl := fmt.Sprintf("%s/%s", args[0], line)
		req, err := http.NewRequest("GET", strings.TrimSpace(fullUrl), nil)
		if err != nil {
			panic(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode >= 400 {
			continue
		}
		fmt.Printf("/%s: %d\n", strings.TrimSpace(line), resp.StatusCode)
	}
}
