package router

import (
	"fmt"

	"github.com/ak9024/okr-generator/config"
	"github.com/ak9024/okr-generator/docs"
	"github.com/ak9024/okr-generator/internal/handler"
	"github.com/ak9024/okr-generator/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
)

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
func (s *server) New() fiber.Router {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", s.Config.GetString("app.host"), s.Config.GetInt("app.port"))
	docs.SwaggerInfo.Version = s.Config.GetString("app.version")

	app := fiber.New()

	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	service := service.NewService(s.Config)
	h := handler.NewHandler(service)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/okr-generator", h.OKRGeneratorHandler)

	port := fmt.Sprintf(":%d", s.Config.GetInt("app.port"))
	app.Listen(port)

	return app
}
