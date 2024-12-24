package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "porcupine --help",
	Short: "A key-value store cli app.",
	Long: `
\\\\\\\\\\\\\\\
\\ PORCUPINE \\
\\\\\\\\\\\\\\\

A key-value store supporting various in-memory key-value data structures with different trade-offs. 
Porcupine ships with a http server and a cli client.

To start the http server, run 'porcupine server'.
You can then issue commands to the runtime in a different terminal session. 
View commands with 'porcupine --help'.
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
