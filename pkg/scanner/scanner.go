package scanner

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/bugrakocabay/endpoint-brute/cmd"
)

type Impl struct{}

func (s *Impl) Scanner(config cmd.Config) chan string {
	ch := make(chan string)

	go func() {
		var wg sync.WaitGroup

		scanner, err := readFile(config.Wordlist)
		if err != nil {
			panic(err)
		}

		client := &http.Client{}
		for scanner.Scan() {
			line := scanner.Text()
			wg.Add(1)

			go processLine(client, config.Url, line, &wg, ch, config.Success)
			t := time.Duration(1000 / config.Verbosity)
			time.Sleep(time.Millisecond * t)
		}

		wg.Wait()
		close(ch)

		if err = scanner.Err(); err != nil {
			log.Fatalf("error reading file: %v", err)
		}
	}()

	return ch
}

func NewScanner() *Impl {
	return &Impl{}
}
