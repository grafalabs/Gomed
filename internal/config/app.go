package config

import (
	"gomed/internal/delivery/http/route"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	DB       *sqlx.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// Repository

	// Use cases

	// Controller

	// Middleware

	routeConfig := route.RouteConfig{
		App: config.App,
	}

	routeConfig.Setup()

}
