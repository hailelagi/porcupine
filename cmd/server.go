package cmd

import (
	"fmt"

	"github.com/hailelagi/porcupine-go/api"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a server instance with a `Store`",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`\\\\\\\\\\\\\\\
\\ PORCUPINE \\
\\\\\\\\\\\\\\\
		`)
		fmt.Println("starting a server at https://localhost:8080")
		fmt.Println(`Please keep this terminal process alive and issue commands to the server, 
via a new terminal session in the cli or browser`)
		api.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
