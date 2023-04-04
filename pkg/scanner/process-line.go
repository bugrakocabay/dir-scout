package scanner

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func processLine(client *http.Client, baseURL string, line string, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	if !strings.HasSuffix(baseURL, "http") {
		baseURL = fmt.Sprintf("https://%s", baseURL)
	}

	fullURL := fmt.Sprintf("%s/%s", baseURL, line)
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	successString := color.HiMagentaString("+ ") + color.RedString("/"+line) + " " + color.GreenString(strconv.Itoa(resp.StatusCode))
	if resp.StatusCode < 400 {
		fmt.Println(successString)
	}
	ch <- 1
}
