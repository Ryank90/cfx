package cfx

// GlobalFlags stores global flags for the cli.
var GlobalFlags struct {
	OutputToScreen int    // Tells the program to output everything to the terminal.
	ConfigFile     string // Configuration file path.
}
