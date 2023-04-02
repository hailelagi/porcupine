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

A key-value store supporting several data structures with different trade-offs. 
Porcupine ships with an optional http server and a cli client.

To start the http server, run 'porcupine server'.
Or interact with the cli directly 'porcupine --help'
`,
	// todo select default data structe
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.porcupine-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
