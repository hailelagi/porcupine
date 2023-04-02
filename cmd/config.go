/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hailelagi/porcupine-go/porcupine"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure the store",
	Long: `select the data structure for the store, experimenting with tradeoffs. 
	
	If you use the http server, you can configure the store by passing the config as a query param.
	Example: http://localhost:8080/?config=hashmap where it's stored in-memory of the main go routine.
	
	Whereas if you use the cli, config options are persisted in a .txt file in the root directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		var store *porcupine.Porcupine

		if len(args) == 0 {
			store = porcupine.NewPorcupine("hashmap")
		} else {
			store = porcupine.NewPorcupine(args[0])
		}

		fmt.Printf("setup data store as %s", store.Name)
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
