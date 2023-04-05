package config

type Config struct {
	Url       string
	Wordlist  string
	Verbosity int16
	Output    string
	Success   bool
	// add more fields here for your other flags
}
