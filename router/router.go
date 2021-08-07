package route

import (
	"net/http"

	"github.com/StayAlert/blogByF/backend/domains/healthcheck"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func Router() {
	router := fiber.New()

	group := router.Group("/check")
	group.Get("/health-check", healthcheck.HealthCheck)

	http.ListenAndServe(":4000", adaptor.FiberApp(router))

}
