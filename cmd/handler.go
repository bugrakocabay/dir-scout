package cmd

import (
	"github.com/bugrakocabay/dir-scout/cmd/config"
	"github.com/bugrakocabay/dir-scout/cmd/utils"
	"github.com/bugrakocabay/dir-scout/pkg/output"
	"github.com/spf13/cobra"
)

func Do(cmd *cobra.Command, config config.Config) {
	utils.ShowBanner(config.Url, config.Wordlist)

	output.Output(config)
}
