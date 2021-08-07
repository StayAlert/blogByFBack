package route

import (
	"github.com/StayAlert/blogByF/backend/domains/healthcheck"
	"github.com/gofiber/fiber/v2"
)

func Router() *fiber.App {
	router := fiber.New()

	router.Get("health-check", healthcheck.HealthCheck)

	return router
}
