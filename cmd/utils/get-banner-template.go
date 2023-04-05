package utils

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	ASCII = `
██████╗ ██╗██████╗       ███████╗ ██████╗ ██████╗ ██╗   ██╗████████╗
██╔══██╗██║██╔══██╗      ██╔════╝██╔════╝██╔═══██╗██║   ██║╚══██╔══╝
██║  ██║██║██████╔╝█████╗███████╗██║     ██║   ██║██║   ██║   ██║   
██║  ██║██║██╔══██╗╚════╝╚════██║██║     ██║   ██║██║   ██║   ██║   
██████╔╝██║██║  ██║      ███████║╚██████╗╚██████╔╝╚██████╔╝   ██║   
╚═════╝ ╚═╝╚═╝  ╚═╝      ╚══════╝ ╚═════╝ ╚═════╝  ╚═════╝    ╚═╝          `
	ruler = "======================================================"
)

func ShowBanner(baseURL, wordlist string) {
	color.Blue(fmt.Sprint(ASCII, "v0.0.1"))
	color.Blue(ruler)
	color.Magenta("[+] Target: 	%s\n", color.GreenString(baseURL))
	color.Magenta("[+] Wordlist: 	%s\n", color.GreenString(wordlist))
	color.Blue(ruler)
}
