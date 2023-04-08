package cmd

import (
	"fmt"

	"github.com/hailelagi/porcupine-go/porcupine"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure the store",
	Long: `select the data structure for the store, experimenting with tradeoffs. 
	
	After starting the http server, you can configure the store by passing the config as a query param.
	Example: http://localhost:8080/?config=hashmap where it's stored in-memory.`,
	Run: func(cmd *cobra.Command, args []string) {
		var store *porcupine.Porcupine
		var stdout string

		// todo:

		if len(args) == 0 {
			store = porcupine.NewPorcupine("hashmap")
			stdout = fmt.Sprintf("using default data store: %s", store.Name)
		} else {
			store = porcupine.NewPorcupine(args[0])
			stdout = fmt.Sprintf("setup data store as: %s", store.Name)
		}

		fmt.Println(stdout)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
