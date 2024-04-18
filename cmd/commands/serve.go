package commands

import (
	"github.com/gi4nks/cesar/internal"

	"github.com/spf13/cobra"
)

// init initializes the commands package by adding the serveCmd command to the root command.
func init() {
	RootCmd.AddCommand(serveCmd)
}

// serveCmd represents the command to start the server for serving Cesar's APIs.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the Cesar's APIs",
	Long:  `Serve the Cesar's APIs`,
	Run: func(cmd *cobra.Command, args []string) {
		handler := internal.NewHandler()
		handler.StartServer()
	},
}
