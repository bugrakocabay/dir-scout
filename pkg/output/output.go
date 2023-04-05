package output

import (
	"fmt"
	"github.com/bugrakocabay/dir-scout/cmd/config"
	"github.com/bugrakocabay/dir-scout/pkg/scanner"
	"github.com/fatih/color"
	"log"
	"os"
	"strconv"
)

func Output(config config.Config) {
	ch := make(chan map[string]int)
	scanner.Scanner(config, ch)

	if config.Output != "" {
		writeToFile(config.Output, ch)
	} else {
		outputToConsole(config.Success, ch)
	}
}

func writeToFile(filename string, ch chan map[string]int) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for m := range ch {
		for k, v := range m {
			fmt.Fprintf(file, "/%s: %d\n", k, v)
		}
	}
}

func outputToConsole(success bool, ch chan map[string]int) {
	for m := range ch {
		for k, v := range m {
			var responseString string
			if success {
				if v < 400 {
					responseString = color.HiMagentaString("+ "+"/"+k) + " " + color.GreenString(strconv.Itoa(v))
					fmt.Println(responseString)
				}
			} else {
				if v < 400 {
					responseString = color.HiMagentaString("+ "+"/"+k) + " " + color.GreenString(strconv.Itoa(v))
				} else {
					responseString = color.HiMagentaString("- "+"/"+k) + " " + color.RedString(strconv.Itoa(v))
				}
				fmt.Println(responseString)
			}
		}
	}
}
