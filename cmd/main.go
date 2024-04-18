package main

import (
	"fmt"
	"os"

	"github.com/gi4nks/cesar/cmd/commands"
)

// main entry point for the executable. It starts cobra commands.
func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
