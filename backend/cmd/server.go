package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ak9024/okr-generator/config"
	"github.com/ak9024/okr-generator/docs"
	"github.com/ak9024/okr-generator/internal/okr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		// init the configuration
		c := config.Config()
		// init the server
		s := NewServer(c)
		// start the server
		s.StartApp()
	},
}

type server struct {
	Config config.Provider
}

func NewServer(cfg config.Provider) *server {
	return &server{
		Config: cfg,
	}
}

// @title OKR Generator API
// @description This is Official API for OKR Generator API
func (s *server) StartApp() {
	// init swagger
	setupSwaggerInfo(s)

	app := fiber.New()

	// setup fiber middleware
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())

	// declare endpoint for swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// init okr config
	okr := okr.OKR{
		Config: s.Config,
	}

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// POST /api/v1/okr-generator
	v1.Post("/okr-generator", okr.OKRGeneratorHandler)

	getPort := fmt.Sprintf(":%d", s.Config.GetInt("app.port"))

	if s.Config.GetString("app.env") == "development" {
		if err := app.Listen(getPort); err != nil {
			logrus.Error(err)
		}
	}

	if s.Config.GetString("app.env") == "production" {
		go func() {
			if err := app.Listen(getPort); err != nil {
				logrus.Error(err)
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := app.ShutdownWithContext(ctx); err != nil {
			logrus.Fatal(err)
		}
	}
}

func setupSwaggerInfo(s *server) {
	// get hostname (host + port) get from file .config.toml
	getHostName := fmt.Sprintf("%s:%d", s.Config.GetString("app.host"), s.Config.GetInt("app.port"))

	// setup swagger info
	docs.SwaggerInfo.Host = getHostName
	docs.SwaggerInfo.Version = s.Config.GetString("app.version")
}
