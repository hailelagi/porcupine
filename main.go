package main

import "github.com/hailelagi/porcupine-go/cmd"

// "github.com/hailelagi/porcupine-go/cmd"
//"github.com/hailelagi/porcupine-go/porcupine"

// maybe refactor to only use flag?
// https://abhinavg.net/2022/08/13/flag-subcommand/
func main() {
	cmd.Execute()
	// porcupine.PrintFineMap()
}
