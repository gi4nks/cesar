package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// init initializes the commands package by adding the versionCmd command to the RootCmd.
func init() {
	RootCmd.AddCommand(versionCmd)
}

// versionCmd represents the command to print the version number of Cesar.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Cesar",
	Long:  `All software has versions. This is Cesar's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
}
