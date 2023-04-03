package cmd

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
