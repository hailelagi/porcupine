package cmd

import (
	"fmt"

	"github.com/hailelagi/porcupine-go/api"
	porcupine "github.com/hailelagi/porcupine-go/internal"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure the store",
	Long: `select the data structure for the store, experimenting with tradeoffs. 
	After starting the http server, you can configure the store in the UI. Example: http://localhost:8080/configure`,
	Run: func(cmd *cobra.Command, args []string) {
		var store porcupine.Porcupine
		var stdout string

		if len(args) == 0 {
			store = porcupine.NewPorcupine("hashmap")
			stdout = fmt.Sprintf("using default data store: %s", store.Name)
			api.Config(store.Name)
		} else {
			store = porcupine.NewPorcupine(args[0])
			api.Config(store.Name)
			stdout = fmt.Sprintf("setup data store as: %s", store.Name)
		}

		fmt.Println(stdout)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
