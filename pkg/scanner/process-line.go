package scanner

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/fatih/color"
)

func processLine(client *http.Client, baseURL string, line string, wg *sync.WaitGroup, ch chan string, success bool) {
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

	var responseString string
	if success {
		if resp.StatusCode < 400 {
			responseString = color.HiMagentaString("+ "+"/"+line) + " " + color.GreenString(strconv.Itoa(resp.StatusCode))
		}
	} else {
		if resp.StatusCode < 400 {
			responseString = color.HiMagentaString("+ "+"/"+line) + " " + color.GreenString(strconv.Itoa(resp.StatusCode))
		} else {
			responseString = color.HiMagentaString("- "+"/"+line) + " " + color.RedString(strconv.Itoa(resp.StatusCode))
		}
	}

	ch <- responseString
}
