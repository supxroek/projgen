package main

import (
	"os"

	"projgen/cmd"
)

// Entry point for the projgen CLI.
// Wires the Cobra root command and executes the CLI.
func main() {
	if err := cmd.Execute(); err != nil {
		// Cobra already prints user-friendly errors; ensure proper exit code.
		os.Exit(1)
	}
}
 
