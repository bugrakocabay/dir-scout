package scanner

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func processLine(baseURL string, line string, wg *sync.WaitGroup, ch chan map[string]int) {
	defer wg.Done()

	resp, _ := sendRequest(baseURL, line)

	var responseMap = map[string]int{}
	responseMap[line] = resp.StatusCode

	ch <- responseMap
}

func sendRequest(baseURL string, line string) (*http.Response, error) {
	if !strings.HasSuffix(baseURL, "http") {
		baseURL = fmt.Sprintf("https://%s", baseURL)
	}

	fullURL := fmt.Sprintf("%s/%s", baseURL, line)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
