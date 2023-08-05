package cmd

import (
	"github.com/ak9024/okr-generator/cmd/server"
	"github.com/ak9024/okr-generator/config"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		// init the configuration
		c := config.Config()
		// init the server
		s := server.NewServer(c)
		// start the router
		s.Router()
		// start the server
		s.StartApp()
	},
}
