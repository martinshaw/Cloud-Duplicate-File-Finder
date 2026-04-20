package server

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"martinshaw.co/cloudduplicatefilefinder/config"
)

type Server struct {
	App  *fiber.App
	Port string
}

func LoadServer(config *config.Config) *Server {
	app := fiber.New(fiber.Config{
		AppName: config.AppName,
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	go func() {
		if err := app.Listen(":"+config.Port, fiber.ListenConfig{
			EnablePrefork:         true,
			DisableStartupMessage: !config.Debug,
		}); err != nil {
			log.Fatalf("Listen: %v", err)
		}
	}()

	return &Server{App: app, Port: config.Port}
}

func (s *Server) Announce() {
	log.Printf("%s is running on port %s", s.App.Config().AppName, s.Port)
}
