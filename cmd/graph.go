package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hailelagi/porcupine-go/pkg"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "draw a graph of the read/write perf of a data structure",
	Long:  `"draw a graph of the read/write perf of a data structure"`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.PlotExample()
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}
