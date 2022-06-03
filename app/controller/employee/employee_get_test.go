package employee

import (
	"api/app/lib"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetEmploye(t *testing.T) {
	services.InitDatabaseForTest()
	app := fiber.New()
	app.Get("/api/v1/employee", GetEmployee)

	res, _, err := lib.GetTest(app, "/api/v1/employee", nil)
	utils.AssertEqual(t, nil, err, "GET /api/v1/employee")
	utils.AssertEqual(t, 200, res.StatusCode, "HTTP Status")
}
