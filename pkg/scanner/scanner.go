package scanner

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func Scanner(url, filename string, verbosity int16) {
	var wg sync.WaitGroup
	ch := make(chan int)

	scanner, err := readFile(filename)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)

		go processLine(client, url, line, &wg, ch)
		t := time.Duration(1000 / verbosity)
		time.Sleep(time.Millisecond * t)
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
