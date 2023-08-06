package cmd

import (
	"fmt"
	"os"

	"github.com/ak9024/okr-generator/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "okr-generator",
	Version: config.Config().GetString("app.version"),
}

func Execute() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(envGenerator)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
