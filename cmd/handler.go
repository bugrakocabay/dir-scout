package cmd

import (
	"fmt"

	"github.com/bugrakocabay/endpoint-brute/cmd/utils"
	"github.com/spf13/cobra"
)

type scanner interface {
	Scanner(config Config) chan string
}

type CMD struct {
	Scan scanner
}

func New(scanner scanner) *CMD {
	return &CMD{Scan: scanner}
}

func Do(cmd *cobra.Command, config Config) {
	utils.ShowBanner(config.Url, config.Wordlist)
	pkg := CMD{}

	ch := pkg.Scan.Scanner(config)
	for v := range ch {
		if config.Success {
			fmt.Println(v)
		}
	}
}
