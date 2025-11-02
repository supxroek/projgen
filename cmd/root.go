package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd is the base command for projgen.
var rootCmd = &cobra.Command{
	Use:   "projgen",
	Short: "Project generator for frontend, backend, and fullstack apps",
	Long:  "projgen scaffolds modern projects with interactive prompts and templates.",
	// Keep usage quiet on errors; subcommands will provide context.
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global persistent flags can be defined here in later phases.
	// Example: rootCmd.PersistentFlags().BoolP("yes", "y", false, "non-interactive yes to all prompts")
}

