package route

import "github.com/gofiber/fiber/v2"

type RouteConfig struct {
	App *fiber.App
}

func (c *RouteConfig) Setup() {
	c.SetupDefaultRoute()
}

func (c *RouteConfig) SetupDefaultRoute() {
	c.App.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("API is alive.")
	})

}

