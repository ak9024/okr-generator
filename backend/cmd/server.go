package cmd

import (
	"github.com/ak9024/okr-generator/config"
	"github.com/ak9024/okr-generator/internal/router"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config()
		s := router.NewServer(c)
		s.New()
	},
}
