package cmd

import (
	"os"
	"strconv"

	envgenerator "github.com/ak9024/okr-generator/cmd/env-generator"
	"github.com/spf13/cobra"
)

var envGenerator = &cobra.Command{
	Use: "env-generator",
	Run: func(cmd *cobra.Command, args []string) {
		// setup new config here
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		host := os.Getenv("HOST")
		version := os.Getenv("VERSION")
		env := os.Getenv("ENV")
		token := os.Getenv("TOKEN")
		googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
		googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
		googleRedirectURL := os.Getenv("GOOGLE_REDIRECT_URL")
		googleClientRedirectURL := os.Getenv("GOOGLE_CLIENT_REDIRECT_URL")
		supabaseURL := os.Getenv("SUPABASE_URL")
		supabaseKey := os.Getenv("SUPABASE_KEY")

		// init env generator
		eg := envgenerator.New(envgenerator.EnvGenerator{
			Port:                    port,
			Host:                    host,
			Version:                 version,
			Env:                     env,
			Token:                   token,
			GoogleClientID:          googleClientID,
			GoogleClientSecret:      googleClientSecret,
			GoogleRedirectURL:       googleRedirectURL,
			GoogleClientRedirectURL: googleClientRedirectURL,
			SupabaseURL:             supabaseURL,
			SupabaseKey:             supabaseKey,
		})

		// run env generator
		eg.Exec()
	},
}
