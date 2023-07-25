package router

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ak9024/okr-generator/config"
	"github.com/ak9024/okr-generator/docs"
	"github.com/ak9024/okr-generator/internal/handler"
	"github.com/ak9024/okr-generator/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
)

var (
	fiberApp *fiber.App
	once     sync.Once
)

// single instance of the fiber.App
func getFiberApp() *fiber.App {
	once.Do(func() {
		fiberApp = fiber.New()
	})

	return fiberApp
}

type server struct {
	Config config.Provider
}

func NewServer(cfg config.Provider) *server {
	return &server{
		Config: cfg,
	}
}

func setupSwaggerInfo(s *server) {
	// get hostname (host + port) get from file .config.toml
	getHostName := fmt.Sprintf("%s:%d", s.Config.GetString("app.host"), s.Config.GetInt("app.port"))

	// setup swagger info
	docs.SwaggerInfo.Host = getHostName
	docs.SwaggerInfo.Version = s.Config.GetString("app.version")
}

// @title OKR Generator API
// @description This is Official API for OKR Generator API
func (s *server) New() {
	setupSwaggerInfo(s)

	app := getFiberApp()

	// setup fiber middleware
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())

	// declare endpoint for swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// init handler and service
	service := service.NewService(s.Config)
	h := handler.NewHandler(service)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// POST /api/v1/okr-generator
	v1.Post("/okr-generator", h.OKRGeneratorHandler)

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
