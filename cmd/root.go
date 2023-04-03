package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "endpoint-discovery [https://url.com] [wordlist.txt]",
	Short: "\nA endpoint discovery tool for recon.",
	Long: `endpoint-discovery is a CLI tool that allows you
to feed a wordlist and brute force them to a given URL.`,
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		Do()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "toggle için yardım mesajı")
}
