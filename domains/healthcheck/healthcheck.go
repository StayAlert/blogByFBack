package healthcheck

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(ctx *fiber.Ctx) error {
	fmt.Println("Health look OK...")
	return nil
}
