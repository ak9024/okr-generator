package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ak9024/okr-generator/config"
	"github.com/ak9024/okr-generator/docs"
	"github.com/ak9024/okr-generator/internal/auth"
	"github.com/ak9024/okr-generator/internal/lib"
	"github.com/ak9024/okr-generator/internal/okr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
)

type server struct {
	Config config.Provider
}

func NewServer(cfg config.Provider) *server {
	return &server{
		Config: cfg,
	}
}

func (s *server) Router() *fiber.App {
	// init swagger
	setupSwaggerInfo(s)

	// init new fiber app
	app := fiber.New()

	// setup fiber middleware
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New())

	// declare endpoint for swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)
	// monitor
	app.Get("/metrics", monitor.New(monitor.Config{
		Title: "OKR Generator Metrics Page",
	}))

	// group /api/
	api := app.Group("/api")
	// init auth
	auth := auth.NewAuth(s.Config)
	// GET /api/auth/google/login
	api.Get("/auth/google/login", auth.GoogleLoginHandler)
	// GET /api/auth/google/callback
	api.Get("/auth/google/callback", auth.GoogleLoginCallback)
	// GET /api/auth/google/logout
	api.Get("/auth/google/logout", auth.GoogleLogoutHandler)

	// group /v1/
	v1 := api.Group("/v1")
	// prevent access to v1 with AuthMiddleware
	// need to add header Authorization <token>
	v1.Use(lib.AuthMiddleware)

	// init okr config
	okr := okr.NewOKR(s.Config)
	// POST /api/v1/okr-generator
	v1.Post("/okr-generator", okr.OKRGeneratorHandler)

	return app
}

func (s *server) StartApp() {
	app := s.Router()

	// Get PORT
	getPort := fmt.Sprintf(":%d", s.Config.GetInt("app.port"))

	// if run in mode development
	if s.Config.GetString("app.env") == "development" {
		if err := app.Listen(getPort); err != nil {
			logrus.Error(err)
		}
	}

	// if run in mode production
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

// setup swagger info
func setupSwaggerInfo(s *server) {
	// get hostname (host + port) get from file .config.toml
	getHostName := fmt.Sprintf("%s:%d", s.Config.GetString("app.host"), s.Config.GetInt("app.port"))

	// setup swagger info
	docs.SwaggerInfo.Host = getHostName
	docs.SwaggerInfo.Version = s.Config.GetString("app.version")
}
