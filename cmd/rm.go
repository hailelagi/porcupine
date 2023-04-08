package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove an entry by key",
	Long:  `Remove the value associated with the key, performance will vary depending on the data structure.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("rm called")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
