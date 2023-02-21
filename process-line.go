package main

import (
	"fmt"
	"log"
	"net/http"
)

func processLine(client *http.Client, baseURL string, line string) {
	fullURL := fmt.Sprintf("%s/%s", baseURL, line)
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < 400 {
		fmt.Printf("/%s: %d\n", line, resp.StatusCode)
	}
}
