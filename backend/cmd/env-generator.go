package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Env struct {
	Port    int
	Host    string
	Version string
	Env     string
	Token   string
}

// Get the OS environment
func GetEnvironment() Env {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")
	version := os.Getenv("VERSION")
	env := os.Getenv("ENV")
	token := os.Getenv("TOKEN")

	return Env{
		Port:    port,
		Host:    host,
		Version: version,
		Env:     env,
		Token:   token,
	}
}

var envGenerator = &cobra.Command{
	Use: "env-generator",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}
