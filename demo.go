package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// demo handler
type DemoHandler struct {
	storage *Storage
}

func NewDemoHandler(lc fx.Lifecycle, storage *Storage) *DemoHandler {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// do something on start
			fmt.Println("DEMO HANDLER HAS BEEN STARTED!!")
			return nil
		},
	})
	return &DemoHandler{
		storage: storage,
	}
}

func (h *DemoHandler) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (h *DemoHandler) GetData(c *fiber.Ctx) error {
	return c.SendString(h.storage.GetData())
}
