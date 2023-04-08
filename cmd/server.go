package cmd

import (
	"fmt"

	"github.com/hailelagi/porcupine-go/api"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a server at localhost:8080",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting server at https://localhost:8080")
		// todo: expose routes or redirect to '/' with query params
		api.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
