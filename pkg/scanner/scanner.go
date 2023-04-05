package scanner

import (
	"log"
	"sync"
	"time"

	"github.com/bugrakocabay/dir-scout/cmd/config"
)

func Scanner(config config.Config, ch chan map[string]int) chan map[string]int {
	go func() {
		var wg sync.WaitGroup

		scanner, err := readFile(config.Wordlist)
		if err != nil {
			panic(err)
		}

		for scanner.Scan() {
			line := scanner.Text()
			wg.Add(1)

			go processLine(config.Url, line, &wg, ch)
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
