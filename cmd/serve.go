package cmd

import (
	"fmt"

	"github.com/hailelagi/porcupine-go/api"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts an instance with a `Store`",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`\\\\\\\\\\\\\\\
\\ PORCUPINE \\
\\\\\\\\\\\\\\\
		`)
		cmd.Println("starting a server at http://localhost:8080")
		cmd.Println(`Please keep this terminal process alive and issue commands to the server, 
via a new terminal session in the cli or browser`)
		api.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
