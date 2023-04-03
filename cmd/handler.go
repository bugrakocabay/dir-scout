package cmd

import (
	"github.com/bugrakocabay/endpoint-brute/cmd/utils"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func Do() {
	var wg sync.WaitGroup
	ch := make(chan int)

	args := os.Args[1:]
	if len(args) != 2 {
		panic("Usage: [https://url.com] [wordlist.txt]")
	}

	utils.ShowBanner(args[0], args[1])
	scanner, err := readFile(args[1])
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)

		go processLine(client, args[0], line, &wg, ch)
		time.Sleep(time.Millisecond * 100)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for range ch {
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}
}
