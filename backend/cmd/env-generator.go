package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Env struct {
	Port               int
	Host               string
	Version            string
	Env                string
	Token              string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	GoogleAuthState    string
	SupabaseURL        string
	SupabaseKey        string
}

// Get the OS environment
func GetEnvironment() Env {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")
	version := os.Getenv("VERSION")
	env := os.Getenv("ENV")
	token := os.Getenv("TOKEN")
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleRedirectURL := os.Getenv("GOOGLE_REDIRECT_URL")
	googleAuthState := os.Getenv("GOOGLE_AUTH_STATE")
	supabaseURL := os.Getenv("SUPABASE_URL")
	supbaseKey := os.Getenv("SUPABSE_KEY")

	return Env{
		Port:               port,
		Host:               host,
		Version:            version,
		Env:                env,
		Token:              token,
		GoogleClientID:     googleClientID,
		GoogleClientSecret: googleClientSecret,
		GoogleRedirectURL:  googleRedirectURL,
		GoogleAuthState:    googleAuthState,
		SupabaseURL:        supabaseURL,
		SupabaseKey:        supbaseKey,
	}
}

func EnvGenerator(cmd *cobra.Command, args []string) {
	// Get OS environment
	env := GetEnvironment()

	// Compose the config
	cfg := map[string]interface{}{
		"app": map[string]interface{}{
			"port":    env.Port,
			"host":    env.Host,
			"version": env.Version,
			"env":     env.Env,
		},
		"chatgpt": map[string]interface{}{
			"token": env.Token,
		},
		"google": map[string]interface{}{
			"redirect_url":  env.GoogleRedirectURL,
			"client_id":     env.GoogleClientID,
			"client_secret": env.GoogleClientSecret,
			"auth_state":    env.GoogleAuthState,
		},
		"supabase": map[string]interface{}{
			"url": env.SupabaseURL,
			"key": env.SupabaseKey,
		},
	}

	// Convert the data to toml structure
	b, errCreate := toml.Marshal(cfg)
	if errCreate != nil {
		fmt.Println(errCreate)
	}

	// Create a file .config.generated.toml
	f, err := os.Create(".config.generated.toml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// Insert the data to file `.config.generated.toml`
	_, errWrite := f.Write(b)
	if errWrite != nil {
		fmt.Println(errWrite)
	}

	fmt.Println("Done!")
}

var envGenerator = &cobra.Command{
	Use: "env-generator",
	Run: EnvGenerator,
}
