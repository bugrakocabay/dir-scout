package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	Url       string
	Wordlist  string
	Verbosity int16
	Output    string
	Success   bool
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
		output, _ := cmd.Flags().GetString("output")
		success, _ := cmd.Flags().GetBool("success")

		config := &Config{
			Url:       args[0],
			Wordlist:  args[1],
			Verbosity: verbosity,
			Output:    output,
			Success:   success,
		}
		Do(cmd, *config)
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
	rootCmd.Flags().StringP("output", "o", "", "output file name")
	rootCmd.Flags().BoolP("success", "s", false, "filters out status codes 400 and above")
}
