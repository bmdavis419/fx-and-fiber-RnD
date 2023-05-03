package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// server creator
func newHTTPServer(lc fx.Lifecycle, demoHandler *DemoHandler) *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/demo", demoHandler.Hello)

	app.Get("/demo/storage", demoHandler.GetData)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting HTTP server on 8080")
			go app.Listen(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}

func main() {
	fx.New(
		fx.Provide(NewStorage, NewDemoHandler),
		fx.Invoke(newHTTPServer),
	).Run()
}
