package cmd

var GlobalFlags struct {
	Verbose    int    //Tells the program to print everything to screen (used multiple times for better verbosity).
	ConfigFile string //Config file path (assumed ./.gobot if not specified)
}

// rootFlags provides flag definitions valid for root command.
var rootFlags struct {
	Version bool
}

// initFlags provdes flag definition for init command.
var initFlags struct {
	ConfigFile string
	Exchange   string
	Strategies []struct {
		Market   string
		Strategy string
	}
	BTCAddress string
}

// startFlags provdes flag definition for start command.
var startFlags struct {
	Simulate bool
}
