package cmd

import (
	"github.com/bugrakocabay/endpoint-brute/cmd/utils"
	"github.com/bugrakocabay/endpoint-brute/pkg/scanner"
	"github.com/spf13/cobra"
)

func Do(cmd *cobra.Command, config Config) {
	utils.ShowBanner(config.url, config.wordlist)
	scanner.Scanner(config.url, config.wordlist, config.verbosity)
}
