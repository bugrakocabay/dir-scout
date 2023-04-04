package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Config struct {
	url       string
	wordlist  string
	verbosity int16
	// add more fields here for your other flags
}

var rootCmd = &cobra.Command{
	Use:   "endpoint-discovery [https://url.com] [wordlist.txt]",
	Short: "\nA endpoint discovery tool for recon.",
	Long: `endpoint-discovery is a CLI tool that allows you
to feed a wordlist and brute force them to a given URL.`,
	Version: "0.0.1",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		verbosity, _ := cmd.Flags().GetInt16("verbosity")
		config := Config{
			url:       args[0],
			wordlist:  args[1],
			verbosity: verbosity,
		}
		Do(cmd, config)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Int16P("verbosity", "v", 5, "Provide a value between 0-10")
}
